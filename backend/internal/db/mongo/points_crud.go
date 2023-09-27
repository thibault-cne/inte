package mongo

import (
	"backend/internal/models"
	"context"
)

// CreatePointsTransaction implements db.DBackend.
func (*Backend) CreatePointsTransaction(ctx context.Context, pointsTransaction *models.PointsTransaction) error {
	panic("unimplemented")
}

// DeletePointsTransaction implements db.DBackend.
func (*Backend) DeletePointsTransaction(ctx context.Context, id string) error {
	panic("unimplemented")
}

// GetPointsTransaction implements db.DBackend.
func (*Backend) GetPointsTransaction(ctx context.Context, id string) (*models.PointsTransaction, error) {
	panic("unimplemented")
}

// RestorePointsTransaction implements db.DBackend.
func (*Backend) RestorePointsTransaction(ctx context.Context, id string) error {
	panic("unimplemented")
}

// UpdatePointsTransaction implements db.DBackend.
func (*Backend) UpdatePointsTransaction(ctx context.Context, pointsTransaction *models.PointsTransaction) error {
	panic("unimplemented")
}