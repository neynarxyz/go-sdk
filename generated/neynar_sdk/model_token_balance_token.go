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

// checks if the TokenBalanceToken type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TokenBalanceToken{}

// TokenBalanceToken struct for TokenBalanceToken
type TokenBalanceToken struct {
	Object string `json:"object"`
	// The token name e.g. \"Ethereum\"
	Name string `json:"name"`
	// The token symbol e.g. \"ETH\"
	Symbol string `json:"symbol"`
	// The contract address of the token (omitted for native token)
	Address *string `json:"address,omitempty"`
	// The number of decimals the token uses
	Decimals *int32 `json:"decimals,omitempty"`
}

type _TokenBalanceToken TokenBalanceToken

// NewTokenBalanceToken instantiates a new TokenBalanceToken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenBalanceToken(object string, name string, symbol string) *TokenBalanceToken {
	this := TokenBalanceToken{}
	this.Object = object
	this.Name = name
	this.Symbol = symbol
	return &this
}

// NewTokenBalanceTokenWithDefaults instantiates a new TokenBalanceToken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenBalanceTokenWithDefaults() *TokenBalanceToken {
	this := TokenBalanceToken{}
	return &this
}

// GetObject returns the Object field value
func (o *TokenBalanceToken) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *TokenBalanceToken) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *TokenBalanceToken) SetObject(v string) {
	o.Object = v
}

// GetName returns the Name field value
func (o *TokenBalanceToken) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TokenBalanceToken) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TokenBalanceToken) SetName(v string) {
	o.Name = v
}

// GetSymbol returns the Symbol field value
func (o *TokenBalanceToken) GetSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *TokenBalanceToken) GetSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *TokenBalanceToken) SetSymbol(v string) {
	o.Symbol = v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *TokenBalanceToken) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenBalanceToken) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *TokenBalanceToken) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *TokenBalanceToken) SetAddress(v string) {
	o.Address = &v
}

// GetDecimals returns the Decimals field value if set, zero value otherwise.
func (o *TokenBalanceToken) GetDecimals() int32 {
	if o == nil || IsNil(o.Decimals) {
		var ret int32
		return ret
	}
	return *o.Decimals
}

// GetDecimalsOk returns a tuple with the Decimals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenBalanceToken) GetDecimalsOk() (*int32, bool) {
	if o == nil || IsNil(o.Decimals) {
		return nil, false
	}
	return o.Decimals, true
}

// HasDecimals returns a boolean if a field has been set.
func (o *TokenBalanceToken) HasDecimals() bool {
	if o != nil && !IsNil(o.Decimals) {
		return true
	}

	return false
}

// SetDecimals gets a reference to the given int32 and assigns it to the Decimals field.
func (o *TokenBalanceToken) SetDecimals(v int32) {
	o.Decimals = &v
}

func (o TokenBalanceToken) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TokenBalanceToken) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["object"] = o.Object
	toSerialize["name"] = o.Name
	toSerialize["symbol"] = o.Symbol
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.Decimals) {
		toSerialize["decimals"] = o.Decimals
	}
	return toSerialize, nil
}

func (o *TokenBalanceToken) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"object",
		"name",
		"symbol",
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
	varTokenBalanceToken := _TokenBalanceToken{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTokenBalanceToken)

	if err != nil {
		return err
	}

	*o = TokenBalanceToken(varTokenBalanceToken)

	return err
}

type NullableTokenBalanceToken struct {
	value *TokenBalanceToken
	isSet bool
}

func (v NullableTokenBalanceToken) Get() *TokenBalanceToken {
	return v.value
}

func (v *NullableTokenBalanceToken) Set(val *TokenBalanceToken) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenBalanceToken) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenBalanceToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenBalanceToken(val *TokenBalanceToken) *NullableTokenBalanceToken {
	return &NullableTokenBalanceToken{value: val, isSet: true}
}

func (v NullableTokenBalanceToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenBalanceToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
