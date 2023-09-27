package mongo

import (
	"backend/internal/models"
	"context"
)

// CountDeletedStars implements db.DBackend.
func (*Backend) CountDeletedStars(ctx context.Context, search string) (uint64, error) {
	panic("unimplemented")
}

// GetDeletedStars implements db.DBackend.
func (*Backend) GetDeletedStars(ctx context.Context, page uint64, size uint64, search string) ([]*models.Star, error) {
	panic("unimplemented")
}

// GetStars implements db.DBackend.
func (*Backend) GetStars(ctx context.Context, page uint64, size uint64, from string, to string) ([]*models.Star, error) {
	panic("unimplemented")
}