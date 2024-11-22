// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type StaticChunkingStrategy struct {
	// The maximum number of tokens in each chunk. The default value is `800`. The minimum value is `100` and the maximum value is `4096`.
	MaxChunkSizeTokens int64 `json:"max_chunk_size_tokens"`
	// The number of tokens that overlap between chunks. The default value is `400`.
	//
	// Note that the overlap must not exceed half of `max_chunk_size_tokens`.
	//
	ChunkOverlapTokens int64 `json:"chunk_overlap_tokens"`
}

func (o *StaticChunkingStrategy) GetMaxChunkSizeTokens() int64 {
	if o == nil {
		return 0
	}
	return o.MaxChunkSizeTokens
}

func (o *StaticChunkingStrategy) GetChunkOverlapTokens() int64 {
	if o == nil {
		return 0
	}
	return o.ChunkOverlapTokens
}
