package common

type TransactionHash struct {
	/**
	 * TX hash of successful transaction.
	 * @type {string}
	 * @memberof TransactionHash
	 */
	TxId string
}

func (tx TransactionHash) New(txId string) *TransactionHash {
	p := new(TransactionHash)
	p.TxId = txId
	return p
}
