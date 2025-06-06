# FrameV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | **string** | Version of the mini app, &#39;next&#39; for v2, &#39;vNext&#39; for v1 | 
**Image** | **string** | URL of the image | 
**FramesUrl** | **string** | Launch URL of the mini app | 
**Buttons** | Pointer to [**[]FrameActionButton**](FrameActionButton.md) |  | [optional] 
**PostUrl** | Pointer to **string** | Post URL to take an action on this mini app | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**ImageAspectRatio** | Pointer to **string** |  | [optional] 
**Input** | Pointer to [**FrameV1AllOfInput**](FrameV1AllOfInput.md) |  | [optional] 
**State** | Pointer to [**FrameV1AllOfState**](FrameV1AllOfState.md) |  | [optional] 

## Methods

### NewFrameV1

`func NewFrameV1(version string, image string, framesUrl string, ) *FrameV1`

NewFrameV1 instantiates a new FrameV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFrameV1WithDefaults

`func NewFrameV1WithDefaults() *FrameV1`

NewFrameV1WithDefaults instantiates a new FrameV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *FrameV1) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *FrameV1) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *FrameV1) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetImage

`func (o *FrameV1) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *FrameV1) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *FrameV1) SetImage(v string)`

SetImage sets Image field to given value.


### GetFramesUrl

`func (o *FrameV1) GetFramesUrl() string`

GetFramesUrl returns the FramesUrl field if non-nil, zero value otherwise.

### GetFramesUrlOk

`func (o *FrameV1) GetFramesUrlOk() (*string, bool)`

GetFramesUrlOk returns a tuple with the FramesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFramesUrl

`func (o *FrameV1) SetFramesUrl(v string)`

SetFramesUrl sets FramesUrl field to given value.


### GetButtons

`func (o *FrameV1) GetButtons() []FrameActionButton`

GetButtons returns the Buttons field if non-nil, zero value otherwise.

### GetButtonsOk

`func (o *FrameV1) GetButtonsOk() (*[]FrameActionButton, bool)`

GetButtonsOk returns a tuple with the Buttons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButtons

`func (o *FrameV1) SetButtons(v []FrameActionButton)`

SetButtons sets Buttons field to given value.

### HasButtons

`func (o *FrameV1) HasButtons() bool`

HasButtons returns a boolean if a field has been set.

### GetPostUrl

`func (o *FrameV1) GetPostUrl() string`

GetPostUrl returns the PostUrl field if non-nil, zero value otherwise.

### GetPostUrlOk

`func (o *FrameV1) GetPostUrlOk() (*string, bool)`

GetPostUrlOk returns a tuple with the PostUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostUrl

`func (o *FrameV1) SetPostUrl(v string)`

SetPostUrl sets PostUrl field to given value.

### HasPostUrl

`func (o *FrameV1) HasPostUrl() bool`

HasPostUrl returns a boolean if a field has been set.

### GetTitle

`func (o *FrameV1) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *FrameV1) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *FrameV1) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *FrameV1) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetImageAspectRatio

`func (o *FrameV1) GetImageAspectRatio() string`

GetImageAspectRatio returns the ImageAspectRatio field if non-nil, zero value otherwise.

### GetImageAspectRatioOk

`func (o *FrameV1) GetImageAspectRatioOk() (*string, bool)`

GetImageAspectRatioOk returns a tuple with the ImageAspectRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageAspectRatio

`func (o *FrameV1) SetImageAspectRatio(v string)`

SetImageAspectRatio sets ImageAspectRatio field to given value.

### HasImageAspectRatio

`func (o *FrameV1) HasImageAspectRatio() bool`

HasImageAspectRatio returns a boolean if a field has been set.

### GetInput

`func (o *FrameV1) GetInput() FrameV1AllOfInput`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *FrameV1) GetInputOk() (*FrameV1AllOfInput, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *FrameV1) SetInput(v FrameV1AllOfInput)`

SetInput sets Input field to given value.

### HasInput

`func (o *FrameV1) HasInput() bool`

HasInput returns a boolean if a field has been set.

### GetState

`func (o *FrameV1) GetState() FrameV1AllOfState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *FrameV1) GetStateOk() (*FrameV1AllOfState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *FrameV1) SetState(v FrameV1AllOfState)`

SetState sets State field to given value.

### HasState

`func (o *FrameV1) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


