package api

import (
	"backend/autogen"

	"github.com/labstack/echo/v4"
)

// (DELETE /star)
func (s *Server) DeleteStarById(ctx echo.Context, params autogen.DeleteStarByIdParams) error {
	panic("unimplemented")
}

// (GET /star)
func (s *Server) GetStars(ctx echo.Context, params autogen.GetStarsParams) error {
	panic("unimplemented")
}

// (PATCH /star)
func (s *Server) EditStar(ctx echo.Context) error {
	panic("unimplemented")
}

// (POST /star)
func (s *Server) PostUserStarsById(ctx echo.Context, params autogen.PostUserStarsByIdParams) error {
	panic("unimplemented")
}

// (PATCH /star/{star_id}/upvote)
func (s *Server) PatchStarUpvote(ctx echo.Context, starId autogen.UUID) error {
	panic("unimplemented")
}

// (POST /star/{star_id}/validate)
func (s *Server) PostStarValidation(ctx echo.Context, starId autogen.UUID) error {
	panic("unimplemented")
}