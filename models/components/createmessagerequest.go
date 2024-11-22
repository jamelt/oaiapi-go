// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jamelt/openai-api/internal/utils"
)

// CreateMessageRequestRole - The role of the entity that is creating the message. Allowed values include:
// - `user`: Indicates the message is sent by an actual user and should be used in most cases to represent user-generated messages.
// - `assistant`: Indicates the message is generated by the assistant. Use this value to insert messages from the assistant into the conversation.
type CreateMessageRequestRole string

const (
	CreateMessageRequestRoleUser      CreateMessageRequestRole = "user"
	CreateMessageRequestRoleAssistant CreateMessageRequestRole = "assistant"
)

func (e CreateMessageRequestRole) ToPointer() *CreateMessageRequestRole {
	return &e
}
func (e *CreateMessageRequestRole) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "user":
		fallthrough
	case "assistant":
		*e = CreateMessageRequestRole(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateMessageRequestRole: %v", v)
	}
}

type Content2Type string

const (
	Content2TypeMessageContentImageFileObject   Content2Type = "MessageContentImageFileObject"
	Content2TypeMessageContentImageURLObject    Content2Type = "MessageContentImageUrlObject"
	Content2TypeMessageRequestContentTextObject Content2Type = "MessageRequestContentTextObject"
)

type Content2 struct {
	MessageContentImageFileObject   *MessageContentImageFileObject
	MessageContentImageURLObject    *MessageContentImageURLObject
	MessageRequestContentTextObject *MessageRequestContentTextObject

	Type Content2Type
}

func CreateContent2MessageContentImageFileObject(messageContentImageFileObject MessageContentImageFileObject) Content2 {
	typ := Content2TypeMessageContentImageFileObject

	return Content2{
		MessageContentImageFileObject: &messageContentImageFileObject,
		Type:                          typ,
	}
}

func CreateContent2MessageContentImageURLObject(messageContentImageURLObject MessageContentImageURLObject) Content2 {
	typ := Content2TypeMessageContentImageURLObject

	return Content2{
		MessageContentImageURLObject: &messageContentImageURLObject,
		Type:                         typ,
	}
}

func CreateContent2MessageRequestContentTextObject(messageRequestContentTextObject MessageRequestContentTextObject) Content2 {
	typ := Content2TypeMessageRequestContentTextObject

	return Content2{
		MessageRequestContentTextObject: &messageRequestContentTextObject,
		Type:                            typ,
	}
}

func (u *Content2) UnmarshalJSON(data []byte) error {

	var messageContentImageFileObject MessageContentImageFileObject = MessageContentImageFileObject{}
	if err := utils.UnmarshalJSON(data, &messageContentImageFileObject, "", true, true); err == nil {
		u.MessageContentImageFileObject = &messageContentImageFileObject
		u.Type = Content2TypeMessageContentImageFileObject
		return nil
	}

	var messageContentImageURLObject MessageContentImageURLObject = MessageContentImageURLObject{}
	if err := utils.UnmarshalJSON(data, &messageContentImageURLObject, "", true, true); err == nil {
		u.MessageContentImageURLObject = &messageContentImageURLObject
		u.Type = Content2TypeMessageContentImageURLObject
		return nil
	}

	var messageRequestContentTextObject MessageRequestContentTextObject = MessageRequestContentTextObject{}
	if err := utils.UnmarshalJSON(data, &messageRequestContentTextObject, "", true, true); err == nil {
		u.MessageRequestContentTextObject = &messageRequestContentTextObject
		u.Type = Content2TypeMessageRequestContentTextObject
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for Content2", string(data))
}

func (u Content2) MarshalJSON() ([]byte, error) {
	if u.MessageContentImageFileObject != nil {
		return utils.MarshalJSON(u.MessageContentImageFileObject, "", true)
	}

	if u.MessageContentImageURLObject != nil {
		return utils.MarshalJSON(u.MessageContentImageURLObject, "", true)
	}

	if u.MessageRequestContentTextObject != nil {
		return utils.MarshalJSON(u.MessageRequestContentTextObject, "", true)
	}

	return nil, errors.New("could not marshal union type Content2: all fields are null")
}

type CreateMessageRequestContentType string

const (
	CreateMessageRequestContentTypeStr             CreateMessageRequestContentType = "str"
	CreateMessageRequestContentTypeArrayOfContent2 CreateMessageRequestContentType = "arrayOfContent2"
)

type CreateMessageRequestContent struct {
	Str             *string
	ArrayOfContent2 []Content2

	Type CreateMessageRequestContentType
}

func CreateCreateMessageRequestContentStr(str string) CreateMessageRequestContent {
	typ := CreateMessageRequestContentTypeStr

	return CreateMessageRequestContent{
		Str:  &str,
		Type: typ,
	}
}

func CreateCreateMessageRequestContentArrayOfContent2(arrayOfContent2 []Content2) CreateMessageRequestContent {
	typ := CreateMessageRequestContentTypeArrayOfContent2

	return CreateMessageRequestContent{
		ArrayOfContent2: arrayOfContent2,
		Type:            typ,
	}
}

func (u *CreateMessageRequestContent) UnmarshalJSON(data []byte) error {

	var str string = ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = CreateMessageRequestContentTypeStr
		return nil
	}

	var arrayOfContent2 []Content2 = []Content2{}
	if err := utils.UnmarshalJSON(data, &arrayOfContent2, "", true, true); err == nil {
		u.ArrayOfContent2 = arrayOfContent2
		u.Type = CreateMessageRequestContentTypeArrayOfContent2
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for CreateMessageRequestContent", string(data))
}

func (u CreateMessageRequestContent) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.ArrayOfContent2 != nil {
		return utils.MarshalJSON(u.ArrayOfContent2, "", true)
	}

	return nil, errors.New("could not marshal union type CreateMessageRequestContent: all fields are null")
}

type CreateMessageRequestToolsType string

const (
	CreateMessageRequestToolsTypeAssistantToolsCode               CreateMessageRequestToolsType = "AssistantToolsCode"
	CreateMessageRequestToolsTypeAssistantToolsFileSearchTypeOnly CreateMessageRequestToolsType = "AssistantToolsFileSearchTypeOnly"
)

type CreateMessageRequestTools struct {
	AssistantToolsCode               *AssistantToolsCode
	AssistantToolsFileSearchTypeOnly *AssistantToolsFileSearchTypeOnly

	Type CreateMessageRequestToolsType
}

func CreateCreateMessageRequestToolsAssistantToolsCode(assistantToolsCode AssistantToolsCode) CreateMessageRequestTools {
	typ := CreateMessageRequestToolsTypeAssistantToolsCode

	return CreateMessageRequestTools{
		AssistantToolsCode: &assistantToolsCode,
		Type:               typ,
	}
}

func CreateCreateMessageRequestToolsAssistantToolsFileSearchTypeOnly(assistantToolsFileSearchTypeOnly AssistantToolsFileSearchTypeOnly) CreateMessageRequestTools {
	typ := CreateMessageRequestToolsTypeAssistantToolsFileSearchTypeOnly

	return CreateMessageRequestTools{
		AssistantToolsFileSearchTypeOnly: &assistantToolsFileSearchTypeOnly,
		Type:                             typ,
	}
}

func (u *CreateMessageRequestTools) UnmarshalJSON(data []byte) error {

	var assistantToolsCode AssistantToolsCode = AssistantToolsCode{}
	if err := utils.UnmarshalJSON(data, &assistantToolsCode, "", true, true); err == nil {
		u.AssistantToolsCode = &assistantToolsCode
		u.Type = CreateMessageRequestToolsTypeAssistantToolsCode
		return nil
	}

	var assistantToolsFileSearchTypeOnly AssistantToolsFileSearchTypeOnly = AssistantToolsFileSearchTypeOnly{}
	if err := utils.UnmarshalJSON(data, &assistantToolsFileSearchTypeOnly, "", true, true); err == nil {
		u.AssistantToolsFileSearchTypeOnly = &assistantToolsFileSearchTypeOnly
		u.Type = CreateMessageRequestToolsTypeAssistantToolsFileSearchTypeOnly
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for CreateMessageRequestTools", string(data))
}

func (u CreateMessageRequestTools) MarshalJSON() ([]byte, error) {
	if u.AssistantToolsCode != nil {
		return utils.MarshalJSON(u.AssistantToolsCode, "", true)
	}

	if u.AssistantToolsFileSearchTypeOnly != nil {
		return utils.MarshalJSON(u.AssistantToolsFileSearchTypeOnly, "", true)
	}

	return nil, errors.New("could not marshal union type CreateMessageRequestTools: all fields are null")
}

type Attachments struct {
	// The ID of the file to attach to the message.
	FileID *string `json:"file_id,omitempty"`
	// The tools to add this file to.
	Tools []CreateMessageRequestTools `json:"tools,omitempty"`
}

func (o *Attachments) GetFileID() *string {
	if o == nil {
		return nil
	}
	return o.FileID
}

func (o *Attachments) GetTools() []CreateMessageRequestTools {
	if o == nil {
		return nil
	}
	return o.Tools
}

// CreateMessageRequestMetadata - Set of 16 key-value pairs that can be attached to an object. This can be useful for storing additional information about the object in a structured format. Keys can be a maximum of 64 characters long and values can be a maximum of 512 characters long.
type CreateMessageRequestMetadata struct {
}

type CreateMessageRequest struct {
	// The role of the entity that is creating the message. Allowed values include:
	// - `user`: Indicates the message is sent by an actual user and should be used in most cases to represent user-generated messages.
	// - `assistant`: Indicates the message is generated by the assistant. Use this value to insert messages from the assistant into the conversation.
	//
	Role    CreateMessageRequestRole    `json:"role"`
	Content CreateMessageRequestContent `json:"content"`
	// A list of files attached to the message, and the tools they should be added to.
	Attachments []Attachments `json:"attachments,omitempty"`
	// Set of 16 key-value pairs that can be attached to an object. This can be useful for storing additional information about the object in a structured format. Keys can be a maximum of 64 characters long and values can be a maximum of 512 characters long.
	//
	Metadata *CreateMessageRequestMetadata `json:"metadata,omitempty"`
}

func (o *CreateMessageRequest) GetRole() CreateMessageRequestRole {
	if o == nil {
		return CreateMessageRequestRole("")
	}
	return o.Role
}

func (o *CreateMessageRequest) GetContent() CreateMessageRequestContent {
	if o == nil {
		return CreateMessageRequestContent{}
	}
	return o.Content
}

func (o *CreateMessageRequest) GetAttachments() []Attachments {
	if o == nil {
		return nil
	}
	return o.Attachments
}

func (o *CreateMessageRequest) GetMetadata() *CreateMessageRequestMetadata {
	if o == nil {
		return nil
	}
	return o.Metadata
}