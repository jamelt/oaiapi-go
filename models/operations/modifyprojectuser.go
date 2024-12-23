// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/jamelt/openai-api/models/components"
)

type ModifyProjectUserRequest struct {
	// The ID of the project.
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
	// The ID of the user.
	UserID string `pathParam:"style=simple,explode=false,name=user_id"`
	// The project user update request payload.
	ProjectUserUpdateRequest components.ProjectUserUpdateRequest `request:"mediaType=application/json"`
}

func (o *ModifyProjectUserRequest) GetProjectID() string {
	if o == nil {
		return ""
	}
	return o.ProjectID
}

func (o *ModifyProjectUserRequest) GetUserID() string {
	if o == nil {
		return ""
	}
	return o.UserID
}

func (o *ModifyProjectUserRequest) GetProjectUserUpdateRequest() components.ProjectUserUpdateRequest {
	if o == nil {
		return components.ProjectUserUpdateRequest{}
	}
	return o.ProjectUserUpdateRequest
}

type ModifyProjectUserResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Project user's role updated successfully.
	ProjectUser *components.ProjectUser
}

func (o *ModifyProjectUserResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ModifyProjectUserResponse) GetProjectUser() *components.ProjectUser {
	if o == nil {
		return nil
	}
	return o.ProjectUser
}
