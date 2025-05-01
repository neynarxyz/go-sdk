# FrameV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | **string** | Version of the mini app, &#39;next&#39; for v2, &#39;vNext&#39; for v1 | 
**Image** | **string** | URL of the image | 
**FramesUrl** | **string** | Launch URL of the mini app | 
**Title** | Pointer to **string** | Button title of a mini app | [optional] 
**Manifest** | Pointer to [**FarcasterManifest**](FarcasterManifest.md) |  | [optional] 
**Author** | Pointer to [**UserDehydrated**](UserDehydrated.md) |  | [optional] 
**Metadata** | Pointer to [**FrameV2AllOfMetadata**](FrameV2AllOfMetadata.md) |  | [optional] 

## Methods

### NewFrameV2

`func NewFrameV2(version string, image string, framesUrl string, ) *FrameV2`

NewFrameV2 instantiates a new FrameV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFrameV2WithDefaults

`func NewFrameV2WithDefaults() *FrameV2`

NewFrameV2WithDefaults instantiates a new FrameV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *FrameV2) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *FrameV2) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *FrameV2) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetImage

`func (o *FrameV2) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *FrameV2) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *FrameV2) SetImage(v string)`

SetImage sets Image field to given value.


### GetFramesUrl

`func (o *FrameV2) GetFramesUrl() string`

GetFramesUrl returns the FramesUrl field if non-nil, zero value otherwise.

### GetFramesUrlOk

`func (o *FrameV2) GetFramesUrlOk() (*string, bool)`

GetFramesUrlOk returns a tuple with the FramesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFramesUrl

`func (o *FrameV2) SetFramesUrl(v string)`

SetFramesUrl sets FramesUrl field to given value.


### GetTitle

`func (o *FrameV2) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *FrameV2) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *FrameV2) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *FrameV2) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetManifest

`func (o *FrameV2) GetManifest() FarcasterManifest`

GetManifest returns the Manifest field if non-nil, zero value otherwise.

### GetManifestOk

`func (o *FrameV2) GetManifestOk() (*FarcasterManifest, bool)`

GetManifestOk returns a tuple with the Manifest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManifest

`func (o *FrameV2) SetManifest(v FarcasterManifest)`

SetManifest sets Manifest field to given value.

### HasManifest

`func (o *FrameV2) HasManifest() bool`

HasManifest returns a boolean if a field has been set.

### GetAuthor

`func (o *FrameV2) GetAuthor() UserDehydrated`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *FrameV2) GetAuthorOk() (*UserDehydrated, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *FrameV2) SetAuthor(v UserDehydrated)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *FrameV2) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetMetadata

`func (o *FrameV2) GetMetadata() FrameV2AllOfMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *FrameV2) GetMetadataOk() (*FrameV2AllOfMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *FrameV2) SetMetadata(v FrameV2AllOfMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *FrameV2) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


