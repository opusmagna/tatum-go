package btc

type Tx struct {
	/**
	 * Transaction hash.
	 * @type {string}
	 * @memberof BtcTx
	 */
	Hash string
	/**
	 * Witness hash in case of a SegWit transaction.
	 * @type {string}
	 * @memberof BtcTx
	 */
	WitnessHash string
	/**
	 * Fee paid for this transaction, in satoshis.
	 * @type {number}
	 * @memberof BtcTx
	 */
	Fee uint64
	/**
	 *
	 * @type {number}
	 * @memberof BtcTx
	 */
	Rate uint64
	/**
	 *
	 * @type {number}
	 * @memberof BtcTx
	 */
	Mtime uint64
	/**
	 * Height of the block this transaction belongs to.
	 * @type {number}
	 * @memberof BtcTx
	 */
	Height uint64
	/**
	 * Hash of the block this transaction belongs to.
	 * @type {string}
	 * @memberof BtcTx
	 */
	Block string
	/**
	 * Time of the transaction.
	 * @type {number}
	 * @memberof BtcTx
	 */
	Time uint64
	/**
	 * Index of the transaction in the block.
	 * @type {number}
	 * @memberof BtcTx
	 */
	Index int
	/**
	 * Index of the transaction.
	 * @type {number}
	 * @memberof BtcTx
	 */
	Version int
	/**
	 *
	 * @type {Array<BtcTxInputs>}
	 * @memberof BtcTx
	 */
	Inputs []TxInputs
	/**
	 *
	 * @type {Array<BtcTxOutputs>}
	 * @memberof BtcTx
	 */
	Outputs []TxOutputs
	/**
	 * Block this transaction was included in.
	 * @type {number}
	 * @memberof BtcTx
	 */
	Locktime uint64
}

type TxCoin struct {

	/**
	 *
	 * @type {number}
	 * @memberof BtcTxCoin
	 */
	Version int
	/**
	 *
	 * @type {number}
	 * @memberof BtcTxCoin
	 */
	Height uint64
	/**
	 *
	 * @type {number}
	 * @memberof BtcTxCoin
	 */
	Value uint64
	/**
	 *
	 * @type {string}
	 * @memberof BtcTxCoin
	 */
	Script string
	/**
	 * Sender address.
	 * @type {string}
	 * @memberof BtcTxCoin
	 */
	Address string
	/**
	 * Coinbase transaction - miner fee.
	 * @type {boolean}
	 * @memberof BtcTxCoin
	 */
	Coinbase bool
}

type TxInputs struct {

	/**
	 *
	 * @type {BtcTxPrevout}
	 * @memberof BtcTxInputs
	 */
	Prevout TxPrevout
	/**
	 * Data generated by a spender which is almost always used as variables to satisfy a pubkey script.
	 * @type {string}
	 * @memberof BtcTxInputs
	 */
	Script string
	/**
	 * Transaction witness.
	 * @type {string}
	 * @memberof BtcTxInputs
	 */
	Witness string
	/**
	 *
	 * @type {number}
	 * @memberof BtcTxInputs
	 */
	Sequence uint64
	/**
	 *
	 * @type {BtcTxCoin}
	 * @memberof BtcTxInputs
	 */
	Coin TxCoin
}

type TxOutputs struct {

	/**
	 * Sent amount in satoshis.
	 * @type {number}
	 * @memberof BtcTxOutputs
	 */
	Value uint64
	/**
	 * Transaction script.
	 * @type {string}
	 * @memberof BtcTxOutputs
	 */
	Script string
	/**
	 * Recipient address.
	 * @type {string}
	 * @memberof BtcTxOutputs
	 */
	Address string
}

type TxPrevout struct {

	/**
	 * Transaction hash of the input.
	 * @type {string}
	 * @memberof BtcTxPrevout
	 */
	Hash string
	/**
	 * Transaction index of the input.
	 * @type {number}
	 * @memberof BtcTxPrevout
	 */
	Index int
}
