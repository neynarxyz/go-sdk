# TransactionFrameLineItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the line item in transaction | 
**Description** | **string** | Description of the line item in transaction | 
**Image** | Pointer to **string** | Optional image URL for the line item in transaction | [optional] 

## Methods

### NewTransactionFrameLineItem

`func NewTransactionFrameLineItem(name string, description string, ) *TransactionFrameLineItem`

NewTransactionFrameLineItem instantiates a new TransactionFrameLineItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionFrameLineItemWithDefaults

`func NewTransactionFrameLineItemWithDefaults() *TransactionFrameLineItem`

NewTransactionFrameLineItemWithDefaults instantiates a new TransactionFrameLineItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TransactionFrameLineItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TransactionFrameLineItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TransactionFrameLineItem) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *TransactionFrameLineItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TransactionFrameLineItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TransactionFrameLineItem) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetImage

`func (o *TransactionFrameLineItem) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *TransactionFrameLineItem) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *TransactionFrameLineItem) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *TransactionFrameLineItem) HasImage() bool`

HasImage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


