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

// checks if the PostCastResponseCast type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostCastResponseCast{}

// PostCastResponseCast struct for PostCastResponseCast
type PostCastResponseCast struct {
	// Cast Hash
	Hash   string                     `json:"hash"`
	Author PostCastResponseCastAuthor `json:"author"`
	Text   string                     `json:"text"`
}

type _PostCastResponseCast PostCastResponseCast

// NewPostCastResponseCast instantiates a new PostCastResponseCast object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostCastResponseCast(hash string, author PostCastResponseCastAuthor, text string) *PostCastResponseCast {
	this := PostCastResponseCast{}
	this.Hash = hash
	this.Author = author
	this.Text = text
	return &this
}

// NewPostCastResponseCastWithDefaults instantiates a new PostCastResponseCast object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostCastResponseCastWithDefaults() *PostCastResponseCast {
	this := PostCastResponseCast{}
	var hash string = "0xfe90f9de682273e05b201629ad2338bdcd89b6be"
	this.Hash = hash
	return &this
}

// GetHash returns the Hash field value
func (o *PostCastResponseCast) GetHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hash
}

// GetHashOk returns a tuple with the Hash field value
// and a boolean to check if the value has been set.
func (o *PostCastResponseCast) GetHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hash, true
}

// SetHash sets field value
func (o *PostCastResponseCast) SetHash(v string) {
	o.Hash = v
}

// GetDefaultHash returns the default value "0xfe90f9de682273e05b201629ad2338bdcd89b6be" of the Hash field.
func (o *PostCastResponseCast) GetDefaultHash() interface{} {
	return "0xfe90f9de682273e05b201629ad2338bdcd89b6be"
}

// GetAuthor returns the Author field value
func (o *PostCastResponseCast) GetAuthor() PostCastResponseCastAuthor {
	if o == nil {
		var ret PostCastResponseCastAuthor
		return ret
	}

	return o.Author
}

// GetAuthorOk returns a tuple with the Author field value
// and a boolean to check if the value has been set.
func (o *PostCastResponseCast) GetAuthorOk() (*PostCastResponseCastAuthor, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Author, true
}

// SetAuthor sets field value
func (o *PostCastResponseCast) SetAuthor(v PostCastResponseCastAuthor) {
	o.Author = v
}

// GetText returns the Text field value
func (o *PostCastResponseCast) GetText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Text
}

// GetTextOk returns a tuple with the Text field value
// and a boolean to check if the value has been set.
func (o *PostCastResponseCast) GetTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Text, true
}

// SetText sets field value
func (o *PostCastResponseCast) SetText(v string) {
	o.Text = v
}

func (o PostCastResponseCast) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostCastResponseCast) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if _, exists := toSerialize["hash"]; !exists {
		toSerialize["hash"] = o.GetDefaultHash()
	}
	toSerialize["hash"] = o.Hash
	toSerialize["author"] = o.Author
	toSerialize["text"] = o.Text
	return toSerialize, nil
}

func (o *PostCastResponseCast) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"hash",
		"author",
		"text",
	}

	// defaultValueFuncMap captures the default values for required properties.
	// These values are used when required properties are missing from the payload.
	defaultValueFuncMap := map[string]func() interface{}{
		"hash": o.GetDefaultHash,
	}
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
	varPostCastResponseCast := _PostCastResponseCast{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostCastResponseCast)

	if err != nil {
		return err
	}

	*o = PostCastResponseCast(varPostCastResponseCast)

	return err
}

type NullablePostCastResponseCast struct {
	value *PostCastResponseCast
	isSet bool
}

func (v NullablePostCastResponseCast) Get() *PostCastResponseCast {
	return v.value
}

func (v *NullablePostCastResponseCast) Set(val *PostCastResponseCast) {
	v.value = val
	v.isSet = true
}

func (v NullablePostCastResponseCast) IsSet() bool {
	return v.isSet
}

func (v *NullablePostCastResponseCast) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostCastResponseCast(val *PostCastResponseCast) *NullablePostCastResponseCast {
	return &NullablePostCastResponseCast{value: val, isSet: true}
}

func (v NullablePostCastResponseCast) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostCastResponseCast) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
