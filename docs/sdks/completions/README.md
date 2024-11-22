# Completions
(*Completions*)

## Overview

Given a prompt, the model will return one or more predicted completions, and can also return the probabilities of alternative tokens at each position.

### Available Operations

* [Create](#create) - Creates a completion for the provided prompt and parameters.

## Create

Creates a completion for the provided prompt and parameters.

### Example Usage

```go
package main

import(
	"os"
	openaiapi "github.com/jamelt/openai-api"
	"context"
	"github.com/jamelt/openai-api/models/components"
	"log"
)

func main() {
    s := openaiapi.New(
        openaiapi.WithSecurity(os.Getenv("OPENAIAPI_API_KEY_AUTH")),
    )

    ctx := context.Background()
    res, err := s.Completions.Create(ctx, components.CreateCompletionRequest{
        Model: components.CreateCreateCompletionRequestModelStr(
            "<value>",
        ),
        Prompt: components.CreatePromptStr(
            "This is a test.",
        ),
        MaxTokens: openaiapi.Int64(16),
        N: openaiapi.Int64(1),
        Stop: openaiapi.Pointer(components.CreateCreateCompletionRequestStopArrayOfStr(
            []string{
                "[\"\n\"]",
            },
        )),
        Suffix: openaiapi.String("test."),
        Temperature: openaiapi.Float64(1),
        TopP: openaiapi.Float64(1),
        User: openaiapi.String("user-1234"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateCompletionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [components.CreateCompletionRequest](../../models/components/createcompletionrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.CreateCompletionResponse](../../models/operations/createcompletionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |