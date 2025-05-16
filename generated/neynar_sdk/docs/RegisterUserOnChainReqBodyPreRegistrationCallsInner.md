# RegisterUserOnChainReqBodyPreRegistrationCallsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Target** | **string** | Ethereum address | 
**Value** | Pointer to **int32** | Value in wei to send with the transaction. This is not the amount of ETH that will be sent, but rather the value of the transaction. | [optional] [default to 0]
**Data** | **string** | Hexadecimal number expressed as string with &#39;0x&#39; prefix | 
**AllowFailure** | Pointer to **bool** | Set it to true if you want to ignore the failure of this call. If set to false, the registration will fail if this call fails. | [optional] [default to false]

## Methods

### NewRegisterUserOnChainReqBodyPreRegistrationCallsInner

`func NewRegisterUserOnChainReqBodyPreRegistrationCallsInner(target string, data string, ) *RegisterUserOnChainReqBodyPreRegistrationCallsInner`

NewRegisterUserOnChainReqBodyPreRegistrationCallsInner instantiates a new RegisterUserOnChainReqBodyPreRegistrationCallsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterUserOnChainReqBodyPreRegistrationCallsInnerWithDefaults

`func NewRegisterUserOnChainReqBodyPreRegistrationCallsInnerWithDefaults() *RegisterUserOnChainReqBodyPreRegistrationCallsInner`

NewRegisterUserOnChainReqBodyPreRegistrationCallsInnerWithDefaults instantiates a new RegisterUserOnChainReqBodyPreRegistrationCallsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTarget

`func (o *RegisterUserOnChainReqBodyPreRegistrationCallsInner) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *RegisterUserOnChainReqBodyPreRegistrationCallsInner) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *RegisterUserOnChainReqBodyPreRegistrationCallsInner) SetTarget(v string)`

SetTarget sets Target field to given value.


### GetValue

`func (o *RegisterUserOnChainReqBodyPreRegistrationCallsInner) GetValue() int32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *RegisterUserOnChainReqBodyPreRegistrationCallsInner) GetValueOk() (*int32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *RegisterUserOnChainReqBodyPreRegistrationCallsInner) SetValue(v int32)`

SetValue sets Value field to given value.

### HasValue

`func (o *RegisterUserOnChainReqBodyPreRegistrationCallsInner) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetData

`func (o *RegisterUserOnChainReqBodyPreRegistrationCallsInner) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *RegisterUserOnChainReqBodyPreRegistrationCallsInner) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *RegisterUserOnChainReqBodyPreRegistrationCallsInner) SetData(v string)`

SetData sets Data field to given value.


### GetAllowFailure

`func (o *RegisterUserOnChainReqBodyPreRegistrationCallsInner) GetAllowFailure() bool`

GetAllowFailure returns the AllowFailure field if non-nil, zero value otherwise.

### GetAllowFailureOk

`func (o *RegisterUserOnChainReqBodyPreRegistrationCallsInner) GetAllowFailureOk() (*bool, bool)`

GetAllowFailureOk returns a tuple with the AllowFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowFailure

`func (o *RegisterUserOnChainReqBodyPreRegistrationCallsInner) SetAllowFailure(v bool)`

SetAllowFailure sets AllowFailure field to given value.

### HasAllowFailure

`func (o *RegisterUserOnChainReqBodyPreRegistrationCallsInner) HasAllowFailure() bool`

HasAllowFailure returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


