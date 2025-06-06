# CastViewerContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Liked** | **bool** | Indicates if the viewer liked the cast. | 
**Recasted** | **bool** | Indicates if the viewer recasted the cast. | 

## Methods

### NewCastViewerContext

`func NewCastViewerContext(liked bool, recasted bool, ) *CastViewerContext`

NewCastViewerContext instantiates a new CastViewerContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCastViewerContextWithDefaults

`func NewCastViewerContextWithDefaults() *CastViewerContext`

NewCastViewerContextWithDefaults instantiates a new CastViewerContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLiked

`func (o *CastViewerContext) GetLiked() bool`

GetLiked returns the Liked field if non-nil, zero value otherwise.

### GetLikedOk

`func (o *CastViewerContext) GetLikedOk() (*bool, bool)`

GetLikedOk returns a tuple with the Liked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiked

`func (o *CastViewerContext) SetLiked(v bool)`

SetLiked sets Liked field to given value.


### GetRecasted

`func (o *CastViewerContext) GetRecasted() bool`

GetRecasted returns the Recasted field if non-nil, zero value otherwise.

### GetRecastedOk

`func (o *CastViewerContext) GetRecastedOk() (*bool, bool)`

GetRecastedOk returns a tuple with the Recasted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecasted

`func (o *CastViewerContext) SetRecasted(v bool)`

SetRecasted sets Recasted field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


