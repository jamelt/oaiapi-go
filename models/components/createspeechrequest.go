// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jamelt/openai-api/internal/utils"
)

type Model2 string

const (
	Model2Tts1   Model2 = "tts-1"
	Model2Tts1Hd Model2 = "tts-1-hd"
)

func (e Model2) ToPointer() *Model2 {
	return &e
}
func (e *Model2) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "tts-1":
		fallthrough
	case "tts-1-hd":
		*e = Model2(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Model2: %v", v)
	}
}

type CreateSpeechRequestModelType string

const (
	CreateSpeechRequestModelTypeStr    CreateSpeechRequestModelType = "str"
	CreateSpeechRequestModelTypeModel2 CreateSpeechRequestModelType = "model_2"
)

// CreateSpeechRequestModel - One of the available [TTS models](/docs/models#tts): `tts-1` or `tts-1-hd`
type CreateSpeechRequestModel struct {
	Str    *string
	Model2 *Model2

	Type CreateSpeechRequestModelType
}

func CreateCreateSpeechRequestModelStr(str string) CreateSpeechRequestModel {
	typ := CreateSpeechRequestModelTypeStr

	return CreateSpeechRequestModel{
		Str:  &str,
		Type: typ,
	}
}

func CreateCreateSpeechRequestModelModel2(model2 Model2) CreateSpeechRequestModel {
	typ := CreateSpeechRequestModelTypeModel2

	return CreateSpeechRequestModel{
		Model2: &model2,
		Type:   typ,
	}
}

func (u *CreateSpeechRequestModel) UnmarshalJSON(data []byte) error {

	var str string = ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = CreateSpeechRequestModelTypeStr
		return nil
	}

	var model2 Model2 = Model2("")
	if err := utils.UnmarshalJSON(data, &model2, "", true, true); err == nil {
		u.Model2 = &model2
		u.Type = CreateSpeechRequestModelTypeModel2
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for CreateSpeechRequestModel", string(data))
}

func (u CreateSpeechRequestModel) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.Model2 != nil {
		return utils.MarshalJSON(u.Model2, "", true)
	}

	return nil, errors.New("could not marshal union type CreateSpeechRequestModel: all fields are null")
}

// Voice - The voice to use when generating the audio. Supported voices are `alloy`, `echo`, `fable`, `onyx`, `nova`, and `shimmer`. Previews of the voices are available in the [Text to speech guide](/docs/guides/text-to-speech#voice-options).
type Voice string

const (
	VoiceAlloy   Voice = "alloy"
	VoiceEcho    Voice = "echo"
	VoiceFable   Voice = "fable"
	VoiceOnyx    Voice = "onyx"
	VoiceNova    Voice = "nova"
	VoiceShimmer Voice = "shimmer"
)

func (e Voice) ToPointer() *Voice {
	return &e
}
func (e *Voice) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "alloy":
		fallthrough
	case "echo":
		fallthrough
	case "fable":
		fallthrough
	case "onyx":
		fallthrough
	case "nova":
		fallthrough
	case "shimmer":
		*e = Voice(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Voice: %v", v)
	}
}

// ResponseFormat - The format to audio in. Supported formats are `mp3`, `opus`, `aac`, `flac`, `wav`, and `pcm`.
type ResponseFormat string

const (
	ResponseFormatMp3  ResponseFormat = "mp3"
	ResponseFormatOpus ResponseFormat = "opus"
	ResponseFormatAac  ResponseFormat = "aac"
	ResponseFormatFlac ResponseFormat = "flac"
	ResponseFormatWav  ResponseFormat = "wav"
	ResponseFormatPcm  ResponseFormat = "pcm"
)

func (e ResponseFormat) ToPointer() *ResponseFormat {
	return &e
}
func (e *ResponseFormat) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "mp3":
		fallthrough
	case "opus":
		fallthrough
	case "aac":
		fallthrough
	case "flac":
		fallthrough
	case "wav":
		fallthrough
	case "pcm":
		*e = ResponseFormat(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ResponseFormat: %v", v)
	}
}

type CreateSpeechRequest struct {
	// One of the available [TTS models](/docs/models#tts): `tts-1` or `tts-1-hd`
	//
	Model CreateSpeechRequestModel `json:"model"`
	// The text to generate audio for. The maximum length is 4096 characters.
	Input string `json:"input"`
	// The voice to use when generating the audio. Supported voices are `alloy`, `echo`, `fable`, `onyx`, `nova`, and `shimmer`. Previews of the voices are available in the [Text to speech guide](/docs/guides/text-to-speech#voice-options).
	Voice Voice `json:"voice"`
	// The format to audio in. Supported formats are `mp3`, `opus`, `aac`, `flac`, `wav`, and `pcm`.
	ResponseFormat *ResponseFormat `default:"mp3" json:"response_format"`
	// The speed of the generated audio. Select a value from `0.25` to `4.0`. `1.0` is the default.
	Speed *float64 `default:"1" json:"speed"`
}

func (c CreateSpeechRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateSpeechRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateSpeechRequest) GetModel() CreateSpeechRequestModel {
	if o == nil {
		return CreateSpeechRequestModel{}
	}
	return o.Model
}

func (o *CreateSpeechRequest) GetInput() string {
	if o == nil {
		return ""
	}
	return o.Input
}

func (o *CreateSpeechRequest) GetVoice() Voice {
	if o == nil {
		return Voice("")
	}
	return o.Voice
}

func (o *CreateSpeechRequest) GetResponseFormat() *ResponseFormat {
	if o == nil {
		return nil
	}
	return o.ResponseFormat
}

func (o *CreateSpeechRequest) GetSpeed() *float64 {
	if o == nil {
		return nil
	}
	return o.Speed
}