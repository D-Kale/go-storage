// storage.go
package storage

import (
	"context"
	"io"
)

// Storage define el comportamiento que cualquier driver debe cumplir
type Storage interface {
	Upload(ctx context.Context, path string, file io.Reader) (string, error)
	GetURL(path string) string
	Delete(ctx context.Context, path string) error
}
