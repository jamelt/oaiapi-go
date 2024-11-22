// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

// ChatCompletionRequestMessageContentPartRefusalType - The type of the content part.
type ChatCompletionRequestMessageContentPartRefusalType string

const (
	ChatCompletionRequestMessageContentPartRefusalTypeRefusal ChatCompletionRequestMessageContentPartRefusalType = "refusal"
)

func (e ChatCompletionRequestMessageContentPartRefusalType) ToPointer() *ChatCompletionRequestMessageContentPartRefusalType {
	return &e
}
func (e *ChatCompletionRequestMessageContentPartRefusalType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "refusal":
		*e = ChatCompletionRequestMessageContentPartRefusalType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ChatCompletionRequestMessageContentPartRefusalType: %v", v)
	}
}

type ChatCompletionRequestMessageContentPartRefusal struct {
	// The type of the content part.
	Type ChatCompletionRequestMessageContentPartRefusalType `json:"type"`
	// The refusal message generated by the model.
	Refusal string `json:"refusal"`
}

func (o *ChatCompletionRequestMessageContentPartRefusal) GetType() ChatCompletionRequestMessageContentPartRefusalType {
	if o == nil {
		return ChatCompletionRequestMessageContentPartRefusalType("")
	}
	return o.Type
}

func (o *ChatCompletionRequestMessageContentPartRefusal) GetRefusal() string {
	if o == nil {
		return ""
	}
	return o.Refusal
}