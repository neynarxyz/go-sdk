# FetchUserData200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hash** | **string** |  | 
**HashScheme** | [**HashScheme**](HashScheme.md) |  | [default to HASH_SCHEME_BLAKE3]
**Signature** | **string** |  | 
**SignatureScheme** | [**SignatureScheme**](SignatureScheme.md) |  | [default to SIGNATURE_SCHEME_ED25519]
**Signer** | **string** |  | 
**Data** | [**UserDataAddAllOfData**](UserDataAddAllOfData.md) |  | 
**Messages** | [**[]UserDataAdd**](UserDataAdd.md) |  | 
**NextPageToken** | **string** | Base64-encoded pagination token for fetching the next page of results. An empty value indicates there are no more pages to return. Used in conjunction with the pageSize parameter to implement pagination across large result sets. | 

## Methods

### NewFetchUserData200Response

`func NewFetchUserData200Response(hash string, hashScheme HashScheme, signature string, signatureScheme SignatureScheme, signer string, data UserDataAddAllOfData, messages []UserDataAdd, nextPageToken string, ) *FetchUserData200Response`

NewFetchUserData200Response instantiates a new FetchUserData200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFetchUserData200ResponseWithDefaults

`func NewFetchUserData200ResponseWithDefaults() *FetchUserData200Response`

NewFetchUserData200ResponseWithDefaults instantiates a new FetchUserData200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHash

`func (o *FetchUserData200Response) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *FetchUserData200Response) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *FetchUserData200Response) SetHash(v string)`

SetHash sets Hash field to given value.


### GetHashScheme

`func (o *FetchUserData200Response) GetHashScheme() HashScheme`

GetHashScheme returns the HashScheme field if non-nil, zero value otherwise.

### GetHashSchemeOk

`func (o *FetchUserData200Response) GetHashSchemeOk() (*HashScheme, bool)`

GetHashSchemeOk returns a tuple with the HashScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashScheme

`func (o *FetchUserData200Response) SetHashScheme(v HashScheme)`

SetHashScheme sets HashScheme field to given value.


### GetSignature

`func (o *FetchUserData200Response) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *FetchUserData200Response) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *FetchUserData200Response) SetSignature(v string)`

SetSignature sets Signature field to given value.


### GetSignatureScheme

`func (o *FetchUserData200Response) GetSignatureScheme() SignatureScheme`

GetSignatureScheme returns the SignatureScheme field if non-nil, zero value otherwise.

### GetSignatureSchemeOk

`func (o *FetchUserData200Response) GetSignatureSchemeOk() (*SignatureScheme, bool)`

GetSignatureSchemeOk returns a tuple with the SignatureScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureScheme

`func (o *FetchUserData200Response) SetSignatureScheme(v SignatureScheme)`

SetSignatureScheme sets SignatureScheme field to given value.


### GetSigner

`func (o *FetchUserData200Response) GetSigner() string`

GetSigner returns the Signer field if non-nil, zero value otherwise.

### GetSignerOk

`func (o *FetchUserData200Response) GetSignerOk() (*string, bool)`

GetSignerOk returns a tuple with the Signer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigner

`func (o *FetchUserData200Response) SetSigner(v string)`

SetSigner sets Signer field to given value.


### GetData

`func (o *FetchUserData200Response) GetData() UserDataAddAllOfData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *FetchUserData200Response) GetDataOk() (*UserDataAddAllOfData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *FetchUserData200Response) SetData(v UserDataAddAllOfData)`

SetData sets Data field to given value.


### GetMessages

`func (o *FetchUserData200Response) GetMessages() []UserDataAdd`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *FetchUserData200Response) GetMessagesOk() (*[]UserDataAdd, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *FetchUserData200Response) SetMessages(v []UserDataAdd)`

SetMessages sets Messages field to given value.


### GetNextPageToken

`func (o *FetchUserData200Response) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *FetchUserData200Response) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *FetchUserData200Response) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


