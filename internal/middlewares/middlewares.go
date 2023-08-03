package middlewares

import (
	"github.com/dmishashkov/SimpleAPI/internal/auth"
	"github.com/dmishashkov/SimpleAPI/internal/schemas"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
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
