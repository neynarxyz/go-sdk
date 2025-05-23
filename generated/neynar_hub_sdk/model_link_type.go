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

// LinkType Defines the type of social connection between users. - follow: Establishes a following relationship where the user will receive updates from the target user in their feed
type LinkType string

// List of LinkType
const (
	LINKTYPE_FOLLOW LinkType = "follow"
)

// All allowed values of LinkType enum
var AllowedLinkTypeEnumValues = []LinkType{
	"follow",
}

func (v *LinkType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := LinkType(value)
	for _, existing := range AllowedLinkTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid LinkType", value)
}

// NewLinkTypeFromValue returns a pointer to a valid LinkType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLinkTypeFromValue(v string) (*LinkType, error) {
	ev := LinkType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LinkType: valid values are %v", v, AllowedLinkTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LinkType) IsValid() bool {
	for _, existing := range AllowedLinkTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LinkType value
func (v LinkType) Ptr() *LinkType {
	return &v
}

type NullableLinkType struct {
	value *LinkType
	isSet bool
}

func (v NullableLinkType) Get() *LinkType {
	return v.value
}

func (v *NullableLinkType) Set(val *LinkType) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkType) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkType(val *LinkType) *NullableLinkType {
	return &NullableLinkType{value: val, isSet: true}
}

func (v NullableLinkType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
