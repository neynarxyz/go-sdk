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
	"gopkg.in/validator.v2"
)

// ChannelMemberUser - struct for ChannelMemberUser
type ChannelMemberUser struct {
	User           *User
	UserDehydrated *UserDehydrated
}

// UserAsChannelMemberUser is a convenience function that returns User wrapped in ChannelMemberUser
func UserAsChannelMemberUser(v *User) ChannelMemberUser {
	return ChannelMemberUser{
		User: v,
	}
}

// UserDehydratedAsChannelMemberUser is a convenience function that returns UserDehydrated wrapped in ChannelMemberUser
func UserDehydratedAsChannelMemberUser(v *UserDehydrated) ChannelMemberUser {
	return ChannelMemberUser{
		UserDehydrated: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ChannelMemberUser) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into User
	err = newStrictDecoder(data).Decode(&dst.User)
	if err == nil {
		jsonUser, _ := json.Marshal(dst.User)
		if string(jsonUser) == "{}" { // empty struct
			dst.User = nil
		} else {
			if err = validator.Validate(dst.User); err != nil {
				dst.User = nil
			} else {
				match++
			}
		}
	} else {
		dst.User = nil
	}

	// try to unmarshal data into UserDehydrated
	err = newStrictDecoder(data).Decode(&dst.UserDehydrated)
	if err == nil {
		jsonUserDehydrated, _ := json.Marshal(dst.UserDehydrated)
		if string(jsonUserDehydrated) == "{}" { // empty struct
			dst.UserDehydrated = nil
		} else {
			if err = validator.Validate(dst.UserDehydrated); err != nil {
				dst.UserDehydrated = nil
			} else {
				match++
			}
		}
	} else {
		dst.UserDehydrated = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.User = nil
		dst.UserDehydrated = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ChannelMemberUser)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ChannelMemberUser)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ChannelMemberUser) MarshalJSON() ([]byte, error) {
	if src.User != nil {
		return json.Marshal(&src.User)
	}

	if src.UserDehydrated != nil {
		return json.Marshal(&src.UserDehydrated)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ChannelMemberUser) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.User != nil {
		return obj.User
	}

	if obj.UserDehydrated != nil {
		return obj.UserDehydrated
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj ChannelMemberUser) GetActualInstanceValue() interface{} {
	if obj.User != nil {
		return *obj.User
	}

	if obj.UserDehydrated != nil {
		return *obj.UserDehydrated
	}

	// all schemas are nil
	return nil
}

type NullableChannelMemberUser struct {
	value *ChannelMemberUser
	isSet bool
}

func (v NullableChannelMemberUser) Get() *ChannelMemberUser {
	return v.value
}

func (v *NullableChannelMemberUser) Set(val *ChannelMemberUser) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelMemberUser) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelMemberUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelMemberUser(val *ChannelMemberUser) *NullableChannelMemberUser {
	return &NullableChannelMemberUser{value: val, isSet: true}
}

func (v NullableChannelMemberUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelMemberUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
