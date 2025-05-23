/*
Farcaster API V2

The Farcaster API allows you to interact with the Farcaster protocol. See the [Neynar docs](https://docs.neynar.com/reference) for more details.

API version: 2.43.0
Contact: team@neynar.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package neynar_sdk

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the FrameV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FrameV2{}

// FrameV2 Mini app v2 object
type FrameV2 struct {
	// Version of the mini app, 'next' for v2, 'vNext' for v1
	Version string `json:"version"`
	// URL of the image
	Image string `json:"image"`
	// Launch URL of the mini app
	FramesUrl string `json:"frames_url"`
	// Button title of a mini app
	Title    *string               `json:"title,omitempty"`
	Manifest *FarcasterManifest    `json:"manifest,omitempty"`
	Author   *UserDehydrated       `json:"author,omitempty"`
	Metadata *FrameV2AllOfMetadata `json:"metadata,omitempty"`
}

type _FrameV2 FrameV2

// NewFrameV2 instantiates a new FrameV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFrameV2(version string, image string, framesUrl string) *FrameV2 {
	this := FrameV2{}
	this.Version = version
	this.Image = image
	this.FramesUrl = framesUrl
	return &this
}

// NewFrameV2WithDefaults instantiates a new FrameV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFrameV2WithDefaults() *FrameV2 {
	this := FrameV2{}
	return &this
}

// GetVersion returns the Version field value
func (o *FrameV2) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *FrameV2) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *FrameV2) SetVersion(v string) {
	o.Version = v
}

// GetImage returns the Image field value
func (o *FrameV2) GetImage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Image
}

// GetImageOk returns a tuple with the Image field value
// and a boolean to check if the value has been set.
func (o *FrameV2) GetImageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Image, true
}

// SetImage sets field value
func (o *FrameV2) SetImage(v string) {
	o.Image = v
}

// GetFramesUrl returns the FramesUrl field value
func (o *FrameV2) GetFramesUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FramesUrl
}

// GetFramesUrlOk returns a tuple with the FramesUrl field value
// and a boolean to check if the value has been set.
func (o *FrameV2) GetFramesUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FramesUrl, true
}

// SetFramesUrl sets field value
func (o *FrameV2) SetFramesUrl(v string) {
	o.FramesUrl = v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *FrameV2) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FrameV2) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *FrameV2) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *FrameV2) SetTitle(v string) {
	o.Title = &v
}

// GetManifest returns the Manifest field value if set, zero value otherwise.
func (o *FrameV2) GetManifest() FarcasterManifest {
	if o == nil || IsNil(o.Manifest) {
		var ret FarcasterManifest
		return ret
	}
	return *o.Manifest
}

// GetManifestOk returns a tuple with the Manifest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FrameV2) GetManifestOk() (*FarcasterManifest, bool) {
	if o == nil || IsNil(o.Manifest) {
		return nil, false
	}
	return o.Manifest, true
}

// HasManifest returns a boolean if a field has been set.
func (o *FrameV2) HasManifest() bool {
	if o != nil && !IsNil(o.Manifest) {
		return true
	}

	return false
}

// SetManifest gets a reference to the given FarcasterManifest and assigns it to the Manifest field.
func (o *FrameV2) SetManifest(v FarcasterManifest) {
	o.Manifest = &v
}

// GetAuthor returns the Author field value if set, zero value otherwise.
func (o *FrameV2) GetAuthor() UserDehydrated {
	if o == nil || IsNil(o.Author) {
		var ret UserDehydrated
		return ret
	}
	return *o.Author
}

// GetAuthorOk returns a tuple with the Author field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FrameV2) GetAuthorOk() (*UserDehydrated, bool) {
	if o == nil || IsNil(o.Author) {
		return nil, false
	}
	return o.Author, true
}

// HasAuthor returns a boolean if a field has been set.
func (o *FrameV2) HasAuthor() bool {
	if o != nil && !IsNil(o.Author) {
		return true
	}

	return false
}

// SetAuthor gets a reference to the given UserDehydrated and assigns it to the Author field.
func (o *FrameV2) SetAuthor(v UserDehydrated) {
	o.Author = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *FrameV2) GetMetadata() FrameV2AllOfMetadata {
	if o == nil || IsNil(o.Metadata) {
		var ret FrameV2AllOfMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FrameV2) GetMetadataOk() (*FrameV2AllOfMetadata, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *FrameV2) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given FrameV2AllOfMetadata and assigns it to the Metadata field.
func (o *FrameV2) SetMetadata(v FrameV2AllOfMetadata) {
	o.Metadata = &v
}

func (o FrameV2) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FrameV2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["version"] = o.Version
	toSerialize["image"] = o.Image
	toSerialize["frames_url"] = o.FramesUrl
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Manifest) {
		toSerialize["manifest"] = o.Manifest
	}
	if !IsNil(o.Author) {
		toSerialize["author"] = o.Author
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	return toSerialize, nil
}

func (o *FrameV2) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"version",
		"image",
		"frames_url",
	}

	// defaultValueFuncMap captures the default values for required properties.
	// These values are used when required properties are missing from the payload.
	defaultValueFuncMap := map[string]func() interface{}{}
	var defaultValueApplied bool
	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if value, exists := allProperties[requiredProperty]; !exists || value == "" {
			if _, ok := defaultValueFuncMap[requiredProperty]; ok {
				allProperties[requiredProperty] = defaultValueFuncMap[requiredProperty]()
				defaultValueApplied = true
			}
		}
		if value, exists := allProperties[requiredProperty]; !exists || value == "" {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	if defaultValueApplied {
		data, err = json.Marshal(allProperties)
		if err != nil {
			return err
		}
	}
	varFrameV2 := _FrameV2{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFrameV2)

	if err != nil {
		return err
	}

	*o = FrameV2(varFrameV2)

	return err
}

type NullableFrameV2 struct {
	value *FrameV2
	isSet bool
}

func (v NullableFrameV2) Get() *FrameV2 {
	return v.value
}

func (v *NullableFrameV2) Set(val *FrameV2) {
	v.value = val
	v.isSet = true
}

func (v NullableFrameV2) IsSet() bool {
	return v.isSet
}

func (v *NullableFrameV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFrameV2(val *FrameV2) *NullableFrameV2 {
	return &NullableFrameV2{value: val, isSet: true}
}

func (v NullableFrameV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFrameV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
