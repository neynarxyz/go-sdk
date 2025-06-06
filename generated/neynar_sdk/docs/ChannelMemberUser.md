# ChannelMemberUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | **string** |  | 
**Fid** | **int32** | The unique identifier of a farcaster user or app (unsigned integer) | 
**Username** | **string** |  | 
**DisplayName** | Pointer to **string** |  | [optional] 
**PfpUrl** | Pointer to **string** | The URL of the user&#39;s profile picture | [optional] 
**CustodyAddress** | **string** | Ethereum address | 
**Profile** | [**UserProfile**](UserProfile.md) |  | 
**FollowerCount** | **int32** | The number of followers the user has. | 
**FollowingCount** | **int32** | The number of users the user is following. | 
**Verifications** | **[]string** |  | 
**VerifiedAddresses** | [**UserVerifiedAddresses**](UserVerifiedAddresses.md) |  | 
**VerifiedAccounts** | [**[]UserVerifiedAccountsInner**](UserVerifiedAccountsInner.md) |  | 
**PowerBadge** | **bool** |  | 
**Experimental** | Pointer to [**UserExperimental**](UserExperimental.md) |  | [optional] 
**ViewerContext** | Pointer to [**UserViewerContext**](UserViewerContext.md) |  | [optional] 
**Score** | **float64** | Score that represents the probability that the account is not spam. | 

## Methods

### NewChannelMemberUser

`func NewChannelMemberUser(object string, fid int32, username string, custodyAddress string, profile UserProfile, followerCount int32, followingCount int32, verifications []string, verifiedAddresses UserVerifiedAddresses, verifiedAccounts []UserVerifiedAccountsInner, powerBadge bool, score float64, ) *ChannelMemberUser`

NewChannelMemberUser instantiates a new ChannelMemberUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelMemberUserWithDefaults

`func NewChannelMemberUserWithDefaults() *ChannelMemberUser`

NewChannelMemberUserWithDefaults instantiates a new ChannelMemberUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *ChannelMemberUser) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *ChannelMemberUser) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *ChannelMemberUser) SetObject(v string)`

SetObject sets Object field to given value.


### GetFid

`func (o *ChannelMemberUser) GetFid() int32`

GetFid returns the Fid field if non-nil, zero value otherwise.

### GetFidOk

`func (o *ChannelMemberUser) GetFidOk() (*int32, bool)`

GetFidOk returns a tuple with the Fid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFid

`func (o *ChannelMemberUser) SetFid(v int32)`

SetFid sets Fid field to given value.


### GetUsername

`func (o *ChannelMemberUser) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ChannelMemberUser) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ChannelMemberUser) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetDisplayName

`func (o *ChannelMemberUser) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ChannelMemberUser) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ChannelMemberUser) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ChannelMemberUser) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetPfpUrl

`func (o *ChannelMemberUser) GetPfpUrl() string`

GetPfpUrl returns the PfpUrl field if non-nil, zero value otherwise.

### GetPfpUrlOk

`func (o *ChannelMemberUser) GetPfpUrlOk() (*string, bool)`

GetPfpUrlOk returns a tuple with the PfpUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPfpUrl

`func (o *ChannelMemberUser) SetPfpUrl(v string)`

SetPfpUrl sets PfpUrl field to given value.

### HasPfpUrl

`func (o *ChannelMemberUser) HasPfpUrl() bool`

HasPfpUrl returns a boolean if a field has been set.

### GetCustodyAddress

`func (o *ChannelMemberUser) GetCustodyAddress() string`

GetCustodyAddress returns the CustodyAddress field if non-nil, zero value otherwise.

### GetCustodyAddressOk

`func (o *ChannelMemberUser) GetCustodyAddressOk() (*string, bool)`

GetCustodyAddressOk returns a tuple with the CustodyAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustodyAddress

`func (o *ChannelMemberUser) SetCustodyAddress(v string)`

SetCustodyAddress sets CustodyAddress field to given value.


### GetProfile

`func (o *ChannelMemberUser) GetProfile() UserProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *ChannelMemberUser) GetProfileOk() (*UserProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *ChannelMemberUser) SetProfile(v UserProfile)`

SetProfile sets Profile field to given value.


### GetFollowerCount

`func (o *ChannelMemberUser) GetFollowerCount() int32`

GetFollowerCount returns the FollowerCount field if non-nil, zero value otherwise.

### GetFollowerCountOk

`func (o *ChannelMemberUser) GetFollowerCountOk() (*int32, bool)`

GetFollowerCountOk returns a tuple with the FollowerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowerCount

`func (o *ChannelMemberUser) SetFollowerCount(v int32)`

SetFollowerCount sets FollowerCount field to given value.


### GetFollowingCount

`func (o *ChannelMemberUser) GetFollowingCount() int32`

GetFollowingCount returns the FollowingCount field if non-nil, zero value otherwise.

### GetFollowingCountOk

`func (o *ChannelMemberUser) GetFollowingCountOk() (*int32, bool)`

GetFollowingCountOk returns a tuple with the FollowingCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowingCount

`func (o *ChannelMemberUser) SetFollowingCount(v int32)`

SetFollowingCount sets FollowingCount field to given value.


### GetVerifications

`func (o *ChannelMemberUser) GetVerifications() []string`

GetVerifications returns the Verifications field if non-nil, zero value otherwise.

### GetVerificationsOk

`func (o *ChannelMemberUser) GetVerificationsOk() (*[]string, bool)`

GetVerificationsOk returns a tuple with the Verifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifications

`func (o *ChannelMemberUser) SetVerifications(v []string)`

SetVerifications sets Verifications field to given value.


### GetVerifiedAddresses

`func (o *ChannelMemberUser) GetVerifiedAddresses() UserVerifiedAddresses`

GetVerifiedAddresses returns the VerifiedAddresses field if non-nil, zero value otherwise.

### GetVerifiedAddressesOk

`func (o *ChannelMemberUser) GetVerifiedAddressesOk() (*UserVerifiedAddresses, bool)`

GetVerifiedAddressesOk returns a tuple with the VerifiedAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifiedAddresses

`func (o *ChannelMemberUser) SetVerifiedAddresses(v UserVerifiedAddresses)`

SetVerifiedAddresses sets VerifiedAddresses field to given value.


### GetVerifiedAccounts

`func (o *ChannelMemberUser) GetVerifiedAccounts() []UserVerifiedAccountsInner`

GetVerifiedAccounts returns the VerifiedAccounts field if non-nil, zero value otherwise.

### GetVerifiedAccountsOk

`func (o *ChannelMemberUser) GetVerifiedAccountsOk() (*[]UserVerifiedAccountsInner, bool)`

GetVerifiedAccountsOk returns a tuple with the VerifiedAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifiedAccounts

`func (o *ChannelMemberUser) SetVerifiedAccounts(v []UserVerifiedAccountsInner)`

SetVerifiedAccounts sets VerifiedAccounts field to given value.


### GetPowerBadge

`func (o *ChannelMemberUser) GetPowerBadge() bool`

GetPowerBadge returns the PowerBadge field if non-nil, zero value otherwise.

### GetPowerBadgeOk

`func (o *ChannelMemberUser) GetPowerBadgeOk() (*bool, bool)`

GetPowerBadgeOk returns a tuple with the PowerBadge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerBadge

`func (o *ChannelMemberUser) SetPowerBadge(v bool)`

SetPowerBadge sets PowerBadge field to given value.


### GetExperimental

`func (o *ChannelMemberUser) GetExperimental() UserExperimental`

GetExperimental returns the Experimental field if non-nil, zero value otherwise.

### GetExperimentalOk

`func (o *ChannelMemberUser) GetExperimentalOk() (*UserExperimental, bool)`

GetExperimentalOk returns a tuple with the Experimental field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperimental

`func (o *ChannelMemberUser) SetExperimental(v UserExperimental)`

SetExperimental sets Experimental field to given value.

### HasExperimental

`func (o *ChannelMemberUser) HasExperimental() bool`

HasExperimental returns a boolean if a field has been set.

### GetViewerContext

`func (o *ChannelMemberUser) GetViewerContext() UserViewerContext`

GetViewerContext returns the ViewerContext field if non-nil, zero value otherwise.

### GetViewerContextOk

`func (o *ChannelMemberUser) GetViewerContextOk() (*UserViewerContext, bool)`

GetViewerContextOk returns a tuple with the ViewerContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewerContext

`func (o *ChannelMemberUser) SetViewerContext(v UserViewerContext)`

SetViewerContext sets ViewerContext field to given value.

### HasViewerContext

`func (o *ChannelMemberUser) HasViewerContext() bool`

HasViewerContext returns a boolean if a field has been set.

### GetScore

`func (o *ChannelMemberUser) GetScore() float64`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *ChannelMemberUser) GetScoreOk() (*float64, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *ChannelMemberUser) SetScore(v float64)`

SetScore sets Score field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


