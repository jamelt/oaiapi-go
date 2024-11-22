// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jamelt/openai-api/internal/utils"
)

// ChatCompletionToolChoiceOption1 - `none` means the model will not call any tool and instead generates a message. `auto` means the model can pick between generating a message or calling one or more tools. `required` means the model must call one or more tools.
type ChatCompletionToolChoiceOption1 string

const (
	ChatCompletionToolChoiceOption1None     ChatCompletionToolChoiceOption1 = "none"
	ChatCompletionToolChoiceOption1Auto     ChatCompletionToolChoiceOption1 = "auto"
	ChatCompletionToolChoiceOption1Required ChatCompletionToolChoiceOption1 = "required"
)

func (e ChatCompletionToolChoiceOption1) ToPointer() *ChatCompletionToolChoiceOption1 {
	return &e
}
func (e *ChatCompletionToolChoiceOption1) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "none":
		fallthrough
	case "auto":
		fallthrough
	case "required":
		*e = ChatCompletionToolChoiceOption1(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ChatCompletionToolChoiceOption1: %v", v)
	}
}

type ChatCompletionToolChoiceOptionType string

const (
	ChatCompletionToolChoiceOptionTypeChatCompletionToolChoiceOption1 ChatCompletionToolChoiceOptionType = "ChatCompletionToolChoiceOption_1"
	ChatCompletionToolChoiceOptionTypeChatCompletionNamedToolChoice   ChatCompletionToolChoiceOptionType = "ChatCompletionNamedToolChoice"
)

// ChatCompletionToolChoiceOption - Controls which (if any) tool is called by the model.
// `none` means the model will not call any tool and instead generates a message.
// `auto` means the model can pick between generating a message or calling one or more tools.
// `required` means the model must call one or more tools.
// Specifying a particular tool via `{"type": "function", "function": {"name": "my_function"}}` forces the model to call that tool.
//
// `none` is the default when no tools are present. `auto` is the default if tools are present.
type ChatCompletionToolChoiceOption struct {
	ChatCompletionToolChoiceOption1 *ChatCompletionToolChoiceOption1
	ChatCompletionNamedToolChoice   *ChatCompletionNamedToolChoice

	Type ChatCompletionToolChoiceOptionType
}

func CreateChatCompletionToolChoiceOptionChatCompletionToolChoiceOption1(chatCompletionToolChoiceOption1 ChatCompletionToolChoiceOption1) ChatCompletionToolChoiceOption {
	typ := ChatCompletionToolChoiceOptionTypeChatCompletionToolChoiceOption1

	return ChatCompletionToolChoiceOption{
		ChatCompletionToolChoiceOption1: &chatCompletionToolChoiceOption1,
		Type:                            typ,
	}
}

func CreateChatCompletionToolChoiceOptionChatCompletionNamedToolChoice(chatCompletionNamedToolChoice ChatCompletionNamedToolChoice) ChatCompletionToolChoiceOption {
	typ := ChatCompletionToolChoiceOptionTypeChatCompletionNamedToolChoice

	return ChatCompletionToolChoiceOption{
		ChatCompletionNamedToolChoice: &chatCompletionNamedToolChoice,
		Type:                          typ,
	}
}

func (u *ChatCompletionToolChoiceOption) UnmarshalJSON(data []byte) error {

	var chatCompletionNamedToolChoice ChatCompletionNamedToolChoice = ChatCompletionNamedToolChoice{}
	if err := utils.UnmarshalJSON(data, &chatCompletionNamedToolChoice, "", true, true); err == nil {
		u.ChatCompletionNamedToolChoice = &chatCompletionNamedToolChoice
		u.Type = ChatCompletionToolChoiceOptionTypeChatCompletionNamedToolChoice
		return nil
	}

	var chatCompletionToolChoiceOption1 ChatCompletionToolChoiceOption1 = ChatCompletionToolChoiceOption1("")
	if err := utils.UnmarshalJSON(data, &chatCompletionToolChoiceOption1, "", true, true); err == nil {
		u.ChatCompletionToolChoiceOption1 = &chatCompletionToolChoiceOption1
		u.Type = ChatCompletionToolChoiceOptionTypeChatCompletionToolChoiceOption1
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for ChatCompletionToolChoiceOption", string(data))
}

func (u ChatCompletionToolChoiceOption) MarshalJSON() ([]byte, error) {
	if u.ChatCompletionToolChoiceOption1 != nil {
		return utils.MarshalJSON(u.ChatCompletionToolChoiceOption1, "", true)
	}

	if u.ChatCompletionNamedToolChoice != nil {
		return utils.MarshalJSON(u.ChatCompletionNamedToolChoice, "", true)
	}

	return nil, errors.New("could not marshal union type ChatCompletionToolChoiceOption: all fields are null")
}