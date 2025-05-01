# MuteResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** |  | 
**Message** | Pointer to **string** |  | [optional] 

## Methods

### NewMuteResponse

`func NewMuteResponse(success bool, ) *MuteResponse`

NewMuteResponse instantiates a new MuteResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMuteResponseWithDefaults

`func NewMuteResponseWithDefaults() *MuteResponse`

NewMuteResponseWithDefaults instantiates a new MuteResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *MuteResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *MuteResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *MuteResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetMessage

`func (o *MuteResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *MuteResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *MuteResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *MuteResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


