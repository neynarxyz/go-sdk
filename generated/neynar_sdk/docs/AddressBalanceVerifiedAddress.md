# AddressBalanceVerifiedAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | The wallet address | 
**Network** | [**Network**](Network.md) |  | 

## Methods

### NewAddressBalanceVerifiedAddress

`func NewAddressBalanceVerifiedAddress(address string, network Network, ) *AddressBalanceVerifiedAddress`

NewAddressBalanceVerifiedAddress instantiates a new AddressBalanceVerifiedAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressBalanceVerifiedAddressWithDefaults

`func NewAddressBalanceVerifiedAddressWithDefaults() *AddressBalanceVerifiedAddress`

NewAddressBalanceVerifiedAddressWithDefaults instantiates a new AddressBalanceVerifiedAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *AddressBalanceVerifiedAddress) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *AddressBalanceVerifiedAddress) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *AddressBalanceVerifiedAddress) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetNetwork

`func (o *AddressBalanceVerifiedAddress) GetNetwork() Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *AddressBalanceVerifiedAddress) GetNetworkOk() (*Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *AddressBalanceVerifiedAddress) SetNetwork(v Network)`

SetNetwork sets Network field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


