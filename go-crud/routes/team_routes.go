package routes

import (
	"github.com/aleksjovanovic/go-crud/controllers"
	"github.com/gin-gonic/gin"
)

func TeamRouter(router *gin.Engine) {

	routes := router.Group("api/v1/standings")
	routes.POST("", controllers.TeamCreate)
	routes.GET("", controllers.TeamGet)
	routes.GET("/:id", controllers.TeamGetById)
	routes.PUT("/:id", controllers.TeamUpdate)
	routes.DELETE("/:id", controllers.TeamDelete)

}
