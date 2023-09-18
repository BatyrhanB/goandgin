package repository

import "github.com/Batyrhan21/goandgin/model"

type CryptoSymbolrRepository interface {
	Save(symbols model.CryptoSymbols)
	Update(symbols model.CryptoSymbols)
	Delete(symbolsId int)
	FindAll() []model.CryptoSymbols
}
