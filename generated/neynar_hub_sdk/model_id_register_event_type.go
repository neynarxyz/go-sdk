/*
Farcaster Hub API

Perform basic queries of Farcaster state via the REST API of a Farcaster hub. See the [Neynar docs](https://docs.neynar.com/reference) for more details.

API version: 2.35.0
Contact: team@neynar.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package neynar_hub_sdk

import (
	"encoding/json"
	"fmt"
)

// IdRegisterEventType the model 'IdRegisterEventType'
type IdRegisterEventType string

// List of IdRegisterEventType
const (
	IDREGISTEREVENTTYPE_ID_REGISTER_EVENT_TYPE_REGISTER        IdRegisterEventType = "ID_REGISTER_EVENT_TYPE_REGISTER"
	IDREGISTEREVENTTYPE_ID_REGISTER_EVENT_TYPE_TRANSFER        IdRegisterEventType = "ID_REGISTER_EVENT_TYPE_TRANSFER"
	IDREGISTEREVENTTYPE_ID_REGISTER_EVENT_TYPE_CHANGE_RECOVERY IdRegisterEventType = "ID_REGISTER_EVENT_TYPE_CHANGE_RECOVERY"
)

// All allowed values of IdRegisterEventType enum
var AllowedIdRegisterEventTypeEnumValues = []IdRegisterEventType{
	"ID_REGISTER_EVENT_TYPE_REGISTER",
	"ID_REGISTER_EVENT_TYPE_TRANSFER",
	"ID_REGISTER_EVENT_TYPE_CHANGE_RECOVERY",
}

func (v *IdRegisterEventType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := IdRegisterEventType(value)
	for _, existing := range AllowedIdRegisterEventTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid IdRegisterEventType", value)
}

// NewIdRegisterEventTypeFromValue returns a pointer to a valid IdRegisterEventType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewIdRegisterEventTypeFromValue(v string) (*IdRegisterEventType, error) {
	ev := IdRegisterEventType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for IdRegisterEventType: valid values are %v", v, AllowedIdRegisterEventTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v IdRegisterEventType) IsValid() bool {
	for _, existing := range AllowedIdRegisterEventTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IdRegisterEventType value
func (v IdRegisterEventType) Ptr() *IdRegisterEventType {
	return &v
}

type NullableIdRegisterEventType struct {
	value *IdRegisterEventType
	isSet bool
}

func (v NullableIdRegisterEventType) Get() *IdRegisterEventType {
	return v.value
}

func (v *NullableIdRegisterEventType) Set(val *IdRegisterEventType) {
	v.value = val
	v.isSet = true
}

func (v NullableIdRegisterEventType) IsSet() bool {
	return v.isSet
}

func (v *NullableIdRegisterEventType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdRegisterEventType(val *IdRegisterEventType) *NullableIdRegisterEventType {
	return &NullableIdRegisterEventType{value: val, isSet: true}
}

func (v NullableIdRegisterEventType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdRegisterEventType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
