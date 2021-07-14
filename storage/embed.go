package storage

import "context"

// EmbeddingFunc is a function type for embedding text into a vector.
type EmbeddingFunc func(ctx context.Context, text string) ([]float32, error)

// hobby-session-24

// hobby-session-1

// hobby-session-51
