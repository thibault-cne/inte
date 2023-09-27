package api

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *Server) AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		v := c.Get("userStore")
		userStore, ok := v.(sessions.Store)
		if !ok {
			return echo.NewHTTPError(500, "userStore not found")
		}
		v = c.Get("inteStore")
		inteStore, ok := v.(sessions.Store)
		if !ok {
			return echo.NewHTTPError(500, "inteStore not found")
		}
		v = c.Get("adminStore")
		adminStore, ok := v.(sessions.Store)
		if !ok {
			return echo.NewHTTPError(500, "adminStore not found")
		}

		userSess, err := userStore.Get(c.Request(), "SESS")
		if err != nil {
			return echo.NewHTTPError(500, "session not found")
		}
		inteSess, err := inteStore.Get(c.Request(), "INTE_SESS")
		if err != nil {
			return echo.NewHTTPError(500, "session not found")
		}
		adminSess, err := adminStore.Get(c.Request(), "ADMIN_SESS")
		if err != nil {
			return echo.NewHTTPError(500, "session not found")
		}

		c.Set("userSess", userSess)
		c.Set("inteSess", inteSess)
		c.Set("adminSess", adminSess)

		c.Set("userLogged", false)
		c.Set("inteLogged", false)
		c.Set("adminLogged", false)

		// Get user's account from cookie
		accountID, ok := userSess.Values["account_id"].(string)
		if ok {
			// Get account from database
			account, err := s.DBackend.GetUser(c.Request().Context(), accountID)
			if err != nil {
				if err == mongo.ErrNoDocuments {
					// Delete cookie
					userSess.Options.MaxAge = -1
					userSess.Save(c.Request(), c.Response())
					return ErrorAccNotFound(c)
				}
				logrus.Error(err)
				return Error500(c)
			}

			c.Set("userLogged", true)
			c.Set("userAccountID", accountID)
			c.Set("userAccount", account)
		}

		// Get admin' inte account from cookie
		inteID, ok := inteSess.Values["inte_account_id"].(string)
		if ok {
			// Get account from database
			account, err := s.DBackend.GetUser(c.Request().Context(), inteID)
			if err != nil {
				if err == mongo.ErrNoDocuments {
					// Delete cookie
					userSess.Options.MaxAge = -1
					userSess.Save(c.Request(), c.Response())
					return ErrorAccNotFound(c)
				}
				logrus.Error(err)
				return Error500(c)
			}

			c.Set("inteLogged", true)
			c.Set("inteAccountID", accountID)
			c.Set("inteAccount", account)
		}

		// Get admin account from cookie
		adminId, ok := adminSess.Values["admin_account_id"].(string)
		if ok {
			// Get account from database
			account, err := s.DBackend.GetUser(c.Request().Context(), adminId)
			if err != nil {
				if err == mongo.ErrNoDocuments {
					// Delete cookie
					adminSess.Options.MaxAge = -1
					adminSess.Save(c.Request(), c.Response())
					return ErrorAccNotFound(c)
				}
				logrus.Error(err)
				return Error500(c)
			}

			c.Set("adminLogged", true)
			c.Set("adminAccountID", adminId)
			c.Set("adminAccount", account)
			// c.Set("adminAccountRole", account.Role)
		}

		return next(c)
	}
}