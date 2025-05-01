# FetchUserInteractions200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Interactions** | [**[]Notification**](Notification.md) |  | 

## Methods

### NewFetchUserInteractions200Response

`func NewFetchUserInteractions200Response(interactions []Notification, ) *FetchUserInteractions200Response`

NewFetchUserInteractions200Response instantiates a new FetchUserInteractions200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFetchUserInteractions200ResponseWithDefaults

`func NewFetchUserInteractions200ResponseWithDefaults() *FetchUserInteractions200Response`

NewFetchUserInteractions200ResponseWithDefaults instantiates a new FetchUserInteractions200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInteractions

`func (o *FetchUserInteractions200Response) GetInteractions() []Notification`

GetInteractions returns the Interactions field if non-nil, zero value otherwise.

### GetInteractionsOk

`func (o *FetchUserInteractions200Response) GetInteractionsOk() (*[]Notification, bool)`

GetInteractionsOk returns a tuple with the Interactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInteractions

`func (o *FetchUserInteractions200Response) SetInteractions(v []Notification)`

SetInteractions sets Interactions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


