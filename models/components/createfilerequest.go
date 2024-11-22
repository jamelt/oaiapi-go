// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

type CreateFileRequestFile struct {
	FileName string `multipartForm:"name=file"`
	// This field accepts []byte data or io.Reader implementations, such as *os.File.
	Content any `multipartForm:"content"`
}

func (o *CreateFileRequestFile) GetFileName() string {
	if o == nil {
		return ""
	}
	return o.FileName
}

func (o *CreateFileRequestFile) GetContent() any {
	if o == nil {
		return nil
	}
	return o.Content
}

// CreateFileRequestPurpose - The intended purpose of the uploaded file.
//
// Use "assistants" for [Assistants](/docs/api-reference/assistants) and [Message](/docs/api-reference/messages) files, "vision" for Assistants image file inputs, "batch" for [Batch API](/docs/guides/batch), and "fine-tune" for [Fine-tuning](/docs/api-reference/fine-tuning).
type CreateFileRequestPurpose string

const (
	CreateFileRequestPurposeAssistants CreateFileRequestPurpose = "assistants"
	CreateFileRequestPurposeBatch      CreateFileRequestPurpose = "batch"
	CreateFileRequestPurposeFineTune   CreateFileRequestPurpose = "fine-tune"
	CreateFileRequestPurposeVision     CreateFileRequestPurpose = "vision"
)

func (e CreateFileRequestPurpose) ToPointer() *CreateFileRequestPurpose {
	return &e
}
func (e *CreateFileRequestPurpose) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "assistants":
		fallthrough
	case "batch":
		fallthrough
	case "fine-tune":
		fallthrough
	case "vision":
		*e = CreateFileRequestPurpose(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateFileRequestPurpose: %v", v)
	}
}

type CreateFileRequest struct {
	// The File object (not file name) to be uploaded.
	//
	File CreateFileRequestFile `multipartForm:"file"`
	// The intended purpose of the uploaded file.
	//
	// Use "assistants" for [Assistants](/docs/api-reference/assistants) and [Message](/docs/api-reference/messages) files, "vision" for Assistants image file inputs, "batch" for [Batch API](/docs/guides/batch), and "fine-tune" for [Fine-tuning](/docs/api-reference/fine-tuning).
	//
	Purpose CreateFileRequestPurpose `multipartForm:"name=purpose"`
}

func (o *CreateFileRequest) GetFile() CreateFileRequestFile {
	if o == nil {
		return CreateFileRequestFile{}
	}
	return o.File
}

func (o *CreateFileRequest) GetPurpose() CreateFileRequestPurpose {
	if o == nil {
		return CreateFileRequestPurpose("")
	}
	return o.Purpose
}
