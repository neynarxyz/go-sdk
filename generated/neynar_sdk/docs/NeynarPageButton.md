# NeynarPageButton

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** | The title of the button. | 
**Index** | **int32** | The index of the button, first button should have index 1 and so on. | 
**ActionType** | **string** | The type of action that the button performs. | 
**NextPage** | Pointer to [**NeynarPageButtonNextPage**](NeynarPageButtonNextPage.md) |  | [optional] 

## Methods

### NewNeynarPageButton

`func NewNeynarPageButton(title string, index int32, actionType string, ) *NeynarPageButton`

NewNeynarPageButton instantiates a new NeynarPageButton object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNeynarPageButtonWithDefaults

`func NewNeynarPageButtonWithDefaults() *NeynarPageButton`

NewNeynarPageButtonWithDefaults instantiates a new NeynarPageButton object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *NeynarPageButton) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *NeynarPageButton) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *NeynarPageButton) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetIndex

`func (o *NeynarPageButton) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *NeynarPageButton) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *NeynarPageButton) SetIndex(v int32)`

SetIndex sets Index field to given value.


### GetActionType

`func (o *NeynarPageButton) GetActionType() string`

GetActionType returns the ActionType field if non-nil, zero value otherwise.

### GetActionTypeOk

`func (o *NeynarPageButton) GetActionTypeOk() (*string, bool)`

GetActionTypeOk returns a tuple with the ActionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionType

`func (o *NeynarPageButton) SetActionType(v string)`

SetActionType sets ActionType field to given value.


### GetNextPage

`func (o *NeynarPageButton) GetNextPage() NeynarPageButtonNextPage`

GetNextPage returns the NextPage field if non-nil, zero value otherwise.

### GetNextPageOk

`func (o *NeynarPageButton) GetNextPageOk() (*NeynarPageButtonNextPage, bool)`

GetNextPageOk returns a tuple with the NextPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPage

`func (o *NeynarPageButton) SetNextPage(v NeynarPageButtonNextPage)`

SetNextPage sets NextPage field to given value.

### HasNextPage

`func (o *NeynarPageButton) HasNextPage() bool`

HasNextPage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


