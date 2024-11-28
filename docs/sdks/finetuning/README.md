# FineTuning
(*FineTuning*)

## Overview

### Available Operations

* [CreateJob](#createjob) - Creates a fine-tuning job which begins the process of creating a new model from a given dataset.

Response includes details of the enqueued job including job status and the name of the fine-tuned models once complete.

[Learn more about fine-tuning](/docs/guides/fine-tuning)

* [ListJobs](#listjobs) - List your organization's fine-tuning jobs

* [GetJob](#getjob) - Get info about a fine-tuning job.

[Learn more about fine-tuning](/docs/guides/fine-tuning)

* [CancelJob](#canceljob) - Immediately cancel a fine-tune job.

* [ListJobCheckpoints](#listjobcheckpoints) - List checkpoints for a fine-tuning job.

* [ListEvents](#listevents) - Get status updates for a fine-tuning job.


## CreateJob

Creates a fine-tuning job which begins the process of creating a new model from a given dataset.

Response includes details of the enqueued job including job status and the name of the fine-tuned models once complete.

[Learn more about fine-tuning](/docs/guides/fine-tuning)


### Example Usage

```go
package main

import(
	"os"
	openaiapi "github.com/jamelt/openai-api"
	"github.com/jamelt/openai-api/models/components"
	"context"
	"log"
)

func main() {
    s := openaiapi.New(
        openaiapi.WithSecurity(os.Getenv("OPENAIAPI_API_KEY_AUTH")),
    )

    ctx := context.Background()
    res, err := s.FineTuning.CreateJob(ctx, components.CreateFineTuningJobRequest{
        Model: components.CreateCreateFineTuningJobRequestModelStr(
            "gpt-4o-mini",
        ),
        TrainingFile: "file-abc123",
        ValidationFile: openaiapi.String("file-abc123"),
        Integrations: []components.Integrations{
            components.Integrations{
                Type: components.CreateCreateFineTuningJobRequestTypeType1(
                    components.Type1Wandb,
                ),
                Wandb: components.CreateFineTuningJobRequestWandb{
                    Project: "my-wandb-project",
                    Tags: []string{
                        "custom-tag",
                    },
                },
            },
        },
        Seed: openaiapi.Int64(42),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.FineTuningJob != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [components.CreateFineTuningJobRequest](../../models/components/createfinetuningjobrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.CreateFineTuningJobResponse](../../models/operations/createfinetuningjobresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## ListJobs

List your organization's fine-tuning jobs


### Example Usage

```go
package main

import(
	"os"
	openaiapi "github.com/jamelt/openai-api"
	"context"
	"log"
)

func main() {
    s := openaiapi.New(
        openaiapi.WithSecurity(os.Getenv("OPENAIAPI_API_KEY_AUTH")),
    )

    ctx := context.Background()
    res, err := s.FineTuning.ListJobs(ctx, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ListPaginatedFineTuningJobsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                         | Type                                                              | Required                                                          | Description                                                       |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| `ctx`                                                             | [context.Context](https://pkg.go.dev/context#Context)             | :heavy_check_mark:                                                | The context to use for the request.                               |
| `after`                                                           | **string*                                                         | :heavy_minus_sign:                                                | Identifier for the last job from the previous pagination request. |
| `limit`                                                           | **int64*                                                          | :heavy_minus_sign:                                                | Number of fine-tuning jobs to retrieve.                           |
| `opts`                                                            | [][operations.Option](../../models/operations/option.md)          | :heavy_minus_sign:                                                | The options for this request.                                     |

### Response

**[*operations.ListPaginatedFineTuningJobsResponse](../../models/operations/listpaginatedfinetuningjobsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetJob

Get info about a fine-tuning job.

[Learn more about fine-tuning](/docs/guides/fine-tuning)


### Example Usage

```go
package main

import(
	"os"
	openaiapi "github.com/jamelt/openai-api"
	"context"
	"log"
)

func main() {
    s := openaiapi.New(
        openaiapi.WithSecurity(os.Getenv("OPENAIAPI_API_KEY_AUTH")),
    )

    ctx := context.Background()
    res, err := s.FineTuning.GetJob(ctx, "ft-AF1WoRqd3aJAHsqc9NY7iL8F")
    if err != nil {
        log.Fatal(err)
    }
    if res.FineTuningJob != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `fineTuningJobID`                                        | *string*                                                 | :heavy_check_mark:                                       | The ID of the fine-tuning job.<br/>                      | ft-AF1WoRqd3aJAHsqc9NY7iL8F                              |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.RetrieveFineTuningJobResponse](../../models/operations/retrievefinetuningjobresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## CancelJob

Immediately cancel a fine-tune job.


### Example Usage

```go
package main

import(
	"os"
	openaiapi "github.com/jamelt/openai-api"
	"context"
	"log"
)

func main() {
    s := openaiapi.New(
        openaiapi.WithSecurity(os.Getenv("OPENAIAPI_API_KEY_AUTH")),
    )

    ctx := context.Background()
    res, err := s.FineTuning.CancelJob(ctx, "ft-AF1WoRqd3aJAHsqc9NY7iL8F")
    if err != nil {
        log.Fatal(err)
    }
    if res.FineTuningJob != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `fineTuningJobID`                                        | *string*                                                 | :heavy_check_mark:                                       | The ID of the fine-tuning job to cancel.<br/>            | ft-AF1WoRqd3aJAHsqc9NY7iL8F                              |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.CancelFineTuningJobResponse](../../models/operations/cancelfinetuningjobresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## ListJobCheckpoints

List checkpoints for a fine-tuning job.


### Example Usage

```go
package main

import(
	"os"
	openaiapi "github.com/jamelt/openai-api"
	"context"
	"log"
)

func main() {
    s := openaiapi.New(
        openaiapi.WithSecurity(os.Getenv("OPENAIAPI_API_KEY_AUTH")),
    )

    ctx := context.Background()
    res, err := s.FineTuning.ListJobCheckpoints(ctx, "ft-AF1WoRqd3aJAHsqc9NY7iL8F", nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ListFineTuningJobCheckpointsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                   | Type                                                                        | Required                                                                    | Description                                                                 | Example                                                                     |
| --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- |
| `ctx`                                                                       | [context.Context](https://pkg.go.dev/context#Context)                       | :heavy_check_mark:                                                          | The context to use for the request.                                         |                                                                             |
| `fineTuningJobID`                                                           | *string*                                                                    | :heavy_check_mark:                                                          | The ID of the fine-tuning job to get checkpoints for.<br/>                  | ft-AF1WoRqd3aJAHsqc9NY7iL8F                                                 |
| `after`                                                                     | **string*                                                                   | :heavy_minus_sign:                                                          | Identifier for the last checkpoint ID from the previous pagination request. |                                                                             |
| `limit`                                                                     | **int64*                                                                    | :heavy_minus_sign:                                                          | Number of checkpoints to retrieve.                                          |                                                                             |
| `opts`                                                                      | [][operations.Option](../../models/operations/option.md)                    | :heavy_minus_sign:                                                          | The options for this request.                                               |                                                                             |

### Response

**[*operations.ListFineTuningJobCheckpointsResponse](../../models/operations/listfinetuningjobcheckpointsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## ListEvents

Get status updates for a fine-tuning job.


### Example Usage

```go
package main

import(
	"os"
	openaiapi "github.com/jamelt/openai-api"
	"context"
	"log"
)

func main() {
    s := openaiapi.New(
        openaiapi.WithSecurity(os.Getenv("OPENAIAPI_API_KEY_AUTH")),
    )

    ctx := context.Background()
    res, err := s.FineTuning.ListEvents(ctx, "ft-AF1WoRqd3aJAHsqc9NY7iL8F", nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ListFineTuningJobEventsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         | Example                                                             |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `ctx`                                                               | [context.Context](https://pkg.go.dev/context#Context)               | :heavy_check_mark:                                                  | The context to use for the request.                                 |                                                                     |
| `fineTuningJobID`                                                   | *string*                                                            | :heavy_check_mark:                                                  | The ID of the fine-tuning job to get events for.<br/>               | ft-AF1WoRqd3aJAHsqc9NY7iL8F                                         |
| `after`                                                             | **string*                                                           | :heavy_minus_sign:                                                  | Identifier for the last event from the previous pagination request. |                                                                     |
| `limit`                                                             | **int64*                                                            | :heavy_minus_sign:                                                  | Number of events to retrieve.                                       |                                                                     |
| `opts`                                                              | [][operations.Option](../../models/operations/option.md)            | :heavy_minus_sign:                                                  | The options for this request.                                       |                                                                     |

### Response

**[*operations.ListFineTuningEventsResponse](../../models/operations/listfinetuningeventsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |