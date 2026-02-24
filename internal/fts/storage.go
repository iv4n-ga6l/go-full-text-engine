package fts

import "sync"

// Global variables for storing documents and indexes.
var (
	// documents stores all the documents indexed by their unique ID.
	documents = make(map[string]Document)

	// index represents the inverted index for full-text search.
	index = make(map[string][]string) // Maps terms to document IDs

	// mutex ensures thread-safe access to documents and index.
	mutex sync.RWMutex
)
