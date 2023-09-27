package mongo

import (
	"backend/internal/models"
	"context"
)

// CreateStar implements db.DBackend.
func (*Backend) CreateStar(ctx context.Context, star *models.Star) error {
	panic("unimplemented")
}

// DeleteStar implements db.DBackend.
func (*Backend) DeleteStar(ctx context.Context, id string) error {
	panic("unimplemented")
}

// GetStar implements db.DBackend.
func (*Backend) GetStar(ctx context.Context, id string) (*models.Star, error) {
	panic("unimplemented")
}

// RestoreStar implements db.DBackend.
func (*Backend) RestoreStar(ctx context.Context, id string) error {
	panic("unimplemented")
}

// UpdateStar implements db.DBackend.
func (*Backend) UpdateStar(ctx context.Context, star *models.Star) error {
	panic("unimplemented")
}