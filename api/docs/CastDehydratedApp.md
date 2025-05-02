# CastDehydratedApp

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

### NewCastDehydratedApp

`func NewCastDehydratedApp(object string, fid int32, ) *CastDehydratedApp`

NewCastDehydratedApp instantiates a new CastDehydratedApp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCastDehydratedAppWithDefaults

`func NewCastDehydratedAppWithDefaults() *CastDehydratedApp`

NewCastDehydratedAppWithDefaults instantiates a new CastDehydratedApp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *CastDehydratedApp) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *CastDehydratedApp) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *CastDehydratedApp) SetObject(v string)`

SetObject sets Object field to given value.


### GetFid

`func (o *CastDehydratedApp) GetFid() int32`

GetFid returns the Fid field if non-nil, zero value otherwise.

### GetFidOk

`func (o *CastDehydratedApp) GetFidOk() (*int32, bool)`

GetFidOk returns a tuple with the Fid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFid

`func (o *CastDehydratedApp) SetFid(v int32)`

SetFid sets Fid field to given value.


### GetUsername

`func (o *CastDehydratedApp) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CastDehydratedApp) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CastDehydratedApp) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *CastDehydratedApp) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetDisplayName

`func (o *CastDehydratedApp) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CastDehydratedApp) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CastDehydratedApp) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CastDehydratedApp) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetPfpUrl

`func (o *CastDehydratedApp) GetPfpUrl() string`

GetPfpUrl returns the PfpUrl field if non-nil, zero value otherwise.

### GetPfpUrlOk

`func (o *CastDehydratedApp) GetPfpUrlOk() (*string, bool)`

GetPfpUrlOk returns a tuple with the PfpUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPfpUrl

`func (o *CastDehydratedApp) SetPfpUrl(v string)`

SetPfpUrl sets PfpUrl field to given value.

### HasPfpUrl

`func (o *CastDehydratedApp) HasPfpUrl() bool`

HasPfpUrl returns a boolean if a field has been set.

### GetCustodyAddress

`func (o *CastDehydratedApp) GetCustodyAddress() string`

GetCustodyAddress returns the CustodyAddress field if non-nil, zero value otherwise.

### GetCustodyAddressOk

`func (o *CastDehydratedApp) GetCustodyAddressOk() (*string, bool)`

GetCustodyAddressOk returns a tuple with the CustodyAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustodyAddress

`func (o *CastDehydratedApp) SetCustodyAddress(v string)`

SetCustodyAddress sets CustodyAddress field to given value.

### HasCustodyAddress

`func (o *CastDehydratedApp) HasCustodyAddress() bool`

HasCustodyAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


