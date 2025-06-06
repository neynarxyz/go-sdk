# FrameSignaturePacketUntrustedData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fid** | Pointer to **int32** | The unique identifier of a farcaster user or app (unsigned integer) | [optional] 
**Url** | Pointer to **string** | URL of the mini app | [optional] 
**MessageHash** | Pointer to **string** | Message hash | [optional] 
**Timestamp** | Pointer to **string** | Timestamp | [optional] 
**Network** | Pointer to **int32** | Network | [optional] 
**ButtonIndex** | Pointer to **int32** | Index of the button | [optional] 
**InputText** | Pointer to **string** | Input text | [optional] 
**State** | Pointer to **string** | State | [optional] 
**TransactionId** | Pointer to **string** | Transaction ID | [optional] 
**Address** | Pointer to **string** | Ethereum address | [optional] 
**CastId** | Pointer to [**CastId**](CastId.md) |  | [optional] 

## Methods

### NewFrameSignaturePacketUntrustedData

`func NewFrameSignaturePacketUntrustedData() *FrameSignaturePacketUntrustedData`

NewFrameSignaturePacketUntrustedData instantiates a new FrameSignaturePacketUntrustedData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFrameSignaturePacketUntrustedDataWithDefaults

`func NewFrameSignaturePacketUntrustedDataWithDefaults() *FrameSignaturePacketUntrustedData`

NewFrameSignaturePacketUntrustedDataWithDefaults instantiates a new FrameSignaturePacketUntrustedData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFid

`func (o *FrameSignaturePacketUntrustedData) GetFid() int32`

GetFid returns the Fid field if non-nil, zero value otherwise.

### GetFidOk

`func (o *FrameSignaturePacketUntrustedData) GetFidOk() (*int32, bool)`

GetFidOk returns a tuple with the Fid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFid

`func (o *FrameSignaturePacketUntrustedData) SetFid(v int32)`

SetFid sets Fid field to given value.

### HasFid

`func (o *FrameSignaturePacketUntrustedData) HasFid() bool`

HasFid returns a boolean if a field has been set.

### GetUrl

`func (o *FrameSignaturePacketUntrustedData) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *FrameSignaturePacketUntrustedData) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *FrameSignaturePacketUntrustedData) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *FrameSignaturePacketUntrustedData) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetMessageHash

`func (o *FrameSignaturePacketUntrustedData) GetMessageHash() string`

GetMessageHash returns the MessageHash field if non-nil, zero value otherwise.

### GetMessageHashOk

`func (o *FrameSignaturePacketUntrustedData) GetMessageHashOk() (*string, bool)`

GetMessageHashOk returns a tuple with the MessageHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageHash

`func (o *FrameSignaturePacketUntrustedData) SetMessageHash(v string)`

SetMessageHash sets MessageHash field to given value.

### HasMessageHash

`func (o *FrameSignaturePacketUntrustedData) HasMessageHash() bool`

HasMessageHash returns a boolean if a field has been set.

### GetTimestamp

`func (o *FrameSignaturePacketUntrustedData) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *FrameSignaturePacketUntrustedData) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *FrameSignaturePacketUntrustedData) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *FrameSignaturePacketUntrustedData) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetNetwork

`func (o *FrameSignaturePacketUntrustedData) GetNetwork() int32`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *FrameSignaturePacketUntrustedData) GetNetworkOk() (*int32, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *FrameSignaturePacketUntrustedData) SetNetwork(v int32)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *FrameSignaturePacketUntrustedData) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetButtonIndex

`func (o *FrameSignaturePacketUntrustedData) GetButtonIndex() int32`

GetButtonIndex returns the ButtonIndex field if non-nil, zero value otherwise.

### GetButtonIndexOk

`func (o *FrameSignaturePacketUntrustedData) GetButtonIndexOk() (*int32, bool)`

GetButtonIndexOk returns a tuple with the ButtonIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButtonIndex

`func (o *FrameSignaturePacketUntrustedData) SetButtonIndex(v int32)`

SetButtonIndex sets ButtonIndex field to given value.

### HasButtonIndex

`func (o *FrameSignaturePacketUntrustedData) HasButtonIndex() bool`

HasButtonIndex returns a boolean if a field has been set.

### GetInputText

`func (o *FrameSignaturePacketUntrustedData) GetInputText() string`

GetInputText returns the InputText field if non-nil, zero value otherwise.

### GetInputTextOk

`func (o *FrameSignaturePacketUntrustedData) GetInputTextOk() (*string, bool)`

GetInputTextOk returns a tuple with the InputText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputText

`func (o *FrameSignaturePacketUntrustedData) SetInputText(v string)`

SetInputText sets InputText field to given value.

### HasInputText

`func (o *FrameSignaturePacketUntrustedData) HasInputText() bool`

HasInputText returns a boolean if a field has been set.

### GetState

`func (o *FrameSignaturePacketUntrustedData) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *FrameSignaturePacketUntrustedData) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *FrameSignaturePacketUntrustedData) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *FrameSignaturePacketUntrustedData) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTransactionId

`func (o *FrameSignaturePacketUntrustedData) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *FrameSignaturePacketUntrustedData) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *FrameSignaturePacketUntrustedData) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *FrameSignaturePacketUntrustedData) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.

### GetAddress

`func (o *FrameSignaturePacketUntrustedData) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *FrameSignaturePacketUntrustedData) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *FrameSignaturePacketUntrustedData) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *FrameSignaturePacketUntrustedData) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetCastId

`func (o *FrameSignaturePacketUntrustedData) GetCastId() CastId`

GetCastId returns the CastId field if non-nil, zero value otherwise.

### GetCastIdOk

`func (o *FrameSignaturePacketUntrustedData) GetCastIdOk() (*CastId, bool)`

GetCastIdOk returns a tuple with the CastId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCastId

`func (o *FrameSignaturePacketUntrustedData) SetCastId(v CastId)`

SetCastId sets CastId field to given value.

### HasCastId

`func (o *FrameSignaturePacketUntrustedData) HasCastId() bool`

HasCastId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


