// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

// CreateUploadRequestPurpose - The intended purpose of the uploaded file.
//
// See the [documentation on File purposes](/docs/api-reference/files/create#files-create-purpose).
type CreateUploadRequestPurpose string

const (
	CreateUploadRequestPurposeAssistants CreateUploadRequestPurpose = "assistants"
	CreateUploadRequestPurposeBatch      CreateUploadRequestPurpose = "batch"
	CreateUploadRequestPurposeFineTune   CreateUploadRequestPurpose = "fine-tune"
	CreateUploadRequestPurposeVision     CreateUploadRequestPurpose = "vision"
)

func (e CreateUploadRequestPurpose) ToPointer() *CreateUploadRequestPurpose {
	return &e
}
func (e *CreateUploadRequestPurpose) UnmarshalJSON(data []byte) error {
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
		*e = CreateUploadRequestPurpose(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateUploadRequestPurpose: %v", v)
	}
}

type CreateUploadRequest struct {
	// The name of the file to upload.
	//
	Filename string `json:"filename"`
	// The intended purpose of the uploaded file.
	//
	// See the [documentation on File purposes](/docs/api-reference/files/create#files-create-purpose).
	//
	Purpose CreateUploadRequestPurpose `json:"purpose"`
	// The number of bytes in the file you are uploading.
	//
	Bytes int64 `json:"bytes"`
	// The MIME type of the file.
	//
	// This must fall within the supported MIME types for your file purpose. See the supported MIME types for assistants and vision.
	//
	MimeType string `json:"mime_type"`
}

func (o *CreateUploadRequest) GetFilename() string {
	if o == nil {
		return ""
	}
	return o.Filename
}

func (o *CreateUploadRequest) GetPurpose() CreateUploadRequestPurpose {
	if o == nil {
		return CreateUploadRequestPurpose("")
	}
	return o.Purpose
}

func (o *CreateUploadRequest) GetBytes() int64 {
	if o == nil {
		return 0
	}
	return o.Bytes
}

func (o *CreateUploadRequest) GetMimeType() string {
	if o == nil {
		return ""
	}
	return o.MimeType
}
