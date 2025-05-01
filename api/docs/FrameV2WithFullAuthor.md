# FrameV2WithFullAuthor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | **string** | Version of the mini app, &#39;next&#39; for v2, &#39;vNext&#39; for v1 | 
**Image** | **string** | URL of the image | 
**FramesUrl** | **string** | Launch URL of the mini app | 
**Title** | Pointer to **string** | Button title of a mini app | [optional] 
**Manifest** | Pointer to [**FarcasterManifest**](FarcasterManifest.md) |  | [optional] 
**Author** | Pointer to [**User**](User.md) |  | [optional] 
**Metadata** | Pointer to [**FrameV2AllOfMetadata**](FrameV2AllOfMetadata.md) |  | [optional] 

## Methods

### NewFrameV2WithFullAuthor

`func NewFrameV2WithFullAuthor(version string, image string, framesUrl string, ) *FrameV2WithFullAuthor`

NewFrameV2WithFullAuthor instantiates a new FrameV2WithFullAuthor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFrameV2WithFullAuthorWithDefaults

`func NewFrameV2WithFullAuthorWithDefaults() *FrameV2WithFullAuthor`

NewFrameV2WithFullAuthorWithDefaults instantiates a new FrameV2WithFullAuthor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *FrameV2WithFullAuthor) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *FrameV2WithFullAuthor) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *FrameV2WithFullAuthor) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetImage

`func (o *FrameV2WithFullAuthor) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *FrameV2WithFullAuthor) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *FrameV2WithFullAuthor) SetImage(v string)`

SetImage sets Image field to given value.


### GetFramesUrl

`func (o *FrameV2WithFullAuthor) GetFramesUrl() string`

GetFramesUrl returns the FramesUrl field if non-nil, zero value otherwise.

### GetFramesUrlOk

`func (o *FrameV2WithFullAuthor) GetFramesUrlOk() (*string, bool)`

GetFramesUrlOk returns a tuple with the FramesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFramesUrl

`func (o *FrameV2WithFullAuthor) SetFramesUrl(v string)`

SetFramesUrl sets FramesUrl field to given value.


### GetTitle

`func (o *FrameV2WithFullAuthor) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *FrameV2WithFullAuthor) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *FrameV2WithFullAuthor) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *FrameV2WithFullAuthor) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetManifest

`func (o *FrameV2WithFullAuthor) GetManifest() FarcasterManifest`

GetManifest returns the Manifest field if non-nil, zero value otherwise.

### GetManifestOk

`func (o *FrameV2WithFullAuthor) GetManifestOk() (*FarcasterManifest, bool)`

GetManifestOk returns a tuple with the Manifest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManifest

`func (o *FrameV2WithFullAuthor) SetManifest(v FarcasterManifest)`

SetManifest sets Manifest field to given value.

### HasManifest

`func (o *FrameV2WithFullAuthor) HasManifest() bool`

HasManifest returns a boolean if a field has been set.

### GetAuthor

`func (o *FrameV2WithFullAuthor) GetAuthor() User`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *FrameV2WithFullAuthor) GetAuthorOk() (*User, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *FrameV2WithFullAuthor) SetAuthor(v User)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *FrameV2WithFullAuthor) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetMetadata

`func (o *FrameV2WithFullAuthor) GetMetadata() FrameV2AllOfMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *FrameV2WithFullAuthor) GetMetadataOk() (*FrameV2AllOfMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *FrameV2WithFullAuthor) SetMetadata(v FrameV2AllOfMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *FrameV2WithFullAuthor) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


