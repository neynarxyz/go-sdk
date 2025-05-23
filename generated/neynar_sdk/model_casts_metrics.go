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
	"time"
)

// checks if the CastsMetrics type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CastsMetrics{}

// CastsMetrics struct for CastsMetrics
type CastsMetrics struct {
	Start               time.Time `json:"start"`
	ResolutionInSeconds int32     `json:"resolution_in_seconds"`
	CastCount           int32     `json:"cast_count"`
}

type _CastsMetrics CastsMetrics

// NewCastsMetrics instantiates a new CastsMetrics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCastsMetrics(start time.Time, resolutionInSeconds int32, castCount int32) *CastsMetrics {
	this := CastsMetrics{}
	this.Start = start
	this.ResolutionInSeconds = resolutionInSeconds
	this.CastCount = castCount
	return &this
}

// NewCastsMetricsWithDefaults instantiates a new CastsMetrics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCastsMetricsWithDefaults() *CastsMetrics {
	this := CastsMetrics{}
	return &this
}

// GetStart returns the Start field value
func (o *CastsMetrics) GetStart() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Start
}

// GetStartOk returns a tuple with the Start field value
// and a boolean to check if the value has been set.
func (o *CastsMetrics) GetStartOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Start, true
}

// SetStart sets field value
func (o *CastsMetrics) SetStart(v time.Time) {
	o.Start = v
}

// GetResolutionInSeconds returns the ResolutionInSeconds field value
func (o *CastsMetrics) GetResolutionInSeconds() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ResolutionInSeconds
}

// GetResolutionInSecondsOk returns a tuple with the ResolutionInSeconds field value
// and a boolean to check if the value has been set.
func (o *CastsMetrics) GetResolutionInSecondsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResolutionInSeconds, true
}

// SetResolutionInSeconds sets field value
func (o *CastsMetrics) SetResolutionInSeconds(v int32) {
	o.ResolutionInSeconds = v
}

// GetCastCount returns the CastCount field value
func (o *CastsMetrics) GetCastCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CastCount
}

// GetCastCountOk returns a tuple with the CastCount field value
// and a boolean to check if the value has been set.
func (o *CastsMetrics) GetCastCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CastCount, true
}

// SetCastCount sets field value
func (o *CastsMetrics) SetCastCount(v int32) {
	o.CastCount = v
}

func (o CastsMetrics) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CastsMetrics) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["start"] = o.Start
	toSerialize["resolution_in_seconds"] = o.ResolutionInSeconds
	toSerialize["cast_count"] = o.CastCount
	return toSerialize, nil
}

func (o *CastsMetrics) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"start",
		"resolution_in_seconds",
		"cast_count",
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
	varCastsMetrics := _CastsMetrics{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCastsMetrics)

	if err != nil {
		return err
	}

	*o = CastsMetrics(varCastsMetrics)

	return err
}

type NullableCastsMetrics struct {
	value *CastsMetrics
	isSet bool
}

func (v NullableCastsMetrics) Get() *CastsMetrics {
	return v.value
}

func (v *NullableCastsMetrics) Set(val *CastsMetrics) {
	v.value = val
	v.isSet = true
}

func (v NullableCastsMetrics) IsSet() bool {
	return v.isSet
}

func (v *NullableCastsMetrics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCastsMetrics(val *CastsMetrics) *NullableCastsMetrics {
	return &NullableCastsMetrics{value: val, isSet: true}
}

func (v NullableCastsMetrics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCastsMetrics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
