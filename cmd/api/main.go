package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/slavajs/SimpleAPI/internal/controllers"
	"github.com/slavajs/SimpleAPI/internal/middlewares"
	"log"
)

func main() {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.Default())
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
	log.Print("[main] Successfully started server")
}
