package auth

import (
	//"errors"
	"github.com/gin-gonic/gin"
	//"errors"
	//"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/slavajs/SimpleAPI/config"
	"github.com/slavajs/SimpleAPI/internal/schemas"
	"log"
	"strings"
	"time"
)

func GenerateToken(cfg schemas.JWTConfig, login string, id int) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(cfg.TokenDuration)).Unix()
	claims["login"] = login
	claims["id"] = id
	tokenString, err := token.SignedString([]byte(cfg.SecretString))
	log.Println(err)
	return tokenString, err
}

func ParseToken(c *gin.Context) string { // TODO: add error handling
	header := c.Request.Header["Authorization"][0]
	headerSplitted := strings.Split(header, " ")
	token := headerSplitted[1]
	return token
}
func VerifyCallback(cfg schemas.JWTConfig) func(*jwt.Token) (interface{}, error) {
	return func(token *jwt.Token) (interface{}, error) {
		return []byte(cfg.SecretString), nil
	}
}

func VerifyToken(tokenString string) (bool, *jwt.Token) {
	parsedToken, err := jwt.Parse(tokenString, VerifyCallback(config.ProjectConfig.JWT))
	log.Println(err)
	return parsedToken.Valid, parsedToken
}
