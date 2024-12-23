// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

// ModelObject - The object type, which is always "model".
type ModelObject string

const (
	ModelObjectModel ModelObject = "model"
)

func (e ModelObject) ToPointer() *ModelObject {
	return &e
}
func (e *ModelObject) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "model":
		*e = ModelObject(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ModelObject: %v", v)
	}
}

// Model - Describes an OpenAI model offering that can be used with the API.
type Model struct {
	// The model identifier, which can be referenced in the API endpoints.
	ID string `json:"id"`
	// The Unix timestamp (in seconds) when the model was created.
	Created int64 `json:"created"`
	// The object type, which is always "model".
	Object ModelObject `json:"object"`
	// The organization that owns the model.
	OwnedBy string `json:"owned_by"`
}

func (o *Model) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *Model) GetCreated() int64 {
	if o == nil {
		return 0
	}
	return o.Created
}

func (o *Model) GetObject() ModelObject {
	if o == nil {
		return ModelObject("")
	}
	return o.Object
}

func (o *Model) GetOwnedBy() string {
	if o == nil {
		return ""
	}
	return o.OwnedBy
}
