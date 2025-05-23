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

// checks if the BulkFollowResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BulkFollowResponse{}

// BulkFollowResponse struct for BulkFollowResponse
type BulkFollowResponse struct {
	Success bool             `json:"success"`
	Details []FollowResponse `json:"details"`
}

type _BulkFollowResponse BulkFollowResponse

// NewBulkFollowResponse instantiates a new BulkFollowResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkFollowResponse(success bool, details []FollowResponse) *BulkFollowResponse {
	this := BulkFollowResponse{}
	this.Success = success
	this.Details = details
	return &this
}

// NewBulkFollowResponseWithDefaults instantiates a new BulkFollowResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkFollowResponseWithDefaults() *BulkFollowResponse {
	this := BulkFollowResponse{}
	return &this
}

// GetSuccess returns the Success field value
func (o *BulkFollowResponse) GetSuccess() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value
// and a boolean to check if the value has been set.
func (o *BulkFollowResponse) GetSuccessOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Success, true
}

// SetSuccess sets field value
func (o *BulkFollowResponse) SetSuccess(v bool) {
	o.Success = v
}

// GetDetails returns the Details field value
func (o *BulkFollowResponse) GetDetails() []FollowResponse {
	if o == nil {
		var ret []FollowResponse
		return ret
	}

	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value
// and a boolean to check if the value has been set.
func (o *BulkFollowResponse) GetDetailsOk() ([]FollowResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Details, true
}

// SetDetails sets field value
func (o *BulkFollowResponse) SetDetails(v []FollowResponse) {
	o.Details = v
}

func (o BulkFollowResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BulkFollowResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["success"] = o.Success
	toSerialize["details"] = o.Details
	return toSerialize, nil
}

func (o *BulkFollowResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"success",
		"details",
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
	varBulkFollowResponse := _BulkFollowResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBulkFollowResponse)

	if err != nil {
		return err
	}

	*o = BulkFollowResponse(varBulkFollowResponse)

	return err
}

type NullableBulkFollowResponse struct {
	value *BulkFollowResponse
	isSet bool
}

func (v NullableBulkFollowResponse) Get() *BulkFollowResponse {
	return v.value
}

func (v *NullableBulkFollowResponse) Set(val *BulkFollowResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkFollowResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkFollowResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkFollowResponse(val *BulkFollowResponse) *NullableBulkFollowResponse {
	return &NullableBulkFollowResponse{value: val, isSet: true}
}

func (v NullableBulkFollowResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkFollowResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
