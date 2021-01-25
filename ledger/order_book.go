package ledger

import (
	"encoding/json"
	"fmt"
	"github.com/tatumio/tatum-go/model/request"
	"github.com/tatumio/tatum-go/model/response/common"
	"github.com/tatumio/tatum-go/model/response/ledger"
	"net/url"
	"strconv"
)

type OrderBook struct {
}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/getHistoricalTrades" target="_blank">Tatum API documentation</a>
 */
func (o *OrderBook) GetHistoricalTrades(pageSize uint16, offset uint16) *[]ledger.OrderBookResponse {

	url, _ := url.Parse("/v3/trade/history")
	q := url.Query()
	q.Add("offset", strconv.FormatUint(uint64(offset), 10))
	q.Add("pageSize", strconv.FormatUint(uint64(pageSize), 10))
	url.RawQuery = q.Encode()
	fmt.Println(url.String())

	var orderBooks []ledger.OrderBookResponse
	res, err := sender.SendGet(url.String(), nil)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	err = json.Unmarshal([]byte(res), &orderBooks)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return &orderBooks
}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/getBuyTrades" target="_blank">Tatum API documentation</a>
 */
func (o *OrderBook) GetActiveBuyTrades(id string, pageSize uint16, offset uint16) *[]ledger.OrderBookResponse {

	url, _ := url.Parse("/v3/trade/buy")
	q := url.Query()
	q.Add("id", id)
	q.Add("offset", strconv.FormatUint(uint64(offset), 10))
	q.Add("pageSize", strconv.FormatUint(uint64(pageSize), 10))
	url.RawQuery = q.Encode()
	fmt.Println(url.String())

	var orderBooks []ledger.OrderBookResponse
	res, err := sender.SendGet(url.String(), nil)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	err = json.Unmarshal([]byte(res), &orderBooks)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return &orderBooks
}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/getSellTrades" target="_blank">Tatum API documentation</a>
 */
func (o *OrderBook) GetActiveSellTrades(id string, pageSize uint16, offset uint16) *[]ledger.OrderBookResponse {

	url, _ := url.Parse("/v3/trade/sell")
	q := url.Query()
	q.Add("id", id)
	q.Add("offset", strconv.FormatUint(uint64(offset), 10))
	q.Add("pageSize", strconv.FormatUint(uint64(pageSize), 10))
	url.RawQuery = q.Encode()
	fmt.Println(url.String())

	var orderBooks []ledger.OrderBookResponse
	res, err := sender.SendGet(url.String(), nil)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	err = json.Unmarshal([]byte(res), &orderBooks)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return &orderBooks
}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/storeTrade" target="_blank">Tatum API documentation</a>
 */
func (o *OrderBook) StoreTrade(data request.OrderBookRequest) *common.Id {

	url, _ := url.Parse("/v3/trade")

	requestJSON, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	res, err := sender.SendPost(url.String(), requestJSON)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	var idRes common.Id
	err = json.Unmarshal([]byte(res), &idRes)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return &idRes
}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/getTradeById" target="_blank">Tatum API documentation</a>
 */
func (o *OrderBook) GetTradeById(id string) *ledger.OrderBookResponse {

	url, _ := url.Parse("/v3/trade/" + id)

	var orderBook ledger.OrderBookResponse
	res, err := sender.SendGet(url.String(), nil)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	if res == "" {
		return &orderBook
	}

	err = json.Unmarshal([]byte(res), &orderBook)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return &orderBook
}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/deleteTrade" target="_blank">Tatum API documentation</a>
 */
func (o *OrderBook) DeleteTrade(id string) {
	url, _ := url.Parse("/v3/trade/" + id)
	sender.SendDel(url.String(), nil)
}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/deleteAccountTrades" target="_blank">Tatum API documentation</a>
 */
func (o *OrderBook) DeleteAccountTrades(id string) {
	url, _ := url.Parse("/v3/trade/" + id)
	sender.SendDel(url.String(), nil)
}
