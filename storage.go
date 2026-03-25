// storage.go
package storage

import (
	"context"
	"io"
)

// Storage define el comportamiento que cualquier driver debe cumplir
type Storage interface {
	// Upload sube un archivo y devuelve la URL o el Path
	Upload(ctx context.Context, path string, file io.Reader) (string, error)
	// Delete elimina un archivo por su ruta
	Delete(ctx context.Context, path string) error
	// Exists comprueba si el archivo está ahí
	Exists(ctx context.Context, path string) (bool, error)
	// GetURL devuelve la URL pública o la ruta del archivo
	GetURL(path string) string
}
