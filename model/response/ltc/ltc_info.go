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
	Chain string
	/**
	 * Last block.
	 * @type {number}
	 * @memberof LtcInfo
	 */
	Blocks uint64
	/**
	 * Last headers.
	 * @type {number}
	 * @memberof LtcInfo
	 */
	Headers uint64
	/**
	 * Hash of the last block.
	 * @type {string}
	 * @memberof LtcInfo
	 */
	Bestblockhash string
	/**
	 * Difficulty of the PoW algorithm.
	 * @type {number}
	 * @memberof LtcInfo
	 */
	Difficulty decimal.Decimal
}
