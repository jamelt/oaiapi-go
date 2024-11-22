// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jamelt/openai-api/internal/utils"
)

// MessageContentTextObjectType - Always `text`.
type MessageContentTextObjectType string

const (
	MessageContentTextObjectTypeText MessageContentTextObjectType = "text"
)

func (e MessageContentTextObjectType) ToPointer() *MessageContentTextObjectType {
	return &e
}
func (e *MessageContentTextObjectType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "text":
		*e = MessageContentTextObjectType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MessageContentTextObjectType: %v", v)
	}
}

type AnnotationsType string

const (
	AnnotationsTypeMessageContentTextAnnotationsFileCitationObject AnnotationsType = "MessageContentTextAnnotationsFileCitationObject"
	AnnotationsTypeMessageContentTextAnnotationsFilePathObject     AnnotationsType = "MessageContentTextAnnotationsFilePathObject"
)

type Annotations struct {
	MessageContentTextAnnotationsFileCitationObject *MessageContentTextAnnotationsFileCitationObject
	MessageContentTextAnnotationsFilePathObject     *MessageContentTextAnnotationsFilePathObject

	Type AnnotationsType
}

func CreateAnnotationsMessageContentTextAnnotationsFileCitationObject(messageContentTextAnnotationsFileCitationObject MessageContentTextAnnotationsFileCitationObject) Annotations {
	typ := AnnotationsTypeMessageContentTextAnnotationsFileCitationObject

	return Annotations{
		MessageContentTextAnnotationsFileCitationObject: &messageContentTextAnnotationsFileCitationObject,
		Type: typ,
	}
}

func CreateAnnotationsMessageContentTextAnnotationsFilePathObject(messageContentTextAnnotationsFilePathObject MessageContentTextAnnotationsFilePathObject) Annotations {
	typ := AnnotationsTypeMessageContentTextAnnotationsFilePathObject

	return Annotations{
		MessageContentTextAnnotationsFilePathObject: &messageContentTextAnnotationsFilePathObject,
		Type: typ,
	}
}

func (u *Annotations) UnmarshalJSON(data []byte) error {

	var messageContentTextAnnotationsFileCitationObject MessageContentTextAnnotationsFileCitationObject = MessageContentTextAnnotationsFileCitationObject{}
	if err := utils.UnmarshalJSON(data, &messageContentTextAnnotationsFileCitationObject, "", true, true); err == nil {
		u.MessageContentTextAnnotationsFileCitationObject = &messageContentTextAnnotationsFileCitationObject
		u.Type = AnnotationsTypeMessageContentTextAnnotationsFileCitationObject
		return nil
	}

	var messageContentTextAnnotationsFilePathObject MessageContentTextAnnotationsFilePathObject = MessageContentTextAnnotationsFilePathObject{}
	if err := utils.UnmarshalJSON(data, &messageContentTextAnnotationsFilePathObject, "", true, true); err == nil {
		u.MessageContentTextAnnotationsFilePathObject = &messageContentTextAnnotationsFilePathObject
		u.Type = AnnotationsTypeMessageContentTextAnnotationsFilePathObject
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for Annotations", string(data))
}

func (u Annotations) MarshalJSON() ([]byte, error) {
	if u.MessageContentTextAnnotationsFileCitationObject != nil {
		return utils.MarshalJSON(u.MessageContentTextAnnotationsFileCitationObject, "", true)
	}

	if u.MessageContentTextAnnotationsFilePathObject != nil {
		return utils.MarshalJSON(u.MessageContentTextAnnotationsFilePathObject, "", true)
	}

	return nil, errors.New("could not marshal union type Annotations: all fields are null")
}

type Text struct {
	// The data that makes up the text.
	Value       string        `json:"value"`
	Annotations []Annotations `json:"annotations"`
}

func (o *Text) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

func (o *Text) GetAnnotations() []Annotations {
	if o == nil {
		return []Annotations{}
	}
	return o.Annotations
}

// MessageContentTextObject - The text content that is part of a message.
type MessageContentTextObject struct {
	// Always `text`.
	Type MessageContentTextObjectType `json:"type"`
	Text Text                         `json:"text"`
}

func (o *MessageContentTextObject) GetType() MessageContentTextObjectType {
	if o == nil {
		return MessageContentTextObjectType("")
	}
	return o.Type
}

func (o *MessageContentTextObject) GetText() Text {
	if o == nil {
		return Text{}
	}
	return o.Text
}