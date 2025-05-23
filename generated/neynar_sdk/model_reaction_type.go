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

// ReactionType the model 'ReactionType'
type ReactionType string

// List of ReactionType
const (
	REACTIONTYPE_LIKE   ReactionType = "like"
	REACTIONTYPE_RECAST ReactionType = "recast"
)

// All allowed values of ReactionType enum
var AllowedReactionTypeEnumValues = []ReactionType{
	"like",
	"recast",
}

func (v *ReactionType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ReactionType(value)
	for _, existing := range AllowedReactionTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ReactionType", value)
}

// NewReactionTypeFromValue returns a pointer to a valid ReactionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReactionTypeFromValue(v string) (*ReactionType, error) {
	ev := ReactionType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReactionType: valid values are %v", v, AllowedReactionTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReactionType) IsValid() bool {
	for _, existing := range AllowedReactionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ReactionType value
func (v ReactionType) Ptr() *ReactionType {
	return &v
}

type NullableReactionType struct {
	value *ReactionType
	isSet bool
}

func (v NullableReactionType) Get() *ReactionType {
	return v.value
}

func (v *NullableReactionType) Set(val *ReactionType) {
	v.value = val
	v.isSet = true
}

func (v NullableReactionType) IsSet() bool {
	return v.isSet
}

func (v *NullableReactionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReactionType(val *ReactionType) *NullableReactionType {
	return &NullableReactionType{value: val, isSet: true}
}

func (v NullableReactionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReactionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
