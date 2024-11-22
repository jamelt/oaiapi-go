// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type ImagesResponse struct {
	Created int64   `json:"created"`
	Data    []Image `json:"data"`
}

func (o *ImagesResponse) GetCreated() int64 {
	if o == nil {
		return 0
	}
	return o.Created
}

func (o *ImagesResponse) GetData() []Image {
	if o == nil {
		return []Image{}
	}
	return o.Data
}