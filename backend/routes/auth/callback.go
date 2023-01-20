package auth

import (
	"backend/models"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yyewolf/goth"
	"github.com/yyewolf/goth/gothic"
	"golang.org/x/oauth2"
	admin "google.golang.org/api/admin/directory/v1"
	"google.golang.org/api/option"
)

func Callback(ctx *gin.Context) {
	provider, err := goth.GetProvider("google")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error 1"})
		return
	}

	value, err := gothic.GetFromSession(provider.Name(), ctx.Request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error 2"})
		return
	}

	sess, err := provider.UnmarshalSession(value)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error 3"})
		return
	}

	user, err := provider.FetchUser(sess)
	if err != nil {
		params := ctx.Request.URL.Query()
		if params.Encode() == "" && ctx.Request.Method == "POST" {
			ctx.Request.ParseForm()
			params = ctx.Request.Form
		}

		// get new token and retry fetch
		_, err = sess.Authorize(provider, params)
		if err != nil {
			fmt.Printf("%s\n", err.Error())
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error 4"})
			return
		}

		err = gothic.StoreInSession(provider.Name(), sess.Marshal(), ctx.Request, ctx.Writer)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error 5"})
			return
		}

		gu, err := provider.FetchUser(sess)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error 6"})
			return
		}

		user = gu
	}

	c := context.Background()
	oconfig := &oauth2.Config{}
	token, err := googleProvider.RefreshToken(user.RefreshToken)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Error refreshing token 7"})
		return
	}
	adminService, err := admin.NewService(c, option.WithTokenSource(oconfig.TokenSource(c, token)))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error 8"})
		return
	}

	t, err := adminService.Users.Get(user.UserID).Projection("custom").CustomFieldMask("Education").ViewType("domain_public").Do()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error 9"})
		return
	}
	edc := &Education{}
	err = json.Unmarshal(t.CustomSchemas["Education"], edc)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error 10"})
		return
	}
	// Saves to database

	temp, _ := models.GetUser(user.UserID)

	u := &models.User{
		ID:             user.UserID,
		Email:          user.Email,
		Name:           fmt.Sprintf("%s %s", user.FirstName, user.LastName),
		PromotionYear: edc.Promo,
		UserType:      temp.UserType,
	}

	err = u.Save()

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error : cannot save user"})
	}

	ctx.Redirect(http.StatusFound, redirectFront)
}
