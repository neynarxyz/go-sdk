# AppHostPostEventBodyOneOf1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SignerUuid** | **string** | UUID of the signer. &#x60;signer_uuid&#x60; is paired with API key, can&#39;t use a &#x60;uuid&#x60; made with a different API key. | 
**AppDomain** | **string** | Domain of the mini app | 
**Fid** | **int32** | The unique identifier of a farcaster user or app (unsigned integer) | 
**Event** | [**AppHostEventType**](AppHostEventType.md) |  | 

## Methods

### NewAppHostPostEventBodyOneOf1

`func NewAppHostPostEventBodyOneOf1(signerUuid string, appDomain string, fid int32, event AppHostEventType, ) *AppHostPostEventBodyOneOf1`

NewAppHostPostEventBodyOneOf1 instantiates a new AppHostPostEventBodyOneOf1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppHostPostEventBodyOneOf1WithDefaults

`func NewAppHostPostEventBodyOneOf1WithDefaults() *AppHostPostEventBodyOneOf1`

NewAppHostPostEventBodyOneOf1WithDefaults instantiates a new AppHostPostEventBodyOneOf1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSignerUuid

`func (o *AppHostPostEventBodyOneOf1) GetSignerUuid() string`

GetSignerUuid returns the SignerUuid field if non-nil, zero value otherwise.

### GetSignerUuidOk

`func (o *AppHostPostEventBodyOneOf1) GetSignerUuidOk() (*string, bool)`

GetSignerUuidOk returns a tuple with the SignerUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignerUuid

`func (o *AppHostPostEventBodyOneOf1) SetSignerUuid(v string)`

SetSignerUuid sets SignerUuid field to given value.


### GetAppDomain

`func (o *AppHostPostEventBodyOneOf1) GetAppDomain() string`

GetAppDomain returns the AppDomain field if non-nil, zero value otherwise.

### GetAppDomainOk

`func (o *AppHostPostEventBodyOneOf1) GetAppDomainOk() (*string, bool)`

GetAppDomainOk returns a tuple with the AppDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppDomain

`func (o *AppHostPostEventBodyOneOf1) SetAppDomain(v string)`

SetAppDomain sets AppDomain field to given value.


### GetFid

`func (o *AppHostPostEventBodyOneOf1) GetFid() int32`

GetFid returns the Fid field if non-nil, zero value otherwise.

### GetFidOk

`func (o *AppHostPostEventBodyOneOf1) GetFidOk() (*int32, bool)`

GetFidOk returns a tuple with the Fid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFid

`func (o *AppHostPostEventBodyOneOf1) SetFid(v int32)`

SetFid sets Fid field to given value.


### GetEvent

`func (o *AppHostPostEventBodyOneOf1) GetEvent() AppHostEventType`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *AppHostPostEventBodyOneOf1) GetEventOk() (*AppHostEventType, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *AppHostPostEventBodyOneOf1) SetEvent(v AppHostEventType)`

SetEvent sets Event field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


