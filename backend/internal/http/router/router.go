package router

import (
	middlewares "github.com/Julia-Marcal/reusable-api/helpers/middlewares"
	controllers "github.com/Julia-Marcal/reusable-api/internal/http/controllers"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

func StartRouter() {
	router := gin.Default()

	rateLimiter := middlewares.RateLimiting()

	api := router.Group("/api")
	{
		api.POST("/login", rateLimiter, controllers.GenerateToken)
		api.GET("/metrics", rateLimiter, middlewares.PrometheusHandler())
		api.POST("/users", rateLimiter, controllers.CreateUser)

		authorized := api.Group("/v1/").Use(middlewares.Auth())
		{
			authorized.GET("user/", rateLimiter, controllers.GetUser)
			authorized.GET("users/", rateLimiter, controllers.GetAllUsers)
			authorized.DELETE("users/:id", rateLimiter, controllers.DeleteUser)
		}
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run()
}
