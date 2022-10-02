package model121

// Code generated by ../../generators/models.go. DO NOT EDIT.
// You can regenerate it by running 'go generate ./models' from the directory above.

import (
	"github.com/andig/gosunspec"
	"github.com/andig/gosunspec/smdx"
	"github.com/andig/gosunspec/typelabel"
)

// Block121 - Basic Settings - Inverter Controls Basic Settings

const (
	ModelID          = 121
	ModelLabel       = "Basic Settings"
	ModelDescription = "Inverter Controls Basic Settings "
)

const (
	ClcTotVA     = "ClcTotVA"
	ConnPh       = "ConnPh"
	ECPNomHz     = "ECPNomHz"
	ECPNomHz_SF  = "ECPNomHz_SF"
	MaxRmpRte    = "MaxRmpRte"
	MaxRmpRte_SF = "MaxRmpRte_SF"
	PFMinQ1      = "PFMinQ1"
	PFMinQ2      = "PFMinQ2"
	PFMinQ3      = "PFMinQ3"
	PFMinQ4      = "PFMinQ4"
	PFMin_SF     = "PFMin_SF"
	VAMax        = "VAMax"
	VAMax_SF     = "VAMax_SF"
	VArAct       = "VArAct"
	VArMaxQ1     = "VArMaxQ1"
	VArMaxQ2     = "VArMaxQ2"
	VArMaxQ3     = "VArMaxQ3"
	VArMaxQ4     = "VArMaxQ4"
	VArMax_SF    = "VArMax_SF"
	VMax         = "VMax"
	VMin         = "VMin"
	VMinMax_SF   = "VMinMax_SF"
	VRef         = "VRef"
	VRefOfs      = "VRefOfs"
	VRefOfs_SF   = "VRefOfs_SF"
	VRef_SF      = "VRef_SF"
	WGra         = "WGra"
	WGra_SF      = "WGra_SF"
	WMax         = "WMax"
	WMax_SF      = "WMax_SF"
)

type Block121 struct {
	WMax         uint16              `sunspec:"offset=0,len=1,sf=WMax_SF,access=rw"`
	VRef         uint16              `sunspec:"offset=1,len=1,sf=VRef_SF,access=rw"`
	VRefOfs      int16               `sunspec:"offset=2,len=1,sf=VRefOfs_SF,access=rw"`
	VMax         uint16              `sunspec:"offset=3,len=1,sf=VMinMax_SF,access=rw"`
	VMin         uint16              `sunspec:"offset=4,len=1,sf=VMinMax_SF,access=rw"`
	VAMax        uint16              `sunspec:"offset=5,len=1,sf=VAMax_SF,access=rw"`
	VArMaxQ1     int16               `sunspec:"offset=6,len=1,sf=VArMax_SF,access=rw"`
	VArMaxQ2     int16               `sunspec:"offset=7,len=1,sf=VArMax_SF,access=rw"`
	VArMaxQ3     int16               `sunspec:"offset=8,len=1,sf=VArMax_SF,access=rw"`
	VArMaxQ4     int16               `sunspec:"offset=9,len=1,sf=VArMax_SF,access=rw"`
	WGra         uint16              `sunspec:"offset=10,len=1,sf=WGra_SF,access=rw"`
	PFMinQ1      int16               `sunspec:"offset=11,len=1,sf=PFMin_SF,access=rw"`
	PFMinQ2      int16               `sunspec:"offset=12,len=1,sf=PFMin_SF,access=rw"`
	PFMinQ3      int16               `sunspec:"offset=13,len=1,sf=PFMin_SF,access=rw"`
	PFMinQ4      int16               `sunspec:"offset=14,len=1,sf=PFMin_SF,access=rw"`
	VArAct       sunspec.Enum16      `sunspec:"offset=15,len=1,access=rw"`
	ClcTotVA     sunspec.Enum16      `sunspec:"offset=16,len=1,access=rw"`
	MaxRmpRte    uint16              `sunspec:"offset=17,len=1,sf=MaxRmpRte_SF,access=rw"`
	ECPNomHz     uint16              `sunspec:"offset=18,len=1,sf=ECPNomHz_SF,access=rw"`
	ConnPh       sunspec.Enum16      `sunspec:"offset=19,len=1,access=rw"`
	WMax_SF      sunspec.ScaleFactor `sunspec:"offset=20,len=1,access=r"`
	VRef_SF      sunspec.ScaleFactor `sunspec:"offset=21,len=1,access=r"`
	VRefOfs_SF   sunspec.ScaleFactor `sunspec:"offset=22,len=1,access=r"`
	VMinMax_SF   sunspec.ScaleFactor `sunspec:"offset=23,len=1,access=r"`
	VAMax_SF     sunspec.ScaleFactor `sunspec:"offset=24,len=1,access=r"`
	VArMax_SF    sunspec.ScaleFactor `sunspec:"offset=25,len=1,access=r"`
	WGra_SF      sunspec.ScaleFactor `sunspec:"offset=26,len=1,access=r"`
	PFMin_SF     sunspec.ScaleFactor `sunspec:"offset=27,len=1,access=r"`
	MaxRmpRte_SF sunspec.ScaleFactor `sunspec:"offset=28,len=1,access=r"`
	ECPNomHz_SF  sunspec.ScaleFactor `sunspec:"offset=29,len=1,access=r"`
}

func (block *Block121) GetId() sunspec.ModelId {
	return ModelID
}

func init() {
	smdx.RegisterModel(&smdx.ModelElement{
		Id:     ModelID,
		Name:   "settings",
		Length: 30,
		Blocks: []smdx.BlockElement{
			{
				Length: 30,
				Type:   "fixed",
				Points: []smdx.PointElement{
					{Id: WMax, Offset: 0, Type: typelabel.Uint16, ScaleFactor: "WMax_SF", Units: "W", Access: "rw", Length: 1, Mandatory: true, Label: "WMax", Description: "Setting for maximum power output. Default to WRtg."},
					{Id: VRef, Offset: 1, Type: typelabel.Uint16, ScaleFactor: "VRef_SF", Units: "V", Access: "rw", Length: 1, Mandatory: true, Label: "VRef", Description: "Voltage at the PCC."},
					{Id: VRefOfs, Offset: 2, Type: typelabel.Int16, ScaleFactor: "VRefOfs_SF", Units: "V", Access: "rw", Length: 1, Mandatory: true, Label: "VRefOfs", Description: "Offset  from PCC to inverter."},
					{Id: VMax, Offset: 3, Type: typelabel.Uint16, ScaleFactor: "VMinMax_SF", Units: "V", Access: "rw", Length: 1, Label: "VMax", Description: "Setpoint for maximum voltage."},
					{Id: VMin, Offset: 4, Type: typelabel.Uint16, ScaleFactor: "VMinMax_SF", Units: "V", Access: "rw", Length: 1, Label: "VMin", Description: "Setpoint for minimum voltage."},
					{Id: VAMax, Offset: 5, Type: typelabel.Uint16, ScaleFactor: "VAMax_SF", Units: "VA", Access: "rw", Length: 1, Label: "VAMax", Description: "Setpoint for maximum apparent power. Default to VARtg."},
					{Id: VArMaxQ1, Offset: 6, Type: typelabel.Int16, ScaleFactor: "VArMax_SF", Units: "var", Access: "rw", Length: 1, Label: "VArMaxQ1", Description: "Setting for maximum reactive power in quadrant 1. Default to VArRtgQ1."},
					{Id: VArMaxQ2, Offset: 7, Type: typelabel.Int16, ScaleFactor: "VArMax_SF", Units: "var", Access: "rw", Length: 1, Label: "VArMaxQ2", Description: "Setting for maximum reactive power in quadrant 2. Default to VArRtgQ2."},
					{Id: VArMaxQ3, Offset: 8, Type: typelabel.Int16, ScaleFactor: "VArMax_SF", Units: "var", Access: "rw", Length: 1, Label: "VArMaxQ3", Description: "Setting for maximum reactive power in quadrant 3. Default to VArRtgQ3."},
					{Id: VArMaxQ4, Offset: 9, Type: typelabel.Int16, ScaleFactor: "VArMax_SF", Units: "var", Access: "rw", Length: 1, Label: "VArMaxQ4", Description: "Setting for maximum reactive power in quadrant 4. Default to VArRtgQ4."},
					{Id: WGra, Offset: 10, Type: typelabel.Uint16, ScaleFactor: "WGra_SF", Units: "% WMax/sec", Access: "rw", Length: 1, Label: "WGra", Description: "Default ramp rate of change of active power due to command or internal action."},
					{Id: PFMinQ1, Offset: 11, Type: typelabel.Int16, ScaleFactor: "PFMin_SF", Units: "cos()", Access: "rw", Length: 1, Label: "PFMinQ1", Description: "Setpoint for minimum power factor value in quadrant 1. Default to PFRtgQ1."},
					{Id: PFMinQ2, Offset: 12, Type: typelabel.Int16, ScaleFactor: "PFMin_SF", Units: "cos()", Access: "rw", Length: 1, Label: "PFMinQ2", Description: "Setpoint for minimum power factor value in quadrant 2. Default to PFRtgQ2."},
					{Id: PFMinQ3, Offset: 13, Type: typelabel.Int16, ScaleFactor: "PFMin_SF", Units: "cos()", Access: "rw", Length: 1, Label: "PFMinQ3", Description: "Setpoint for minimum power factor value in quadrant 3. Default to PFRtgQ3."},
					{Id: PFMinQ4, Offset: 14, Type: typelabel.Int16, ScaleFactor: "PFMin_SF", Units: "cos()", Access: "rw", Length: 1, Label: "PFMinQ4", Description: "Setpoint for minimum power factor value in quadrant 4. Default to PFRtgQ4."},
					{Id: VArAct, Offset: 15, Type: typelabel.Enum16, Access: "rw", Length: 1, Label: "VArAct", Description: "VAR action on change between charging and discharging: 1=switch 2=maintain VAR characterization."},
					{Id: ClcTotVA, Offset: 16, Type: typelabel.Enum16, Access: "rw", Length: 1, Label: "ClcTotVA", Description: "Calculation method for total apparent power. 1=vector 2=arithmetic."},
					{Id: MaxRmpRte, Offset: 17, Type: typelabel.Uint16, ScaleFactor: "MaxRmpRte_SF", Units: "% WGra", Access: "rw", Length: 1, Label: "MaxRmpRte", Description: "Setpoint for maximum ramp rate as percentage of nominal maximum ramp rate. This setting will limit the rate that watts delivery to the grid can increase or decrease in response to intermittent PV generation."},
					{Id: ECPNomHz, Offset: 18, Type: typelabel.Uint16, ScaleFactor: "ECPNomHz_SF", Units: "Hz", Access: "rw", Length: 1, Label: "ECPNomHz", Description: "Setpoint for nominal frequency at the ECP."},
					{Id: ConnPh, Offset: 19, Type: typelabel.Enum16, Access: "rw", Length: 1, Label: "ConnPh", Description: "Identity of connected phase for single phase inverters. A=1 B=2 C=3."},
					{Id: WMax_SF, Offset: 20, Type: typelabel.ScaleFactor, Access: "r", Length: 1, Mandatory: true, Label: "WMax_SF", Description: "Scale factor for real power."},
					{Id: VRef_SF, Offset: 21, Type: typelabel.ScaleFactor, Access: "r", Length: 1, Mandatory: true, Label: "VRef_SF", Description: "Scale factor for voltage at the PCC."},
					{Id: VRefOfs_SF, Offset: 22, Type: typelabel.ScaleFactor, Access: "r", Length: 1, Mandatory: true, Label: "VRefOfs_SF", Description: "Scale factor for offset voltage."},
					{Id: VMinMax_SF, Offset: 23, Type: typelabel.ScaleFactor, Access: "r", Length: 1, Label: "VMinMax_SF", Description: "Scale factor for min/max voltages."},
					{Id: VAMax_SF, Offset: 24, Type: typelabel.ScaleFactor, Access: "r", Length: 1, Label: "VAMax_SF", Description: "Scale factor for apparent power."},
					{Id: VArMax_SF, Offset: 25, Type: typelabel.ScaleFactor, Access: "r", Length: 1, Label: "VArMax_SF", Description: "Scale factor for reactive power."},
					{Id: WGra_SF, Offset: 26, Type: typelabel.ScaleFactor, Access: "r", Length: 1, Label: "WGra_SF", Description: "Scale factor for default ramp rate."},
					{Id: PFMin_SF, Offset: 27, Type: typelabel.ScaleFactor, Access: "r", Length: 1, Label: "PFMin_SF", Description: "Scale factor for minimum power factor."},
					{Id: MaxRmpRte_SF, Offset: 28, Type: typelabel.ScaleFactor, Access: "r", Length: 1, Label: "MaxRmpRte_SF", Description: "Scale factor for maximum ramp percentage."},
					{Id: ECPNomHz_SF, Offset: 29, Type: typelabel.ScaleFactor, Access: "r", Length: 1, Label: "ECPNomHz_SF", Description: "Scale factor for nominal frequency."},
				},
			},
		}})
}
