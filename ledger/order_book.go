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

	_url, _ := url.Parse("/v3/trade/history")
	q := _url.Query()
	q.Add("offset", strconv.FormatUint(uint64(offset), 10))
	q.Add("pageSize", strconv.FormatUint(uint64(pageSize), 10))
	_url.RawQuery = q.Encode()
	fmt.Println(_url.String())

	var orderBooks []ledger.OrderBookResponse
	res, err := sender.SendGet(_url.String(), nil)
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

	_url, _ := url.Parse("/v3/trade/buy")
	q := _url.Query()
	q.Add("id", id)
	q.Add("offset", strconv.FormatUint(uint64(offset), 10))
	q.Add("pageSize", strconv.FormatUint(uint64(pageSize), 10))
	_url.RawQuery = q.Encode()
	fmt.Println(_url.String())

	var orderBooks []ledger.OrderBookResponse
	res, err := sender.SendGet(_url.String(), nil)
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

	_url, _ := url.Parse("/v3/trade/sell")
	q := _url.Query()
	q.Add("id", id)
	q.Add("offset", strconv.FormatUint(uint64(offset), 10))
	q.Add("pageSize", strconv.FormatUint(uint64(pageSize), 10))
	_url.RawQuery = q.Encode()
	fmt.Println(_url.String())

	var orderBooks []ledger.OrderBookResponse
	res, err := sender.SendGet(_url.String(), nil)
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

	_url, _ := url.Parse("/v3/trade")

	requestJSON, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	res, err := sender.SendPost(_url.String(), requestJSON)
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

	_url, _ := url.Parse("/v3/trade/" + id)

	var orderBook ledger.OrderBookResponse
	res, err := sender.SendGet(_url.String(), nil)
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
	_url, _ := url.Parse("/v3/trade/" + id)
	_, err := sender.SendDel(_url.String(), nil)
	if err != nil {
		fmt.Println(err)
	}
}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/deleteAccountTrades" target="_blank">Tatum API documentation</a>
 */
func (o *OrderBook) DeleteAccountTrades(id string) {
	_url, _ := url.Parse("/v3/trade/" + id)
	_, err := sender.SendDel(_url.String(), nil)
	if err != nil {
		fmt.Println(err)
	}
}
