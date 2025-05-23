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

// checks if the ProfileUrl type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProfileUrl{}

// ProfileUrl struct for ProfileUrl
type ProfileUrl struct {
	Pfp ProfileUrlPfp `json:"pfp"`
}

type _ProfileUrl ProfileUrl

// NewProfileUrl instantiates a new ProfileUrl object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProfileUrl(pfp ProfileUrlPfp) *ProfileUrl {
	this := ProfileUrl{}
	this.Pfp = pfp
	return &this
}

// NewProfileUrlWithDefaults instantiates a new ProfileUrl object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProfileUrlWithDefaults() *ProfileUrl {
	this := ProfileUrl{}
	return &this
}

// GetPfp returns the Pfp field value
func (o *ProfileUrl) GetPfp() ProfileUrlPfp {
	if o == nil {
		var ret ProfileUrlPfp
		return ret
	}

	return o.Pfp
}

// GetPfpOk returns a tuple with the Pfp field value
// and a boolean to check if the value has been set.
func (o *ProfileUrl) GetPfpOk() (*ProfileUrlPfp, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pfp, true
}

// SetPfp sets field value
func (o *ProfileUrl) SetPfp(v ProfileUrlPfp) {
	o.Pfp = v
}

func (o ProfileUrl) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProfileUrl) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pfp"] = o.Pfp
	return toSerialize, nil
}

func (o *ProfileUrl) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pfp",
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
	varProfileUrl := _ProfileUrl{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varProfileUrl)

	if err != nil {
		return err
	}

	*o = ProfileUrl(varProfileUrl)

	return err
}

type NullableProfileUrl struct {
	value *ProfileUrl
	isSet bool
}

func (v NullableProfileUrl) Get() *ProfileUrl {
	return v.value
}

func (v *NullableProfileUrl) Set(val *ProfileUrl) {
	v.value = val
	v.isSet = true
}

func (v NullableProfileUrl) IsSet() bool {
	return v.isSet
}

func (v *NullableProfileUrl) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProfileUrl(val *ProfileUrl) *NullableProfileUrl {
	return &NullableProfileUrl{value: val, isSet: true}
}

func (v NullableProfileUrl) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProfileUrl) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
