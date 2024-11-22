<!-- Start SDK Example Usage [usage] -->
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