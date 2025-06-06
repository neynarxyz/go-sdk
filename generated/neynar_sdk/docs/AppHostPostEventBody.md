# AppHostPostEventBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SignedMessage** | **string** | JFS-signed message containing the event payload. The message must be properly signed and contain valid event information. | 
**AppDomain** | **string** | Domain of the mini app | 
**SignerUuid** | **string** | UUID of the signer. &#x60;signer_uuid&#x60; is paired with API key, can&#39;t use a &#x60;uuid&#x60; made with a different API key. | 
**Fid** | **int32** | The unique identifier of a farcaster user or app (unsigned integer) | 
**Event** | [**AppHostEventType**](AppHostEventType.md) |  | 

## Methods

### NewAppHostPostEventBody

`func NewAppHostPostEventBody(signedMessage string, appDomain string, signerUuid string, fid int32, event AppHostEventType, ) *AppHostPostEventBody`

NewAppHostPostEventBody instantiates a new AppHostPostEventBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppHostPostEventBodyWithDefaults

`func NewAppHostPostEventBodyWithDefaults() *AppHostPostEventBody`

NewAppHostPostEventBodyWithDefaults instantiates a new AppHostPostEventBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSignedMessage

`func (o *AppHostPostEventBody) GetSignedMessage() string`

GetSignedMessage returns the SignedMessage field if non-nil, zero value otherwise.

### GetSignedMessageOk

`func (o *AppHostPostEventBody) GetSignedMessageOk() (*string, bool)`

GetSignedMessageOk returns a tuple with the SignedMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignedMessage

`func (o *AppHostPostEventBody) SetSignedMessage(v string)`

SetSignedMessage sets SignedMessage field to given value.


### GetAppDomain

`func (o *AppHostPostEventBody) GetAppDomain() string`

GetAppDomain returns the AppDomain field if non-nil, zero value otherwise.

### GetAppDomainOk

`func (o *AppHostPostEventBody) GetAppDomainOk() (*string, bool)`

GetAppDomainOk returns a tuple with the AppDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppDomain

`func (o *AppHostPostEventBody) SetAppDomain(v string)`

SetAppDomain sets AppDomain field to given value.


### GetSignerUuid

`func (o *AppHostPostEventBody) GetSignerUuid() string`

GetSignerUuid returns the SignerUuid field if non-nil, zero value otherwise.

### GetSignerUuidOk

`func (o *AppHostPostEventBody) GetSignerUuidOk() (*string, bool)`

GetSignerUuidOk returns a tuple with the SignerUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignerUuid

`func (o *AppHostPostEventBody) SetSignerUuid(v string)`

SetSignerUuid sets SignerUuid field to given value.


### GetFid

`func (o *AppHostPostEventBody) GetFid() int32`

GetFid returns the Fid field if non-nil, zero value otherwise.

### GetFidOk

`func (o *AppHostPostEventBody) GetFidOk() (*int32, bool)`

GetFidOk returns a tuple with the Fid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFid

`func (o *AppHostPostEventBody) SetFid(v int32)`

SetFid sets Fid field to given value.


### GetEvent

`func (o *AppHostPostEventBody) GetEvent() AppHostEventType`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *AppHostPostEventBody) GetEventOk() (*AppHostEventType, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *AppHostPostEventBody) SetEvent(v AppHostEventType)`

SetEvent sets Event field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


