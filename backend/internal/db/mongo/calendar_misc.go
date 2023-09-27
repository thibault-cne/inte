package mongo

import (
	"backend/internal/models"
	"context"
)

// ClearCalendar implements db.DBackend.
func (*Backend) ClearCalendar(ctx context.Context) error {
	panic("unimplemented")
}

// CountDeletedCalendarDays implements db.DBackend.
func (*Backend) CountDeletedCalendarDays(ctx context.Context, search string) (uint64, error) {
	panic("unimplemented")
}

// GetCalendarDays implements db.DBackend.
func (*Backend) GetCalendarDays(ctx context.Context, page uint64, size uint64) ([]*models.CalendarDay, error) {
	panic("unimplemented")
}

// GetDeletedCalendarDays implements db.DBackend.
func (*Backend) GetDeletedCalendarDays(ctx context.Context, page uint64, size uint64, search string) ([]*models.CalendarDay, error) {
	panic("unimplemented")
}