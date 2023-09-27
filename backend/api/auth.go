package api

import (
	"backend/autogen"
	"backend/internal/config"
	"backend/internal/models"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/patrickmn/go-cache"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/oauth2"
	admin "google.golang.org/api/admin/directory/v1"
	"google.golang.org/api/option"
)

var redirectCache = cache.New(5*time.Minute, 10*time.Minute)

var scopes = []string{
	"https://www.googleapis.com/auth/userinfo.profile",
	"https://www.googleapis.com/auth/userinfo.email",
	"https://www.googleapis.com/auth/admin.directory.user.readonly",
}

type education struct {
	Promo  uint64 `json:"Promotion"`
	Spé    string `json:"Approfondissement"`
	Statut uint64 `json:"Statut"`
}

type googleUser struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	FirstName string `json:"given_name"`
	LastName  string `json:"family_name"`
	Link      string `json:"link"`
	Picture   string `json:"picture"`
}

func DefaultRedirect(c echo.Context) error {
	conf := config.GetConfig()
	return c.Redirect(http.StatusPermanentRedirect, conf.BaseURI.Front+"/login")
}

// (GET /auth/callback)
func (s *Server) Callback(ctx echo.Context, params autogen.CallbackParams) error {
	notRegisterd := false

	conf := config.GetConfig()

	// Get token from Google
	oauth2Config := oauth2.Config{
		ClientID:     conf.OauthConfig.GoogleClientID,
		ClientSecret: conf.OauthConfig.GoogleClientSecret,
		RedirectURL:  fmt.Sprintf("%s/auth/callback", conf.BaseURI.Back),
		Scopes:       scopes,
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://accounts.google.com/o/oauth2/auth",
			TokenURL: "https://oauth2.googleapis.com/token",
		},
	}

	token, err := oauth2Config.Exchange(ctx.Request().Context(), params.Code)
	if err != nil {
		logrus.Error(err)
		return DefaultRedirect(ctx)
	}

	// Get user from Google
	client := oauth2Config.Client(ctx.Request().Context(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		logrus.Error(err)
		return DefaultRedirect(ctx)
	}
	defer resp.Body.Close()

	usr := &googleUser{}
	err = json.NewDecoder(resp.Body).Decode(usr)
	if err != nil {
		logrus.Error(err)
		return DefaultRedirect(ctx)
	}

	account, err := s.DBackend.GetUser(ctx.Request().Context(), usr.ID)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			logrus.Error(err)
			return DefaultRedirect(ctx)
		}
		
		notRegisterd = true
		account = &models.User{
			User: autogen.User{
				Points: 0,
				Role: autogen.UserFreshman,
			},
		}
	}

	adminService, err := admin.NewService(ctx.Request().Context(), option.WithTokenSource(oauth2Config.TokenSource(ctx.Request().Context(), token)))
	if err != nil {
		logrus.Error(err)
		return DefaultRedirect(ctx)
	}

	t, err := adminService.Users.Get(usr.ID).Projection("custom").CustomFieldMask("Education").ViewType("domain_public").Do()
	if err != nil {
		logrus.Error(err)
		return DefaultRedirect(ctx)
	}
	edc := &education{}
	err = json.Unmarshal(t.CustomSchemas["Education"], edc)
	if err != nil {
		logrus.Error(err)
		return DefaultRedirect(ctx)
	}

	logrus.Info(fmt.Sprintf("%+v", account))

	account.FirstName = usr.FirstName
	account.LastName = usr.LastName
	account.EmailAddress = usr.Email
	account.GoogleId = usr.ID
	account.GooglePicture = &usr.Picture

	if notRegisterd {

		err = s.DBackend.CreateUser(ctx.Request().Context(), account)
		if err != nil {
			logrus.Error(err)
			return DefaultRedirect(ctx)
		}
	} else {
		err = s.DBackend.UpdateUser(ctx.Request().Context(), account)
		if err != nil {
			logrus.Error(err)
			return DefaultRedirect(ctx)
		}
	}

	// TODO: add the websocket
	// BroadcastToRoom(accountID.(string), []byte("connected"))

	r, found := redirectCache.Get(params.State)
	if !found {
		return DefaultRedirect(ctx)
	}
	redirectCache.Delete(params.State)

	s.SetCookie(ctx, account)
	return ctx.Redirect(http.StatusPermanentRedirect, r.(string))
}

// (GET /auth/login)
func (s *Server) Connect(ctx echo.Context, p autogen.ConnectParams) error {
	conf := config.GetConfig()

	// Get ?r=
	/* rel := *p.R

	// Check if it's a safe redirect (TODO: check if this is correct)
	switch rel {
	case "admin":
		rel = conf.BaseURI.Front + "/admin"
	} */
	// Init OAuth2 flow with Google
	oauth2Config := oauth2.Config{
		ClientID:     conf.OauthConfig.GoogleClientID,
		ClientSecret: conf.OauthConfig.GoogleClientSecret,
		RedirectURL:  fmt.Sprintf("%s/auth/callback", conf.BaseURI.Back),
		Scopes:       scopes,
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://accounts.google.com/o/oauth2/auth",
			TokenURL: "https://oauth2.googleapis.com/token",
		},
	}

	// state is not nonce
	state := uuid.NewString()

	redirectCache.Set(state, conf.BaseURI.Front, cache.DefaultExpiration)

	hostDomainOption := oauth2.SetAuthURLParam("hd", "telecomnancy.net")
	// Redirect to Google
	url := oauth2Config.AuthCodeURL(state, oauth2.AccessTypeOffline, hostDomainOption)

	return ctx.Redirect(http.StatusTemporaryRedirect, url)
}

// (GET /auth/logout)
func (s *Server) Logout(ctx echo.Context) error {
	s.RemoveCookies(ctx)

	autogen.Logout204Response{}.VisitLogoutResponse(ctx.Response())
	return nil
}