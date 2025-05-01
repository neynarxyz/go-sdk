# CastComposerActionsListResponseActionsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the action. | [optional] 
**Icon** | Pointer to **string** | The icon representing the action. | [optional] 
**Description** | Pointer to **string** | A brief description of the action. | [optional] 
**AboutUrl** | Pointer to **string** | URL to learn more about the action. | [optional] 
**ImageUrl** | Pointer to **string** | URL of the action&#39;s image. | [optional] 
**ActionUrl** | Pointer to **string** | URL to perform the action. | [optional] 
**Action** | Pointer to [**CastComposerActionsListResponseActionsInnerAction**](CastComposerActionsListResponseActionsInnerAction.md) |  | [optional] 
**Octicon** | Pointer to **string** | Icon name for the action. | [optional] 
**AddedCount** | Pointer to **int32** | Number of times the action has been added. | [optional] 
**AppName** | Pointer to **string** | Name of the application providing the action. | [optional] 
**AuthorFid** | Pointer to **int32** | Author&#39;s Farcaster ID. | [optional] 
**Category** | Pointer to **string** | Category of the action. | [optional] 
**Object** | Pointer to **string** | Object type, which is \&quot;composer_action\&quot;. | [optional] 

## Methods

### NewCastComposerActionsListResponseActionsInner

`func NewCastComposerActionsListResponseActionsInner() *CastComposerActionsListResponseActionsInner`

NewCastComposerActionsListResponseActionsInner instantiates a new CastComposerActionsListResponseActionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCastComposerActionsListResponseActionsInnerWithDefaults

`func NewCastComposerActionsListResponseActionsInnerWithDefaults() *CastComposerActionsListResponseActionsInner`

NewCastComposerActionsListResponseActionsInnerWithDefaults instantiates a new CastComposerActionsListResponseActionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CastComposerActionsListResponseActionsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CastComposerActionsListResponseActionsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CastComposerActionsListResponseActionsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CastComposerActionsListResponseActionsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIcon

`func (o *CastComposerActionsListResponseActionsInner) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *CastComposerActionsListResponseActionsInner) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *CastComposerActionsListResponseActionsInner) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *CastComposerActionsListResponseActionsInner) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetDescription

`func (o *CastComposerActionsListResponseActionsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CastComposerActionsListResponseActionsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CastComposerActionsListResponseActionsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CastComposerActionsListResponseActionsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAboutUrl

`func (o *CastComposerActionsListResponseActionsInner) GetAboutUrl() string`

GetAboutUrl returns the AboutUrl field if non-nil, zero value otherwise.

### GetAboutUrlOk

`func (o *CastComposerActionsListResponseActionsInner) GetAboutUrlOk() (*string, bool)`

GetAboutUrlOk returns a tuple with the AboutUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAboutUrl

`func (o *CastComposerActionsListResponseActionsInner) SetAboutUrl(v string)`

SetAboutUrl sets AboutUrl field to given value.

### HasAboutUrl

`func (o *CastComposerActionsListResponseActionsInner) HasAboutUrl() bool`

HasAboutUrl returns a boolean if a field has been set.

### GetImageUrl

`func (o *CastComposerActionsListResponseActionsInner) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *CastComposerActionsListResponseActionsInner) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *CastComposerActionsListResponseActionsInner) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *CastComposerActionsListResponseActionsInner) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### GetActionUrl

`func (o *CastComposerActionsListResponseActionsInner) GetActionUrl() string`

GetActionUrl returns the ActionUrl field if non-nil, zero value otherwise.

### GetActionUrlOk

`func (o *CastComposerActionsListResponseActionsInner) GetActionUrlOk() (*string, bool)`

GetActionUrlOk returns a tuple with the ActionUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionUrl

`func (o *CastComposerActionsListResponseActionsInner) SetActionUrl(v string)`

SetActionUrl sets ActionUrl field to given value.

### HasActionUrl

`func (o *CastComposerActionsListResponseActionsInner) HasActionUrl() bool`

HasActionUrl returns a boolean if a field has been set.

### GetAction

`func (o *CastComposerActionsListResponseActionsInner) GetAction() CastComposerActionsListResponseActionsInnerAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *CastComposerActionsListResponseActionsInner) GetActionOk() (*CastComposerActionsListResponseActionsInnerAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *CastComposerActionsListResponseActionsInner) SetAction(v CastComposerActionsListResponseActionsInnerAction)`

SetAction sets Action field to given value.

### HasAction

`func (o *CastComposerActionsListResponseActionsInner) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetOcticon

`func (o *CastComposerActionsListResponseActionsInner) GetOcticon() string`

GetOcticon returns the Octicon field if non-nil, zero value otherwise.

### GetOcticonOk

`func (o *CastComposerActionsListResponseActionsInner) GetOcticonOk() (*string, bool)`

GetOcticonOk returns a tuple with the Octicon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOcticon

`func (o *CastComposerActionsListResponseActionsInner) SetOcticon(v string)`

SetOcticon sets Octicon field to given value.

### HasOcticon

`func (o *CastComposerActionsListResponseActionsInner) HasOcticon() bool`

HasOcticon returns a boolean if a field has been set.

### GetAddedCount

`func (o *CastComposerActionsListResponseActionsInner) GetAddedCount() int32`

GetAddedCount returns the AddedCount field if non-nil, zero value otherwise.

### GetAddedCountOk

`func (o *CastComposerActionsListResponseActionsInner) GetAddedCountOk() (*int32, bool)`

GetAddedCountOk returns a tuple with the AddedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedCount

`func (o *CastComposerActionsListResponseActionsInner) SetAddedCount(v int32)`

SetAddedCount sets AddedCount field to given value.

### HasAddedCount

`func (o *CastComposerActionsListResponseActionsInner) HasAddedCount() bool`

HasAddedCount returns a boolean if a field has been set.

### GetAppName

`func (o *CastComposerActionsListResponseActionsInner) GetAppName() string`

GetAppName returns the AppName field if non-nil, zero value otherwise.

### GetAppNameOk

`func (o *CastComposerActionsListResponseActionsInner) GetAppNameOk() (*string, bool)`

GetAppNameOk returns a tuple with the AppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppName

`func (o *CastComposerActionsListResponseActionsInner) SetAppName(v string)`

SetAppName sets AppName field to given value.

### HasAppName

`func (o *CastComposerActionsListResponseActionsInner) HasAppName() bool`

HasAppName returns a boolean if a field has been set.

### GetAuthorFid

`func (o *CastComposerActionsListResponseActionsInner) GetAuthorFid() int32`

GetAuthorFid returns the AuthorFid field if non-nil, zero value otherwise.

### GetAuthorFidOk

`func (o *CastComposerActionsListResponseActionsInner) GetAuthorFidOk() (*int32, bool)`

GetAuthorFidOk returns a tuple with the AuthorFid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorFid

`func (o *CastComposerActionsListResponseActionsInner) SetAuthorFid(v int32)`

SetAuthorFid sets AuthorFid field to given value.

### HasAuthorFid

`func (o *CastComposerActionsListResponseActionsInner) HasAuthorFid() bool`

HasAuthorFid returns a boolean if a field has been set.

### GetCategory

`func (o *CastComposerActionsListResponseActionsInner) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *CastComposerActionsListResponseActionsInner) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *CastComposerActionsListResponseActionsInner) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *CastComposerActionsListResponseActionsInner) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetObject

`func (o *CastComposerActionsListResponseActionsInner) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *CastComposerActionsListResponseActionsInner) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *CastComposerActionsListResponseActionsInner) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *CastComposerActionsListResponseActionsInner) HasObject() bool`

HasObject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


