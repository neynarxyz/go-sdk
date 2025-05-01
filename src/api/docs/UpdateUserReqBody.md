# UpdateUserReqBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SignerUuid** | **string** | UUID of the signer. &#x60;signer_uuid&#x60; is paired with API key, can&#39;t use a &#x60;uuid&#x60; made with a different API key.  | 
**Bio** | Pointer to **string** |  | [optional] 
**PfpUrl** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Location** | Pointer to [**UpdateUserReqBodyLocation**](UpdateUserReqBodyLocation.md) |  | [optional] 
**VerifiedAccounts** | Pointer to [**UpdateUserReqBodyVerifiedAccounts**](UpdateUserReqBodyVerifiedAccounts.md) |  | [optional] 

## Methods

### NewUpdateUserReqBody

`func NewUpdateUserReqBody(signerUuid string, ) *UpdateUserReqBody`

NewUpdateUserReqBody instantiates a new UpdateUserReqBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateUserReqBodyWithDefaults

`func NewUpdateUserReqBodyWithDefaults() *UpdateUserReqBody`

NewUpdateUserReqBodyWithDefaults instantiates a new UpdateUserReqBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSignerUuid

`func (o *UpdateUserReqBody) GetSignerUuid() string`

GetSignerUuid returns the SignerUuid field if non-nil, zero value otherwise.

### GetSignerUuidOk

`func (o *UpdateUserReqBody) GetSignerUuidOk() (*string, bool)`

GetSignerUuidOk returns a tuple with the SignerUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignerUuid

`func (o *UpdateUserReqBody) SetSignerUuid(v string)`

SetSignerUuid sets SignerUuid field to given value.


### GetBio

`func (o *UpdateUserReqBody) GetBio() string`

GetBio returns the Bio field if non-nil, zero value otherwise.

### GetBioOk

`func (o *UpdateUserReqBody) GetBioOk() (*string, bool)`

GetBioOk returns a tuple with the Bio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBio

`func (o *UpdateUserReqBody) SetBio(v string)`

SetBio sets Bio field to given value.

### HasBio

`func (o *UpdateUserReqBody) HasBio() bool`

HasBio returns a boolean if a field has been set.

### GetPfpUrl

`func (o *UpdateUserReqBody) GetPfpUrl() string`

GetPfpUrl returns the PfpUrl field if non-nil, zero value otherwise.

### GetPfpUrlOk

`func (o *UpdateUserReqBody) GetPfpUrlOk() (*string, bool)`

GetPfpUrlOk returns a tuple with the PfpUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPfpUrl

`func (o *UpdateUserReqBody) SetPfpUrl(v string)`

SetPfpUrl sets PfpUrl field to given value.

### HasPfpUrl

`func (o *UpdateUserReqBody) HasPfpUrl() bool`

HasPfpUrl returns a boolean if a field has been set.

### GetUrl

`func (o *UpdateUserReqBody) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *UpdateUserReqBody) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *UpdateUserReqBody) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *UpdateUserReqBody) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetUsername

`func (o *UpdateUserReqBody) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UpdateUserReqBody) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UpdateUserReqBody) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *UpdateUserReqBody) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetDisplayName

`func (o *UpdateUserReqBody) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *UpdateUserReqBody) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *UpdateUserReqBody) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *UpdateUserReqBody) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetLocation

`func (o *UpdateUserReqBody) GetLocation() UpdateUserReqBodyLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *UpdateUserReqBody) GetLocationOk() (*UpdateUserReqBodyLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *UpdateUserReqBody) SetLocation(v UpdateUserReqBodyLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *UpdateUserReqBody) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetVerifiedAccounts

`func (o *UpdateUserReqBody) GetVerifiedAccounts() UpdateUserReqBodyVerifiedAccounts`

GetVerifiedAccounts returns the VerifiedAccounts field if non-nil, zero value otherwise.

### GetVerifiedAccountsOk

`func (o *UpdateUserReqBody) GetVerifiedAccountsOk() (*UpdateUserReqBodyVerifiedAccounts, bool)`

GetVerifiedAccountsOk returns a tuple with the VerifiedAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifiedAccounts

`func (o *UpdateUserReqBody) SetVerifiedAccounts(v UpdateUserReqBodyVerifiedAccounts)`

SetVerifiedAccounts sets VerifiedAccounts field to given value.

### HasVerifiedAccounts

`func (o *UpdateUserReqBody) HasVerifiedAccounts() bool`

HasVerifiedAccounts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


