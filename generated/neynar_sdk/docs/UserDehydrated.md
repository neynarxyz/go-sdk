# UserDehydrated

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

### NewUserDehydrated

`func NewUserDehydrated(object string, fid int32, ) *UserDehydrated`

NewUserDehydrated instantiates a new UserDehydrated object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserDehydratedWithDefaults

`func NewUserDehydratedWithDefaults() *UserDehydrated`

NewUserDehydratedWithDefaults instantiates a new UserDehydrated object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *UserDehydrated) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *UserDehydrated) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *UserDehydrated) SetObject(v string)`

SetObject sets Object field to given value.


### GetFid

`func (o *UserDehydrated) GetFid() int32`

GetFid returns the Fid field if non-nil, zero value otherwise.

### GetFidOk

`func (o *UserDehydrated) GetFidOk() (*int32, bool)`

GetFidOk returns a tuple with the Fid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFid

`func (o *UserDehydrated) SetFid(v int32)`

SetFid sets Fid field to given value.


### GetUsername

`func (o *UserDehydrated) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UserDehydrated) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UserDehydrated) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *UserDehydrated) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetDisplayName

`func (o *UserDehydrated) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *UserDehydrated) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *UserDehydrated) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *UserDehydrated) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetPfpUrl

`func (o *UserDehydrated) GetPfpUrl() string`

GetPfpUrl returns the PfpUrl field if non-nil, zero value otherwise.

### GetPfpUrlOk

`func (o *UserDehydrated) GetPfpUrlOk() (*string, bool)`

GetPfpUrlOk returns a tuple with the PfpUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPfpUrl

`func (o *UserDehydrated) SetPfpUrl(v string)`

SetPfpUrl sets PfpUrl field to given value.

### HasPfpUrl

`func (o *UserDehydrated) HasPfpUrl() bool`

HasPfpUrl returns a boolean if a field has been set.

### GetCustodyAddress

`func (o *UserDehydrated) GetCustodyAddress() string`

GetCustodyAddress returns the CustodyAddress field if non-nil, zero value otherwise.

### GetCustodyAddressOk

`func (o *UserDehydrated) GetCustodyAddressOk() (*string, bool)`

GetCustodyAddressOk returns a tuple with the CustodyAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustodyAddress

`func (o *UserDehydrated) SetCustodyAddress(v string)`

SetCustodyAddress sets CustodyAddress field to given value.

### HasCustodyAddress

`func (o *UserDehydrated) HasCustodyAddress() bool`

HasCustodyAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


