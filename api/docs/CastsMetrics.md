# CastsMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Start** | **time.Time** |  | 
**ResolutionInSeconds** | **int32** |  | 
**CastCount** | **int32** |  | 

## Methods

### NewCastsMetrics

`func NewCastsMetrics(start time.Time, resolutionInSeconds int32, castCount int32, ) *CastsMetrics`

NewCastsMetrics instantiates a new CastsMetrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCastsMetricsWithDefaults

`func NewCastsMetricsWithDefaults() *CastsMetrics`

NewCastsMetricsWithDefaults instantiates a new CastsMetrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStart

`func (o *CastsMetrics) GetStart() time.Time`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *CastsMetrics) GetStartOk() (*time.Time, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *CastsMetrics) SetStart(v time.Time)`

SetStart sets Start field to given value.


### GetResolutionInSeconds

`func (o *CastsMetrics) GetResolutionInSeconds() int32`

GetResolutionInSeconds returns the ResolutionInSeconds field if non-nil, zero value otherwise.

### GetResolutionInSecondsOk

`func (o *CastsMetrics) GetResolutionInSecondsOk() (*int32, bool)`

GetResolutionInSecondsOk returns a tuple with the ResolutionInSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolutionInSeconds

`func (o *CastsMetrics) SetResolutionInSeconds(v int32)`

SetResolutionInSeconds sets ResolutionInSeconds field to given value.


### GetCastCount

`func (o *CastsMetrics) GetCastCount() int32`

GetCastCount returns the CastCount field if non-nil, zero value otherwise.

### GetCastCountOk

`func (o *CastsMetrics) GetCastCountOk() (*int32, bool)`

GetCastCountOk returns a tuple with the CastCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCastCount

`func (o *CastsMetrics) SetCastCount(v int32)`

SetCastCount sets CastCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


