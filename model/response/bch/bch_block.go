package bch

import "github.com/shopspring/decimal"

type Block struct {

	/**
	 * Hash of block.
	 * @type {string}
	 * @memberof BchBlock
	 */
	Hash string
	/**
	 * Block size.
	 * @type {number}
	 * @memberof BchBlock
	 */
	Size uint32
	/**
	 * The number of blocks preceding a particular block on a block chain.
	 * @type {number}
	 * @memberof BchBlock
	 */
	Height uint64
	/**
	 * Block version.
	 * @type {number}
	 * @memberof BchBlock
	 */
	Version uint32
	/**
	 * The root node of a merkle tree, a descendant of all the hashed pairs in the tree.
	 * @type {string}
	 * @memberof BchBlock
	 */
	Merkleroot string
	/**
	 * List of transactions present in the block.
	 * @type {Array<BchTx>}
	 * @memberof BchBlock
	 */
	Tx []Tx
	/**
	 * Time of the block.
	 * @type {number}
	 * @memberof BchBlock
	 */
	Time uint64
	/**
	 * Arbitrary number that is used in Bitcoin's proof of work consensus algorithm.
	 * @type {number}
	 * @memberof BchBlock
	 */
	Nonce uint64
	/**
	 *
	 * @type {number}
	 * @memberof BchBlock
	 */
	Difficulty decimal.Decimal
	/**
	 * Number of blocks mined after this block.
	 * @type {number}
	 * @memberof BchBlock
	 */
	Confirmations uint32
	/**
	 * Hash of the previous block.
	 * @type {string}
	 * @memberof BchBlock
	 */
	Previousblockhash string
	/**
	 * Hash of the next block.
	 * @type {string}
	 * @memberof BchBlock
	 */
	Nextblockhash string
}
