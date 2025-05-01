# UserDataBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**UserDataType**](UserDataType.md) | Specifies which profile field is being updated (e.g., profile picture, display name, bio). | [default to PFP]
**Value** | **string** | The new value for the specified profile field. The format depends on the type: URLs for profile pictures, plain text for display names and bios, etc. | 

## Methods

### NewUserDataBody

`func NewUserDataBody(type_ UserDataType, value string, ) *UserDataBody`

NewUserDataBody instantiates a new UserDataBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserDataBodyWithDefaults

`func NewUserDataBodyWithDefaults() *UserDataBody`

NewUserDataBodyWithDefaults instantiates a new UserDataBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *UserDataBody) GetType() UserDataType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserDataBody) GetTypeOk() (*UserDataType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserDataBody) SetType(v UserDataType)`

SetType sets Type field to given value.


### GetValue

`func (o *UserDataBody) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *UserDataBody) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *UserDataBody) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


