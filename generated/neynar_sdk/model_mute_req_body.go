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

// checks if the MuteReqBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MuteReqBody{}

// MuteReqBody struct for MuteReqBody
type MuteReqBody struct {
	// The unique identifier of a farcaster user (unsigned integer)
	Fid int32 `json:"fid"`
	// The unique identifier of a farcaster user (unsigned integer)
	MutedFid int32 `json:"muted_fid"`
}

type _MuteReqBody MuteReqBody

// NewMuteReqBody instantiates a new MuteReqBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMuteReqBody(fid int32, mutedFid int32) *MuteReqBody {
	this := MuteReqBody{}
	this.Fid = fid
	this.MutedFid = mutedFid
	return &this
}

// NewMuteReqBodyWithDefaults instantiates a new MuteReqBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMuteReqBodyWithDefaults() *MuteReqBody {
	this := MuteReqBody{}
	return &this
}

// GetFid returns the Fid field value
func (o *MuteReqBody) GetFid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Fid
}

// GetFidOk returns a tuple with the Fid field value
// and a boolean to check if the value has been set.
func (o *MuteReqBody) GetFidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fid, true
}

// SetFid sets field value
func (o *MuteReqBody) SetFid(v int32) {
	o.Fid = v
}

// GetMutedFid returns the MutedFid field value
func (o *MuteReqBody) GetMutedFid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MutedFid
}

// GetMutedFidOk returns a tuple with the MutedFid field value
// and a boolean to check if the value has been set.
func (o *MuteReqBody) GetMutedFidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MutedFid, true
}

// SetMutedFid sets field value
func (o *MuteReqBody) SetMutedFid(v int32) {
	o.MutedFid = v
}

func (o MuteReqBody) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MuteReqBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fid"] = o.Fid
	toSerialize["muted_fid"] = o.MutedFid
	return toSerialize, nil
}

func (o *MuteReqBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fid",
		"muted_fid",
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
	varMuteReqBody := _MuteReqBody{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMuteReqBody)

	if err != nil {
		return err
	}

	*o = MuteReqBody(varMuteReqBody)

	return err
}

type NullableMuteReqBody struct {
	value *MuteReqBody
	isSet bool
}

func (v NullableMuteReqBody) Get() *MuteReqBody {
	return v.value
}

func (v *NullableMuteReqBody) Set(val *MuteReqBody) {
	v.value = val
	v.isSet = true
}

func (v NullableMuteReqBody) IsSet() bool {
	return v.isSet
}

func (v *NullableMuteReqBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMuteReqBody(val *MuteReqBody) *NullableMuteReqBody {
	return &NullableMuteReqBody{value: val, isSet: true}
}

func (v NullableMuteReqBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMuteReqBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
