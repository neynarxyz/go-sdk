# RegisterUserResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** |  | 
**Message** | **string** |  | 
**Signer** | [**Signer**](Signer.md) |  | 
**User** | Pointer to [**User**](User.md) |  | [optional] 

## Methods

### NewRegisterUserResponse

`func NewRegisterUserResponse(success bool, message string, signer Signer, ) *RegisterUserResponse`

NewRegisterUserResponse instantiates a new RegisterUserResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterUserResponseWithDefaults

`func NewRegisterUserResponseWithDefaults() *RegisterUserResponse`

NewRegisterUserResponseWithDefaults instantiates a new RegisterUserResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *RegisterUserResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *RegisterUserResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *RegisterUserResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetMessage

`func (o *RegisterUserResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *RegisterUserResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *RegisterUserResponse) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetSigner

`func (o *RegisterUserResponse) GetSigner() Signer`

GetSigner returns the Signer field if non-nil, zero value otherwise.

### GetSignerOk

`func (o *RegisterUserResponse) GetSignerOk() (*Signer, bool)`

GetSignerOk returns a tuple with the Signer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigner

`func (o *RegisterUserResponse) SetSigner(v Signer)`

SetSigner sets Signer field to given value.


### GetUser

`func (o *RegisterUserResponse) GetUser() User`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *RegisterUserResponse) GetUserOk() (*User, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *RegisterUserResponse) SetUser(v User)`

SetUser sets User field to given value.

### HasUser

`func (o *RegisterUserResponse) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


