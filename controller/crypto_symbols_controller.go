package controller

import (
	"github.com/Batyrhan21/goandgin/data/response"
	"github.com/Batyrhan21/goandgin/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CryptoSymbolsController struct {
	symbolsRepository repository.CryptoSymbolrRepository
}

func NewCryptoSymbolsController(repository repository.CryptoSymbolrRepository) *CryptoSymbolsController {
	return &CryptoSymbolsController{symbolsRepository: repository}
}

func (controller *CryptoSymbolsController) GetCryptoSymbols(ctx *gin.Context) {
	users := controller.symbolsRepository.FindAll()
	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully fetch all crypto symbols data!",
		Data:    users,
	}

	ctx.JSON(http.StatusOK, webResponse)
}
