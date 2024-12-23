// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type ProjectUpdateRequest struct {
	// The updated name of the project, this name appears in reports.
	Name string `json:"name"`
}

func (o *ProjectUpdateRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}
