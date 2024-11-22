# RequestCounts

The request counts for different statuses within the batch.


## Fields

| Field                                                     | Type                                                      | Required                                                  | Description                                               |
| --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- |
| `Total`                                                   | *int64*                                                   | :heavy_check_mark:                                        | Total number of requests in the batch.                    |
| `Completed`                                               | *int64*                                                   | :heavy_check_mark:                                        | Number of requests that have been completed successfully. |
| `Failed`                                                  | *int64*                                                   | :heavy_check_mark:                                        | Number of requests that have failed.                      |