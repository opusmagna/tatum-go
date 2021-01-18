package bch

import "github.com/shopspring/decimal"

type Info struct {

	/**
	 * Chain of the blockchain, main or test.
	 * @type {string}
	 * @memberof BchInfo
	 */
	Chain string
	/**
	 * Last block.
	 * @type {number}
	 * @memberof BchInfo
	 */
	Blocks uint64
	/**
	 * Hash of the last block.
	 * @type {string}
	 * @memberof BchInfo
	 */
	Bestblockhash string
	/**
	 * Difficulty of the PoW algorithm.
	 * @type {number}
	 * @memberof BchInfo
	 */
	Difficulty decimal.Decimal
}
