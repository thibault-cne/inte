package api

import (
	"backend/autogen"

	"github.com/labstack/echo/v4"
)

// (GET /user)
func (s *Server) GetUsers(ctx echo.Context, params autogen.GetUsersParams) error {
	panic("unimplemented")
}

// (GET /user/me)
func (s *Server) GetUserMe(ctx echo.Context) error {
	account, err := MustGetUser(ctx)
	if err != nil {
		return nil
	}

	// Return account
	resp := autogen.GetUserMe200JSONResponse{
		User: &account.User,
	}
	resp.VisitGetUserMeResponse(ctx.Response())
	return nil
}

// (PATCH /user/me)
func (s *Server) PatchUserMe(ctx echo.Context) error {
	panic("unimplemented")
}

// (GET /user/{user_id})
func (s *Server) GetUserById(ctx echo.Context, userId autogen.GoogleId) error {
	panic("unimplemented")
}

// (PATCH /user/{user_id})
func (s *Server) PatchUserById(ctx echo.Context, userId autogen.GoogleId) error {
	panic("unimplemented")
}