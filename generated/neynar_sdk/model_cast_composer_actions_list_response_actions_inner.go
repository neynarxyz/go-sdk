/*
Farcaster API V2

The Farcaster API allows you to interact with the Farcaster protocol. See the [Neynar docs](https://docs.neynar.com/reference) for more details.

API version: 2.43.0
Contact: team@neynar.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package neynar_sdk

import (
	"encoding/json"
)

// checks if the CastComposerActionsListResponseActionsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CastComposerActionsListResponseActionsInner{}

// CastComposerActionsListResponseActionsInner struct for CastComposerActionsListResponseActionsInner
type CastComposerActionsListResponseActionsInner struct {
	// The name of the action.
	Name *string `json:"name,omitempty"`
	// The icon representing the action.
	Icon *string `json:"icon,omitempty"`
	// A brief description of the action.
	Description *string `json:"description,omitempty"`
	// URL to learn more about the action.
	AboutUrl *string `json:"about_url,omitempty"`
	// URL of the action's image.
	ImageUrl *string `json:"image_url,omitempty"`
	// URL to perform the action.
	ActionUrl *string                                            `json:"action_url,omitempty"`
	Action    *CastComposerActionsListResponseActionsInnerAction `json:"action,omitempty"`
	// Icon name for the action.
	Octicon *string `json:"octicon,omitempty"`
	// Number of times the action has been added.
	AddedCount *int32 `json:"added_count,omitempty"`
	// Name of the application providing the action.
	AppName *string `json:"app_name,omitempty"`
	// Author's Farcaster ID.
	AuthorFid *int32 `json:"author_fid,omitempty"`
	// Category of the action.
	Category *string `json:"category,omitempty"`
	// Object type, which is \"composer_action\".
	Object *string `json:"object,omitempty"`
}

// NewCastComposerActionsListResponseActionsInner instantiates a new CastComposerActionsListResponseActionsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCastComposerActionsListResponseActionsInner() *CastComposerActionsListResponseActionsInner {
	this := CastComposerActionsListResponseActionsInner{}
	return &this
}

// NewCastComposerActionsListResponseActionsInnerWithDefaults instantiates a new CastComposerActionsListResponseActionsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCastComposerActionsListResponseActionsInnerWithDefaults() *CastComposerActionsListResponseActionsInner {
	this := CastComposerActionsListResponseActionsInner{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CastComposerActionsListResponseActionsInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CastComposerActionsListResponseActionsInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CastComposerActionsListResponseActionsInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CastComposerActionsListResponseActionsInner) SetName(v string) {
	o.Name = &v
}

// GetIcon returns the Icon field value if set, zero value otherwise.
func (o *CastComposerActionsListResponseActionsInner) GetIcon() string {
	if o == nil || IsNil(o.Icon) {
		var ret string
		return ret
	}
	return *o.Icon
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CastComposerActionsListResponseActionsInner) GetIconOk() (*string, bool) {
	if o == nil || IsNil(o.Icon) {
		return nil, false
	}
	return o.Icon, true
}

// HasIcon returns a boolean if a field has been set.
func (o *CastComposerActionsListResponseActionsInner) HasIcon() bool {
	if o != nil && !IsNil(o.Icon) {
		return true
	}

	return false
}

// SetIcon gets a reference to the given string and assigns it to the Icon field.
func (o *CastComposerActionsListResponseActionsInner) SetIcon(v string) {
	o.Icon = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CastComposerActionsListResponseActionsInner) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CastComposerActionsListResponseActionsInner) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CastComposerActionsListResponseActionsInner) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CastComposerActionsListResponseActionsInner) SetDescription(v string) {
	o.Description = &v
}

// GetAboutUrl returns the AboutUrl field value if set, zero value otherwise.
func (o *CastComposerActionsListResponseActionsInner) GetAboutUrl() string {
	if o == nil || IsNil(o.AboutUrl) {
		var ret string
		return ret
	}
	return *o.AboutUrl
}

// GetAboutUrlOk returns a tuple with the AboutUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CastComposerActionsListResponseActionsInner) GetAboutUrlOk() (*string, bool) {
	if o == nil || IsNil(o.AboutUrl) {
		return nil, false
	}
	return o.AboutUrl, true
}

// HasAboutUrl returns a boolean if a field has been set.
func (o *CastComposerActionsListResponseActionsInner) HasAboutUrl() bool {
	if o != nil && !IsNil(o.AboutUrl) {
		return true
	}

	return false
}

// SetAboutUrl gets a reference to the given string and assigns it to the AboutUrl field.
func (o *CastComposerActionsListResponseActionsInner) SetAboutUrl(v string) {
	o.AboutUrl = &v
}

// GetImageUrl returns the ImageUrl field value if set, zero value otherwise.
func (o *CastComposerActionsListResponseActionsInner) GetImageUrl() string {
	if o == nil || IsNil(o.ImageUrl) {
		var ret string
		return ret
	}
	return *o.ImageUrl
}

// GetImageUrlOk returns a tuple with the ImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CastComposerActionsListResponseActionsInner) GetImageUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ImageUrl) {
		return nil, false
	}
	return o.ImageUrl, true
}

// HasImageUrl returns a boolean if a field has been set.
func (o *CastComposerActionsListResponseActionsInner) HasImageUrl() bool {
	if o != nil && !IsNil(o.ImageUrl) {
		return true
	}

	return false
}

// SetImageUrl gets a reference to the given string and assigns it to the ImageUrl field.
func (o *CastComposerActionsListResponseActionsInner) SetImageUrl(v string) {
	o.ImageUrl = &v
}

// GetActionUrl returns the ActionUrl field value if set, zero value otherwise.
func (o *CastComposerActionsListResponseActionsInner) GetActionUrl() string {
	if o == nil || IsNil(o.ActionUrl) {
		var ret string
		return ret
	}
	return *o.ActionUrl
}

// GetActionUrlOk returns a tuple with the ActionUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CastComposerActionsListResponseActionsInner) GetActionUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ActionUrl) {
		return nil, false
	}
	return o.ActionUrl, true
}

// HasActionUrl returns a boolean if a field has been set.
func (o *CastComposerActionsListResponseActionsInner) HasActionUrl() bool {
	if o != nil && !IsNil(o.ActionUrl) {
		return true
	}

	return false
}

// SetActionUrl gets a reference to the given string and assigns it to the ActionUrl field.
func (o *CastComposerActionsListResponseActionsInner) SetActionUrl(v string) {
	o.ActionUrl = &v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *CastComposerActionsListResponseActionsInner) GetAction() CastComposerActionsListResponseActionsInnerAction {
	if o == nil || IsNil(o.Action) {
		var ret CastComposerActionsListResponseActionsInnerAction
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CastComposerActionsListResponseActionsInner) GetActionOk() (*CastComposerActionsListResponseActionsInnerAction, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *CastComposerActionsListResponseActionsInner) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given CastComposerActionsListResponseActionsInnerAction and assigns it to the Action field.
func (o *CastComposerActionsListResponseActionsInner) SetAction(v CastComposerActionsListResponseActionsInnerAction) {
	o.Action = &v
}

// GetOcticon returns the Octicon field value if set, zero value otherwise.
func (o *CastComposerActionsListResponseActionsInner) GetOcticon() string {
	if o == nil || IsNil(o.Octicon) {
		var ret string
		return ret
	}
	return *o.Octicon
}

// GetOcticonOk returns a tuple with the Octicon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CastComposerActionsListResponseActionsInner) GetOcticonOk() (*string, bool) {
	if o == nil || IsNil(o.Octicon) {
		return nil, false
	}
	return o.Octicon, true
}

// HasOcticon returns a boolean if a field has been set.
func (o *CastComposerActionsListResponseActionsInner) HasOcticon() bool {
	if o != nil && !IsNil(o.Octicon) {
		return true
	}

	return false
}

// SetOcticon gets a reference to the given string and assigns it to the Octicon field.
func (o *CastComposerActionsListResponseActionsInner) SetOcticon(v string) {
	o.Octicon = &v
}

// GetAddedCount returns the AddedCount field value if set, zero value otherwise.
func (o *CastComposerActionsListResponseActionsInner) GetAddedCount() int32 {
	if o == nil || IsNil(o.AddedCount) {
		var ret int32
		return ret
	}
	return *o.AddedCount
}

// GetAddedCountOk returns a tuple with the AddedCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CastComposerActionsListResponseActionsInner) GetAddedCountOk() (*int32, bool) {
	if o == nil || IsNil(o.AddedCount) {
		return nil, false
	}
	return o.AddedCount, true
}

// HasAddedCount returns a boolean if a field has been set.
func (o *CastComposerActionsListResponseActionsInner) HasAddedCount() bool {
	if o != nil && !IsNil(o.AddedCount) {
		return true
	}

	return false
}

// SetAddedCount gets a reference to the given int32 and assigns it to the AddedCount field.
func (o *CastComposerActionsListResponseActionsInner) SetAddedCount(v int32) {
	o.AddedCount = &v
}

// GetAppName returns the AppName field value if set, zero value otherwise.
func (o *CastComposerActionsListResponseActionsInner) GetAppName() string {
	if o == nil || IsNil(o.AppName) {
		var ret string
		return ret
	}
	return *o.AppName
}

// GetAppNameOk returns a tuple with the AppName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CastComposerActionsListResponseActionsInner) GetAppNameOk() (*string, bool) {
	if o == nil || IsNil(o.AppName) {
		return nil, false
	}
	return o.AppName, true
}

// HasAppName returns a boolean if a field has been set.
func (o *CastComposerActionsListResponseActionsInner) HasAppName() bool {
	if o != nil && !IsNil(o.AppName) {
		return true
	}

	return false
}

// SetAppName gets a reference to the given string and assigns it to the AppName field.
func (o *CastComposerActionsListResponseActionsInner) SetAppName(v string) {
	o.AppName = &v
}

// GetAuthorFid returns the AuthorFid field value if set, zero value otherwise.
func (o *CastComposerActionsListResponseActionsInner) GetAuthorFid() int32 {
	if o == nil || IsNil(o.AuthorFid) {
		var ret int32
		return ret
	}
	return *o.AuthorFid
}

// GetAuthorFidOk returns a tuple with the AuthorFid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CastComposerActionsListResponseActionsInner) GetAuthorFidOk() (*int32, bool) {
	if o == nil || IsNil(o.AuthorFid) {
		return nil, false
	}
	return o.AuthorFid, true
}

// HasAuthorFid returns a boolean if a field has been set.
func (o *CastComposerActionsListResponseActionsInner) HasAuthorFid() bool {
	if o != nil && !IsNil(o.AuthorFid) {
		return true
	}

	return false
}

// SetAuthorFid gets a reference to the given int32 and assigns it to the AuthorFid field.
func (o *CastComposerActionsListResponseActionsInner) SetAuthorFid(v int32) {
	o.AuthorFid = &v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *CastComposerActionsListResponseActionsInner) GetCategory() string {
	if o == nil || IsNil(o.Category) {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CastComposerActionsListResponseActionsInner) GetCategoryOk() (*string, bool) {
	if o == nil || IsNil(o.Category) {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *CastComposerActionsListResponseActionsInner) HasCategory() bool {
	if o != nil && !IsNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *CastComposerActionsListResponseActionsInner) SetCategory(v string) {
	o.Category = &v
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *CastComposerActionsListResponseActionsInner) GetObject() string {
	if o == nil || IsNil(o.Object) {
		var ret string
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CastComposerActionsListResponseActionsInner) GetObjectOk() (*string, bool) {
	if o == nil || IsNil(o.Object) {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *CastComposerActionsListResponseActionsInner) HasObject() bool {
	if o != nil && !IsNil(o.Object) {
		return true
	}

	return false
}

// SetObject gets a reference to the given string and assigns it to the Object field.
func (o *CastComposerActionsListResponseActionsInner) SetObject(v string) {
	o.Object = &v
}

func (o CastComposerActionsListResponseActionsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CastComposerActionsListResponseActionsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Icon) {
		toSerialize["icon"] = o.Icon
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.AboutUrl) {
		toSerialize["about_url"] = o.AboutUrl
	}
	if !IsNil(o.ImageUrl) {
		toSerialize["image_url"] = o.ImageUrl
	}
	if !IsNil(o.ActionUrl) {
		toSerialize["action_url"] = o.ActionUrl
	}
	if !IsNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	if !IsNil(o.Octicon) {
		toSerialize["octicon"] = o.Octicon
	}
	if !IsNil(o.AddedCount) {
		toSerialize["added_count"] = o.AddedCount
	}
	if !IsNil(o.AppName) {
		toSerialize["app_name"] = o.AppName
	}
	if !IsNil(o.AuthorFid) {
		toSerialize["author_fid"] = o.AuthorFid
	}
	if !IsNil(o.Category) {
		toSerialize["category"] = o.Category
	}
	if !IsNil(o.Object) {
		toSerialize["object"] = o.Object
	}
	return toSerialize, nil
}

type NullableCastComposerActionsListResponseActionsInner struct {
	value *CastComposerActionsListResponseActionsInner
	isSet bool
}

func (v NullableCastComposerActionsListResponseActionsInner) Get() *CastComposerActionsListResponseActionsInner {
	return v.value
}

func (v *NullableCastComposerActionsListResponseActionsInner) Set(val *CastComposerActionsListResponseActionsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCastComposerActionsListResponseActionsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCastComposerActionsListResponseActionsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCastComposerActionsListResponseActionsInner(val *CastComposerActionsListResponseActionsInner) *NullableCastComposerActionsListResponseActionsInner {
	return &NullableCastComposerActionsListResponseActionsInner{value: val, isSet: true}
}

func (v NullableCastComposerActionsListResponseActionsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCastComposerActionsListResponseActionsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
