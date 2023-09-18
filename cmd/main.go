package main

import (
	"log"
	"net/http"
	"time"

	"github.com/Batyrhan21/goandgin/config"
	"github.com/Batyrhan21/goandgin/controller"
	"github.com/Batyrhan21/goandgin/helper"
	"github.com/Batyrhan21/goandgin/model"
	"github.com/Batyrhan21/goandgin/repository"
	"github.com/Batyrhan21/goandgin/router"
	"github.com/Batyrhan21/goandgin/service"

	"github.com/go-playground/validator/v10"
)

func main() {

	loadConfig, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	//Database
	db := config.ConnectionDB(&loadConfig)
	validate := validator.New()

	db.Table("users").AutoMigrate(&model.Users{})
	db.Table("crypto_symbols").AutoMigrate(&model.CryptoSymbols{})

	//Init Repository
	userRepository := repository.NewUsersRepositoryImpl(db)
	cryptoSymbolsRepository := repository.NewCryptoSymbolsRepositoryImpl(db)

	//Init Service
	authenticationService := service.NewAuthenticationServiceImpl(userRepository, validate)

	//Init controller
	authenticationController := controller.NewAuthenticationController(authenticationService)
	usersController := controller.NewUsersController(userRepository)

	cryptoSymbolsController := controller.NewCryptoSymbolsController(cryptoSymbolsRepository)

	routes := router.NewRouter(userRepository, authenticationController, usersController, cryptoSymbolsController)

	server := &http.Server{
		Addr:           ":" + loadConfig.ServerPort,
		Handler:        routes,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	server_err := server.ListenAndServe()
	helper.ErrorPanic(server_err)
}
