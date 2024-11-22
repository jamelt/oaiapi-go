# VectorStoreExpirationAfter

The expiration policy for a vector store.


## Fields

| Field                                                                                            | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `Anchor`                                                                                         | [components.Anchor](../../models/components/anchor.md)                                           | :heavy_check_mark:                                                                               | Anchor timestamp after which the expiration policy applies. Supported anchors: `last_active_at`. |
| `Days`                                                                                           | *int64*                                                                                          | :heavy_check_mark:                                                                               | The number of days after the anchor time that the vector store will expire.                      |