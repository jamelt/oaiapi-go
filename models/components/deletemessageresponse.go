// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

type DeleteMessageResponseObject string

const (
	DeleteMessageResponseObjectThreadMessageDeleted DeleteMessageResponseObject = "thread.message.deleted"
)

func (e DeleteMessageResponseObject) ToPointer() *DeleteMessageResponseObject {
	return &e
}
func (e *DeleteMessageResponseObject) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "thread.message.deleted":
		*e = DeleteMessageResponseObject(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DeleteMessageResponseObject: %v", v)
	}
}

type DeleteMessageResponse struct {
	ID      string                      `json:"id"`
	Deleted bool                        `json:"deleted"`
	Object  DeleteMessageResponseObject `json:"object"`
}

func (o *DeleteMessageResponse) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *DeleteMessageResponse) GetDeleted() bool {
	if o == nil {
		return false
	}
	return o.Deleted
}

func (o *DeleteMessageResponse) GetObject() DeleteMessageResponseObject {
	if o == nil {
		return DeleteMessageResponseObject("")
	}
	return o.Object
}
