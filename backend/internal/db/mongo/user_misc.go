package mongo

import (
	"backend/internal/models"
	"context"
)

// GetUsers implements db.DBackend.
func (*Backend) GetUsers(ctx context.Context, page uint64, size uint64, search string) ([]*models.User, error) {
	panic("unimplemented")
}

// CountDeletedUsers implements db.DBackend.
func (*Backend) CountDeletedUsers(ctx context.Context, search string) (uint64, error) {
	panic("unimplemented")
}

// GetDeletedUsers implements db.DBackend.
func (*Backend) GetDeletedUsers(ctx context.Context, page uint64, size uint64, search string) ([]*models.User, error) {
	panic("unimplemented")
}