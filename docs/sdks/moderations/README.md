# Moderations
(*Moderations*)

## Overview

Given text and/or image inputs, classifies if those inputs are potentially harmful.

### Available Operations

* [Create](#create) - Classifies if text and/or image inputs are potentially harmful. Learn
more in the [moderation guide](/docs/guides/moderation).


## Create

Classifies if text and/or image inputs are potentially harmful. Learn
more in the [moderation guide](/docs/guides/moderation).


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
    res, err := s.Moderations.Create(ctx, components.CreateModerationRequest{
        Input: components.CreateCreateModerationRequestInputArrayOf3(
            []components.Three{
                components.CreateThreeThree1(
                    components.Three1{
                        Type: components.ThreeTypeImageURL,
                        ImageURL: components.ThreeImageURL{
                            URL: "https://example.com/image.jpg",
                        },
                    },
                ),
                components.CreateThreeThree2(
                    components.Three2{
                        Type: components.CreateModerationRequest3TypeText,
                        Text: "I want to kill them",
                    },
                ),
                components.CreateThreeThree1(
                    components.Three1{
                        Type: components.ThreeTypeImageURL,
                        ImageURL: components.ThreeImageURL{
                            URL: "https://example.com/image.jpg",
                        },
                    },
                ),
            },
        ),
        Model: openaiapi.Pointer(components.CreateCreateModerationRequestModelStr(
            "omni-moderation-2024-09-26",
        )),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateModerationResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [components.CreateModerationRequest](../../models/components/createmoderationrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.CreateModerationResponse](../../models/operations/createmoderationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |