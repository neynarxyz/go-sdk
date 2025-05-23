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

// checks if the IdRegisterEventBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdRegisterEventBody{}

// IdRegisterEventBody struct for IdRegisterEventBody
type IdRegisterEventBody struct {
	To              string              `json:"to" validate:"regexp=^0x[a-fA-F0-9]*$"`
	EventType       IdRegisterEventType `json:"eventType"`
	From            string              `json:"from" validate:"regexp=^0x[a-fA-F0-9]*$|^$"`
	RecoveryAddress string              `json:"recoveryAddress" validate:"regexp=^0x[a-fA-F0-9]*$"`
}

type _IdRegisterEventBody IdRegisterEventBody

// NewIdRegisterEventBody instantiates a new IdRegisterEventBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdRegisterEventBody(to string, eventType IdRegisterEventType, from string, recoveryAddress string) *IdRegisterEventBody {
	this := IdRegisterEventBody{}
	this.To = to
	this.EventType = eventType
	this.From = from
	this.RecoveryAddress = recoveryAddress
	return &this
}

// NewIdRegisterEventBodyWithDefaults instantiates a new IdRegisterEventBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdRegisterEventBodyWithDefaults() *IdRegisterEventBody {
	this := IdRegisterEventBody{}
	var eventType IdRegisterEventType = IDREGISTEREVENTTYPE_ID_REGISTER_EVENT_TYPE_REGISTER
	this.EventType = eventType
	return &this
}

// GetTo returns the To field value
func (o *IdRegisterEventBody) GetTo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.To
}

// GetToOk returns a tuple with the To field value
// and a boolean to check if the value has been set.
func (o *IdRegisterEventBody) GetToOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.To, true
}

// SetTo sets field value
func (o *IdRegisterEventBody) SetTo(v string) {
	o.To = v
}

// GetEventType returns the EventType field value
func (o *IdRegisterEventBody) GetEventType() IdRegisterEventType {
	if o == nil {
		var ret IdRegisterEventType
		return ret
	}

	return o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value
// and a boolean to check if the value has been set.
func (o *IdRegisterEventBody) GetEventTypeOk() (*IdRegisterEventType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventType, true
}

// SetEventType sets field value
func (o *IdRegisterEventBody) SetEventType(v IdRegisterEventType) {
	o.EventType = v
}

// GetDefaultEventType returns the default value IDREGISTEREVENTTYPE_ID_REGISTER_EVENT_TYPE_REGISTER of the EventType field.
func (o *IdRegisterEventBody) GetDefaultEventType() interface{} {
	return IDREGISTEREVENTTYPE_ID_REGISTER_EVENT_TYPE_REGISTER
}

// GetFrom returns the From field value
func (o *IdRegisterEventBody) GetFrom() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.From
}

// GetFromOk returns a tuple with the From field value
// and a boolean to check if the value has been set.
func (o *IdRegisterEventBody) GetFromOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.From, true
}

// SetFrom sets field value
func (o *IdRegisterEventBody) SetFrom(v string) {
	o.From = v
}

// GetRecoveryAddress returns the RecoveryAddress field value
func (o *IdRegisterEventBody) GetRecoveryAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RecoveryAddress
}

// GetRecoveryAddressOk returns a tuple with the RecoveryAddress field value
// and a boolean to check if the value has been set.
func (o *IdRegisterEventBody) GetRecoveryAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RecoveryAddress, true
}

// SetRecoveryAddress sets field value
func (o *IdRegisterEventBody) SetRecoveryAddress(v string) {
	o.RecoveryAddress = v
}

func (o IdRegisterEventBody) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdRegisterEventBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["to"] = o.To
	if _, exists := toSerialize["eventType"]; !exists {
		toSerialize["eventType"] = o.GetDefaultEventType()
	}
	toSerialize["eventType"] = o.EventType
	toSerialize["from"] = o.From
	toSerialize["recoveryAddress"] = o.RecoveryAddress
	return toSerialize, nil
}

func (o *IdRegisterEventBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"to",
		"eventType",
		"from",
		"recoveryAddress",
	}

	// defaultValueFuncMap captures the default values for required properties.
	// These values are used when required properties are missing from the payload.
	defaultValueFuncMap := map[string]func() interface{}{
		"eventType": o.GetDefaultEventType,
	}
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
	varIdRegisterEventBody := _IdRegisterEventBody{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varIdRegisterEventBody)

	if err != nil {
		return err
	}

	*o = IdRegisterEventBody(varIdRegisterEventBody)

	return err
}

type NullableIdRegisterEventBody struct {
	value *IdRegisterEventBody
	isSet bool
}

func (v NullableIdRegisterEventBody) Get() *IdRegisterEventBody {
	return v.value
}

func (v *NullableIdRegisterEventBody) Set(val *IdRegisterEventBody) {
	v.value = val
	v.isSet = true
}

func (v NullableIdRegisterEventBody) IsSet() bool {
	return v.isSet
}

func (v *NullableIdRegisterEventBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdRegisterEventBody(val *IdRegisterEventBody) *NullableIdRegisterEventBody {
	return &NullableIdRegisterEventBody{value: val, isSet: true}
}

func (v NullableIdRegisterEventBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdRegisterEventBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
