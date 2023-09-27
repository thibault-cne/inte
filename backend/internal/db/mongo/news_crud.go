package mongo

import (
	"backend/internal/models"
	"context"
)

// CreateNews implements db.DBackend.
func (*Backend) CreateNews(ctx context.Context, news *models.News) error {
	panic("unimplemented")
}

// DeleteNews implements db.DBackend.
func (*Backend) DeleteNews(ctx context.Context, id string) error {
	panic("unimplemented")
}

// GetNews implements db.DBackend.
func (*Backend) GetNews(ctx context.Context, id string) (*models.News, error) {
	panic("unimplemented")
}

// RestoreNews implements db.DBackend.
func (*Backend) RestoreNews(ctx context.Context, id string) error {
	panic("unimplemented")
}

// UpdateNews implements db.DBackend.
func (*Backend) UpdateNews(ctx context.Context, news *models.News) error {
	panic("unimplemented")
}