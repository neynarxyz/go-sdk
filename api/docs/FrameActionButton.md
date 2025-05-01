# FrameActionButton

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** | Title of the button | [optional] 
**Index** | **int32** | Index of the button | 
**ActionType** | [**FrameButtonActionType**](FrameButtonActionType.md) |  | 
**Target** | Pointer to **string** | Target of the button | [optional] 
**PostUrl** | Pointer to **string** | Used specifically for the tx action type to post a successful transaction hash | [optional] 

## Methods

### NewFrameActionButton

`func NewFrameActionButton(index int32, actionType FrameButtonActionType, ) *FrameActionButton`

NewFrameActionButton instantiates a new FrameActionButton object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFrameActionButtonWithDefaults

`func NewFrameActionButtonWithDefaults() *FrameActionButton`

NewFrameActionButtonWithDefaults instantiates a new FrameActionButton object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *FrameActionButton) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *FrameActionButton) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *FrameActionButton) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *FrameActionButton) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetIndex

`func (o *FrameActionButton) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *FrameActionButton) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *FrameActionButton) SetIndex(v int32)`

SetIndex sets Index field to given value.


### GetActionType

`func (o *FrameActionButton) GetActionType() FrameButtonActionType`

GetActionType returns the ActionType field if non-nil, zero value otherwise.

### GetActionTypeOk

`func (o *FrameActionButton) GetActionTypeOk() (*FrameButtonActionType, bool)`

GetActionTypeOk returns a tuple with the ActionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionType

`func (o *FrameActionButton) SetActionType(v FrameButtonActionType)`

SetActionType sets ActionType field to given value.


### GetTarget

`func (o *FrameActionButton) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *FrameActionButton) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *FrameActionButton) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *FrameActionButton) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetPostUrl

`func (o *FrameActionButton) GetPostUrl() string`

GetPostUrl returns the PostUrl field if non-nil, zero value otherwise.

### GetPostUrlOk

`func (o *FrameActionButton) GetPostUrlOk() (*string, bool)`

GetPostUrlOk returns a tuple with the PostUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostUrl

`func (o *FrameActionButton) SetPostUrl(v string)`

SetPostUrl sets PostUrl field to given value.

### HasPostUrl

`func (o *FrameActionButton) HasPostUrl() bool`

HasPostUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


