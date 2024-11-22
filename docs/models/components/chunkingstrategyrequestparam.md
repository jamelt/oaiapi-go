# ChunkingStrategyRequestParam

The chunking strategy used to chunk the file(s). If not set, will use the `auto` strategy.


## Supported Types

### AutoChunkingStrategyRequestParam

```go
chunkingStrategyRequestParam := components.CreateChunkingStrategyRequestParamAutoChunkingStrategyRequestParam(components.AutoChunkingStrategyRequestParam{/* values here */})
```

### StaticChunkingStrategyRequestParam

```go
chunkingStrategyRequestParam := components.CreateChunkingStrategyRequestParamStaticChunkingStrategyRequestParam(components.StaticChunkingStrategyRequestParam{/* values here */})
```

