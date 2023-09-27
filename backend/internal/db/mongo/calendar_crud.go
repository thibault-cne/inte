package mongo

import (
	"backend/internal/models"
	"context"
)

// CreateCalendarDay implements db.DBackend.
func (*Backend) CreateCalendarDay(ctx context.Context, calendarDay *models.CalendarDay) error {
	panic("unimplemented")
}

// DeleteCalendarDay implements db.DBackend.
func (*Backend) DeleteCalendarDay(ctx context.Context, id string) error {
	panic("unimplemented")
}

// GetCalendarDay implements db.DBackend.
func (*Backend) GetCalendarDay(ctx context.Context, id string) (*models.CalendarDay, error) {
	panic("unimplemented")
}

// RestoreCalendarDay implements db.DBackend.
func (*Backend) RestoreCalendarDay(ctx context.Context, id string) error {
	panic("unimplemented")
}

// UpdateCalendarDay implements db.DBackend.
func (*Backend) UpdateCalendarDay(ctx context.Context, calendarDay *models.CalendarDay) error {
	panic("unimplemented")
}