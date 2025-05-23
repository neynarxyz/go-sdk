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

// FeedType the model 'FeedType'
type FeedType string

// List of FeedType
const (
	FEEDTYPE_FOLLOWING FeedType = "following"
	FEEDTYPE_FILTER    FeedType = "filter"
)

// All allowed values of FeedType enum
var AllowedFeedTypeEnumValues = []FeedType{
	"following",
	"filter",
}

func (v *FeedType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FeedType(value)
	for _, existing := range AllowedFeedTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FeedType", value)
}

// NewFeedTypeFromValue returns a pointer to a valid FeedType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFeedTypeFromValue(v string) (*FeedType, error) {
	ev := FeedType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FeedType: valid values are %v", v, AllowedFeedTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FeedType) IsValid() bool {
	for _, existing := range AllowedFeedTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FeedType value
func (v FeedType) Ptr() *FeedType {
	return &v
}

type NullableFeedType struct {
	value *FeedType
	isSet bool
}

func (v NullableFeedType) Get() *FeedType {
	return v.value
}

func (v *NullableFeedType) Set(val *FeedType) {
	v.value = val
	v.isSet = true
}

func (v NullableFeedType) IsSet() bool {
	return v.isSet
}

func (v *NullableFeedType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeedType(val *FeedType) *NullableFeedType {
	return &NullableFeedType{value: val, isSet: true}
}

func (v NullableFeedType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeedType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
