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

// checks if the CastAddAllOfData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CastAddAllOfData{}

// CastAddAllOfData struct for CastAddAllOfData
type CastAddAllOfData struct {
	Type MessageType `json:"type"`
	// The unique identifier (FID) of the user who created this message. FIDs are assigned sequentially when users register on the network and cannot be changed.
	Fid int32 `json:"fid"`
	// Seconds since Farcaster Epoch (2021-01-01T00:00:00Z). Used to order messages chronologically and determine the most recent state. Must be within 10 minutes of the current time when the message is created.
	Timestamp int64            `json:"timestamp"`
	Network   FarcasterNetwork `json:"network"`
	// The content and metadata of the new cast, including the text, mentions, embeds, and any parent references for replies.
	CastAddBody CastAddBody `json:"castAddBody"`
}

type _CastAddAllOfData CastAddAllOfData

// NewCastAddAllOfData instantiates a new CastAddAllOfData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCastAddAllOfData(type_ MessageType, fid int32, timestamp int64, network FarcasterNetwork, castAddBody CastAddBody) *CastAddAllOfData {
	this := CastAddAllOfData{}
	this.Type = type_
	this.Fid = fid
	this.Timestamp = timestamp
	this.Network = network
	this.CastAddBody = castAddBody
	return &this
}

// NewCastAddAllOfDataWithDefaults instantiates a new CastAddAllOfData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCastAddAllOfDataWithDefaults() *CastAddAllOfData {
	this := CastAddAllOfData{}
	var type_ MessageType = MESSAGETYPE_MESSAGE_TYPE_CAST_ADD
	this.Type = type_
	var network FarcasterNetwork = FARCASTERNETWORK_FARCASTER_NETWORK_MAINNET
	this.Network = network
	return &this
}

// GetType returns the Type field value
func (o *CastAddAllOfData) GetType() MessageType {
	if o == nil {
		var ret MessageType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CastAddAllOfData) GetTypeOk() (*MessageType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CastAddAllOfData) SetType(v MessageType) {
	o.Type = v
}

// GetDefaultType returns the default value MESSAGETYPE_MESSAGE_TYPE_CAST_ADD of the Type field.
func (o *CastAddAllOfData) GetDefaultType() interface{} {
	return MESSAGETYPE_MESSAGE_TYPE_CAST_ADD
}

// GetFid returns the Fid field value
func (o *CastAddAllOfData) GetFid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Fid
}

// GetFidOk returns a tuple with the Fid field value
// and a boolean to check if the value has been set.
func (o *CastAddAllOfData) GetFidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fid, true
}

// SetFid sets field value
func (o *CastAddAllOfData) SetFid(v int32) {
	o.Fid = v
}

// GetTimestamp returns the Timestamp field value
func (o *CastAddAllOfData) GetTimestamp() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *CastAddAllOfData) GetTimestampOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *CastAddAllOfData) SetTimestamp(v int64) {
	o.Timestamp = v
}

// GetNetwork returns the Network field value
func (o *CastAddAllOfData) GetNetwork() FarcasterNetwork {
	if o == nil {
		var ret FarcasterNetwork
		return ret
	}

	return o.Network
}

// GetNetworkOk returns a tuple with the Network field value
// and a boolean to check if the value has been set.
func (o *CastAddAllOfData) GetNetworkOk() (*FarcasterNetwork, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Network, true
}

// SetNetwork sets field value
func (o *CastAddAllOfData) SetNetwork(v FarcasterNetwork) {
	o.Network = v
}

// GetDefaultNetwork returns the default value FARCASTERNETWORK_FARCASTER_NETWORK_MAINNET of the Network field.
func (o *CastAddAllOfData) GetDefaultNetwork() interface{} {
	return FARCASTERNETWORK_FARCASTER_NETWORK_MAINNET
}

// GetCastAddBody returns the CastAddBody field value
func (o *CastAddAllOfData) GetCastAddBody() CastAddBody {
	if o == nil {
		var ret CastAddBody
		return ret
	}

	return o.CastAddBody
}

// GetCastAddBodyOk returns a tuple with the CastAddBody field value
// and a boolean to check if the value has been set.
func (o *CastAddAllOfData) GetCastAddBodyOk() (*CastAddBody, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CastAddBody, true
}

// SetCastAddBody sets field value
func (o *CastAddAllOfData) SetCastAddBody(v CastAddBody) {
	o.CastAddBody = v
}

func (o CastAddAllOfData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CastAddAllOfData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if _, exists := toSerialize["type"]; !exists {
		toSerialize["type"] = o.GetDefaultType()
	}
	toSerialize["type"] = o.Type
	toSerialize["fid"] = o.Fid
	toSerialize["timestamp"] = o.Timestamp
	if _, exists := toSerialize["network"]; !exists {
		toSerialize["network"] = o.GetDefaultNetwork()
	}
	toSerialize["network"] = o.Network
	toSerialize["castAddBody"] = o.CastAddBody
	return toSerialize, nil
}

func (o *CastAddAllOfData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"fid",
		"timestamp",
		"network",
		"castAddBody",
	}

	// defaultValueFuncMap captures the default values for required properties.
	// These values are used when required properties are missing from the payload.
	defaultValueFuncMap := map[string]func() interface{}{
		"type":    o.GetDefaultType,
		"network": o.GetDefaultNetwork,
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
	varCastAddAllOfData := _CastAddAllOfData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCastAddAllOfData)

	if err != nil {
		return err
	}

	*o = CastAddAllOfData(varCastAddAllOfData)

	return err
}

type NullableCastAddAllOfData struct {
	value *CastAddAllOfData
	isSet bool
}

func (v NullableCastAddAllOfData) Get() *CastAddAllOfData {
	return v.value
}

func (v *NullableCastAddAllOfData) Set(val *CastAddAllOfData) {
	v.value = val
	v.isSet = true
}

func (v NullableCastAddAllOfData) IsSet() bool {
	return v.isSet
}

func (v *NullableCastAddAllOfData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCastAddAllOfData(val *CastAddAllOfData) *NullableCastAddAllOfData {
	return &NullableCastAddAllOfData{value: val, isSet: true}
}

func (v NullableCastAddAllOfData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCastAddAllOfData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
