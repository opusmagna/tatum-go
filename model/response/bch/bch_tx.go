package bch

/**
 *
 * @export
 * @interface BchTx
 */
type Tx struct {
	/**
	 *
	 * @type {string}
	 * @memberof BchTx
	 */
	Txid string
	/**
	 *
	 * @type {number}
	 * @memberof BchTx
	 */
	Version uint32
	/**
	 *
	 * @type {number}
	 * @memberof BchTx
	 */
	Locktime uint64
	/**
	 *
	 * @type {Array<BchTxVin>}
	 * @memberof BchTx
	 */
	Vin []BchTxVin
	/**
	 *
	 * @type {Array<BchTxVout>}
	 * @memberof BchTx
	 */
	Vout []BchTxVout
}

/**
 *
 * @export
 * @interface BchTxScriptPubKey
 */
type TxScriptPubKey struct {

	/**
	 *
	 * @type {string}
	 * @memberof BchTxScriptPubKey
	 */
	Hex string
	/**
	 *
	 * @type {string}
	 * @memberof BchTxScriptPubKey
	 */
	Asm string
	/**
	 *
	 * @type {Array<string>}
	 * @memberof BchTxScriptPubKey
	 */
	Addresses []string
	/**
	 *
	 * @type {string}
	 * @memberof BchTxScriptPubKey
	 */
	Type string
}

/**
 *
 * @export
 * @interface BchTxScriptSig
 */
type TxScriptSig struct {
	/**
	 *
	 * @type {string}
	 * @memberof BchTxScriptSig
	 */
	Hex string
	/**
	 *
	 * @type {string}
	 * @memberof BchTxScriptSig
	 */
	Asm string
}

/**
 *
 * @export
 * @interface BchTxVin
 */
type BchTxVin struct {

	/**
	 *
	 * @type {string}
	 * @memberof BchTxVin
	 */
	Txid string
	/**
	 *
	 * @type {number}
	 * @memberof BchTxVin
	 */
	Vout uint32
	/**
	 *
	 * @type {BchTxScriptSig}
	 * @memberof BchTxVin
	 */
	ScriptSig TxScriptSig
	/**
	 *
	 * @type {string}
	 * @memberof BchTxVin
	 */
	Coinbase string
	/**
	 *
	 * @type {number}
	 * @memberof BchTxVin
	 */
	Sequence uint64
}

/**
 *
 * @export
 * @interface BchTxVout
 */
type BchTxVout struct {

	/**
	 *
	 * @type {string}
	 * @memberof BchTxVout
	 */
	Value float64
	/**
	 *
	 * @type {number}
	 * @memberof BchTxVout
	 */
	N uint32
	/**
	 *
	 * @type {BchTxScriptPubKey}
	 * @memberof BchTxVout
	 */
	ScriptPubKey TxScriptPubKey
}
