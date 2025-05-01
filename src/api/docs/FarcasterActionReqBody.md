# FarcasterActionReqBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SignerUuid** | **string** | The signer_uuid of the user on behalf of whom the action is being performed.  | 
**BaseUrl** | **string** | The base URL of the app on which the action is being performed.  | 
**Action** | [**FarcasterActionReqBodyAction**](FarcasterActionReqBodyAction.md) |  | 

## Methods

### NewFarcasterActionReqBody

`func NewFarcasterActionReqBody(signerUuid string, baseUrl string, action FarcasterActionReqBodyAction, ) *FarcasterActionReqBody`

NewFarcasterActionReqBody instantiates a new FarcasterActionReqBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFarcasterActionReqBodyWithDefaults

`func NewFarcasterActionReqBodyWithDefaults() *FarcasterActionReqBody`

NewFarcasterActionReqBodyWithDefaults instantiates a new FarcasterActionReqBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSignerUuid

`func (o *FarcasterActionReqBody) GetSignerUuid() string`

GetSignerUuid returns the SignerUuid field if non-nil, zero value otherwise.

### GetSignerUuidOk

`func (o *FarcasterActionReqBody) GetSignerUuidOk() (*string, bool)`

GetSignerUuidOk returns a tuple with the SignerUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignerUuid

`func (o *FarcasterActionReqBody) SetSignerUuid(v string)`

SetSignerUuid sets SignerUuid field to given value.


### GetBaseUrl

`func (o *FarcasterActionReqBody) GetBaseUrl() string`

GetBaseUrl returns the BaseUrl field if non-nil, zero value otherwise.

### GetBaseUrlOk

`func (o *FarcasterActionReqBody) GetBaseUrlOk() (*string, bool)`

GetBaseUrlOk returns a tuple with the BaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUrl

`func (o *FarcasterActionReqBody) SetBaseUrl(v string)`

SetBaseUrl sets BaseUrl field to given value.


### GetAction

`func (o *FarcasterActionReqBody) GetAction() FarcasterActionReqBodyAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *FarcasterActionReqBody) GetActionOk() (*FarcasterActionReqBodyAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *FarcasterActionReqBody) SetAction(v FarcasterActionReqBodyAction)`

SetAction sets Action field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


