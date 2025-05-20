# TransactionSendFungiblesReceipt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fid** | **int32** | The unique identifier of a farcaster user (unsigned integer) | 
**Amount** | **float32** |  | 
**Status** | **string** |  | 
**Reason** | Pointer to **string** | Reason for failure (if status is failed) | [optional] 

## Methods

### NewTransactionSendFungiblesReceipt

`func NewTransactionSendFungiblesReceipt(fid int32, amount float32, status string, ) *TransactionSendFungiblesReceipt`

NewTransactionSendFungiblesReceipt instantiates a new TransactionSendFungiblesReceipt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionSendFungiblesReceiptWithDefaults

`func NewTransactionSendFungiblesReceiptWithDefaults() *TransactionSendFungiblesReceipt`

NewTransactionSendFungiblesReceiptWithDefaults instantiates a new TransactionSendFungiblesReceipt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFid

`func (o *TransactionSendFungiblesReceipt) GetFid() int32`

GetFid returns the Fid field if non-nil, zero value otherwise.

### GetFidOk

`func (o *TransactionSendFungiblesReceipt) GetFidOk() (*int32, bool)`

GetFidOk returns a tuple with the Fid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFid

`func (o *TransactionSendFungiblesReceipt) SetFid(v int32)`

SetFid sets Fid field to given value.


### GetAmount

`func (o *TransactionSendFungiblesReceipt) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TransactionSendFungiblesReceipt) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TransactionSendFungiblesReceipt) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetStatus

`func (o *TransactionSendFungiblesReceipt) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TransactionSendFungiblesReceipt) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TransactionSendFungiblesReceipt) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetReason

`func (o *TransactionSendFungiblesReceipt) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *TransactionSendFungiblesReceipt) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *TransactionSendFungiblesReceipt) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *TransactionSendFungiblesReceipt) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


