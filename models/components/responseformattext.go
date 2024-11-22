// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

// ResponseFormatTextType - The type of response format being defined: `text`
type ResponseFormatTextType string

const (
	ResponseFormatTextTypeText ResponseFormatTextType = "text"
)

func (e ResponseFormatTextType) ToPointer() *ResponseFormatTextType {
	return &e
}
func (e *ResponseFormatTextType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "text":
		*e = ResponseFormatTextType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ResponseFormatTextType: %v", v)
	}
}

type ResponseFormatText struct {
	// The type of response format being defined: `text`
	Type ResponseFormatTextType `json:"type"`
}

func (o *ResponseFormatText) GetType() ResponseFormatTextType {
	if o == nil {
		return ResponseFormatTextType("")
	}
	return o.Type
}