// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
	"github.com/jamelt/openai-api/internal/utils"
)

// MessageContentImageURLObjectType - The type of the content part.
type MessageContentImageURLObjectType string

const (
	MessageContentImageURLObjectTypeImageURL MessageContentImageURLObjectType = "image_url"
)

func (e MessageContentImageURLObjectType) ToPointer() *MessageContentImageURLObjectType {
	return &e
}
func (e *MessageContentImageURLObjectType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "image_url":
		*e = MessageContentImageURLObjectType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MessageContentImageURLObjectType: %v", v)
	}
}

// MessageContentImageURLObjectDetail - Specifies the detail level of the image. `low` uses fewer tokens, you can opt in to high resolution using `high`. Default value is `auto`
type MessageContentImageURLObjectDetail string

const (
	MessageContentImageURLObjectDetailAuto MessageContentImageURLObjectDetail = "auto"
	MessageContentImageURLObjectDetailLow  MessageContentImageURLObjectDetail = "low"
	MessageContentImageURLObjectDetailHigh MessageContentImageURLObjectDetail = "high"
)

func (e MessageContentImageURLObjectDetail) ToPointer() *MessageContentImageURLObjectDetail {
	return &e
}
func (e *MessageContentImageURLObjectDetail) UnmarshalJSON(data []byte) error {
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
		*e = MessageContentImageURLObjectDetail(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MessageContentImageURLObjectDetail: %v", v)
	}
}

type MessageContentImageURLObjectImageURL struct {
	// The external URL of the image, must be a supported image types: jpeg, jpg, png, gif, webp.
	URL string `json:"url"`
	// Specifies the detail level of the image. `low` uses fewer tokens, you can opt in to high resolution using `high`. Default value is `auto`
	Detail *MessageContentImageURLObjectDetail `default:"auto" json:"detail"`
}

func (m MessageContentImageURLObjectImageURL) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(m, "", false)
}

func (m *MessageContentImageURLObjectImageURL) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &m, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *MessageContentImageURLObjectImageURL) GetURL() string {
	if o == nil {
		return ""
	}
	return o.URL
}

func (o *MessageContentImageURLObjectImageURL) GetDetail() *MessageContentImageURLObjectDetail {
	if o == nil {
		return nil
	}
	return o.Detail
}

// MessageContentImageURLObject - References an image URL in the content of a message.
type MessageContentImageURLObject struct {
	// The type of the content part.
	Type     MessageContentImageURLObjectType     `json:"type"`
	ImageURL MessageContentImageURLObjectImageURL `json:"image_url"`
}

func (o *MessageContentImageURLObject) GetType() MessageContentImageURLObjectType {
	if o == nil {
		return MessageContentImageURLObjectType("")
	}
	return o.Type
}

func (o *MessageContentImageURLObject) GetImageURL() MessageContentImageURLObjectImageURL {
	if o == nil {
		return MessageContentImageURLObjectImageURL{}
	}
	return o.ImageURL
}
