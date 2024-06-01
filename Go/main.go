package main

import (
	"fmt"
	"net/http"

	"github.com/yasir16/simple-api-rbac/config"
	"github.com/yasir16/simple-api-rbac/controller"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	setupServer().Run()
}
func setupServer() *gin.Engine {
	db := config.DBInit()
	inDB := &controller.InDB{DB: db}

	router := gin.Default()

	router.GET("/books", inDB.GetBooks)
	router.GET("/books/:id", inDB.GetBook)
	router.PUT("/books/:id", auth, inDB.UpdateBook)
	router.DELETE("/books/:id", auth, inDB.DeleteBook)
	router.POST("/books", auth, inDB.CreateBook)

	return router
}
func auth(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	authPermission := make(map[string]bool)
	authPermission["adminadmin"] = true
	// if token.Valid && err == nil {
	if authPermission[token] {
		fmt.Println("token verified")
	} else {
		result := gin.H{
			"message": "not authorized",
			"error":   "not authorized",
		}
		c.JSON(http.StatusUnauthorized, result)
		c.Abort()
	}
}
