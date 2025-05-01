# PostCastResponseCast

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hash** | **string** | Cast Hash | [default to "0xfe90f9de682273e05b201629ad2338bdcd89b6be"]
**Author** | [**CastEmbeddedParentAuthor**](CastEmbeddedParentAuthor.md) |  | 
**Text** | **string** |  | 

## Methods

### NewPostCastResponseCast

`func NewPostCastResponseCast(hash string, author CastEmbeddedParentAuthor, text string, ) *PostCastResponseCast`

NewPostCastResponseCast instantiates a new PostCastResponseCast object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostCastResponseCastWithDefaults

`func NewPostCastResponseCastWithDefaults() *PostCastResponseCast`

NewPostCastResponseCastWithDefaults instantiates a new PostCastResponseCast object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHash

`func (o *PostCastResponseCast) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *PostCastResponseCast) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *PostCastResponseCast) SetHash(v string)`

SetHash sets Hash field to given value.


### GetAuthor

`func (o *PostCastResponseCast) GetAuthor() CastEmbeddedParentAuthor`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *PostCastResponseCast) GetAuthorOk() (*CastEmbeddedParentAuthor, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *PostCastResponseCast) SetAuthor(v CastEmbeddedParentAuthor)`

SetAuthor sets Author field to given value.


### GetText

`func (o *PostCastResponseCast) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *PostCastResponseCast) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *PostCastResponseCast) SetText(v string)`

SetText sets Text field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


