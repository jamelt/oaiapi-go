// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

// RunStepDetailsToolCallsFunctionObjectType - The type of tool call. This is always going to be `function` for this type of tool call.
type RunStepDetailsToolCallsFunctionObjectType string

const (
	RunStepDetailsToolCallsFunctionObjectTypeFunction RunStepDetailsToolCallsFunctionObjectType = "function"
)

func (e RunStepDetailsToolCallsFunctionObjectType) ToPointer() *RunStepDetailsToolCallsFunctionObjectType {
	return &e
}
func (e *RunStepDetailsToolCallsFunctionObjectType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "function":
		*e = RunStepDetailsToolCallsFunctionObjectType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RunStepDetailsToolCallsFunctionObjectType: %v", v)
	}
}

// RunStepDetailsToolCallsFunctionObjectFunction - The definition of the function that was called.
type RunStepDetailsToolCallsFunctionObjectFunction struct {
	// The name of the function.
	Name string `json:"name"`
	// The arguments passed to the function.
	Arguments string `json:"arguments"`
	// The output of the function. This will be `null` if the outputs have not been [submitted](/docs/api-reference/runs/submitToolOutputs) yet.
	Output *string `json:"output"`
}

func (o *RunStepDetailsToolCallsFunctionObjectFunction) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *RunStepDetailsToolCallsFunctionObjectFunction) GetArguments() string {
	if o == nil {
		return ""
	}
	return o.Arguments
}

func (o *RunStepDetailsToolCallsFunctionObjectFunction) GetOutput() *string {
	if o == nil {
		return nil
	}
	return o.Output
}

type RunStepDetailsToolCallsFunctionObject struct {
	// The ID of the tool call object.
	ID string `json:"id"`
	// The type of tool call. This is always going to be `function` for this type of tool call.
	Type RunStepDetailsToolCallsFunctionObjectType `json:"type"`
	// The definition of the function that was called.
	Function RunStepDetailsToolCallsFunctionObjectFunction `json:"function"`
}

func (o *RunStepDetailsToolCallsFunctionObject) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *RunStepDetailsToolCallsFunctionObject) GetType() RunStepDetailsToolCallsFunctionObjectType {
	if o == nil {
		return RunStepDetailsToolCallsFunctionObjectType("")
	}
	return o.Type
}

func (o *RunStepDetailsToolCallsFunctionObject) GetFunction() RunStepDetailsToolCallsFunctionObjectFunction {
	if o == nil {
		return RunStepDetailsToolCallsFunctionObjectFunction{}
	}
	return o.Function
}
