# Endpoint

The endpoint to be used for all requests in the batch. Currently `/v1/chat/completions`, `/v1/embeddings`, and `/v1/completions` are supported. Note that `/v1/embeddings` batches are also restricted to a maximum of 50,000 embedding inputs across all requests in the batch.


## Values

| Name                            | Value                           |
| ------------------------------- | ------------------------------- |
| `EndpointRootV1ChatCompletions` | /v1/chat/completions            |
| `EndpointRootV1Embeddings`      | /v1/embeddings                  |
| `EndpointRootV1Completions`     | /v1/completions                 |