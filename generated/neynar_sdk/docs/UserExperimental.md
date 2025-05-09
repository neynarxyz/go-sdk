# UserExperimental

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeprecationNotice** | Pointer to **string** |  | [optional] 
**NeynarUserScore** | **float64** | Score that represents the probability that the account is not spam. | 

## Methods

### NewUserExperimental

`func NewUserExperimental(neynarUserScore float64, ) *UserExperimental`

NewUserExperimental instantiates a new UserExperimental object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserExperimentalWithDefaults

`func NewUserExperimentalWithDefaults() *UserExperimental`

NewUserExperimentalWithDefaults instantiates a new UserExperimental object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeprecationNotice

`func (o *UserExperimental) GetDeprecationNotice() string`

GetDeprecationNotice returns the DeprecationNotice field if non-nil, zero value otherwise.

### GetDeprecationNoticeOk

`func (o *UserExperimental) GetDeprecationNoticeOk() (*string, bool)`

GetDeprecationNoticeOk returns a tuple with the DeprecationNotice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecationNotice

`func (o *UserExperimental) SetDeprecationNotice(v string)`

SetDeprecationNotice sets DeprecationNotice field to given value.

### HasDeprecationNotice

`func (o *UserExperimental) HasDeprecationNotice() bool`

HasDeprecationNotice returns a boolean if a field has been set.

### GetNeynarUserScore

`func (o *UserExperimental) GetNeynarUserScore() float64`

GetNeynarUserScore returns the NeynarUserScore field if non-nil, zero value otherwise.

### GetNeynarUserScoreOk

`func (o *UserExperimental) GetNeynarUserScoreOk() (*float64, bool)`

GetNeynarUserScoreOk returns a tuple with the NeynarUserScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeynarUserScore

`func (o *UserExperimental) SetNeynarUserScore(v float64)`

SetNeynarUserScore sets NeynarUserScore field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


