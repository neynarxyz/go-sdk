# CastEmbeddedApp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | **string** |  | 
**Fid** | **int32** | The unique identifier of a farcaster user (unsigned integer) | 
**Username** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**PfpUrl** | Pointer to **string** |  | [optional] 
**CustodyAddress** | Pointer to **string** | Ethereum address | [optional] 

## Methods

### NewCastEmbeddedApp

`func NewCastEmbeddedApp(object string, fid int32, ) *CastEmbeddedApp`

NewCastEmbeddedApp instantiates a new CastEmbeddedApp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCastEmbeddedAppWithDefaults

`func NewCastEmbeddedAppWithDefaults() *CastEmbeddedApp`

NewCastEmbeddedAppWithDefaults instantiates a new CastEmbeddedApp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *CastEmbeddedApp) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *CastEmbeddedApp) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *CastEmbeddedApp) SetObject(v string)`

SetObject sets Object field to given value.


### GetFid

`func (o *CastEmbeddedApp) GetFid() int32`

GetFid returns the Fid field if non-nil, zero value otherwise.

### GetFidOk

`func (o *CastEmbeddedApp) GetFidOk() (*int32, bool)`

GetFidOk returns a tuple with the Fid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFid

`func (o *CastEmbeddedApp) SetFid(v int32)`

SetFid sets Fid field to given value.


### GetUsername

`func (o *CastEmbeddedApp) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CastEmbeddedApp) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CastEmbeddedApp) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *CastEmbeddedApp) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetDisplayName

`func (o *CastEmbeddedApp) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CastEmbeddedApp) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CastEmbeddedApp) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CastEmbeddedApp) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetPfpUrl

`func (o *CastEmbeddedApp) GetPfpUrl() string`

GetPfpUrl returns the PfpUrl field if non-nil, zero value otherwise.

### GetPfpUrlOk

`func (o *CastEmbeddedApp) GetPfpUrlOk() (*string, bool)`

GetPfpUrlOk returns a tuple with the PfpUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPfpUrl

`func (o *CastEmbeddedApp) SetPfpUrl(v string)`

SetPfpUrl sets PfpUrl field to given value.

### HasPfpUrl

`func (o *CastEmbeddedApp) HasPfpUrl() bool`

HasPfpUrl returns a boolean if a field has been set.

### GetCustodyAddress

`func (o *CastEmbeddedApp) GetCustodyAddress() string`

GetCustodyAddress returns the CustodyAddress field if non-nil, zero value otherwise.

### GetCustodyAddressOk

`func (o *CastEmbeddedApp) GetCustodyAddressOk() (*string, bool)`

GetCustodyAddressOk returns a tuple with the CustodyAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustodyAddress

`func (o *CastEmbeddedApp) SetCustodyAddress(v string)`

SetCustodyAddress sets CustodyAddress field to given value.

### HasCustodyAddress

`func (o *CastEmbeddedApp) HasCustodyAddress() bool`

HasCustodyAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


