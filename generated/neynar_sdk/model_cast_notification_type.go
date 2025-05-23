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

// CastNotificationType The notification type of a cast.
type CastNotificationType string

// List of CastNotificationType
const (
	CASTNOTIFICATIONTYPE_CAST_MENTION CastNotificationType = "cast-mention"
	CASTNOTIFICATIONTYPE_CAST_REPLY   CastNotificationType = "cast-reply"
)

// All allowed values of CastNotificationType enum
var AllowedCastNotificationTypeEnumValues = []CastNotificationType{
	"cast-mention",
	"cast-reply",
}

func (v *CastNotificationType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CastNotificationType(value)
	for _, existing := range AllowedCastNotificationTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CastNotificationType", value)
}

// NewCastNotificationTypeFromValue returns a pointer to a valid CastNotificationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCastNotificationTypeFromValue(v string) (*CastNotificationType, error) {
	ev := CastNotificationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CastNotificationType: valid values are %v", v, AllowedCastNotificationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CastNotificationType) IsValid() bool {
	for _, existing := range AllowedCastNotificationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CastNotificationType value
func (v CastNotificationType) Ptr() *CastNotificationType {
	return &v
}

type NullableCastNotificationType struct {
	value *CastNotificationType
	isSet bool
}

func (v NullableCastNotificationType) Get() *CastNotificationType {
	return v.value
}

func (v *NullableCastNotificationType) Set(val *CastNotificationType) {
	v.value = val
	v.isSet = true
}

func (v NullableCastNotificationType) IsSet() bool {
	return v.isSet
}

func (v *NullableCastNotificationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCastNotificationType(val *CastNotificationType) *NullableCastNotificationType {
	return &NullableCastNotificationType{value: val, isSet: true}
}

func (v NullableCastNotificationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCastNotificationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
