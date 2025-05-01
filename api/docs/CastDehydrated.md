# CastDehydrated

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | **string** |  | 
**Hash** | **string** |  | 
**Author** | Pointer to [**UserDehydrated**](UserDehydrated.md) |  | [optional] 
**App** | Pointer to [**CastEmbeddedApp**](CastEmbeddedApp.md) |  | [optional] 

## Methods

### NewCastDehydrated

`func NewCastDehydrated(object string, hash string, ) *CastDehydrated`

NewCastDehydrated instantiates a new CastDehydrated object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCastDehydratedWithDefaults

`func NewCastDehydratedWithDefaults() *CastDehydrated`

NewCastDehydratedWithDefaults instantiates a new CastDehydrated object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *CastDehydrated) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *CastDehydrated) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *CastDehydrated) SetObject(v string)`

SetObject sets Object field to given value.


### GetHash

`func (o *CastDehydrated) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *CastDehydrated) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *CastDehydrated) SetHash(v string)`

SetHash sets Hash field to given value.


### GetAuthor

`func (o *CastDehydrated) GetAuthor() UserDehydrated`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *CastDehydrated) GetAuthorOk() (*UserDehydrated, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *CastDehydrated) SetAuthor(v UserDehydrated)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *CastDehydrated) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetApp

`func (o *CastDehydrated) GetApp() CastEmbeddedApp`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *CastDehydrated) GetAppOk() (*CastEmbeddedApp, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *CastDehydrated) SetApp(v CastEmbeddedApp)`

SetApp sets App field to given value.

### HasApp

`func (o *CastDehydrated) HasApp() bool`

HasApp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


