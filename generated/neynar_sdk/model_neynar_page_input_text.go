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

// checks if the NeynarPageInputText type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NeynarPageInputText{}

// NeynarPageInputText struct for NeynarPageInputText
type NeynarPageInputText struct {
	// Indicates if text input is enabled.
	Enabled bool `json:"enabled"`
	// The placeholder text for the input.
	Placeholder *string `json:"placeholder,omitempty"`
}

type _NeynarPageInputText NeynarPageInputText

// NewNeynarPageInputText instantiates a new NeynarPageInputText object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNeynarPageInputText(enabled bool) *NeynarPageInputText {
	this := NeynarPageInputText{}
	this.Enabled = enabled
	return &this
}

// NewNeynarPageInputTextWithDefaults instantiates a new NeynarPageInputText object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNeynarPageInputTextWithDefaults() *NeynarPageInputText {
	this := NeynarPageInputText{}
	var enabled bool = false
	this.Enabled = enabled
	return &this
}

// GetEnabled returns the Enabled field value
func (o *NeynarPageInputText) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *NeynarPageInputText) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *NeynarPageInputText) SetEnabled(v bool) {
	o.Enabled = v
}

// GetDefaultEnabled returns the default value false of the Enabled field.
func (o *NeynarPageInputText) GetDefaultEnabled() interface{} {
	return false
}

// GetPlaceholder returns the Placeholder field value if set, zero value otherwise.
func (o *NeynarPageInputText) GetPlaceholder() string {
	if o == nil || IsNil(o.Placeholder) {
		var ret string
		return ret
	}
	return *o.Placeholder
}

// GetPlaceholderOk returns a tuple with the Placeholder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NeynarPageInputText) GetPlaceholderOk() (*string, bool) {
	if o == nil || IsNil(o.Placeholder) {
		return nil, false
	}
	return o.Placeholder, true
}

// HasPlaceholder returns a boolean if a field has been set.
func (o *NeynarPageInputText) HasPlaceholder() bool {
	if o != nil && !IsNil(o.Placeholder) {
		return true
	}

	return false
}

// SetPlaceholder gets a reference to the given string and assigns it to the Placeholder field.
func (o *NeynarPageInputText) SetPlaceholder(v string) {
	o.Placeholder = &v
}

func (o NeynarPageInputText) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NeynarPageInputText) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if _, exists := toSerialize["enabled"]; !exists {
		toSerialize["enabled"] = o.GetDefaultEnabled()
	}
	toSerialize["enabled"] = o.Enabled
	if !IsNil(o.Placeholder) {
		toSerialize["placeholder"] = o.Placeholder
	}
	return toSerialize, nil
}

func (o *NeynarPageInputText) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"enabled",
	}

	// defaultValueFuncMap captures the default values for required properties.
	// These values are used when required properties are missing from the payload.
	defaultValueFuncMap := map[string]func() interface{}{
		"enabled": o.GetDefaultEnabled,
	}
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
	varNeynarPageInputText := _NeynarPageInputText{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varNeynarPageInputText)

	if err != nil {
		return err
	}

	*o = NeynarPageInputText(varNeynarPageInputText)

	return err
}

type NullableNeynarPageInputText struct {
	value *NeynarPageInputText
	isSet bool
}

func (v NullableNeynarPageInputText) Get() *NeynarPageInputText {
	return v.value
}

func (v *NullableNeynarPageInputText) Set(val *NeynarPageInputText) {
	v.value = val
	v.isSet = true
}

func (v NullableNeynarPageInputText) IsSet() bool {
	return v.isSet
}

func (v *NullableNeynarPageInputText) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNeynarPageInputText(val *NeynarPageInputText) *NullableNeynarPageInputText {
	return &NullableNeynarPageInputText{value: val, isSet: true}
}

func (v NullableNeynarPageInputText) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNeynarPageInputText) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
