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

// checks if the FollowResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FollowResponse{}

// FollowResponse struct for FollowResponse
type FollowResponse struct {
	Success bool `json:"success"`
	// The unique identifier of a farcaster user (unsigned integer)
	TargetFid int32  `json:"target_fid"`
	Hash      string `json:"hash"`
}

type _FollowResponse FollowResponse

// NewFollowResponse instantiates a new FollowResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFollowResponse(success bool, targetFid int32, hash string) *FollowResponse {
	this := FollowResponse{}
	this.Success = success
	this.TargetFid = targetFid
	this.Hash = hash
	return &this
}

// NewFollowResponseWithDefaults instantiates a new FollowResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFollowResponseWithDefaults() *FollowResponse {
	this := FollowResponse{}
	return &this
}

// GetSuccess returns the Success field value
func (o *FollowResponse) GetSuccess() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value
// and a boolean to check if the value has been set.
func (o *FollowResponse) GetSuccessOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Success, true
}

// SetSuccess sets field value
func (o *FollowResponse) SetSuccess(v bool) {
	o.Success = v
}

// GetTargetFid returns the TargetFid field value
func (o *FollowResponse) GetTargetFid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TargetFid
}

// GetTargetFidOk returns a tuple with the TargetFid field value
// and a boolean to check if the value has been set.
func (o *FollowResponse) GetTargetFidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetFid, true
}

// SetTargetFid sets field value
func (o *FollowResponse) SetTargetFid(v int32) {
	o.TargetFid = v
}

// GetHash returns the Hash field value
func (o *FollowResponse) GetHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hash
}

// GetHashOk returns a tuple with the Hash field value
// and a boolean to check if the value has been set.
func (o *FollowResponse) GetHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hash, true
}

// SetHash sets field value
func (o *FollowResponse) SetHash(v string) {
	o.Hash = v
}

func (o FollowResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FollowResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["success"] = o.Success
	toSerialize["target_fid"] = o.TargetFid
	toSerialize["hash"] = o.Hash
	return toSerialize, nil
}

func (o *FollowResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"success",
		"target_fid",
		"hash",
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
	varFollowResponse := _FollowResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFollowResponse)

	if err != nil {
		return err
	}

	*o = FollowResponse(varFollowResponse)

	return err
}

type NullableFollowResponse struct {
	value *FollowResponse
	isSet bool
}

func (v NullableFollowResponse) Get() *FollowResponse {
	return v.value
}

func (v *NullableFollowResponse) Set(val *FollowResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFollowResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFollowResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFollowResponse(val *FollowResponse) *NullableFollowResponse {
	return &NullableFollowResponse{value: val, isSet: true}
}

func (v NullableFollowResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFollowResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
