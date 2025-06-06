# FollowReqBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SignerUuid** | **string** | UUID of the signer. &#x60;signer_uuid&#x60; is paired with API key, can&#39;t use a &#x60;uuid&#x60; made with a different API key. | 
**TargetFids** | **[]int32** |  | 

## Methods

### NewFollowReqBody

`func NewFollowReqBody(signerUuid string, targetFids []int32, ) *FollowReqBody`

NewFollowReqBody instantiates a new FollowReqBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFollowReqBodyWithDefaults

`func NewFollowReqBodyWithDefaults() *FollowReqBody`

NewFollowReqBodyWithDefaults instantiates a new FollowReqBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSignerUuid

`func (o *FollowReqBody) GetSignerUuid() string`

GetSignerUuid returns the SignerUuid field if non-nil, zero value otherwise.

### GetSignerUuidOk

`func (o *FollowReqBody) GetSignerUuidOk() (*string, bool)`

GetSignerUuidOk returns a tuple with the SignerUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignerUuid

`func (o *FollowReqBody) SetSignerUuid(v string)`

SetSignerUuid sets SignerUuid field to given value.


### GetTargetFids

`func (o *FollowReqBody) GetTargetFids() []int32`

GetTargetFids returns the TargetFids field if non-nil, zero value otherwise.

### GetTargetFidsOk

`func (o *FollowReqBody) GetTargetFidsOk() (*[]int32, bool)`

GetTargetFidsOk returns a tuple with the TargetFids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetFids

`func (o *FollowReqBody) SetTargetFids(v []int32)`

SetTargetFids sets TargetFids field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


