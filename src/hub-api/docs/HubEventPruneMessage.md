# HubEventPruneMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | **int32** |  | 
**PruneMessageBody** | [**PruneMessageBody**](PruneMessageBody.md) |  | 

## Methods

### NewHubEventPruneMessage

`func NewHubEventPruneMessage(type_ string, id int32, pruneMessageBody PruneMessageBody, ) *HubEventPruneMessage`

NewHubEventPruneMessage instantiates a new HubEventPruneMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHubEventPruneMessageWithDefaults

`func NewHubEventPruneMessageWithDefaults() *HubEventPruneMessage`

NewHubEventPruneMessageWithDefaults instantiates a new HubEventPruneMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *HubEventPruneMessage) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HubEventPruneMessage) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HubEventPruneMessage) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *HubEventPruneMessage) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HubEventPruneMessage) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HubEventPruneMessage) SetId(v int32)`

SetId sets Id field to given value.


### GetPruneMessageBody

`func (o *HubEventPruneMessage) GetPruneMessageBody() PruneMessageBody`

GetPruneMessageBody returns the PruneMessageBody field if non-nil, zero value otherwise.

### GetPruneMessageBodyOk

`func (o *HubEventPruneMessage) GetPruneMessageBodyOk() (*PruneMessageBody, bool)`

GetPruneMessageBodyOk returns a tuple with the PruneMessageBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPruneMessageBody

`func (o *HubEventPruneMessage) SetPruneMessageBody(v PruneMessageBody)`

SetPruneMessageBody sets PruneMessageBody field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


