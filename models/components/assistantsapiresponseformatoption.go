// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jamelt/openai-api/internal/utils"
)

// AssistantsAPIResponseFormatOption1 - `auto` is the default value
type AssistantsAPIResponseFormatOption1 string

const (
	AssistantsAPIResponseFormatOption1Auto AssistantsAPIResponseFormatOption1 = "auto"
)

func (e AssistantsAPIResponseFormatOption1) ToPointer() *AssistantsAPIResponseFormatOption1 {
	return &e
}
func (e *AssistantsAPIResponseFormatOption1) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "auto":
		*e = AssistantsAPIResponseFormatOption1(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AssistantsAPIResponseFormatOption1: %v", v)
	}
}

type AssistantsAPIResponseFormatOptionType string

const (
	AssistantsAPIResponseFormatOptionTypeAssistantsAPIResponseFormatOption1 AssistantsAPIResponseFormatOptionType = "AssistantsApiResponseFormatOption_1"
	AssistantsAPIResponseFormatOptionTypeResponseFormatText                 AssistantsAPIResponseFormatOptionType = "ResponseFormatText"
	AssistantsAPIResponseFormatOptionTypeResponseFormatJSONObject           AssistantsAPIResponseFormatOptionType = "ResponseFormatJsonObject"
	AssistantsAPIResponseFormatOptionTypeResponseFormatJSONSchema           AssistantsAPIResponseFormatOptionType = "ResponseFormatJsonSchema"
)

// AssistantsAPIResponseFormatOption - Specifies the format that the model must output. Compatible with [GPT-4o](/docs/models#gpt-4o), [GPT-4 Turbo](/docs/models#gpt-4-turbo-and-gpt-4), and all GPT-3.5 Turbo models since `gpt-3.5-turbo-1106`.
//
// Setting to `{ "type": "json_schema", "json_schema": {...} }` enables Structured Outputs which ensures the model will match your supplied JSON schema. Learn more in the [Structured Outputs guide](/docs/guides/structured-outputs).
//
// Setting to `{ "type": "json_object" }` enables JSON mode, which ensures the message the model generates is valid JSON.
//
// **Important:** when using JSON mode, you **must** also instruct the model to produce JSON yourself via a system or user message. Without this, the model may generate an unending stream of whitespace until the generation reaches the token limit, resulting in a long-running and seemingly "stuck" request. Also note that the message content may be partially cut off if `finish_reason="length"`, which indicates the generation exceeded `max_tokens` or the conversation exceeded the max context length.
type AssistantsAPIResponseFormatOption struct {
	AssistantsAPIResponseFormatOption1 *AssistantsAPIResponseFormatOption1
	ResponseFormatText                 *ResponseFormatText
	ResponseFormatJSONObject           *ResponseFormatJSONObject
	ResponseFormatJSONSchema           *ResponseFormatJSONSchema

	Type AssistantsAPIResponseFormatOptionType
}

func CreateAssistantsAPIResponseFormatOptionAssistantsAPIResponseFormatOption1(assistantsAPIResponseFormatOption1 AssistantsAPIResponseFormatOption1) AssistantsAPIResponseFormatOption {
	typ := AssistantsAPIResponseFormatOptionTypeAssistantsAPIResponseFormatOption1

	return AssistantsAPIResponseFormatOption{
		AssistantsAPIResponseFormatOption1: &assistantsAPIResponseFormatOption1,
		Type:                               typ,
	}
}

func CreateAssistantsAPIResponseFormatOptionResponseFormatText(responseFormatText ResponseFormatText) AssistantsAPIResponseFormatOption {
	typ := AssistantsAPIResponseFormatOptionTypeResponseFormatText

	return AssistantsAPIResponseFormatOption{
		ResponseFormatText: &responseFormatText,
		Type:               typ,
	}
}

func CreateAssistantsAPIResponseFormatOptionResponseFormatJSONObject(responseFormatJSONObject ResponseFormatJSONObject) AssistantsAPIResponseFormatOption {
	typ := AssistantsAPIResponseFormatOptionTypeResponseFormatJSONObject

	return AssistantsAPIResponseFormatOption{
		ResponseFormatJSONObject: &responseFormatJSONObject,
		Type:                     typ,
	}
}

func CreateAssistantsAPIResponseFormatOptionResponseFormatJSONSchema(responseFormatJSONSchema ResponseFormatJSONSchema) AssistantsAPIResponseFormatOption {
	typ := AssistantsAPIResponseFormatOptionTypeResponseFormatJSONSchema

	return AssistantsAPIResponseFormatOption{
		ResponseFormatJSONSchema: &responseFormatJSONSchema,
		Type:                     typ,
	}
}

func (u *AssistantsAPIResponseFormatOption) UnmarshalJSON(data []byte) error {

	var responseFormatText ResponseFormatText = ResponseFormatText{}
	if err := utils.UnmarshalJSON(data, &responseFormatText, "", true, true); err == nil {
		u.ResponseFormatText = &responseFormatText
		u.Type = AssistantsAPIResponseFormatOptionTypeResponseFormatText
		return nil
	}

	var responseFormatJSONObject ResponseFormatJSONObject = ResponseFormatJSONObject{}
	if err := utils.UnmarshalJSON(data, &responseFormatJSONObject, "", true, true); err == nil {
		u.ResponseFormatJSONObject = &responseFormatJSONObject
		u.Type = AssistantsAPIResponseFormatOptionTypeResponseFormatJSONObject
		return nil
	}

	var responseFormatJSONSchema ResponseFormatJSONSchema = ResponseFormatJSONSchema{}
	if err := utils.UnmarshalJSON(data, &responseFormatJSONSchema, "", true, true); err == nil {
		u.ResponseFormatJSONSchema = &responseFormatJSONSchema
		u.Type = AssistantsAPIResponseFormatOptionTypeResponseFormatJSONSchema
		return nil
	}

	var assistantsAPIResponseFormatOption1 AssistantsAPIResponseFormatOption1 = AssistantsAPIResponseFormatOption1("")
	if err := utils.UnmarshalJSON(data, &assistantsAPIResponseFormatOption1, "", true, true); err == nil {
		u.AssistantsAPIResponseFormatOption1 = &assistantsAPIResponseFormatOption1
		u.Type = AssistantsAPIResponseFormatOptionTypeAssistantsAPIResponseFormatOption1
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for AssistantsAPIResponseFormatOption", string(data))
}

func (u AssistantsAPIResponseFormatOption) MarshalJSON() ([]byte, error) {
	if u.AssistantsAPIResponseFormatOption1 != nil {
		return utils.MarshalJSON(u.AssistantsAPIResponseFormatOption1, "", true)
	}

	if u.ResponseFormatText != nil {
		return utils.MarshalJSON(u.ResponseFormatText, "", true)
	}

	if u.ResponseFormatJSONObject != nil {
		return utils.MarshalJSON(u.ResponseFormatJSONObject, "", true)
	}

	if u.ResponseFormatJSONSchema != nil {
		return utils.MarshalJSON(u.ResponseFormatJSONSchema, "", true)
	}

	return nil, errors.New("could not marshal union type AssistantsAPIResponseFormatOption: all fields are null")
}
