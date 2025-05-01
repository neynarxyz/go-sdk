# MarkNotificationsAsSeenReqBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SignerUuid** | **string** | The UUID of the signer. Signer should have atleast one write permission  | 
**Type** | Pointer to [**NotificationType**](NotificationType.md) |  | [optional] 

## Methods

### NewMarkNotificationsAsSeenReqBody

`func NewMarkNotificationsAsSeenReqBody(signerUuid string, ) *MarkNotificationsAsSeenReqBody`

NewMarkNotificationsAsSeenReqBody instantiates a new MarkNotificationsAsSeenReqBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarkNotificationsAsSeenReqBodyWithDefaults

`func NewMarkNotificationsAsSeenReqBodyWithDefaults() *MarkNotificationsAsSeenReqBody`

NewMarkNotificationsAsSeenReqBodyWithDefaults instantiates a new MarkNotificationsAsSeenReqBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSignerUuid

`func (o *MarkNotificationsAsSeenReqBody) GetSignerUuid() string`

GetSignerUuid returns the SignerUuid field if non-nil, zero value otherwise.

### GetSignerUuidOk

`func (o *MarkNotificationsAsSeenReqBody) GetSignerUuidOk() (*string, bool)`

GetSignerUuidOk returns a tuple with the SignerUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignerUuid

`func (o *MarkNotificationsAsSeenReqBody) SetSignerUuid(v string)`

SetSignerUuid sets SignerUuid field to given value.


### GetType

`func (o *MarkNotificationsAsSeenReqBody) GetType() NotificationType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MarkNotificationsAsSeenReqBody) GetTypeOk() (*NotificationType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MarkNotificationsAsSeenReqBody) SetType(v NotificationType)`

SetType sets Type field to given value.

### HasType

`func (o *MarkNotificationsAsSeenReqBody) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


