# AppHostPostEventBodyOneOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SignedMessage** | **string** | JFS-signed message containing the event payload. The message must be properly signed and contain valid event information. | 
**AppDomain** | **string** | Domain of the mini app | 

## Methods

### NewAppHostPostEventBodyOneOf

`func NewAppHostPostEventBodyOneOf(signedMessage string, appDomain string, ) *AppHostPostEventBodyOneOf`

NewAppHostPostEventBodyOneOf instantiates a new AppHostPostEventBodyOneOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppHostPostEventBodyOneOfWithDefaults

`func NewAppHostPostEventBodyOneOfWithDefaults() *AppHostPostEventBodyOneOf`

NewAppHostPostEventBodyOneOfWithDefaults instantiates a new AppHostPostEventBodyOneOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSignedMessage

`func (o *AppHostPostEventBodyOneOf) GetSignedMessage() string`

GetSignedMessage returns the SignedMessage field if non-nil, zero value otherwise.

### GetSignedMessageOk

`func (o *AppHostPostEventBodyOneOf) GetSignedMessageOk() (*string, bool)`

GetSignedMessageOk returns a tuple with the SignedMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignedMessage

`func (o *AppHostPostEventBodyOneOf) SetSignedMessage(v string)`

SetSignedMessage sets SignedMessage field to given value.


### GetAppDomain

`func (o *AppHostPostEventBodyOneOf) GetAppDomain() string`

GetAppDomain returns the AppDomain field if non-nil, zero value otherwise.

### GetAppDomainOk

`func (o *AppHostPostEventBodyOneOf) GetAppDomainOk() (*string, bool)`

GetAppDomainOk returns a tuple with the AppDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppDomain

`func (o *AppHostPostEventBodyOneOf) SetAppDomain(v string)`

SetAppDomain sets AppDomain field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


