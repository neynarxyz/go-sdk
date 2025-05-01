# TransactionFrame

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier for the transaction mini app | 
**Url** | **string** | URL that can be used to access the transaction mini app | 
**Type** | [**TransactionFrameType**](TransactionFrameType.md) |  | 
**Config** | [**TransactionFrameConfig**](TransactionFrameConfig.md) |  | 
**Status** | [**TransactionFrameStatus**](TransactionFrameStatus.md) |  | 
**Transaction** | [**TransactionFramePayAllOfTransaction**](TransactionFramePayAllOfTransaction.md) |  | 

## Methods

### NewTransactionFrame

`func NewTransactionFrame(id string, url string, type_ TransactionFrameType, config TransactionFrameConfig, status TransactionFrameStatus, transaction TransactionFramePayAllOfTransaction, ) *TransactionFrame`

NewTransactionFrame instantiates a new TransactionFrame object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionFrameWithDefaults

`func NewTransactionFrameWithDefaults() *TransactionFrame`

NewTransactionFrameWithDefaults instantiates a new TransactionFrame object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TransactionFrame) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TransactionFrame) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TransactionFrame) SetId(v string)`

SetId sets Id field to given value.


### GetUrl

`func (o *TransactionFrame) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *TransactionFrame) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *TransactionFrame) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetType

`func (o *TransactionFrame) GetType() TransactionFrameType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TransactionFrame) GetTypeOk() (*TransactionFrameType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TransactionFrame) SetType(v TransactionFrameType)`

SetType sets Type field to given value.


### GetConfig

`func (o *TransactionFrame) GetConfig() TransactionFrameConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *TransactionFrame) GetConfigOk() (*TransactionFrameConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *TransactionFrame) SetConfig(v TransactionFrameConfig)`

SetConfig sets Config field to given value.


### GetStatus

`func (o *TransactionFrame) GetStatus() TransactionFrameStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TransactionFrame) GetStatusOk() (*TransactionFrameStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TransactionFrame) SetStatus(v TransactionFrameStatus)`

SetStatus sets Status field to given value.


### GetTransaction

`func (o *TransactionFrame) GetTransaction() TransactionFramePayAllOfTransaction`

GetTransaction returns the Transaction field if non-nil, zero value otherwise.

### GetTransactionOk

`func (o *TransactionFrame) GetTransactionOk() (*TransactionFramePayAllOfTransaction, bool)`

GetTransactionOk returns a tuple with the Transaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransaction

`func (o *TransactionFrame) SetTransaction(v TransactionFramePayAllOfTransaction)`

SetTransaction sets Transaction field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


