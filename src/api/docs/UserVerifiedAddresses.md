# UserVerifiedAddresses

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EthAddresses** | **[]string** | List of verified Ethereum addresses of the user sorted by oldest to most recent. | 
**SolAddresses** | **[]string** | List of verified Solana addresses of the user sorted by oldest to most recent. | 
**Primary** | [**UserVerifiedAddressesPrimary**](UserVerifiedAddressesPrimary.md) |  | 

## Methods

### NewUserVerifiedAddresses

`func NewUserVerifiedAddresses(ethAddresses []string, solAddresses []string, primary UserVerifiedAddressesPrimary, ) *UserVerifiedAddresses`

NewUserVerifiedAddresses instantiates a new UserVerifiedAddresses object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserVerifiedAddressesWithDefaults

`func NewUserVerifiedAddressesWithDefaults() *UserVerifiedAddresses`

NewUserVerifiedAddressesWithDefaults instantiates a new UserVerifiedAddresses object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEthAddresses

`func (o *UserVerifiedAddresses) GetEthAddresses() []string`

GetEthAddresses returns the EthAddresses field if non-nil, zero value otherwise.

### GetEthAddressesOk

`func (o *UserVerifiedAddresses) GetEthAddressesOk() (*[]string, bool)`

GetEthAddressesOk returns a tuple with the EthAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthAddresses

`func (o *UserVerifiedAddresses) SetEthAddresses(v []string)`

SetEthAddresses sets EthAddresses field to given value.


### GetSolAddresses

`func (o *UserVerifiedAddresses) GetSolAddresses() []string`

GetSolAddresses returns the SolAddresses field if non-nil, zero value otherwise.

### GetSolAddressesOk

`func (o *UserVerifiedAddresses) GetSolAddressesOk() (*[]string, bool)`

GetSolAddressesOk returns a tuple with the SolAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSolAddresses

`func (o *UserVerifiedAddresses) SetSolAddresses(v []string)`

SetSolAddresses sets SolAddresses field to given value.


### GetPrimary

`func (o *UserVerifiedAddresses) GetPrimary() UserVerifiedAddressesPrimary`

GetPrimary returns the Primary field if non-nil, zero value otherwise.

### GetPrimaryOk

`func (o *UserVerifiedAddresses) GetPrimaryOk() (*UserVerifiedAddressesPrimary, bool)`

GetPrimaryOk returns a tuple with the Primary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimary

`func (o *UserVerifiedAddresses) SetPrimary(v UserVerifiedAddressesPrimary)`

SetPrimary sets Primary field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


