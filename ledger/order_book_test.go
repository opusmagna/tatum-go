package ledger

import (
	"fmt"
	"github.com/tatumio/tatum-go/model/request"
	"testing"
)

var orderBook = OrderBook{}

func TestOrderBook_GetTradeById(t *testing.T) {
	res := orderBook.GetTradeById("600d81987c280e37b12f5a52")
	fmt.Println(res)
}

func TestOrderBook_GetHistoricalTrades(t *testing.T) {
	res := orderBook.GetHistoricalTrades(50, 0)
	fmt.Println(res)
}

func TestOrderBook_StoreTrade(t *testing.T) {
	data := request.OrderBookRequest{}
	data.Type = request.BUY
	data.Currency2AccountId = "600e8cc1975ff77f498c9399"
	data.Currency1AccountId = "600e8cd2d918dc412c98739a"
	data.Pair = "EUR\\EUR\\/ETH"
	data.Price = "123.123"
	data.Amount = "456.456"
	res := orderBook.StoreTrade(data)
	fmt.Println(res)
}

func TestOrderBook_GetActiveBuyTrades(t *testing.T) {
	res := orderBook.GetActiveBuyTrades("600e8cc1975ff77f498c9399", 20, 0)
	fmt.Println(res)
}

func TestOrderBook_GetActiveSellTrades(t *testing.T) {
	res := orderBook.GetActiveBuyTrades("600e8cc1975ff77f498c9399", 20, 0)
	fmt.Println(res)
}
