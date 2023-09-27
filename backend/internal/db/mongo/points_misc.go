package mongo

import (
	"backend/internal/models"
	"context"
)

// CountDeletedPointsTransactions implements db.DBackend.
func (*Backend) CountDeletedPointsTransactions(ctx context.Context, search string) (uint64, error) {
	panic("unimplemented")
}

// GetDeletedPointsTransactions implements db.DBackend.
func (*Backend) GetDeletedPointsTransactions(ctx context.Context, page uint64, size uint64, search string) ([]*models.PointsTransaction, error) {
	panic("unimplemented")
}

// GetPointsTransactions implements db.DBackend.
func (*Backend) GetPointsTransactions(ctx context.Context, page uint64, size uint64, from string, to string) ([]*models.PointsTransaction, error) {
	panic("unimplemented")
}