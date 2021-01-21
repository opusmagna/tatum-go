package request

type OrderBookRequest struct {
	Type               TradeType
	Price              string
	Amount             string
	Pair               string
	Currency1AccountId string
	Currency2AccountId string
	Fee                uint32
	FeeAccountId       string
}
