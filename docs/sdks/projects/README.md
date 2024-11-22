# Projects
(*Projects*)

## Overview

### Available Operations

* [List](#list) - Returns a list of projects.
* [Create](#create) - Create a new project in the organization. Projects can be created and archived, but cannot be deleted.
* [Get](#get) - Retrieves a project.
* [Update](#update) - Modifies a project in the organization.
* [ListAPIKeys](#listapikeys) - Returns a list of API keys in the project.
* [GetAPIKey](#getapikey) - Retrieves an API key in the project.
* [DeleteAPIKey](#deleteapikey) - Deletes an API key from the project.
* [Archive](#archive) - Archives a project in the organization. Archived projects cannot be used or updated.
* [ListServiceAccounts](#listserviceaccounts) - Returns a list of service accounts in the project.
* [CreateServiceAccount](#createserviceaccount) - Creates a new service account in the project. This also returns an unredacted API key for the service account.
* [GetServiceAccount](#getserviceaccount) - Retrieves a service account in the project.
* [DeleteServiceAccount](#deleteserviceaccount) - Deletes a service account from the project.
* [ListUsers](#listusers) - Returns a list of users in the project.
* [CreateUser](#createuser) - Adds a user to the project. Users must already be members of the organization to be added to a project.
* [GetUser](#getuser) - Retrieves a user in the project.
* [ModifyUser](#modifyuser) - Modifies a user's role in the project.
* [DeleteUser](#deleteuser) - Deletes a user from the project.

## List

Returns a list of projects.

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
    res, err := s.Projects.List(ctx, nil, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ProjectListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                                                                        | Type                                                                                                                                                                                                                                                                             | Required                                                                                                                                                                                                                                                                         | Description                                                                                                                                                                                                                                                                      |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                                                            | :heavy_check_mark:                                                                                                                                                                                                                                                               | The context to use for the request.                                                                                                                                                                                                                                              |
| `limit`                                                                                                                                                                                                                                                                          | **int64*                                                                                                                                                                                                                                                                         | :heavy_minus_sign:                                                                                                                                                                                                                                                               | A limit on the number of objects to be returned. Limit can range between 1 and 100, and the default is 20.<br/>                                                                                                                                                                  |
| `after`                                                                                                                                                                                                                                                                          | **string*                                                                                                                                                                                                                                                                        | :heavy_minus_sign:                                                                                                                                                                                                                                                               | A cursor for use in pagination. `after` is an object ID that defines your place in the list. For instance, if you make a list request and receive 100 objects, ending with obj_foo, your subsequent call can include after=obj_foo in order to fetch the next page of the list.<br/> |
| `includeArchived`                                                                                                                                                                                                                                                                | **bool*                                                                                                                                                                                                                                                                          | :heavy_minus_sign:                                                                                                                                                                                                                                                               | If `true` returns all projects including those that have been `archived`. Archived projects are not included by default.                                                                                                                                                         |
| `opts`                                                                                                                                                                                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                                                                                                                                                                                         | :heavy_minus_sign:                                                                                                                                                                                                                                                               | The options for this request.                                                                                                                                                                                                                                                    |

### Response

**[*operations.ListProjectsResponse](../../models/operations/listprojectsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Create

Create a new project in the organization. Projects can be created and archived, but cannot be deleted.

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
    res, err := s.Projects.Create(ctx, components.ProjectCreateRequest{
        Name: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Project != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [components.ProjectCreateRequest](../../models/components/projectcreaterequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.CreateProjectResponse](../../models/operations/createprojectresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Get

Retrieves a project.

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
    res, err := s.Projects.Get(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.Project != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `projectID`                                              | *string*                                                 | :heavy_check_mark:                                       | The ID of the project.                                   |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.RetrieveProjectResponse](../../models/operations/retrieveprojectresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Update

Modifies a project in the organization.

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
    res, err := s.Projects.Update(ctx, "<id>", components.ProjectUpdateRequest{
        Name: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Project != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `projectID`                                                                        | *string*                                                                           | :heavy_check_mark:                                                                 | The ID of the project.                                                             |
| `projectUpdateRequest`                                                             | [components.ProjectUpdateRequest](../../models/components/projectupdaterequest.md) | :heavy_check_mark:                                                                 | The project update request payload.                                                |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.ModifyProjectResponse](../../models/operations/modifyprojectresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400                     | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## ListAPIKeys

Returns a list of API keys in the project.

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
    res, err := s.Projects.ListAPIKeys(ctx, "<id>", nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ProjectAPIKeyListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                                                                        | Type                                                                                                                                                                                                                                                                             | Required                                                                                                                                                                                                                                                                         | Description                                                                                                                                                                                                                                                                      |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                                                            | :heavy_check_mark:                                                                                                                                                                                                                                                               | The context to use for the request.                                                                                                                                                                                                                                              |
| `projectID`                                                                                                                                                                                                                                                                      | *string*                                                                                                                                                                                                                                                                         | :heavy_check_mark:                                                                                                                                                                                                                                                               | The ID of the project.                                                                                                                                                                                                                                                           |
| `limit`                                                                                                                                                                                                                                                                          | **int64*                                                                                                                                                                                                                                                                         | :heavy_minus_sign:                                                                                                                                                                                                                                                               | A limit on the number of objects to be returned. Limit can range between 1 and 100, and the default is 20.<br/>                                                                                                                                                                  |
| `after`                                                                                                                                                                                                                                                                          | **string*                                                                                                                                                                                                                                                                        | :heavy_minus_sign:                                                                                                                                                                                                                                                               | A cursor for use in pagination. `after` is an object ID that defines your place in the list. For instance, if you make a list request and receive 100 objects, ending with obj_foo, your subsequent call can include after=obj_foo in order to fetch the next page of the list.<br/> |
| `opts`                                                                                                                                                                                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                                                                                                                                                                                         | :heavy_minus_sign:                                                                                                                                                                                                                                                               | The options for this request.                                                                                                                                                                                                                                                    |

### Response

**[*operations.ListProjectAPIKeysResponse](../../models/operations/listprojectapikeysresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetAPIKey

Retrieves an API key in the project.

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
    res, err := s.Projects.GetAPIKey(ctx, "<id>", "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.ProjectAPIKey != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `projectID`                                              | *string*                                                 | :heavy_check_mark:                                       | The ID of the project.                                   |
| `keyID`                                                  | *string*                                                 | :heavy_check_mark:                                       | The ID of the API key.                                   |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.RetrieveProjectAPIKeyResponse](../../models/operations/retrieveprojectapikeyresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## DeleteAPIKey

Deletes an API key from the project.

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
    res, err := s.Projects.DeleteAPIKey(ctx, "<id>", "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.ProjectAPIKeyDeleteResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `projectID`                                              | *string*                                                 | :heavy_check_mark:                                       | The ID of the project.                                   |
| `keyID`                                                  | *string*                                                 | :heavy_check_mark:                                       | The ID of the API key.                                   |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteProjectAPIKeyResponse](../../models/operations/deleteprojectapikeyresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400                     | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## Archive

Archives a project in the organization. Archived projects cannot be used or updated.

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
    res, err := s.Projects.Archive(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.Project != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `projectID`                                              | *string*                                                 | :heavy_check_mark:                                       | The ID of the project.                                   |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ArchiveProjectResponse](../../models/operations/archiveprojectresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## ListServiceAccounts

Returns a list of service accounts in the project.

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
    res, err := s.Projects.ListServiceAccounts(ctx, "<id>", nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ProjectServiceAccountListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                                                                        | Type                                                                                                                                                                                                                                                                             | Required                                                                                                                                                                                                                                                                         | Description                                                                                                                                                                                                                                                                      |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                                                            | :heavy_check_mark:                                                                                                                                                                                                                                                               | The context to use for the request.                                                                                                                                                                                                                                              |
| `projectID`                                                                                                                                                                                                                                                                      | *string*                                                                                                                                                                                                                                                                         | :heavy_check_mark:                                                                                                                                                                                                                                                               | The ID of the project.                                                                                                                                                                                                                                                           |
| `limit`                                                                                                                                                                                                                                                                          | **int64*                                                                                                                                                                                                                                                                         | :heavy_minus_sign:                                                                                                                                                                                                                                                               | A limit on the number of objects to be returned. Limit can range between 1 and 100, and the default is 20.<br/>                                                                                                                                                                  |
| `after`                                                                                                                                                                                                                                                                          | **string*                                                                                                                                                                                                                                                                        | :heavy_minus_sign:                                                                                                                                                                                                                                                               | A cursor for use in pagination. `after` is an object ID that defines your place in the list. For instance, if you make a list request and receive 100 objects, ending with obj_foo, your subsequent call can include after=obj_foo in order to fetch the next page of the list.<br/> |
| `opts`                                                                                                                                                                                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                                                                                                                                                                                         | :heavy_minus_sign:                                                                                                                                                                                                                                                               | The options for this request.                                                                                                                                                                                                                                                    |

### Response

**[*operations.ListProjectServiceAccountsResponse](../../models/operations/listprojectserviceaccountsresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400                     | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## CreateServiceAccount

Creates a new service account in the project. This also returns an unredacted API key for the service account.

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
    res, err := s.Projects.CreateServiceAccount(ctx, "<id>", components.ProjectServiceAccountCreateRequest{
        Name: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ProjectServiceAccountCreateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `projectID`                                                                                                    | *string*                                                                                                       | :heavy_check_mark:                                                                                             | The ID of the project.                                                                                         |
| `projectServiceAccountCreateRequest`                                                                           | [components.ProjectServiceAccountCreateRequest](../../models/components/projectserviceaccountcreaterequest.md) | :heavy_check_mark:                                                                                             | The project service account create request payload.                                                            |
| `opts`                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.CreateProjectServiceAccountResponse](../../models/operations/createprojectserviceaccountresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400                     | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## GetServiceAccount

Retrieves a service account in the project.

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
    res, err := s.Projects.GetServiceAccount(ctx, "<id>", "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.ProjectServiceAccount != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `projectID`                                              | *string*                                                 | :heavy_check_mark:                                       | The ID of the project.                                   |
| `serviceAccountID`                                       | *string*                                                 | :heavy_check_mark:                                       | The ID of the service account.                           |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.RetrieveProjectServiceAccountResponse](../../models/operations/retrieveprojectserviceaccountresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## DeleteServiceAccount

Deletes a service account from the project.

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
    res, err := s.Projects.DeleteServiceAccount(ctx, "<id>", "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.ProjectServiceAccountDeleteResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `projectID`                                              | *string*                                                 | :heavy_check_mark:                                       | The ID of the project.                                   |
| `serviceAccountID`                                       | *string*                                                 | :heavy_check_mark:                                       | The ID of the service account.                           |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteProjectServiceAccountResponse](../../models/operations/deleteprojectserviceaccountresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## ListUsers

Returns a list of users in the project.

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
    res, err := s.Projects.ListUsers(ctx, "<id>", nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ProjectUserListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                                                                        | Type                                                                                                                                                                                                                                                                             | Required                                                                                                                                                                                                                                                                         | Description                                                                                                                                                                                                                                                                      |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                                                            | :heavy_check_mark:                                                                                                                                                                                                                                                               | The context to use for the request.                                                                                                                                                                                                                                              |
| `projectID`                                                                                                                                                                                                                                                                      | *string*                                                                                                                                                                                                                                                                         | :heavy_check_mark:                                                                                                                                                                                                                                                               | The ID of the project.                                                                                                                                                                                                                                                           |
| `limit`                                                                                                                                                                                                                                                                          | **int64*                                                                                                                                                                                                                                                                         | :heavy_minus_sign:                                                                                                                                                                                                                                                               | A limit on the number of objects to be returned. Limit can range between 1 and 100, and the default is 20.<br/>                                                                                                                                                                  |
| `after`                                                                                                                                                                                                                                                                          | **string*                                                                                                                                                                                                                                                                        | :heavy_minus_sign:                                                                                                                                                                                                                                                               | A cursor for use in pagination. `after` is an object ID that defines your place in the list. For instance, if you make a list request and receive 100 objects, ending with obj_foo, your subsequent call can include after=obj_foo in order to fetch the next page of the list.<br/> |
| `opts`                                                                                                                                                                                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                                                                                                                                                                                         | :heavy_minus_sign:                                                                                                                                                                                                                                                               | The options for this request.                                                                                                                                                                                                                                                    |

### Response

**[*operations.ListProjectUsersResponse](../../models/operations/listprojectusersresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400                     | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## CreateUser

Adds a user to the project. Users must already be members of the organization to be added to a project.

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
    res, err := s.Projects.CreateUser(ctx, "<id>", components.ProjectUserCreateRequest{
        UserID: "<id>",
        Role: components.ProjectUserCreateRequestRoleMember,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ProjectUser != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `projectID`                                                                                | *string*                                                                                   | :heavy_check_mark:                                                                         | The ID of the project.                                                                     |
| `projectUserCreateRequest`                                                                 | [components.ProjectUserCreateRequest](../../models/components/projectusercreaterequest.md) | :heavy_check_mark:                                                                         | The project user create request payload.                                                   |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.CreateProjectUserResponse](../../models/operations/createprojectuserresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400                     | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## GetUser

Retrieves a user in the project.

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
    res, err := s.Projects.GetUser(ctx, "<id>", "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.ProjectUser != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `projectID`                                              | *string*                                                 | :heavy_check_mark:                                       | The ID of the project.                                   |
| `userID`                                                 | *string*                                                 | :heavy_check_mark:                                       | The ID of the user.                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.RetrieveProjectUserResponse](../../models/operations/retrieveprojectuserresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## ModifyUser

Modifies a user's role in the project.

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
    res, err := s.Projects.ModifyUser(ctx, "<id>", "<id>", components.ProjectUserUpdateRequest{
        Role: components.ProjectUserUpdateRequestRoleMember,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ProjectUser != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `projectID`                                                                                | *string*                                                                                   | :heavy_check_mark:                                                                         | The ID of the project.                                                                     |
| `userID`                                                                                   | *string*                                                                                   | :heavy_check_mark:                                                                         | The ID of the user.                                                                        |
| `projectUserUpdateRequest`                                                                 | [components.ProjectUserUpdateRequest](../../models/components/projectuserupdaterequest.md) | :heavy_check_mark:                                                                         | The project user update request payload.                                                   |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.ModifyProjectUserResponse](../../models/operations/modifyprojectuserresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400                     | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## DeleteUser

Deletes a user from the project.

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
    res, err := s.Projects.DeleteUser(ctx, "<id>", "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.ProjectUserDeleteResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `projectID`                                              | *string*                                                 | :heavy_check_mark:                                       | The ID of the project.                                   |
| `userID`                                                 | *string*                                                 | :heavy_check_mark:                                       | The ID of the user.                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteProjectUserResponse](../../models/operations/deleteprojectuserresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400                     | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |