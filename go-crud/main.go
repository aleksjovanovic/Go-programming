package main

import (
	"net/http"

	"github.com/aleksjovanovic/go-crud/configs"
	"github.com/aleksjovanovic/go-crud/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	configs.ConnectToDB()
}

func main() {

	router := gin.Default()

	routes.TeamRouter(router)

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Go server is running ...",
		})
	})
	router.Run()
}
