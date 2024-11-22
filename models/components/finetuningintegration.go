// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

// FineTuningIntegrationType - The type of the integration being enabled for the fine-tuning job
type FineTuningIntegrationType string

const (
	FineTuningIntegrationTypeWandb FineTuningIntegrationType = "wandb"
)

func (e FineTuningIntegrationType) ToPointer() *FineTuningIntegrationType {
	return &e
}
func (e *FineTuningIntegrationType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "wandb":
		*e = FineTuningIntegrationType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FineTuningIntegrationType: %v", v)
	}
}

// Wandb - The settings for your integration with Weights and Biases. This payload specifies the project that
// metrics will be sent to. Optionally, you can set an explicit display name for your run, add tags
// to your run, and set a default entity (team, username, etc) to be associated with your run.
type Wandb struct {
	// The name of the project that the new run will be created under.
	//
	Project string `json:"project"`
	// A display name to set for the run. If not set, we will use the Job ID as the name.
	//
	Name *string `json:"name,omitempty"`
	// The entity to use for the run. This allows you to set the team or username of the WandB user that you would
	// like associated with the run. If not set, the default entity for the registered WandB API key is used.
	//
	Entity *string `json:"entity,omitempty"`
	// A list of tags to be attached to the newly created run. These tags are passed through directly to WandB. Some
	// default tags are generated by OpenAI: "openai/finetune", "openai/{base-model}", "openai/{ftjob-abcdef}".
	//
	Tags []string `json:"tags,omitempty"`
}

func (o *Wandb) GetProject() string {
	if o == nil {
		return ""
	}
	return o.Project
}

func (o *Wandb) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *Wandb) GetEntity() *string {
	if o == nil {
		return nil
	}
	return o.Entity
}

func (o *Wandb) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

type FineTuningIntegration struct {
	// The type of the integration being enabled for the fine-tuning job
	Type FineTuningIntegrationType `json:"type"`
	// The settings for your integration with Weights and Biases. This payload specifies the project that
	// metrics will be sent to. Optionally, you can set an explicit display name for your run, add tags
	// to your run, and set a default entity (team, username, etc) to be associated with your run.
	//
	Wandb Wandb `json:"wandb"`
}

func (o *FineTuningIntegration) GetType() FineTuningIntegrationType {
	if o == nil {
		return FineTuningIntegrationType("")
	}
	return o.Type
}

func (o *FineTuningIntegration) GetWandb() Wandb {
	if o == nil {
		return Wandb{}
	}
	return o.Wandb
}