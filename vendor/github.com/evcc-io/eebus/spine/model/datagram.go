// Package ns_p contains models for http://docs.eebus.org/spine/xsd/v1
package model

// Code generated by github.com/andig/xsd2go. DO NOT EDIT.

import "github.com/evcc-io/eebus/util"

// CmiDatagramType message container
type CmiDatagramType struct {
	Datagram DatagramType `json:"datagram"`
}

// DatagramType complex type
type DatagramType struct {
	Header  HeaderType  `json:"header"`
	Payload PayloadType `json:"payload"`
}

// MarshalJSON is the SHIP serialization marshaller
func (m DatagramType) MarshalJSON() ([]byte, error) {
	return util.Marshal(m)
}

// UnmarshalJSON is the SHIP serialization unmarshaller
func (m *DatagramType) UnmarshalJSON(data []byte) error {
	return util.Unmarshal(data, &m)
}

// HeaderType complex type
type HeaderType struct {
	SpecificationVersion *SpecificationVersionType `json:"specificationVersion,omitempty"`
	AddressSource        *FeatureAddressType       `json:"addressSource,omitempty"`
	AddressDestination   *FeatureAddressType       `json:"addressDestination,omitempty"`
	AddressOriginator    *FeatureAddressType       `json:"addressOriginator,omitempty"`
	MsgCounter           *MsgCounterType           `json:"msgCounter,omitempty"`
	MsgCounterReference  *MsgCounterType           `json:"msgCounterReference,omitempty"`
	CmdClassifier        *CmdClassifierType        `json:"cmdClassifier,omitempty"`
	AckRequest           *bool                     `json:"ackRequest,omitempty"`
	Timestamp            *string                   `json:"timestamp,omitempty"`
}

// MarshalJSON is the SHIP serialization marshaller
func (m HeaderType) MarshalJSON() ([]byte, error) {
	return util.Marshal(m)
}

// UnmarshalJSON is the SHIP serialization unmarshaller
func (m *HeaderType) UnmarshalJSON(data []byte) error {
	return util.Unmarshal(data, &m)
}

// PayloadType complex type
type PayloadType struct {
	Cmd []CmdType `json:"cmd"`
}

// MarshalJSON is the SHIP serialization marshaller
func (m PayloadType) MarshalJSON() ([]byte, error) {
	return util.Marshal(m)
}

// UnmarshalJSON is the SHIP serialization unmarshaller
func (m *PayloadType) UnmarshalJSON(data []byte) error {
	return util.Unmarshal(data, &m)
}