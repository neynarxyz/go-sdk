# FarcasterManifest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountAssociation** | [**FarcasterManifestAccountAssociation**](FarcasterManifestAccountAssociation.md) |  | 
**Frame** | Pointer to [**FarcasterManifestFrame**](FarcasterManifestFrame.md) |  | [optional] 
**Triggers** | Pointer to [**[]FarcasterManifestTriggersInner**](FarcasterManifestTriggersInner.md) |  | [optional] 

## Methods

### NewFarcasterManifest

`func NewFarcasterManifest(accountAssociation FarcasterManifestAccountAssociation, ) *FarcasterManifest`

NewFarcasterManifest instantiates a new FarcasterManifest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFarcasterManifestWithDefaults

`func NewFarcasterManifestWithDefaults() *FarcasterManifest`

NewFarcasterManifestWithDefaults instantiates a new FarcasterManifest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountAssociation

`func (o *FarcasterManifest) GetAccountAssociation() FarcasterManifestAccountAssociation`

GetAccountAssociation returns the AccountAssociation field if non-nil, zero value otherwise.

### GetAccountAssociationOk

`func (o *FarcasterManifest) GetAccountAssociationOk() (*FarcasterManifestAccountAssociation, bool)`

GetAccountAssociationOk returns a tuple with the AccountAssociation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountAssociation

`func (o *FarcasterManifest) SetAccountAssociation(v FarcasterManifestAccountAssociation)`

SetAccountAssociation sets AccountAssociation field to given value.


### GetFrame

`func (o *FarcasterManifest) GetFrame() FarcasterManifestFrame`

GetFrame returns the Frame field if non-nil, zero value otherwise.

### GetFrameOk

`func (o *FarcasterManifest) GetFrameOk() (*FarcasterManifestFrame, bool)`

GetFrameOk returns a tuple with the Frame field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrame

`func (o *FarcasterManifest) SetFrame(v FarcasterManifestFrame)`

SetFrame sets Frame field to given value.

### HasFrame

`func (o *FarcasterManifest) HasFrame() bool`

HasFrame returns a boolean if a field has been set.

### GetTriggers

`func (o *FarcasterManifest) GetTriggers() []FarcasterManifestTriggersInner`

GetTriggers returns the Triggers field if non-nil, zero value otherwise.

### GetTriggersOk

`func (o *FarcasterManifest) GetTriggersOk() (*[]FarcasterManifestTriggersInner, bool)`

GetTriggersOk returns a tuple with the Triggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggers

`func (o *FarcasterManifest) SetTriggers(v []FarcasterManifestTriggersInner)`

SetTriggers sets Triggers field to given value.

### HasTriggers

`func (o *FarcasterManifest) HasTriggers() bool`

HasTriggers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


