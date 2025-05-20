# TransactionSendFungiblesRecipient

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fid** | **int32** | The unique identifier of a farcaster user (unsigned integer) | 
**Amount** | **float32** | Amount to send (must be greater than 0) | 

## Methods

### NewTransactionSendFungiblesRecipient

`func NewTransactionSendFungiblesRecipient(fid int32, amount float32, ) *TransactionSendFungiblesRecipient`

NewTransactionSendFungiblesRecipient instantiates a new TransactionSendFungiblesRecipient object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionSendFungiblesRecipientWithDefaults

`func NewTransactionSendFungiblesRecipientWithDefaults() *TransactionSendFungiblesRecipient`

NewTransactionSendFungiblesRecipientWithDefaults instantiates a new TransactionSendFungiblesRecipient object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFid

`func (o *TransactionSendFungiblesRecipient) GetFid() int32`

GetFid returns the Fid field if non-nil, zero value otherwise.

### GetFidOk

`func (o *TransactionSendFungiblesRecipient) GetFidOk() (*int32, bool)`

GetFidOk returns a tuple with the Fid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFid

`func (o *TransactionSendFungiblesRecipient) SetFid(v int32)`

SetFid sets Fid field to given value.


### GetAmount

`func (o *TransactionSendFungiblesRecipient) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TransactionSendFungiblesRecipient) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TransactionSendFungiblesRecipient) SetAmount(v float32)`

SetAmount sets Amount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


