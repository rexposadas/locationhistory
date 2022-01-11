package server

import (
	"github.com/gin-gonic/gin"
	"github.com/rexposadas/locationhistory/controllers"
	"github.com/rexposadas/locationhistory/middleware"
	"github.com/rexposadas/locationhistory/models"
)

// Setup .
func Setup(expiry int) *gin.Engine {
	db := models.NewOrderDB()

	r := gin.Default()
	r.Use(middleware.SetContext(db, expiry))

	v1 := r.Group("/location")
	{
		v1.POST("/:order_id/now", controllers.Now)
		v1.GET("/:order_id", controllers.List)
		v1.DELETE("/", controllers.Delete)
	}

	return r
}
