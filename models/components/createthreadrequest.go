// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jamelt/openai-api/internal/utils"
)

type CreateThreadRequestCodeInterpreter struct {
	// A list of [file](/docs/api-reference/files) IDs made available to the `code_interpreter` tool. There can be a maximum of 20 files associated with the tool.
	//
	FileIds []string `json:"file_ids,omitempty"`
}

func (o *CreateThreadRequestCodeInterpreter) GetFileIds() []string {
	if o == nil {
		return nil
	}
	return o.FileIds
}

// CreateThreadRequestChunkingStrategyToolResourcesFileSearch2Type - Always `static`.
type CreateThreadRequestChunkingStrategyToolResourcesFileSearch2Type string

const (
	CreateThreadRequestChunkingStrategyToolResourcesFileSearch2TypeStatic CreateThreadRequestChunkingStrategyToolResourcesFileSearch2Type = "static"
)

func (e CreateThreadRequestChunkingStrategyToolResourcesFileSearch2Type) ToPointer() *CreateThreadRequestChunkingStrategyToolResourcesFileSearch2Type {
	return &e
}
func (e *CreateThreadRequestChunkingStrategyToolResourcesFileSearch2Type) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "static":
		*e = CreateThreadRequestChunkingStrategyToolResourcesFileSearch2Type(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateThreadRequestChunkingStrategyToolResourcesFileSearch2Type: %v", v)
	}
}

type CreateThreadRequestChunkingStrategyToolResourcesStatic struct {
	// The maximum number of tokens in each chunk. The default value is `800`. The minimum value is `100` and the maximum value is `4096`.
	MaxChunkSizeTokens int64 `json:"max_chunk_size_tokens"`
	// The number of tokens that overlap between chunks. The default value is `400`.
	//
	// Note that the overlap must not exceed half of `max_chunk_size_tokens`.
	//
	ChunkOverlapTokens int64 `json:"chunk_overlap_tokens"`
}

func (o *CreateThreadRequestChunkingStrategyToolResourcesStatic) GetMaxChunkSizeTokens() int64 {
	if o == nil {
		return 0
	}
	return o.MaxChunkSizeTokens
}

func (o *CreateThreadRequestChunkingStrategyToolResourcesStatic) GetChunkOverlapTokens() int64 {
	if o == nil {
		return 0
	}
	return o.ChunkOverlapTokens
}

type CreateThreadRequestChunkingStrategyToolResourcesStaticChunkingStrategy struct {
	// Always `static`.
	Type   CreateThreadRequestChunkingStrategyToolResourcesFileSearch2Type `json:"type"`
	Static CreateThreadRequestChunkingStrategyToolResourcesStatic          `json:"static"`
}

func (o *CreateThreadRequestChunkingStrategyToolResourcesStaticChunkingStrategy) GetType() CreateThreadRequestChunkingStrategyToolResourcesFileSearch2Type {
	if o == nil {
		return CreateThreadRequestChunkingStrategyToolResourcesFileSearch2Type("")
	}
	return o.Type
}

func (o *CreateThreadRequestChunkingStrategyToolResourcesStaticChunkingStrategy) GetStatic() CreateThreadRequestChunkingStrategyToolResourcesStatic {
	if o == nil {
		return CreateThreadRequestChunkingStrategyToolResourcesStatic{}
	}
	return o.Static
}

// CreateThreadRequestChunkingStrategyToolResourcesFileSearchType - Always `auto`.
type CreateThreadRequestChunkingStrategyToolResourcesFileSearchType string

const (
	CreateThreadRequestChunkingStrategyToolResourcesFileSearchTypeAuto CreateThreadRequestChunkingStrategyToolResourcesFileSearchType = "auto"
)

func (e CreateThreadRequestChunkingStrategyToolResourcesFileSearchType) ToPointer() *CreateThreadRequestChunkingStrategyToolResourcesFileSearchType {
	return &e
}
func (e *CreateThreadRequestChunkingStrategyToolResourcesFileSearchType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "auto":
		*e = CreateThreadRequestChunkingStrategyToolResourcesFileSearchType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateThreadRequestChunkingStrategyToolResourcesFileSearchType: %v", v)
	}
}

// CreateThreadRequestChunkingStrategyToolResourcesAutoChunkingStrategy - The default strategy. This strategy currently uses a `max_chunk_size_tokens` of `800` and `chunk_overlap_tokens` of `400`.
type CreateThreadRequestChunkingStrategyToolResourcesAutoChunkingStrategy struct {
	// Always `auto`.
	Type CreateThreadRequestChunkingStrategyToolResourcesFileSearchType `json:"type"`
}

func (o *CreateThreadRequestChunkingStrategyToolResourcesAutoChunkingStrategy) GetType() CreateThreadRequestChunkingStrategyToolResourcesFileSearchType {
	if o == nil {
		return CreateThreadRequestChunkingStrategyToolResourcesFileSearchType("")
	}
	return o.Type
}

type CreateThreadRequestFileSearchToolResourcesChunkingStrategyType string

const (
	CreateThreadRequestFileSearchToolResourcesChunkingStrategyTypeCreateThreadRequestChunkingStrategyToolResourcesAutoChunkingStrategy   CreateThreadRequestFileSearchToolResourcesChunkingStrategyType = "CreateThreadRequest_chunking_strategy_tool_resources_Auto Chunking Strategy"
	CreateThreadRequestFileSearchToolResourcesChunkingStrategyTypeCreateThreadRequestChunkingStrategyToolResourcesStaticChunkingStrategy CreateThreadRequestFileSearchToolResourcesChunkingStrategyType = "CreateThreadRequest_chunking_strategy_tool_resources_Static Chunking Strategy"
)

// CreateThreadRequestFileSearchToolResourcesChunkingStrategy - The chunking strategy used to chunk the file(s). If not set, will use the `auto` strategy.
type CreateThreadRequestFileSearchToolResourcesChunkingStrategy struct {
	CreateThreadRequestChunkingStrategyToolResourcesAutoChunkingStrategy   *CreateThreadRequestChunkingStrategyToolResourcesAutoChunkingStrategy
	CreateThreadRequestChunkingStrategyToolResourcesStaticChunkingStrategy *CreateThreadRequestChunkingStrategyToolResourcesStaticChunkingStrategy

	Type CreateThreadRequestFileSearchToolResourcesChunkingStrategyType
}

func CreateCreateThreadRequestFileSearchToolResourcesChunkingStrategyCreateThreadRequestChunkingStrategyToolResourcesAutoChunkingStrategy(createThreadRequestChunkingStrategyToolResourcesAutoChunkingStrategy CreateThreadRequestChunkingStrategyToolResourcesAutoChunkingStrategy) CreateThreadRequestFileSearchToolResourcesChunkingStrategy {
	typ := CreateThreadRequestFileSearchToolResourcesChunkingStrategyTypeCreateThreadRequestChunkingStrategyToolResourcesAutoChunkingStrategy

	return CreateThreadRequestFileSearchToolResourcesChunkingStrategy{
		CreateThreadRequestChunkingStrategyToolResourcesAutoChunkingStrategy: &createThreadRequestChunkingStrategyToolResourcesAutoChunkingStrategy,
		Type: typ,
	}
}

func CreateCreateThreadRequestFileSearchToolResourcesChunkingStrategyCreateThreadRequestChunkingStrategyToolResourcesStaticChunkingStrategy(createThreadRequestChunkingStrategyToolResourcesStaticChunkingStrategy CreateThreadRequestChunkingStrategyToolResourcesStaticChunkingStrategy) CreateThreadRequestFileSearchToolResourcesChunkingStrategy {
	typ := CreateThreadRequestFileSearchToolResourcesChunkingStrategyTypeCreateThreadRequestChunkingStrategyToolResourcesStaticChunkingStrategy

	return CreateThreadRequestFileSearchToolResourcesChunkingStrategy{
		CreateThreadRequestChunkingStrategyToolResourcesStaticChunkingStrategy: &createThreadRequestChunkingStrategyToolResourcesStaticChunkingStrategy,
		Type: typ,
	}
}

func (u *CreateThreadRequestFileSearchToolResourcesChunkingStrategy) UnmarshalJSON(data []byte) error {

	var createThreadRequestChunkingStrategyToolResourcesAutoChunkingStrategy CreateThreadRequestChunkingStrategyToolResourcesAutoChunkingStrategy = CreateThreadRequestChunkingStrategyToolResourcesAutoChunkingStrategy{}
	if err := utils.UnmarshalJSON(data, &createThreadRequestChunkingStrategyToolResourcesAutoChunkingStrategy, "", true, true); err == nil {
		u.CreateThreadRequestChunkingStrategyToolResourcesAutoChunkingStrategy = &createThreadRequestChunkingStrategyToolResourcesAutoChunkingStrategy
		u.Type = CreateThreadRequestFileSearchToolResourcesChunkingStrategyTypeCreateThreadRequestChunkingStrategyToolResourcesAutoChunkingStrategy
		return nil
	}

	var createThreadRequestChunkingStrategyToolResourcesStaticChunkingStrategy CreateThreadRequestChunkingStrategyToolResourcesStaticChunkingStrategy = CreateThreadRequestChunkingStrategyToolResourcesStaticChunkingStrategy{}
	if err := utils.UnmarshalJSON(data, &createThreadRequestChunkingStrategyToolResourcesStaticChunkingStrategy, "", true, true); err == nil {
		u.CreateThreadRequestChunkingStrategyToolResourcesStaticChunkingStrategy = &createThreadRequestChunkingStrategyToolResourcesStaticChunkingStrategy
		u.Type = CreateThreadRequestFileSearchToolResourcesChunkingStrategyTypeCreateThreadRequestChunkingStrategyToolResourcesStaticChunkingStrategy
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for CreateThreadRequestFileSearchToolResourcesChunkingStrategy", string(data))
}

func (u CreateThreadRequestFileSearchToolResourcesChunkingStrategy) MarshalJSON() ([]byte, error) {
	if u.CreateThreadRequestChunkingStrategyToolResourcesAutoChunkingStrategy != nil {
		return utils.MarshalJSON(u.CreateThreadRequestChunkingStrategyToolResourcesAutoChunkingStrategy, "", true)
	}

	if u.CreateThreadRequestChunkingStrategyToolResourcesStaticChunkingStrategy != nil {
		return utils.MarshalJSON(u.CreateThreadRequestChunkingStrategyToolResourcesStaticChunkingStrategy, "", true)
	}

	return nil, errors.New("could not marshal union type CreateThreadRequestFileSearchToolResourcesChunkingStrategy: all fields are null")
}

// CreateThreadRequestFileSearchToolResourcesMetadata - Set of 16 key-value pairs that can be attached to a vector store. This can be useful for storing additional information about the vector store in a structured format. Keys can be a maximum of 64 characters long and values can be a maximum of 512 characters long.
type CreateThreadRequestFileSearchToolResourcesMetadata struct {
}

type CreateThreadRequestFileSearchToolResourcesVectorStores struct {
	// A list of [file](/docs/api-reference/files) IDs to add to the vector store. There can be a maximum of 10000 files in a vector store.
	//
	FileIds []string `json:"file_ids,omitempty"`
	// The chunking strategy used to chunk the file(s). If not set, will use the `auto` strategy.
	ChunkingStrategy *CreateThreadRequestFileSearchToolResourcesChunkingStrategy `json:"chunking_strategy,omitempty"`
	// Set of 16 key-value pairs that can be attached to a vector store. This can be useful for storing additional information about the vector store in a structured format. Keys can be a maximum of 64 characters long and values can be a maximum of 512 characters long.
	//
	Metadata *CreateThreadRequestFileSearchToolResourcesMetadata `json:"metadata,omitempty"`
}

func (o *CreateThreadRequestFileSearchToolResourcesVectorStores) GetFileIds() []string {
	if o == nil {
		return nil
	}
	return o.FileIds
}

func (o *CreateThreadRequestFileSearchToolResourcesVectorStores) GetChunkingStrategy() *CreateThreadRequestFileSearchToolResourcesChunkingStrategy {
	if o == nil {
		return nil
	}
	return o.ChunkingStrategy
}

func (o *CreateThreadRequestFileSearchToolResourcesVectorStores) GetMetadata() *CreateThreadRequestFileSearchToolResourcesMetadata {
	if o == nil {
		return nil
	}
	return o.Metadata
}

type CreateThreadRequestFileSearch2 struct {
	// The [vector store](/docs/api-reference/vector-stores/object) attached to this thread. There can be a maximum of 1 vector store attached to the thread.
	//
	VectorStoreIds []string `json:"vector_store_ids,omitempty"`
	// A helper to create a [vector store](/docs/api-reference/vector-stores/object) with file_ids and attach it to this thread. There can be a maximum of 1 vector store attached to the thread.
	//
	VectorStores []CreateThreadRequestFileSearchToolResourcesVectorStores `json:"vector_stores"`
}

func (o *CreateThreadRequestFileSearch2) GetVectorStoreIds() []string {
	if o == nil {
		return nil
	}
	return o.VectorStoreIds
}

func (o *CreateThreadRequestFileSearch2) GetVectorStores() []CreateThreadRequestFileSearchToolResourcesVectorStores {
	if o == nil {
		return []CreateThreadRequestFileSearchToolResourcesVectorStores{}
	}
	return o.VectorStores
}

// CreateThreadRequestChunkingStrategyToolResourcesType - Always `static`.
type CreateThreadRequestChunkingStrategyToolResourcesType string

const (
	CreateThreadRequestChunkingStrategyToolResourcesTypeStatic CreateThreadRequestChunkingStrategyToolResourcesType = "static"
)

func (e CreateThreadRequestChunkingStrategyToolResourcesType) ToPointer() *CreateThreadRequestChunkingStrategyToolResourcesType {
	return &e
}
func (e *CreateThreadRequestChunkingStrategyToolResourcesType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "static":
		*e = CreateThreadRequestChunkingStrategyToolResourcesType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateThreadRequestChunkingStrategyToolResourcesType: %v", v)
	}
}

type CreateThreadRequestChunkingStrategyStatic struct {
	// The maximum number of tokens in each chunk. The default value is `800`. The minimum value is `100` and the maximum value is `4096`.
	MaxChunkSizeTokens int64 `json:"max_chunk_size_tokens"`
	// The number of tokens that overlap between chunks. The default value is `400`.
	//
	// Note that the overlap must not exceed half of `max_chunk_size_tokens`.
	//
	ChunkOverlapTokens int64 `json:"chunk_overlap_tokens"`
}

func (o *CreateThreadRequestChunkingStrategyStatic) GetMaxChunkSizeTokens() int64 {
	if o == nil {
		return 0
	}
	return o.MaxChunkSizeTokens
}

func (o *CreateThreadRequestChunkingStrategyStatic) GetChunkOverlapTokens() int64 {
	if o == nil {
		return 0
	}
	return o.ChunkOverlapTokens
}

type CreateThreadRequestChunkingStrategyStaticChunkingStrategy struct {
	// Always `static`.
	Type   CreateThreadRequestChunkingStrategyToolResourcesType `json:"type"`
	Static CreateThreadRequestChunkingStrategyStatic            `json:"static"`
}

func (o *CreateThreadRequestChunkingStrategyStaticChunkingStrategy) GetType() CreateThreadRequestChunkingStrategyToolResourcesType {
	if o == nil {
		return CreateThreadRequestChunkingStrategyToolResourcesType("")
	}
	return o.Type
}

func (o *CreateThreadRequestChunkingStrategyStaticChunkingStrategy) GetStatic() CreateThreadRequestChunkingStrategyStatic {
	if o == nil {
		return CreateThreadRequestChunkingStrategyStatic{}
	}
	return o.Static
}

// CreateThreadRequestChunkingStrategyType - Always `auto`.
type CreateThreadRequestChunkingStrategyType string

const (
	CreateThreadRequestChunkingStrategyTypeAuto CreateThreadRequestChunkingStrategyType = "auto"
)

func (e CreateThreadRequestChunkingStrategyType) ToPointer() *CreateThreadRequestChunkingStrategyType {
	return &e
}
func (e *CreateThreadRequestChunkingStrategyType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "auto":
		*e = CreateThreadRequestChunkingStrategyType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateThreadRequestChunkingStrategyType: %v", v)
	}
}

// CreateThreadRequestChunkingStrategyAutoChunkingStrategy - The default strategy. This strategy currently uses a `max_chunk_size_tokens` of `800` and `chunk_overlap_tokens` of `400`.
type CreateThreadRequestChunkingStrategyAutoChunkingStrategy struct {
	// Always `auto`.
	Type CreateThreadRequestChunkingStrategyType `json:"type"`
}

func (o *CreateThreadRequestChunkingStrategyAutoChunkingStrategy) GetType() CreateThreadRequestChunkingStrategyType {
	if o == nil {
		return CreateThreadRequestChunkingStrategyType("")
	}
	return o.Type
}

type CreateThreadRequestFileSearchChunkingStrategyType string

const (
	CreateThreadRequestFileSearchChunkingStrategyTypeCreateThreadRequestChunkingStrategyAutoChunkingStrategy   CreateThreadRequestFileSearchChunkingStrategyType = "CreateThreadRequest_chunking_strategy_Auto Chunking Strategy"
	CreateThreadRequestFileSearchChunkingStrategyTypeCreateThreadRequestChunkingStrategyStaticChunkingStrategy CreateThreadRequestFileSearchChunkingStrategyType = "CreateThreadRequest_chunking_strategy_Static Chunking Strategy"
)

// CreateThreadRequestFileSearchChunkingStrategy - The chunking strategy used to chunk the file(s). If not set, will use the `auto` strategy.
type CreateThreadRequestFileSearchChunkingStrategy struct {
	CreateThreadRequestChunkingStrategyAutoChunkingStrategy   *CreateThreadRequestChunkingStrategyAutoChunkingStrategy
	CreateThreadRequestChunkingStrategyStaticChunkingStrategy *CreateThreadRequestChunkingStrategyStaticChunkingStrategy

	Type CreateThreadRequestFileSearchChunkingStrategyType
}

func CreateCreateThreadRequestFileSearchChunkingStrategyCreateThreadRequestChunkingStrategyAutoChunkingStrategy(createThreadRequestChunkingStrategyAutoChunkingStrategy CreateThreadRequestChunkingStrategyAutoChunkingStrategy) CreateThreadRequestFileSearchChunkingStrategy {
	typ := CreateThreadRequestFileSearchChunkingStrategyTypeCreateThreadRequestChunkingStrategyAutoChunkingStrategy

	return CreateThreadRequestFileSearchChunkingStrategy{
		CreateThreadRequestChunkingStrategyAutoChunkingStrategy: &createThreadRequestChunkingStrategyAutoChunkingStrategy,
		Type: typ,
	}
}

func CreateCreateThreadRequestFileSearchChunkingStrategyCreateThreadRequestChunkingStrategyStaticChunkingStrategy(createThreadRequestChunkingStrategyStaticChunkingStrategy CreateThreadRequestChunkingStrategyStaticChunkingStrategy) CreateThreadRequestFileSearchChunkingStrategy {
	typ := CreateThreadRequestFileSearchChunkingStrategyTypeCreateThreadRequestChunkingStrategyStaticChunkingStrategy

	return CreateThreadRequestFileSearchChunkingStrategy{
		CreateThreadRequestChunkingStrategyStaticChunkingStrategy: &createThreadRequestChunkingStrategyStaticChunkingStrategy,
		Type: typ,
	}
}

func (u *CreateThreadRequestFileSearchChunkingStrategy) UnmarshalJSON(data []byte) error {

	var createThreadRequestChunkingStrategyAutoChunkingStrategy CreateThreadRequestChunkingStrategyAutoChunkingStrategy = CreateThreadRequestChunkingStrategyAutoChunkingStrategy{}
	if err := utils.UnmarshalJSON(data, &createThreadRequestChunkingStrategyAutoChunkingStrategy, "", true, true); err == nil {
		u.CreateThreadRequestChunkingStrategyAutoChunkingStrategy = &createThreadRequestChunkingStrategyAutoChunkingStrategy
		u.Type = CreateThreadRequestFileSearchChunkingStrategyTypeCreateThreadRequestChunkingStrategyAutoChunkingStrategy
		return nil
	}

	var createThreadRequestChunkingStrategyStaticChunkingStrategy CreateThreadRequestChunkingStrategyStaticChunkingStrategy = CreateThreadRequestChunkingStrategyStaticChunkingStrategy{}
	if err := utils.UnmarshalJSON(data, &createThreadRequestChunkingStrategyStaticChunkingStrategy, "", true, true); err == nil {
		u.CreateThreadRequestChunkingStrategyStaticChunkingStrategy = &createThreadRequestChunkingStrategyStaticChunkingStrategy
		u.Type = CreateThreadRequestFileSearchChunkingStrategyTypeCreateThreadRequestChunkingStrategyStaticChunkingStrategy
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for CreateThreadRequestFileSearchChunkingStrategy", string(data))
}

func (u CreateThreadRequestFileSearchChunkingStrategy) MarshalJSON() ([]byte, error) {
	if u.CreateThreadRequestChunkingStrategyAutoChunkingStrategy != nil {
		return utils.MarshalJSON(u.CreateThreadRequestChunkingStrategyAutoChunkingStrategy, "", true)
	}

	if u.CreateThreadRequestChunkingStrategyStaticChunkingStrategy != nil {
		return utils.MarshalJSON(u.CreateThreadRequestChunkingStrategyStaticChunkingStrategy, "", true)
	}

	return nil, errors.New("could not marshal union type CreateThreadRequestFileSearchChunkingStrategy: all fields are null")
}

// CreateThreadRequestFileSearchMetadata - Set of 16 key-value pairs that can be attached to a vector store. This can be useful for storing additional information about the vector store in a structured format. Keys can be a maximum of 64 characters long and values can be a maximum of 512 characters long.
type CreateThreadRequestFileSearchMetadata struct {
}

type CreateThreadRequestFileSearchVectorStores struct {
	// A list of [file](/docs/api-reference/files) IDs to add to the vector store. There can be a maximum of 10000 files in a vector store.
	//
	FileIds []string `json:"file_ids,omitempty"`
	// The chunking strategy used to chunk the file(s). If not set, will use the `auto` strategy.
	ChunkingStrategy *CreateThreadRequestFileSearchChunkingStrategy `json:"chunking_strategy,omitempty"`
	// Set of 16 key-value pairs that can be attached to a vector store. This can be useful for storing additional information about the vector store in a structured format. Keys can be a maximum of 64 characters long and values can be a maximum of 512 characters long.
	//
	Metadata *CreateThreadRequestFileSearchMetadata `json:"metadata,omitempty"`
}

func (o *CreateThreadRequestFileSearchVectorStores) GetFileIds() []string {
	if o == nil {
		return nil
	}
	return o.FileIds
}

func (o *CreateThreadRequestFileSearchVectorStores) GetChunkingStrategy() *CreateThreadRequestFileSearchChunkingStrategy {
	if o == nil {
		return nil
	}
	return o.ChunkingStrategy
}

func (o *CreateThreadRequestFileSearchVectorStores) GetMetadata() *CreateThreadRequestFileSearchMetadata {
	if o == nil {
		return nil
	}
	return o.Metadata
}

type CreateThreadRequestFileSearch1 struct {
	// The [vector store](/docs/api-reference/vector-stores/object) attached to this thread. There can be a maximum of 1 vector store attached to the thread.
	//
	VectorStoreIds []string `json:"vector_store_ids"`
	// A helper to create a [vector store](/docs/api-reference/vector-stores/object) with file_ids and attach it to this thread. There can be a maximum of 1 vector store attached to the thread.
	//
	VectorStores []CreateThreadRequestFileSearchVectorStores `json:"vector_stores,omitempty"`
}

func (o *CreateThreadRequestFileSearch1) GetVectorStoreIds() []string {
	if o == nil {
		return []string{}
	}
	return o.VectorStoreIds
}

func (o *CreateThreadRequestFileSearch1) GetVectorStores() []CreateThreadRequestFileSearchVectorStores {
	if o == nil {
		return nil
	}
	return o.VectorStores
}

type CreateThreadRequestFileSearchType string

const (
	CreateThreadRequestFileSearchTypeCreateThreadRequestFileSearch1 CreateThreadRequestFileSearchType = "CreateThreadRequest_file_search_1"
	CreateThreadRequestFileSearchTypeCreateThreadRequestFileSearch2 CreateThreadRequestFileSearchType = "CreateThreadRequest_file_search_2"
)

type CreateThreadRequestFileSearch struct {
	CreateThreadRequestFileSearch1 *CreateThreadRequestFileSearch1
	CreateThreadRequestFileSearch2 *CreateThreadRequestFileSearch2

	Type CreateThreadRequestFileSearchType
}

func CreateCreateThreadRequestFileSearchCreateThreadRequestFileSearch1(createThreadRequestFileSearch1 CreateThreadRequestFileSearch1) CreateThreadRequestFileSearch {
	typ := CreateThreadRequestFileSearchTypeCreateThreadRequestFileSearch1

	return CreateThreadRequestFileSearch{
		CreateThreadRequestFileSearch1: &createThreadRequestFileSearch1,
		Type:                           typ,
	}
}

func CreateCreateThreadRequestFileSearchCreateThreadRequestFileSearch2(createThreadRequestFileSearch2 CreateThreadRequestFileSearch2) CreateThreadRequestFileSearch {
	typ := CreateThreadRequestFileSearchTypeCreateThreadRequestFileSearch2

	return CreateThreadRequestFileSearch{
		CreateThreadRequestFileSearch2: &createThreadRequestFileSearch2,
		Type:                           typ,
	}
}

func (u *CreateThreadRequestFileSearch) UnmarshalJSON(data []byte) error {

	var createThreadRequestFileSearch1 CreateThreadRequestFileSearch1 = CreateThreadRequestFileSearch1{}
	if err := utils.UnmarshalJSON(data, &createThreadRequestFileSearch1, "", true, true); err == nil {
		u.CreateThreadRequestFileSearch1 = &createThreadRequestFileSearch1
		u.Type = CreateThreadRequestFileSearchTypeCreateThreadRequestFileSearch1
		return nil
	}

	var createThreadRequestFileSearch2 CreateThreadRequestFileSearch2 = CreateThreadRequestFileSearch2{}
	if err := utils.UnmarshalJSON(data, &createThreadRequestFileSearch2, "", true, true); err == nil {
		u.CreateThreadRequestFileSearch2 = &createThreadRequestFileSearch2
		u.Type = CreateThreadRequestFileSearchTypeCreateThreadRequestFileSearch2
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for CreateThreadRequestFileSearch", string(data))
}

func (u CreateThreadRequestFileSearch) MarshalJSON() ([]byte, error) {
	if u.CreateThreadRequestFileSearch1 != nil {
		return utils.MarshalJSON(u.CreateThreadRequestFileSearch1, "", true)
	}

	if u.CreateThreadRequestFileSearch2 != nil {
		return utils.MarshalJSON(u.CreateThreadRequestFileSearch2, "", true)
	}

	return nil, errors.New("could not marshal union type CreateThreadRequestFileSearch: all fields are null")
}

// CreateThreadRequestToolResources - A set of resources that are made available to the assistant's tools in this thread. The resources are specific to the type of tool. For example, the `code_interpreter` tool requires a list of file IDs, while the `file_search` tool requires a list of vector store IDs.
type CreateThreadRequestToolResources struct {
	CodeInterpreter *CreateThreadRequestCodeInterpreter `json:"code_interpreter,omitempty"`
	FileSearch      *CreateThreadRequestFileSearch      `json:"file_search,omitempty"`
}

func (o *CreateThreadRequestToolResources) GetCodeInterpreter() *CreateThreadRequestCodeInterpreter {
	if o == nil {
		return nil
	}
	return o.CodeInterpreter
}

func (o *CreateThreadRequestToolResources) GetFileSearch() *CreateThreadRequestFileSearch {
	if o == nil {
		return nil
	}
	return o.FileSearch
}

// CreateThreadRequestMetadata - Set of 16 key-value pairs that can be attached to an object. This can be useful for storing additional information about the object in a structured format. Keys can be a maximum of 64 characters long and values can be a maximum of 512 characters long.
type CreateThreadRequestMetadata struct {
}

type CreateThreadRequest struct {
	// A list of [messages](/docs/api-reference/messages) to start the thread with.
	Messages []CreateMessageRequest `json:"messages,omitempty"`
	// A set of resources that are made available to the assistant's tools in this thread. The resources are specific to the type of tool. For example, the `code_interpreter` tool requires a list of file IDs, while the `file_search` tool requires a list of vector store IDs.
	//
	ToolResources *CreateThreadRequestToolResources `json:"tool_resources,omitempty"`
	// Set of 16 key-value pairs that can be attached to an object. This can be useful for storing additional information about the object in a structured format. Keys can be a maximum of 64 characters long and values can be a maximum of 512 characters long.
	//
	Metadata *CreateThreadRequestMetadata `json:"metadata,omitempty"`
}

func (o *CreateThreadRequest) GetMessages() []CreateMessageRequest {
	if o == nil {
		return nil
	}
	return o.Messages
}

func (o *CreateThreadRequest) GetToolResources() *CreateThreadRequestToolResources {
	if o == nil {
		return nil
	}
	return o.ToolResources
}

func (o *CreateThreadRequest) GetMetadata() *CreateThreadRequestMetadata {
	if o == nil {
		return nil
	}
	return o.Metadata
}
