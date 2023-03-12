package controllers

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/slavajs/SimpleAPI/config"
	"github.com/slavajs/SimpleAPI/internal/auth"
	"github.com/slavajs/SimpleAPI/internal/db"
	"github.com/slavajs/SimpleAPI/internal/schemas"
	"net/http"
	"strconv"
)

func GetArticleByID(c *gin.Context) { // TODO: Adequate response
	var article schemas.Article
	id := c.Param("id")
	database := db.GetDB()
	statement := `SELECT * FROM articles WHERE "articleID" = ($1)`
	row := database.QueryRow(statement, id)
	err := row.Scan(&article.ArticleID, &article.Body, &article.Title, &article.AuthorID)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusOK, schemas.Response[[]schemas.Article]{
			"No articles found",
			make([]schemas.Article, 0),
		})
		return
	}
	c.JSON(http.StatusOK, schemas.Response[schemas.Article]{

		Body: article,
	})
	return
}

func GetArticles(c *gin.Context) { // TODO: handle errors and adequate response
	var articles []schemas.Article
	database := db.GetDB()
	rows, err := database.Query("SELECT * FROM articles")
	if err != nil {
		panic(err)
	}
	var flag bool = true
	for rows.Next() {
		var article schemas.Article
		flag = false
		rows.Scan(&article.ArticleID, &article.Title, &article.Body, &article.AuthorID)
		articles = append(articles, article)
	}
	if flag {
		c.JSON(http.StatusOK, schemas.Response[[]schemas.Article]{
			"No articles found",
			make([]schemas.Article, 0),
		})
		return
	}
	c.JSON(http.StatusOK, schemas.Response[[]schemas.Article]{
		Body: articles,
	})
	rows.Close()
}

func EditArticle(c *gin.Context) {
	id := c.Param("id")
	authorID, _ := c.Get("id")
	intAuthorID := int64(authorID.(float64))
	var editedArticle = &schemas.Article{}
	editedArticle.ArticleID = -1
	database := db.GetDB()
	err := c.BindJSON(editedArticle)
	searchAuthor := `SELECT "authorID" FROM articles WHERE "articleID" = ($1)`
	errQ := database.QueryRow(searchAuthor, id).Scan(&editedArticle.AuthorID)
	if errQ == sql.ErrNoRows {
		c.JSON(http.StatusForbidden, schemas.Response[[]schemas.Article]{
			Error: "Article with given id does not exist",
		})
		return
	}
	if editedArticle.AuthorID != intAuthorID {
		c.JSON(http.StatusForbidden, schemas.Response[[]schemas.Article]{
			"Permission denied",
			make([]schemas.Article, 0),
		})
		return
	}
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, schemas.Response[[]schemas.Article]{
			"Wrong entity",
			make([]schemas.Article, 0),
		})
		return
	}

	editedArticle.ArticleID, _ = strconv.ParseInt(id, 10, 64)
	statement := `UPDATE articles SET title = ($1), body = ($2) WHERE "articleID" = ($3)`
	database.Exec(statement, editedArticle.Title, editedArticle.Body, editedArticle.ArticleID)
	c.JSON(http.StatusOK, schemas.Response[string]{
		Body: "Successfully edited article",
	})
}

func RemoveArticle(c *gin.Context) { // TODO: Error if there is no id and adequate response
	var origAuthorID int64
	database := db.GetDB()
	id := c.Param("id")
	authorID, _ := c.Get("id")
	intAuthorID := int64(authorID.(float64))
	searchAuthor := `SELECT "authorID" FROM articles WHERE "articleID" = ($1)`
	err := database.QueryRow(searchAuthor, id).Scan(&origAuthorID)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusForbidden, schemas.Response[[]schemas.Article]{
			"Article with given id does not exist",
			make([]schemas.Article, 0),
		})
		return
	}
	if intAuthorID != origAuthorID {
		c.JSON(http.StatusForbidden, schemas.Response[[]schemas.Article]{
			"Permission denied",
			make([]schemas.Article, 0),
		})
		return
	}
	statement := `DELETE FROM articles WHERE "articleID" = ($1)`
	database.Exec(statement, id)
	c.JSON(http.StatusOK, schemas.Response[string]{
		Body: "Successfully deleted",
	})
}

func PostArticle(c *gin.Context) { // TODO: Adequate response
	authorID, _ := c.Get("id")
	intAuthorID := int64(authorID.(float64))
	database := db.GetDB()
	var newArticle = &schemas.Article{}
	if err := c.BindJSON(newArticle); err != nil {
		c.JSON(http.StatusUnprocessableEntity, schemas.Response[[]schemas.Article]{
			Error: "Unprocessable entity",
			Body:  make([]schemas.Article, 0),
		})
		return
	}
	newArticle.AuthorID = intAuthorID
	statement := `INSERT INTO articles (title, body, "authorID") VALUES ($1, $2, $3) RETURNING "articleID"`
	database.QueryRow(statement, newArticle.Title, newArticle.Body, newArticle.AuthorID).Scan(&newArticle.ArticleID)
	c.JSON(http.StatusOK, schemas.Response[schemas.Article]{
		Body: *newArticle,
	})
}

func RegisterUser(c *gin.Context) {
	login := c.PostForm("login")
	password := c.PostForm("password")
	var newUser = schemas.UserAuthData{login, password, 0}
	database := db.GetDB()
	if auth.CheckUserExists(login) {
		c.JSON(http.StatusConflict, schemas.Response[string]{
			Body: "User already exists",
		})
		return
	}
	hashedPassword, _ := auth.HashPassword(password)
	statement := `INSERT INTO "usersAuthData" (login, password) VALUES ($1, $2) RETURNING "userID"`
	database.QueryRow(statement, login, hashedPassword).Scan(&newUser.UserID) // TODO handle this
	c.JSON(http.StatusOK, schemas.Response[string]{
		Body: "New user registered",
	})
}

func AuthUser(c *gin.Context) {
	database := db.ConnectToDB(config.ProjectConfig.DB)
	login := c.PostForm("login")
	password := c.PostForm("password")
	var userPassword string
	var userID int
	statement := `SELECT password, "userID" from "usersAuthData" WHERE login = ($1)`
	err := database.QueryRow(statement, login).Scan(&userPassword, &userID)
	if err != nil {
		c.JSON(http.StatusUnauthorized, schemas.Response[string]{
			Body: "Wrong auth data",
		})
		return
	}
	ans, _ := auth.CheckPassword([]byte(password), []byte(userPassword))
	if ans {
		token, _ := auth.GenerateToken(config.ProjectConfig.JWT, login, userID) // generate token
		c.SetCookie("token", token, 3600, "/", "localhost", false, true)
		c.JSON(http.StatusOK, schemas.Response[string]{
			Body: "Successfully authenticated",
		})
		return
	}
	c.JSON(http.StatusUnauthorized, schemas.Response[string]{
		Body: "Wrong authdata",
	})
}
