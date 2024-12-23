// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
	"github.com/jamelt/openai-api/internal/utils"
)

// ChatCompletionRequestMessageContentPartImageType - The type of the content part.
type ChatCompletionRequestMessageContentPartImageType string

const (
	ChatCompletionRequestMessageContentPartImageTypeImageURL ChatCompletionRequestMessageContentPartImageType = "image_url"
)

func (e ChatCompletionRequestMessageContentPartImageType) ToPointer() *ChatCompletionRequestMessageContentPartImageType {
	return &e
}
func (e *ChatCompletionRequestMessageContentPartImageType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "image_url":
		*e = ChatCompletionRequestMessageContentPartImageType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ChatCompletionRequestMessageContentPartImageType: %v", v)
	}
}

// Detail - Specifies the detail level of the image. Learn more in the [Vision guide](/docs/guides/vision#low-or-high-fidelity-image-understanding).
type Detail string

const (
	DetailAuto Detail = "auto"
	DetailLow  Detail = "low"
	DetailHigh Detail = "high"
)

func (e Detail) ToPointer() *Detail {
	return &e
}
func (e *Detail) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "auto":
		fallthrough
	case "low":
		fallthrough
	case "high":
		*e = Detail(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Detail: %v", v)
	}
}

type ImageURL struct {
	// Either a URL of the image or the base64 encoded image data.
	URL string `json:"url"`
	// Specifies the detail level of the image. Learn more in the [Vision guide](/docs/guides/vision#low-or-high-fidelity-image-understanding).
	Detail *Detail `default:"auto" json:"detail"`
}

func (i ImageURL) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *ImageURL) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ImageURL) GetURL() string {
	if o == nil {
		return ""
	}
	return o.URL
}

func (o *ImageURL) GetDetail() *Detail {
	if o == nil {
		return nil
	}
	return o.Detail
}

// ChatCompletionRequestMessageContentPartImage - Learn about [image inputs](/docs/guides/vision).
type ChatCompletionRequestMessageContentPartImage struct {
	// The type of the content part.
	Type     ChatCompletionRequestMessageContentPartImageType `json:"type"`
	ImageURL ImageURL                                         `json:"image_url"`
}

func (o *ChatCompletionRequestMessageContentPartImage) GetType() ChatCompletionRequestMessageContentPartImageType {
	if o == nil {
		return ChatCompletionRequestMessageContentPartImageType("")
	}
	return o.Type
}

func (o *ChatCompletionRequestMessageContentPartImage) GetImageURL() ImageURL {
	if o == nil {
		return ImageURL{}
	}
	return o.ImageURL
}
