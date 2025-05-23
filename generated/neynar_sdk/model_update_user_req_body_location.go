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

// checks if the UpdateUserReqBodyLocation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateUserReqBodyLocation{}

// UpdateUserReqBodyLocation struct for UpdateUserReqBodyLocation
type UpdateUserReqBodyLocation struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type _UpdateUserReqBodyLocation UpdateUserReqBodyLocation

// NewUpdateUserReqBodyLocation instantiates a new UpdateUserReqBodyLocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateUserReqBodyLocation(latitude float64, longitude float64) *UpdateUserReqBodyLocation {
	this := UpdateUserReqBodyLocation{}
	this.Latitude = latitude
	this.Longitude = longitude
	return &this
}

// NewUpdateUserReqBodyLocationWithDefaults instantiates a new UpdateUserReqBodyLocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateUserReqBodyLocationWithDefaults() *UpdateUserReqBodyLocation {
	this := UpdateUserReqBodyLocation{}
	return &this
}

// GetLatitude returns the Latitude field value
func (o *UpdateUserReqBodyLocation) GetLatitude() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Latitude
}

// GetLatitudeOk returns a tuple with the Latitude field value
// and a boolean to check if the value has been set.
func (o *UpdateUserReqBodyLocation) GetLatitudeOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Latitude, true
}

// SetLatitude sets field value
func (o *UpdateUserReqBodyLocation) SetLatitude(v float64) {
	o.Latitude = v
}

// GetLongitude returns the Longitude field value
func (o *UpdateUserReqBodyLocation) GetLongitude() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Longitude
}

// GetLongitudeOk returns a tuple with the Longitude field value
// and a boolean to check if the value has been set.
func (o *UpdateUserReqBodyLocation) GetLongitudeOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Longitude, true
}

// SetLongitude sets field value
func (o *UpdateUserReqBodyLocation) SetLongitude(v float64) {
	o.Longitude = v
}

func (o UpdateUserReqBodyLocation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateUserReqBodyLocation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["latitude"] = o.Latitude
	toSerialize["longitude"] = o.Longitude
	return toSerialize, nil
}

func (o *UpdateUserReqBodyLocation) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"latitude",
		"longitude",
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
	varUpdateUserReqBodyLocation := _UpdateUserReqBodyLocation{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpdateUserReqBodyLocation)

	if err != nil {
		return err
	}

	*o = UpdateUserReqBodyLocation(varUpdateUserReqBodyLocation)

	return err
}

type NullableUpdateUserReqBodyLocation struct {
	value *UpdateUserReqBodyLocation
	isSet bool
}

func (v NullableUpdateUserReqBodyLocation) Get() *UpdateUserReqBodyLocation {
	return v.value
}

func (v *NullableUpdateUserReqBodyLocation) Set(val *UpdateUserReqBodyLocation) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateUserReqBodyLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateUserReqBodyLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateUserReqBodyLocation(val *UpdateUserReqBodyLocation) *NullableUpdateUserReqBodyLocation {
	return &NullableUpdateUserReqBodyLocation{value: val, isSet: true}
}

func (v NullableUpdateUserReqBodyLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateUserReqBodyLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
