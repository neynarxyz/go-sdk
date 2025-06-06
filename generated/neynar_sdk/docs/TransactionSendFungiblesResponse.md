# TransactionSendFungiblesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SendReceipts** | [**[]TransactionSendFungiblesReceipt**](TransactionSendFungiblesReceipt.md) |  | 
**Transactions** | [**[]TransactionSendTxInfo**](TransactionSendTxInfo.md) |  | 

## Methods

### NewTransactionSendFungiblesResponse

`func NewTransactionSendFungiblesResponse(sendReceipts []TransactionSendFungiblesReceipt, transactions []TransactionSendTxInfo, ) *TransactionSendFungiblesResponse`

NewTransactionSendFungiblesResponse instantiates a new TransactionSendFungiblesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionSendFungiblesResponseWithDefaults

`func NewTransactionSendFungiblesResponseWithDefaults() *TransactionSendFungiblesResponse`

NewTransactionSendFungiblesResponseWithDefaults instantiates a new TransactionSendFungiblesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSendReceipts

`func (o *TransactionSendFungiblesResponse) GetSendReceipts() []TransactionSendFungiblesReceipt`

GetSendReceipts returns the SendReceipts field if non-nil, zero value otherwise.

### GetSendReceiptsOk

`func (o *TransactionSendFungiblesResponse) GetSendReceiptsOk() (*[]TransactionSendFungiblesReceipt, bool)`

GetSendReceiptsOk returns a tuple with the SendReceipts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendReceipts

`func (o *TransactionSendFungiblesResponse) SetSendReceipts(v []TransactionSendFungiblesReceipt)`

SetSendReceipts sets SendReceipts field to given value.


### GetTransactions

`func (o *TransactionSendFungiblesResponse) GetTransactions() []TransactionSendTxInfo`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *TransactionSendFungiblesResponse) GetTransactionsOk() (*[]TransactionSendTxInfo, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *TransactionSendFungiblesResponse) SetTransactions(v []TransactionSendTxInfo)`

SetTransactions sets Transactions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


