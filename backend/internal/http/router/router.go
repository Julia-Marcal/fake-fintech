package router

import (
	"time"

	middlewares "github.com/Julia-Marcal/fake-fintech/helpers/middlewares"
	acoes "github.com/Julia-Marcal/fake-fintech/internal/http/controllers/acoes"
	users "github.com/Julia-Marcal/fake-fintech/internal/http/controllers/users"
	wallet "github.com/Julia-Marcal/fake-fintech/internal/http/controllers/wallet"
	wallet_acoes "github.com/Julia-Marcal/fake-fintech/internal/http/controllers/wallet_acoes"

	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/cors"
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

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"POST", "GET", "PUT", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization", "Accept", "User-Agent", "Cache-Control", "Pragma"}
	config.ExposeHeaders = []string{"Content-Length"}
	config.AllowCredentials = true
	config.MaxAge = 12 * time.Hour

	router.Use(cors.New(config))

	rateLimiter := middlewares.RateLimiting()

	api := router.Group("/api")
	{
		api.POST("/login", rateLimiter, users.GenerateToken)
		api.GET("/metrics", rateLimiter, middlewares.PrometheusHandler())
		api.POST("/users", rateLimiter, users.CreateUser)

		authorized := api.Group("/v1")
		authorized.Use(middlewares.Auth())
		{
			userGroup := authorized.Group("/user")
			{
				userGroup.GET("/:id_user", rateLimiter, middlewares.RoleBasedAccess("admin"), users.GetUser)
				userGroup.PATCH("/:id_user", rateLimiter, middlewares.RoleBasedAccess("admin"), users.UpdateUser)
				userGroup.GET("/total_amount/:id_user", rateLimiter, middlewares.RoleBasedAccess("admin"), users.TotalAmount)
			}

			usersGroup := authorized.Group("/users")
			{
				usersGroup.GET("/*limit", rateLimiter, users.GetAllUsers)
				usersGroup.DELETE("/:id", rateLimiter, users.DeleteUser)
			}

			walletGroup := authorized.Group("/wallet")
			{
				walletGroup.POST("/", rateLimiter, wallet.CreateWallet)
				walletGroup.GET("/:id_wallet", rateLimiter, wallet.GetWallet)
				walletGroup.GET("/all_wallets/:id_user", rateLimiter, wallet.GetAllWallets)
			}

			acoesGroup := authorized.Group("/acoes")
			{
				acoesGroup.POST("/", rateLimiter, acoes.CreateAcoes)
				acoesGroup.GET("/:id_acao", rateLimiter, acoes.GetAcoes)
			}

			walletAcoesGroup := authorized.Group("/wallet_acoes")
			{
				walletAcoesGroup.GET("/:id_wallet", rateLimiter, wallet_acoes.GetWalletAcoes)
			}
		}
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run()
}
