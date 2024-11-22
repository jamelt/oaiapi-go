// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jamelt/openai-api/internal/utils"
)

type InputType string

const (
	InputTypeStr                   InputType = "str"
	InputTypeArrayOfStr            InputType = "arrayOfStr"
	InputTypeArrayOfInteger        InputType = "arrayOfInteger"
	InputTypeArrayOfArrayOfInteger InputType = "arrayOfArrayOfInteger"
)

// Input text to embed, encoded as a string or array of tokens. To embed multiple inputs in a single request, pass an array of strings or array of token arrays. The input must not exceed the max input tokens for the model (8192 tokens for `text-embedding-ada-002`), cannot be an empty string, and any array must be 2048 dimensions or less. [Example Python code](https://cookbook.openai.com/examples/how_to_count_tokens_with_tiktoken) for counting tokens.
type Input struct {
	Str                   *string
	ArrayOfStr            []string
	ArrayOfInteger        []int64
	ArrayOfArrayOfInteger [][]int64

	Type InputType
}

func CreateInputStr(str string) Input {
	typ := InputTypeStr

	return Input{
		Str:  &str,
		Type: typ,
	}
}

func CreateInputArrayOfStr(arrayOfStr []string) Input {
	typ := InputTypeArrayOfStr

	return Input{
		ArrayOfStr: arrayOfStr,
		Type:       typ,
	}
}

func CreateInputArrayOfInteger(arrayOfInteger []int64) Input {
	typ := InputTypeArrayOfInteger

	return Input{
		ArrayOfInteger: arrayOfInteger,
		Type:           typ,
	}
}

func CreateInputArrayOfArrayOfInteger(arrayOfArrayOfInteger [][]int64) Input {
	typ := InputTypeArrayOfArrayOfInteger

	return Input{
		ArrayOfArrayOfInteger: arrayOfArrayOfInteger,
		Type:                  typ,
	}
}

func (u *Input) UnmarshalJSON(data []byte) error {

	var str string = ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = InputTypeStr
		return nil
	}

	var arrayOfStr []string = []string{}
	if err := utils.UnmarshalJSON(data, &arrayOfStr, "", true, true); err == nil {
		u.ArrayOfStr = arrayOfStr
		u.Type = InputTypeArrayOfStr
		return nil
	}

	var arrayOfInteger []int64 = []int64{}
	if err := utils.UnmarshalJSON(data, &arrayOfInteger, "", true, true); err == nil {
		u.ArrayOfInteger = arrayOfInteger
		u.Type = InputTypeArrayOfInteger
		return nil
	}

	var arrayOfArrayOfInteger [][]int64 = [][]int64{}
	if err := utils.UnmarshalJSON(data, &arrayOfArrayOfInteger, "", true, true); err == nil {
		u.ArrayOfArrayOfInteger = arrayOfArrayOfInteger
		u.Type = InputTypeArrayOfArrayOfInteger
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for Input", string(data))
}

func (u Input) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.ArrayOfStr != nil {
		return utils.MarshalJSON(u.ArrayOfStr, "", true)
	}

	if u.ArrayOfInteger != nil {
		return utils.MarshalJSON(u.ArrayOfInteger, "", true)
	}

	if u.ArrayOfArrayOfInteger != nil {
		return utils.MarshalJSON(u.ArrayOfArrayOfInteger, "", true)
	}

	return nil, errors.New("could not marshal union type Input: all fields are null")
}

type CreateEmbeddingRequestModel2 string

const (
	CreateEmbeddingRequestModel2TextEmbeddingAda002 CreateEmbeddingRequestModel2 = "text-embedding-ada-002"
	CreateEmbeddingRequestModel2TextEmbedding3Small CreateEmbeddingRequestModel2 = "text-embedding-3-small"
	CreateEmbeddingRequestModel2TextEmbedding3Large CreateEmbeddingRequestModel2 = "text-embedding-3-large"
)

func (e CreateEmbeddingRequestModel2) ToPointer() *CreateEmbeddingRequestModel2 {
	return &e
}
func (e *CreateEmbeddingRequestModel2) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "text-embedding-ada-002":
		fallthrough
	case "text-embedding-3-small":
		fallthrough
	case "text-embedding-3-large":
		*e = CreateEmbeddingRequestModel2(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateEmbeddingRequestModel2: %v", v)
	}
}

type CreateEmbeddingRequestModelType string

const (
	CreateEmbeddingRequestModelTypeStr                          CreateEmbeddingRequestModelType = "str"
	CreateEmbeddingRequestModelTypeCreateEmbeddingRequestModel2 CreateEmbeddingRequestModelType = "CreateEmbeddingRequest_model_2"
)

// CreateEmbeddingRequestModel - ID of the model to use. You can use the [List models](/docs/api-reference/models/list) API to see all of your available models, or see our [Model overview](/docs/models) for descriptions of them.
type CreateEmbeddingRequestModel struct {
	Str                          *string
	CreateEmbeddingRequestModel2 *CreateEmbeddingRequestModel2

	Type CreateEmbeddingRequestModelType
}

func CreateCreateEmbeddingRequestModelStr(str string) CreateEmbeddingRequestModel {
	typ := CreateEmbeddingRequestModelTypeStr

	return CreateEmbeddingRequestModel{
		Str:  &str,
		Type: typ,
	}
}

func CreateCreateEmbeddingRequestModelCreateEmbeddingRequestModel2(createEmbeddingRequestModel2 CreateEmbeddingRequestModel2) CreateEmbeddingRequestModel {
	typ := CreateEmbeddingRequestModelTypeCreateEmbeddingRequestModel2

	return CreateEmbeddingRequestModel{
		CreateEmbeddingRequestModel2: &createEmbeddingRequestModel2,
		Type:                         typ,
	}
}

func (u *CreateEmbeddingRequestModel) UnmarshalJSON(data []byte) error {

	var str string = ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = CreateEmbeddingRequestModelTypeStr
		return nil
	}

	var createEmbeddingRequestModel2 CreateEmbeddingRequestModel2 = CreateEmbeddingRequestModel2("")
	if err := utils.UnmarshalJSON(data, &createEmbeddingRequestModel2, "", true, true); err == nil {
		u.CreateEmbeddingRequestModel2 = &createEmbeddingRequestModel2
		u.Type = CreateEmbeddingRequestModelTypeCreateEmbeddingRequestModel2
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for CreateEmbeddingRequestModel", string(data))
}

func (u CreateEmbeddingRequestModel) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.CreateEmbeddingRequestModel2 != nil {
		return utils.MarshalJSON(u.CreateEmbeddingRequestModel2, "", true)
	}

	return nil, errors.New("could not marshal union type CreateEmbeddingRequestModel: all fields are null")
}

// EncodingFormat - The format to return the embeddings in. Can be either `float` or [`base64`](https://pypi.org/project/pybase64/).
type EncodingFormat string

const (
	EncodingFormatFloat  EncodingFormat = "float"
	EncodingFormatBase64 EncodingFormat = "base64"
)

func (e EncodingFormat) ToPointer() *EncodingFormat {
	return &e
}
func (e *EncodingFormat) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "float":
		fallthrough
	case "base64":
		*e = EncodingFormat(v)
		return nil
	default:
		return fmt.Errorf("invalid value for EncodingFormat: %v", v)
	}
}

type CreateEmbeddingRequest struct {
	// Input text to embed, encoded as a string or array of tokens. To embed multiple inputs in a single request, pass an array of strings or array of token arrays. The input must not exceed the max input tokens for the model (8192 tokens for `text-embedding-ada-002`), cannot be an empty string, and any array must be 2048 dimensions or less. [Example Python code](https://cookbook.openai.com/examples/how_to_count_tokens_with_tiktoken) for counting tokens.
	//
	Input Input `json:"input"`
	// ID of the model to use. You can use the [List models](/docs/api-reference/models/list) API to see all of your available models, or see our [Model overview](/docs/models) for descriptions of them.
	//
	Model CreateEmbeddingRequestModel `json:"model"`
	// The format to return the embeddings in. Can be either `float` or [`base64`](https://pypi.org/project/pybase64/).
	EncodingFormat *EncodingFormat `default:"float" json:"encoding_format"`
	// The number of dimensions the resulting output embeddings should have. Only supported in `text-embedding-3` and later models.
	//
	Dimensions *int64 `json:"dimensions,omitempty"`
	// A unique identifier representing your end-user, which can help OpenAI to monitor and detect abuse. [Learn more](/docs/guides/safety-best-practices#end-user-ids).
	//
	User *string `json:"user,omitempty"`
}

func (c CreateEmbeddingRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateEmbeddingRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateEmbeddingRequest) GetInput() Input {
	if o == nil {
		return Input{}
	}
	return o.Input
}

func (o *CreateEmbeddingRequest) GetModel() CreateEmbeddingRequestModel {
	if o == nil {
		return CreateEmbeddingRequestModel{}
	}
	return o.Model
}

func (o *CreateEmbeddingRequest) GetEncodingFormat() *EncodingFormat {
	if o == nil {
		return nil
	}
	return o.EncodingFormat
}

func (o *CreateEmbeddingRequest) GetDimensions() *int64 {
	if o == nil {
		return nil
	}
	return o.Dimensions
}

func (o *CreateEmbeddingRequest) GetUser() *string {
	if o == nil {
		return nil
	}
	return o.User
}
