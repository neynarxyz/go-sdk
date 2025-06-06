# HtmlMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Favicon** | Pointer to **string** |  | [optional] 
**ModifiedTime** | Pointer to **string** |  | [optional] 
**OgArticleAuthor** | Pointer to **string** |  | [optional] 
**OgArticleExpirationTime** | Pointer to **string** |  | [optional] 
**OgArticleModifiedTime** | Pointer to **string** |  | [optional] 
**OgArticlePublishedTime** | Pointer to **string** |  | [optional] 
**OgArticlePublisher** | Pointer to **string** |  | [optional] 
**OgArticleSection** | Pointer to **string** |  | [optional] 
**OgArticleTag** | Pointer to **string** |  | [optional] 
**OgAudio** | Pointer to **string** |  | [optional] 
**OgAudioSecureURL** | Pointer to **string** |  | [optional] 
**OgAudioType** | Pointer to **string** |  | [optional] 
**OgAudioURL** | Pointer to **string** |  | [optional] 
**OgAvailability** | Pointer to **string** |  | [optional] 
**OgDate** | Pointer to **string** |  | [optional] 
**OgDescription** | Pointer to **string** |  | [optional] 
**OgDeterminer** | Pointer to **string** |  | [optional] 
**OgEpisode** | Pointer to **string** |  | [optional] 
**OgImage** | Pointer to [**[]ImageObject**](ImageObject.md) |  | [optional] 
**OgLocale** | Pointer to **string** |  | [optional] 
**OgLocaleAlternate** | Pointer to **string** |  | [optional] 
**OgLogo** | Pointer to **string** |  | [optional] 
**OgMovie** | Pointer to **string** |  | [optional] 
**OgPriceAmount** | Pointer to **string** |  | [optional] 
**OgPriceCurrency** | Pointer to **string** |  | [optional] 
**OgProductAvailability** | Pointer to **string** |  | [optional] 
**OgProductCondition** | Pointer to **string** |  | [optional] 
**OgProductPriceAmount** | Pointer to **string** |  | [optional] 
**OgProductPriceCurrency** | Pointer to **string** |  | [optional] 
**OgProductRetailerItemId** | Pointer to **string** |  | [optional] 
**OgSiteName** | Pointer to **string** |  | [optional] 
**OgTitle** | Pointer to **string** |  | [optional] 
**OgType** | Pointer to **string** |  | [optional] 
**OgUrl** | Pointer to **string** |  | [optional] 
**OgVideo** | Pointer to [**[]VideoObject**](VideoObject.md) |  | [optional] 
**OgVideoActor** | Pointer to **string** |  | [optional] 
**OgVideoActorId** | Pointer to **string** |  | [optional] 
**OgVideoActorRole** | Pointer to **string** |  | [optional] 
**OgVideoDirector** | Pointer to **string** |  | [optional] 
**OgVideoDuration** | Pointer to **string** |  | [optional] 
**OgVideoOther** | Pointer to **string** |  | [optional] 
**OgVideoReleaseDate** | Pointer to **string** |  | [optional] 
**OgVideoSecureURL** | Pointer to **string** |  | [optional] 
**OgVideoSeries** | Pointer to **string** |  | [optional] 
**OgVideoTag** | Pointer to **string** |  | [optional] 
**OgVideoTvShow** | Pointer to **string** |  | [optional] 
**OgVideoWriter** | Pointer to **string** |  | [optional] 
**OgWebsite** | Pointer to **string** |  | [optional] 
**UpdatedTime** | Pointer to **string** |  | [optional] 
**Oembed** | Pointer to [**HtmlMetadataOembed**](HtmlMetadataOembed.md) |  | [optional] 

## Methods

### NewHtmlMetadata

`func NewHtmlMetadata() *HtmlMetadata`

NewHtmlMetadata instantiates a new HtmlMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHtmlMetadataWithDefaults

`func NewHtmlMetadataWithDefaults() *HtmlMetadata`

NewHtmlMetadataWithDefaults instantiates a new HtmlMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFavicon

`func (o *HtmlMetadata) GetFavicon() string`

GetFavicon returns the Favicon field if non-nil, zero value otherwise.

### GetFaviconOk

`func (o *HtmlMetadata) GetFaviconOk() (*string, bool)`

GetFaviconOk returns a tuple with the Favicon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFavicon

`func (o *HtmlMetadata) SetFavicon(v string)`

SetFavicon sets Favicon field to given value.

### HasFavicon

`func (o *HtmlMetadata) HasFavicon() bool`

HasFavicon returns a boolean if a field has been set.

### GetModifiedTime

`func (o *HtmlMetadata) GetModifiedTime() string`

GetModifiedTime returns the ModifiedTime field if non-nil, zero value otherwise.

### GetModifiedTimeOk

`func (o *HtmlMetadata) GetModifiedTimeOk() (*string, bool)`

GetModifiedTimeOk returns a tuple with the ModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedTime

`func (o *HtmlMetadata) SetModifiedTime(v string)`

SetModifiedTime sets ModifiedTime field to given value.

### HasModifiedTime

`func (o *HtmlMetadata) HasModifiedTime() bool`

HasModifiedTime returns a boolean if a field has been set.

### GetOgArticleAuthor

`func (o *HtmlMetadata) GetOgArticleAuthor() string`

GetOgArticleAuthor returns the OgArticleAuthor field if non-nil, zero value otherwise.

### GetOgArticleAuthorOk

`func (o *HtmlMetadata) GetOgArticleAuthorOk() (*string, bool)`

GetOgArticleAuthorOk returns a tuple with the OgArticleAuthor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOgArticleAuthor

`func (o *HtmlMetadata) SetOgArticleAuthor(v string)`

SetOgArticleAuthor sets OgArticleAuthor field to given value.

### HasOgArticleAuthor

`func (o *HtmlMetadata) HasOgArticleAuthor() bool`

HasOgArticleAuthor returns a boolean if a field has been set.

### GetOgArticleExpirationTime

`func (o *HtmlMetadata) GetOgArticleExpirationTime() string`

GetOgArticleExpirationTime returns the OgArticleExpirationTime field if non-nil, zero value otherwise.

### GetOgArticleExpirationTimeOk

`func (o *HtmlMetadata) GetOgArticleExpirationTimeOk() (*string, bool)`

GetOgArticleExpirationTimeOk returns a tuple with the OgArticleExpirationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOgArticleExpirationTime

`func (o *HtmlMetadata) SetOgArticleExpirationTime(v string)`

SetOgArticleExpirationTime sets OgArticleExpirationTime field to given value.

### HasOgArticleExpirationTime

`func (o *HtmlMetadata) HasOgArticleExpirationTime() bool`

HasOgArticleExpirationTime returns a boolean if a field has been set.

### GetOgArticleModifiedTime

`func (o *HtmlMetadata) GetOgArticleModifiedTime() string`

GetOgArticleModifiedTime returns the OgArticleModifiedTime field if non-nil, zero value otherwise.

### GetOgArticleModifiedTimeOk

`func (o *HtmlMetadata) GetOgArticleModifiedTimeOk() (*string, bool)`

GetOgArticleModifiedTimeOk returns a tuple with the OgArticleModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOgArticleModifiedTime

`func (o *HtmlMetadata) SetOgArticleModifiedTime(v string)`

SetOgArticleModifiedTime sets OgArticleModifiedTime field to given value.

### HasOgArticleModifiedTime

`func (o *HtmlMetadata) HasOgArticleModifiedTime() bool`

HasOgArticleModifiedTime returns a boolean if a field has been set.

### GetOgArticlePublishedTime

`func (o *HtmlMetadata) GetOgArticlePublishedTime() string`

GetOgArticlePublishedTime returns the OgArticlePublishedTime field if non-nil, zero value otherwise.

### GetOgArticlePublishedTimeOk

`func (o *HtmlMetadata) GetOgArticlePublishedTimeOk() (*string, bool)`

GetOgArticlePublishedTimeOk returns a tuple with the OgArticlePublishedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOgArticlePublishedTime

`func (o *HtmlMetadata) SetOgArticlePublishedTime(v string)`

SetOgArticlePublishedTime sets OgArticlePublishedTime field to given value.

### HasOgArticlePublishedTime

`func (o *HtmlMetadata) HasOgArticlePublishedTime() bool`

HasOgArticlePublishedTime returns a boolean if a field has been set.

### GetOgArticlePublisher

`func (o *HtmlMetadata) GetOgArticlePublisher() string`

GetOgArticlePublisher returns the OgArticlePublisher field if non-nil, zero value otherwise.

### GetOgArticlePublisherOk

`func (o *HtmlMetadata) GetOgArticlePublisherOk() (*string, bool)`

GetOgArticlePublisherOk returns a tuple with the OgArticlePublisher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOgArticlePublisher

`func (o *HtmlMetadata) SetOgArticlePublisher(v string)`

SetOgArticlePublisher sets OgArticlePublisher field to given value.

### HasOgArticlePublisher

`func (o *HtmlMetadata) HasOgArticlePublisher() bool`

HasOgArticlePublisher returns a boolean if a field has been set.

### GetOgArticleSection

`func (o *HtmlMetadata) GetOgArticleSection() string`

GetOgArticleSection returns the OgArticleSection field if non-nil, zero value otherwise.

### GetOgArticleSectionOk

`func (o *HtmlMetadata) GetOgArticleSectionOk() (*string, bool)`

GetOgArticleSectionOk returns a tuple with the OgArticleSection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOgArticleSection

`func (o *HtmlMetadata) SetOgArticleSection(v string)`

SetOgArticleSection sets OgArticleSection field to given value.

### HasOgArticleSection

`func (o *HtmlMetadata) HasOgArticleSection() bool`

HasOgArticleSection returns a boolean if a field has been set.

### GetOgArticleTag

`func (o *HtmlMetadata) GetOgArticleTag() string`

GetOgArticleTag returns the OgArticleTag field if non-nil, zero value otherwise.

### GetOgArticleTagOk

`func (o *HtmlMetadata) GetOgArticleTagOk() (*string, bool)`

GetOgArticleTagOk returns a tuple with the OgArticleTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOgArticleTag

`func (o *HtmlMetadata) SetOgArticleTag(v string)`

SetOgArticleTag sets OgArticleTag field to given value.

### HasOgArticleTag

`func (o *HtmlMetadata) HasOgArticleTag() bool`

HasOgArticleTag returns a boolean if a field has been set.

### GetOgAudio

`func (o *HtmlMetadata) GetOgAudio() string`

GetOgAudio returns the OgAudio field if non-nil, zero value otherwise.

### GetOgAudioOk

`func (o *HtmlMetadata) GetOgAudioOk() (*string, bool)`

GetOgAudioOk returns a tuple with the OgAudio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOgAudio

`func (o *HtmlMetadata) SetOgAudio(v string)`

SetOgAudio sets OgAudio field to given value.

### HasOgAudio

`func (o *HtmlMetadata) HasOgAudio() bool`

HasOgAudio returns a boolean if a field has been set.

### GetOgAudioSecureURL

`func (o *HtmlMetadata) GetOgAudioSecureURL() string`

GetOgAudioSecureURL returns the OgAudioSecureURL field if non-nil, zero value otherwise.

### GetOgAudioSecureURLOk

`func (o *HtmlMetadata) GetOgAudioSecureURLOk() (*string, bool)`

GetOgAudioSecureURLOk returns a tuple with the OgAudioSecureURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOgAudioSecureURL

`func (o *HtmlMetadata) SetOgAudioSecureURL(v string)`

SetOgAudioSecureURL sets OgAudioSecureURL field to given value.

### HasOgAudioSecureURL

`func (o *HtmlMetadata) HasOgAudioSecureURL() bool`

HasOgAudioSecureURL returns a boolean if a field has been set.

### GetOgAudioType

`func (o *HtmlMetadata) GetOgAudioType() string`

GetOgAudioType returns the OgAudioType field if non-nil, zero value otherwise.

### GetOgAudioTypeOk

`func (o *HtmlMetadata) GetOgAudioTypeOk() (*string, bool)`

GetOgAudioTypeOk returns a tuple with the OgAudioType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOgAudioType

`func (o *HtmlMetadata) SetOgAudioType(v string)`

SetOgAudioType sets OgAudioType field to given value.

### HasOgAudioType

`func (o *HtmlMetadata) HasOgAudioType() bool`

HasOgAudioType returns a boolean if a field has been set.

### GetOgAudioURL

`func (o *HtmlMetadata) GetOgAudioURL() string`

GetOgAudioURL returns the OgAudioURL field if non-nil, zero value otherwise.

### GetOgAudioURLOk

`func (o *HtmlMetadata) GetOgAudioURLOk() (*string, bool)`

GetOgAudioURLOk returns a tuple with the OgAudioURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOgAudioURL

`func (o *HtmlMetadata) SetOgAudioURL(v string)`

SetOgAudioURL sets OgAudioURL field to given value.

### HasOgAudioURL

`func (o *HtmlMetadata) HasOgAudioURL() bool`

HasOgAudioURL returns a boolean if a field has been set.

### GetOgAvailability

`func (o *HtmlMetadata) GetOgAvailability() string`

GetOgAvailability returns the OgAvailability field if non-nil, zero value otherwise.

### GetOgAvailabilityOk

`func (o *HtmlMetadata) GetOgAvailabilityOk() (*string, bool)`

GetOgAvailabilityOk returns a tuple with the OgAvailability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOgAvailability

`func (o *HtmlMetadata) SetOgAvailability(v string)`

SetOgAvailability sets OgAvailability field to given value.

### HasOgAvailability

`func (o *HtmlMetadata) HasOgAvailability() bool`

HasOgAvailability returns a boolean if a field has been set.

### GetOgDate

`func (o *HtmlMetadata) GetOgDate() string`

GetOgDate returns the OgDate field if non-nil, zero value otherwise.

### GetOgDateOk

`func (o *HtmlMetadata) GetOgDateOk() (*string, bool)`

GetOgDateOk returns a tuple with the OgDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOgDate

`func (o *HtmlMetadata) SetOgDate(v string)`

SetOgDate sets OgDate field to given value.

### HasOgDate

`func (o *HtmlMetadata) HasOgDate() bool`

HasOgDate returns a boolean if a field has been set.

### GetOgDescription

`func (o *HtmlMetadata) GetOgDescription() string`

GetOgDescription returns the OgDescription field if non-nil, zero value otherwise.

### GetOgDescriptionOk

`func (o *HtmlMetadata) GetOgDescriptionOk() (*string, bool)`

GetOgDescriptionOk returns a tuple with the OgDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOgDescription

`func (o *HtmlMetadata) SetOgDescription(v string)`

SetOgDescription sets OgDescription field to given value.

### HasOgDescription

`func (o *HtmlMetadata) HasOgDescription() bool`

HasOgDescription returns a boolean if a field has been set.

### GetOgDeterminer

`func (o *HtmlMetadata) GetOgDeterminer() string`

GetOgDeterminer returns the OgDeterminer field if non-nil, zero value otherwise.

### GetOgDeterminerOk

`func (o *HtmlMetadata) GetOgDeterminerOk() (*string, bool)`

GetOgDeterminerOk returns a tuple with the OgDeterminer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOgDeterminer

`func (o *HtmlMetadata) SetOgDeterminer(v string)`

SetOgDeterminer sets OgDeterminer field to given value.

### HasOgDeterminer

`func (o *HtmlMetadata) HasOgDeterminer() bool`

HasOgDeterminer returns a boolean if a field has been set.

### GetOgEpisode

`func (o *HtmlMetadata) GetOgEpisode() string`

GetOgEpisode returns the OgEpisode field if non-nil, zero value otherwise.

### GetOgEpisodeOk

`func (o *HtmlMetadata) GetOgEpisodeOk() (*string, bool)`

GetOgEpisodeOk returns a tuple with the OgEpisode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOgEpisode

`func (o *HtmlMetadata) SetOgEpisode(v string)`

SetOgEpisode sets OgEpisode field to given value.

### HasOgEpisode

`func (o *HtmlMetadata) HasOgEpisode() bool`

HasOgEpisode returns a boolean if a field has been set.

### GetOgImage

`func (o *HtmlMetadata) GetOgImage() []ImageObject`

GetOgImage returns the OgImage field if non-nil, zero value otherwise.

### GetOgImageOk

`func (o *HtmlMetadata) GetOgImageOk() (*[]ImageObject, bool)`

GetOgImageOk returns a tuple with the OgImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOgImage

`func (o *HtmlMetadata) SetOgImage(v []ImageObject)`

SetOgImage sets OgImage field to given value.

### HasOgImage

`func (o *HtmlMetadata) HasOgImage() bool`

HasOgImage returns a boolean if a field has been set.

### GetOgLocale

`func (o *HtmlMetadata) GetOgLocale() string`

GetOgLocale returns the OgLocale field if non-nil, zero value otherwise.

### GetOgLocaleOk

`func (o *HtmlMetadata) GetOgLocaleOk() (*string, bool)`

GetOgLocaleOk returns a tuple with the OgLocale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOgLocale

`func (o *HtmlMetadata) SetOgLocale(v string)`

SetOgLocale sets OgLocale field to given value.

### HasOgLocale

`func (o *HtmlMetadata) HasOgLocale() bool`

HasOgLocale returns a boolean if a field has been set.

### GetOgLocaleAlternate

`func (o *HtmlMetadata) GetOgLocaleAlternate() string`

GetOgLocaleAlternate returns the OgLocaleAlternate field if non-nil, zero value otherwise.

### GetOgLocaleAlternateOk

`func (o *HtmlMetadata) GetOgLocaleAlternateOk() (*string, bool)`

GetOgLocaleAlternateOk returns a tuple with the OgLocaleAlternate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOgLocaleAlternate

`func (o *HtmlMetadata) SetOgLocaleAlternate(v string)`

SetOgLocaleAlternate sets OgLocaleAlternate field to given value.

### HasOgLocaleAlternate

`func (o *HtmlMetadata) HasOgLocaleAlternate() bool`

HasOgLocaleAlternate returns a boolean if a field has been set.

### GetOgLogo

`func (o *HtmlMetadata) GetOgLogo() string`

GetOgLogo returns the OgLogo field if non-nil, zero value otherwise.

### GetOgLogoOk

`func (o *HtmlMetadata) GetOgLogoOk() (*string, bool)`

GetOgLogoOk returns a tuple with the OgLogo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOgLogo

`func (o *HtmlMetadata) SetOgLogo(v string)`

SetOgLogo sets OgLogo field to given value.

### HasOgLogo

`func (o *HtmlMetadata) HasOgLogo() bool`

HasOgLogo returns a boolean if a field has been set.

### GetOgMovie

`func (o *HtmlMetadata) GetOgMovie() string`

GetOgMovie returns the OgMovie field if non-nil, zero value otherwise.

### GetOgMovieOk

`func (o *HtmlMetadata) GetOgMovieOk() (*string, bool)`

GetOgMovieOk returns a tuple with the OgMovie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOgMovie

`func (o *HtmlMetadata) SetOgMovie(v string)`

SetOgMovie sets OgMovie field to given value.

### HasOgMovie

`func (o *HtmlMetadata) HasOgMovie() bool`

HasOgMovie returns a boolean if a field has been set.

### GetOgPriceAmount

`func (o *HtmlMetadata) GetOgPriceAmount() string`

GetOgPriceAmount returns the OgPriceAmount field if non-nil, zero value otherwise.

### GetOgPriceAmountOk

`func (o *HtmlMetadata) GetOgPriceAmountOk() (*string, bool)`

GetOgPriceAmountOk returns a tuple with the OgPriceAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOgPriceAmount

`func (o *HtmlMetadata) SetOgPriceAmount(v string)`

SetOgPriceAmount sets OgPriceAmount field to given value.

### HasOgPriceAmount

`func (o *HtmlMetadata) HasOgPriceAmount() bool`

HasOgPriceAmount returns a boolean if a field has been set.

### GetOgPriceCurrency

`func (o *HtmlMetadata) GetOgPriceCurrency() string`

GetOgPriceCurrency returns the OgPriceCurrency field if non-nil, zero value otherwise.

### GetOgPriceCurrencyOk

`func (o *HtmlMetadata) GetOgPriceCurrencyOk() (*string, bool)`

GetOgPriceCurrencyOk returns a tuple with the OgPriceCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOgPriceCurrency

`func (o *HtmlMetadata) SetOgPriceCurrency(v string)`

SetOgPriceCurrency sets OgPriceCurrency field to given value.

### HasOgPriceCurrency

`func (o *HtmlMetadata) HasOgPriceCurrency() bool`

HasOgPriceCurrency returns a boolean if a field has been set.

### GetOgProductAvailability

`func (o *HtmlMetadata) GetOgProductAvailability() string`

GetOgProductAvailability returns the OgProductAvailability field if non-nil, zero value otherwise.

### GetOgProductAvailabilityOk

`func (o *HtmlMetadata) GetOgProductAvailabilityOk() (*string, bool)`

GetOgProductAvailabilityOk returns a tuple with the OgProductAvailability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOgProductAvailability

`func (o *HtmlMetadata) SetOgProductAvailability(v string)`

SetOgProductAvailability sets OgProductAvailability field to given value.

### HasOgProductAvailability

`func (o *HtmlMetadata) HasOgProductAvailability() bool`

HasOgProductAvailability returns a boolean if a field has been set.

### GetOgProductCondition

`func (o *HtmlMetadata) GetOgProductCondition() string`

GetOgProductCondition returns the OgProductCondition field if non-nil, zero value otherwise.

### GetOgProductConditionOk

`func (o *HtmlMetadata) GetOgProductConditionOk() (*string, bool)`

GetOgProductConditionOk returns a tuple with the OgProductCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOgProductCondition

`func (o *HtmlMetadata) SetOgProductCondition(v string)`

SetOgProductCondition sets OgProductCondition field to given value.

### HasOgProductCondition

`func (o *HtmlMetadata) HasOgProductCondition() bool`

HasOgProductCondition returns a boolean if a field has been set.

### GetOgProductPriceAmount

`func (o *HtmlMetadata) GetOgProductPriceAmount() string`

GetOgProductPriceAmount returns the OgProductPriceAmount field if non-nil, zero value otherwise.

### GetOgProductPriceAmountOk

`func (o *HtmlMetadata) GetOgProductPriceAmountOk() (*string, bool)`

GetOgProductPriceAmountOk returns a tuple with the OgProductPriceAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOgProductPriceAmount

`func (o *HtmlMetadata) SetOgProductPriceAmount(v string)`

SetOgProductPriceAmount sets OgProductPriceAmount field to given value.

### HasOgProductPriceAmount

`func (o *HtmlMetadata) HasOgProductPriceAmount() bool`

HasOgProductPriceAmount returns a boolean if a field has been set.

### GetOgProductPriceCurrency

`func (o *HtmlMetadata) GetOgProductPriceCurrency() string`

GetOgProductPriceCurrency returns the OgProductPriceCurrency field if non-nil, zero value otherwise.

### GetOgProductPriceCurrencyOk

`func (o *HtmlMetadata) GetOgProductPriceCurrencyOk() (*string, bool)`

GetOgProductPriceCurrencyOk returns a tuple with the OgProductPriceCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOgProductPriceCurrency

`func (o *HtmlMetadata) SetOgProductPriceCurrency(v string)`

SetOgProductPriceCurrency sets OgProductPriceCurrency field to given value.

### HasOgProductPriceCurrency

`func (o *HtmlMetadata) HasOgProductPriceCurrency() bool`

HasOgProductPriceCurrency returns a boolean if a field has been set.

### GetOgProductRetailerItemId

`func (o *HtmlMetadata) GetOgProductRetailerItemId() string`

GetOgProductRetailerItemId returns the OgProductRetailerItemId field if non-nil, zero value otherwise.

### GetOgProductRetailerItemIdOk

`func (o *HtmlMetadata) GetOgProductRetailerItemIdOk() (*string, bool)`

GetOgProductRetailerItemIdOk returns a tuple with the OgProductRetailerItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOgProductRetailerItemId

`func (o *HtmlMetadata) SetOgProductRetailerItemId(v string)`

SetOgProductRetailerItemId sets OgProductRetailerItemId field to given value.

### HasOgProductRetailerItemId

`func (o *HtmlMetadata) HasOgProductRetailerItemId() bool`

HasOgProductRetailerItemId returns a boolean if a field has been set.

### GetOgSiteName

`func (o *HtmlMetadata) GetOgSiteName() string`

GetOgSiteName returns the OgSiteName field if non-nil, zero value otherwise.

### GetOgSiteNameOk

`func (o *HtmlMetadata) GetOgSiteNameOk() (*string, bool)`

GetOgSiteNameOk returns a tuple with the OgSiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOgSiteName

`func (o *HtmlMetadata) SetOgSiteName(v string)`

SetOgSiteName sets OgSiteName field to given value.

### HasOgSiteName

`func (o *HtmlMetadata) HasOgSiteName() bool`

HasOgSiteName returns a boolean if a field has been set.

### GetOgTitle

`func (o *HtmlMetadata) GetOgTitle() string`

GetOgTitle returns the OgTitle field if non-nil, zero value otherwise.

### GetOgTitleOk

`func (o *HtmlMetadata) GetOgTitleOk() (*string, bool)`

GetOgTitleOk returns a tuple with the OgTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOgTitle

`func (o *HtmlMetadata) SetOgTitle(v string)`

SetOgTitle sets OgTitle field to given value.

### HasOgTitle

`func (o *HtmlMetadata) HasOgTitle() bool`

HasOgTitle returns a boolean if a field has been set.

### GetOgType

`func (o *HtmlMetadata) GetOgType() string`

GetOgType returns the OgType field if non-nil, zero value otherwise.

### GetOgTypeOk

`func (o *HtmlMetadata) GetOgTypeOk() (*string, bool)`

GetOgTypeOk returns a tuple with the OgType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOgType

`func (o *HtmlMetadata) SetOgType(v string)`

SetOgType sets OgType field to given value.

### HasOgType

`func (o *HtmlMetadata) HasOgType() bool`

HasOgType returns a boolean if a field has been set.

### GetOgUrl

`func (o *HtmlMetadata) GetOgUrl() string`

GetOgUrl returns the OgUrl field if non-nil, zero value otherwise.

### GetOgUrlOk

`func (o *HtmlMetadata) GetOgUrlOk() (*string, bool)`

GetOgUrlOk returns a tuple with the OgUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOgUrl

`func (o *HtmlMetadata) SetOgUrl(v string)`

SetOgUrl sets OgUrl field to given value.

### HasOgUrl

`func (o *HtmlMetadata) HasOgUrl() bool`

HasOgUrl returns a boolean if a field has been set.

### GetOgVideo

`func (o *HtmlMetadata) GetOgVideo() []VideoObject`

GetOgVideo returns the OgVideo field if non-nil, zero value otherwise.

### GetOgVideoOk

`func (o *HtmlMetadata) GetOgVideoOk() (*[]VideoObject, bool)`

GetOgVideoOk returns a tuple with the OgVideo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOgVideo

`func (o *HtmlMetadata) SetOgVideo(v []VideoObject)`

SetOgVideo sets OgVideo field to given value.

### HasOgVideo

`func (o *HtmlMetadata) HasOgVideo() bool`

HasOgVideo returns a boolean if a field has been set.

### GetOgVideoActor

`func (o *HtmlMetadata) GetOgVideoActor() string`

GetOgVideoActor returns the OgVideoActor field if non-nil, zero value otherwise.

### GetOgVideoActorOk

`func (o *HtmlMetadata) GetOgVideoActorOk() (*string, bool)`

GetOgVideoActorOk returns a tuple with the OgVideoActor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOgVideoActor

`func (o *HtmlMetadata) SetOgVideoActor(v string)`

SetOgVideoActor sets OgVideoActor field to given value.

### HasOgVideoActor

`func (o *HtmlMetadata) HasOgVideoActor() bool`

HasOgVideoActor returns a boolean if a field has been set.

### GetOgVideoActorId

`func (o *HtmlMetadata) GetOgVideoActorId() string`

GetOgVideoActorId returns the OgVideoActorId field if non-nil, zero value otherwise.

### GetOgVideoActorIdOk

`func (o *HtmlMetadata) GetOgVideoActorIdOk() (*string, bool)`

GetOgVideoActorIdOk returns a tuple with the OgVideoActorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOgVideoActorId

`func (o *HtmlMetadata) SetOgVideoActorId(v string)`

SetOgVideoActorId sets OgVideoActorId field to given value.

### HasOgVideoActorId

`func (o *HtmlMetadata) HasOgVideoActorId() bool`

HasOgVideoActorId returns a boolean if a field has been set.

### GetOgVideoActorRole

`func (o *HtmlMetadata) GetOgVideoActorRole() string`

GetOgVideoActorRole returns the OgVideoActorRole field if non-nil, zero value otherwise.

### GetOgVideoActorRoleOk

`func (o *HtmlMetadata) GetOgVideoActorRoleOk() (*string, bool)`

GetOgVideoActorRoleOk returns a tuple with the OgVideoActorRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOgVideoActorRole

`func (o *HtmlMetadata) SetOgVideoActorRole(v string)`

SetOgVideoActorRole sets OgVideoActorRole field to given value.

### HasOgVideoActorRole

`func (o *HtmlMetadata) HasOgVideoActorRole() bool`

HasOgVideoActorRole returns a boolean if a field has been set.

### GetOgVideoDirector

`func (o *HtmlMetadata) GetOgVideoDirector() string`

GetOgVideoDirector returns the OgVideoDirector field if non-nil, zero value otherwise.

### GetOgVideoDirectorOk

`func (o *HtmlMetadata) GetOgVideoDirectorOk() (*string, bool)`

GetOgVideoDirectorOk returns a tuple with the OgVideoDirector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOgVideoDirector

`func (o *HtmlMetadata) SetOgVideoDirector(v string)`

SetOgVideoDirector sets OgVideoDirector field to given value.

### HasOgVideoDirector

`func (o *HtmlMetadata) HasOgVideoDirector() bool`

HasOgVideoDirector returns a boolean if a field has been set.

### GetOgVideoDuration

`func (o *HtmlMetadata) GetOgVideoDuration() string`

GetOgVideoDuration returns the OgVideoDuration field if non-nil, zero value otherwise.

### GetOgVideoDurationOk

`func (o *HtmlMetadata) GetOgVideoDurationOk() (*string, bool)`

GetOgVideoDurationOk returns a tuple with the OgVideoDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOgVideoDuration

`func (o *HtmlMetadata) SetOgVideoDuration(v string)`

SetOgVideoDuration sets OgVideoDuration field to given value.

### HasOgVideoDuration

`func (o *HtmlMetadata) HasOgVideoDuration() bool`

HasOgVideoDuration returns a boolean if a field has been set.

### GetOgVideoOther

`func (o *HtmlMetadata) GetOgVideoOther() string`

GetOgVideoOther returns the OgVideoOther field if non-nil, zero value otherwise.

### GetOgVideoOtherOk

`func (o *HtmlMetadata) GetOgVideoOtherOk() (*string, bool)`

GetOgVideoOtherOk returns a tuple with the OgVideoOther field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOgVideoOther

`func (o *HtmlMetadata) SetOgVideoOther(v string)`

SetOgVideoOther sets OgVideoOther field to given value.

### HasOgVideoOther

`func (o *HtmlMetadata) HasOgVideoOther() bool`

HasOgVideoOther returns a boolean if a field has been set.

### GetOgVideoReleaseDate

`func (o *HtmlMetadata) GetOgVideoReleaseDate() string`

GetOgVideoReleaseDate returns the OgVideoReleaseDate field if non-nil, zero value otherwise.

### GetOgVideoReleaseDateOk

`func (o *HtmlMetadata) GetOgVideoReleaseDateOk() (*string, bool)`

GetOgVideoReleaseDateOk returns a tuple with the OgVideoReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOgVideoReleaseDate

`func (o *HtmlMetadata) SetOgVideoReleaseDate(v string)`

SetOgVideoReleaseDate sets OgVideoReleaseDate field to given value.

### HasOgVideoReleaseDate

`func (o *HtmlMetadata) HasOgVideoReleaseDate() bool`

HasOgVideoReleaseDate returns a boolean if a field has been set.

### GetOgVideoSecureURL

`func (o *HtmlMetadata) GetOgVideoSecureURL() string`

GetOgVideoSecureURL returns the OgVideoSecureURL field if non-nil, zero value otherwise.

### GetOgVideoSecureURLOk

`func (o *HtmlMetadata) GetOgVideoSecureURLOk() (*string, bool)`

GetOgVideoSecureURLOk returns a tuple with the OgVideoSecureURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOgVideoSecureURL

`func (o *HtmlMetadata) SetOgVideoSecureURL(v string)`

SetOgVideoSecureURL sets OgVideoSecureURL field to given value.

### HasOgVideoSecureURL

`func (o *HtmlMetadata) HasOgVideoSecureURL() bool`

HasOgVideoSecureURL returns a boolean if a field has been set.

### GetOgVideoSeries

`func (o *HtmlMetadata) GetOgVideoSeries() string`

GetOgVideoSeries returns the OgVideoSeries field if non-nil, zero value otherwise.

### GetOgVideoSeriesOk

`func (o *HtmlMetadata) GetOgVideoSeriesOk() (*string, bool)`

GetOgVideoSeriesOk returns a tuple with the OgVideoSeries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOgVideoSeries

`func (o *HtmlMetadata) SetOgVideoSeries(v string)`

SetOgVideoSeries sets OgVideoSeries field to given value.

### HasOgVideoSeries

`func (o *HtmlMetadata) HasOgVideoSeries() bool`

HasOgVideoSeries returns a boolean if a field has been set.

### GetOgVideoTag

`func (o *HtmlMetadata) GetOgVideoTag() string`

GetOgVideoTag returns the OgVideoTag field if non-nil, zero value otherwise.

### GetOgVideoTagOk

`func (o *HtmlMetadata) GetOgVideoTagOk() (*string, bool)`

GetOgVideoTagOk returns a tuple with the OgVideoTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOgVideoTag

`func (o *HtmlMetadata) SetOgVideoTag(v string)`

SetOgVideoTag sets OgVideoTag field to given value.

### HasOgVideoTag

`func (o *HtmlMetadata) HasOgVideoTag() bool`

HasOgVideoTag returns a boolean if a field has been set.

### GetOgVideoTvShow

`func (o *HtmlMetadata) GetOgVideoTvShow() string`

GetOgVideoTvShow returns the OgVideoTvShow field if non-nil, zero value otherwise.

### GetOgVideoTvShowOk

`func (o *HtmlMetadata) GetOgVideoTvShowOk() (*string, bool)`

GetOgVideoTvShowOk returns a tuple with the OgVideoTvShow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOgVideoTvShow

`func (o *HtmlMetadata) SetOgVideoTvShow(v string)`

SetOgVideoTvShow sets OgVideoTvShow field to given value.

### HasOgVideoTvShow

`func (o *HtmlMetadata) HasOgVideoTvShow() bool`

HasOgVideoTvShow returns a boolean if a field has been set.

### GetOgVideoWriter

`func (o *HtmlMetadata) GetOgVideoWriter() string`

GetOgVideoWriter returns the OgVideoWriter field if non-nil, zero value otherwise.

### GetOgVideoWriterOk

`func (o *HtmlMetadata) GetOgVideoWriterOk() (*string, bool)`

GetOgVideoWriterOk returns a tuple with the OgVideoWriter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOgVideoWriter

`func (o *HtmlMetadata) SetOgVideoWriter(v string)`

SetOgVideoWriter sets OgVideoWriter field to given value.

### HasOgVideoWriter

`func (o *HtmlMetadata) HasOgVideoWriter() bool`

HasOgVideoWriter returns a boolean if a field has been set.

### GetOgWebsite

`func (o *HtmlMetadata) GetOgWebsite() string`

GetOgWebsite returns the OgWebsite field if non-nil, zero value otherwise.

### GetOgWebsiteOk

`func (o *HtmlMetadata) GetOgWebsiteOk() (*string, bool)`

GetOgWebsiteOk returns a tuple with the OgWebsite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOgWebsite

`func (o *HtmlMetadata) SetOgWebsite(v string)`

SetOgWebsite sets OgWebsite field to given value.

### HasOgWebsite

`func (o *HtmlMetadata) HasOgWebsite() bool`

HasOgWebsite returns a boolean if a field has been set.

### GetUpdatedTime

`func (o *HtmlMetadata) GetUpdatedTime() string`

GetUpdatedTime returns the UpdatedTime field if non-nil, zero value otherwise.

### GetUpdatedTimeOk

`func (o *HtmlMetadata) GetUpdatedTimeOk() (*string, bool)`

GetUpdatedTimeOk returns a tuple with the UpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTime

`func (o *HtmlMetadata) SetUpdatedTime(v string)`

SetUpdatedTime sets UpdatedTime field to given value.

### HasUpdatedTime

`func (o *HtmlMetadata) HasUpdatedTime() bool`

HasUpdatedTime returns a boolean if a field has been set.

### GetOembed

`func (o *HtmlMetadata) GetOembed() HtmlMetadataOembed`

GetOembed returns the Oembed field if non-nil, zero value otherwise.

### GetOembedOk

`func (o *HtmlMetadata) GetOembedOk() (*HtmlMetadataOembed, bool)`

GetOembedOk returns a tuple with the Oembed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOembed

`func (o *HtmlMetadata) SetOembed(v HtmlMetadataOembed)`

SetOembed sets Oembed field to given value.

### HasOembed

`func (o *HtmlMetadata) HasOembed() bool`

HasOembed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


