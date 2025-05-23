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

// checks if the UsernameProofsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UsernameProofsResponse{}

// UsernameProofsResponse struct for UsernameProofsResponse
type UsernameProofsResponse struct {
	Proofs []UserNameProof `json:"proofs"`
}

type _UsernameProofsResponse UsernameProofsResponse

// NewUsernameProofsResponse instantiates a new UsernameProofsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsernameProofsResponse(proofs []UserNameProof) *UsernameProofsResponse {
	this := UsernameProofsResponse{}
	this.Proofs = proofs
	return &this
}

// NewUsernameProofsResponseWithDefaults instantiates a new UsernameProofsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsernameProofsResponseWithDefaults() *UsernameProofsResponse {
	this := UsernameProofsResponse{}
	return &this
}

// GetProofs returns the Proofs field value
func (o *UsernameProofsResponse) GetProofs() []UserNameProof {
	if o == nil {
		var ret []UserNameProof
		return ret
	}

	return o.Proofs
}

// GetProofsOk returns a tuple with the Proofs field value
// and a boolean to check if the value has been set.
func (o *UsernameProofsResponse) GetProofsOk() ([]UserNameProof, bool) {
	if o == nil {
		return nil, false
	}
	return o.Proofs, true
}

// SetProofs sets field value
func (o *UsernameProofsResponse) SetProofs(v []UserNameProof) {
	o.Proofs = v
}

func (o UsernameProofsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UsernameProofsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["proofs"] = o.Proofs
	return toSerialize, nil
}

func (o *UsernameProofsResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"proofs",
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
	varUsernameProofsResponse := _UsernameProofsResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUsernameProofsResponse)

	if err != nil {
		return err
	}

	*o = UsernameProofsResponse(varUsernameProofsResponse)

	return err
}

type NullableUsernameProofsResponse struct {
	value *UsernameProofsResponse
	isSet bool
}

func (v NullableUsernameProofsResponse) Get() *UsernameProofsResponse {
	return v.value
}

func (v *NullableUsernameProofsResponse) Set(val *UsernameProofsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUsernameProofsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUsernameProofsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsernameProofsResponse(val *UsernameProofsResponse) *NullableUsernameProofsResponse {
	return &NullableUsernameProofsResponse{value: val, isSet: true}
}

func (v NullableUsernameProofsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsernameProofsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
