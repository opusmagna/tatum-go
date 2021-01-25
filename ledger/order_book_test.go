package ledger

import (
	"fmt"
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
