package vet

/**
 *
 * @export
 * @interface VetBlock
 */
type VetBlock struct {
	/**
	 * block number (height)
	 * @type {number}
	 * @memberof VetBlock
	 */
	Number uint32
	/**
	 * block identifier
	 * @type {string}
	 * @memberof VetBlock
	 */
	Id string
	/**
	 * RLP encoded block size in bytes
	 * @type {number}
	 * @memberof VetBlock
	 */
	Size uint32
	/**
	 * parent block ID
	 * @type {string}
	 * @memberof VetBlock
	 */
	ParentID string
	/**
	 * block unix timestamp
	 * @type {number}
	 * @memberof VetBlock
	 */
	Timestamp uint64
	/**
	 * block gas limit (max allowed accumulative gas usage of transactions)
	 * @type {number}
	 * @memberof VetBlock
	 */
	GasLimit uint64
	/**
	 * Address of account to receive block reward
	 * @type {string}
	 * @memberof VetBlock
	 */
	Beneficiary string
	/**
	 * accumulative gas usage of transactions
	 * @type {number}
	 * @memberof VetBlock
	 */
	GasUsed uint64
	/**
	 * sum of all ancestral blocks' score
	 * @type {number}
	 * @memberof VetBlock
	 */
	TotalScore uint32
	/**
	 * root hash of transactions in the block
	 * @type {string}
	 * @memberof VetBlock
	 */
	TxsRoot string
	/**
	 * supported txs features bitset
	 * @type {number}
	 * @memberof VetBlock
	 */
	TxsFeatures uint32
	/**
	 * root hash of accounts state
	 * @type {string}
	 * @memberof VetBlock
	 */
	StateRoot string
	/**
	 * root hash of transaction receipts
	 * @type {string}
	 * @memberof VetBlock
	 */
	ReceiptsRoot string
	/**
	 * the one who signed this block
	 * @type {string}
	 * @memberof VetBlock
	 */
	Signer string
	/**
	 * transactions IDs
	 * @type {Array<string>}
	 * @memberof VetBlock
	 */
	Transactions []string
}
