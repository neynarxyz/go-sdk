# SendFrameNotificationsReqBodyFiltersNearLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Latitude** | **float64** |  | 
**Longitude** | **float64** |  | 
**Address** | Pointer to [**LocationAddress**](LocationAddress.md) |  | [optional] 
**Radius** | Pointer to **float32** | The radius in meters for the location search. Any location within this radius will be returned. | [optional] 

## Methods

### NewSendFrameNotificationsReqBodyFiltersNearLocation

`func NewSendFrameNotificationsReqBodyFiltersNearLocation(latitude float64, longitude float64, ) *SendFrameNotificationsReqBodyFiltersNearLocation`

NewSendFrameNotificationsReqBodyFiltersNearLocation instantiates a new SendFrameNotificationsReqBodyFiltersNearLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendFrameNotificationsReqBodyFiltersNearLocationWithDefaults

`func NewSendFrameNotificationsReqBodyFiltersNearLocationWithDefaults() *SendFrameNotificationsReqBodyFiltersNearLocation`

NewSendFrameNotificationsReqBodyFiltersNearLocationWithDefaults instantiates a new SendFrameNotificationsReqBodyFiltersNearLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLatitude

`func (o *SendFrameNotificationsReqBodyFiltersNearLocation) GetLatitude() float64`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *SendFrameNotificationsReqBodyFiltersNearLocation) GetLatitudeOk() (*float64, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *SendFrameNotificationsReqBodyFiltersNearLocation) SetLatitude(v float64)`

SetLatitude sets Latitude field to given value.


### GetLongitude

`func (o *SendFrameNotificationsReqBodyFiltersNearLocation) GetLongitude() float64`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *SendFrameNotificationsReqBodyFiltersNearLocation) GetLongitudeOk() (*float64, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *SendFrameNotificationsReqBodyFiltersNearLocation) SetLongitude(v float64)`

SetLongitude sets Longitude field to given value.


### GetAddress

`func (o *SendFrameNotificationsReqBodyFiltersNearLocation) GetAddress() LocationAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *SendFrameNotificationsReqBodyFiltersNearLocation) GetAddressOk() (*LocationAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *SendFrameNotificationsReqBodyFiltersNearLocation) SetAddress(v LocationAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *SendFrameNotificationsReqBodyFiltersNearLocation) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetRadius

`func (o *SendFrameNotificationsReqBodyFiltersNearLocation) GetRadius() float32`

GetRadius returns the Radius field if non-nil, zero value otherwise.

### GetRadiusOk

`func (o *SendFrameNotificationsReqBodyFiltersNearLocation) GetRadiusOk() (*float32, bool)`

GetRadiusOk returns a tuple with the Radius field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadius

`func (o *SendFrameNotificationsReqBodyFiltersNearLocation) SetRadius(v float32)`

SetRadius sets Radius field to given value.

### HasRadius

`func (o *SendFrameNotificationsReqBodyFiltersNearLocation) HasRadius() bool`

HasRadius returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


