# AppHostUserStateResponseNotificationsEnabledInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | **string** | Domain of the mini app | 
**Status** | **string** | Status of notifications for this domain (usually &#39;valid&#39;) | 
**UpdatedAt** | **time.Time** | When the notification preference was last updated | 

## Methods

### NewAppHostUserStateResponseNotificationsEnabledInner

`func NewAppHostUserStateResponseNotificationsEnabledInner(domain string, status string, updatedAt time.Time, ) *AppHostUserStateResponseNotificationsEnabledInner`

NewAppHostUserStateResponseNotificationsEnabledInner instantiates a new AppHostUserStateResponseNotificationsEnabledInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppHostUserStateResponseNotificationsEnabledInnerWithDefaults

`func NewAppHostUserStateResponseNotificationsEnabledInnerWithDefaults() *AppHostUserStateResponseNotificationsEnabledInner`

NewAppHostUserStateResponseNotificationsEnabledInnerWithDefaults instantiates a new AppHostUserStateResponseNotificationsEnabledInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *AppHostUserStateResponseNotificationsEnabledInner) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *AppHostUserStateResponseNotificationsEnabledInner) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *AppHostUserStateResponseNotificationsEnabledInner) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetStatus

`func (o *AppHostUserStateResponseNotificationsEnabledInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AppHostUserStateResponseNotificationsEnabledInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AppHostUserStateResponseNotificationsEnabledInner) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetUpdatedAt

`func (o *AppHostUserStateResponseNotificationsEnabledInner) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AppHostUserStateResponseNotificationsEnabledInner) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AppHostUserStateResponseNotificationsEnabledInner) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


