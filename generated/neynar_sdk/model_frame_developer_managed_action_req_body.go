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

// checks if the FrameDeveloperManagedActionReqBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FrameDeveloperManagedActionReqBody{}

// FrameDeveloperManagedActionReqBody struct for FrameDeveloperManagedActionReqBody
type FrameDeveloperManagedActionReqBody struct {
	// Cast Hash
	CastHash        *string              `json:"cast_hash,omitempty"`
	Action          FrameAction          `json:"action"`
	SignaturePacket FrameSignaturePacket `json:"signature_packet"`
}

type _FrameDeveloperManagedActionReqBody FrameDeveloperManagedActionReqBody

// NewFrameDeveloperManagedActionReqBody instantiates a new FrameDeveloperManagedActionReqBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFrameDeveloperManagedActionReqBody(action FrameAction, signaturePacket FrameSignaturePacket) *FrameDeveloperManagedActionReqBody {
	this := FrameDeveloperManagedActionReqBody{}
	var castHash string = "0xfe90f9de682273e05b201629ad2338bdcd89b6be"
	this.CastHash = &castHash
	this.Action = action
	this.SignaturePacket = signaturePacket
	return &this
}

// NewFrameDeveloperManagedActionReqBodyWithDefaults instantiates a new FrameDeveloperManagedActionReqBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFrameDeveloperManagedActionReqBodyWithDefaults() *FrameDeveloperManagedActionReqBody {
	this := FrameDeveloperManagedActionReqBody{}
	var castHash string = "0xfe90f9de682273e05b201629ad2338bdcd89b6be"
	this.CastHash = &castHash
	return &this
}

// GetCastHash returns the CastHash field value if set, zero value otherwise.
func (o *FrameDeveloperManagedActionReqBody) GetCastHash() string {
	if o == nil || IsNil(o.CastHash) {
		var ret string
		return ret
	}
	return *o.CastHash
}

// GetCastHashOk returns a tuple with the CastHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FrameDeveloperManagedActionReqBody) GetCastHashOk() (*string, bool) {
	if o == nil || IsNil(o.CastHash) {
		return nil, false
	}
	return o.CastHash, true
}

// HasCastHash returns a boolean if a field has been set.
func (o *FrameDeveloperManagedActionReqBody) HasCastHash() bool {
	if o != nil && !IsNil(o.CastHash) {
		return true
	}

	return false
}

// SetCastHash gets a reference to the given string and assigns it to the CastHash field.
func (o *FrameDeveloperManagedActionReqBody) SetCastHash(v string) {
	o.CastHash = &v
}

// GetAction returns the Action field value
func (o *FrameDeveloperManagedActionReqBody) GetAction() FrameAction {
	if o == nil {
		var ret FrameAction
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *FrameDeveloperManagedActionReqBody) GetActionOk() (*FrameAction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *FrameDeveloperManagedActionReqBody) SetAction(v FrameAction) {
	o.Action = v
}

// GetSignaturePacket returns the SignaturePacket field value
func (o *FrameDeveloperManagedActionReqBody) GetSignaturePacket() FrameSignaturePacket {
	if o == nil {
		var ret FrameSignaturePacket
		return ret
	}

	return o.SignaturePacket
}

// GetSignaturePacketOk returns a tuple with the SignaturePacket field value
// and a boolean to check if the value has been set.
func (o *FrameDeveloperManagedActionReqBody) GetSignaturePacketOk() (*FrameSignaturePacket, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SignaturePacket, true
}

// SetSignaturePacket sets field value
func (o *FrameDeveloperManagedActionReqBody) SetSignaturePacket(v FrameSignaturePacket) {
	o.SignaturePacket = v
}

func (o FrameDeveloperManagedActionReqBody) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FrameDeveloperManagedActionReqBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CastHash) {
		toSerialize["cast_hash"] = o.CastHash
	}
	toSerialize["action"] = o.Action
	toSerialize["signature_packet"] = o.SignaturePacket
	return toSerialize, nil
}

func (o *FrameDeveloperManagedActionReqBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"action",
		"signature_packet",
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
	varFrameDeveloperManagedActionReqBody := _FrameDeveloperManagedActionReqBody{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFrameDeveloperManagedActionReqBody)

	if err != nil {
		return err
	}

	*o = FrameDeveloperManagedActionReqBody(varFrameDeveloperManagedActionReqBody)

	return err
}

type NullableFrameDeveloperManagedActionReqBody struct {
	value *FrameDeveloperManagedActionReqBody
	isSet bool
}

func (v NullableFrameDeveloperManagedActionReqBody) Get() *FrameDeveloperManagedActionReqBody {
	return v.value
}

func (v *NullableFrameDeveloperManagedActionReqBody) Set(val *FrameDeveloperManagedActionReqBody) {
	v.value = val
	v.isSet = true
}

func (v NullableFrameDeveloperManagedActionReqBody) IsSet() bool {
	return v.isSet
}

func (v *NullableFrameDeveloperManagedActionReqBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFrameDeveloperManagedActionReqBody(val *FrameDeveloperManagedActionReqBody) *NullableFrameDeveloperManagedActionReqBody {
	return &NullableFrameDeveloperManagedActionReqBody{value: val, isSet: true}
}

func (v NullableFrameDeveloperManagedActionReqBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFrameDeveloperManagedActionReqBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
