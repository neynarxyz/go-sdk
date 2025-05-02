# TransactionFrameConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowlistFids** | Pointer to **[]int32** | Optional list of FIDs that are allowed to use this transaction mini app | [optional] 
**LineItems** | [**[]TransactionFrameLineItem**](TransactionFrameLineItem.md) | List of items included in the transaction | 
**Action** | Pointer to [**TransactionFrameAction**](TransactionFrameAction.md) |  | [optional] 

## Methods

### NewTransactionFrameConfig

`func NewTransactionFrameConfig(lineItems []TransactionFrameLineItem, ) *TransactionFrameConfig`

NewTransactionFrameConfig instantiates a new TransactionFrameConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionFrameConfigWithDefaults

`func NewTransactionFrameConfigWithDefaults() *TransactionFrameConfig`

NewTransactionFrameConfigWithDefaults instantiates a new TransactionFrameConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowlistFids

`func (o *TransactionFrameConfig) GetAllowlistFids() []int32`

GetAllowlistFids returns the AllowlistFids field if non-nil, zero value otherwise.

### GetAllowlistFidsOk

`func (o *TransactionFrameConfig) GetAllowlistFidsOk() (*[]int32, bool)`

GetAllowlistFidsOk returns a tuple with the AllowlistFids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowlistFids

`func (o *TransactionFrameConfig) SetAllowlistFids(v []int32)`

SetAllowlistFids sets AllowlistFids field to given value.

### HasAllowlistFids

`func (o *TransactionFrameConfig) HasAllowlistFids() bool`

HasAllowlistFids returns a boolean if a field has been set.

### GetLineItems

`func (o *TransactionFrameConfig) GetLineItems() []TransactionFrameLineItem`

GetLineItems returns the LineItems field if non-nil, zero value otherwise.

### GetLineItemsOk

`func (o *TransactionFrameConfig) GetLineItemsOk() (*[]TransactionFrameLineItem, bool)`

GetLineItemsOk returns a tuple with the LineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItems

`func (o *TransactionFrameConfig) SetLineItems(v []TransactionFrameLineItem)`

SetLineItems sets LineItems field to given value.


### GetAction

`func (o *TransactionFrameConfig) GetAction() TransactionFrameAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *TransactionFrameConfig) GetActionOk() (*TransactionFrameAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *TransactionFrameConfig) SetAction(v TransactionFrameAction)`

SetAction sets Action field to given value.

### HasAction

`func (o *TransactionFrameConfig) HasAction() bool`

HasAction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


