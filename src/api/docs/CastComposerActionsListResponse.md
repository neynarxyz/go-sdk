# CastComposerActionsListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | Pointer to [**[]CastComposerActionsListResponseActionsInner**](CastComposerActionsListResponseActionsInner.md) |  | [optional] 
**Next** | Pointer to [**NextCursor**](NextCursor.md) |  | [optional] 

## Methods

### NewCastComposerActionsListResponse

`func NewCastComposerActionsListResponse() *CastComposerActionsListResponse`

NewCastComposerActionsListResponse instantiates a new CastComposerActionsListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCastComposerActionsListResponseWithDefaults

`func NewCastComposerActionsListResponseWithDefaults() *CastComposerActionsListResponse`

NewCastComposerActionsListResponseWithDefaults instantiates a new CastComposerActionsListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *CastComposerActionsListResponse) GetActions() []CastComposerActionsListResponseActionsInner`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *CastComposerActionsListResponse) GetActionsOk() (*[]CastComposerActionsListResponseActionsInner, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *CastComposerActionsListResponse) SetActions(v []CastComposerActionsListResponseActionsInner)`

SetActions sets Actions field to given value.

### HasActions

`func (o *CastComposerActionsListResponse) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetNext

`func (o *CastComposerActionsListResponse) GetNext() NextCursor`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *CastComposerActionsListResponse) GetNextOk() (*NextCursor, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *CastComposerActionsListResponse) SetNext(v NextCursor)`

SetNext sets Next field to given value.

### HasNext

`func (o *CastComposerActionsListResponse) HasNext() bool`

HasNext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


