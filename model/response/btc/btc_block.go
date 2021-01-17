package btc

import "math/big"

type Block struct {
	/**
	 * Hash of block.
	 * @type {string}
	 * @memberof Block
	 */
	Hash string
	/**
	 * The number of blocks preceding a particular block on a block chain.
	 * @type {number}
	 * @memberof Block
	 */
	Height uint64
	/**
	 * The number of blocks following a particular block on a block chain, including current one.
	 * @type {number}
	 * @memberof Block
	 */
	Depth uint64
	/**
	 * Block version.
	 * @type {number}
	 * @memberof Block
	 */
	Version int
	/**
	 * Hash of the previous block.
	 * @type {string}
	 * @memberof Block
	 */
	PrevBlock string
	/**
	 * The root node of a merkle tree, a descendant of all the hashed pairs in the tree.
	 * @type {string}
	 * @memberof Block
	 */
	MerkleRoot string
	/**
	 * Time of the block.
	 * @type {number}
	 * @memberof Block
	 */
	Time uint64
	/**
	 *
	 * @type {number}
	 * @memberof Block
	 */
	Bits uint64
	/**
	 * Arbitrary number that is used in Bitcoin's proof of work consensus algorithm.
	 * @type {number}
	 * @memberof Block
	 */
	Nonce big.Int
	/**
	 *
	 * @type {Array<BtcTx>}
	 * @memberof Block
	 */
	Txs []Tx
}
