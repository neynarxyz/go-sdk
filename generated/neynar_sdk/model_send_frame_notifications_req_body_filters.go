/*
Farcaster API V2

The Farcaster API allows you to interact with the Farcaster protocol. See the [Neynar docs](https://docs.neynar.com/reference) for more details.

API version: 2.43.0
Contact: team@neynar.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package neynar_sdk

import (
	"encoding/json"
)

// checks if the SendFrameNotificationsReqBodyFilters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SendFrameNotificationsReqBodyFilters{}

// SendFrameNotificationsReqBodyFilters Filters to apply to the target_fids set. All filters are additive, so only users matching all filters will be notified.
type SendFrameNotificationsReqBodyFilters struct {
	// Only send notifications to users who are not in the given FIDs.
	ExcludeFids []int32 `json:"exclude_fids,omitempty"`
	// Only send notifications to users who follow the given FID.
	FollowingFid *int32 `json:"following_fid,omitempty"`
	// Only send notifications to users with a score greater than or equal to this value.
	MinimumUserScore *float32  `json:"minimum_user_score,omitempty"`
	NearLocation     *Location `json:"near_location,omitempty"`
}

// NewSendFrameNotificationsReqBodyFilters instantiates a new SendFrameNotificationsReqBodyFilters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSendFrameNotificationsReqBodyFilters() *SendFrameNotificationsReqBodyFilters {
	this := SendFrameNotificationsReqBodyFilters{}
	return &this
}

// NewSendFrameNotificationsReqBodyFiltersWithDefaults instantiates a new SendFrameNotificationsReqBodyFilters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSendFrameNotificationsReqBodyFiltersWithDefaults() *SendFrameNotificationsReqBodyFilters {
	this := SendFrameNotificationsReqBodyFilters{}
	return &this
}

// GetExcludeFids returns the ExcludeFids field value if set, zero value otherwise.
func (o *SendFrameNotificationsReqBodyFilters) GetExcludeFids() []int32 {
	if o == nil || IsNil(o.ExcludeFids) {
		var ret []int32
		return ret
	}
	return o.ExcludeFids
}

// GetExcludeFidsOk returns a tuple with the ExcludeFids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendFrameNotificationsReqBodyFilters) GetExcludeFidsOk() ([]int32, bool) {
	if o == nil || IsNil(o.ExcludeFids) {
		return nil, false
	}
	return o.ExcludeFids, true
}

// HasExcludeFids returns a boolean if a field has been set.
func (o *SendFrameNotificationsReqBodyFilters) HasExcludeFids() bool {
	if o != nil && !IsNil(o.ExcludeFids) {
		return true
	}

	return false
}

// SetExcludeFids gets a reference to the given []int32 and assigns it to the ExcludeFids field.
func (o *SendFrameNotificationsReqBodyFilters) SetExcludeFids(v []int32) {
	o.ExcludeFids = v
}

// GetFollowingFid returns the FollowingFid field value if set, zero value otherwise.
func (o *SendFrameNotificationsReqBodyFilters) GetFollowingFid() int32 {
	if o == nil || IsNil(o.FollowingFid) {
		var ret int32
		return ret
	}
	return *o.FollowingFid
}

// GetFollowingFidOk returns a tuple with the FollowingFid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendFrameNotificationsReqBodyFilters) GetFollowingFidOk() (*int32, bool) {
	if o == nil || IsNil(o.FollowingFid) {
		return nil, false
	}
	return o.FollowingFid, true
}

// HasFollowingFid returns a boolean if a field has been set.
func (o *SendFrameNotificationsReqBodyFilters) HasFollowingFid() bool {
	if o != nil && !IsNil(o.FollowingFid) {
		return true
	}

	return false
}

// SetFollowingFid gets a reference to the given int32 and assigns it to the FollowingFid field.
func (o *SendFrameNotificationsReqBodyFilters) SetFollowingFid(v int32) {
	o.FollowingFid = &v
}

// GetMinimumUserScore returns the MinimumUserScore field value if set, zero value otherwise.
func (o *SendFrameNotificationsReqBodyFilters) GetMinimumUserScore() float32 {
	if o == nil || IsNil(o.MinimumUserScore) {
		var ret float32
		return ret
	}
	return *o.MinimumUserScore
}

// GetMinimumUserScoreOk returns a tuple with the MinimumUserScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendFrameNotificationsReqBodyFilters) GetMinimumUserScoreOk() (*float32, bool) {
	if o == nil || IsNil(o.MinimumUserScore) {
		return nil, false
	}
	return o.MinimumUserScore, true
}

// HasMinimumUserScore returns a boolean if a field has been set.
func (o *SendFrameNotificationsReqBodyFilters) HasMinimumUserScore() bool {
	if o != nil && !IsNil(o.MinimumUserScore) {
		return true
	}

	return false
}

// SetMinimumUserScore gets a reference to the given float32 and assigns it to the MinimumUserScore field.
func (o *SendFrameNotificationsReqBodyFilters) SetMinimumUserScore(v float32) {
	o.MinimumUserScore = &v
}

// GetNearLocation returns the NearLocation field value if set, zero value otherwise.
func (o *SendFrameNotificationsReqBodyFilters) GetNearLocation() Location {
	if o == nil || IsNil(o.NearLocation) {
		var ret Location
		return ret
	}
	return *o.NearLocation
}

// GetNearLocationOk returns a tuple with the NearLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendFrameNotificationsReqBodyFilters) GetNearLocationOk() (*Location, bool) {
	if o == nil || IsNil(o.NearLocation) {
		return nil, false
	}
	return o.NearLocation, true
}

// HasNearLocation returns a boolean if a field has been set.
func (o *SendFrameNotificationsReqBodyFilters) HasNearLocation() bool {
	if o != nil && !IsNil(o.NearLocation) {
		return true
	}

	return false
}

// SetNearLocation gets a reference to the given Location and assigns it to the NearLocation field.
func (o *SendFrameNotificationsReqBodyFilters) SetNearLocation(v Location) {
	o.NearLocation = &v
}

func (o SendFrameNotificationsReqBodyFilters) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SendFrameNotificationsReqBodyFilters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExcludeFids) {
		toSerialize["exclude_fids"] = o.ExcludeFids
	}
	if !IsNil(o.FollowingFid) {
		toSerialize["following_fid"] = o.FollowingFid
	}
	if !IsNil(o.MinimumUserScore) {
		toSerialize["minimum_user_score"] = o.MinimumUserScore
	}
	if !IsNil(o.NearLocation) {
		toSerialize["near_location"] = o.NearLocation
	}
	return toSerialize, nil
}

type NullableSendFrameNotificationsReqBodyFilters struct {
	value *SendFrameNotificationsReqBodyFilters
	isSet bool
}

func (v NullableSendFrameNotificationsReqBodyFilters) Get() *SendFrameNotificationsReqBodyFilters {
	return v.value
}

func (v *NullableSendFrameNotificationsReqBodyFilters) Set(val *SendFrameNotificationsReqBodyFilters) {
	v.value = val
	v.isSet = true
}

func (v NullableSendFrameNotificationsReqBodyFilters) IsSet() bool {
	return v.isSet
}

func (v *NullableSendFrameNotificationsReqBodyFilters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSendFrameNotificationsReqBodyFilters(val *SendFrameNotificationsReqBodyFilters) *NullableSendFrameNotificationsReqBodyFilters {
	return &NullableSendFrameNotificationsReqBodyFilters{value: val, isSet: true}
}

func (v NullableSendFrameNotificationsReqBodyFilters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSendFrameNotificationsReqBodyFilters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
