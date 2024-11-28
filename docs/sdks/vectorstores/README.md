# VectorStores
(*VectorStores*)

## Overview

### Available Operations

* [List](#list) - Returns a list of vector stores.
* [Create](#create) - Create a vector store.
* [Get](#get) - Retrieves a vector store.
* [Modify](#modify) - Modifies a vector store.
* [Delete](#delete) - Delete a vector store.
* [CreateFileBatch](#createfilebatch) - Create a vector store file batch.
* [GetFileBatch](#getfilebatch) - Retrieves a vector store file batch.
* [CancelFileBatch](#cancelfilebatch) - Cancel a vector store file batch. This attempts to cancel the processing of files in this batch as soon as possible.
* [ListBatchFiles](#listbatchfiles) - Returns a list of vector store files in a batch.
* [ListFiles](#listfiles) - Returns a list of vector store files.
* [CreateFile](#createfile) - Create a vector store file by attaching a [File](/docs/api-reference/files) to a [vector store](/docs/api-reference/vector-stores/object).
* [GetFile](#getfile) - Retrieves a vector store file.
* [DeleteFile](#deletefile) - Delete a vector store file. This will remove the file from the vector store but the file itself will not be deleted. To delete the file, use the [delete file](/docs/api-reference/files/delete) endpoint.

## List

Returns a list of vector stores.

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
    res, err := s.VectorStores.List(ctx, nil, nil, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ListVectorStoresResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                                                                                | Type                                                                                                                                                                                                                                                                                     | Required                                                                                                                                                                                                                                                                                 | Description                                                                                                                                                                                                                                                                              |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                                                                    | :heavy_check_mark:                                                                                                                                                                                                                                                                       | The context to use for the request.                                                                                                                                                                                                                                                      |
| `limit`                                                                                                                                                                                                                                                                                  | **int64*                                                                                                                                                                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                                                                                                                                       | A limit on the number of objects to be returned. Limit can range between 1 and 100, and the default is 20.<br/>                                                                                                                                                                          |
| `order`                                                                                                                                                                                                                                                                                  | [*operations.ListVectorStoresQueryParamOrder](../../models/operations/listvectorstoresqueryparamorder.md)                                                                                                                                                                                | :heavy_minus_sign:                                                                                                                                                                                                                                                                       | Sort order by the `created_at` timestamp of the objects. `asc` for ascending order and `desc` for descending order.<br/>                                                                                                                                                                 |
| `after`                                                                                                                                                                                                                                                                                  | **string*                                                                                                                                                                                                                                                                                | :heavy_minus_sign:                                                                                                                                                                                                                                                                       | A cursor for use in pagination. `after` is an object ID that defines your place in the list. For instance, if you make a list request and receive 100 objects, ending with obj_foo, your subsequent call can include after=obj_foo in order to fetch the next page of the list.<br/>     |
| `before`                                                                                                                                                                                                                                                                                 | **string*                                                                                                                                                                                                                                                                                | :heavy_minus_sign:                                                                                                                                                                                                                                                                       | A cursor for use in pagination. `before` is an object ID that defines your place in the list. For instance, if you make a list request and receive 100 objects, starting with obj_foo, your subsequent call can include before=obj_foo in order to fetch the previous page of the list.<br/> |
| `opts`                                                                                                                                                                                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                                                                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                                                                                                                                       | The options for this request.                                                                                                                                                                                                                                                            |

### Response

**[*operations.ListVectorStoresResponse](../../models/operations/listvectorstoresresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Create

Create a vector store.

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
    res, err := s.VectorStores.Create(ctx, components.CreateVectorStoreRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.VectorStoreObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [components.CreateVectorStoreRequest](../../models/components/createvectorstorerequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.CreateVectorStoreResponse](../../models/operations/createvectorstoreresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Get

Retrieves a vector store.

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
    res, err := s.VectorStores.Get(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.VectorStoreObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `vectorStoreID`                                          | *string*                                                 | :heavy_check_mark:                                       | The ID of the vector store to retrieve.                  |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetVectorStoreResponse](../../models/operations/getvectorstoreresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Modify

Modifies a vector store.

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
    res, err := s.VectorStores.Modify(ctx, "<id>", components.UpdateVectorStoreRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.VectorStoreObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `vectorStoreID`                                                                            | *string*                                                                                   | :heavy_check_mark:                                                                         | The ID of the vector store to modify.                                                      |
| `updateVectorStoreRequest`                                                                 | [components.UpdateVectorStoreRequest](../../models/components/updatevectorstorerequest.md) | :heavy_check_mark:                                                                         | N/A                                                                                        |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.ModifyVectorStoreResponse](../../models/operations/modifyvectorstoreresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Delete

Delete a vector store.

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
    res, err := s.VectorStores.Delete(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.DeleteVectorStoreResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `vectorStoreID`                                          | *string*                                                 | :heavy_check_mark:                                       | The ID of the vector store to delete.                    |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteVectorStoreResponse](../../models/operations/deletevectorstoreresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## CreateFileBatch

Create a vector store file batch.

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
    res, err := s.VectorStores.CreateFileBatch(ctx, "vs_abc123", components.CreateVectorStoreFileBatchRequest{
        FileIds: []string{
            "<value>",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.VectorStoreFileBatchObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  | Example                                                                                                      |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |                                                                                                              |
| `vectorStoreID`                                                                                              | *string*                                                                                                     | :heavy_check_mark:                                                                                           | The ID of the vector store for which to create a File Batch.<br/>                                            | vs_abc123                                                                                                    |
| `createVectorStoreFileBatchRequest`                                                                          | [components.CreateVectorStoreFileBatchRequest](../../models/components/createvectorstorefilebatchrequest.md) | :heavy_check_mark:                                                                                           | N/A                                                                                                          |                                                                                                              |
| `opts`                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |                                                                                                              |

### Response

**[*operations.CreateVectorStoreFileBatchResponse](../../models/operations/createvectorstorefilebatchresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetFileBatch

Retrieves a vector store file batch.

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
    res, err := s.VectorStores.GetFileBatch(ctx, "vs_abc123", "vsfb_abc123")
    if err != nil {
        log.Fatal(err)
    }
    if res.VectorStoreFileBatchObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                  | Type                                                       | Required                                                   | Description                                                | Example                                                    |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| `ctx`                                                      | [context.Context](https://pkg.go.dev/context#Context)      | :heavy_check_mark:                                         | The context to use for the request.                        |                                                            |
| `vectorStoreID`                                            | *string*                                                   | :heavy_check_mark:                                         | The ID of the vector store that the file batch belongs to. | vs_abc123                                                  |
| `batchID`                                                  | *string*                                                   | :heavy_check_mark:                                         | The ID of the file batch being retrieved.                  | vsfb_abc123                                                |
| `opts`                                                     | [][operations.Option](../../models/operations/option.md)   | :heavy_minus_sign:                                         | The options for this request.                              |                                                            |

### Response

**[*operations.GetVectorStoreFileBatchResponse](../../models/operations/getvectorstorefilebatchresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## CancelFileBatch

Cancel a vector store file batch. This attempts to cancel the processing of files in this batch as soon as possible.

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
    res, err := s.VectorStores.CancelFileBatch(ctx, "<id>", "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.VectorStoreFileBatchObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                  | Type                                                       | Required                                                   | Description                                                |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| `ctx`                                                      | [context.Context](https://pkg.go.dev/context#Context)      | :heavy_check_mark:                                         | The context to use for the request.                        |
| `vectorStoreID`                                            | *string*                                                   | :heavy_check_mark:                                         | The ID of the vector store that the file batch belongs to. |
| `batchID`                                                  | *string*                                                   | :heavy_check_mark:                                         | The ID of the file batch to cancel.                        |
| `opts`                                                     | [][operations.Option](../../models/operations/option.md)   | :heavy_minus_sign:                                         | The options for this request.                              |

### Response

**[*operations.CancelVectorStoreFileBatchResponse](../../models/operations/cancelvectorstorefilebatchresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## ListBatchFiles

Returns a list of vector store files in a batch.

### Example Usage

```go
package main

import(
	"os"
	openaiapi "github.com/jamelt/openai-api"
	"github.com/jamelt/openai-api/models/operations"
	"context"
	"log"
)

func main() {
    s := openaiapi.New(
        openaiapi.WithSecurity(os.Getenv("OPENAIAPI_API_KEY_AUTH")),
    )

    ctx := context.Background()
    res, err := s.VectorStores.ListBatchFiles(ctx, operations.ListFilesInVectorStoreBatchRequest{
        VectorStoreID: "<id>",
        BatchID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListVectorStoreFilesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.ListFilesInVectorStoreBatchRequest](../../models/operations/listfilesinvectorstorebatchrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.ListFilesInVectorStoreBatchResponse](../../models/operations/listfilesinvectorstorebatchresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## ListFiles

Returns a list of vector store files.

### Example Usage

```go
package main

import(
	"os"
	openaiapi "github.com/jamelt/openai-api"
	"github.com/jamelt/openai-api/models/operations"
	"context"
	"log"
)

func main() {
    s := openaiapi.New(
        openaiapi.WithSecurity(os.Getenv("OPENAIAPI_API_KEY_AUTH")),
    )

    ctx := context.Background()
    res, err := s.VectorStores.ListFiles(ctx, operations.ListVectorStoreFilesRequest{
        VectorStoreID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListVectorStoreFilesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.ListVectorStoreFilesRequest](../../models/operations/listvectorstorefilesrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.ListVectorStoreFilesResponse](../../models/operations/listvectorstorefilesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## CreateFile

Create a vector store file by attaching a [File](/docs/api-reference/files) to a [vector store](/docs/api-reference/vector-stores/object).

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
    res, err := s.VectorStores.CreateFile(ctx, "vs_abc123", components.CreateVectorStoreFileRequest{
        FileID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.VectorStoreFileObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        | Example                                                                                            |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |                                                                                                    |
| `vectorStoreID`                                                                                    | *string*                                                                                           | :heavy_check_mark:                                                                                 | The ID of the vector store for which to create a File.<br/>                                        | vs_abc123                                                                                          |
| `createVectorStoreFileRequest`                                                                     | [components.CreateVectorStoreFileRequest](../../models/components/createvectorstorefilerequest.md) | :heavy_check_mark:                                                                                 | N/A                                                                                                |                                                                                                    |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |                                                                                                    |

### Response

**[*operations.CreateVectorStoreFileResponse](../../models/operations/createvectorstorefileresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetFile

Retrieves a vector store file.

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
    res, err := s.VectorStores.GetFile(ctx, "vs_abc123", "file-abc123")
    if err != nil {
        log.Fatal(err)
    }
    if res.VectorStoreFileObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `vectorStoreID`                                          | *string*                                                 | :heavy_check_mark:                                       | The ID of the vector store that the file belongs to.     | vs_abc123                                                |
| `fileID`                                                 | *string*                                                 | :heavy_check_mark:                                       | The ID of the file being retrieved.                      | file-abc123                                              |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetVectorStoreFileResponse](../../models/operations/getvectorstorefileresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## DeleteFile

Delete a vector store file. This will remove the file from the vector store but the file itself will not be deleted. To delete the file, use the [delete file](/docs/api-reference/files/delete) endpoint.

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
    res, err := s.VectorStores.DeleteFile(ctx, "<id>", "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.DeleteVectorStoreFileResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `vectorStoreID`                                          | *string*                                                 | :heavy_check_mark:                                       | The ID of the vector store that the file belongs to.     |
| `fileID`                                                 | *string*                                                 | :heavy_check_mark:                                       | The ID of the file to delete.                            |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteVectorStoreFileResponse](../../models/operations/deletevectorstorefileresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |