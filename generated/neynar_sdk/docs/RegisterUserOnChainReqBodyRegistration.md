# RegisterUserOnChainReqBodyRegistration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Deadline** | **int32** |  | 
**Signature** | **string** | Hexadecimal number expressed as string with &#39;0x&#39; prefix | 
**CustodyAddress** | **string** | Ethereum address | 
**RecoveryAddress** | **string** | Ethereum address | 

## Methods

### NewRegisterUserOnChainReqBodyRegistration

`func NewRegisterUserOnChainReqBodyRegistration(deadline int32, signature string, custodyAddress string, recoveryAddress string, ) *RegisterUserOnChainReqBodyRegistration`

NewRegisterUserOnChainReqBodyRegistration instantiates a new RegisterUserOnChainReqBodyRegistration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterUserOnChainReqBodyRegistrationWithDefaults

`func NewRegisterUserOnChainReqBodyRegistrationWithDefaults() *RegisterUserOnChainReqBodyRegistration`

NewRegisterUserOnChainReqBodyRegistrationWithDefaults instantiates a new RegisterUserOnChainReqBodyRegistration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeadline

`func (o *RegisterUserOnChainReqBodyRegistration) GetDeadline() int32`

GetDeadline returns the Deadline field if non-nil, zero value otherwise.

### GetDeadlineOk

`func (o *RegisterUserOnChainReqBodyRegistration) GetDeadlineOk() (*int32, bool)`

GetDeadlineOk returns a tuple with the Deadline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadline

`func (o *RegisterUserOnChainReqBodyRegistration) SetDeadline(v int32)`

SetDeadline sets Deadline field to given value.


### GetSignature

`func (o *RegisterUserOnChainReqBodyRegistration) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *RegisterUserOnChainReqBodyRegistration) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *RegisterUserOnChainReqBodyRegistration) SetSignature(v string)`

SetSignature sets Signature field to given value.


### GetCustodyAddress

`func (o *RegisterUserOnChainReqBodyRegistration) GetCustodyAddress() string`

GetCustodyAddress returns the CustodyAddress field if non-nil, zero value otherwise.

### GetCustodyAddressOk

`func (o *RegisterUserOnChainReqBodyRegistration) GetCustodyAddressOk() (*string, bool)`

GetCustodyAddressOk returns a tuple with the CustodyAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustodyAddress

`func (o *RegisterUserOnChainReqBodyRegistration) SetCustodyAddress(v string)`

SetCustodyAddress sets CustodyAddress field to given value.


### GetRecoveryAddress

`func (o *RegisterUserOnChainReqBodyRegistration) GetRecoveryAddress() string`

GetRecoveryAddress returns the RecoveryAddress field if non-nil, zero value otherwise.

### GetRecoveryAddressOk

`func (o *RegisterUserOnChainReqBodyRegistration) GetRecoveryAddressOk() (*string, bool)`

GetRecoveryAddressOk returns a tuple with the RecoveryAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryAddress

`func (o *RegisterUserOnChainReqBodyRegistration) SetRecoveryAddress(v string)`

SetRecoveryAddress sets RecoveryAddress field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


