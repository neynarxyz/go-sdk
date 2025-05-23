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

// checks if the UserSearchResponseResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserSearchResponseResult{}

// UserSearchResponseResult struct for UserSearchResponseResult
type UserSearchResponseResult struct {
	Users []SearchedUser `json:"users"`
	Next  *NextCursor    `json:"next,omitempty"`
}

type _UserSearchResponseResult UserSearchResponseResult

// NewUserSearchResponseResult instantiates a new UserSearchResponseResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserSearchResponseResult(users []SearchedUser) *UserSearchResponseResult {
	this := UserSearchResponseResult{}
	this.Users = users
	return &this
}

// NewUserSearchResponseResultWithDefaults instantiates a new UserSearchResponseResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserSearchResponseResultWithDefaults() *UserSearchResponseResult {
	this := UserSearchResponseResult{}
	return &this
}

// GetUsers returns the Users field value
func (o *UserSearchResponseResult) GetUsers() []SearchedUser {
	if o == nil {
		var ret []SearchedUser
		return ret
	}

	return o.Users
}

// GetUsersOk returns a tuple with the Users field value
// and a boolean to check if the value has been set.
func (o *UserSearchResponseResult) GetUsersOk() ([]SearchedUser, bool) {
	if o == nil {
		return nil, false
	}
	return o.Users, true
}

// SetUsers sets field value
func (o *UserSearchResponseResult) SetUsers(v []SearchedUser) {
	o.Users = v
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *UserSearchResponseResult) GetNext() NextCursor {
	if o == nil || IsNil(o.Next) {
		var ret NextCursor
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSearchResponseResult) GetNextOk() (*NextCursor, bool) {
	if o == nil || IsNil(o.Next) {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *UserSearchResponseResult) HasNext() bool {
	if o != nil && !IsNil(o.Next) {
		return true
	}

	return false
}

// SetNext gets a reference to the given NextCursor and assigns it to the Next field.
func (o *UserSearchResponseResult) SetNext(v NextCursor) {
	o.Next = &v
}

func (o UserSearchResponseResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserSearchResponseResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["users"] = o.Users
	if !IsNil(o.Next) {
		toSerialize["next"] = o.Next
	}
	return toSerialize, nil
}

func (o *UserSearchResponseResult) UnmarshalJSON(data []byte) (err error) {
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
	varUserSearchResponseResult := _UserSearchResponseResult{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserSearchResponseResult)

	if err != nil {
		return err
	}

	*o = UserSearchResponseResult(varUserSearchResponseResult)

	return err
}

type NullableUserSearchResponseResult struct {
	value *UserSearchResponseResult
	isSet bool
}

func (v NullableUserSearchResponseResult) Get() *UserSearchResponseResult {
	return v.value
}

func (v *NullableUserSearchResponseResult) Set(val *UserSearchResponseResult) {
	v.value = val
	v.isSet = true
}

func (v NullableUserSearchResponseResult) IsSet() bool {
	return v.isSet
}

func (v *NullableUserSearchResponseResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserSearchResponseResult(val *UserSearchResponseResult) *NullableUserSearchResponseResult {
	return &NullableUserSearchResponseResult{value: val, isSet: true}
}

func (v NullableUserSearchResponseResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserSearchResponseResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
