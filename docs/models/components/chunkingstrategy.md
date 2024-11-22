# ChunkingStrategy

The chunking strategy used to chunk the file(s). If not set, will use the `auto` strategy. Only applicable if `file_ids` is non-empty.


## Supported Types

### AutoChunkingStrategyRequestParam

```go
chunkingStrategy := components.CreateChunkingStrategyAutoChunkingStrategyRequestParam(components.AutoChunkingStrategyRequestParam{/* values here */})
```

### StaticChunkingStrategyRequestParam

```go
chunkingStrategy := components.CreateChunkingStrategyStaticChunkingStrategyRequestParam(components.StaticChunkingStrategyRequestParam{/* values here */})
```

