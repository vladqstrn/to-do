package app

import (
	"github.com/gin-gonic/gin"
	"github.com/vladqstrn/to-do/internal/handlers"
)

func Routes(incRoutes *gin.Engine) {
	incRoutes.POST("/create", handlers.Create)
	incRoutes.GET("/read", handlers.Read)
	incRoutes.PUT("/update", handlers.Update)
	incRoutes.DELETE("/delete", handlers.Delete)
}
