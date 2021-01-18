package eth

/**
 *
 * @export
 * @interface EthBlock
 */
type Block struct {

	/**
	 * Difficulty for this block.
	 * @type {string}
	 * @memberof EthBlock
	 */
	Difficulty string
	/**
	 * The 'extra data' field of this block.
	 * @type {string}
	 * @memberof EthBlock
	 */
	ExtraData string
	/**
	 * The maximum gas allowed in this block.
	 * @type {number}
	 * @memberof EthBlock
	 */
	GasLimit uint64
	/**
	 * The total used gas by all transactions in this block.
	 * @type {number}
	 * @memberof EthBlock
	 */
	GasUsed uint64
	/**
	 * Hash of the block. 'null' when its pending block.
	 * @type {string}
	 * @memberof EthBlock
	 */
	Hash string
	/**
	 * The bloom filter for the logs of the block. 'null' when its pending block.
	 * @type {string}
	 * @memberof EthBlock
	 */
	LogsBloom string
	/**
	 * The address of the beneficiary to whom the mining rewards were given.
	 * @type {string}
	 * @memberof EthBlock
	 */
	Miner string
	/**
	 *
	 * @type {string}
	 * @memberof EthBlock
	 */
	MixHash string
	/**
	 * Hash of the generated proof-of-work. 'null' when its pending block.
	 * @type {string}
	 * @memberof EthBlock
	 */
	Nonce string
	/**
	 * The block number. 'null' when its pending block.
	 * @type {number}
	 * @memberof EthBlock
	 */
	Number uint32
	/**
	 * Hash of the parent block.
	 * @type {string}
	 * @memberof EthBlock
	 */
	ParentHash string
	/**
	 *
	 * @type {string}
	 * @memberof EthBlock
	 */
	ReceiptsRoot string
	/**
	 * SHA3 of the uncles data in the block.
	 * @type {string}
	 * @memberof EthBlock
	 */
	Sha3Uncles string
	/**
	 * The size of this block in bytes.
	 * @type {number}
	 * @memberof EthBlock
	 */
	Size uint32
	/**
	 * The root of the final state trie of the block.
	 * @type {string}
	 * @memberof EthBlock
	 */
	StateRoot string
	/**
	 * The unix timestamp for when the block was collated.
	 * @type {number}
	 * @memberof EthBlock
	 */
	Timestamp uint32
	/**
	 * Total difficulty of the chain until this block.
	 * @type {string}
	 * @memberof EthBlock
	 */
	TotalDifficulty string
	/**
	 * Array of transactions.
	 * @type {Array<Eth_tx>}
	 * @memberof EthBlock
	 */
	Transactions []EthTx
	/**
	 * The root of the transaction trie of the block.
	 * @type {string}
	 * @memberof EthBlock
	 */
	TransactionsRoot string
}
