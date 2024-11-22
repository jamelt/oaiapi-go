// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type CreateVectorStoreFileRequest struct {
	// A [File](/docs/api-reference/files) ID that the vector store should use. Useful for tools like `file_search` that can access files.
	FileID string `json:"file_id"`
	// The chunking strategy used to chunk the file(s). If not set, will use the `auto` strategy.
	ChunkingStrategy *ChunkingStrategyRequestParam `json:"chunking_strategy,omitempty"`
}

func (o *CreateVectorStoreFileRequest) GetFileID() string {
	if o == nil {
		return ""
	}
	return o.FileID
}

func (o *CreateVectorStoreFileRequest) GetChunkingStrategy() *ChunkingStrategyRequestParam {
	if o == nil {
		return nil
	}
	return o.ChunkingStrategy
}
