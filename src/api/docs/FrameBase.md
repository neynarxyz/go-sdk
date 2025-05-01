# FrameBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | **string** | Version of the mini app, &#39;next&#39; for v2, &#39;vNext&#39; for v1 | 
**Image** | **string** | URL of the image | 
**FramesUrl** | **string** | Launch URL of the mini app | 

## Methods

### NewFrameBase

`func NewFrameBase(version string, image string, framesUrl string, ) *FrameBase`

NewFrameBase instantiates a new FrameBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFrameBaseWithDefaults

`func NewFrameBaseWithDefaults() *FrameBase`

NewFrameBaseWithDefaults instantiates a new FrameBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *FrameBase) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *FrameBase) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *FrameBase) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetImage

`func (o *FrameBase) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *FrameBase) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *FrameBase) SetImage(v string)`

SetImage sets Image field to given value.


### GetFramesUrl

`func (o *FrameBase) GetFramesUrl() string`

GetFramesUrl returns the FramesUrl field if non-nil, zero value otherwise.

### GetFramesUrlOk

`func (o *FrameBase) GetFramesUrlOk() (*string, bool)`

GetFramesUrlOk returns a tuple with the FramesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFramesUrl

`func (o *FrameBase) SetFramesUrl(v string)`

SetFramesUrl sets FramesUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


