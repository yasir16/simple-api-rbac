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

	router.GET("/books", auth, inDB.GetBooks)
	router.GET("/books/:id", auth, inDB.GetUser)
	router.PUT("/books/:id", auth, inDB.UpdateUser)
	router.DELETE("/books/:id", auth, inDB.DeleteUser)
	router.POST("/books", auth, inDB.CreateUser)

	return router
}
func auth(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")

	// if token.Valid && err == nil {
	if token == "initestyasir" {
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
