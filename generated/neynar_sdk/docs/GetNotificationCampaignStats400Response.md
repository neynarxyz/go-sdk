# GetNotificationCampaignStats400Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** |  | 
**Code** | **string** |  | 
**Errors** | [**[]ZodErrorErrorsInner**](ZodErrorErrorsInner.md) |  | 
**Property** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **int32** |  | [optional] 

## Methods

### NewGetNotificationCampaignStats400Response

`func NewGetNotificationCampaignStats400Response(message string, code string, errors []ZodErrorErrorsInner, ) *GetNotificationCampaignStats400Response`

NewGetNotificationCampaignStats400Response instantiates a new GetNotificationCampaignStats400Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNotificationCampaignStats400ResponseWithDefaults

`func NewGetNotificationCampaignStats400ResponseWithDefaults() *GetNotificationCampaignStats400Response`

NewGetNotificationCampaignStats400ResponseWithDefaults instantiates a new GetNotificationCampaignStats400Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *GetNotificationCampaignStats400Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetNotificationCampaignStats400Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetNotificationCampaignStats400Response) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetCode

`func (o *GetNotificationCampaignStats400Response) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetNotificationCampaignStats400Response) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetNotificationCampaignStats400Response) SetCode(v string)`

SetCode sets Code field to given value.


### GetErrors

`func (o *GetNotificationCampaignStats400Response) GetErrors() []ZodErrorErrorsInner`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *GetNotificationCampaignStats400Response) GetErrorsOk() (*[]ZodErrorErrorsInner, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *GetNotificationCampaignStats400Response) SetErrors(v []ZodErrorErrorsInner)`

SetErrors sets Errors field to given value.


### GetProperty

`func (o *GetNotificationCampaignStats400Response) GetProperty() string`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *GetNotificationCampaignStats400Response) GetPropertyOk() (*string, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *GetNotificationCampaignStats400Response) SetProperty(v string)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *GetNotificationCampaignStats400Response) HasProperty() bool`

HasProperty returns a boolean if a field has been set.

### GetStatus

`func (o *GetNotificationCampaignStats400Response) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetNotificationCampaignStats400Response) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetNotificationCampaignStats400Response) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetNotificationCampaignStats400Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


