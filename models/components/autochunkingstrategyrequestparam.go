// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

// AutoChunkingStrategyRequestParamType - Always `auto`.
type AutoChunkingStrategyRequestParamType string

const (
	AutoChunkingStrategyRequestParamTypeAuto AutoChunkingStrategyRequestParamType = "auto"
)

func (e AutoChunkingStrategyRequestParamType) ToPointer() *AutoChunkingStrategyRequestParamType {
	return &e
}
func (e *AutoChunkingStrategyRequestParamType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "auto":
		*e = AutoChunkingStrategyRequestParamType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AutoChunkingStrategyRequestParamType: %v", v)
	}
}

// AutoChunkingStrategyRequestParam - The default strategy. This strategy currently uses a `max_chunk_size_tokens` of `800` and `chunk_overlap_tokens` of `400`.
type AutoChunkingStrategyRequestParam struct {
	// Always `auto`.
	Type AutoChunkingStrategyRequestParamType `json:"type"`
}

func (o *AutoChunkingStrategyRequestParam) GetType() AutoChunkingStrategyRequestParamType {
	if o == nil {
		return AutoChunkingStrategyRequestParamType("")
	}
	return o.Type
}
