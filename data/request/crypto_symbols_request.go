package request

type UpdateCryptoSymbolsRequest struct {
	Id     int    `validate:"required"`
	Symbol string `validate:"required,max=200,min=2" json:"symbol"`
}
