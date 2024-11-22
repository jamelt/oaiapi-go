# Chat
(*Chat*)

## Overview

Given a list of messages comprising a conversation, the model will return a response.

### Available Operations

* [Create](#create) - Creates a model response for the given chat conversation. Learn more in the
[text generation](/docs/guides/text-generation), [vision](/docs/guides/vision),
and [audio](/docs/guides/audio) guides.


## Create

Creates a model response for the given chat conversation. Learn more in the
[text generation](/docs/guides/text-generation), [vision](/docs/guides/vision),
and [audio](/docs/guides/audio) guides.


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
    res, err := s.Chat.Create(ctx, components.CreateChatCompletionRequest{
        Messages: []components.ChatCompletionRequestMessage{
            components.CreateChatCompletionRequestMessageChatCompletionRequestUserMessage(
                components.ChatCompletionRequestUserMessage{
                    Content: components.CreateChatCompletionRequestUserMessageContentStr(
                        "<value>",
                    ),
                    Role: components.ChatCompletionRequestUserMessageRoleUser,
                },
            ),
            components.CreateChatCompletionRequestMessageChatCompletionRequestUserMessage(
                components.ChatCompletionRequestUserMessage{
                    Content: components.CreateChatCompletionRequestUserMessageContentArrayOfChatCompletionRequestUserMessageContentPart(
                        []components.ChatCompletionRequestUserMessageContentPart{
                            components.CreateChatCompletionRequestUserMessageContentPartChatCompletionRequestMessageContentPartImage(
                                components.ChatCompletionRequestMessageContentPartImage{
                                    Type: components.ChatCompletionRequestMessageContentPartImageTypeImageURL,
                                    ImageURL: components.ImageURL{
                                        URL: "https://rewarding-drug.com",
                                    },
                                },
                            ),
                            components.CreateChatCompletionRequestUserMessageContentPartChatCompletionRequestMessageContentPartText(
                                components.ChatCompletionRequestMessageContentPartText{
                                    Type: components.ChatCompletionRequestMessageContentPartTextTypeText,
                                    Text: "<value>",
                                },
                            ),
                        },
                    ),
                    Role: components.ChatCompletionRequestUserMessageRoleUser,
                },
            ),
            components.CreateChatCompletionRequestMessageChatCompletionRequestToolMessage(
                components.ChatCompletionRequestToolMessage{
                    Role: components.ChatCompletionRequestToolMessageRoleTool,
                    Content: components.CreateChatCompletionRequestToolMessageContentStr(
                        "<value>",
                    ),
                    ToolCallID: "<id>",
                },
            ),
        },
        Model: components.CreateCreateChatCompletionRequestModelStr(
            "gpt-4o",
        ),
        N: openaiapi.Int64(1),
        Temperature: openaiapi.Float64(1),
        TopP: openaiapi.Float64(1),
        User: openaiapi.String("user-1234"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateChatCompletionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [components.CreateChatCompletionRequest](../../models/components/createchatcompletionrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.CreateChatCompletionResponse](../../models/operations/createchatcompletionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |