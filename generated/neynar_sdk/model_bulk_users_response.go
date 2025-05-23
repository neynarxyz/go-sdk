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

// checks if the BulkUsersResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BulkUsersResponse{}

// BulkUsersResponse struct for BulkUsersResponse
type BulkUsersResponse struct {
	Users []User `json:"users"`
}

type _BulkUsersResponse BulkUsersResponse

// NewBulkUsersResponse instantiates a new BulkUsersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkUsersResponse(users []User) *BulkUsersResponse {
	this := BulkUsersResponse{}
	this.Users = users
	return &this
}

// NewBulkUsersResponseWithDefaults instantiates a new BulkUsersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkUsersResponseWithDefaults() *BulkUsersResponse {
	this := BulkUsersResponse{}
	return &this
}

// GetUsers returns the Users field value
func (o *BulkUsersResponse) GetUsers() []User {
	if o == nil {
		var ret []User
		return ret
	}

	return o.Users
}

// GetUsersOk returns a tuple with the Users field value
// and a boolean to check if the value has been set.
func (o *BulkUsersResponse) GetUsersOk() ([]User, bool) {
	if o == nil {
		return nil, false
	}
	return o.Users, true
}

// SetUsers sets field value
func (o *BulkUsersResponse) SetUsers(v []User) {
	o.Users = v
}

func (o BulkUsersResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BulkUsersResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["users"] = o.Users
	return toSerialize, nil
}

func (o *BulkUsersResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"users",
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
	varBulkUsersResponse := _BulkUsersResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBulkUsersResponse)

	if err != nil {
		return err
	}

	*o = BulkUsersResponse(varBulkUsersResponse)

	return err
}

type NullableBulkUsersResponse struct {
	value *BulkUsersResponse
	isSet bool
}

func (v NullableBulkUsersResponse) Get() *BulkUsersResponse {
	return v.value
}

func (v *NullableBulkUsersResponse) Set(val *BulkUsersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkUsersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkUsersResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkUsersResponse(val *BulkUsersResponse) *NullableBulkUsersResponse {
	return &NullableBulkUsersResponse{value: val, isSet: true}
}

func (v NullableBulkUsersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkUsersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
