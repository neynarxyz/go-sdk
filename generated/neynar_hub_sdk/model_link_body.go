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

// checks if the LinkBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LinkBody{}

// LinkBody Defines a social connection between users in the Farcaster network. Currently used to establish following relationships, allowing users to curate their content feed.
type LinkBody struct {
	Type LinkType `json:"type"`
	// User-defined timestamp that preserves the original creation time when message.data.timestamp needs to be updated for compaction. Useful for maintaining chronological order in user feeds.
	DisplayTimestamp *int64 `json:"displayTimestamp,omitempty"`
	// Farcaster ID (FID). A unique identifier assigned to each user in the Farcaster network. This number is permanent and cannot be changed. FIDs are assigned sequentially when users register on the network.
	TargetFid int32 `json:"targetFid"`
}

type _LinkBody LinkBody

// NewLinkBody instantiates a new LinkBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkBody(type_ LinkType, targetFid int32) *LinkBody {
	this := LinkBody{}
	this.Type = type_
	this.TargetFid = targetFid
	return &this
}

// NewLinkBodyWithDefaults instantiates a new LinkBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkBodyWithDefaults() *LinkBody {
	this := LinkBody{}
	var type_ LinkType = LINKTYPE_FOLLOW
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *LinkBody) GetType() LinkType {
	if o == nil {
		var ret LinkType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *LinkBody) GetTypeOk() (*LinkType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *LinkBody) SetType(v LinkType) {
	o.Type = v
}

// GetDefaultType returns the default value LINKTYPE_FOLLOW of the Type field.
func (o *LinkBody) GetDefaultType() interface{} {
	return LINKTYPE_FOLLOW
}

// GetDisplayTimestamp returns the DisplayTimestamp field value if set, zero value otherwise.
func (o *LinkBody) GetDisplayTimestamp() int64 {
	if o == nil || IsNil(o.DisplayTimestamp) {
		var ret int64
		return ret
	}
	return *o.DisplayTimestamp
}

// GetDisplayTimestampOk returns a tuple with the DisplayTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkBody) GetDisplayTimestampOk() (*int64, bool) {
	if o == nil || IsNil(o.DisplayTimestamp) {
		return nil, false
	}
	return o.DisplayTimestamp, true
}

// HasDisplayTimestamp returns a boolean if a field has been set.
func (o *LinkBody) HasDisplayTimestamp() bool {
	if o != nil && !IsNil(o.DisplayTimestamp) {
		return true
	}

	return false
}

// SetDisplayTimestamp gets a reference to the given int64 and assigns it to the DisplayTimestamp field.
func (o *LinkBody) SetDisplayTimestamp(v int64) {
	o.DisplayTimestamp = &v
}

// GetTargetFid returns the TargetFid field value
func (o *LinkBody) GetTargetFid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TargetFid
}

// GetTargetFidOk returns a tuple with the TargetFid field value
// and a boolean to check if the value has been set.
func (o *LinkBody) GetTargetFidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetFid, true
}

// SetTargetFid sets field value
func (o *LinkBody) SetTargetFid(v int32) {
	o.TargetFid = v
}

func (o LinkBody) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LinkBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if _, exists := toSerialize["type"]; !exists {
		toSerialize["type"] = o.GetDefaultType()
	}
	toSerialize["type"] = o.Type
	if !IsNil(o.DisplayTimestamp) {
		toSerialize["displayTimestamp"] = o.DisplayTimestamp
	}
	toSerialize["targetFid"] = o.TargetFid
	return toSerialize, nil
}

func (o *LinkBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"targetFid",
	}

	// defaultValueFuncMap captures the default values for required properties.
	// These values are used when required properties are missing from the payload.
	defaultValueFuncMap := map[string]func() interface{}{
		"type": o.GetDefaultType,
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
	varLinkBody := _LinkBody{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varLinkBody)

	if err != nil {
		return err
	}

	*o = LinkBody(varLinkBody)

	return err
}

type NullableLinkBody struct {
	value *LinkBody
	isSet bool
}

func (v NullableLinkBody) Get() *LinkBody {
	return v.value
}

func (v *NullableLinkBody) Set(val *LinkBody) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkBody) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkBody(val *LinkBody) *NullableLinkBody {
	return &NullableLinkBody{value: val, isSet: true}
}

func (v NullableLinkBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
