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

// checks if the FrameValidateAnalyticsInputText type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FrameValidateAnalyticsInputText{}

// FrameValidateAnalyticsInputText struct for FrameValidateAnalyticsInputText
type FrameValidateAnalyticsInputText struct {
	InputTexts []FrameValidateAnalyticsInputTextInputTextsInner `json:"input_texts"`
}

type _FrameValidateAnalyticsInputText FrameValidateAnalyticsInputText

// NewFrameValidateAnalyticsInputText instantiates a new FrameValidateAnalyticsInputText object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFrameValidateAnalyticsInputText(inputTexts []FrameValidateAnalyticsInputTextInputTextsInner) *FrameValidateAnalyticsInputText {
	this := FrameValidateAnalyticsInputText{}
	this.InputTexts = inputTexts
	return &this
}

// NewFrameValidateAnalyticsInputTextWithDefaults instantiates a new FrameValidateAnalyticsInputText object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFrameValidateAnalyticsInputTextWithDefaults() *FrameValidateAnalyticsInputText {
	this := FrameValidateAnalyticsInputText{}
	return &this
}

// GetInputTexts returns the InputTexts field value
func (o *FrameValidateAnalyticsInputText) GetInputTexts() []FrameValidateAnalyticsInputTextInputTextsInner {
	if o == nil {
		var ret []FrameValidateAnalyticsInputTextInputTextsInner
		return ret
	}

	return o.InputTexts
}

// GetInputTextsOk returns a tuple with the InputTexts field value
// and a boolean to check if the value has been set.
func (o *FrameValidateAnalyticsInputText) GetInputTextsOk() ([]FrameValidateAnalyticsInputTextInputTextsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.InputTexts, true
}

// SetInputTexts sets field value
func (o *FrameValidateAnalyticsInputText) SetInputTexts(v []FrameValidateAnalyticsInputTextInputTextsInner) {
	o.InputTexts = v
}

func (o FrameValidateAnalyticsInputText) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FrameValidateAnalyticsInputText) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["input_texts"] = o.InputTexts
	return toSerialize, nil
}

func (o *FrameValidateAnalyticsInputText) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"input_texts",
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
	varFrameValidateAnalyticsInputText := _FrameValidateAnalyticsInputText{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFrameValidateAnalyticsInputText)

	if err != nil {
		return err
	}

	*o = FrameValidateAnalyticsInputText(varFrameValidateAnalyticsInputText)

	return err
}

type NullableFrameValidateAnalyticsInputText struct {
	value *FrameValidateAnalyticsInputText
	isSet bool
}

func (v NullableFrameValidateAnalyticsInputText) Get() *FrameValidateAnalyticsInputText {
	return v.value
}

func (v *NullableFrameValidateAnalyticsInputText) Set(val *FrameValidateAnalyticsInputText) {
	v.value = val
	v.isSet = true
}

func (v NullableFrameValidateAnalyticsInputText) IsSet() bool {
	return v.isSet
}

func (v *NullableFrameValidateAnalyticsInputText) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFrameValidateAnalyticsInputText(val *FrameValidateAnalyticsInputText) *NullableFrameValidateAnalyticsInputText {
	return &NullableFrameValidateAnalyticsInputText{value: val, isSet: true}
}

func (v NullableFrameValidateAnalyticsInputText) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFrameValidateAnalyticsInputText) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
