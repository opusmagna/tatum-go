package request

type OrderBookRequest struct {
	Type               TradeType `json:"type" validate:"required"`
	Price              string    `json:"price" validate:"required,numeric,max=38"`
	Amount             string    `json:"amount" validate:"required,numeric,max=38"`
	Pair               string    `json:"pair" validate:"required,alphanumunicode,min=3,max=30"`
	Currency1AccountId string    `json:"currency1AccountId" validate:"required,len=24"`
	Currency2AccountId string    `json:"currency2AccountId" validate:"required,len=24"`
	Fee                uint32    `json:"fee" validate:"required,min=0,max=100"`
	FeeAccountId       string    `json:"feeAccountId" validate:"required,len=24"`
}
