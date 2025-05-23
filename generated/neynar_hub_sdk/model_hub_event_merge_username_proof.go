/*
Farcaster Hub API

Perform basic queries of Farcaster state via the REST API of a Farcaster hub. See the [Neynar docs](https://docs.neynar.com/reference) for more details.

API version: 2.35.0
Contact: team@neynar.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package neynar_hub_sdk

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the HubEventMergeUsernameProof type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HubEventMergeUsernameProof{}

// HubEventMergeUsernameProof struct for HubEventMergeUsernameProof
type HubEventMergeUsernameProof struct {
	Type                   string                 `json:"type"`
	Id                     int32                  `json:"id"`
	MergeUsernameProofBody MergeUserNameProofBody `json:"mergeUsernameProofBody"`
}

type _HubEventMergeUsernameProof HubEventMergeUsernameProof

// NewHubEventMergeUsernameProof instantiates a new HubEventMergeUsernameProof object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHubEventMergeUsernameProof(type_ string, id int32, mergeUsernameProofBody MergeUserNameProofBody) *HubEventMergeUsernameProof {
	this := HubEventMergeUsernameProof{}
	this.Type = type_
	this.Id = id
	this.MergeUsernameProofBody = mergeUsernameProofBody
	return &this
}

// NewHubEventMergeUsernameProofWithDefaults instantiates a new HubEventMergeUsernameProof object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHubEventMergeUsernameProofWithDefaults() *HubEventMergeUsernameProof {
	this := HubEventMergeUsernameProof{}
	return &this
}

// GetType returns the Type field value
func (o *HubEventMergeUsernameProof) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *HubEventMergeUsernameProof) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *HubEventMergeUsernameProof) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *HubEventMergeUsernameProof) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *HubEventMergeUsernameProof) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *HubEventMergeUsernameProof) SetId(v int32) {
	o.Id = v
}

// GetMergeUsernameProofBody returns the MergeUsernameProofBody field value
func (o *HubEventMergeUsernameProof) GetMergeUsernameProofBody() MergeUserNameProofBody {
	if o == nil {
		var ret MergeUserNameProofBody
		return ret
	}

	return o.MergeUsernameProofBody
}

// GetMergeUsernameProofBodyOk returns a tuple with the MergeUsernameProofBody field value
// and a boolean to check if the value has been set.
func (o *HubEventMergeUsernameProof) GetMergeUsernameProofBodyOk() (*MergeUserNameProofBody, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MergeUsernameProofBody, true
}

// SetMergeUsernameProofBody sets field value
func (o *HubEventMergeUsernameProof) SetMergeUsernameProofBody(v MergeUserNameProofBody) {
	o.MergeUsernameProofBody = v
}

func (o HubEventMergeUsernameProof) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HubEventMergeUsernameProof) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	toSerialize["mergeUsernameProofBody"] = o.MergeUsernameProofBody
	return toSerialize, nil
}

func (o *HubEventMergeUsernameProof) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"id",
		"mergeUsernameProofBody",
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
	varHubEventMergeUsernameProof := _HubEventMergeUsernameProof{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varHubEventMergeUsernameProof)

	if err != nil {
		return err
	}

	*o = HubEventMergeUsernameProof(varHubEventMergeUsernameProof)

	return err
}

type NullableHubEventMergeUsernameProof struct {
	value *HubEventMergeUsernameProof
	isSet bool
}

func (v NullableHubEventMergeUsernameProof) Get() *HubEventMergeUsernameProof {
	return v.value
}

func (v *NullableHubEventMergeUsernameProof) Set(val *HubEventMergeUsernameProof) {
	v.value = val
	v.isSet = true
}

func (v NullableHubEventMergeUsernameProof) IsSet() bool {
	return v.isSet
}

func (v *NullableHubEventMergeUsernameProof) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHubEventMergeUsernameProof(val *HubEventMergeUsernameProof) *NullableHubEventMergeUsernameProof {
	return &NullableHubEventMergeUsernameProof{value: val, isSet: true}
}

func (v NullableHubEventMergeUsernameProof) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHubEventMergeUsernameProof) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
