package route

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"github.com/cesc1802/go_training/internal/controllers"
	"github.com/cesc1802/go_training/internal/middlewares"
)

func SetupRoutes(db *gorm.DB) {
	r := gin.Default()

	controllers.AuthController().Routes(r)
	r.Use(middlewares.TokenAuthMiddleware())
	controllers.TaskController().Routes(r)

	r.Run() // listen and serve on 8080 port
}