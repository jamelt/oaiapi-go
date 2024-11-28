# Users
(*Users*)

## Overview

### Available Operations

* [List](#list) - Lists all of the users in the organization.
* [Get](#get) - Retrieves a user by their identifier.
* [Modify](#modify) - Modifies a user's role in the organization.
* [Delete](#delete) - Deletes a user from the organization.

## List

Lists all of the users in the organization.

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
    res, err := s.Users.List(ctx, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.UserListResponse != nil {
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
| `opts`                                                                                                                                                                                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                                                                                                                                                                                         | :heavy_minus_sign:                                                                                                                                                                                                                                                               | The options for this request.                                                                                                                                                                                                                                                    |

### Response

**[*operations.ListUsersResponse](../../models/operations/listusersresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Get

Retrieves a user by their identifier.

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
    res, err := s.Users.Get(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.User != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `userID`                                                 | *string*                                                 | :heavy_check_mark:                                       | The ID of the user.                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.RetrieveUserResponse](../../models/operations/retrieveuserresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Modify

Modifies a user's role in the organization.

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
    res, err := s.Users.Modify(ctx, "<id>", components.UserRoleUpdateRequest{
        Role: components.UserRoleUpdateRequestRoleReader,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.User != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `userID`                                                                             | *string*                                                                             | :heavy_check_mark:                                                                   | The ID of the user.                                                                  |
| `userRoleUpdateRequest`                                                              | [components.UserRoleUpdateRequest](../../models/components/userroleupdaterequest.md) | :heavy_check_mark:                                                                   | The new user role to modify. This must be one of `owner` or `member`.                |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.ModifyUserResponse](../../models/operations/modifyuserresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Delete

Deletes a user from the organization.

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
    res, err := s.Users.Delete(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.UserDeleteResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `userID`                                                 | *string*                                                 | :heavy_check_mark:                                       | The ID of the user.                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteUserResponse](../../models/operations/deleteuserresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |