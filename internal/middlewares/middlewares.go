package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/slavajs/SimpleAPI/internal/auth"
	"github.com/slavajs/SimpleAPI/internal/schemas"
	"net/http"
)

func CheckAuthorized(c *gin.Context) {
	token := auth.ParseToken(c)
	res, parsedToken := auth.VerifyToken(token)
	if !res {
		c.JSON(http.StatusUnauthorized, schemas.Response[string]{
			Error: "Wrong JSON token",
		})
		c.Abort()
		return
	}
	c.Set("login", parsedToken.Claims.(jwt.MapClaims)["login"])
	c.Set("id", parsedToken.Claims.(jwt.MapClaims)["id"])
	c.Next()
}
