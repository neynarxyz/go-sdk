/*
Farcaster API V2

The Farcaster API allows you to interact with the Farcaster protocol. See the [Neynar docs](https://docs.neynar.com/reference) for more details.

API version: 2.43.0
Contact: team@neynar.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package neynar_sdk

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the AddressBalanceVerifiedAddress type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddressBalanceVerifiedAddress{}

// AddressBalanceVerifiedAddress struct for AddressBalanceVerifiedAddress
type AddressBalanceVerifiedAddress struct {
	// The wallet address
	Address string  `json:"address"`
	Network Network `json:"network"`
}

type _AddressBalanceVerifiedAddress AddressBalanceVerifiedAddress

// NewAddressBalanceVerifiedAddress instantiates a new AddressBalanceVerifiedAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressBalanceVerifiedAddress(address string, network Network) *AddressBalanceVerifiedAddress {
	this := AddressBalanceVerifiedAddress{}
	this.Address = address
	this.Network = network
	return &this
}

// NewAddressBalanceVerifiedAddressWithDefaults instantiates a new AddressBalanceVerifiedAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressBalanceVerifiedAddressWithDefaults() *AddressBalanceVerifiedAddress {
	this := AddressBalanceVerifiedAddress{}
	return &this
}

// GetAddress returns the Address field value
func (o *AddressBalanceVerifiedAddress) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *AddressBalanceVerifiedAddress) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *AddressBalanceVerifiedAddress) SetAddress(v string) {
	o.Address = v
}

// GetNetwork returns the Network field value
func (o *AddressBalanceVerifiedAddress) GetNetwork() Network {
	if o == nil {
		var ret Network
		return ret
	}

	return o.Network
}

// GetNetworkOk returns a tuple with the Network field value
// and a boolean to check if the value has been set.
func (o *AddressBalanceVerifiedAddress) GetNetworkOk() (*Network, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Network, true
}

// SetNetwork sets field value
func (o *AddressBalanceVerifiedAddress) SetNetwork(v Network) {
	o.Network = v
}

func (o AddressBalanceVerifiedAddress) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddressBalanceVerifiedAddress) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["address"] = o.Address
	toSerialize["network"] = o.Network
	return toSerialize, nil
}

func (o *AddressBalanceVerifiedAddress) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"address",
		"network",
	}

	// defaultValueFuncMap captures the default values for required properties.
	// These values are used when required properties are missing from the payload.
	defaultValueFuncMap := map[string]func() interface{}{}
	var defaultValueApplied bool
	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if value, exists := allProperties[requiredProperty]; !exists || value == "" {
			if _, ok := defaultValueFuncMap[requiredProperty]; ok {
				allProperties[requiredProperty] = defaultValueFuncMap[requiredProperty]()
				defaultValueApplied = true
			}
		}
		if value, exists := allProperties[requiredProperty]; !exists || value == "" {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	if defaultValueApplied {
		data, err = json.Marshal(allProperties)
		if err != nil {
			return err
		}
	}
	varAddressBalanceVerifiedAddress := _AddressBalanceVerifiedAddress{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAddressBalanceVerifiedAddress)

	if err != nil {
		return err
	}

	*o = AddressBalanceVerifiedAddress(varAddressBalanceVerifiedAddress)

	return err
}

type NullableAddressBalanceVerifiedAddress struct {
	value *AddressBalanceVerifiedAddress
	isSet bool
}

func (v NullableAddressBalanceVerifiedAddress) Get() *AddressBalanceVerifiedAddress {
	return v.value
}

func (v *NullableAddressBalanceVerifiedAddress) Set(val *AddressBalanceVerifiedAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressBalanceVerifiedAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressBalanceVerifiedAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressBalanceVerifiedAddress(val *AddressBalanceVerifiedAddress) *NullableAddressBalanceVerifiedAddress {
	return &NullableAddressBalanceVerifiedAddress{value: val, isSet: true}
}

func (v NullableAddressBalanceVerifiedAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressBalanceVerifiedAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
