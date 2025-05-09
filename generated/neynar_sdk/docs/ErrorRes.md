# ErrorRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** |  | [optional] 
**Message** | **string** |  | 
**Property** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **int32** |  | [optional] 

## Methods

### NewErrorRes

`func NewErrorRes(message string, ) *ErrorRes`

NewErrorRes instantiates a new ErrorRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorResWithDefaults

`func NewErrorResWithDefaults() *ErrorRes`

NewErrorResWithDefaults instantiates a new ErrorRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ErrorRes) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ErrorRes) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ErrorRes) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ErrorRes) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessage

`func (o *ErrorRes) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ErrorRes) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ErrorRes) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetProperty

`func (o *ErrorRes) GetProperty() string`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *ErrorRes) GetPropertyOk() (*string, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *ErrorRes) SetProperty(v string)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *ErrorRes) HasProperty() bool`

HasProperty returns a boolean if a field has been set.

### GetStatus

`func (o *ErrorRes) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ErrorRes) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ErrorRes) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ErrorRes) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


