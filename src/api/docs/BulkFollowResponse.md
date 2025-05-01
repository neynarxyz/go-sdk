# BulkFollowResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** |  | 
**Details** | [**[]FollowResponse**](FollowResponse.md) |  | 

## Methods

### NewBulkFollowResponse

`func NewBulkFollowResponse(success bool, details []FollowResponse, ) *BulkFollowResponse`

NewBulkFollowResponse instantiates a new BulkFollowResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkFollowResponseWithDefaults

`func NewBulkFollowResponseWithDefaults() *BulkFollowResponse`

NewBulkFollowResponseWithDefaults instantiates a new BulkFollowResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *BulkFollowResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *BulkFollowResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *BulkFollowResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetDetails

`func (o *BulkFollowResponse) GetDetails() []FollowResponse`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *BulkFollowResponse) GetDetailsOk() (*[]FollowResponse, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *BulkFollowResponse) SetDetails(v []FollowResponse)`

SetDetails sets Details field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


