package api

import (
	"backend/autogen"

	"github.com/labstack/echo/v4"
)

// (DELETE /points)
func (s *Server) DeletePoints(ctx echo.Context, params autogen.DeletePointsParams) error {
	panic("unimplemented")
}

// (GET /points)
func (s *Server) GetPointsTransactions(ctx echo.Context, params autogen.GetPointsTransactionsParams) error {
	panic("unimplemented")
}

// (POST /points)
func (s *Server) AddPoints(ctx echo.Context) error {
	panic("unimplemented")
}