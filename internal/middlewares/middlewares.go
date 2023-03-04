package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/slavajs/SimpleAPI/config"
	"github.com/slavajs/SimpleAPI/internal/db"
	"github.com/slavajs/SimpleAPI/internal/schemas"
	"log"
	"net/http"
)

func GetArticleByID(c *gin.Context) { // TODO: Adequate response if there is
	var article schemas.Article
	id := c.Param("id")
	fmt.Println(id)
	database := db.ConnectToDB(config.ProjectConfig.DB)
	statement := `SELECT * FROM articles WHERE id = ($1)`
	row := database.QueryRow(statement, id)
	row.Scan(&article.ID, &article.Body, &article.Title)
	c.JSON(http.StatusOK, article)
	database.Close()
}

func GetArticles(c *gin.Context) { // TODO: handle errors and adequate response
	var articles []schemas.Article
	database := db.ConnectToDB(config.ProjectConfig.DB)
	rows, err := database.Query("SELECT * FROM articles")
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var article schemas.Article
		err := rows.Scan(&article.ID, &article.Title, &article.Body)
		if err != nil {
			log.Println(err)
		}
		articles = append(articles, article)
	}
	c.JSON(http.StatusOK, articles)
	rows.Close()
	database.Close()
}

func EditArticle(c *gin.Context) {
	
}

func RemoveArticle(c *gin.Context) { // TODO: Error if there is no id and adequate response
	database := db.ConnectToDB(config.ProjectConfig.DB)
	id := c.Query("id")
	statement := `DELETE FROM articles WHERE id = ($1)`
	database.Exec(statement, id)
	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.WriteString("Successfully deleted article")
	database.Close()
}

func PostArticle(c *gin.Context) { // TODO: Adequate response
	database := db.ConnectToDB(config.ProjectConfig.DB)
	var newArticle = &schemas.Article{}
	if c.BindJSON(newArticle) != nil {
		c.Writer.WriteHeader(http.StatusUnprocessableEntity)
		c.Writer.WriteString("Bad entity")
		return
	}
	statement := "INSERT INTO articles (title, body) VALUES ($1, $2) RETURNING id"
	database.QueryRow(statement, newArticle.Title, newArticle.Body).Scan(&newArticle.ID)
	c.JSON(http.StatusOK, newArticle)
	database.Close()
}
