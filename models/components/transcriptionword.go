// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type TranscriptionWord struct {
	// The text content of the word.
	Word string `json:"word"`
	// Start time of the word in seconds.
	Start float32 `json:"start"`
	// End time of the word in seconds.
	End float32 `json:"end"`
}

func (o *TranscriptionWord) GetWord() string {
	if o == nil {
		return ""
	}
	return o.Word
}

func (o *TranscriptionWord) GetStart() float32 {
	if o == nil {
		return 0.0
	}
	return o.Start
}

func (o *TranscriptionWord) GetEnd() float32 {
	if o == nil {
		return 0.0
	}
	return o.End
}