# NeynarFrameCreationReqBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the mini app. | 
**Pages** | [**[]NeynarFramePage**](NeynarFramePage.md) |  | 

## Methods

### NewNeynarFrameCreationReqBody

`func NewNeynarFrameCreationReqBody(name string, pages []NeynarFramePage, ) *NeynarFrameCreationReqBody`

NewNeynarFrameCreationReqBody instantiates a new NeynarFrameCreationReqBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNeynarFrameCreationReqBodyWithDefaults

`func NewNeynarFrameCreationReqBodyWithDefaults() *NeynarFrameCreationReqBody`

NewNeynarFrameCreationReqBodyWithDefaults instantiates a new NeynarFrameCreationReqBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NeynarFrameCreationReqBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NeynarFrameCreationReqBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NeynarFrameCreationReqBody) SetName(v string)`

SetName sets Name field to given value.


### GetPages

`func (o *NeynarFrameCreationReqBody) GetPages() []NeynarFramePage`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *NeynarFrameCreationReqBody) GetPagesOk() (*[]NeynarFramePage, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *NeynarFrameCreationReqBody) SetPages(v []NeynarFramePage)`

SetPages sets Pages field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


