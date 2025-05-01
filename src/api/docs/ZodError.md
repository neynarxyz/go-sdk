# ZodError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** |  | 
**Code** | **string** |  | 
**Errors** | [**[]ZodErrorErrorsInner**](ZodErrorErrorsInner.md) |  | 

## Methods

### NewZodError

`func NewZodError(message string, code string, errors []ZodErrorErrorsInner, ) *ZodError`

NewZodError instantiates a new ZodError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZodErrorWithDefaults

`func NewZodErrorWithDefaults() *ZodError`

NewZodErrorWithDefaults instantiates a new ZodError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ZodError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ZodError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ZodError) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetCode

`func (o *ZodError) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ZodError) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ZodError) SetCode(v string)`

SetCode sets Code field to given value.


### GetErrors

`func (o *ZodError) GetErrors() []ZodErrorErrorsInner`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ZodError) GetErrorsOk() (*[]ZodErrorErrorsInner, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ZodError) SetErrors(v []ZodErrorErrorsInner)`

SetErrors sets Errors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


