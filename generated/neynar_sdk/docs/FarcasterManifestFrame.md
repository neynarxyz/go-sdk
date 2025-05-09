# FarcasterManifestFrame

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | **string** |  | 
**Name** | **string** |  | 
**HomeUrl** | **string** |  | 
**IconUrl** | **string** |  | 
**ImageUrl** | Pointer to **string** |  | [optional] 
**ButtonTitle** | Pointer to **string** |  | [optional] 
**SplashImageUrl** | Pointer to **string** |  | [optional] 
**SplashBackgroundColor** | Pointer to **string** |  | [optional] 
**WebhookUrl** | Pointer to **string** |  | [optional] 
**Subtitle** | Pointer to **string** | Short subtitle for the frame | [optional] 
**Description** | Pointer to **string** | Detailed description of the frame | [optional] 
**ScreenshotUrls** | Pointer to **[]string** | URLs of screenshots showcasing the frame | [optional] 
**PrimaryCategory** | Pointer to **string** | Primary category the frame belongs to | [optional] 
**Tags** | Pointer to **[]string** | Tags associated with the frame | [optional] 
**HeroImageUrl** | Pointer to **string** | URL of the hero image displayed for the frame | [optional] 
**Tagline** | Pointer to **string** | Short tagline for the frame | [optional] 
**OgTitle** | Pointer to **string** | Title used for Open Graph previews | [optional] 
**OgDescription** | Pointer to **string** | Description used for Open Graph previews | [optional] 
**OgImageUrl** | Pointer to **string** | Image URL used for Open Graph previews | [optional] 
**Noindex** | Pointer to **bool** | Whether search engines should not index this frame | [optional] 

## Methods

### NewFarcasterManifestFrame

`func NewFarcasterManifestFrame(version string, name string, homeUrl string, iconUrl string, ) *FarcasterManifestFrame`

NewFarcasterManifestFrame instantiates a new FarcasterManifestFrame object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFarcasterManifestFrameWithDefaults

`func NewFarcasterManifestFrameWithDefaults() *FarcasterManifestFrame`

NewFarcasterManifestFrameWithDefaults instantiates a new FarcasterManifestFrame object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *FarcasterManifestFrame) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *FarcasterManifestFrame) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *FarcasterManifestFrame) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetName

`func (o *FarcasterManifestFrame) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FarcasterManifestFrame) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FarcasterManifestFrame) SetName(v string)`

SetName sets Name field to given value.


### GetHomeUrl

`func (o *FarcasterManifestFrame) GetHomeUrl() string`

GetHomeUrl returns the HomeUrl field if non-nil, zero value otherwise.

### GetHomeUrlOk

`func (o *FarcasterManifestFrame) GetHomeUrlOk() (*string, bool)`

GetHomeUrlOk returns a tuple with the HomeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeUrl

`func (o *FarcasterManifestFrame) SetHomeUrl(v string)`

SetHomeUrl sets HomeUrl field to given value.


### GetIconUrl

`func (o *FarcasterManifestFrame) GetIconUrl() string`

GetIconUrl returns the IconUrl field if non-nil, zero value otherwise.

### GetIconUrlOk

`func (o *FarcasterManifestFrame) GetIconUrlOk() (*string, bool)`

GetIconUrlOk returns a tuple with the IconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrl

`func (o *FarcasterManifestFrame) SetIconUrl(v string)`

SetIconUrl sets IconUrl field to given value.


### GetImageUrl

`func (o *FarcasterManifestFrame) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *FarcasterManifestFrame) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *FarcasterManifestFrame) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *FarcasterManifestFrame) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### GetButtonTitle

`func (o *FarcasterManifestFrame) GetButtonTitle() string`

GetButtonTitle returns the ButtonTitle field if non-nil, zero value otherwise.

### GetButtonTitleOk

`func (o *FarcasterManifestFrame) GetButtonTitleOk() (*string, bool)`

GetButtonTitleOk returns a tuple with the ButtonTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButtonTitle

`func (o *FarcasterManifestFrame) SetButtonTitle(v string)`

SetButtonTitle sets ButtonTitle field to given value.

### HasButtonTitle

`func (o *FarcasterManifestFrame) HasButtonTitle() bool`

HasButtonTitle returns a boolean if a field has been set.

### GetSplashImageUrl

`func (o *FarcasterManifestFrame) GetSplashImageUrl() string`

GetSplashImageUrl returns the SplashImageUrl field if non-nil, zero value otherwise.

### GetSplashImageUrlOk

`func (o *FarcasterManifestFrame) GetSplashImageUrlOk() (*string, bool)`

GetSplashImageUrlOk returns a tuple with the SplashImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplashImageUrl

`func (o *FarcasterManifestFrame) SetSplashImageUrl(v string)`

SetSplashImageUrl sets SplashImageUrl field to given value.

### HasSplashImageUrl

`func (o *FarcasterManifestFrame) HasSplashImageUrl() bool`

HasSplashImageUrl returns a boolean if a field has been set.

### GetSplashBackgroundColor

`func (o *FarcasterManifestFrame) GetSplashBackgroundColor() string`

GetSplashBackgroundColor returns the SplashBackgroundColor field if non-nil, zero value otherwise.

### GetSplashBackgroundColorOk

`func (o *FarcasterManifestFrame) GetSplashBackgroundColorOk() (*string, bool)`

GetSplashBackgroundColorOk returns a tuple with the SplashBackgroundColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplashBackgroundColor

`func (o *FarcasterManifestFrame) SetSplashBackgroundColor(v string)`

SetSplashBackgroundColor sets SplashBackgroundColor field to given value.

### HasSplashBackgroundColor

`func (o *FarcasterManifestFrame) HasSplashBackgroundColor() bool`

HasSplashBackgroundColor returns a boolean if a field has been set.

### GetWebhookUrl

`func (o *FarcasterManifestFrame) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *FarcasterManifestFrame) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *FarcasterManifestFrame) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.

### HasWebhookUrl

`func (o *FarcasterManifestFrame) HasWebhookUrl() bool`

HasWebhookUrl returns a boolean if a field has been set.

### GetSubtitle

`func (o *FarcasterManifestFrame) GetSubtitle() string`

GetSubtitle returns the Subtitle field if non-nil, zero value otherwise.

### GetSubtitleOk

`func (o *FarcasterManifestFrame) GetSubtitleOk() (*string, bool)`

GetSubtitleOk returns a tuple with the Subtitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtitle

`func (o *FarcasterManifestFrame) SetSubtitle(v string)`

SetSubtitle sets Subtitle field to given value.

### HasSubtitle

`func (o *FarcasterManifestFrame) HasSubtitle() bool`

HasSubtitle returns a boolean if a field has been set.

### GetDescription

`func (o *FarcasterManifestFrame) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FarcasterManifestFrame) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FarcasterManifestFrame) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FarcasterManifestFrame) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetScreenshotUrls

`func (o *FarcasterManifestFrame) GetScreenshotUrls() []string`

GetScreenshotUrls returns the ScreenshotUrls field if non-nil, zero value otherwise.

### GetScreenshotUrlsOk

`func (o *FarcasterManifestFrame) GetScreenshotUrlsOk() (*[]string, bool)`

GetScreenshotUrlsOk returns a tuple with the ScreenshotUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScreenshotUrls

`func (o *FarcasterManifestFrame) SetScreenshotUrls(v []string)`

SetScreenshotUrls sets ScreenshotUrls field to given value.

### HasScreenshotUrls

`func (o *FarcasterManifestFrame) HasScreenshotUrls() bool`

HasScreenshotUrls returns a boolean if a field has been set.

### GetPrimaryCategory

`func (o *FarcasterManifestFrame) GetPrimaryCategory() string`

GetPrimaryCategory returns the PrimaryCategory field if non-nil, zero value otherwise.

### GetPrimaryCategoryOk

`func (o *FarcasterManifestFrame) GetPrimaryCategoryOk() (*string, bool)`

GetPrimaryCategoryOk returns a tuple with the PrimaryCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryCategory

`func (o *FarcasterManifestFrame) SetPrimaryCategory(v string)`

SetPrimaryCategory sets PrimaryCategory field to given value.

### HasPrimaryCategory

`func (o *FarcasterManifestFrame) HasPrimaryCategory() bool`

HasPrimaryCategory returns a boolean if a field has been set.

### GetTags

`func (o *FarcasterManifestFrame) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *FarcasterManifestFrame) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *FarcasterManifestFrame) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *FarcasterManifestFrame) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetHeroImageUrl

`func (o *FarcasterManifestFrame) GetHeroImageUrl() string`

GetHeroImageUrl returns the HeroImageUrl field if non-nil, zero value otherwise.

### GetHeroImageUrlOk

`func (o *FarcasterManifestFrame) GetHeroImageUrlOk() (*string, bool)`

GetHeroImageUrlOk returns a tuple with the HeroImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeroImageUrl

`func (o *FarcasterManifestFrame) SetHeroImageUrl(v string)`

SetHeroImageUrl sets HeroImageUrl field to given value.

### HasHeroImageUrl

`func (o *FarcasterManifestFrame) HasHeroImageUrl() bool`

HasHeroImageUrl returns a boolean if a field has been set.

### GetTagline

`func (o *FarcasterManifestFrame) GetTagline() string`

GetTagline returns the Tagline field if non-nil, zero value otherwise.

### GetTaglineOk

`func (o *FarcasterManifestFrame) GetTaglineOk() (*string, bool)`

GetTaglineOk returns a tuple with the Tagline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagline

`func (o *FarcasterManifestFrame) SetTagline(v string)`

SetTagline sets Tagline field to given value.

### HasTagline

`func (o *FarcasterManifestFrame) HasTagline() bool`

HasTagline returns a boolean if a field has been set.

### GetOgTitle

`func (o *FarcasterManifestFrame) GetOgTitle() string`

GetOgTitle returns the OgTitle field if non-nil, zero value otherwise.

### GetOgTitleOk

`func (o *FarcasterManifestFrame) GetOgTitleOk() (*string, bool)`

GetOgTitleOk returns a tuple with the OgTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOgTitle

`func (o *FarcasterManifestFrame) SetOgTitle(v string)`

SetOgTitle sets OgTitle field to given value.

### HasOgTitle

`func (o *FarcasterManifestFrame) HasOgTitle() bool`

HasOgTitle returns a boolean if a field has been set.

### GetOgDescription

`func (o *FarcasterManifestFrame) GetOgDescription() string`

GetOgDescription returns the OgDescription field if non-nil, zero value otherwise.

### GetOgDescriptionOk

`func (o *FarcasterManifestFrame) GetOgDescriptionOk() (*string, bool)`

GetOgDescriptionOk returns a tuple with the OgDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOgDescription

`func (o *FarcasterManifestFrame) SetOgDescription(v string)`

SetOgDescription sets OgDescription field to given value.

### HasOgDescription

`func (o *FarcasterManifestFrame) HasOgDescription() bool`

HasOgDescription returns a boolean if a field has been set.

### GetOgImageUrl

`func (o *FarcasterManifestFrame) GetOgImageUrl() string`

GetOgImageUrl returns the OgImageUrl field if non-nil, zero value otherwise.

### GetOgImageUrlOk

`func (o *FarcasterManifestFrame) GetOgImageUrlOk() (*string, bool)`

GetOgImageUrlOk returns a tuple with the OgImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOgImageUrl

`func (o *FarcasterManifestFrame) SetOgImageUrl(v string)`

SetOgImageUrl sets OgImageUrl field to given value.

### HasOgImageUrl

`func (o *FarcasterManifestFrame) HasOgImageUrl() bool`

HasOgImageUrl returns a boolean if a field has been set.

### GetNoindex

`func (o *FarcasterManifestFrame) GetNoindex() bool`

GetNoindex returns the Noindex field if non-nil, zero value otherwise.

### GetNoindexOk

`func (o *FarcasterManifestFrame) GetNoindexOk() (*bool, bool)`

GetNoindexOk returns a tuple with the Noindex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoindex

`func (o *FarcasterManifestFrame) SetNoindex(v bool)`

SetNoindex sets Noindex field to given value.

### HasNoindex

`func (o *FarcasterManifestFrame) HasNoindex() bool`

HasNoindex returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


