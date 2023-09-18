package model

type CryptoSymbols struct {
	Id     int    `gorm:"type:int;primary_key"`
	Symbol string `gorm:"type:varchar(255);not null"`
}
