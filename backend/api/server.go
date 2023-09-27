package api

import (
	"backend/autogen"
	"backend/internal/config"
	"backend/internal/db"
	"os"

	"github.com/gorilla/context"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
	echoLog "github.com/labstack/gommon/log"
	middleware "github.com/neko-neko/echo-logrus/v2"
	"github.com/neko-neko/echo-logrus/v2/log"
)

type Server struct {
	db.DBackend
}

func NewServer(db db.DBackend) *Server {
	s := &Server{
		db,
	}
	return s
}

func (s *Server) getUserSess(c echo.Context) *sessions.Session {
	return c.Get("userSess").(*sessions.Session)
}

func (s *Server) getInteSess(c echo.Context) *sessions.Session {
	return c.Get("inteSess").(*sessions.Session)
}

func (s *Server) getAdminSess(c echo.Context) *sessions.Session {
	return c.Get("adminSess").(*sessions.Session)
}

func (s *Server) Serve(c *config.Config) error {
	e := echo.New()

	// Logger
	log.Logger().SetOutput(os.Stdout)
	log.Logger().SetLevel(echoLog.INFO)
	e.Logger = log.Logger()
	e.Use(middleware.Logger())

	userStore := sessions.NewCookieStore([]byte(c.ApiConfig.SessionSecret))
	inteStore := sessions.NewCookieStore([]byte(c.ApiConfig.AdminSessionSecret))
	adminStore := sessions.NewCookieStore([]byte(c.ApiConfig.AdminSessionSecret))
	
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			defer context.Clear(c.Request())
			c.Set("userStore", userStore)
			c.Set("inteStore", inteStore)
			c.Set("adminStore", adminStore)
			return next(c)
		}
	})

	// CORS
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Response().Header().Set("Access-Control-Allow-Origin", c.Request().Header.Get("Origin"))
			c.Response().Header().Set("Access-Control-Allow-Credentials", "true")
			c.Response().Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
			c.Response().Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization, Cookie, Cookies, X-Local-Token")
			return next(c)
		}
	})

	e.Use(s.AuthMiddleware)

	autogen.RegisterHandlers(e, s)

	if err := e.Start(c.ApiConfig.Port); err != nil {
		return err
	}

	return nil
}