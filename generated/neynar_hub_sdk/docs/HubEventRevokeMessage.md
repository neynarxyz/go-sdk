# HubEventRevokeMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | **int32** |  | 
**RevokeMessageBody** | [**RevokeMessageBody**](RevokeMessageBody.md) |  | 

## Methods

### NewHubEventRevokeMessage

`func NewHubEventRevokeMessage(type_ string, id int32, revokeMessageBody RevokeMessageBody, ) *HubEventRevokeMessage`

NewHubEventRevokeMessage instantiates a new HubEventRevokeMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHubEventRevokeMessageWithDefaults

`func NewHubEventRevokeMessageWithDefaults() *HubEventRevokeMessage`

NewHubEventRevokeMessageWithDefaults instantiates a new HubEventRevokeMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *HubEventRevokeMessage) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HubEventRevokeMessage) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HubEventRevokeMessage) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *HubEventRevokeMessage) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HubEventRevokeMessage) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HubEventRevokeMessage) SetId(v int32)`

SetId sets Id field to given value.


### GetRevokeMessageBody

`func (o *HubEventRevokeMessage) GetRevokeMessageBody() RevokeMessageBody`

GetRevokeMessageBody returns the RevokeMessageBody field if non-nil, zero value otherwise.

### GetRevokeMessageBodyOk

`func (o *HubEventRevokeMessage) GetRevokeMessageBodyOk() (*RevokeMessageBody, bool)`

GetRevokeMessageBodyOk returns a tuple with the RevokeMessageBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevokeMessageBody

`func (o *HubEventRevokeMessage) SetRevokeMessageBody(v RevokeMessageBody)`

SetRevokeMessageBody sets RevokeMessageBody field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


