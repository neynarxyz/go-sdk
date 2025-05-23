/*
Farcaster Hub API

Perform basic queries of Farcaster state via the REST API of a Farcaster hub. See the [Neynar docs](https://docs.neynar.com/reference) for more details.

API version: 2.35.0
Contact: team@neynar.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package neynar_hub_sdk

import (
	"encoding/json"
)

// checks if the MergeUserNameProofBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MergeUserNameProofBody{}

// MergeUserNameProofBody struct for MergeUserNameProofBody
type MergeUserNameProofBody struct {
	UsernameProof               *UserNameProof `json:"usernameProof,omitempty"`
	DeletedUsernameProof        *UserNameProof `json:"deletedUsernameProof,omitempty"`
	UsernameProofMessage        *Message       `json:"usernameProofMessage,omitempty"`
	DeletedUsernameProofMessage *Message       `json:"deletedUsernameProofMessage,omitempty"`
}

// NewMergeUserNameProofBody instantiates a new MergeUserNameProofBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMergeUserNameProofBody() *MergeUserNameProofBody {
	this := MergeUserNameProofBody{}
	return &this
}

// NewMergeUserNameProofBodyWithDefaults instantiates a new MergeUserNameProofBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMergeUserNameProofBodyWithDefaults() *MergeUserNameProofBody {
	this := MergeUserNameProofBody{}
	return &this
}

// GetUsernameProof returns the UsernameProof field value if set, zero value otherwise.
func (o *MergeUserNameProofBody) GetUsernameProof() UserNameProof {
	if o == nil || IsNil(o.UsernameProof) {
		var ret UserNameProof
		return ret
	}
	return *o.UsernameProof
}

// GetUsernameProofOk returns a tuple with the UsernameProof field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MergeUserNameProofBody) GetUsernameProofOk() (*UserNameProof, bool) {
	if o == nil || IsNil(o.UsernameProof) {
		return nil, false
	}
	return o.UsernameProof, true
}

// HasUsernameProof returns a boolean if a field has been set.
func (o *MergeUserNameProofBody) HasUsernameProof() bool {
	if o != nil && !IsNil(o.UsernameProof) {
		return true
	}

	return false
}

// SetUsernameProof gets a reference to the given UserNameProof and assigns it to the UsernameProof field.
func (o *MergeUserNameProofBody) SetUsernameProof(v UserNameProof) {
	o.UsernameProof = &v
}

// GetDeletedUsernameProof returns the DeletedUsernameProof field value if set, zero value otherwise.
func (o *MergeUserNameProofBody) GetDeletedUsernameProof() UserNameProof {
	if o == nil || IsNil(o.DeletedUsernameProof) {
		var ret UserNameProof
		return ret
	}
	return *o.DeletedUsernameProof
}

// GetDeletedUsernameProofOk returns a tuple with the DeletedUsernameProof field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MergeUserNameProofBody) GetDeletedUsernameProofOk() (*UserNameProof, bool) {
	if o == nil || IsNil(o.DeletedUsernameProof) {
		return nil, false
	}
	return o.DeletedUsernameProof, true
}

// HasDeletedUsernameProof returns a boolean if a field has been set.
func (o *MergeUserNameProofBody) HasDeletedUsernameProof() bool {
	if o != nil && !IsNil(o.DeletedUsernameProof) {
		return true
	}

	return false
}

// SetDeletedUsernameProof gets a reference to the given UserNameProof and assigns it to the DeletedUsernameProof field.
func (o *MergeUserNameProofBody) SetDeletedUsernameProof(v UserNameProof) {
	o.DeletedUsernameProof = &v
}

// GetUsernameProofMessage returns the UsernameProofMessage field value if set, zero value otherwise.
func (o *MergeUserNameProofBody) GetUsernameProofMessage() Message {
	if o == nil || IsNil(o.UsernameProofMessage) {
		var ret Message
		return ret
	}
	return *o.UsernameProofMessage
}

// GetUsernameProofMessageOk returns a tuple with the UsernameProofMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MergeUserNameProofBody) GetUsernameProofMessageOk() (*Message, bool) {
	if o == nil || IsNil(o.UsernameProofMessage) {
		return nil, false
	}
	return o.UsernameProofMessage, true
}

// HasUsernameProofMessage returns a boolean if a field has been set.
func (o *MergeUserNameProofBody) HasUsernameProofMessage() bool {
	if o != nil && !IsNil(o.UsernameProofMessage) {
		return true
	}

	return false
}

// SetUsernameProofMessage gets a reference to the given Message and assigns it to the UsernameProofMessage field.
func (o *MergeUserNameProofBody) SetUsernameProofMessage(v Message) {
	o.UsernameProofMessage = &v
}

// GetDeletedUsernameProofMessage returns the DeletedUsernameProofMessage field value if set, zero value otherwise.
func (o *MergeUserNameProofBody) GetDeletedUsernameProofMessage() Message {
	if o == nil || IsNil(o.DeletedUsernameProofMessage) {
		var ret Message
		return ret
	}
	return *o.DeletedUsernameProofMessage
}

// GetDeletedUsernameProofMessageOk returns a tuple with the DeletedUsernameProofMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MergeUserNameProofBody) GetDeletedUsernameProofMessageOk() (*Message, bool) {
	if o == nil || IsNil(o.DeletedUsernameProofMessage) {
		return nil, false
	}
	return o.DeletedUsernameProofMessage, true
}

// HasDeletedUsernameProofMessage returns a boolean if a field has been set.
func (o *MergeUserNameProofBody) HasDeletedUsernameProofMessage() bool {
	if o != nil && !IsNil(o.DeletedUsernameProofMessage) {
		return true
	}

	return false
}

// SetDeletedUsernameProofMessage gets a reference to the given Message and assigns it to the DeletedUsernameProofMessage field.
func (o *MergeUserNameProofBody) SetDeletedUsernameProofMessage(v Message) {
	o.DeletedUsernameProofMessage = &v
}

func (o MergeUserNameProofBody) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MergeUserNameProofBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UsernameProof) {
		toSerialize["usernameProof"] = o.UsernameProof
	}
	if !IsNil(o.DeletedUsernameProof) {
		toSerialize["deletedUsernameProof"] = o.DeletedUsernameProof
	}
	if !IsNil(o.UsernameProofMessage) {
		toSerialize["usernameProofMessage"] = o.UsernameProofMessage
	}
	if !IsNil(o.DeletedUsernameProofMessage) {
		toSerialize["deletedUsernameProofMessage"] = o.DeletedUsernameProofMessage
	}
	return toSerialize, nil
}

type NullableMergeUserNameProofBody struct {
	value *MergeUserNameProofBody
	isSet bool
}

func (v NullableMergeUserNameProofBody) Get() *MergeUserNameProofBody {
	return v.value
}

func (v *NullableMergeUserNameProofBody) Set(val *MergeUserNameProofBody) {
	v.value = val
	v.isSet = true
}

func (v NullableMergeUserNameProofBody) IsSet() bool {
	return v.isSet
}

func (v *NullableMergeUserNameProofBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMergeUserNameProofBody(val *MergeUserNameProofBody) *NullableMergeUserNameProofBody {
	return &NullableMergeUserNameProofBody{value: val, isSet: true}
}

func (v NullableMergeUserNameProofBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMergeUserNameProofBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
