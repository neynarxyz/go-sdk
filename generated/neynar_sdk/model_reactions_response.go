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

// checks if the ReactionsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReactionsResponse{}

// ReactionsResponse struct for ReactionsResponse
type ReactionsResponse struct {
	Reactions []ReactionWithCastInfo `json:"reactions"`
	Next      NextCursor             `json:"next"`
}

type _ReactionsResponse ReactionsResponse

// NewReactionsResponse instantiates a new ReactionsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReactionsResponse(reactions []ReactionWithCastInfo, next NextCursor) *ReactionsResponse {
	this := ReactionsResponse{}
	this.Reactions = reactions
	this.Next = next
	return &this
}

// NewReactionsResponseWithDefaults instantiates a new ReactionsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReactionsResponseWithDefaults() *ReactionsResponse {
	this := ReactionsResponse{}
	return &this
}

// GetReactions returns the Reactions field value
func (o *ReactionsResponse) GetReactions() []ReactionWithCastInfo {
	if o == nil {
		var ret []ReactionWithCastInfo
		return ret
	}

	return o.Reactions
}

// GetReactionsOk returns a tuple with the Reactions field value
// and a boolean to check if the value has been set.
func (o *ReactionsResponse) GetReactionsOk() ([]ReactionWithCastInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.Reactions, true
}

// SetReactions sets field value
func (o *ReactionsResponse) SetReactions(v []ReactionWithCastInfo) {
	o.Reactions = v
}

// GetNext returns the Next field value
func (o *ReactionsResponse) GetNext() NextCursor {
	if o == nil {
		var ret NextCursor
		return ret
	}

	return o.Next
}

// GetNextOk returns a tuple with the Next field value
// and a boolean to check if the value has been set.
func (o *ReactionsResponse) GetNextOk() (*NextCursor, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Next, true
}

// SetNext sets field value
func (o *ReactionsResponse) SetNext(v NextCursor) {
	o.Next = v
}

func (o ReactionsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReactionsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["reactions"] = o.Reactions
	toSerialize["next"] = o.Next
	return toSerialize, nil
}

func (o *ReactionsResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"reactions",
		"next",
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
	varReactionsResponse := _ReactionsResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varReactionsResponse)

	if err != nil {
		return err
	}

	*o = ReactionsResponse(varReactionsResponse)

	return err
}

type NullableReactionsResponse struct {
	value *ReactionsResponse
	isSet bool
}

func (v NullableReactionsResponse) Get() *ReactionsResponse {
	return v.value
}

func (v *NullableReactionsResponse) Set(val *ReactionsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableReactionsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableReactionsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReactionsResponse(val *ReactionsResponse) *NullableReactionsResponse {
	return &NullableReactionsResponse{value: val, isSet: true}
}

func (v NullableReactionsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReactionsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
