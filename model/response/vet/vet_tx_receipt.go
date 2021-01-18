package vet

/**
 *
 * @export
 * @interface VetTxReceipt
 */
type VetTxReceipt struct {
	/**
	 *
	 * @type {number}
	 * @memberof VetTxReceipt
	 */
	GasUsed uint64
	/**
	 *
	 * @type {string}
	 * @memberof VetTxReceipt
	 */
	GasPayer string
	/**
	 *
	 * @type {string}
	 * @memberof VetTxReceipt
	 */
	Paid string
	/**
	 *
	 * @type {string}
	 * @memberof VetTxReceipt
	 */
	Reward string
	/**
	 *
	 * @type {boolean}
	 * @memberof VetTxReceipt
	 */
	Reverted bool
	/**
	 *
	 * @type {VetTxReceiptMeta}
	 * @memberof VetTxReceipt
	 */
	Meta VetTxReceiptMeta
	/**
	 *
	 * @type {Array<VetTxReceiptOutputs>}
	 * @memberof VetTxReceipt
	 */
	Outputs []VetTxReceiptOutputs
	/**
	 *
	 * @type {number}
	 * @memberof VetTxReceipt
	 */
	BlockNumber uint32
	/**
	 *
	 * @type {string}
	 * @memberof VetTxReceipt
	 */
	BlockHash string
	/**
	 *
	 * @type {string}
	 * @memberof VetTxReceipt
	 */
	TransactionHash string
	/**
	 *
	 * @type {string}
	 * @memberof VetTxReceipt
	 */
	Status string
}

/**
 *
 * @export
 * @interface VetTxReceiptMeta
 */
type VetTxReceiptMeta struct {

	/**
	 *
	 * @type {string}
	 * @memberof VetTxReceiptMeta
	 */
	BlockID string
	/**
	 *
	 * @type {number}
	 * @memberof VetTxReceiptMeta
	 */
	BlockNumber uint32
	/**
	 *
	 * @type {number}
	 * @memberof VetTxReceiptMeta
	 */
	BlockTimestamp uint64
	/**
	 *
	 * @type {string}
	 * @memberof VetTxReceiptMeta
	 */
	TxID string
	/**
	 *
	 * @type {string}
	 * @memberof VetTxReceiptMeta
	 */
	TxOrigin string
}

/**
 *
 * @export
 * @interface VetTxReceiptOutputs
 */
type VetTxReceiptOutputs struct {

	/**
	 *
	 * @type {Array<any>}
	 * @memberof VetTxReceiptOutputs
	 */
	Events []interface{}
	/**
	 *
	 * @type {Array<VetTxReceiptTransfers>}
	 * @memberof VetTxReceiptOutputs
	 */
	Transfers []VetTxReceiptTransfers
}

/**
 *
 * @export
 * @interface VetTxReceiptTransfers
 */
type VetTxReceiptTransfers struct {

	/**
	 *
	 * @type {string}
	 * @memberof VetTxReceiptTransfers
	 */
	Sender string
	/**
	 *
	 * @type {string}
	 * @memberof VetTxReceiptTransfers
	 */
	Recipient string
	/**
	 *
	 * @type {string}
	 * @memberof VetTxReceiptTransfers
	 */
	Amount string
}
