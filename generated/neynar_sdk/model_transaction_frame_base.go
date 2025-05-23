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

// checks if the TransactionFrameBase type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransactionFrameBase{}

// TransactionFrameBase struct for TransactionFrameBase
type TransactionFrameBase struct {
	// Unique identifier for the transaction mini app
	Id string `json:"id"`
	// URL that can be used to access the transaction mini app
	Url    string                 `json:"url"`
	Type   TransactionFrameType   `json:"type"`
	Config TransactionFrameConfig `json:"config"`
	Status TransactionFrameStatus `json:"status"`
}

type _TransactionFrameBase TransactionFrameBase

// NewTransactionFrameBase instantiates a new TransactionFrameBase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionFrameBase(id string, url string, type_ TransactionFrameType, config TransactionFrameConfig, status TransactionFrameStatus) *TransactionFrameBase {
	this := TransactionFrameBase{}
	this.Id = id
	this.Url = url
	this.Type = type_
	this.Config = config
	this.Status = status
	return &this
}

// NewTransactionFrameBaseWithDefaults instantiates a new TransactionFrameBase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionFrameBaseWithDefaults() *TransactionFrameBase {
	this := TransactionFrameBase{}
	return &this
}

// GetId returns the Id field value
func (o *TransactionFrameBase) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TransactionFrameBase) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TransactionFrameBase) SetId(v string) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *TransactionFrameBase) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *TransactionFrameBase) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *TransactionFrameBase) SetUrl(v string) {
	o.Url = v
}

// GetType returns the Type field value
func (o *TransactionFrameBase) GetType() TransactionFrameType {
	if o == nil {
		var ret TransactionFrameType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TransactionFrameBase) GetTypeOk() (*TransactionFrameType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TransactionFrameBase) SetType(v TransactionFrameType) {
	o.Type = v
}

// GetConfig returns the Config field value
func (o *TransactionFrameBase) GetConfig() TransactionFrameConfig {
	if o == nil {
		var ret TransactionFrameConfig
		return ret
	}

	return o.Config
}

// GetConfigOk returns a tuple with the Config field value
// and a boolean to check if the value has been set.
func (o *TransactionFrameBase) GetConfigOk() (*TransactionFrameConfig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Config, true
}

// SetConfig sets field value
func (o *TransactionFrameBase) SetConfig(v TransactionFrameConfig) {
	o.Config = v
}

// GetStatus returns the Status field value
func (o *TransactionFrameBase) GetStatus() TransactionFrameStatus {
	if o == nil {
		var ret TransactionFrameStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *TransactionFrameBase) GetStatusOk() (*TransactionFrameStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *TransactionFrameBase) SetStatus(v TransactionFrameStatus) {
	o.Status = v
}

func (o TransactionFrameBase) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionFrameBase) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	toSerialize["type"] = o.Type
	toSerialize["config"] = o.Config
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *TransactionFrameBase) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"url",
		"type",
		"config",
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
	varTransactionFrameBase := _TransactionFrameBase{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTransactionFrameBase)

	if err != nil {
		return err
	}

	*o = TransactionFrameBase(varTransactionFrameBase)

	return err
}

type NullableTransactionFrameBase struct {
	value *TransactionFrameBase
	isSet bool
}

func (v NullableTransactionFrameBase) Get() *TransactionFrameBase {
	return v.value
}

func (v *NullableTransactionFrameBase) Set(val *TransactionFrameBase) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionFrameBase) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionFrameBase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionFrameBase(val *TransactionFrameBase) *NullableTransactionFrameBase {
	return &NullableTransactionFrameBase{value: val, isSet: true}
}

func (v NullableTransactionFrameBase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionFrameBase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
