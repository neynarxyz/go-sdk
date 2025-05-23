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

// checks if the FetchVerificationsByFid200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FetchVerificationsByFid200Response{}

// FetchVerificationsByFid200Response struct for FetchVerificationsByFid200Response
type FetchVerificationsByFid200Response struct {
	Messages      []Verification `json:"messages"`
	NextPageToken string         `json:"nextPageToken" validate:"regexp=^(?:[A-Za-z0-9+\\/]{4})*(?:[A-Za-z0-9+\\/]{2}==|[A-Za-z0-9+\\/]{3}=)?$"`
}

type _FetchVerificationsByFid200Response FetchVerificationsByFid200Response

// NewFetchVerificationsByFid200Response instantiates a new FetchVerificationsByFid200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFetchVerificationsByFid200Response(messages []Verification, nextPageToken string) *FetchVerificationsByFid200Response {
	this := FetchVerificationsByFid200Response{}
	this.Messages = messages
	this.NextPageToken = nextPageToken
	return &this
}

// NewFetchVerificationsByFid200ResponseWithDefaults instantiates a new FetchVerificationsByFid200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFetchVerificationsByFid200ResponseWithDefaults() *FetchVerificationsByFid200Response {
	this := FetchVerificationsByFid200Response{}
	return &this
}

// GetMessages returns the Messages field value
func (o *FetchVerificationsByFid200Response) GetMessages() []Verification {
	if o == nil {
		var ret []Verification
		return ret
	}

	return o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value
// and a boolean to check if the value has been set.
func (o *FetchVerificationsByFid200Response) GetMessagesOk() ([]Verification, bool) {
	if o == nil {
		return nil, false
	}
	return o.Messages, true
}

// SetMessages sets field value
func (o *FetchVerificationsByFid200Response) SetMessages(v []Verification) {
	o.Messages = v
}

// GetNextPageToken returns the NextPageToken field value
func (o *FetchVerificationsByFid200Response) GetNextPageToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value
// and a boolean to check if the value has been set.
func (o *FetchVerificationsByFid200Response) GetNextPageTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NextPageToken, true
}

// SetNextPageToken sets field value
func (o *FetchVerificationsByFid200Response) SetNextPageToken(v string) {
	o.NextPageToken = v
}

func (o FetchVerificationsByFid200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FetchVerificationsByFid200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["messages"] = o.Messages
	toSerialize["nextPageToken"] = o.NextPageToken
	return toSerialize, nil
}

func (o *FetchVerificationsByFid200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"messages",
		"nextPageToken",
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
	varFetchVerificationsByFid200Response := _FetchVerificationsByFid200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFetchVerificationsByFid200Response)

	if err != nil {
		return err
	}

	*o = FetchVerificationsByFid200Response(varFetchVerificationsByFid200Response)

	return err
}

type NullableFetchVerificationsByFid200Response struct {
	value *FetchVerificationsByFid200Response
	isSet bool
}

func (v NullableFetchVerificationsByFid200Response) Get() *FetchVerificationsByFid200Response {
	return v.value
}

func (v *NullableFetchVerificationsByFid200Response) Set(val *FetchVerificationsByFid200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableFetchVerificationsByFid200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableFetchVerificationsByFid200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFetchVerificationsByFid200Response(val *FetchVerificationsByFid200Response) *NullableFetchVerificationsByFid200Response {
	return &NullableFetchVerificationsByFid200Response{value: val, isSet: true}
}

func (v NullableFetchVerificationsByFid200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFetchVerificationsByFid200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
