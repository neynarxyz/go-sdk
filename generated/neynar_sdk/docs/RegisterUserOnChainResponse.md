# RegisterUserOnChainResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**TransactionHash** | Pointer to **string** |  | [optional] 
**User** | Pointer to [**User**](User.md) |  | [optional] 

## Methods

### NewRegisterUserOnChainResponse

`func NewRegisterUserOnChainResponse() *RegisterUserOnChainResponse`

NewRegisterUserOnChainResponse instantiates a new RegisterUserOnChainResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterUserOnChainResponseWithDefaults

`func NewRegisterUserOnChainResponseWithDefaults() *RegisterUserOnChainResponse`

NewRegisterUserOnChainResponseWithDefaults instantiates a new RegisterUserOnChainResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *RegisterUserOnChainResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *RegisterUserOnChainResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *RegisterUserOnChainResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *RegisterUserOnChainResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetTransactionHash

`func (o *RegisterUserOnChainResponse) GetTransactionHash() string`

GetTransactionHash returns the TransactionHash field if non-nil, zero value otherwise.

### GetTransactionHashOk

`func (o *RegisterUserOnChainResponse) GetTransactionHashOk() (*string, bool)`

GetTransactionHashOk returns a tuple with the TransactionHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionHash

`func (o *RegisterUserOnChainResponse) SetTransactionHash(v string)`

SetTransactionHash sets TransactionHash field to given value.

### HasTransactionHash

`func (o *RegisterUserOnChainResponse) HasTransactionHash() bool`

HasTransactionHash returns a boolean if a field has been set.

### GetUser

`func (o *RegisterUserOnChainResponse) GetUser() User`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *RegisterUserOnChainResponse) GetUserOk() (*User, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *RegisterUserOnChainResponse) SetUser(v User)`

SetUser sets User field to given value.

### HasUser

`func (o *RegisterUserOnChainResponse) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


