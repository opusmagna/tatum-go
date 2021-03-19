package common

type TxHash struct {

	/**
	 * TX hash of successful transaction.
	 * @type {string}
	 * @memberof TxHash
	 */
	TxId string

	/**
	 * Whethet withdrawal was completed in Tatum's internal ledger. If not, it must be done manually.
	 * @type {boolean}
	 * @memberof TxHash
	 */
	Completed bool
}
