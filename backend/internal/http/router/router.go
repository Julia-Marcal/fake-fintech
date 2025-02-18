package router

import (
	middlewares "github.com/Julia-Marcal/fake-fintech/helpers/middlewares"
	users "github.com/Julia-Marcal/fake-fintech/internal/http/controllers/users"
	wallet "github.com/Julia-Marcal/fake-fintech/internal/http/controllers/wallet"

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
		api.POST("/login", rateLimiter, users.GenerateToken)
		api.GET("/metrics", rateLimiter, middlewares.PrometheusHandler())
		api.POST("/users", rateLimiter, users.CreateUser)

		authorized := api.Group("/v1/").Use(middlewares.Auth())
		{
			authorized.GET("user/", rateLimiter, users.GetUser)
			authorized.GET("users/*limit", rateLimiter, users.GetAllUsers)
			authorized.DELETE("users/:id", rateLimiter, users.DeleteUser)

			authorized.POST("/wallet/", rateLimiter, wallet.CreateWallet)
			authorized.GET("/wallet/:id_wallet", rateLimiter, wallet.GetWallet)
			authorized.GET("/all_wallets/:id_user", rateLimiter, wallet.GetAllWallets)
		}
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run()
}
