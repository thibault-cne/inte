package api

import (
	"backend/autogen"

	"github.com/labstack/echo/v4"
)

// (GET /mentoring)
func (s *Server) GetMentoringProcess(ctx echo.Context) error {
	panic("unimplemented")
}

// (PATCH /mentoring)
func (s *Server) ToggleMentoringProcess(ctx echo.Context) error {
	panic("unimplemented")
}

// (GET /mentoring/freshman)
func (s *Server) GetMentoringFreshmans(ctx echo.Context, params autogen.GetMentoringFreshmansParams) error {
	panic("unimplemented")
}

// (POST /mentoring/step)
func (s *Server) StartAssociation(ctx echo.Context) error {
	panic("unimplemented")
}

// (GET /mentoring/wish)
func (s *Server) GetWish(ctx echo.Context) error {
	panic("unimplemented")
}

// (POST /mentoring/wish)
func (s *Server) AddWish(ctx echo.Context) error {
	panic("unimplemented")
}