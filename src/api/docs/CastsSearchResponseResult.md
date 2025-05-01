# CastsSearchResponseResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Casts** | [**[]CastWithInteractions**](CastWithInteractions.md) |  | 
**Next** | [**NextCursor**](NextCursor.md) |  | 

## Methods

### NewCastsSearchResponseResult

`func NewCastsSearchResponseResult(casts []CastWithInteractions, next NextCursor, ) *CastsSearchResponseResult`

NewCastsSearchResponseResult instantiates a new CastsSearchResponseResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCastsSearchResponseResultWithDefaults

`func NewCastsSearchResponseResultWithDefaults() *CastsSearchResponseResult`

NewCastsSearchResponseResultWithDefaults instantiates a new CastsSearchResponseResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCasts

`func (o *CastsSearchResponseResult) GetCasts() []CastWithInteractions`

GetCasts returns the Casts field if non-nil, zero value otherwise.

### GetCastsOk

`func (o *CastsSearchResponseResult) GetCastsOk() (*[]CastWithInteractions, bool)`

GetCastsOk returns a tuple with the Casts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCasts

`func (o *CastsSearchResponseResult) SetCasts(v []CastWithInteractions)`

SetCasts sets Casts field to given value.


### GetNext

`func (o *CastsSearchResponseResult) GetNext() NextCursor`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *CastsSearchResponseResult) GetNextOk() (*NextCursor, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *CastsSearchResponseResult) SetNext(v NextCursor)`

SetNext sets Next field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


