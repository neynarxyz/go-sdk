# TransactionFrameBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier for the transaction mini app | 
**Url** | **string** | URL that can be used to access the transaction mini app | 
**Type** | [**TransactionFrameType**](TransactionFrameType.md) |  | 
**Config** | [**TransactionFrameConfig**](TransactionFrameConfig.md) |  | 
**Status** | [**TransactionFrameStatus**](TransactionFrameStatus.md) |  | 

## Methods

### NewTransactionFrameBase

`func NewTransactionFrameBase(id string, url string, type_ TransactionFrameType, config TransactionFrameConfig, status TransactionFrameStatus, ) *TransactionFrameBase`

NewTransactionFrameBase instantiates a new TransactionFrameBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionFrameBaseWithDefaults

`func NewTransactionFrameBaseWithDefaults() *TransactionFrameBase`

NewTransactionFrameBaseWithDefaults instantiates a new TransactionFrameBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TransactionFrameBase) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TransactionFrameBase) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TransactionFrameBase) SetId(v string)`

SetId sets Id field to given value.


### GetUrl

`func (o *TransactionFrameBase) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *TransactionFrameBase) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *TransactionFrameBase) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetType

`func (o *TransactionFrameBase) GetType() TransactionFrameType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TransactionFrameBase) GetTypeOk() (*TransactionFrameType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TransactionFrameBase) SetType(v TransactionFrameType)`

SetType sets Type field to given value.


### GetConfig

`func (o *TransactionFrameBase) GetConfig() TransactionFrameConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *TransactionFrameBase) GetConfigOk() (*TransactionFrameConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *TransactionFrameBase) SetConfig(v TransactionFrameConfig)`

SetConfig sets Config field to given value.


### GetStatus

`func (o *TransactionFrameBase) GetStatus() TransactionFrameStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TransactionFrameBase) GetStatusOk() (*TransactionFrameStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TransactionFrameBase) SetStatus(v TransactionFrameStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


