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

// checks if the FetchRelevantFrames200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FetchRelevantFrames200Response{}

// FetchRelevantFrames200Response struct for FetchRelevantFrames200Response
type FetchRelevantFrames200Response struct {
	RelevantFrames []FetchRelevantFrames200ResponseRelevantFramesInner `json:"relevant_frames"`
}

type _FetchRelevantFrames200Response FetchRelevantFrames200Response

// NewFetchRelevantFrames200Response instantiates a new FetchRelevantFrames200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFetchRelevantFrames200Response(relevantFrames []FetchRelevantFrames200ResponseRelevantFramesInner) *FetchRelevantFrames200Response {
	this := FetchRelevantFrames200Response{}
	this.RelevantFrames = relevantFrames
	return &this
}

// NewFetchRelevantFrames200ResponseWithDefaults instantiates a new FetchRelevantFrames200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFetchRelevantFrames200ResponseWithDefaults() *FetchRelevantFrames200Response {
	this := FetchRelevantFrames200Response{}
	return &this
}

// GetRelevantFrames returns the RelevantFrames field value
func (o *FetchRelevantFrames200Response) GetRelevantFrames() []FetchRelevantFrames200ResponseRelevantFramesInner {
	if o == nil {
		var ret []FetchRelevantFrames200ResponseRelevantFramesInner
		return ret
	}

	return o.RelevantFrames
}

// GetRelevantFramesOk returns a tuple with the RelevantFrames field value
// and a boolean to check if the value has been set.
func (o *FetchRelevantFrames200Response) GetRelevantFramesOk() ([]FetchRelevantFrames200ResponseRelevantFramesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.RelevantFrames, true
}

// SetRelevantFrames sets field value
func (o *FetchRelevantFrames200Response) SetRelevantFrames(v []FetchRelevantFrames200ResponseRelevantFramesInner) {
	o.RelevantFrames = v
}

func (o FetchRelevantFrames200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FetchRelevantFrames200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["relevant_frames"] = o.RelevantFrames
	return toSerialize, nil
}

func (o *FetchRelevantFrames200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"relevant_frames",
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
	varFetchRelevantFrames200Response := _FetchRelevantFrames200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFetchRelevantFrames200Response)

	if err != nil {
		return err
	}

	*o = FetchRelevantFrames200Response(varFetchRelevantFrames200Response)

	return err
}

type NullableFetchRelevantFrames200Response struct {
	value *FetchRelevantFrames200Response
	isSet bool
}

func (v NullableFetchRelevantFrames200Response) Get() *FetchRelevantFrames200Response {
	return v.value
}

func (v *NullableFetchRelevantFrames200Response) Set(val *FetchRelevantFrames200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableFetchRelevantFrames200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableFetchRelevantFrames200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFetchRelevantFrames200Response(val *FetchRelevantFrames200Response) *NullableFetchRelevantFrames200Response {
	return &NullableFetchRelevantFrames200Response{value: val, isSet: true}
}

func (v NullableFetchRelevantFrames200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFetchRelevantFrames200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
