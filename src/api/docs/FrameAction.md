# FrameAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Image** | Pointer to **string** |  | [optional] 
**Button** | [**FrameActionButton**](FrameActionButton.md) |  | 
**Input** | Pointer to [**FrameInput**](FrameInput.md) |  | [optional] 
**State** | Pointer to [**FrameState**](FrameState.md) |  | [optional] 
**Transaction** | Pointer to [**FrameTransaction**](FrameTransaction.md) |  | [optional] 
**Address** | Pointer to **string** | The connected wallet address of the interacting user. | [optional] 
**FramesUrl** | **string** | URL of the mini apps | 
**PostUrl** | **string** | URL of the post to get the next mini app | 

## Methods

### NewFrameAction

`func NewFrameAction(button FrameActionButton, framesUrl string, postUrl string, ) *FrameAction`

NewFrameAction instantiates a new FrameAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFrameActionWithDefaults

`func NewFrameActionWithDefaults() *FrameAction`

NewFrameActionWithDefaults instantiates a new FrameAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *FrameAction) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *FrameAction) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *FrameAction) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *FrameAction) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetTitle

`func (o *FrameAction) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *FrameAction) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *FrameAction) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *FrameAction) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetImage

`func (o *FrameAction) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *FrameAction) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *FrameAction) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *FrameAction) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetButton

`func (o *FrameAction) GetButton() FrameActionButton`

GetButton returns the Button field if non-nil, zero value otherwise.

### GetButtonOk

`func (o *FrameAction) GetButtonOk() (*FrameActionButton, bool)`

GetButtonOk returns a tuple with the Button field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButton

`func (o *FrameAction) SetButton(v FrameActionButton)`

SetButton sets Button field to given value.


### GetInput

`func (o *FrameAction) GetInput() FrameInput`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *FrameAction) GetInputOk() (*FrameInput, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *FrameAction) SetInput(v FrameInput)`

SetInput sets Input field to given value.

### HasInput

`func (o *FrameAction) HasInput() bool`

HasInput returns a boolean if a field has been set.

### GetState

`func (o *FrameAction) GetState() FrameState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *FrameAction) GetStateOk() (*FrameState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *FrameAction) SetState(v FrameState)`

SetState sets State field to given value.

### HasState

`func (o *FrameAction) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTransaction

`func (o *FrameAction) GetTransaction() FrameTransaction`

GetTransaction returns the Transaction field if non-nil, zero value otherwise.

### GetTransactionOk

`func (o *FrameAction) GetTransactionOk() (*FrameTransaction, bool)`

GetTransactionOk returns a tuple with the Transaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransaction

`func (o *FrameAction) SetTransaction(v FrameTransaction)`

SetTransaction sets Transaction field to given value.

### HasTransaction

`func (o *FrameAction) HasTransaction() bool`

HasTransaction returns a boolean if a field has been set.

### GetAddress

`func (o *FrameAction) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *FrameAction) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *FrameAction) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *FrameAction) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetFramesUrl

`func (o *FrameAction) GetFramesUrl() string`

GetFramesUrl returns the FramesUrl field if non-nil, zero value otherwise.

### GetFramesUrlOk

`func (o *FrameAction) GetFramesUrlOk() (*string, bool)`

GetFramesUrlOk returns a tuple with the FramesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFramesUrl

`func (o *FrameAction) SetFramesUrl(v string)`

SetFramesUrl sets FramesUrl field to given value.


### GetPostUrl

`func (o *FrameAction) GetPostUrl() string`

GetPostUrl returns the PostUrl field if non-nil, zero value otherwise.

### GetPostUrlOk

`func (o *FrameAction) GetPostUrlOk() (*string, bool)`

GetPostUrlOk returns a tuple with the PostUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostUrl

`func (o *FrameAction) SetPostUrl(v string)`

SetPostUrl sets PostUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


