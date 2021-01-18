package ltc

import "github.com/shopspring/decimal"

/**
 *
 * @export
 * @interface LtcInfo
 */
type Info struct {

	/**
	 * Chain of the blockchain, main or test.
	 * @type {string}
	 * @memberof LtcInfo
	 */
	chain string
	/**
	 * Last block.
	 * @type {number}
	 * @memberof LtcInfo
	 */
	blocks uint32
	/**
	 * Last headers.
	 * @type {number}
	 * @memberof LtcInfo
	 */
	headers uint32
	/**
	 * Hash of the last block.
	 * @type {string}
	 * @memberof LtcInfo
	 */
	bestblockhash string
	/**
	 * Difficulty of the PoW algorithm.
	 * @type {number}
	 * @memberof LtcInfo
	 */
	difficulty decimal.Decimal
}
