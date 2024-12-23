// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

// EmbeddingObject - The object type, which is always "embedding".
type EmbeddingObject string

const (
	EmbeddingObjectEmbedding EmbeddingObject = "embedding"
)

func (e EmbeddingObject) ToPointer() *EmbeddingObject {
	return &e
}
func (e *EmbeddingObject) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "embedding":
		*e = EmbeddingObject(v)
		return nil
	default:
		return fmt.Errorf("invalid value for EmbeddingObject: %v", v)
	}
}

// Embedding - Represents an embedding vector returned by embedding endpoint.
type Embedding struct {
	// The index of the embedding in the list of embeddings.
	Index int64 `json:"index"`
	// The embedding vector, which is a list of floats. The length of vector depends on the model as listed in the [embedding guide](/docs/guides/embeddings).
	//
	Embedding []float64 `json:"embedding"`
	// The object type, which is always "embedding".
	Object EmbeddingObject `json:"object"`
}

func (o *Embedding) GetIndex() int64 {
	if o == nil {
		return 0
	}
	return o.Index
}

func (o *Embedding) GetEmbedding() []float64 {
	if o == nil {
		return []float64{}
	}
	return o.Embedding
}

func (o *Embedding) GetObject() EmbeddingObject {
	if o == nil {
		return EmbeddingObject("")
	}
	return o.Object
}
