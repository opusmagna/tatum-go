package ltc

/**
 *
 * @export
 * @interface LtcBlock
 */
type Block struct {

	/**
	 * Hash of block.
	 * @type {string}
	 * @memberof LtcBlock
	 */
	Hash string
	/**
	 * The number of blocks preceding a particular block on a block chain.
	 * @type {number}
	 * @memberof LtcBlock
	 */
	Height uint32
	/**
	 * Block version.
	 * @type {number}
	 * @memberof LtcBlock
	 */
	Version uint32
	/**
	 * Hash of the previous block.
	 * @type {string}
	 * @memberof LtcBlock
	 */
	PrevBlock string
	/**
	 * The root node of a merkle tree, a descendant of all the hashed pairs in the tree.
	 * @type {string}
	 * @memberof LtcBlock
	 */
	MerkleRoot string
	/**
	 * Time of the block.
	 * @type {number}
	 * @memberof LtcBlock
	 */
	Ts uint64
	/**
	 *
	 * @type {number}
	 * @memberof LtcBlock
	 */
	Bits uint64
	/**
	 * Arbitrary number that is used in Litecoin's proof of work consensus algorithm.
	 * @type {number}
	 * @memberof LtcBlock
	 */
	Nonce uint64
	/**
	 *
	 * @type {Array<Ltc_tx>}
	 * @memberof LtcBlock
	 */
	Txs []LtcTx
}
