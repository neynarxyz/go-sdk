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

// checks if the Signer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Signer{}

// Signer struct for Signer
type Signer struct {
	Object *string `json:"object,omitempty"`
	// UUID of the signer. `signer_uuid` is paired with API key, can't use a `uuid` made with a different API key.
	SignerUuid string `json:"signer_uuid"`
	// Ed25519 public key
	PublicKey         string  `json:"public_key" validate:"regexp=^0x[a-fA-F0-9]{64}$"`
	Status            string  `json:"status"`
	SignerApprovalUrl *string `json:"signer_approval_url,omitempty"`
	// The unique identifier of a farcaster user (unsigned integer)
	Fid         *int32                   `json:"fid,omitempty"`
	Permissions []SharedSignerPermission `json:"permissions,omitempty"`
}

type _Signer Signer

// NewSigner instantiates a new Signer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSigner(signerUuid string, publicKey string, status string) *Signer {
	this := Signer{}
	this.SignerUuid = signerUuid
	this.PublicKey = publicKey
	this.Status = status
	return &this
}

// NewSignerWithDefaults instantiates a new Signer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignerWithDefaults() *Signer {
	this := Signer{}
	return &this
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *Signer) GetObject() string {
	if o == nil || IsNil(o.Object) {
		var ret string
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Signer) GetObjectOk() (*string, bool) {
	if o == nil || IsNil(o.Object) {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *Signer) HasObject() bool {
	if o != nil && !IsNil(o.Object) {
		return true
	}

	return false
}

// SetObject gets a reference to the given string and assigns it to the Object field.
func (o *Signer) SetObject(v string) {
	o.Object = &v
}

// GetSignerUuid returns the SignerUuid field value
func (o *Signer) GetSignerUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SignerUuid
}

// GetSignerUuidOk returns a tuple with the SignerUuid field value
// and a boolean to check if the value has been set.
func (o *Signer) GetSignerUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SignerUuid, true
}

// SetSignerUuid sets field value
func (o *Signer) SetSignerUuid(v string) {
	o.SignerUuid = v
}

// GetPublicKey returns the PublicKey field value
func (o *Signer) GetPublicKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PublicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value
// and a boolean to check if the value has been set.
func (o *Signer) GetPublicKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublicKey, true
}

// SetPublicKey sets field value
func (o *Signer) SetPublicKey(v string) {
	o.PublicKey = v
}

// GetStatus returns the Status field value
func (o *Signer) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Signer) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Signer) SetStatus(v string) {
	o.Status = v
}

// GetSignerApprovalUrl returns the SignerApprovalUrl field value if set, zero value otherwise.
func (o *Signer) GetSignerApprovalUrl() string {
	if o == nil || IsNil(o.SignerApprovalUrl) {
		var ret string
		return ret
	}
	return *o.SignerApprovalUrl
}

// GetSignerApprovalUrlOk returns a tuple with the SignerApprovalUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Signer) GetSignerApprovalUrlOk() (*string, bool) {
	if o == nil || IsNil(o.SignerApprovalUrl) {
		return nil, false
	}
	return o.SignerApprovalUrl, true
}

// HasSignerApprovalUrl returns a boolean if a field has been set.
func (o *Signer) HasSignerApprovalUrl() bool {
	if o != nil && !IsNil(o.SignerApprovalUrl) {
		return true
	}

	return false
}

// SetSignerApprovalUrl gets a reference to the given string and assigns it to the SignerApprovalUrl field.
func (o *Signer) SetSignerApprovalUrl(v string) {
	o.SignerApprovalUrl = &v
}

// GetFid returns the Fid field value if set, zero value otherwise.
func (o *Signer) GetFid() int32 {
	if o == nil || IsNil(o.Fid) {
		var ret int32
		return ret
	}
	return *o.Fid
}

// GetFidOk returns a tuple with the Fid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Signer) GetFidOk() (*int32, bool) {
	if o == nil || IsNil(o.Fid) {
		return nil, false
	}
	return o.Fid, true
}

// HasFid returns a boolean if a field has been set.
func (o *Signer) HasFid() bool {
	if o != nil && !IsNil(o.Fid) {
		return true
	}

	return false
}

// SetFid gets a reference to the given int32 and assigns it to the Fid field.
func (o *Signer) SetFid(v int32) {
	o.Fid = &v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *Signer) GetPermissions() []SharedSignerPermission {
	if o == nil || IsNil(o.Permissions) {
		var ret []SharedSignerPermission
		return ret
	}
	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Signer) GetPermissionsOk() ([]SharedSignerPermission, bool) {
	if o == nil || IsNil(o.Permissions) {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *Signer) HasPermissions() bool {
	if o != nil && !IsNil(o.Permissions) {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given []SharedSignerPermission and assigns it to the Permissions field.
func (o *Signer) SetPermissions(v []SharedSignerPermission) {
	o.Permissions = v
}

func (o Signer) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Signer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Object) {
		toSerialize["object"] = o.Object
	}
	toSerialize["signer_uuid"] = o.SignerUuid
	toSerialize["public_key"] = o.PublicKey
	toSerialize["status"] = o.Status
	if !IsNil(o.SignerApprovalUrl) {
		toSerialize["signer_approval_url"] = o.SignerApprovalUrl
	}
	if !IsNil(o.Fid) {
		toSerialize["fid"] = o.Fid
	}
	if !IsNil(o.Permissions) {
		toSerialize["permissions"] = o.Permissions
	}
	return toSerialize, nil
}

func (o *Signer) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"signer_uuid",
		"public_key",
		"status",
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
	varSigner := _Signer{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSigner)

	if err != nil {
		return err
	}

	*o = Signer(varSigner)

	return err
}

type NullableSigner struct {
	value *Signer
	isSet bool
}

func (v NullableSigner) Get() *Signer {
	return v.value
}

func (v *NullableSigner) Set(val *Signer) {
	v.value = val
	v.isSet = true
}

func (v NullableSigner) IsSet() bool {
	return v.isSet
}

func (v *NullableSigner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSigner(val *Signer) *NullableSigner {
	return &NullableSigner{value: val, isSet: true}
}

func (v NullableSigner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSigner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
