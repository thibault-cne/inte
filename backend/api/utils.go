package api

import (
	"backend/internal/models"
	"errors"

	"github.com/labstack/echo/v4"
)

func (s *Server) SetCookie(c echo.Context, account *models.User) {
	sess := s.getUserSess(c)
	sess.Options.MaxAge = 60 * 60 * 24 * 7 // 1 week
	sess.Options.HttpOnly = true
	sess.Options.Secure = true
	sess.Values["account_id"] = account.User.GoogleId
	sess.Save(c.Request(), c.Response())

	if account.IsInteAdmin() {
		sess := s.getInteSess(c)
		sess.Options.MaxAge = 60 * 60 * 24 * 7 // 1 week
		sess.Options.HttpOnly = true
		sess.Options.Secure = true
		sess.Values["inte_account_id"] = account.User.GoogleId
		sess.Save(c.Request(), c.Response())
	}

	if account.IsAdmin() {
		sess := s.getAdminSess(c)
		sess.Options.MaxAge = 60 * 60 * 24 * 7 // 1 week
		sess.Options.HttpOnly = true
		sess.Options.Secure = true
		sess.Values["admin_account_id"] = account.User.GoogleId
		sess.Save(c.Request(), c.Response())
	}
}

func (s *Server) RemoveCookies(c echo.Context) {
	sess := s.getUserSess(c)
	sess.Options.MaxAge = -1
	sess.Save(c.Request(), c.Response())

	sess = s.getInteSess(c)
	sess.Options.MaxAge = -1
	sess.Save(c.Request(), c.Response())

	sess = s.getAdminSess(c)
	sess.Options.MaxAge = -1
	sess.Save(c.Request(), c.Response())
}

func MustGetUser(c echo.Context) (*models.User, error) {
	logged := c.Get("userLogged").(bool)
	if !logged {
		ErrorNotAuthenticated(c)
		return nil, errors.New("not authenticated")
	}

	account := c.Get("userAccount").(*models.User)
	return account, nil
}