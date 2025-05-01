# CastEmbeddedChannel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Object** | **string** |  | 
**ImageUrl** | Pointer to **string** |  | [optional] 
**ViewerContext** | Pointer to [**ChannelUserContext**](ChannelUserContext.md) |  | [optional] 

## Methods

### NewCastEmbeddedChannel

`func NewCastEmbeddedChannel(id string, name string, object string, ) *CastEmbeddedChannel`

NewCastEmbeddedChannel instantiates a new CastEmbeddedChannel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCastEmbeddedChannelWithDefaults

`func NewCastEmbeddedChannelWithDefaults() *CastEmbeddedChannel`

NewCastEmbeddedChannelWithDefaults instantiates a new CastEmbeddedChannel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CastEmbeddedChannel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CastEmbeddedChannel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CastEmbeddedChannel) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *CastEmbeddedChannel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CastEmbeddedChannel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CastEmbeddedChannel) SetName(v string)`

SetName sets Name field to given value.


### GetObject

`func (o *CastEmbeddedChannel) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *CastEmbeddedChannel) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *CastEmbeddedChannel) SetObject(v string)`

SetObject sets Object field to given value.


### GetImageUrl

`func (o *CastEmbeddedChannel) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *CastEmbeddedChannel) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *CastEmbeddedChannel) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *CastEmbeddedChannel) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### GetViewerContext

`func (o *CastEmbeddedChannel) GetViewerContext() ChannelUserContext`

GetViewerContext returns the ViewerContext field if non-nil, zero value otherwise.

### GetViewerContextOk

`func (o *CastEmbeddedChannel) GetViewerContextOk() (*ChannelUserContext, bool)`

GetViewerContextOk returns a tuple with the ViewerContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewerContext

`func (o *CastEmbeddedChannel) SetViewerContext(v ChannelUserContext)`

SetViewerContext sets ViewerContext field to given value.

### HasViewerContext

`func (o *CastEmbeddedChannel) HasViewerContext() bool`

HasViewerContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


