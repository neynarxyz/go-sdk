# NeynarFramePage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | **string** | Unique identifier for the page. | 
**Version** | Pointer to **string** | The version of the page schema. | [optional] [default to "vNext"]
**Title** | **string** | The title of the page. | 
**Image** | [**NeynarPageImage**](NeynarPageImage.md) |  | 
**Buttons** | Pointer to [**[]NeynarPageButton**](NeynarPageButton.md) |  | [optional] 
**Input** | Pointer to [**NeynarPageInput**](NeynarPageInput.md) |  | [optional] 

## Methods

### NewNeynarFramePage

`func NewNeynarFramePage(uuid string, title string, image NeynarPageImage, ) *NeynarFramePage`

NewNeynarFramePage instantiates a new NeynarFramePage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNeynarFramePageWithDefaults

`func NewNeynarFramePageWithDefaults() *NeynarFramePage`

NewNeynarFramePageWithDefaults instantiates a new NeynarFramePage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *NeynarFramePage) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *NeynarFramePage) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *NeynarFramePage) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetVersion

`func (o *NeynarFramePage) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *NeynarFramePage) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *NeynarFramePage) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *NeynarFramePage) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetTitle

`func (o *NeynarFramePage) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *NeynarFramePage) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *NeynarFramePage) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetImage

`func (o *NeynarFramePage) GetImage() NeynarPageImage`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *NeynarFramePage) GetImageOk() (*NeynarPageImage, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *NeynarFramePage) SetImage(v NeynarPageImage)`

SetImage sets Image field to given value.


### GetButtons

`func (o *NeynarFramePage) GetButtons() []NeynarPageButton`

GetButtons returns the Buttons field if non-nil, zero value otherwise.

### GetButtonsOk

`func (o *NeynarFramePage) GetButtonsOk() (*[]NeynarPageButton, bool)`

GetButtonsOk returns a tuple with the Buttons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButtons

`func (o *NeynarFramePage) SetButtons(v []NeynarPageButton)`

SetButtons sets Buttons field to given value.

### HasButtons

`func (o *NeynarFramePage) HasButtons() bool`

HasButtons returns a boolean if a field has been set.

### GetInput

`func (o *NeynarFramePage) GetInput() NeynarPageInput`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *NeynarFramePage) GetInputOk() (*NeynarPageInput, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *NeynarFramePage) SetInput(v NeynarPageInput)`

SetInput sets Input field to given value.

### HasInput

`func (o *NeynarFramePage) HasInput() bool`

HasInput returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


