# github.com/jamelt/openai-api

Developer-friendly & type-safe Go SDK specifically catered to leverage *github.com/jamelt/openai-api* API.

<div align="left">
    <a href="https://www.speakeasy.com/?utm_source=github-com/jamelt/openai-api&utm_campaign=go"><img src="https://custom-icon-badges.demolab.com/badge/-Built%20By%20Speakeasy-212015?style=for-the-badge&logoColor=FBE331&logo=speakeasy&labelColor=545454" /></a>
    <a href="https://opensource.org/licenses/MIT">
        <img src="https://img.shields.io/badge/License-MIT-blue.svg" style="width: 100px; height: 28px;" />
    </a>
</div>


<br /><br />
> [!IMPORTANT]
> This SDK is not yet ready for production use. To complete setup please follow the steps outlined in your [workspace](https://app.speakeasy.com/org/jamelt/jamelt). Delete this section before > publishing to a package manager.

<!-- Start Summary [summary] -->
## Summary

OpenAI API: The OpenAI REST API. Please see https://platform.openai.com/docs/api-reference for more details.
<!-- End Summary [summary] -->

<!-- Start Table of Contents [toc] -->
## Table of Contents

* [SDK Installation](#sdk-installation)
* [SDK Example Usage](#sdk-example-usage)
* [Available Resources and Operations](#available-resources-and-operations)
* [Retries](#retries)
* [Error Handling](#error-handling)
* [Server Selection](#server-selection)
* [Custom HTTP Client](#custom-http-client)
* [Authentication](#authentication)
<!-- End Table of Contents [toc] -->

<!-- Start SDK Installation [installation] -->
## SDK Installation

To add the SDK as a dependency to your project:
```bash
go get github.com/jamelt/openai-api
```
<!-- End SDK Installation [installation] -->

<!-- Start SDK Example Usage [usage] -->
## SDK Example Usage

### Example

```go
package main

import (
	"context"
	openaiapi "github.com/jamelt/openai-api"
	"log"
	"os"
)

func main() {
	s := openaiapi.New(
		openaiapi.WithSecurity(os.Getenv("OPENAIAPI_API_KEY_AUTH")),
	)

	ctx := context.Background()
	res, err := s.Assistants.List(ctx, nil, nil, nil, nil)
	if err != nil {
		log.Fatal(err)
	}
	if res.ListAssistantsResponse != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

<details open>
<summary>Available methods</summary>

### [Assistants](docs/sdks/assistants/README.md)

* [List](docs/sdks/assistants/README.md#list) - Returns a list of assistants.
* [Create](docs/sdks/assistants/README.md#create) - Create an assistant with a model and instructions.
* [Get](docs/sdks/assistants/README.md#get) - Retrieves an assistant.
* [Modify](docs/sdks/assistants/README.md#modify) - Modifies an assistant.
* [Delete](docs/sdks/assistants/README.md#delete) - Delete an assistant.
* [CreateThread](docs/sdks/assistants/README.md#createthread) - Create a thread.
* [CreateThreadAndRun](docs/sdks/assistants/README.md#createthreadandrun) - Create a thread and run it in one request.
* [GetThread](docs/sdks/assistants/README.md#getthread) - Retrieves a thread.
* [ModifyThread](docs/sdks/assistants/README.md#modifythread) - Modifies a thread.
* [DeleteThread](docs/sdks/assistants/README.md#deletethread) - Delete a thread.
* [ListMessages](docs/sdks/assistants/README.md#listmessages) - Returns a list of messages for a given thread.
* [CreateMessage](docs/sdks/assistants/README.md#createmessage) - Create a message.
* [GetMessage](docs/sdks/assistants/README.md#getmessage) - Retrieve a message.
* [ModifyMessage](docs/sdks/assistants/README.md#modifymessage) - Modifies a message.
* [DeleteMessage](docs/sdks/assistants/README.md#deletemessage) - Deletes a message.
* [ListRuns](docs/sdks/assistants/README.md#listruns) - Returns a list of runs belonging to a thread.
* [CreateRun](docs/sdks/assistants/README.md#createrun) - Create a run.
* [GetRun](docs/sdks/assistants/README.md#getrun) - Retrieves a run.
* [ModifyRun](docs/sdks/assistants/README.md#modifyrun) - Modifies a run.
* [CancelRun](docs/sdks/assistants/README.md#cancelrun) - Cancels a run that is `in_progress`.
* [ListRunSteps](docs/sdks/assistants/README.md#listrunsteps) - Returns a list of run steps belonging to a run.
* [GetRunStep](docs/sdks/assistants/README.md#getrunstep) - Retrieves a run step.
* [SubmitToolOutputsToRun](docs/sdks/assistants/README.md#submittooloutputstorun) - When a run has the `status: "requires_action"` and `required_action.type` is `submit_tool_outputs`, this endpoint can be used to submit the outputs from the tool calls once they're all completed. All outputs must be submitted in a single request.


### [Audio](docs/sdks/audio/README.md)

* [CreateSpeech](docs/sdks/audio/README.md#createspeech) - Generates audio from the input text.
* [CreateTranscription](docs/sdks/audio/README.md#createtranscription) - Transcribes audio into the input language.
* [CreateTranslation](docs/sdks/audio/README.md#createtranslation) - Translates audio into English.

### [AuditLogs](docs/sdks/auditlogs/README.md)

* [List](docs/sdks/auditlogs/README.md#list) - List user actions and configuration changes within this organization.

### [Batch](docs/sdks/batch/README.md)

* [Get](docs/sdks/batch/README.md#get) - Retrieves a batch.

### [Batches](docs/sdks/batches/README.md)

* [Create](docs/sdks/batches/README.md#create) - Creates and executes a batch from an uploaded file of requests
* [List](docs/sdks/batches/README.md#list) - List your organization's batches.
* [Cancel](docs/sdks/batches/README.md#cancel) - Cancels an in-progress batch. The batch will be in status `cancelling` for up to 10 minutes, before changing to `cancelled`, where it will have partial results (if any) available in the output file.

### [Chat](docs/sdks/chat/README.md)

* [Create](docs/sdks/chat/README.md#create) - Creates a model response for the given chat conversation. Learn more in the
[text generation](/docs/guides/text-generation), [vision](/docs/guides/vision),
and [audio](/docs/guides/audio) guides.


### [Completions](docs/sdks/completions/README.md)

* [Create](docs/sdks/completions/README.md#create) - Creates a completion for the provided prompt and parameters.

### [Embeddings](docs/sdks/embeddings/README.md)

* [Create](docs/sdks/embeddings/README.md#create) - Creates an embedding vector representing the input text.

### [Files](docs/sdks/files/README.md)

* [List](docs/sdks/files/README.md#list) - Returns a list of files.
* [Create](docs/sdks/files/README.md#create) - Upload a file that can be used across various endpoints. Individual files can be up to 512 MB, and the size of all files uploaded by one organization can be up to 100 GB.

The Assistants API supports files up to 2 million tokens and of specific file types. See the [Assistants Tools guide](/docs/assistants/tools) for details.

The Fine-tuning API only supports `.jsonl` files. The input also has certain required formats for fine-tuning [chat](/docs/api-reference/fine-tuning/chat-input) or [completions](/docs/api-reference/fine-tuning/completions-input) models.

The Batch API only supports `.jsonl` files up to 100 MB in size. The input also has a specific required [format](/docs/api-reference/batch/request-input).

Please [contact us](https://help.openai.com/) if you need to increase these storage limits.

* [Delete](docs/sdks/files/README.md#delete) - Delete a file.
* [Get](docs/sdks/files/README.md#get) - Returns information about a specific file.
* [Download](docs/sdks/files/README.md#download) - Returns the contents of the specified file.

### [FineTuning](docs/sdks/finetuning/README.md)

* [CreateJob](docs/sdks/finetuning/README.md#createjob) - Creates a fine-tuning job which begins the process of creating a new model from a given dataset.

Response includes details of the enqueued job including job status and the name of the fine-tuned models once complete.

[Learn more about fine-tuning](/docs/guides/fine-tuning)

* [ListJobs](docs/sdks/finetuning/README.md#listjobs) - List your organization's fine-tuning jobs

* [GetJob](docs/sdks/finetuning/README.md#getjob) - Get info about a fine-tuning job.

[Learn more about fine-tuning](/docs/guides/fine-tuning)

* [CancelJob](docs/sdks/finetuning/README.md#canceljob) - Immediately cancel a fine-tune job.

* [ListJobCheckpoints](docs/sdks/finetuning/README.md#listjobcheckpoints) - List checkpoints for a fine-tuning job.

* [ListEvents](docs/sdks/finetuning/README.md#listevents) - Get status updates for a fine-tuning job.


### [Images](docs/sdks/images/README.md)

* [CreateEdit](docs/sdks/images/README.md#createedit) - Creates an edited or extended image given an original image and a prompt.
* [Create](docs/sdks/images/README.md#create) - Creates an image given a prompt.
* [CreateVariation](docs/sdks/images/README.md#createvariation) - Creates a variation of a given image.

### [Invites](docs/sdks/invites/README.md)

* [List](docs/sdks/invites/README.md#list) - Returns a list of invites in the organization.
* [Create](docs/sdks/invites/README.md#create) - Create an invite for a user to the organization. The invite must be accepted by the user before they have access to the organization.
* [Get](docs/sdks/invites/README.md#get) - Retrieves an invite.
* [Delete](docs/sdks/invites/README.md#delete) - Delete an invite. If the invite has already been accepted, it cannot be deleted.

### [Models](docs/sdks/models/README.md)

* [List](docs/sdks/models/README.md#list) - Lists the currently available models, and provides basic information about each one such as the owner and availability.
* [Get](docs/sdks/models/README.md#get) - Retrieves a model instance, providing basic information about the model such as the owner and permissioning.
* [Delete](docs/sdks/models/README.md#delete) - Delete a fine-tuned model. You must have the Owner role in your organization to delete a model.

### [Moderations](docs/sdks/moderations/README.md)

* [Create](docs/sdks/moderations/README.md#create) - Classifies if text and/or image inputs are potentially harmful. Learn
more in the [moderation guide](/docs/guides/moderation).



### [Projects](docs/sdks/projects/README.md)

* [List](docs/sdks/projects/README.md#list) - Returns a list of projects.
* [Create](docs/sdks/projects/README.md#create) - Create a new project in the organization. Projects can be created and archived, but cannot be deleted.
* [Get](docs/sdks/projects/README.md#get) - Retrieves a project.
* [Update](docs/sdks/projects/README.md#update) - Modifies a project in the organization.
* [ListAPIKeys](docs/sdks/projects/README.md#listapikeys) - Returns a list of API keys in the project.
* [GetAPIKey](docs/sdks/projects/README.md#getapikey) - Retrieves an API key in the project.
* [DeleteAPIKey](docs/sdks/projects/README.md#deleteapikey) - Deletes an API key from the project.
* [Archive](docs/sdks/projects/README.md#archive) - Archives a project in the organization. Archived projects cannot be used or updated.
* [ListServiceAccounts](docs/sdks/projects/README.md#listserviceaccounts) - Returns a list of service accounts in the project.
* [CreateServiceAccount](docs/sdks/projects/README.md#createserviceaccount) - Creates a new service account in the project. This also returns an unredacted API key for the service account.
* [GetServiceAccount](docs/sdks/projects/README.md#getserviceaccount) - Retrieves a service account in the project.
* [DeleteServiceAccount](docs/sdks/projects/README.md#deleteserviceaccount) - Deletes a service account from the project.
* [ListUsers](docs/sdks/projects/README.md#listusers) - Returns a list of users in the project.
* [CreateUser](docs/sdks/projects/README.md#createuser) - Adds a user to the project. Users must already be members of the organization to be added to a project.
* [GetUser](docs/sdks/projects/README.md#getuser) - Retrieves a user in the project.
* [ModifyUser](docs/sdks/projects/README.md#modifyuser) - Modifies a user's role in the project.
* [DeleteUser](docs/sdks/projects/README.md#deleteuser) - Deletes a user from the project.

### [Uploads](docs/sdks/uploads/README.md)

* [Create](docs/sdks/uploads/README.md#create) - Creates an intermediate [Upload](/docs/api-reference/uploads/object) object that you can add [Parts](/docs/api-reference/uploads/part-object) to. Currently, an Upload can accept at most 8 GB in total and expires after an hour after you create it.

Once you complete the Upload, we will create a [File](/docs/api-reference/files/object) object that contains all the parts you uploaded. This File is usable in the rest of our platform as a regular File object.

For certain `purpose`s, the correct `mime_type` must be specified. Please refer to documentation for the supported MIME types for your use case:
- [Assistants](/docs/assistants/tools/file-search#supported-files)

For guidance on the proper filename extensions for each purpose, please follow the documentation on [creating a File](/docs/api-reference/files/create).

* [Cancel](docs/sdks/uploads/README.md#cancel) - Cancels the Upload. No Parts may be added after an Upload is cancelled.

* [Complete](docs/sdks/uploads/README.md#complete) - Completes the [Upload](/docs/api-reference/uploads/object). 

Within the returned Upload object, there is a nested [File](/docs/api-reference/files/object) object that is ready to use in the rest of the platform.

You can specify the order of the Parts by passing in an ordered list of the Part IDs.

The number of bytes uploaded upon completion must match the number of bytes initially specified when creating the Upload object. No Parts may be added after an Upload is completed.

* [AddPart](docs/sdks/uploads/README.md#addpart) - Adds a [Part](/docs/api-reference/uploads/part-object) to an [Upload](/docs/api-reference/uploads/object) object. A Part represents a chunk of bytes from the file you are trying to upload. 

Each Part can be at most 64 MB, and you can add Parts until you hit the Upload maximum of 8 GB.

It is possible to add multiple Parts in parallel. You can decide the intended order of the Parts when you [complete the Upload](/docs/api-reference/uploads/complete).


### [Users](docs/sdks/users/README.md)

* [List](docs/sdks/users/README.md#list) - Lists all of the users in the organization.
* [Get](docs/sdks/users/README.md#get) - Retrieves a user by their identifier.
* [Modify](docs/sdks/users/README.md#modify) - Modifies a user's role in the organization.
* [Delete](docs/sdks/users/README.md#delete) - Deletes a user from the organization.

### [VectorStores](docs/sdks/vectorstores/README.md)

* [List](docs/sdks/vectorstores/README.md#list) - Returns a list of vector stores.
* [Create](docs/sdks/vectorstores/README.md#create) - Create a vector store.
* [Get](docs/sdks/vectorstores/README.md#get) - Retrieves a vector store.
* [Modify](docs/sdks/vectorstores/README.md#modify) - Modifies a vector store.
* [Delete](docs/sdks/vectorstores/README.md#delete) - Delete a vector store.
* [CreateFileBatch](docs/sdks/vectorstores/README.md#createfilebatch) - Create a vector store file batch.
* [GetFileBatch](docs/sdks/vectorstores/README.md#getfilebatch) - Retrieves a vector store file batch.
* [CancelFileBatch](docs/sdks/vectorstores/README.md#cancelfilebatch) - Cancel a vector store file batch. This attempts to cancel the processing of files in this batch as soon as possible.
* [ListBatchFiles](docs/sdks/vectorstores/README.md#listbatchfiles) - Returns a list of vector store files in a batch.
* [ListFiles](docs/sdks/vectorstores/README.md#listfiles) - Returns a list of vector store files.
* [CreateFile](docs/sdks/vectorstores/README.md#createfile) - Create a vector store file by attaching a [File](/docs/api-reference/files) to a [vector store](/docs/api-reference/vector-stores/object).
* [GetFile](docs/sdks/vectorstores/README.md#getfile) - Retrieves a vector store file.
* [DeleteFile](docs/sdks/vectorstores/README.md#deletefile) - Delete a vector store file. This will remove the file from the vector store but the file itself will not be deleted. To delete the file, use the [delete file](/docs/api-reference/files/delete) endpoint.

</details>
<!-- End Available Resources and Operations [operations] -->

<!-- Start Retries [retries] -->
## Retries

Some of the endpoints in this SDK support retries. If you use the SDK without any configuration, it will fall back to the default retry strategy provided by the API. However, the default retry strategy can be overridden on a per-operation basis, or across the entire SDK.

To change the default retry strategy for a single API call, simply provide a `retry.Config` object to the call by using the `WithRetries` option:
```go
package main

import (
	"context"
	openaiapi "github.com/jamelt/openai-api"
	"github.com/jamelt/openai-api/retry"
	"log"
	"models/operations"
	"os"
)

func main() {
	s := openaiapi.New(
		openaiapi.WithSecurity(os.Getenv("OPENAIAPI_API_KEY_AUTH")),
	)

	ctx := context.Background()
	res, err := s.Assistants.List(ctx, nil, nil, nil, nil, operations.WithRetries(
		retry.Config{
			Strategy: "backoff",
			Backoff: &retry.BackoffStrategy{
				InitialInterval: 1,
				MaxInterval:     50,
				Exponent:        1.1,
				MaxElapsedTime:  100,
			},
			RetryConnectionErrors: false,
		}))
	if err != nil {
		log.Fatal(err)
	}
	if res.ListAssistantsResponse != nil {
		// handle response
	}
}

```

If you'd like to override the default retry strategy for all operations that support retries, you can use the `WithRetryConfig` option at SDK initialization:
```go
package main

import (
	"context"
	openaiapi "github.com/jamelt/openai-api"
	"github.com/jamelt/openai-api/retry"
	"log"
	"os"
)

func main() {
	s := openaiapi.New(
		openaiapi.WithRetryConfig(
			retry.Config{
				Strategy: "backoff",
				Backoff: &retry.BackoffStrategy{
					InitialInterval: 1,
					MaxInterval:     50,
					Exponent:        1.1,
					MaxElapsedTime:  100,
				},
				RetryConnectionErrors: false,
			}),
		openaiapi.WithSecurity(os.Getenv("OPENAIAPI_API_KEY_AUTH")),
	)

	ctx := context.Background()
	res, err := s.Assistants.List(ctx, nil, nil, nil, nil)
	if err != nil {
		log.Fatal(err)
	}
	if res.ListAssistantsResponse != nil {
		// handle response
	}
}

```
<!-- End Retries [retries] -->

<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations. All operations return a response object or an error, they will never return both.

By Default, an API error will return `apierrors.APIError`. When custom error responses are specified for an operation, the SDK may also return their associated error. You can refer to respective *Errors* tables in SDK docs for more details on possible error types for each operation.

For example, the `Update` function may return the following errors:

| Error Type              | Status Code | Content Type     |
| ----------------------- | ----------- | ---------------- |
| apierrors.ErrorResponse | 400         | application/json |
| apierrors.APIError      | 4XX, 5XX    | \*/\*            |

### Example

```go
package main

import (
	"context"
	"errors"
	openaiapi "github.com/jamelt/openai-api"
	"github.com/jamelt/openai-api/models/apierrors"
	"github.com/jamelt/openai-api/models/components"
	"log"
	"os"
)

func main() {
	s := openaiapi.New(
		openaiapi.WithSecurity(os.Getenv("OPENAIAPI_API_KEY_AUTH")),
	)

	ctx := context.Background()
	res, err := s.Projects.Update(ctx, "<id>", components.ProjectUpdateRequest{
		Name: "<value>",
	})
	if err != nil {

		var e *apierrors.ErrorResponse
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *apierrors.APIError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}
	}
}

```
<!-- End Error Handling [errors] -->

<!-- Start Server Selection [server] -->
## Server Selection

### Override Server URL Per-Client

The default server can also be overridden globally using the `WithServerURL(serverURL string)` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	openaiapi "github.com/jamelt/openai-api"
	"log"
	"os"
)

func main() {
	s := openaiapi.New(
		openaiapi.WithServerURL("https://api.openai.com/v1"),
		openaiapi.WithSecurity(os.Getenv("OPENAIAPI_API_KEY_AUTH")),
	)

	ctx := context.Background()
	res, err := s.Assistants.List(ctx, nil, nil, nil, nil)
	if err != nil {
		log.Fatal(err)
	}
	if res.ListAssistantsResponse != nil {
		// handle response
	}
}

```
<!-- End Server Selection [server] -->

<!-- Start Custom HTTP Client [http-client] -->
## Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"
	"github.com/myorg/your-go-sdk"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = sdk.New(sdk.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client [http-client] -->

<!-- Start Authentication [security] -->
## Authentication

### Per-Client Security Schemes

This SDK supports the following security scheme globally:

| Name         | Type | Scheme      | Environment Variable     |
| ------------ | ---- | ----------- | ------------------------ |
| `APIKeyAuth` | http | HTTP Bearer | `OPENAIAPI_API_KEY_AUTH` |

You can configure it using the `WithSecurity` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	openaiapi "github.com/jamelt/openai-api"
	"log"
	"os"
)

func main() {
	s := openaiapi.New(
		openaiapi.WithSecurity(os.Getenv("OPENAIAPI_API_KEY_AUTH")),
	)

	ctx := context.Background()
	res, err := s.Assistants.List(ctx, nil, nil, nil, nil)
	if err != nil {
		log.Fatal(err)
	}
	if res.ListAssistantsResponse != nil {
		// handle response
	}
}

```
<!-- End Authentication [security] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->

# Development

## Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

## Contributions

While we value open-source contributions to this SDK, this library is generated programmatically. Any manual changes added to internal files will be overwritten on the next generation. 
We look forward to hearing your feedback. Feel free to open a PR or an issue with a proof of concept and we'll do our best to include it in a future release. 

### SDK Created by [Speakeasy](https://www.speakeasy.com/?utm_source=github-com/jamelt/openai-api&utm_campaign=go)
