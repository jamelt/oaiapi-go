// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

type ListModelsResponseObject string

const (
	ListModelsResponseObjectList ListModelsResponseObject = "list"
)

func (e ListModelsResponseObject) ToPointer() *ListModelsResponseObject {
	return &e
}
func (e *ListModelsResponseObject) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "list":
		*e = ListModelsResponseObject(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListModelsResponseObject: %v", v)
	}
}

type ListModelsResponse struct {
	Object ListModelsResponseObject `json:"object"`
	Data   []Model                  `json:"data"`
}

func (o *ListModelsResponse) GetObject() ListModelsResponseObject {
	if o == nil {
		return ListModelsResponseObject("")
	}
	return o.Object
}

func (o *ListModelsResponse) GetData() []Model {
	if o == nil {
		return []Model{}
	}
	return o.Data
}
