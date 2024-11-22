// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jamelt/openai-api/internal/utils"
)

type CreateImageVariationRequestImage struct {
	FileName string `multipartForm:"name=image"`
	// This field accepts []byte data or io.Reader implementations, such as *os.File.
	Content any `multipartForm:"content"`
}

func (o *CreateImageVariationRequestImage) GetFileName() string {
	if o == nil {
		return ""
	}
	return o.FileName
}

func (o *CreateImageVariationRequestImage) GetContent() any {
	if o == nil {
		return nil
	}
	return o.Content
}

type CreateImageVariationRequestModel2 string

const (
	CreateImageVariationRequestModel2DallE2 CreateImageVariationRequestModel2 = "dall-e-2"
)

func (e CreateImageVariationRequestModel2) ToPointer() *CreateImageVariationRequestModel2 {
	return &e
}
func (e *CreateImageVariationRequestModel2) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "dall-e-2":
		*e = CreateImageVariationRequestModel2(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateImageVariationRequestModel2: %v", v)
	}
}

type CreateImageVariationRequestModelType string

const (
	CreateImageVariationRequestModelTypeStr                               CreateImageVariationRequestModelType = "str"
	CreateImageVariationRequestModelTypeCreateImageVariationRequestModel2 CreateImageVariationRequestModelType = "CreateImageVariationRequest_model_2"
)

// CreateImageVariationRequestModel - The model to use for image generation. Only `dall-e-2` is supported at this time.
type CreateImageVariationRequestModel struct {
	Str                               *string
	CreateImageVariationRequestModel2 *CreateImageVariationRequestModel2

	Type CreateImageVariationRequestModelType
}

func CreateCreateImageVariationRequestModelStr(str string) CreateImageVariationRequestModel {
	typ := CreateImageVariationRequestModelTypeStr

	return CreateImageVariationRequestModel{
		Str:  &str,
		Type: typ,
	}
}

func CreateCreateImageVariationRequestModelCreateImageVariationRequestModel2(createImageVariationRequestModel2 CreateImageVariationRequestModel2) CreateImageVariationRequestModel {
	typ := CreateImageVariationRequestModelTypeCreateImageVariationRequestModel2

	return CreateImageVariationRequestModel{
		CreateImageVariationRequestModel2: &createImageVariationRequestModel2,
		Type:                              typ,
	}
}

func (u *CreateImageVariationRequestModel) UnmarshalJSON(data []byte) error {

	var str string = ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = CreateImageVariationRequestModelTypeStr
		return nil
	}

	var createImageVariationRequestModel2 CreateImageVariationRequestModel2 = CreateImageVariationRequestModel2("")
	if err := utils.UnmarshalJSON(data, &createImageVariationRequestModel2, "", true, true); err == nil {
		u.CreateImageVariationRequestModel2 = &createImageVariationRequestModel2
		u.Type = CreateImageVariationRequestModelTypeCreateImageVariationRequestModel2
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for CreateImageVariationRequestModel", string(data))
}

func (u CreateImageVariationRequestModel) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.CreateImageVariationRequestModel2 != nil {
		return utils.MarshalJSON(u.CreateImageVariationRequestModel2, "", true)
	}

	return nil, errors.New("could not marshal union type CreateImageVariationRequestModel: all fields are null")
}

// CreateImageVariationRequestResponseFormat - The format in which the generated images are returned. Must be one of `url` or `b64_json`. URLs are only valid for 60 minutes after the image has been generated.
type CreateImageVariationRequestResponseFormat string

const (
	CreateImageVariationRequestResponseFormatURL     CreateImageVariationRequestResponseFormat = "url"
	CreateImageVariationRequestResponseFormatB64JSON CreateImageVariationRequestResponseFormat = "b64_json"
)

func (e CreateImageVariationRequestResponseFormat) ToPointer() *CreateImageVariationRequestResponseFormat {
	return &e
}
func (e *CreateImageVariationRequestResponseFormat) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "url":
		fallthrough
	case "b64_json":
		*e = CreateImageVariationRequestResponseFormat(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateImageVariationRequestResponseFormat: %v", v)
	}
}

// CreateImageVariationRequestSize - The size of the generated images. Must be one of `256x256`, `512x512`, or `1024x1024`.
type CreateImageVariationRequestSize string

const (
	CreateImageVariationRequestSizeTwoHundredAndFiftySixx256     CreateImageVariationRequestSize = "256x256"
	CreateImageVariationRequestSizeFiveHundredAndTwelvex512      CreateImageVariationRequestSize = "512x512"
	CreateImageVariationRequestSizeOneThousandAndTwentyFourx1024 CreateImageVariationRequestSize = "1024x1024"
)

func (e CreateImageVariationRequestSize) ToPointer() *CreateImageVariationRequestSize {
	return &e
}
func (e *CreateImageVariationRequestSize) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "256x256":
		fallthrough
	case "512x512":
		fallthrough
	case "1024x1024":
		*e = CreateImageVariationRequestSize(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateImageVariationRequestSize: %v", v)
	}
}

type CreateImageVariationRequest struct {
	// The image to use as the basis for the variation(s). Must be a valid PNG file, less than 4MB, and square.
	Image CreateImageVariationRequestImage `multipartForm:"file"`
	// The model to use for image generation. Only `dall-e-2` is supported at this time.
	Model *CreateImageVariationRequestModel `multipartForm:"name=model"`
	// The number of images to generate. Must be between 1 and 10. For `dall-e-3`, only `n=1` is supported.
	N *int64 `default:"1" multipartForm:"name=n"`
	// The format in which the generated images are returned. Must be one of `url` or `b64_json`. URLs are only valid for 60 minutes after the image has been generated.
	ResponseFormat *CreateImageVariationRequestResponseFormat `default:"url" multipartForm:"name=response_format"`
	// The size of the generated images. Must be one of `256x256`, `512x512`, or `1024x1024`.
	Size *CreateImageVariationRequestSize `default:"1024x1024" multipartForm:"name=size"`
	// A unique identifier representing your end-user, which can help OpenAI to monitor and detect abuse. [Learn more](/docs/guides/safety-best-practices#end-user-ids).
	//
	User *string `multipartForm:"name=user"`
}

func (c CreateImageVariationRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateImageVariationRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateImageVariationRequest) GetImage() CreateImageVariationRequestImage {
	if o == nil {
		return CreateImageVariationRequestImage{}
	}
	return o.Image
}

func (o *CreateImageVariationRequest) GetModel() *CreateImageVariationRequestModel {
	if o == nil {
		return nil
	}
	return o.Model
}

func (o *CreateImageVariationRequest) GetN() *int64 {
	if o == nil {
		return nil
	}
	return o.N
}

func (o *CreateImageVariationRequest) GetResponseFormat() *CreateImageVariationRequestResponseFormat {
	if o == nil {
		return nil
	}
	return o.ResponseFormat
}

func (o *CreateImageVariationRequest) GetSize() *CreateImageVariationRequestSize {
	if o == nil {
		return nil
	}
	return o.Size
}

func (o *CreateImageVariationRequest) GetUser() *string {
	if o == nil {
		return nil
	}
	return o.User
}
