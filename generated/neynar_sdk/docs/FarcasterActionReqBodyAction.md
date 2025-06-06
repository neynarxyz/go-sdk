# FarcasterActionReqBodyAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of action being performed.  | 
**Payload** | Pointer to **map[string]interface{}** | The payload of the action being performed.  | [optional] 

## Methods

### NewFarcasterActionReqBodyAction

`func NewFarcasterActionReqBodyAction(type_ string, ) *FarcasterActionReqBodyAction`

NewFarcasterActionReqBodyAction instantiates a new FarcasterActionReqBodyAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFarcasterActionReqBodyActionWithDefaults

`func NewFarcasterActionReqBodyActionWithDefaults() *FarcasterActionReqBodyAction`

NewFarcasterActionReqBodyActionWithDefaults instantiates a new FarcasterActionReqBodyAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FarcasterActionReqBodyAction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FarcasterActionReqBodyAction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FarcasterActionReqBodyAction) SetType(v string)`

SetType sets Type field to given value.


### GetPayload

`func (o *FarcasterActionReqBodyAction) GetPayload() map[string]interface{}`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *FarcasterActionReqBodyAction) GetPayloadOk() (*map[string]interface{}, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *FarcasterActionReqBodyAction) SetPayload(v map[string]interface{})`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *FarcasterActionReqBodyAction) HasPayload() bool`

HasPayload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


