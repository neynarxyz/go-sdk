/*
Farcaster API V2

The Farcaster API allows you to interact with the Farcaster protocol. See the [Neynar docs](https://docs.neynar.com/reference) for more details.

API version: 2.43.0
Contact: team@neynar.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package neynar_sdk

import (
	"encoding/json"
	"fmt"
)

// FrameType Type of identifier (either 'uuid' or 'url')
type FrameType string

// List of FrameType
const (
	FRAMETYPE_UUID FrameType = "uuid"
	FRAMETYPE_URL  FrameType = "url"
)

// All allowed values of FrameType enum
var AllowedFrameTypeEnumValues = []FrameType{
	"uuid",
	"url",
}

func (v *FrameType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FrameType(value)
	for _, existing := range AllowedFrameTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FrameType", value)
}

// NewFrameTypeFromValue returns a pointer to a valid FrameType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFrameTypeFromValue(v string) (*FrameType, error) {
	ev := FrameType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FrameType: valid values are %v", v, AllowedFrameTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FrameType) IsValid() bool {
	for _, existing := range AllowedFrameTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FrameType value
func (v FrameType) Ptr() *FrameType {
	return &v
}

type NullableFrameType struct {
	value *FrameType
	isSet bool
}

func (v NullableFrameType) Get() *FrameType {
	return v.value
}

func (v *NullableFrameType) Set(val *FrameType) {
	v.value = val
	v.isSet = true
}

func (v NullableFrameType) IsSet() bool {
	return v.isSet
}

func (v *NullableFrameType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFrameType(val *FrameType) *NullableFrameType {
	return &NullableFrameType{value: val, isSet: true}
}

func (v NullableFrameType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFrameType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
