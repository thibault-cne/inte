package db

import (
	"backend/internal/models"
	"context"
	"time"
)

type DatabaseOptions struct {
	ConnString   string
	Database     string
	QueryTimeout time.Duration
}

func NewDatabaseOptions(connString, database string, timeout time.Duration) *DatabaseOptions {
	return &DatabaseOptions{
		Database:     database,
		ConnString:   connString,
		QueryTimeout: timeout,
	}
}

type DBackend interface {
	Connect() error
	Disconnect() error

	// User's CRUD operations
	CreateUser(ctx context.Context, user *models.User) error
	GetUser(ctx context.Context, id string) (*models.User, error)
	UpdateUser(ctx context.Context, user *models.User) error
	DeleteUser(ctx context.Context, id string) error
	RestoreUser(ctx context.Context, id string) error

	// User's MISC operations
	GetUsers(ctx context.Context, page uint64, size uint64, search string) ([]*models.User, error)
	GetDeletedUsers(ctx context.Context, page uint64, size uint64, search string) ([]*models.User, error)
	CountDeletedUsers(ctx context.Context, search string) (uint64, error)

	// Star's CRUD operations
	CreateStar(ctx context.Context, star *models.Star) error
	GetStar(ctx context.Context, id string) (*models.Star, error)
	UpdateStar(ctx context.Context, star *models.Star) error
	DeleteStar(ctx context.Context, id string) error
	RestoreStar(ctx context.Context, id string) error

	// Star's MISC operations
	GetStars(ctx context.Context, page uint64, size uint64, from, to string) ([]*models.Star, error)
	GetDeletedStars(ctx context.Context, page uint64, size uint64, search string) ([]*models.Star, error)
	CountDeletedStars(ctx context.Context, search string) (uint64, error)

	// Suggestion's CRUD operations
	CreateSuggestion(ctx context.Context, suggestion *models.Suggestion) error
	GetSuggestion(ctx context.Context, id string) (*models.Suggestion, error)
	UpdateSuggestion(ctx context.Context, suggestion *models.Suggestion) error
	DeleteSuggestion(ctx context.Context, id string) error
	RestoreSuggestion(ctx context.Context, id string) error

	// Suggestion's MISC operations
	GetDeletedSuggestions(ctx context.Context, page uint64, size uint64, search string) ([]*models.Suggestion, error)
	CountDeletedSuggestions(ctx context.Context, search string) (uint64, error)

	// PointsTransaction's CRUD operations
	CreatePointsTransaction(ctx context.Context, pointsTransaction *models.PointsTransaction) error
	GetPointsTransaction(ctx context.Context, id string) (*models.PointsTransaction, error)
	UpdatePointsTransaction(ctx context.Context, pointsTransaction *models.PointsTransaction) error
	DeletePointsTransaction(ctx context.Context, id string) error
	RestorePointsTransaction(ctx context.Context, id string) error

	// PointsTransaction's MISC operations
	GetPointsTransactions(ctx context.Context, page uint64, size uint64, from, to string) ([]*models.PointsTransaction, error)
	GetDeletedPointsTransactions(ctx context.Context, page uint64, size uint64, search string) ([]*models.PointsTransaction, error)
	CountDeletedPointsTransactions(ctx context.Context, search string) (uint64, error)

	// News's CRUD operations
	CreateNews(ctx context.Context, news *models.News) error
	GetNews(ctx context.Context, id string) (*models.News, error)
	UpdateNews(ctx context.Context, news *models.News) error
	DeleteNews(ctx context.Context, id string) error
	RestoreNews(ctx context.Context, id string) error

	// News's MISC operations
	GetNewsList(ctx context.Context, page uint64, size uint64) ([]*models.News, error)
	GetDeletedNews(ctx context.Context, page uint64, size uint64, search string) ([]*models.News, error)
	CountDeletedNews(ctx context.Context, search string) (uint64, error)

	// CalendarDay's CRUD operations
	CreateCalendarDay(ctx context.Context, calendarDay *models.CalendarDay) error
	GetCalendarDay(ctx context.Context, id string) (*models.CalendarDay, error)
	UpdateCalendarDay(ctx context.Context, calendarDay *models.CalendarDay) error
	DeleteCalendarDay(ctx context.Context, id string) error
	RestoreCalendarDay(ctx context.Context, id string) error

	// CalendarDay's MISC operations
	ClearCalendar(ctx context.Context) error
	GetCalendarDays(ctx context.Context, page uint64, size uint64) ([]*models.CalendarDay, error)
	GetDeletedCalendarDays(ctx context.Context, page uint64, size uint64, search string) ([]*models.CalendarDay, error)
	CountDeletedCalendarDays(ctx context.Context, search string) (uint64, error)
}