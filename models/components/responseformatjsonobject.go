// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

// ResponseFormatJSONObjectType - The type of response format being defined: `json_object`
type ResponseFormatJSONObjectType string

const (
	ResponseFormatJSONObjectTypeJSONObject ResponseFormatJSONObjectType = "json_object"
)

func (e ResponseFormatJSONObjectType) ToPointer() *ResponseFormatJSONObjectType {
	return &e
}
func (e *ResponseFormatJSONObjectType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "json_object":
		*e = ResponseFormatJSONObjectType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ResponseFormatJSONObjectType: %v", v)
	}
}

type ResponseFormatJSONObject struct {
	// The type of response format being defined: `json_object`
	Type ResponseFormatJSONObjectType `json:"type"`
}

func (o *ResponseFormatJSONObject) GetType() ResponseFormatJSONObjectType {
	if o == nil {
		return ResponseFormatJSONObjectType("")
	}
	return o.Type
}