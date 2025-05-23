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

// checks if the RelevantFollowersResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RelevantFollowersResponse{}

// RelevantFollowersResponse struct for RelevantFollowersResponse
type RelevantFollowersResponse struct {
	TopRelevantFollowersHydrated   []Follower           `json:"top_relevant_followers_hydrated"`
	AllRelevantFollowersDehydrated []FollowerDehydrated `json:"all_relevant_followers_dehydrated"`
}

type _RelevantFollowersResponse RelevantFollowersResponse

// NewRelevantFollowersResponse instantiates a new RelevantFollowersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelevantFollowersResponse(topRelevantFollowersHydrated []Follower, allRelevantFollowersDehydrated []FollowerDehydrated) *RelevantFollowersResponse {
	this := RelevantFollowersResponse{}
	this.TopRelevantFollowersHydrated = topRelevantFollowersHydrated
	this.AllRelevantFollowersDehydrated = allRelevantFollowersDehydrated
	return &this
}

// NewRelevantFollowersResponseWithDefaults instantiates a new RelevantFollowersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelevantFollowersResponseWithDefaults() *RelevantFollowersResponse {
	this := RelevantFollowersResponse{}
	return &this
}

// GetTopRelevantFollowersHydrated returns the TopRelevantFollowersHydrated field value
func (o *RelevantFollowersResponse) GetTopRelevantFollowersHydrated() []Follower {
	if o == nil {
		var ret []Follower
		return ret
	}

	return o.TopRelevantFollowersHydrated
}

// GetTopRelevantFollowersHydratedOk returns a tuple with the TopRelevantFollowersHydrated field value
// and a boolean to check if the value has been set.
func (o *RelevantFollowersResponse) GetTopRelevantFollowersHydratedOk() ([]Follower, bool) {
	if o == nil {
		return nil, false
	}
	return o.TopRelevantFollowersHydrated, true
}

// SetTopRelevantFollowersHydrated sets field value
func (o *RelevantFollowersResponse) SetTopRelevantFollowersHydrated(v []Follower) {
	o.TopRelevantFollowersHydrated = v
}

// GetAllRelevantFollowersDehydrated returns the AllRelevantFollowersDehydrated field value
func (o *RelevantFollowersResponse) GetAllRelevantFollowersDehydrated() []FollowerDehydrated {
	if o == nil {
		var ret []FollowerDehydrated
		return ret
	}

	return o.AllRelevantFollowersDehydrated
}

// GetAllRelevantFollowersDehydratedOk returns a tuple with the AllRelevantFollowersDehydrated field value
// and a boolean to check if the value has been set.
func (o *RelevantFollowersResponse) GetAllRelevantFollowersDehydratedOk() ([]FollowerDehydrated, bool) {
	if o == nil {
		return nil, false
	}
	return o.AllRelevantFollowersDehydrated, true
}

// SetAllRelevantFollowersDehydrated sets field value
func (o *RelevantFollowersResponse) SetAllRelevantFollowersDehydrated(v []FollowerDehydrated) {
	o.AllRelevantFollowersDehydrated = v
}

func (o RelevantFollowersResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RelevantFollowersResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["top_relevant_followers_hydrated"] = o.TopRelevantFollowersHydrated
	toSerialize["all_relevant_followers_dehydrated"] = o.AllRelevantFollowersDehydrated
	return toSerialize, nil
}

func (o *RelevantFollowersResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"top_relevant_followers_hydrated",
		"all_relevant_followers_dehydrated",
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
	varRelevantFollowersResponse := _RelevantFollowersResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRelevantFollowersResponse)

	if err != nil {
		return err
	}

	*o = RelevantFollowersResponse(varRelevantFollowersResponse)

	return err
}

type NullableRelevantFollowersResponse struct {
	value *RelevantFollowersResponse
	isSet bool
}

func (v NullableRelevantFollowersResponse) Get() *RelevantFollowersResponse {
	return v.value
}

func (v *NullableRelevantFollowersResponse) Set(val *RelevantFollowersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRelevantFollowersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRelevantFollowersResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRelevantFollowersResponse(val *RelevantFollowersResponse) *NullableRelevantFollowersResponse {
	return &NullableRelevantFollowersResponse{value: val, isSet: true}
}

func (v NullableRelevantFollowersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRelevantFollowersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
