package api

import (
	"backend/autogen"

	"github.com/labstack/echo/v4"
)

// (GET /team)
func (s *Server) GetTeam(ctx echo.Context) error {
	panic("unimplemented")
}

// (POST /team)
func (s *Server) PostTeam(ctx echo.Context) error {
	panic("unimplemented")
}

// (DELETE /team/{id})
func (s *Server) DeleteTeam(ctx echo.Context, id autogen.UUID) error {
	panic("unimplemented")
}

// (GET /team/{id})
func (s *Server) GetTeamById(ctx echo.Context, id autogen.UUID) error {
	panic("unimplemented")
}

// (POST /team/{id})
func (s *Server) AddUserTeam(ctx echo.Context, id autogen.UUID) error {
	panic("unimplemented")
}