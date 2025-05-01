# ValidatedFrameAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | **string** |  | 
**Url** | **string** |  | 
**Interactor** | [**User**](User.md) |  | 
**TappedButton** | [**ValidatedFrameActionTappedButton**](ValidatedFrameActionTappedButton.md) |  | 
**Input** | Pointer to [**FrameInput**](FrameInput.md) |  | [optional] 
**State** | [**FrameState**](FrameState.md) |  | 
**Cast** | [**CastWithInteractions**](CastWithInteractions.md) |  | 
**Timestamp** | **time.Time** |  | 
**Signer** | Pointer to [**ValidatedFrameActionSigner**](ValidatedFrameActionSigner.md) |  | [optional] 
**Transaction** | Pointer to [**FrameTransaction**](FrameTransaction.md) |  | [optional] 
**Address** | Pointer to **string** | The connected wallet address of the interacting user. | [optional] 

## Methods

### NewValidatedFrameAction

`func NewValidatedFrameAction(object string, url string, interactor User, tappedButton ValidatedFrameActionTappedButton, state FrameState, cast CastWithInteractions, timestamp time.Time, ) *ValidatedFrameAction`

NewValidatedFrameAction instantiates a new ValidatedFrameAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidatedFrameActionWithDefaults

`func NewValidatedFrameActionWithDefaults() *ValidatedFrameAction`

NewValidatedFrameActionWithDefaults instantiates a new ValidatedFrameAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *ValidatedFrameAction) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *ValidatedFrameAction) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *ValidatedFrameAction) SetObject(v string)`

SetObject sets Object field to given value.


### GetUrl

`func (o *ValidatedFrameAction) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ValidatedFrameAction) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ValidatedFrameAction) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetInteractor

`func (o *ValidatedFrameAction) GetInteractor() User`

GetInteractor returns the Interactor field if non-nil, zero value otherwise.

### GetInteractorOk

`func (o *ValidatedFrameAction) GetInteractorOk() (*User, bool)`

GetInteractorOk returns a tuple with the Interactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInteractor

`func (o *ValidatedFrameAction) SetInteractor(v User)`

SetInteractor sets Interactor field to given value.


### GetTappedButton

`func (o *ValidatedFrameAction) GetTappedButton() ValidatedFrameActionTappedButton`

GetTappedButton returns the TappedButton field if non-nil, zero value otherwise.

### GetTappedButtonOk

`func (o *ValidatedFrameAction) GetTappedButtonOk() (*ValidatedFrameActionTappedButton, bool)`

GetTappedButtonOk returns a tuple with the TappedButton field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTappedButton

`func (o *ValidatedFrameAction) SetTappedButton(v ValidatedFrameActionTappedButton)`

SetTappedButton sets TappedButton field to given value.


### GetInput

`func (o *ValidatedFrameAction) GetInput() FrameInput`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *ValidatedFrameAction) GetInputOk() (*FrameInput, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *ValidatedFrameAction) SetInput(v FrameInput)`

SetInput sets Input field to given value.

### HasInput

`func (o *ValidatedFrameAction) HasInput() bool`

HasInput returns a boolean if a field has been set.

### GetState

`func (o *ValidatedFrameAction) GetState() FrameState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ValidatedFrameAction) GetStateOk() (*FrameState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ValidatedFrameAction) SetState(v FrameState)`

SetState sets State field to given value.


### GetCast

`func (o *ValidatedFrameAction) GetCast() CastWithInteractions`

GetCast returns the Cast field if non-nil, zero value otherwise.

### GetCastOk

`func (o *ValidatedFrameAction) GetCastOk() (*CastWithInteractions, bool)`

GetCastOk returns a tuple with the Cast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCast

`func (o *ValidatedFrameAction) SetCast(v CastWithInteractions)`

SetCast sets Cast field to given value.


### GetTimestamp

`func (o *ValidatedFrameAction) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ValidatedFrameAction) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ValidatedFrameAction) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetSigner

`func (o *ValidatedFrameAction) GetSigner() ValidatedFrameActionSigner`

GetSigner returns the Signer field if non-nil, zero value otherwise.

### GetSignerOk

`func (o *ValidatedFrameAction) GetSignerOk() (*ValidatedFrameActionSigner, bool)`

GetSignerOk returns a tuple with the Signer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigner

`func (o *ValidatedFrameAction) SetSigner(v ValidatedFrameActionSigner)`

SetSigner sets Signer field to given value.

### HasSigner

`func (o *ValidatedFrameAction) HasSigner() bool`

HasSigner returns a boolean if a field has been set.

### GetTransaction

`func (o *ValidatedFrameAction) GetTransaction() FrameTransaction`

GetTransaction returns the Transaction field if non-nil, zero value otherwise.

### GetTransactionOk

`func (o *ValidatedFrameAction) GetTransactionOk() (*FrameTransaction, bool)`

GetTransactionOk returns a tuple with the Transaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransaction

`func (o *ValidatedFrameAction) SetTransaction(v FrameTransaction)`

SetTransaction sets Transaction field to given value.

### HasTransaction

`func (o *ValidatedFrameAction) HasTransaction() bool`

HasTransaction returns a boolean if a field has been set.

### GetAddress

`func (o *ValidatedFrameAction) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ValidatedFrameAction) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ValidatedFrameAction) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *ValidatedFrameAction) HasAddress() bool`

HasAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


