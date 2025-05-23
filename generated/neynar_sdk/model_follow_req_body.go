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

// checks if the FollowReqBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FollowReqBody{}

// FollowReqBody struct for FollowReqBody
type FollowReqBody struct {
	// UUID of the signer. `signer_uuid` is paired with API key, can't use a `uuid` made with a different API key.
	SignerUuid string  `json:"signer_uuid"`
	TargetFids []int32 `json:"target_fids"`
}

type _FollowReqBody FollowReqBody

// NewFollowReqBody instantiates a new FollowReqBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFollowReqBody(signerUuid string, targetFids []int32) *FollowReqBody {
	this := FollowReqBody{}
	this.SignerUuid = signerUuid
	this.TargetFids = targetFids
	return &this
}

// NewFollowReqBodyWithDefaults instantiates a new FollowReqBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFollowReqBodyWithDefaults() *FollowReqBody {
	this := FollowReqBody{}
	return &this
}

// GetSignerUuid returns the SignerUuid field value
func (o *FollowReqBody) GetSignerUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SignerUuid
}

// GetSignerUuidOk returns a tuple with the SignerUuid field value
// and a boolean to check if the value has been set.
func (o *FollowReqBody) GetSignerUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SignerUuid, true
}

// SetSignerUuid sets field value
func (o *FollowReqBody) SetSignerUuid(v string) {
	o.SignerUuid = v
}

// GetTargetFids returns the TargetFids field value
func (o *FollowReqBody) GetTargetFids() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.TargetFids
}

// GetTargetFidsOk returns a tuple with the TargetFids field value
// and a boolean to check if the value has been set.
func (o *FollowReqBody) GetTargetFidsOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TargetFids, true
}

// SetTargetFids sets field value
func (o *FollowReqBody) SetTargetFids(v []int32) {
	o.TargetFids = v
}

func (o FollowReqBody) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FollowReqBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["signer_uuid"] = o.SignerUuid
	toSerialize["target_fids"] = o.TargetFids
	return toSerialize, nil
}

func (o *FollowReqBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"signer_uuid",
		"target_fids",
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
	varFollowReqBody := _FollowReqBody{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFollowReqBody)

	if err != nil {
		return err
	}

	*o = FollowReqBody(varFollowReqBody)

	return err
}

type NullableFollowReqBody struct {
	value *FollowReqBody
	isSet bool
}

func (v NullableFollowReqBody) Get() *FollowReqBody {
	return v.value
}

func (v *NullableFollowReqBody) Set(val *FollowReqBody) {
	v.value = val
	v.isSet = true
}

func (v NullableFollowReqBody) IsSet() bool {
	return v.isSet
}

func (v *NullableFollowReqBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFollowReqBody(val *FollowReqBody) *NullableFollowReqBody {
	return &NullableFollowReqBody{value: val, isSet: true}
}

func (v NullableFollowReqBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFollowReqBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
