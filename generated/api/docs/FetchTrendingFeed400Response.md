# FetchTrendingFeed400Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** |  | 
**Code** | **string** |  | 
**Errors** | [**[]ZodErrorErrorsInner**](ZodErrorErrorsInner.md) |  | 
**Property** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **int32** |  | [optional] 

## Methods

### NewFetchTrendingFeed400Response

`func NewFetchTrendingFeed400Response(message string, code string, errors []ZodErrorErrorsInner, ) *FetchTrendingFeed400Response`

NewFetchTrendingFeed400Response instantiates a new FetchTrendingFeed400Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFetchTrendingFeed400ResponseWithDefaults

`func NewFetchTrendingFeed400ResponseWithDefaults() *FetchTrendingFeed400Response`

NewFetchTrendingFeed400ResponseWithDefaults instantiates a new FetchTrendingFeed400Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *FetchTrendingFeed400Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *FetchTrendingFeed400Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *FetchTrendingFeed400Response) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetCode

`func (o *FetchTrendingFeed400Response) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *FetchTrendingFeed400Response) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *FetchTrendingFeed400Response) SetCode(v string)`

SetCode sets Code field to given value.


### GetErrors

`func (o *FetchTrendingFeed400Response) GetErrors() []ZodErrorErrorsInner`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *FetchTrendingFeed400Response) GetErrorsOk() (*[]ZodErrorErrorsInner, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *FetchTrendingFeed400Response) SetErrors(v []ZodErrorErrorsInner)`

SetErrors sets Errors field to given value.


### GetProperty

`func (o *FetchTrendingFeed400Response) GetProperty() string`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *FetchTrendingFeed400Response) GetPropertyOk() (*string, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *FetchTrendingFeed400Response) SetProperty(v string)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *FetchTrendingFeed400Response) HasProperty() bool`

HasProperty returns a boolean if a field has been set.

### GetStatus

`func (o *FetchTrendingFeed400Response) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FetchTrendingFeed400Response) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FetchTrendingFeed400Response) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FetchTrendingFeed400Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


