// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jamelt/openai-api/internal/utils"
)

type CreateTranslationRequestFile struct {
	FileName string `multipartForm:"name=file"`
	// This field accepts []byte data or io.Reader implementations, such as *os.File.
	Content any `multipartForm:"content"`
}

func (o *CreateTranslationRequestFile) GetFileName() string {
	if o == nil {
		return ""
	}
	return o.FileName
}

func (o *CreateTranslationRequestFile) GetContent() any {
	if o == nil {
		return nil
	}
	return o.Content
}

type CreateTranslationRequestModel2 string

const (
	CreateTranslationRequestModel2Whisper1 CreateTranslationRequestModel2 = "whisper-1"
)

func (e CreateTranslationRequestModel2) ToPointer() *CreateTranslationRequestModel2 {
	return &e
}
func (e *CreateTranslationRequestModel2) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "whisper-1":
		*e = CreateTranslationRequestModel2(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateTranslationRequestModel2: %v", v)
	}
}

type CreateTranslationRequestModelType string

const (
	CreateTranslationRequestModelTypeStr                            CreateTranslationRequestModelType = "str"
	CreateTranslationRequestModelTypeCreateTranslationRequestModel2 CreateTranslationRequestModelType = "CreateTranslationRequest_model_2"
)

// CreateTranslationRequestModel - ID of the model to use. Only `whisper-1` (which is powered by our open source Whisper V2 model) is currently available.
type CreateTranslationRequestModel struct {
	Str                            *string
	CreateTranslationRequestModel2 *CreateTranslationRequestModel2

	Type CreateTranslationRequestModelType
}

func CreateCreateTranslationRequestModelStr(str string) CreateTranslationRequestModel {
	typ := CreateTranslationRequestModelTypeStr

	return CreateTranslationRequestModel{
		Str:  &str,
		Type: typ,
	}
}

func CreateCreateTranslationRequestModelCreateTranslationRequestModel2(createTranslationRequestModel2 CreateTranslationRequestModel2) CreateTranslationRequestModel {
	typ := CreateTranslationRequestModelTypeCreateTranslationRequestModel2

	return CreateTranslationRequestModel{
		CreateTranslationRequestModel2: &createTranslationRequestModel2,
		Type:                           typ,
	}
}

func (u *CreateTranslationRequestModel) UnmarshalJSON(data []byte) error {

	var str string = ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = CreateTranslationRequestModelTypeStr
		return nil
	}

	var createTranslationRequestModel2 CreateTranslationRequestModel2 = CreateTranslationRequestModel2("")
	if err := utils.UnmarshalJSON(data, &createTranslationRequestModel2, "", true, true); err == nil {
		u.CreateTranslationRequestModel2 = &createTranslationRequestModel2
		u.Type = CreateTranslationRequestModelTypeCreateTranslationRequestModel2
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for CreateTranslationRequestModel", string(data))
}

func (u CreateTranslationRequestModel) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.CreateTranslationRequestModel2 != nil {
		return utils.MarshalJSON(u.CreateTranslationRequestModel2, "", true)
	}

	return nil, errors.New("could not marshal union type CreateTranslationRequestModel: all fields are null")
}

type CreateTranslationRequest struct {
	// The audio file object (not file name) translate, in one of these formats: flac, mp3, mp4, mpeg, mpga, m4a, ogg, wav, or webm.
	//
	File CreateTranslationRequestFile `multipartForm:"file"`
	// ID of the model to use. Only `whisper-1` (which is powered by our open source Whisper V2 model) is currently available.
	//
	Model CreateTranslationRequestModel `multipartForm:"name=model"`
	// An optional text to guide the model's style or continue a previous audio segment. The [prompt](/docs/guides/speech-to-text#prompting) should be in English.
	//
	Prompt *string `multipartForm:"name=prompt"`
	// The format of the output, in one of these options: `json`, `text`, `srt`, `verbose_json`, or `vtt`.
	//
	ResponseFormat *AudioResponseFormat `default:"json" multipartForm:"name=response_format"`
	// The sampling temperature, between 0 and 1. Higher values like 0.8 will make the output more random, while lower values like 0.2 will make it more focused and deterministic. If set to 0, the model will use [log probability](https://en.wikipedia.org/wiki/Log_probability) to automatically increase the temperature until certain thresholds are hit.
	//
	Temperature *float64 `default:"0" multipartForm:"name=temperature"`
}

func (c CreateTranslationRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateTranslationRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateTranslationRequest) GetFile() CreateTranslationRequestFile {
	if o == nil {
		return CreateTranslationRequestFile{}
	}
	return o.File
}

func (o *CreateTranslationRequest) GetModel() CreateTranslationRequestModel {
	if o == nil {
		return CreateTranslationRequestModel{}
	}
	return o.Model
}

func (o *CreateTranslationRequest) GetPrompt() *string {
	if o == nil {
		return nil
	}
	return o.Prompt
}

func (o *CreateTranslationRequest) GetResponseFormat() *AudioResponseFormat {
	if o == nil {
		return nil
	}
	return o.ResponseFormat
}

func (o *CreateTranslationRequest) GetTemperature() *float64 {
	if o == nil {
		return nil
	}
	return o.Temperature
}
