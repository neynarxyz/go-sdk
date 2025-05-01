# VerificationRemoveBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | The Ethereum address (0x-prefixed) for which the verification should be removed. Must match a previously verified address in the user&#39;s profile. | 

## Methods

### NewVerificationRemoveBody

`func NewVerificationRemoveBody(address string, ) *VerificationRemoveBody`

NewVerificationRemoveBody instantiates a new VerificationRemoveBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerificationRemoveBodyWithDefaults

`func NewVerificationRemoveBodyWithDefaults() *VerificationRemoveBody`

NewVerificationRemoveBodyWithDefaults instantiates a new VerificationRemoveBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *VerificationRemoveBody) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *VerificationRemoveBody) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *VerificationRemoveBody) SetAddress(v string)`

SetAddress sets Address field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


