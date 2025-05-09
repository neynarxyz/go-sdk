# RemoveVerificationReqBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SignerUuid** | **string** | UUID of the signer. &#x60;signer_uuid&#x60; is paired with API key, can&#39;t use a &#x60;uuid&#x60; made with a different API key.  | 
**Address** | **string** | Ethereum address | 
**BlockHash** | **string** |  | 

## Methods

### NewRemoveVerificationReqBody

`func NewRemoveVerificationReqBody(signerUuid string, address string, blockHash string, ) *RemoveVerificationReqBody`

NewRemoveVerificationReqBody instantiates a new RemoveVerificationReqBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoveVerificationReqBodyWithDefaults

`func NewRemoveVerificationReqBodyWithDefaults() *RemoveVerificationReqBody`

NewRemoveVerificationReqBodyWithDefaults instantiates a new RemoveVerificationReqBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSignerUuid

`func (o *RemoveVerificationReqBody) GetSignerUuid() string`

GetSignerUuid returns the SignerUuid field if non-nil, zero value otherwise.

### GetSignerUuidOk

`func (o *RemoveVerificationReqBody) GetSignerUuidOk() (*string, bool)`

GetSignerUuidOk returns a tuple with the SignerUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignerUuid

`func (o *RemoveVerificationReqBody) SetSignerUuid(v string)`

SetSignerUuid sets SignerUuid field to given value.


### GetAddress

`func (o *RemoveVerificationReqBody) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *RemoveVerificationReqBody) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *RemoveVerificationReqBody) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetBlockHash

`func (o *RemoveVerificationReqBody) GetBlockHash() string`

GetBlockHash returns the BlockHash field if non-nil, zero value otherwise.

### GetBlockHashOk

`func (o *RemoveVerificationReqBody) GetBlockHashOk() (*string, bool)`

GetBlockHashOk returns a tuple with the BlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHash

`func (o *RemoveVerificationReqBody) SetBlockHash(v string)`

SetBlockHash sets BlockHash field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


