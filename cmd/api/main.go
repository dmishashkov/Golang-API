package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/slavajs/SimpleAPI/internal/middlewares"
	"log"
)

func main() {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.Default())
	//authorized := r.Group("/")
	r.GET("/article/:id", middlewares.GetArticleByID)
	r.GET("/articles", middlewares.GetArticles)
	r.POST("/article", middlewares.PostArticle)
	r.PUT("/article/:id", middlewares.EditArticle)
	r.DELETE("/article/:id", middlewares.RemoveArticle)
	r.Run(":5050")
	log.Print("[main] Successfully started server")
}
