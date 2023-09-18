package repository

import (
	"github.com/Batyrhan21/goandgin/data/request"
	"github.com/Batyrhan21/goandgin/helper"
	"github.com/Batyrhan21/goandgin/model"
	"gorm.io/gorm"
)

type CryptoSymbolsRepositoryImpl struct {
	Db *gorm.DB
}

func (c *CryptoSymbolsRepositoryImpl) Save(symbols model.CryptoSymbols) {
	result := c.Db.Create(&symbols)
	helper.ErrorPanic(result.Error)
}

func (c *CryptoSymbolsRepositoryImpl) Update(symbols model.CryptoSymbols) {
	var updateSymbols = request.UpdateCryptoSymbolsRequest{
		Id:     symbols.Id,
		Symbol: symbols.Symbol,
	}
	result := c.Db.Model(&symbols).Updates(updateSymbols)
	helper.ErrorPanic(result.Error)
}

func (c *CryptoSymbolsRepositoryImpl) Delete(symbolsId int) {
	var symbols model.CryptoSymbols
	result := c.Db.Where("id = ?", symbolsId).Delete(&symbols)
	helper.ErrorPanic(result.Error)
}

func (c *CryptoSymbolsRepositoryImpl) FindAll() []model.CryptoSymbols {
	var symbols []model.CryptoSymbols
	results := c.Db.Find(&symbols)
	helper.ErrorPanic(results.Error)
	return symbols
}

func NewCryptoSymbolsRepositoryImpl(Db *gorm.DB) CryptoSymbolrRepository {
	return &CryptoSymbolsRepositoryImpl{Db: Db}
}
