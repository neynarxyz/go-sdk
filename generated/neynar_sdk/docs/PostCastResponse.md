# PostCastResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** |  | 
**Cast** | [**PostCastResponseCast**](PostCastResponseCast.md) |  | 

## Methods

### NewPostCastResponse

`func NewPostCastResponse(success bool, cast PostCastResponseCast, ) *PostCastResponse`

NewPostCastResponse instantiates a new PostCastResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostCastResponseWithDefaults

`func NewPostCastResponseWithDefaults() *PostCastResponse`

NewPostCastResponseWithDefaults instantiates a new PostCastResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *PostCastResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *PostCastResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *PostCastResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetCast

`func (o *PostCastResponse) GetCast() PostCastResponseCast`

GetCast returns the Cast field if non-nil, zero value otherwise.

### GetCastOk

`func (o *PostCastResponse) GetCastOk() (*PostCastResponseCast, bool)`

GetCastOk returns a tuple with the Cast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCast

`func (o *PostCastResponse) SetCast(v PostCastResponseCast)`

SetCast sets Cast field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


