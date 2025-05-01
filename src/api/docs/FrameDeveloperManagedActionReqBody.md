# FrameDeveloperManagedActionReqBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CastHash** | Pointer to **string** | Cast Hash | [optional] [default to "0xfe90f9de682273e05b201629ad2338bdcd89b6be"]
**Action** | [**FrameAction**](FrameAction.md) |  | 
**SignaturePacket** | [**FrameSignaturePacket**](FrameSignaturePacket.md) |  | 

## Methods

### NewFrameDeveloperManagedActionReqBody

`func NewFrameDeveloperManagedActionReqBody(action FrameAction, signaturePacket FrameSignaturePacket, ) *FrameDeveloperManagedActionReqBody`

NewFrameDeveloperManagedActionReqBody instantiates a new FrameDeveloperManagedActionReqBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFrameDeveloperManagedActionReqBodyWithDefaults

`func NewFrameDeveloperManagedActionReqBodyWithDefaults() *FrameDeveloperManagedActionReqBody`

NewFrameDeveloperManagedActionReqBodyWithDefaults instantiates a new FrameDeveloperManagedActionReqBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCastHash

`func (o *FrameDeveloperManagedActionReqBody) GetCastHash() string`

GetCastHash returns the CastHash field if non-nil, zero value otherwise.

### GetCastHashOk

`func (o *FrameDeveloperManagedActionReqBody) GetCastHashOk() (*string, bool)`

GetCastHashOk returns a tuple with the CastHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCastHash

`func (o *FrameDeveloperManagedActionReqBody) SetCastHash(v string)`

SetCastHash sets CastHash field to given value.

### HasCastHash

`func (o *FrameDeveloperManagedActionReqBody) HasCastHash() bool`

HasCastHash returns a boolean if a field has been set.

### GetAction

`func (o *FrameDeveloperManagedActionReqBody) GetAction() FrameAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *FrameDeveloperManagedActionReqBody) GetActionOk() (*FrameAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *FrameDeveloperManagedActionReqBody) SetAction(v FrameAction)`

SetAction sets Action field to given value.


### GetSignaturePacket

`func (o *FrameDeveloperManagedActionReqBody) GetSignaturePacket() FrameSignaturePacket`

GetSignaturePacket returns the SignaturePacket field if non-nil, zero value otherwise.

### GetSignaturePacketOk

`func (o *FrameDeveloperManagedActionReqBody) GetSignaturePacketOk() (*FrameSignaturePacket, bool)`

GetSignaturePacketOk returns a tuple with the SignaturePacket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignaturePacket

`func (o *FrameDeveloperManagedActionReqBody) SetSignaturePacket(v FrameSignaturePacket)`

SetSignaturePacket sets SignaturePacket field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


