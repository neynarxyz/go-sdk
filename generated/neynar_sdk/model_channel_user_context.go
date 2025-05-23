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

// checks if the ChannelUserContext type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChannelUserContext{}

// ChannelUserContext Adds context on the viewer's or author's role in the channel.
type ChannelUserContext struct {
	// Indicates if the user is following the channel.
	Following bool               `json:"following"`
	Role      *ChannelMemberRole `json:"role,omitempty"`
}

type _ChannelUserContext ChannelUserContext

// NewChannelUserContext instantiates a new ChannelUserContext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelUserContext(following bool) *ChannelUserContext {
	this := ChannelUserContext{}
	this.Following = following
	return &this
}

// NewChannelUserContextWithDefaults instantiates a new ChannelUserContext object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelUserContextWithDefaults() *ChannelUserContext {
	this := ChannelUserContext{}
	return &this
}

// GetFollowing returns the Following field value
func (o *ChannelUserContext) GetFollowing() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Following
}

// GetFollowingOk returns a tuple with the Following field value
// and a boolean to check if the value has been set.
func (o *ChannelUserContext) GetFollowingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Following, true
}

// SetFollowing sets field value
func (o *ChannelUserContext) SetFollowing(v bool) {
	o.Following = v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *ChannelUserContext) GetRole() ChannelMemberRole {
	if o == nil || IsNil(o.Role) {
		var ret ChannelMemberRole
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelUserContext) GetRoleOk() (*ChannelMemberRole, bool) {
	if o == nil || IsNil(o.Role) {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *ChannelUserContext) HasRole() bool {
	if o != nil && !IsNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given ChannelMemberRole and assigns it to the Role field.
func (o *ChannelUserContext) SetRole(v ChannelMemberRole) {
	o.Role = &v
}

func (o ChannelUserContext) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChannelUserContext) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["following"] = o.Following
	if !IsNil(o.Role) {
		toSerialize["role"] = o.Role
	}
	return toSerialize, nil
}

func (o *ChannelUserContext) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"following",
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
	varChannelUserContext := _ChannelUserContext{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varChannelUserContext)

	if err != nil {
		return err
	}

	*o = ChannelUserContext(varChannelUserContext)

	return err
}

type NullableChannelUserContext struct {
	value *ChannelUserContext
	isSet bool
}

func (v NullableChannelUserContext) Get() *ChannelUserContext {
	return v.value
}

func (v *NullableChannelUserContext) Set(val *ChannelUserContext) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelUserContext) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelUserContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelUserContext(val *ChannelUserContext) *NullableChannelUserContext {
	return &NullableChannelUserContext{value: val, isSet: true}
}

func (v NullableChannelUserContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelUserContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
