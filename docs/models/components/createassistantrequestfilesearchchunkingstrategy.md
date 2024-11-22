# CreateAssistantRequestFileSearchChunkingStrategy

The chunking strategy used to chunk the file(s). If not set, will use the `auto` strategy.


## Supported Types

### AutoChunkingStrategy

```go
createAssistantRequestFileSearchChunkingStrategy := components.CreateCreateAssistantRequestFileSearchChunkingStrategyAutoChunkingStrategy(components.AutoChunkingStrategy{/* values here */})
```

### CreateAssistantRequestChunkingStrategyStaticChunkingStrategy

```go
createAssistantRequestFileSearchChunkingStrategy := components.CreateCreateAssistantRequestFileSearchChunkingStrategyCreateAssistantRequestChunkingStrategyStaticChunkingStrategy(components.CreateAssistantRequestChunkingStrategyStaticChunkingStrategy{/* values here */})
```

