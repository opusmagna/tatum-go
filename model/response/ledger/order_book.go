package ledger

import "github.com/tatumio/tatum-go/model/request"

type OrderBookResponse struct {

	/**
	 * ID of the trade.
	 * @type {string}
	 * @memberof OrderBook
	 */
	Id string

	/**
	 * Type of the trade, BUY or SELL.
	 * @type {string}
	 * @memberof OrderBook
	 */
	Type request.TradeType

	/**
	 * Price to buy / sell.
	 * @type {string}
	 * @memberof OrderBook
	 */
	Price string

	/**
	 * Amount of the trade to be bought / sold.
	 * @type {string}
	 * @memberof OrderBook
	 */
	Amount string

	/**
	 * Trading pair.
	 * @type {string}
	 * @memberof OrderBook
	 */
	Pair string

	/**
	 * How much of the trade was already filled.
	 * @type {string}
	 * @memberof OrderBook
	 */
	Fill string

	/**
	 * ID of the account of the currency 1 trade currency.
	 * @type {string}
	 * @memberof OrderBook
	 */
	Currency1AccountId string

	/**
	 * ID of the account of the currency 2 trade currency.
	 * @type {string}
	 * @memberof OrderBook
	 */
	Currency2AccountId string

	/**
	 * Creation date, UTC millis.
	 * @type {string}
	 * @memberof OrderBook
	 */
	Created uint64
}
