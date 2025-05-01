# UserDataAddAllOfData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fid** | **int32** | The unique identifier (FID) of the user who created this message. FIDs are assigned sequentially when users register on the network and cannot be changed. | 
**Timestamp** | **int64** | Seconds since Farcaster Epoch (2021-01-01T00:00:00Z). Used to order messages chronologically and determine the most recent state. Must be within 10 minutes of the current time when the message is created. | 
**Network** | [**FarcasterNetwork**](FarcasterNetwork.md) |  | [default to MAINNET]
**UserDataBody** | [**UserDataBody**](UserDataBody.md) | Contains the type of profile metadata being updated and its new value. | 
**Type** | [**MessageType**](MessageType.md) |  | [default to CAST_ADD]

## Methods

### NewUserDataAddAllOfData

`func NewUserDataAddAllOfData(fid int32, timestamp int64, network FarcasterNetwork, userDataBody UserDataBody, type_ MessageType, ) *UserDataAddAllOfData`

NewUserDataAddAllOfData instantiates a new UserDataAddAllOfData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserDataAddAllOfDataWithDefaults

`func NewUserDataAddAllOfDataWithDefaults() *UserDataAddAllOfData`

NewUserDataAddAllOfDataWithDefaults instantiates a new UserDataAddAllOfData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFid

`func (o *UserDataAddAllOfData) GetFid() int32`

GetFid returns the Fid field if non-nil, zero value otherwise.

### GetFidOk

`func (o *UserDataAddAllOfData) GetFidOk() (*int32, bool)`

GetFidOk returns a tuple with the Fid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFid

`func (o *UserDataAddAllOfData) SetFid(v int32)`

SetFid sets Fid field to given value.


### GetTimestamp

`func (o *UserDataAddAllOfData) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *UserDataAddAllOfData) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *UserDataAddAllOfData) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.


### GetNetwork

`func (o *UserDataAddAllOfData) GetNetwork() FarcasterNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *UserDataAddAllOfData) GetNetworkOk() (*FarcasterNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *UserDataAddAllOfData) SetNetwork(v FarcasterNetwork)`

SetNetwork sets Network field to given value.


### GetUserDataBody

`func (o *UserDataAddAllOfData) GetUserDataBody() UserDataBody`

GetUserDataBody returns the UserDataBody field if non-nil, zero value otherwise.

### GetUserDataBodyOk

`func (o *UserDataAddAllOfData) GetUserDataBodyOk() (*UserDataBody, bool)`

GetUserDataBodyOk returns a tuple with the UserDataBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDataBody

`func (o *UserDataAddAllOfData) SetUserDataBody(v UserDataBody)`

SetUserDataBody sets UserDataBody field to given value.


### GetType

`func (o *UserDataAddAllOfData) GetType() MessageType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserDataAddAllOfData) GetTypeOk() (*MessageType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserDataAddAllOfData) SetType(v MessageType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


