# Embeddings
(*Embeddings*)

## Overview

Get a vector representation of a given input that can be easily consumed by machine learning models and algorithms.

### Available Operations

* [Create](#create) - Creates an embedding vector representing the input text.

## Create

Creates an embedding vector representing the input text.

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
    res, err := s.Embeddings.Create(ctx, components.CreateEmbeddingRequest{
        Input: components.CreateInputStr(
            "The quick brown fox jumped over the lazy dog",
        ),
        Model: components.CreateCreateEmbeddingRequestModelStr(
            "text-embedding-3-small",
        ),
        EncodingFormat: components.EncodingFormatFloat.ToPointer(),
        User: openaiapi.String("user-1234"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateEmbeddingResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [components.CreateEmbeddingRequest](../../models/components/createembeddingrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.CreateEmbeddingResponse](../../models/operations/createembeddingresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |