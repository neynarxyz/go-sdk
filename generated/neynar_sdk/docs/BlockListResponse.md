# BlockListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Blocks** | [**[]BlockRecord**](BlockRecord.md) |  | 
**Next** | [**NextCursor**](NextCursor.md) |  | 

## Methods

### NewBlockListResponse

`func NewBlockListResponse(blocks []BlockRecord, next NextCursor, ) *BlockListResponse`

NewBlockListResponse instantiates a new BlockListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlockListResponseWithDefaults

`func NewBlockListResponseWithDefaults() *BlockListResponse`

NewBlockListResponseWithDefaults instantiates a new BlockListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlocks

`func (o *BlockListResponse) GetBlocks() []BlockRecord`

GetBlocks returns the Blocks field if non-nil, zero value otherwise.

### GetBlocksOk

`func (o *BlockListResponse) GetBlocksOk() (*[]BlockRecord, bool)`

GetBlocksOk returns a tuple with the Blocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocks

`func (o *BlockListResponse) SetBlocks(v []BlockRecord)`

SetBlocks sets Blocks field to given value.


### GetNext

`func (o *BlockListResponse) GetNext() NextCursor`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *BlockListResponse) GetNextOk() (*NextCursor, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *BlockListResponse) SetNext(v NextCursor)`

SetNext sets Next field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


