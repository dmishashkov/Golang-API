package main

import (
	"github.com/gin-gonic/gin"
	"github.com/slavajs/SimpleAPI/internal/controllers"
	"github.com/slavajs/SimpleAPI/internal/middlewares"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	//r.Use(gin.Logger())
	//r.Use(gin.Recovery())
	//authorized := r.Group("/")
	r.POST("/auth/signup", controllers.RegisterUser)
	r.POST("/auth/signin", controllers.AuthUser)
	r.GET("/article/:id", controllers.GetArticleByID)
	r.GET("/articles", controllers.GetArticles)
	r.POST("/article", middlewares.CheckAuthorized, controllers.PostArticle)
	r.PUT("/article/:id", middlewares.CheckAuthorized, controllers.EditArticle)
	r.DELETE("/article/:id", middlewares.CheckAuthorized, controllers.RemoveArticle)
	r.GET("/test", middlewares.CheckAuthorized)
	r.Run(":5050")
}
