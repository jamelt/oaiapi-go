// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jamelt/openai-api/internal/utils"
)

// VectorStoreFileObjectObject - The object type, which is always `vector_store.file`.
type VectorStoreFileObjectObject string

const (
	VectorStoreFileObjectObjectVectorStoreFile VectorStoreFileObjectObject = "vector_store.file"
)

func (e VectorStoreFileObjectObject) ToPointer() *VectorStoreFileObjectObject {
	return &e
}
func (e *VectorStoreFileObjectObject) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "vector_store.file":
		*e = VectorStoreFileObjectObject(v)
		return nil
	default:
		return fmt.Errorf("invalid value for VectorStoreFileObjectObject: %v", v)
	}
}

// VectorStoreFileObjectStatus - The status of the vector store file, which can be either `in_progress`, `completed`, `cancelled`, or `failed`. The status `completed` indicates that the vector store file is ready for use.
type VectorStoreFileObjectStatus string

const (
	VectorStoreFileObjectStatusInProgress VectorStoreFileObjectStatus = "in_progress"
	VectorStoreFileObjectStatusCompleted  VectorStoreFileObjectStatus = "completed"
	VectorStoreFileObjectStatusCancelled  VectorStoreFileObjectStatus = "cancelled"
	VectorStoreFileObjectStatusFailed     VectorStoreFileObjectStatus = "failed"
)

func (e VectorStoreFileObjectStatus) ToPointer() *VectorStoreFileObjectStatus {
	return &e
}
func (e *VectorStoreFileObjectStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "in_progress":
		fallthrough
	case "completed":
		fallthrough
	case "cancelled":
		fallthrough
	case "failed":
		*e = VectorStoreFileObjectStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for VectorStoreFileObjectStatus: %v", v)
	}
}

// VectorStoreFileObjectCode - One of `server_error` or `rate_limit_exceeded`.
type VectorStoreFileObjectCode string

const (
	VectorStoreFileObjectCodeServerError     VectorStoreFileObjectCode = "server_error"
	VectorStoreFileObjectCodeUnsupportedFile VectorStoreFileObjectCode = "unsupported_file"
	VectorStoreFileObjectCodeInvalidFile     VectorStoreFileObjectCode = "invalid_file"
)

func (e VectorStoreFileObjectCode) ToPointer() *VectorStoreFileObjectCode {
	return &e
}
func (e *VectorStoreFileObjectCode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "server_error":
		fallthrough
	case "unsupported_file":
		fallthrough
	case "invalid_file":
		*e = VectorStoreFileObjectCode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for VectorStoreFileObjectCode: %v", v)
	}
}

// VectorStoreFileObjectLastError - The last error associated with this vector store file. Will be `null` if there are no errors.
type VectorStoreFileObjectLastError struct {
	// One of `server_error` or `rate_limit_exceeded`.
	Code VectorStoreFileObjectCode `json:"code"`
	// A human-readable description of the error.
	Message string `json:"message"`
}

func (o *VectorStoreFileObjectLastError) GetCode() VectorStoreFileObjectCode {
	if o == nil {
		return VectorStoreFileObjectCode("")
	}
	return o.Code
}

func (o *VectorStoreFileObjectLastError) GetMessage() string {
	if o == nil {
		return ""
	}
	return o.Message
}

type VectorStoreFileObjectChunkingStrategyType string

const (
	VectorStoreFileObjectChunkingStrategyTypeStaticChunkingStrategyResponseParam VectorStoreFileObjectChunkingStrategyType = "StaticChunkingStrategyResponseParam"
	VectorStoreFileObjectChunkingStrategyTypeOtherChunkingStrategyResponseParam  VectorStoreFileObjectChunkingStrategyType = "OtherChunkingStrategyResponseParam"
)

// VectorStoreFileObjectChunkingStrategy - The strategy used to chunk the file.
type VectorStoreFileObjectChunkingStrategy struct {
	StaticChunkingStrategyResponseParam *StaticChunkingStrategyResponseParam
	OtherChunkingStrategyResponseParam  *OtherChunkingStrategyResponseParam

	Type VectorStoreFileObjectChunkingStrategyType
}

func CreateVectorStoreFileObjectChunkingStrategyStaticChunkingStrategyResponseParam(staticChunkingStrategyResponseParam StaticChunkingStrategyResponseParam) VectorStoreFileObjectChunkingStrategy {
	typ := VectorStoreFileObjectChunkingStrategyTypeStaticChunkingStrategyResponseParam

	return VectorStoreFileObjectChunkingStrategy{
		StaticChunkingStrategyResponseParam: &staticChunkingStrategyResponseParam,
		Type:                                typ,
	}
}

func CreateVectorStoreFileObjectChunkingStrategyOtherChunkingStrategyResponseParam(otherChunkingStrategyResponseParam OtherChunkingStrategyResponseParam) VectorStoreFileObjectChunkingStrategy {
	typ := VectorStoreFileObjectChunkingStrategyTypeOtherChunkingStrategyResponseParam

	return VectorStoreFileObjectChunkingStrategy{
		OtherChunkingStrategyResponseParam: &otherChunkingStrategyResponseParam,
		Type:                               typ,
	}
}

func (u *VectorStoreFileObjectChunkingStrategy) UnmarshalJSON(data []byte) error {

	var otherChunkingStrategyResponseParam OtherChunkingStrategyResponseParam = OtherChunkingStrategyResponseParam{}
	if err := utils.UnmarshalJSON(data, &otherChunkingStrategyResponseParam, "", true, true); err == nil {
		u.OtherChunkingStrategyResponseParam = &otherChunkingStrategyResponseParam
		u.Type = VectorStoreFileObjectChunkingStrategyTypeOtherChunkingStrategyResponseParam
		return nil
	}

	var staticChunkingStrategyResponseParam StaticChunkingStrategyResponseParam = StaticChunkingStrategyResponseParam{}
	if err := utils.UnmarshalJSON(data, &staticChunkingStrategyResponseParam, "", true, true); err == nil {
		u.StaticChunkingStrategyResponseParam = &staticChunkingStrategyResponseParam
		u.Type = VectorStoreFileObjectChunkingStrategyTypeStaticChunkingStrategyResponseParam
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for VectorStoreFileObjectChunkingStrategy", string(data))
}

func (u VectorStoreFileObjectChunkingStrategy) MarshalJSON() ([]byte, error) {
	if u.StaticChunkingStrategyResponseParam != nil {
		return utils.MarshalJSON(u.StaticChunkingStrategyResponseParam, "", true)
	}

	if u.OtherChunkingStrategyResponseParam != nil {
		return utils.MarshalJSON(u.OtherChunkingStrategyResponseParam, "", true)
	}

	return nil, errors.New("could not marshal union type VectorStoreFileObjectChunkingStrategy: all fields are null")
}

// VectorStoreFileObject - A list of files attached to a vector store.
type VectorStoreFileObject struct {
	// The identifier, which can be referenced in API endpoints.
	ID string `json:"id"`
	// The object type, which is always `vector_store.file`.
	Object VectorStoreFileObjectObject `json:"object"`
	// The total vector store usage in bytes. Note that this may be different from the original file size.
	UsageBytes int64 `json:"usage_bytes"`
	// The Unix timestamp (in seconds) for when the vector store file was created.
	CreatedAt int64 `json:"created_at"`
	// The ID of the [vector store](/docs/api-reference/vector-stores/object) that the [File](/docs/api-reference/files) is attached to.
	VectorStoreID string `json:"vector_store_id"`
	// The status of the vector store file, which can be either `in_progress`, `completed`, `cancelled`, or `failed`. The status `completed` indicates that the vector store file is ready for use.
	Status VectorStoreFileObjectStatus `json:"status"`
	// The last error associated with this vector store file. Will be `null` if there are no errors.
	LastError *VectorStoreFileObjectLastError `json:"last_error"`
	// The strategy used to chunk the file.
	ChunkingStrategy *VectorStoreFileObjectChunkingStrategy `json:"chunking_strategy,omitempty"`
}

func (o *VectorStoreFileObject) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *VectorStoreFileObject) GetObject() VectorStoreFileObjectObject {
	if o == nil {
		return VectorStoreFileObjectObject("")
	}
	return o.Object
}

func (o *VectorStoreFileObject) GetUsageBytes() int64 {
	if o == nil {
		return 0
	}
	return o.UsageBytes
}

func (o *VectorStoreFileObject) GetCreatedAt() int64 {
	if o == nil {
		return 0
	}
	return o.CreatedAt
}

func (o *VectorStoreFileObject) GetVectorStoreID() string {
	if o == nil {
		return ""
	}
	return o.VectorStoreID
}

func (o *VectorStoreFileObject) GetStatus() VectorStoreFileObjectStatus {
	if o == nil {
		return VectorStoreFileObjectStatus("")
	}
	return o.Status
}

func (o *VectorStoreFileObject) GetLastError() *VectorStoreFileObjectLastError {
	if o == nil {
		return nil
	}
	return o.LastError
}

func (o *VectorStoreFileObject) GetChunkingStrategy() *VectorStoreFileObjectChunkingStrategy {
	if o == nil {
		return nil
	}
	return o.ChunkingStrategy
}
