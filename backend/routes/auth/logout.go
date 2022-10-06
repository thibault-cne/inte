package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yyewolf/goth/gothic"
)

func Logout(ctx *gin.Context) {
	gothic.Logout(ctx.Writer, ctx.Request)
	http.Redirect(ctx.Writer, ctx.Request, redirectFront, http.StatusTemporaryRedirect)

	ctx.JSON(http.StatusOK, gin.H{})
}
