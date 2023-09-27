package mongo

import (
	"backend/internal/models"
	"context"
)

// CountDeletedSuggestions implements db.DBackend.
func (*Backend) CountDeletedSuggestions(ctx context.Context, search string) (uint64, error) {
	panic("unimplemented")
}

// GetDeletedSuggestions implements db.DBackend.
func (*Backend) GetDeletedSuggestions(ctx context.Context, page uint64, size uint64, search string) ([]*models.Suggestion, error) {
	panic("unimplemented")
}