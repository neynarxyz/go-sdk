# FrameActionBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | The URL associated with the frame action. This typically points to the frame&#39;s content or the destination that handles the interaction. | 
**ButtonIndex** | **int32** | Identifies which button the user clicked in the frame. Frames can have up to 4 buttons, numbered from 1 to 4. | 
**CastId** | [**CastId**](CastId.md) |  | 

## Methods

### NewFrameActionBody

`func NewFrameActionBody(url string, buttonIndex int32, castId CastId, ) *FrameActionBody`

NewFrameActionBody instantiates a new FrameActionBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFrameActionBodyWithDefaults

`func NewFrameActionBodyWithDefaults() *FrameActionBody`

NewFrameActionBodyWithDefaults instantiates a new FrameActionBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *FrameActionBody) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *FrameActionBody) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *FrameActionBody) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetButtonIndex

`func (o *FrameActionBody) GetButtonIndex() int32`

GetButtonIndex returns the ButtonIndex field if non-nil, zero value otherwise.

### GetButtonIndexOk

`func (o *FrameActionBody) GetButtonIndexOk() (*int32, bool)`

GetButtonIndexOk returns a tuple with the ButtonIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButtonIndex

`func (o *FrameActionBody) SetButtonIndex(v int32)`

SetButtonIndex sets ButtonIndex field to given value.


### GetCastId

`func (o *FrameActionBody) GetCastId() CastId`

GetCastId returns the CastId field if non-nil, zero value otherwise.

### GetCastIdOk

`func (o *FrameActionBody) GetCastIdOk() (*CastId, bool)`

GetCastIdOk returns a tuple with the CastId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCastId

`func (o *FrameActionBody) SetCastId(v CastId)`

SetCastId sets CastId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


