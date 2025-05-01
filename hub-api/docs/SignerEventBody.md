# SignerEventBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** |  | 
**KeyType** | **int64** |  | 
**EventType** | [**SignerEventType**](SignerEventType.md) |  | [default to SIGNER_EVENT_TYPE_ADD]
**Metadata** | **string** |  | 
**MetadataType** | **int64** |  | 

## Methods

### NewSignerEventBody

`func NewSignerEventBody(key string, keyType int64, eventType SignerEventType, metadata string, metadataType int64, ) *SignerEventBody`

NewSignerEventBody instantiates a new SignerEventBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignerEventBodyWithDefaults

`func NewSignerEventBodyWithDefaults() *SignerEventBody`

NewSignerEventBodyWithDefaults instantiates a new SignerEventBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *SignerEventBody) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *SignerEventBody) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *SignerEventBody) SetKey(v string)`

SetKey sets Key field to given value.


### GetKeyType

`func (o *SignerEventBody) GetKeyType() int64`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *SignerEventBody) GetKeyTypeOk() (*int64, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *SignerEventBody) SetKeyType(v int64)`

SetKeyType sets KeyType field to given value.


### GetEventType

`func (o *SignerEventBody) GetEventType() SignerEventType`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *SignerEventBody) GetEventTypeOk() (*SignerEventType, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *SignerEventBody) SetEventType(v SignerEventType)`

SetEventType sets EventType field to given value.


### GetMetadata

`func (o *SignerEventBody) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SignerEventBody) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SignerEventBody) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.


### GetMetadataType

`func (o *SignerEventBody) GetMetadataType() int64`

GetMetadataType returns the MetadataType field if non-nil, zero value otherwise.

### GetMetadataTypeOk

`func (o *SignerEventBody) GetMetadataTypeOk() (*int64, bool)`

GetMetadataTypeOk returns a tuple with the MetadataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataType

`func (o *SignerEventBody) SetMetadataType(v int64)`

SetMetadataType sets MetadataType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


