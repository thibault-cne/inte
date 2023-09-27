package api

import (
	"backend/autogen"

	"github.com/labstack/echo/v4"
)

// (DELETE /calendar)
func (s *Server) DeleteCalendar(ctx echo.Context) error {
	panic("unimplemented")
}

// (GET /calendar)
func (s *Server) GetCalendar(ctx echo.Context, params autogen.GetCalendarParams) error {
	panic("unimplemented")
}

// (POST /calendar)
func (s *Server) AppendCalendarDay(ctx echo.Context) error {
	panic("unimplemented")
}

// (DELETE /calendar/{day})
func (s *Server) DeleteCalendarDay(ctx echo.Context, day int) error {
	panic("unimplemented")
}

// (GET /calendar/{day})
func (s *Server) GetCalendarDay(ctx echo.Context, day int) error {
	panic("unimplemented")
}

// (POST /calendar/{day})
func (s *Server) PostCalendarDay(ctx echo.Context, day int) error {
	panic("unimplemented")
}