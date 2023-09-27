package api

import (
	"backend/autogen"

	"github.com/labstack/echo/v4"
)

// (DELETE /suggestion)
func (s *Server) DeleteSuggestion(ctx echo.Context, params autogen.DeleteSuggestionParams) error {
	panic("unimplemented")
}

// (GET /suggestion)
func (s *Server) GetSuggestions(ctx echo.Context, params autogen.GetSuggestionsParams) error {
	panic("unimplemented")
}

// (PATCH /suggestion)
func (s *Server) PatchSuggestion(ctx echo.Context, params autogen.PatchSuggestionParams) error {
	panic("unimplemented")
}

// (POST /suggestion)
func (s *Server) PostSuggestion(ctx echo.Context) error {
	panic("unimplemented")
}