# FrameSignaturePacket

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UntrustedData** | [**FrameSignaturePacketUntrustedData**](FrameSignaturePacketUntrustedData.md) |  | 
**TrustedData** | [**FrameSignaturePacketTrustedData**](FrameSignaturePacketTrustedData.md) |  | 

## Methods

### NewFrameSignaturePacket

`func NewFrameSignaturePacket(untrustedData FrameSignaturePacketUntrustedData, trustedData FrameSignaturePacketTrustedData, ) *FrameSignaturePacket`

NewFrameSignaturePacket instantiates a new FrameSignaturePacket object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFrameSignaturePacketWithDefaults

`func NewFrameSignaturePacketWithDefaults() *FrameSignaturePacket`

NewFrameSignaturePacketWithDefaults instantiates a new FrameSignaturePacket object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUntrustedData

`func (o *FrameSignaturePacket) GetUntrustedData() FrameSignaturePacketUntrustedData`

GetUntrustedData returns the UntrustedData field if non-nil, zero value otherwise.

### GetUntrustedDataOk

`func (o *FrameSignaturePacket) GetUntrustedDataOk() (*FrameSignaturePacketUntrustedData, bool)`

GetUntrustedDataOk returns a tuple with the UntrustedData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUntrustedData

`func (o *FrameSignaturePacket) SetUntrustedData(v FrameSignaturePacketUntrustedData)`

SetUntrustedData sets UntrustedData field to given value.


### GetTrustedData

`func (o *FrameSignaturePacket) GetTrustedData() FrameSignaturePacketTrustedData`

GetTrustedData returns the TrustedData field if non-nil, zero value otherwise.

### GetTrustedDataOk

`func (o *FrameSignaturePacket) GetTrustedDataOk() (*FrameSignaturePacketTrustedData, bool)`

GetTrustedDataOk returns a tuple with the TrustedData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustedData

`func (o *FrameSignaturePacket) SetTrustedData(v FrameSignaturePacketTrustedData)`

SetTrustedData sets TrustedData field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


