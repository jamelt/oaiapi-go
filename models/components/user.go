// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

// UserObject - The object type, which is always `organization.user`
type UserObject string

const (
	UserObjectOrganizationUser UserObject = "organization.user"
)

func (e UserObject) ToPointer() *UserObject {
	return &e
}
func (e *UserObject) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "organization.user":
		*e = UserObject(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UserObject: %v", v)
	}
}

// UserRole - `owner` or `reader`
type UserRole string

const (
	UserRoleOwner  UserRole = "owner"
	UserRoleReader UserRole = "reader"
)

func (e UserRole) ToPointer() *UserRole {
	return &e
}
func (e *UserRole) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "owner":
		fallthrough
	case "reader":
		*e = UserRole(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UserRole: %v", v)
	}
}

// User - Represents an individual `user` within an organization.
type User struct {
	// The object type, which is always `organization.user`
	Object UserObject `json:"object"`
	// The identifier, which can be referenced in API endpoints
	ID string `json:"id"`
	// The name of the user
	Name string `json:"name"`
	// The email address of the user
	Email string `json:"email"`
	// `owner` or `reader`
	Role UserRole `json:"role"`
	// The Unix timestamp (in seconds) of when the user was added.
	AddedAt int64 `json:"added_at"`
}

func (o *User) GetObject() UserObject {
	if o == nil {
		return UserObject("")
	}
	return o.Object
}

func (o *User) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *User) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *User) GetEmail() string {
	if o == nil {
		return ""
	}
	return o.Email
}

func (o *User) GetRole() UserRole {
	if o == nil {
		return UserRole("")
	}
	return o.Role
}

func (o *User) GetAddedAt() int64 {
	if o == nil {
		return 0
	}
	return o.AddedAt
}
