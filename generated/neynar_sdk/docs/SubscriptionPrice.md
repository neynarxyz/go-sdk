# SubscriptionPrice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PeriodDurationSeconds** | **int32** |  | 
**TokensPerPeriod** | **string** |  | 
**InitialMintPrice** | **string** |  | 

## Methods

### NewSubscriptionPrice

`func NewSubscriptionPrice(periodDurationSeconds int32, tokensPerPeriod string, initialMintPrice string, ) *SubscriptionPrice`

NewSubscriptionPrice instantiates a new SubscriptionPrice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionPriceWithDefaults

`func NewSubscriptionPriceWithDefaults() *SubscriptionPrice`

NewSubscriptionPriceWithDefaults instantiates a new SubscriptionPrice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPeriodDurationSeconds

`func (o *SubscriptionPrice) GetPeriodDurationSeconds() int32`

GetPeriodDurationSeconds returns the PeriodDurationSeconds field if non-nil, zero value otherwise.

### GetPeriodDurationSecondsOk

`func (o *SubscriptionPrice) GetPeriodDurationSecondsOk() (*int32, bool)`

GetPeriodDurationSecondsOk returns a tuple with the PeriodDurationSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodDurationSeconds

`func (o *SubscriptionPrice) SetPeriodDurationSeconds(v int32)`

SetPeriodDurationSeconds sets PeriodDurationSeconds field to given value.


### GetTokensPerPeriod

`func (o *SubscriptionPrice) GetTokensPerPeriod() string`

GetTokensPerPeriod returns the TokensPerPeriod field if non-nil, zero value otherwise.

### GetTokensPerPeriodOk

`func (o *SubscriptionPrice) GetTokensPerPeriodOk() (*string, bool)`

GetTokensPerPeriodOk returns a tuple with the TokensPerPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokensPerPeriod

`func (o *SubscriptionPrice) SetTokensPerPeriod(v string)`

SetTokensPerPeriod sets TokensPerPeriod field to given value.


### GetInitialMintPrice

`func (o *SubscriptionPrice) GetInitialMintPrice() string`

GetInitialMintPrice returns the InitialMintPrice field if non-nil, zero value otherwise.

### GetInitialMintPriceOk

`func (o *SubscriptionPrice) GetInitialMintPriceOk() (*string, bool)`

GetInitialMintPriceOk returns a tuple with the InitialMintPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialMintPrice

`func (o *SubscriptionPrice) SetInitialMintPrice(v string)`

SetInitialMintPrice sets InitialMintPrice field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


