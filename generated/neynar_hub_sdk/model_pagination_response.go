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

// checks if the PaginationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaginationResponse{}

// PaginationResponse struct for PaginationResponse
type PaginationResponse struct {
	// Base64-encoded pagination token for fetching the next page of results. An empty value indicates there are no more pages to return. Used in conjunction with the pageSize parameter to implement pagination across large result sets.
	NextPageToken string `json:"nextPageToken" validate:"regexp=^(?:[A-Za-z0-9+\\/]{4})*(?:[A-Za-z0-9+\\/]{2}==|[A-Za-z0-9+\\/]{3}=)?$"`
}

type _PaginationResponse PaginationResponse

// NewPaginationResponse instantiates a new PaginationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginationResponse(nextPageToken string) *PaginationResponse {
	this := PaginationResponse{}
	this.NextPageToken = nextPageToken
	return &this
}

// NewPaginationResponseWithDefaults instantiates a new PaginationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginationResponseWithDefaults() *PaginationResponse {
	this := PaginationResponse{}
	return &this
}

// GetNextPageToken returns the NextPageToken field value
func (o *PaginationResponse) GetNextPageToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value
// and a boolean to check if the value has been set.
func (o *PaginationResponse) GetNextPageTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NextPageToken, true
}

// SetNextPageToken sets field value
func (o *PaginationResponse) SetNextPageToken(v string) {
	o.NextPageToken = v
}

func (o PaginationResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaginationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["nextPageToken"] = o.NextPageToken
	return toSerialize, nil
}

func (o *PaginationResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"nextPageToken",
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
	varPaginationResponse := _PaginationResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPaginationResponse)

	if err != nil {
		return err
	}

	*o = PaginationResponse(varPaginationResponse)

	return err
}

type NullablePaginationResponse struct {
	value *PaginationResponse
	isSet bool
}

func (v NullablePaginationResponse) Get() *PaginationResponse {
	return v.value
}

func (v *NullablePaginationResponse) Set(val *PaginationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginationResponse(val *PaginationResponse) *NullablePaginationResponse {
	return &NullablePaginationResponse{value: val, isSet: true}
}

func (v NullablePaginationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
