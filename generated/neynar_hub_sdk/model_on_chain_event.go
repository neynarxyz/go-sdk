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
	"gopkg.in/validator.v2"
)

// OnChainEvent - struct for OnChainEvent
type OnChainEvent struct {
	OnChainEventIdRegister     *OnChainEventIdRegister
	OnChainEventSigner         *OnChainEventSigner
	OnChainEventSignerMigrated *OnChainEventSignerMigrated
	OnChainEventStorageRent    *OnChainEventStorageRent
}

// OnChainEventIdRegisterAsOnChainEvent is a convenience function that returns OnChainEventIdRegister wrapped in OnChainEvent
func OnChainEventIdRegisterAsOnChainEvent(v *OnChainEventIdRegister) OnChainEvent {
	return OnChainEvent{
		OnChainEventIdRegister: v,
	}
}

// OnChainEventSignerAsOnChainEvent is a convenience function that returns OnChainEventSigner wrapped in OnChainEvent
func OnChainEventSignerAsOnChainEvent(v *OnChainEventSigner) OnChainEvent {
	return OnChainEvent{
		OnChainEventSigner: v,
	}
}

// OnChainEventSignerMigratedAsOnChainEvent is a convenience function that returns OnChainEventSignerMigrated wrapped in OnChainEvent
func OnChainEventSignerMigratedAsOnChainEvent(v *OnChainEventSignerMigrated) OnChainEvent {
	return OnChainEvent{
		OnChainEventSignerMigrated: v,
	}
}

// OnChainEventStorageRentAsOnChainEvent is a convenience function that returns OnChainEventStorageRent wrapped in OnChainEvent
func OnChainEventStorageRentAsOnChainEvent(v *OnChainEventStorageRent) OnChainEvent {
	return OnChainEvent{
		OnChainEventStorageRent: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *OnChainEvent) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into OnChainEventIdRegister
	err = newStrictDecoder(data).Decode(&dst.OnChainEventIdRegister)
	if err == nil {
		jsonOnChainEventIdRegister, _ := json.Marshal(dst.OnChainEventIdRegister)
		if string(jsonOnChainEventIdRegister) == "{}" { // empty struct
			dst.OnChainEventIdRegister = nil
		} else {
			if err = validator.Validate(dst.OnChainEventIdRegister); err != nil {
				dst.OnChainEventIdRegister = nil
			} else {
				match++
			}
		}
	} else {
		dst.OnChainEventIdRegister = nil
	}

	// try to unmarshal data into OnChainEventSigner
	err = newStrictDecoder(data).Decode(&dst.OnChainEventSigner)
	if err == nil {
		jsonOnChainEventSigner, _ := json.Marshal(dst.OnChainEventSigner)
		if string(jsonOnChainEventSigner) == "{}" { // empty struct
			dst.OnChainEventSigner = nil
		} else {
			if err = validator.Validate(dst.OnChainEventSigner); err != nil {
				dst.OnChainEventSigner = nil
			} else {
				match++
			}
		}
	} else {
		dst.OnChainEventSigner = nil
	}

	// try to unmarshal data into OnChainEventSignerMigrated
	err = newStrictDecoder(data).Decode(&dst.OnChainEventSignerMigrated)
	if err == nil {
		jsonOnChainEventSignerMigrated, _ := json.Marshal(dst.OnChainEventSignerMigrated)
		if string(jsonOnChainEventSignerMigrated) == "{}" { // empty struct
			dst.OnChainEventSignerMigrated = nil
		} else {
			if err = validator.Validate(dst.OnChainEventSignerMigrated); err != nil {
				dst.OnChainEventSignerMigrated = nil
			} else {
				match++
			}
		}
	} else {
		dst.OnChainEventSignerMigrated = nil
	}

	// try to unmarshal data into OnChainEventStorageRent
	err = newStrictDecoder(data).Decode(&dst.OnChainEventStorageRent)
	if err == nil {
		jsonOnChainEventStorageRent, _ := json.Marshal(dst.OnChainEventStorageRent)
		if string(jsonOnChainEventStorageRent) == "{}" { // empty struct
			dst.OnChainEventStorageRent = nil
		} else {
			if err = validator.Validate(dst.OnChainEventStorageRent); err != nil {
				dst.OnChainEventStorageRent = nil
			} else {
				match++
			}
		}
	} else {
		dst.OnChainEventStorageRent = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.OnChainEventIdRegister = nil
		dst.OnChainEventSigner = nil
		dst.OnChainEventSignerMigrated = nil
		dst.OnChainEventStorageRent = nil

		return fmt.Errorf("data matches more than one schema in oneOf(OnChainEvent)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(OnChainEvent)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src OnChainEvent) MarshalJSON() ([]byte, error) {
	if src.OnChainEventIdRegister != nil {
		return json.Marshal(&src.OnChainEventIdRegister)
	}

	if src.OnChainEventSigner != nil {
		return json.Marshal(&src.OnChainEventSigner)
	}

	if src.OnChainEventSignerMigrated != nil {
		return json.Marshal(&src.OnChainEventSignerMigrated)
	}

	if src.OnChainEventStorageRent != nil {
		return json.Marshal(&src.OnChainEventStorageRent)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *OnChainEvent) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.OnChainEventIdRegister != nil {
		return obj.OnChainEventIdRegister
	}

	if obj.OnChainEventSigner != nil {
		return obj.OnChainEventSigner
	}

	if obj.OnChainEventSignerMigrated != nil {
		return obj.OnChainEventSignerMigrated
	}

	if obj.OnChainEventStorageRent != nil {
		return obj.OnChainEventStorageRent
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj OnChainEvent) GetActualInstanceValue() interface{} {
	if obj.OnChainEventIdRegister != nil {
		return *obj.OnChainEventIdRegister
	}

	if obj.OnChainEventSigner != nil {
		return *obj.OnChainEventSigner
	}

	if obj.OnChainEventSignerMigrated != nil {
		return *obj.OnChainEventSignerMigrated
	}

	if obj.OnChainEventStorageRent != nil {
		return *obj.OnChainEventStorageRent
	}

	// all schemas are nil
	return nil
}

type NullableOnChainEvent struct {
	value *OnChainEvent
	isSet bool
}

func (v NullableOnChainEvent) Get() *OnChainEvent {
	return v.value
}

func (v *NullableOnChainEvent) Set(val *OnChainEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableOnChainEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableOnChainEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOnChainEvent(val *OnChainEvent) *NullableOnChainEvent {
	return &NullableOnChainEvent{value: val, isSet: true}
}

func (v NullableOnChainEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOnChainEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
