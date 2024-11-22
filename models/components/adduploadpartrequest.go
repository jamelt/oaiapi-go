// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type Data struct {
	FileName string `multipartForm:"name=data"`
	// This field accepts []byte data or io.Reader implementations, such as *os.File.
	Content any `multipartForm:"content"`
}

func (o *Data) GetFileName() string {
	if o == nil {
		return ""
	}
	return o.FileName
}

func (o *Data) GetContent() any {
	if o == nil {
		return nil
	}
	return o.Content
}

type AddUploadPartRequest struct {
	// The chunk of bytes for this Part.
	//
	Data Data `multipartForm:"file"`
}

func (o *AddUploadPartRequest) GetData() Data {
	if o == nil {
		return Data{}
	}
	return o.Data
}
