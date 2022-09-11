package ocpp

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/evcc-io/evcc/api"
	"github.com/evcc-io/evcc/util"
	"github.com/lorenzodonini/ocpp-go/ocpp1.6/core"
	"github.com/lorenzodonini/ocpp-go/ocpp1.6/types"
)

const timeout = 2 * time.Minute

// Meter Profile Key
const (
	KeyMeterValuesSampledData   = "MeterValuesSampledData"
	KeyMeterValueSampleInterval = "MeterValueSampleInterval"
)

type SmartchargingChargeProfileKey string

// Smart Charging Profile Key
const (
	KeyChargeProfileMaxStackLevel              SmartchargingChargeProfileKey = "ChargeProfileMaxStackLevel"
	KeyChargingScheduleAllowedChargingRateUnit SmartchargingChargeProfileKey = "ChargingScheduleAllowedChargingRateUnit"
	KeyChargingScheduleMaxPeriods              SmartchargingChargeProfileKey = "ChargingScheduleMaxPeriods"
	KeyConnectorSwitch3to1PhaseSupported       SmartchargingChargeProfileKey = "ConnectorSwitch3to1PhaseSupported"
	KeyMaxChargingProfilesInstalled            SmartchargingChargeProfileKey = "MaxChargingProfilesInstalled"
)

type smartChargingProfile struct {
	// Max StackLevel of a ChargingProfile. The number defined also indicates the max allowed
	// number of installed charging scheduls per Charging Profile Purpose
	ChargeProfileMaxStackLevel int
	// A list of supported quantities for use in a ChargingSchedule.
	// Allowed values: 'Current' and 'Power'
	ChargingScheduleAllowedChargingRateUnit []string
	// Maximum number of periods that may be defined per ChargingSchedule
	ChargingScheduleMaxPeriods int
	// Defines if this Charge Point support switching from 3 to 1 phase during a charging session.
	ConnectorSwitch3to1PhaseSupported bool
	// Maximum number of Charging profiles instsalled at a time.
	MaxChargingProfilesInstalled int
}

type CP struct {
	mu  sync.Mutex
	log *util.Logger
	id  string

	updated     time.Time
	initialized *sync.Cond
	boot        *core.BootNotificationRequest
	status      *core.StatusNotificationRequest

	timeout      time.Duration
	meterUpdated time.Time
	measurements map[string]types.SampledValue

	supportedNumberOfConnectors int
	smartChargingCapabilities   smartChargingProfile

	txnCount int // change initial value to the last known global transaction. Needs persistence
	txnId    int
}

func (cp *CP) WatchDog(timeout time.Duration) {
	cp.timeout = timeout

	go func() {
		for ; true; <-time.NewTicker(timeout).C {
			cp.mu.Lock()
			update := cp.txnId != 0 && time.Since(cp.meterUpdated) > timeout
			cp.mu.Unlock()

			if update {
				Instance().TriggerMeterValuesRequest(cp)
			}
		}
	}()
}

func (cp *CP) DetectCapabilities(opts []core.ConfigurationKey) error {
	options := make(map[string]core.ConfigurationKey)
	for _, opt := range opts {
		options[opt.Key] = opt
	}

	var err error
	if cp.supportedNumberOfConnectors, err = parseIntOption("NumberOfConnectors", options); err != nil {
		return err
	}

	if cp.smartChargingCapabilities, err = detectSmartChargingCapabilities(options); err != nil {
		return err
	}

	return nil
}

func (cp *CP) GetNumberOfSupportedConnectors() int {
	return cp.supportedNumberOfConnectors
}

func detectSmartChargingCapabilities(options map[string]core.ConfigurationKey) (smartChargingProfile, error) {
	var profile smartChargingProfile

	{ // required
		val, err := parseIntOption(KeyChargeProfileMaxStackLevel, options)
		if err != nil {
			return profile, err
		}

		profile.ChargeProfileMaxStackLevel = val
	}

	{ // required
		val, err := parseIntOption(KeyChargingScheduleMaxPeriods, options)
		if err != nil {
			return profile, err
		}

		profile.ChargingScheduleMaxPeriods = val
	}

	{ // required
		val, err := parseIntOption(KeyMaxChargingProfilesInstalled, options)
		if err != nil {
			return profile, err
		}

		profile.MaxChargingProfilesInstalled = val
	}

	{ // required
		opt, found := options[string(KeyChargingScheduleAllowedChargingRateUnit)]
		if !found || opt.Value == nil {
			return profile, fmt.Errorf("smart charging key '%s' not found", KeyChargingScheduleAllowedChargingRateUnit)
		}

		vals := strings.Split(*opt.Value, ",")
		profile.ChargingScheduleAllowedChargingRateUnit = append(profile.ChargingScheduleAllowedChargingRateUnit, vals...)
	}

	{ // optional
		var supported bool
		opt, found := options[string(KeyConnectorSwitch3to1PhaseSupported)]
		if found {
			var err error
			supported, err = strconv.ParseBool(*opt.Value)
			if err != nil {
				return profile, fmt.Errorf("invalid value for key: %s", opt.Key)
			}
		}

		profile.ConnectorSwitch3to1PhaseSupported = supported
	}

	return profile, nil
}

func parseIntOption(key SmartchargingChargeProfileKey, options map[string]core.ConfigurationKey) (int, error) {
	opt, found := options[string(key)]
	if !found || opt.Value == nil {
		return 0, fmt.Errorf("smart charging key '%s' not found", key)
	}

	val, err := strconv.Atoi(*opt.Value)
	if err != nil {
		return 0, fmt.Errorf("invalid value for key: %s", key)
	}

	return val, nil
}

// Boot waits for the CP to register itself
func (cp *CP) Boot() error {
	bootC := make(chan struct{})
	go func() {
		cp.mu.Lock()
		defer cp.mu.Unlock()

		for cp.boot == nil || cp.status == nil {
			cp.initialized.Wait()
		}

		close(bootC)
	}()

	select {
	case <-bootC:
		cp.update()
		return nil
	case <-time.After(timeout):
		return api.ErrTimeout
	}
}

// TransactionID returns the current transaction id
func (cp *CP) TransactionID() int {
	cp.mu.Lock()
	defer cp.mu.Unlock()
	return cp.txnId
}

func (cp *CP) Status() (api.ChargeStatus, error) {
	cp.mu.Lock()
	defer cp.mu.Unlock()

	res := api.StatusNone

	cp.log.TRACE.Printf("current transaction ID (last update): %d (%s)", cp.txnId, cp.updated.Format(time.RFC3339))

	if time.Since(cp.updated) > timeout {
		return res, api.ErrTimeout
	}

	if cp.status.ErrorCode != core.NoError {
		cp.log.DEBUG.Printf("chargepoint error: %s: %s", cp.status.ErrorCode, cp.status.Info)
	}

	switch cp.status.Status {
	case core.ChargePointStatusAvailable, // "Available"
		core.ChargePointStatusUnavailable: // "Unavailable"
		res = api.StatusA
	case
		core.ChargePointStatusPreparing,     // "Preparing"
		core.ChargePointStatusSuspendedEVSE, // "SuspendedEVSE"
		core.ChargePointStatusSuspendedEV,   // "SuspendedEV"
		core.ChargePointStatusFinishing:     // "Finishing"
		res = api.StatusB
	case core.ChargePointStatusCharging: // "Charging"
		res = api.StatusC
	case core.ChargePointStatusReserved, // "Reserved"
		core.ChargePointStatusFaulted: // "Faulted"
		return api.StatusF, fmt.Errorf("chargepoint status: %s", cp.status.ErrorCode)
	default:
		return api.StatusNone, fmt.Errorf("invalid chargepoint status: %s", cp.status.Status)
	}

	return res, nil
}

var _ api.Meter = (*CP)(nil)

func (cp *CP) CurrentPower() (float64, error) {
	cp.mu.Lock()
	defer cp.mu.Unlock()

	if cp.timeout > 0 && time.Since(cp.meterUpdated) > cp.timeout {
		return 0, api.ErrNotAvailable
	}

	if power, ok := cp.measurements[string(types.MeasurandPowerActiveImport)]; ok {
		return strconv.ParseFloat(power.Value, 64)
	}

	return 0, api.ErrNotAvailable
}

var _ api.MeterEnergy = (*CP)(nil)

func (cp *CP) TotalEnergy() (float64, error) {
	cp.mu.Lock()
	defer cp.mu.Unlock()

	if cp.timeout > 0 && time.Since(cp.meterUpdated) > cp.timeout {
		return 0, api.ErrNotAvailable
	}

	if power, ok := cp.measurements[string(types.MeasurandEnergyActiveImportRegister)]; ok {
		f, err := strconv.ParseFloat(power.Value, 64)
		return f / 1e3, err
	}

	return 0, api.ErrNotAvailable
}

func getKeyCurrentPhase(phase int) string {
	return string(types.MeasurandCurrentImport) + "@L" + strconv.Itoa(phase)
}

var _ api.MeterCurrent = (*CP)(nil)

func (cp *CP) Currents() (float64, float64, float64, error) {
	cp.mu.Lock()
	defer cp.mu.Unlock()

	if cp.timeout > 0 && time.Since(cp.meterUpdated) > cp.timeout {
		return 0, 0, 0, api.ErrNotAvailable
	}

	currents := make([]float64, 0, 3)

	for phase := 1; phase <= 3; phase++ {
		current, ok := cp.measurements[getKeyCurrentPhase(phase)]
		if !ok {
			return 0, 0, 0, api.ErrNotAvailable
		}

		f, err := strconv.ParseFloat(current.Value, 64)
		if err != nil {
			return 0, 0, 0, fmt.Errorf("invalid current for phase %d: %w", phase, err)
		}

		currents = append(currents, f)
	}

	return currents[0], currents[1], currents[2], nil
}
