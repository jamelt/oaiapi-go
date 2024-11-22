# Images
(*Images*)

## Overview

Given a prompt and/or an input image, the model will generate a new image.

### Available Operations

* [CreateEdit](#createedit) - Creates an edited or extended image given an original image and a prompt.
* [Create](#create) - Creates an image given a prompt.
* [CreateVariation](#createvariation) - Creates a variation of a given image.

## CreateEdit

Creates an edited or extended image given an original image and a prompt.

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

    content, fileErr := os.Open("example.file")
    if fileErr != nil {
        panic(fileErr)
    }

    ctx := context.Background()
    res, err := s.Images.CreateEdit(ctx, components.CreateImageEditRequest{
        Image: components.CreateImageEditRequestImage{
            FileName: "example.file",
            Content: content,
        },
        Prompt: "A cute baby sea otter wearing a beret",
        Model: openaiapi.Pointer(components.CreateCreateImageEditRequestModelStr(
            "dall-e-2",
        )),
        N: openaiapi.Int64(1),
        Size: components.SizeOneThousandAndTwentyFourx1024.ToPointer(),
        ResponseFormat: components.CreateImageEditRequestResponseFormatURL.ToPointer(),
        User: openaiapi.String("user-1234"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ImagesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [components.CreateImageEditRequest](../../models/components/createimageeditrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.CreateImageEditResponse](../../models/operations/createimageeditresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Create

Creates an image given a prompt.

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
    res, err := s.Images.Create(ctx, components.CreateImageRequest{
        Prompt: "A cute baby sea otter",
        Model: openaiapi.Pointer(components.CreateCreateImageRequestModelStr(
            "dall-e-3",
        )),
        N: openaiapi.Int64(1),
        Quality: components.QualityStandard.ToPointer(),
        ResponseFormat: components.CreateImageRequestResponseFormatURL.ToPointer(),
        Size: components.CreateImageRequestSizeOneThousandAndTwentyFourx1024.ToPointer(),
        Style: components.StyleVivid.ToPointer(),
        User: openaiapi.String("user-1234"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ImagesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [components.CreateImageRequest](../../models/components/createimagerequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.CreateImageResponse](../../models/operations/createimageresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## CreateVariation

Creates a variation of a given image.

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

    content, fileErr := os.Open("example.file")
    if fileErr != nil {
        panic(fileErr)
    }

    ctx := context.Background()
    res, err := s.Images.CreateVariation(ctx, components.CreateImageVariationRequest{
        Image: components.CreateImageVariationRequestImage{
            FileName: "example.file",
            Content: content,
        },
        Model: openaiapi.Pointer(components.CreateCreateImageVariationRequestModelStr(
            "dall-e-2",
        )),
        N: openaiapi.Int64(1),
        ResponseFormat: components.CreateImageVariationRequestResponseFormatURL.ToPointer(),
        Size: components.CreateImageVariationRequestSizeOneThousandAndTwentyFourx1024.ToPointer(),
        User: openaiapi.String("user-1234"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ImagesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [components.CreateImageVariationRequest](../../models/components/createimagevariationrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.CreateImageVariationResponse](../../models/operations/createimagevariationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |