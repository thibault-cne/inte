package mongo

import (
	"backend/internal/models"
	"context"
)

// CountDeletedNews implements db.DBackend.
func (*Backend) CountDeletedNews(ctx context.Context, search string) (uint64, error) {
	panic("unimplemented")
}

// GetDeletedNews implements db.DBackend.
func (*Backend) GetDeletedNews(ctx context.Context, page uint64, size uint64, search string) ([]*models.News, error) {
	panic("unimplemented")
}

// GetNewsList implements db.DBackend.
func (*Backend) GetNewsList(ctx context.Context, page uint64, size uint64) ([]*models.News, error) {
	panic("unimplemented")
}