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

// checks if the ChannelResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChannelResponse{}

// ChannelResponse struct for ChannelResponse
type ChannelResponse struct {
	Channel Channel `json:"channel"`
}

type _ChannelResponse ChannelResponse

// NewChannelResponse instantiates a new ChannelResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelResponse(channel Channel) *ChannelResponse {
	this := ChannelResponse{}
	this.Channel = channel
	return &this
}

// NewChannelResponseWithDefaults instantiates a new ChannelResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelResponseWithDefaults() *ChannelResponse {
	this := ChannelResponse{}
	return &this
}

// GetChannel returns the Channel field value
func (o *ChannelResponse) GetChannel() Channel {
	if o == nil {
		var ret Channel
		return ret
	}

	return o.Channel
}

// GetChannelOk returns a tuple with the Channel field value
// and a boolean to check if the value has been set.
func (o *ChannelResponse) GetChannelOk() (*Channel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Channel, true
}

// SetChannel sets field value
func (o *ChannelResponse) SetChannel(v Channel) {
	o.Channel = v
}

func (o ChannelResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChannelResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["channel"] = o.Channel
	return toSerialize, nil
}

func (o *ChannelResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"channel",
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
	varChannelResponse := _ChannelResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varChannelResponse)

	if err != nil {
		return err
	}

	*o = ChannelResponse(varChannelResponse)

	return err
}

type NullableChannelResponse struct {
	value *ChannelResponse
	isSet bool
}

func (v NullableChannelResponse) Get() *ChannelResponse {
	return v.value
}

func (v *NullableChannelResponse) Set(val *ChannelResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelResponse(val *ChannelResponse) *NullableChannelResponse {
	return &NullableChannelResponse{value: val, isSet: true}
}

func (v NullableChannelResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
