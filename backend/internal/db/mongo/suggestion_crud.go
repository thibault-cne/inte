package mongo

import (
	"backend/internal/models"
	"context"
)

// CreateSuggestion implements db.DBackend.
func (*Backend) CreateSuggestion(ctx context.Context, suggestion *models.Suggestion) error {
	panic("unimplemented")
}

// DeleteSuggestion implements db.DBackend.
func (*Backend) DeleteSuggestion(ctx context.Context, id string) error {
	panic("unimplemented")
}

// GetSuggestion implements db.DBackend.
func (*Backend) GetSuggestion(ctx context.Context, id string) (*models.Suggestion, error) {
	panic("unimplemented")
}

// RestoreSuggestion implements db.DBackend.
func (*Backend) RestoreSuggestion(ctx context.Context, id string) error {
	panic("unimplemented")
}

// UpdateSuggestion implements db.DBackend.
func (*Backend) UpdateSuggestion(ctx context.Context, suggestion *models.Suggestion) error {
	panic("unimplemented")
}