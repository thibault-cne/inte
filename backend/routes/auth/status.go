package auth

import (
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yyewolf/goth"
	"github.com/yyewolf/goth/gothic"
	"github.com/yyewolf/goth/providers/google"
)

func Status(ctx *gin.Context) {
	// try to get the user without re-authenticating
	provider, err := goth.GetProvider("google")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	values := ctx.Request.URL.Query()
	values.Add("provider", "google")
	ctx.Request.URL.RawQuery = values.Encode()

	value, err := gothic.GetFromSession(provider.Name(), ctx.Request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"logged": false})
		return
	}

	sess, err := provider.UnmarshalSession(value)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"logged": false})
		return
	}

	user, err := provider.FetchUser(sess)
	if err != nil {
		s := sess.(*google.Session)

		token, err := provider.RefreshToken(s.RefreshToken)
		if err == nil {
			s.AccessToken = token.AccessToken
			s.RefreshToken = token.RefreshToken
			gothic.StoreInSession(provider.Name(), sess.Marshal(), ctx.Request, ctx.Writer)
		}

		user, err = provider.FetchUser(sess)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"logged": false})
			return
		}

		u, err := models.GetUser(user.UserID)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"logged": false})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"user": u, "logged": true})
		return
	}

	token, err := provider.RefreshToken(user.RefreshToken)
	if err == nil {
		user.AccessToken = token.AccessToken
		user.RefreshToken = token.RefreshToken
		gothic.StoreInSession(provider.Name(), sess.Marshal(), ctx.Request, ctx.Writer)
	}

	u, err := models.GetUser(user.UserID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"logged": false})
		return
	}

	// Get user stars
	u.GetStars()

	ctx.JSON(http.StatusOK, gin.H{"user": u, "logged": true})
}
