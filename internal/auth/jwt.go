package auth

import (
	"github.com/gin-gonic/gin"
	//"encoding/json"
	"errors"
	"github.com/golang-jwt/jwt"
	//"github.com/slavajs/SimpleAPI/config"
	"github.com/slavajs/SimpleAPI/internal/schemas"

	"time"
)

func GenerateToken(cfg schemas.JWTConfig) (string, error) {
	token := jwt.New(jwt.SigningMethodEdDSA)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(cfg.TokenDuration))
	tokenString, err := token.SignedString(cfg.SecretString)
	return tokenString, err
}

func ParseToken(c *gin.Context) error { // TODO: finished this and add auth route
	token := c.Request.Header["token"]
	if token == nil {
		return errors.New("token was not provided")
	}
	//token = jwt.Parse()
}
