# FileSearchChunkingStrategy

The chunking strategy used to chunk the file(s). If not set, will use the `auto` strategy.


## Supported Types

### ChunkingStrategyAutoChunkingStrategy

```go
fileSearchChunkingStrategy := components.CreateFileSearchChunkingStrategyChunkingStrategyAutoChunkingStrategy(components.ChunkingStrategyAutoChunkingStrategy{/* values here */})
```

### ChunkingStrategyStaticChunkingStrategy

```go
fileSearchChunkingStrategy := components.CreateFileSearchChunkingStrategyChunkingStrategyStaticChunkingStrategy(components.ChunkingStrategyStaticChunkingStrategy{/* values here */})
```

