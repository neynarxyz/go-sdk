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

// checks if the BalanceResponseUserBalance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BalanceResponseUserBalance{}

// BalanceResponseUserBalance struct for BalanceResponseUserBalance
type BalanceResponseUserBalance struct {
	Object          string           `json:"object"`
	User            UserDehydrated   `json:"user"`
	AddressBalances []AddressBalance `json:"address_balances"`
}

type _BalanceResponseUserBalance BalanceResponseUserBalance

// NewBalanceResponseUserBalance instantiates a new BalanceResponseUserBalance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBalanceResponseUserBalance(object string, user UserDehydrated, addressBalances []AddressBalance) *BalanceResponseUserBalance {
	this := BalanceResponseUserBalance{}
	this.Object = object
	this.User = user
	this.AddressBalances = addressBalances
	return &this
}

// NewBalanceResponseUserBalanceWithDefaults instantiates a new BalanceResponseUserBalance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBalanceResponseUserBalanceWithDefaults() *BalanceResponseUserBalance {
	this := BalanceResponseUserBalance{}
	return &this
}

// GetObject returns the Object field value
func (o *BalanceResponseUserBalance) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *BalanceResponseUserBalance) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *BalanceResponseUserBalance) SetObject(v string) {
	o.Object = v
}

// GetUser returns the User field value
func (o *BalanceResponseUserBalance) GetUser() UserDehydrated {
	if o == nil {
		var ret UserDehydrated
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *BalanceResponseUserBalance) GetUserOk() (*UserDehydrated, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *BalanceResponseUserBalance) SetUser(v UserDehydrated) {
	o.User = v
}

// GetAddressBalances returns the AddressBalances field value
func (o *BalanceResponseUserBalance) GetAddressBalances() []AddressBalance {
	if o == nil {
		var ret []AddressBalance
		return ret
	}

	return o.AddressBalances
}

// GetAddressBalancesOk returns a tuple with the AddressBalances field value
// and a boolean to check if the value has been set.
func (o *BalanceResponseUserBalance) GetAddressBalancesOk() ([]AddressBalance, bool) {
	if o == nil {
		return nil, false
	}
	return o.AddressBalances, true
}

// SetAddressBalances sets field value
func (o *BalanceResponseUserBalance) SetAddressBalances(v []AddressBalance) {
	o.AddressBalances = v
}

func (o BalanceResponseUserBalance) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BalanceResponseUserBalance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["object"] = o.Object
	toSerialize["user"] = o.User
	toSerialize["address_balances"] = o.AddressBalances
	return toSerialize, nil
}

func (o *BalanceResponseUserBalance) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"object",
		"user",
		"address_balances",
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
	varBalanceResponseUserBalance := _BalanceResponseUserBalance{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBalanceResponseUserBalance)

	if err != nil {
		return err
	}

	*o = BalanceResponseUserBalance(varBalanceResponseUserBalance)

	return err
}

type NullableBalanceResponseUserBalance struct {
	value *BalanceResponseUserBalance
	isSet bool
}

func (v NullableBalanceResponseUserBalance) Get() *BalanceResponseUserBalance {
	return v.value
}

func (v *NullableBalanceResponseUserBalance) Set(val *BalanceResponseUserBalance) {
	v.value = val
	v.isSet = true
}

func (v NullableBalanceResponseUserBalance) IsSet() bool {
	return v.isSet
}

func (v *NullableBalanceResponseUserBalance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBalanceResponseUserBalance(val *BalanceResponseUserBalance) *NullableBalanceResponseUserBalance {
	return &NullableBalanceResponseUserBalance{value: val, isSet: true}
}

func (v NullableBalanceResponseUserBalance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBalanceResponseUserBalance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
