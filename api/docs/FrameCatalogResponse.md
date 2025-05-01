# FrameCatalogResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Frames** | [**[]FrameV2WithFullAuthor**](FrameV2WithFullAuthor.md) |  | 
**Next** | [**NextCursor**](NextCursor.md) |  | 

## Methods

### NewFrameCatalogResponse

`func NewFrameCatalogResponse(frames []FrameV2WithFullAuthor, next NextCursor, ) *FrameCatalogResponse`

NewFrameCatalogResponse instantiates a new FrameCatalogResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFrameCatalogResponseWithDefaults

`func NewFrameCatalogResponseWithDefaults() *FrameCatalogResponse`

NewFrameCatalogResponseWithDefaults instantiates a new FrameCatalogResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrames

`func (o *FrameCatalogResponse) GetFrames() []FrameV2WithFullAuthor`

GetFrames returns the Frames field if non-nil, zero value otherwise.

### GetFramesOk

`func (o *FrameCatalogResponse) GetFramesOk() (*[]FrameV2WithFullAuthor, bool)`

GetFramesOk returns a tuple with the Frames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrames

`func (o *FrameCatalogResponse) SetFrames(v []FrameV2WithFullAuthor)`

SetFrames sets Frames field to given value.


### GetNext

`func (o *FrameCatalogResponse) GetNext() NextCursor`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *FrameCatalogResponse) GetNextOk() (*NextCursor, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *FrameCatalogResponse) SetNext(v NextCursor)`

SetNext sets Next field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


