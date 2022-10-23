package auth

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"github.com/yyewolf/goth"
	"github.com/yyewolf/goth/gothic"
	"github.com/yyewolf/goth/providers/google"
	og "golang.org/x/oauth2/google"
	"google.golang.org/api/people/v1"

	_ "github.com/joho/godotenv/autoload"
)

var (
	googleProvider *google.Provider
	googleFile     = os.Getenv("G_CLIENT")
	callbackURL    = os.Getenv("G_REDIRECT_URI")
	redirectFront  = os.Getenv("REDIRECT_FRONT")
	cookieKey      = os.Getenv("COOKIE_KEY")
)

var scopes = []string{
	"https://www.googleapis.com/auth/userinfo.profile",
	"https://www.googleapis.com/auth/userinfo.email",
	"https://www.googleapis.com/auth/admin.directory.user.readonly",
}

type Education struct {
	Promo  int    `json:"Promotion"`
	Spe    string `json:"Approfondissement"`
	Statut int    `json:"Statut"`
}

func init() {
	// On récupère les tokens d'accès google
	b, err := os.ReadFile(googleFile)
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}
	s, err := og.ConfigFromJSON(b, people.ContactsReadonlyScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}

	maxAge := 86400 * 30 // 30 days cookie

	store := sessions.NewCookieStore([]byte(cookieKey))
	store.MaxAge(maxAge)
	store.Options.Path = "/"

	gothic.Store = store

	googleProvider = google.New(s.ClientID, s.ClientSecret, fmt.Sprintf("%sapi/auth/callback", callbackURL), scopes...)
	googleProvider.SetHostedDomain("telecomnancy.net")
	googleProvider.SetPrompt("consent")
	googleProvider.SetAccessType("offline")
	goth.UseProviders(
		googleProvider,
	)
}

func AuthRoutes(path *gin.RouterGroup) {
	subpath := path.Group("/auth")
	subpath.GET("/login", Login)
	subpath.GET("/logout", Logout)
	subpath.GET("/status", Status)
	subpath.GET("/callback", Callback)
}
