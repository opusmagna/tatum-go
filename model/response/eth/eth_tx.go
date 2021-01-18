package eth

/**
 *
 * @export
 * @interface EthTx
 */
type Tx struct {

	/**
	 * Hash of the block where this transaction was in.
	 * @type {string}
	 * @memberof EthTx
	 */
	BlockHash string
	/**
	 * TRUE if the transaction was successful, FALSE, if the EVM reverted the transaction.
	 * @type {boolean}
	 * @memberof EthTx
	 */
	Status bool
	/**
	 * Block number where this transaction was in.
	 * @type {number}
	 * @memberof EthTx
	 */
	BlockNumber uint32
	/**
	 * Address of the sender.
	 * @type {string}
	 * @memberof EthTx
	 */
	From string
	/**
	 * Gas provided by the sender.
	 * @type {number}
	 * @memberof EthTx
	 */
	Gas uint64
	/**
	 * Gas price provided by the sender in wei.
	 * @type {string}
	 * @memberof EthTx
	 */
	GasPrice string
	/**
	 * Hash of the transaction.
	 * @type {string}
	 * @memberof EthTx
	 */
	TransactionHash string
	/**
	 * The data sent along with the transaction.
	 * @type {string}
	 * @memberof EthTx
	 */
	Input string
	/**
	 * The number of transactions made by the sender prior to this one.
	 * @type {number}
	 * @memberof EthTx
	 */
	Nonce uint64
	/**
	 * Address of the receiver. 'null' when its a contract creation transaction.
	 * @type {string}
	 * @memberof EthTx
	 */
	To string
	/**
	 * Integer of the transactions index position in the block.
	 * @type {number}
	 * @memberof EthTx
	 */
	TransactionIndex uint64
	/**
	 * Value transferred in wei.
	 * @type {string}
	 * @memberof EthTx
	 */
	Value string
	/**
	 * The amount of gas used by this specific transaction alone.
	 * @type {number}
	 * @memberof EthTx
	 */
	GasUsed uint64
	/**
	 * The total amount of gas used when this transaction was executed in the block.
	 * @type {number}
	 * @memberof EthTx
	 */
	CumulativeGasUsed uint64
	/**
	 * The contract address created, if the transaction was a contract creation, otherwise null.
	 * @type {string}
	 * @memberof EthTx
	 */
	ContractAddress string
	/**
	 * Log events, that happened in this transaction.
	 * @type {Array<EthTxLogs>}
	 * @memberof EthTx
	 */
	Logs []EthTxLogs
}

/**
 *
 * @export
 * @interface EthTxLogs
 */
type EthTxLogs struct {

	/**
	 * From which this event originated from.
	 * @type {string}
	 * @memberof EthTxLogs
	 */
	address string
	/**
	 * An array with max 4 32 Byte topics, topic 1-3 contains indexed parameters of the log.
	 * @type {Array<string>}
	 * @memberof EthTxLogs
	 */
	topic []string
	/**
	 * The data containing non-indexed log parameter.
	 * @type {string}
	 * @memberof EthTxLogs
	 */
	data string
	/**
	 * Integer of the event index position in the block.
	 * @type {number}
	 * @memberof EthTxLogs
	 */
	logIndex uint32
	/**
	 * Integer of the transaction’s index position, the event was created in.
	 * @type {number}
	 * @memberof EthTxLogs
	 */
	transactionIndex uint32
	/**
	 * Hash of the transaction this event was created in.
	 * @type {string}
	 * @memberof EthTxLogs
	 */
	transactionHash string
}
