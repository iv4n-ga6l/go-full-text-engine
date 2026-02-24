package fts

// Document represents a single document in the Full-Text Search engine.
type Document struct {
	ID      string // Unique identifier for the document
	Content string // The textual content of the document
	Title   string // Optional title for the document
}
