// Package ns_p contains models for http://docs.eebus.org/spine/xsd/v1
package model

// Code generated by github.com/andig/xsd2go. DO NOT EDIT.

import "github.com/evcc-io/eebus/util"

// SubscriptionIdType type
type SubscriptionIdType uint

// SubscriptionManagementEntryDataType complex type
type SubscriptionManagementEntryDataType struct {
	SubscriptionId *SubscriptionIdType `json:"subscriptionId,omitempty"`
	ClientAddress  *FeatureAddressType `json:"clientAddress,omitempty"`
	ServerAddress  *FeatureAddressType `json:"serverAddress,omitempty"`
	Label          *LabelType          `json:"label,omitempty"`
	Description    *DescriptionType    `json:"description,omitempty"`
}

// MarshalJSON is the SHIP serialization marshaller
func (m SubscriptionManagementEntryDataType) MarshalJSON() ([]byte, error) {
	return util.Marshal(m)
}

// UnmarshalJSON is the SHIP serialization unmarshaller
func (m *SubscriptionManagementEntryDataType) UnmarshalJSON(data []byte) error {
	return util.Unmarshal(data, &m)
}

// SubscriptionManagementEntryListDataType complex type
type SubscriptionManagementEntryListDataType struct {
	SubscriptionManagementEntryData []SubscriptionManagementEntryDataType `json:"subscriptionManagementEntryData,omitempty"`
}

// MarshalJSON is the SHIP serialization marshaller
func (m SubscriptionManagementEntryListDataType) MarshalJSON() ([]byte, error) {
	return util.Marshal(m)
}

// UnmarshalJSON is the SHIP serialization unmarshaller
func (m *SubscriptionManagementEntryListDataType) UnmarshalJSON(data []byte) error {
	return util.Unmarshal(data, &m)
}

// SubscriptionManagementEntryListDataSelectorsType complex type
type SubscriptionManagementEntryListDataSelectorsType struct {
	SubscriptionId *SubscriptionIdType `json:"subscriptionId,omitempty"`
	ClientAddress  *FeatureAddressType `json:"clientAddress,omitempty"`
	ServerAddress  *FeatureAddressType `json:"serverAddress,omitempty"`
}

// MarshalJSON is the SHIP serialization marshaller
func (m SubscriptionManagementEntryListDataSelectorsType) MarshalJSON() ([]byte, error) {
	return util.Marshal(m)
}

// UnmarshalJSON is the SHIP serialization unmarshaller
func (m *SubscriptionManagementEntryListDataSelectorsType) UnmarshalJSON(data []byte) error {
	return util.Unmarshal(data, &m)
}

// SubscriptionManagementRequestCallType complex type
type SubscriptionManagementRequestCallType struct {
	ClientAddress     *FeatureAddressType `json:"clientAddress,omitempty"`
	ServerAddress     *FeatureAddressType `json:"serverAddress,omitempty"`
	ServerFeatureType *FeatureTypeType    `json:"serverFeatureType,omitempty"`
}

// MarshalJSON is the SHIP serialization marshaller
func (m SubscriptionManagementRequestCallType) MarshalJSON() ([]byte, error) {
	return util.Marshal(m)
}

// UnmarshalJSON is the SHIP serialization unmarshaller
func (m *SubscriptionManagementRequestCallType) UnmarshalJSON(data []byte) error {
	return util.Unmarshal(data, &m)
}

// SubscriptionManagementDeleteCallType complex type
type SubscriptionManagementDeleteCallType struct {
	SubscriptionId *SubscriptionIdType `json:"subscriptionId,omitempty"`
	ClientAddress  *FeatureAddressType `json:"clientAddress,omitempty"`
	ServerAddress  *FeatureAddressType `json:"serverAddress,omitempty"`
}

// MarshalJSON is the SHIP serialization marshaller
func (m SubscriptionManagementDeleteCallType) MarshalJSON() ([]byte, error) {
	return util.Marshal(m)
}

// UnmarshalJSON is the SHIP serialization unmarshaller
func (m *SubscriptionManagementDeleteCallType) UnmarshalJSON(data []byte) error {
	return util.Unmarshal(data, &m)
}
