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

// checks if the UpdateUserReqBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateUserReqBody{}

// UpdateUserReqBody struct for UpdateUserReqBody
type UpdateUserReqBody struct {
	// UUID of the signer. `signer_uuid` is paired with API key, can't use a `uuid` made with a different API key.
	SignerUuid       string                             `json:"signer_uuid"`
	Bio              *string                            `json:"bio,omitempty"`
	PfpUrl           *string                            `json:"pfp_url,omitempty"`
	Url              *string                            `json:"url,omitempty"`
	Username         *string                            `json:"username,omitempty"`
	DisplayName      *string                            `json:"display_name,omitempty"`
	Location         *UpdateUserReqBodyLocation         `json:"location,omitempty"`
	VerifiedAccounts *UpdateUserReqBodyVerifiedAccounts `json:"verified_accounts,omitempty"`
}

type _UpdateUserReqBody UpdateUserReqBody

// NewUpdateUserReqBody instantiates a new UpdateUserReqBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateUserReqBody(signerUuid string) *UpdateUserReqBody {
	this := UpdateUserReqBody{}
	this.SignerUuid = signerUuid
	return &this
}

// NewUpdateUserReqBodyWithDefaults instantiates a new UpdateUserReqBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateUserReqBodyWithDefaults() *UpdateUserReqBody {
	this := UpdateUserReqBody{}
	return &this
}

// GetSignerUuid returns the SignerUuid field value
func (o *UpdateUserReqBody) GetSignerUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SignerUuid
}

// GetSignerUuidOk returns a tuple with the SignerUuid field value
// and a boolean to check if the value has been set.
func (o *UpdateUserReqBody) GetSignerUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SignerUuid, true
}

// SetSignerUuid sets field value
func (o *UpdateUserReqBody) SetSignerUuid(v string) {
	o.SignerUuid = v
}

// GetBio returns the Bio field value if set, zero value otherwise.
func (o *UpdateUserReqBody) GetBio() string {
	if o == nil || IsNil(o.Bio) {
		var ret string
		return ret
	}
	return *o.Bio
}

// GetBioOk returns a tuple with the Bio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserReqBody) GetBioOk() (*string, bool) {
	if o == nil || IsNil(o.Bio) {
		return nil, false
	}
	return o.Bio, true
}

// HasBio returns a boolean if a field has been set.
func (o *UpdateUserReqBody) HasBio() bool {
	if o != nil && !IsNil(o.Bio) {
		return true
	}

	return false
}

// SetBio gets a reference to the given string and assigns it to the Bio field.
func (o *UpdateUserReqBody) SetBio(v string) {
	o.Bio = &v
}

// GetPfpUrl returns the PfpUrl field value if set, zero value otherwise.
func (o *UpdateUserReqBody) GetPfpUrl() string {
	if o == nil || IsNil(o.PfpUrl) {
		var ret string
		return ret
	}
	return *o.PfpUrl
}

// GetPfpUrlOk returns a tuple with the PfpUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserReqBody) GetPfpUrlOk() (*string, bool) {
	if o == nil || IsNil(o.PfpUrl) {
		return nil, false
	}
	return o.PfpUrl, true
}

// HasPfpUrl returns a boolean if a field has been set.
func (o *UpdateUserReqBody) HasPfpUrl() bool {
	if o != nil && !IsNil(o.PfpUrl) {
		return true
	}

	return false
}

// SetPfpUrl gets a reference to the given string and assigns it to the PfpUrl field.
func (o *UpdateUserReqBody) SetPfpUrl(v string) {
	o.PfpUrl = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *UpdateUserReqBody) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserReqBody) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *UpdateUserReqBody) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *UpdateUserReqBody) SetUrl(v string) {
	o.Url = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *UpdateUserReqBody) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserReqBody) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *UpdateUserReqBody) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *UpdateUserReqBody) SetUsername(v string) {
	o.Username = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *UpdateUserReqBody) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserReqBody) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *UpdateUserReqBody) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *UpdateUserReqBody) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *UpdateUserReqBody) GetLocation() UpdateUserReqBodyLocation {
	if o == nil || IsNil(o.Location) {
		var ret UpdateUserReqBodyLocation
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserReqBody) GetLocationOk() (*UpdateUserReqBodyLocation, bool) {
	if o == nil || IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *UpdateUserReqBody) HasLocation() bool {
	if o != nil && !IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given UpdateUserReqBodyLocation and assigns it to the Location field.
func (o *UpdateUserReqBody) SetLocation(v UpdateUserReqBodyLocation) {
	o.Location = &v
}

// GetVerifiedAccounts returns the VerifiedAccounts field value if set, zero value otherwise.
func (o *UpdateUserReqBody) GetVerifiedAccounts() UpdateUserReqBodyVerifiedAccounts {
	if o == nil || IsNil(o.VerifiedAccounts) {
		var ret UpdateUserReqBodyVerifiedAccounts
		return ret
	}
	return *o.VerifiedAccounts
}

// GetVerifiedAccountsOk returns a tuple with the VerifiedAccounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserReqBody) GetVerifiedAccountsOk() (*UpdateUserReqBodyVerifiedAccounts, bool) {
	if o == nil || IsNil(o.VerifiedAccounts) {
		return nil, false
	}
	return o.VerifiedAccounts, true
}

// HasVerifiedAccounts returns a boolean if a field has been set.
func (o *UpdateUserReqBody) HasVerifiedAccounts() bool {
	if o != nil && !IsNil(o.VerifiedAccounts) {
		return true
	}

	return false
}

// SetVerifiedAccounts gets a reference to the given UpdateUserReqBodyVerifiedAccounts and assigns it to the VerifiedAccounts field.
func (o *UpdateUserReqBody) SetVerifiedAccounts(v UpdateUserReqBodyVerifiedAccounts) {
	o.VerifiedAccounts = &v
}

func (o UpdateUserReqBody) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateUserReqBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["signer_uuid"] = o.SignerUuid
	if !IsNil(o.Bio) {
		toSerialize["bio"] = o.Bio
	}
	if !IsNil(o.PfpUrl) {
		toSerialize["pfp_url"] = o.PfpUrl
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	if !IsNil(o.DisplayName) {
		toSerialize["display_name"] = o.DisplayName
	}
	if !IsNil(o.Location) {
		toSerialize["location"] = o.Location
	}
	if !IsNil(o.VerifiedAccounts) {
		toSerialize["verified_accounts"] = o.VerifiedAccounts
	}
	return toSerialize, nil
}

func (o *UpdateUserReqBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"signer_uuid",
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
	varUpdateUserReqBody := _UpdateUserReqBody{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpdateUserReqBody)

	if err != nil {
		return err
	}

	*o = UpdateUserReqBody(varUpdateUserReqBody)

	return err
}

type NullableUpdateUserReqBody struct {
	value *UpdateUserReqBody
	isSet bool
}

func (v NullableUpdateUserReqBody) Get() *UpdateUserReqBody {
	return v.value
}

func (v *NullableUpdateUserReqBody) Set(val *UpdateUserReqBody) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateUserReqBody) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateUserReqBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateUserReqBody(val *UpdateUserReqBody) *NullableUpdateUserReqBody {
	return &NullableUpdateUserReqBody{value: val, isSet: true}
}

func (v NullableUpdateUserReqBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateUserReqBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
