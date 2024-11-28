// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package openaiapi

import (
	"context"
	"fmt"
	"github.com/jamelt/openai-api/internal/hooks"
	"github.com/jamelt/openai-api/internal/utils"
	"github.com/jamelt/openai-api/models/components"
	"github.com/jamelt/openai-api/retry"
	"net/http"
	"time"
)

// ServerList contains the list of servers available to the SDK
var ServerList = []string{
	"https://api.openai.com/v1",
}

// HTTPClient provides an interface for suplying the SDK with a custom HTTP client
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// String provides a helper function to return a pointer to a string
func String(s string) *string { return &s }

// Bool provides a helper function to return a pointer to a bool
func Bool(b bool) *bool { return &b }

// Int provides a helper function to return a pointer to an int
func Int(i int) *int { return &i }

// Int64 provides a helper function to return a pointer to an int64
func Int64(i int64) *int64 { return &i }

// Float32 provides a helper function to return a pointer to a float32
func Float32(f float32) *float32 { return &f }

// Float64 provides a helper function to return a pointer to a float64
func Float64(f float64) *float64 { return &f }

// Pointer provides a helper function to return a pointer to a type
func Pointer[T any](v T) *T { return &v }

type sdkConfiguration struct {
	Client            HTTPClient
	Security          func(context.Context) (interface{}, error)
	ServerURL         string
	ServerIndex       int
	Language          string
	OpenAPIDocVersion string
	SDKVersion        string
	GenVersion        string
	UserAgent         string
	RetryConfig       *retry.Config
	Hooks             *hooks.Hooks
	Timeout           *time.Duration
}

func (c *sdkConfiguration) GetServerDetails() (string, map[string]string) {
	if c.ServerURL != "" {
		return c.ServerURL, nil
	}

	return ServerList[c.ServerIndex], nil
}

// OpenaiAPI - OpenAI API: The OpenAI REST API. Please see https://platform.openai.com/docs/api-reference for more details.
type OpenaiAPI struct {
	// Build Assistants that can call models and use tools.
	Assistants *Assistants
	// Turn audio into text or text into audio.
	Audio   *Audio
	Batches *Batches
	// Create large batches of API requests to run asynchronously.
	Batch *Batch
	// Given a list of messages comprising a conversation, the model will return a response.
	Chat *Chat
	// Given a prompt, the model will return one or more predicted completions, and can also return the probabilities of alternative tokens at each position.
	Completions *Completions
	// Get a vector representation of a given input that can be easily consumed by machine learning models and algorithms.
	Embeddings *Embeddings
	// Files are used to upload documents that can be used with features like Assistants and Fine-tuning.
	Files      *Files
	FineTuning *FineTuning
	// Given a prompt and/or an input image, the model will generate a new image.
	Images *Images
	// List and describe the various models available in the API.
	Models *Models
	// Given text and/or image inputs, classifies if those inputs are potentially harmful.
	Moderations *Moderations
	AuditLogs   *AuditLogs
	Invites     *Invites
	Projects    *Projects
	Users       *Users
	// Use Uploads to upload large files in multiple parts.
	Uploads      *Uploads
	VectorStores *VectorStores

	sdkConfiguration sdkConfiguration
}

type SDKOption func(*OpenaiAPI)

// WithServerURL allows the overriding of the default server URL
func WithServerURL(serverURL string) SDKOption {
	return func(sdk *OpenaiAPI) {
		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithTemplatedServerURL allows the overriding of the default server URL with a templated URL populated with the provided parameters
func WithTemplatedServerURL(serverURL string, params map[string]string) SDKOption {
	return func(sdk *OpenaiAPI) {
		if params != nil {
			serverURL = utils.ReplaceParameters(serverURL, params)
		}

		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithServerIndex allows the overriding of the default server by index
func WithServerIndex(serverIndex int) SDKOption {
	return func(sdk *OpenaiAPI) {
		if serverIndex < 0 || serverIndex >= len(ServerList) {
			panic(fmt.Errorf("server index %d out of range", serverIndex))
		}

		sdk.sdkConfiguration.ServerIndex = serverIndex
	}
}

// WithClient allows the overriding of the default HTTP client used by the SDK
func WithClient(client HTTPClient) SDKOption {
	return func(sdk *OpenaiAPI) {
		sdk.sdkConfiguration.Client = client
	}
}

// WithSecurity configures the SDK to use the provided security details
func WithSecurity(apiKeyAuth string) SDKOption {
	return func(sdk *OpenaiAPI) {
		security := components.Security{APIKeyAuth: &apiKeyAuth}
		sdk.sdkConfiguration.Security = utils.AsSecuritySource(&security)
	}
}

// WithSecuritySource configures the SDK to invoke the Security Source function on each method call to determine authentication
func WithSecuritySource(security func(context.Context) (components.Security, error)) SDKOption {
	return func(sdk *OpenaiAPI) {
		sdk.sdkConfiguration.Security = func(ctx context.Context) (interface{}, error) {
			return security(ctx)
		}
	}
}

func WithRetryConfig(retryConfig retry.Config) SDKOption {
	return func(sdk *OpenaiAPI) {
		sdk.sdkConfiguration.RetryConfig = &retryConfig
	}
}

// WithTimeout Optional request timeout applied to each operation
func WithTimeout(timeout time.Duration) SDKOption {
	return func(sdk *OpenaiAPI) {
		sdk.sdkConfiguration.Timeout = &timeout
	}
}

// New creates a new instance of the SDK with the provided options
func New(opts ...SDKOption) *OpenaiAPI {
	sdk := &OpenaiAPI{
		sdkConfiguration: sdkConfiguration{
			Language:          "go",
			OpenAPIDocVersion: "2.3.0",
			SDKVersion:        "0.2.0",
			GenVersion:        "2.467.4",
			UserAgent:         "speakeasy-sdk/go 0.2.0 2.467.4 2.3.0 github.com/jamelt/openai-api",
			Hooks:             hooks.New(),
		},
	}
	for _, opt := range opts {
		opt(sdk)
	}

	if sdk.sdkConfiguration.Security == nil {
		var envVarSecurity components.Security
		if utils.PopulateSecurityFromEnv(&envVarSecurity) {
			sdk.sdkConfiguration.Security = utils.AsSecuritySource(envVarSecurity)
		}
	}

	// Use WithClient to override the default client if you would like to customize the timeout
	if sdk.sdkConfiguration.Client == nil {
		sdk.sdkConfiguration.Client = &http.Client{Timeout: 60 * time.Second}
	}

	currentServerURL, _ := sdk.sdkConfiguration.GetServerDetails()
	serverURL := currentServerURL
	serverURL, sdk.sdkConfiguration.Client = sdk.sdkConfiguration.Hooks.SDKInit(currentServerURL, sdk.sdkConfiguration.Client)
	if serverURL != currentServerURL {
		sdk.sdkConfiguration.ServerURL = serverURL
	}

	sdk.Assistants = newAssistants(sdk.sdkConfiguration)

	sdk.Audio = newAudio(sdk.sdkConfiguration)

	sdk.Batches = newBatches(sdk.sdkConfiguration)

	sdk.Batch = newBatch(sdk.sdkConfiguration)

	sdk.Chat = newChat(sdk.sdkConfiguration)

	sdk.Completions = newCompletions(sdk.sdkConfiguration)

	sdk.Embeddings = newEmbeddings(sdk.sdkConfiguration)

	sdk.Files = newFiles(sdk.sdkConfiguration)

	sdk.FineTuning = newFineTuning(sdk.sdkConfiguration)

	sdk.Images = newImages(sdk.sdkConfiguration)

	sdk.Models = newModels(sdk.sdkConfiguration)

	sdk.Moderations = newModerations(sdk.sdkConfiguration)

	sdk.AuditLogs = newAuditLogs(sdk.sdkConfiguration)

	sdk.Invites = newInvites(sdk.sdkConfiguration)

	sdk.Projects = newProjects(sdk.sdkConfiguration)

	sdk.Users = newUsers(sdk.sdkConfiguration)

	sdk.Uploads = newUploads(sdk.sdkConfiguration)

	sdk.VectorStores = newVectorStores(sdk.sdkConfiguration)

	return sdk
}
