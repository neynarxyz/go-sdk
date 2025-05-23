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
)

// checks if the SubscriptionTierPrice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionTierPrice{}

// SubscriptionTierPrice struct for SubscriptionTierPrice
type SubscriptionTierPrice struct {
	PeriodDurationSeconds *int32  `json:"period_duration_seconds,omitempty"`
	TokensPerPeriod       *string `json:"tokens_per_period,omitempty"`
	InitialMintPrice      *string `json:"initial_mint_price,omitempty"`
}

// NewSubscriptionTierPrice instantiates a new SubscriptionTierPrice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionTierPrice() *SubscriptionTierPrice {
	this := SubscriptionTierPrice{}
	return &this
}

// NewSubscriptionTierPriceWithDefaults instantiates a new SubscriptionTierPrice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionTierPriceWithDefaults() *SubscriptionTierPrice {
	this := SubscriptionTierPrice{}
	return &this
}

// GetPeriodDurationSeconds returns the PeriodDurationSeconds field value if set, zero value otherwise.
func (o *SubscriptionTierPrice) GetPeriodDurationSeconds() int32 {
	if o == nil || IsNil(o.PeriodDurationSeconds) {
		var ret int32
		return ret
	}
	return *o.PeriodDurationSeconds
}

// GetPeriodDurationSecondsOk returns a tuple with the PeriodDurationSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionTierPrice) GetPeriodDurationSecondsOk() (*int32, bool) {
	if o == nil || IsNil(o.PeriodDurationSeconds) {
		return nil, false
	}
	return o.PeriodDurationSeconds, true
}

// HasPeriodDurationSeconds returns a boolean if a field has been set.
func (o *SubscriptionTierPrice) HasPeriodDurationSeconds() bool {
	if o != nil && !IsNil(o.PeriodDurationSeconds) {
		return true
	}

	return false
}

// SetPeriodDurationSeconds gets a reference to the given int32 and assigns it to the PeriodDurationSeconds field.
func (o *SubscriptionTierPrice) SetPeriodDurationSeconds(v int32) {
	o.PeriodDurationSeconds = &v
}

// GetTokensPerPeriod returns the TokensPerPeriod field value if set, zero value otherwise.
func (o *SubscriptionTierPrice) GetTokensPerPeriod() string {
	if o == nil || IsNil(o.TokensPerPeriod) {
		var ret string
		return ret
	}
	return *o.TokensPerPeriod
}

// GetTokensPerPeriodOk returns a tuple with the TokensPerPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionTierPrice) GetTokensPerPeriodOk() (*string, bool) {
	if o == nil || IsNil(o.TokensPerPeriod) {
		return nil, false
	}
	return o.TokensPerPeriod, true
}

// HasTokensPerPeriod returns a boolean if a field has been set.
func (o *SubscriptionTierPrice) HasTokensPerPeriod() bool {
	if o != nil && !IsNil(o.TokensPerPeriod) {
		return true
	}

	return false
}

// SetTokensPerPeriod gets a reference to the given string and assigns it to the TokensPerPeriod field.
func (o *SubscriptionTierPrice) SetTokensPerPeriod(v string) {
	o.TokensPerPeriod = &v
}

// GetInitialMintPrice returns the InitialMintPrice field value if set, zero value otherwise.
func (o *SubscriptionTierPrice) GetInitialMintPrice() string {
	if o == nil || IsNil(o.InitialMintPrice) {
		var ret string
		return ret
	}
	return *o.InitialMintPrice
}

// GetInitialMintPriceOk returns a tuple with the InitialMintPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionTierPrice) GetInitialMintPriceOk() (*string, bool) {
	if o == nil || IsNil(o.InitialMintPrice) {
		return nil, false
	}
	return o.InitialMintPrice, true
}

// HasInitialMintPrice returns a boolean if a field has been set.
func (o *SubscriptionTierPrice) HasInitialMintPrice() bool {
	if o != nil && !IsNil(o.InitialMintPrice) {
		return true
	}

	return false
}

// SetInitialMintPrice gets a reference to the given string and assigns it to the InitialMintPrice field.
func (o *SubscriptionTierPrice) SetInitialMintPrice(v string) {
	o.InitialMintPrice = &v
}

func (o SubscriptionTierPrice) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionTierPrice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PeriodDurationSeconds) {
		toSerialize["period_duration_seconds"] = o.PeriodDurationSeconds
	}
	if !IsNil(o.TokensPerPeriod) {
		toSerialize["tokens_per_period"] = o.TokensPerPeriod
	}
	if !IsNil(o.InitialMintPrice) {
		toSerialize["initial_mint_price"] = o.InitialMintPrice
	}
	return toSerialize, nil
}

type NullableSubscriptionTierPrice struct {
	value *SubscriptionTierPrice
	isSet bool
}

func (v NullableSubscriptionTierPrice) Get() *SubscriptionTierPrice {
	return v.value
}

func (v *NullableSubscriptionTierPrice) Set(val *SubscriptionTierPrice) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionTierPrice) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionTierPrice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionTierPrice(val *SubscriptionTierPrice) *NullableSubscriptionTierPrice {
	return &NullableSubscriptionTierPrice{value: val, isSet: true}
}

func (v NullableSubscriptionTierPrice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionTierPrice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
