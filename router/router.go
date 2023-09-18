package router

import (
	"github.com/Batyrhan21/goandgin/controller"
	//"github.com/Batyrhan21/goandgin/middleware"
	"github.com/Batyrhan21/goandgin/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(userRepository repository.UsersRepository, authenticationController *controller.AuthenticationController, usersController *controller.UserController, cryptoSymbolscontroller *controller.CryptoSymbolsController) *gin.Engine {
	service := gin.Default()

	service.GET("", func(context *gin.Context) {
		context.JSON(http.StatusOK, "welcome home")
	})

	service.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	router := service.Group("/api")
	authenticationRouter := router.Group("/authentication")
	authenticationRouter.POST("/register", authenticationController.Register)
	authenticationRouter.POST("/login", authenticationController.Login)

	//usersRouter := router.Group("/users")
	//usersRouter.GET("", middleware.DeserializeUser(userRepository), usersController.GetUsers)

	cryptoRouter := router.Group("/crypto")
	cryptoRouter.GET("/list", cryptoSymbolscontroller.GetCryptoSymbols)

	return service
}
