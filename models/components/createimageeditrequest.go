// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jamelt/openai-api/internal/utils"
)

type CreateImageEditRequestImage struct {
	FileName string `multipartForm:"name=image"`
	// This field accepts []byte data or io.Reader implementations, such as *os.File.
	Content any `multipartForm:"content"`
}

func (o *CreateImageEditRequestImage) GetFileName() string {
	if o == nil {
		return ""
	}
	return o.FileName
}

func (o *CreateImageEditRequestImage) GetContent() any {
	if o == nil {
		return nil
	}
	return o.Content
}

type Mask struct {
	FileName string `multipartForm:"name=mask"`
	// This field accepts []byte data or io.Reader implementations, such as *os.File.
	Content any `multipartForm:"content"`
}

func (o *Mask) GetFileName() string {
	if o == nil {
		return ""
	}
	return o.FileName
}

func (o *Mask) GetContent() any {
	if o == nil {
		return nil
	}
	return o.Content
}

type CreateImageEditRequestModel2 string

const (
	CreateImageEditRequestModel2DallE2 CreateImageEditRequestModel2 = "dall-e-2"
)

func (e CreateImageEditRequestModel2) ToPointer() *CreateImageEditRequestModel2 {
	return &e
}
func (e *CreateImageEditRequestModel2) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "dall-e-2":
		*e = CreateImageEditRequestModel2(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateImageEditRequestModel2: %v", v)
	}
}

type CreateImageEditRequestModelType string

const (
	CreateImageEditRequestModelTypeStr                          CreateImageEditRequestModelType = "str"
	CreateImageEditRequestModelTypeCreateImageEditRequestModel2 CreateImageEditRequestModelType = "CreateImageEditRequest_model_2"
)

// CreateImageEditRequestModel - The model to use for image generation. Only `dall-e-2` is supported at this time.
type CreateImageEditRequestModel struct {
	Str                          *string
	CreateImageEditRequestModel2 *CreateImageEditRequestModel2

	Type CreateImageEditRequestModelType
}

func CreateCreateImageEditRequestModelStr(str string) CreateImageEditRequestModel {
	typ := CreateImageEditRequestModelTypeStr

	return CreateImageEditRequestModel{
		Str:  &str,
		Type: typ,
	}
}

func CreateCreateImageEditRequestModelCreateImageEditRequestModel2(createImageEditRequestModel2 CreateImageEditRequestModel2) CreateImageEditRequestModel {
	typ := CreateImageEditRequestModelTypeCreateImageEditRequestModel2

	return CreateImageEditRequestModel{
		CreateImageEditRequestModel2: &createImageEditRequestModel2,
		Type:                         typ,
	}
}

func (u *CreateImageEditRequestModel) UnmarshalJSON(data []byte) error {

	var str string = ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = CreateImageEditRequestModelTypeStr
		return nil
	}

	var createImageEditRequestModel2 CreateImageEditRequestModel2 = CreateImageEditRequestModel2("")
	if err := utils.UnmarshalJSON(data, &createImageEditRequestModel2, "", true, true); err == nil {
		u.CreateImageEditRequestModel2 = &createImageEditRequestModel2
		u.Type = CreateImageEditRequestModelTypeCreateImageEditRequestModel2
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for CreateImageEditRequestModel", string(data))
}

func (u CreateImageEditRequestModel) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.CreateImageEditRequestModel2 != nil {
		return utils.MarshalJSON(u.CreateImageEditRequestModel2, "", true)
	}

	return nil, errors.New("could not marshal union type CreateImageEditRequestModel: all fields are null")
}

// Size - The size of the generated images. Must be one of `256x256`, `512x512`, or `1024x1024`.
type Size string

const (
	SizeTwoHundredAndFiftySixx256     Size = "256x256"
	SizeFiveHundredAndTwelvex512      Size = "512x512"
	SizeOneThousandAndTwentyFourx1024 Size = "1024x1024"
)

func (e Size) ToPointer() *Size {
	return &e
}
func (e *Size) UnmarshalJSON(data []byte) error {
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
		*e = Size(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Size: %v", v)
	}
}

// CreateImageEditRequestResponseFormat - The format in which the generated images are returned. Must be one of `url` or `b64_json`. URLs are only valid for 60 minutes after the image has been generated.
type CreateImageEditRequestResponseFormat string

const (
	CreateImageEditRequestResponseFormatURL     CreateImageEditRequestResponseFormat = "url"
	CreateImageEditRequestResponseFormatB64JSON CreateImageEditRequestResponseFormat = "b64_json"
)

func (e CreateImageEditRequestResponseFormat) ToPointer() *CreateImageEditRequestResponseFormat {
	return &e
}
func (e *CreateImageEditRequestResponseFormat) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "url":
		fallthrough
	case "b64_json":
		*e = CreateImageEditRequestResponseFormat(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateImageEditRequestResponseFormat: %v", v)
	}
}

type CreateImageEditRequest struct {
	// The image to edit. Must be a valid PNG file, less than 4MB, and square. If mask is not provided, image must have transparency, which will be used as the mask.
	Image CreateImageEditRequestImage `multipartForm:"file"`
	// A text description of the desired image(s). The maximum length is 1000 characters.
	Prompt string `multipartForm:"name=prompt"`
	// An additional image whose fully transparent areas (e.g. where alpha is zero) indicate where `image` should be edited. Must be a valid PNG file, less than 4MB, and have the same dimensions as `image`.
	Mask *Mask `multipartForm:"file"`
	// The model to use for image generation. Only `dall-e-2` is supported at this time.
	Model *CreateImageEditRequestModel `multipartForm:"name=model"`
	// The number of images to generate. Must be between 1 and 10.
	N *int64 `default:"1" multipartForm:"name=n"`
	// The size of the generated images. Must be one of `256x256`, `512x512`, or `1024x1024`.
	Size *Size `default:"1024x1024" multipartForm:"name=size"`
	// The format in which the generated images are returned. Must be one of `url` or `b64_json`. URLs are only valid for 60 minutes after the image has been generated.
	ResponseFormat *CreateImageEditRequestResponseFormat `default:"url" multipartForm:"name=response_format"`
	// A unique identifier representing your end-user, which can help OpenAI to monitor and detect abuse. [Learn more](/docs/guides/safety-best-practices#end-user-ids).
	//
	User *string `multipartForm:"name=user"`
}

func (c CreateImageEditRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateImageEditRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateImageEditRequest) GetImage() CreateImageEditRequestImage {
	if o == nil {
		return CreateImageEditRequestImage{}
	}
	return o.Image
}

func (o *CreateImageEditRequest) GetPrompt() string {
	if o == nil {
		return ""
	}
	return o.Prompt
}

func (o *CreateImageEditRequest) GetMask() *Mask {
	if o == nil {
		return nil
	}
	return o.Mask
}

func (o *CreateImageEditRequest) GetModel() *CreateImageEditRequestModel {
	if o == nil {
		return nil
	}
	return o.Model
}

func (o *CreateImageEditRequest) GetN() *int64 {
	if o == nil {
		return nil
	}
	return o.N
}

func (o *CreateImageEditRequest) GetSize() *Size {
	if o == nil {
		return nil
	}
	return o.Size
}

func (o *CreateImageEditRequest) GetResponseFormat() *CreateImageEditRequestResponseFormat {
	if o == nil {
		return nil
	}
	return o.ResponseFormat
}

func (o *CreateImageEditRequest) GetUser() *string {
	if o == nil {
		return nil
	}
	return o.User
}
