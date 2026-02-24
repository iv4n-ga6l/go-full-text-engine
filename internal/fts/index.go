package fts

import "sync"

// Global variables for managing documents and indexes.
var (
	// documents stores all the documents in the FTS engine.
	documents = make(map[string]Document)

	// index represents the inverted index for the FTS engine.
	index = make(map[string][]string) // Maps terms to document IDs

	// mutex ensures thread-safe access to documents and index.
	mutex = sync.RWMutex{}
)
