package vet

/**
 *
 * @export
 * @interface VetTx
 */
type VetTx struct {

	/**
	 *
	 * @type {string}
	 * @memberof VetTx
	 */
	Id string
	/**
	 *
	 * @type {string}
	 * @memberof VetTx
	 */
	ChainTag string
	/**
	 *
	 * @type {string}
	 * @memberof VetTx
	 */
	BlockRef string
	/**
	 *
	 * @type {number}
	 * @memberof VetTx
	 */
	Expiration uint32
	/**
	 *
	 * @type {Array<VetTxClauses>}
	 * @memberof VetTx
	 */
	Clauses []VetTxClauses
	/**
	 *
	 * @type {number}
	 * @memberof VetTx
	 */
	GasPriceCoef uint64
	/**
	 *
	 * @type {number}
	 * @memberof VetTx
	 */
	Gas uint64
	/**
	 *
	 * @type {string}
	 * @memberof VetTx
	 */
	Origin string
	/**
	 *
	 * @type {string}
	 * @memberof VetTx
	 */
	Nonce string
	/**
	 *
	 * @type {number}
	 * @memberof VetTx
	 */
	Size uint32
	/**
	 *
	 * @type {VetTxMeta}
	 * @memberof VetTx
	 */
	Meta VetTxMeta
	/**
	 *
	 * @type {number}
	 * @memberof VetTx
	 */
	BlockNumber uint32
}

/**
 *
 * @export
 * @interface VetTxClauses
 */
type VetTxClauses struct {
	/**
	 *
	 * @type {string}
	 * @memberof VetTxClauses
	 */
	to string
	/**
	 *
	 * @type {string}
	 * @memberof VetTxClauses
	 */
	value string
	/**
	 *
	 * @type {string}
	 * @memberof VetTxClauses
	 */
	data string
}

/**
 *
 * @export
 * @interface VetTxMeta
 */
type VetTxMeta struct {

	/**
	 *
	 * @type {string}
	 * @memberof VetTxMeta
	 */
	blockID string
	/**
	 *
	 * @type {number}
	 * @memberof VetTxMeta
	 */
	blockNumber uint32
	/**
	 *
	 * @type {number}
	 * @memberof VetTxMeta
	 */
	blockTimestamp uint32
}
