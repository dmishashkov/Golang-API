package main

import (
	"github.com/dmishashkov/SimpleAPI/internal/controllers"
	"github.com/dmishashkov/SimpleAPI/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.POST("/auth/signup", controllers.RegisterUser)
	r.POST("/auth/signin", controllers.AuthUser)
	r.GET("/article/:id", controllers.GetArticleByID)
	r.GET("/articles", controllers.GetArticles)
	r.POST("/article", middlewares.CheckAuthorized, controllers.PostArticle)
	r.PUT("/article/:id", middlewares.CheckAuthorized, controllers.EditArticle)
	r.DELETE("/article/:id", middlewares.CheckAuthorized, controllers.RemoveArticle)
	r.Run(":5050")
}
