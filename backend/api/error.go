package api

import (
	"backend/autogen"

	"github.com/labstack/echo/v4"
)

func Error500(c echo.Context) error {
	resp := autogen.GetUserMe500JSONResponse{
		Message:   autogen.MsgInternalServerError,
		ErrorCode: autogen.ErrInternalServerError,
	}
	resp.VisitGetUserMeResponse(c.Response())
	return nil
}

func Error400(c echo.Context) error {
	resp := autogen.GetCalendar400JSONResponse{
		Message:   autogen.MsgBadRequest,
		ErrorCode: autogen.ErrBadRequest,
	}
	resp.VisitGetCalendarResponse(c.Response())
	return nil
}

func ErrorNotAuthenticated(c echo.Context) error {
	resp := autogen.GetUserMe401JSONResponse{
		Message:   autogen.MsgNotAuthenticated,
		ErrorCode: autogen.ErrNotAuthenticated,
	}
	resp.VisitGetUserMeResponse(c.Response())
	return nil
}

func ErrorForbidden(c echo.Context) error {
	resp := autogen.GetUserById403JSONResponse{
		Message:   autogen.Messages(autogen.ErrForbidden),
		ErrorCode: autogen.ErrForbidden,
	}
	resp.VisitGetUserByIdResponse(c.Response())
	return nil
}

func ErrorAccNotFound(c echo.Context) error {
	resp := autogen.GetCalendarDay404JSONResponse{
		Message:   autogen.MsgUserNotFound,
		ErrorCode: autogen.ErrUserNotFound,
	}
	resp.VisitGetCalendarDayResponse(c.Response())
	return nil
}
