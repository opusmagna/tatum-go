package offchain

/**
 *
 * @export
 * @interface WithdrawalResponse
 */

type WithdrawalResponse struct {

	/**
	 * Transaction reference of the transaction connected to this withdrawal.
	 * @type {string}
	 * @memberof WithdrawalResponse
	 */
	Reference string
	/**
	 *
	 * @type {Array<WithdrawalResponseData>}
	 * @memberof WithdrawalResponse
	 */
	Data []WithdrawalResponseData
	/**
	 * ID of withdrawal
	 * @type {string}
	 * @memberof WithdrawalResponse
	 */
	Id string
}

/**
 *
 * @export
 * @interface WithdrawalResponseData
 */
type WithdrawalResponseData struct {

	/**
	 *
	 * @type {Address}
	 * @memberof WithdrawalResponseData
	 */
	Address Address
	/**
	 * Amount of unprocessed transaction outputs, that can be used for withdrawal. Bitcoin, Litecoin, Bitcoin Cash only.
	 * @type {number}
	 * @memberof WithdrawalResponseData
	 */
	Amount uint64
	/**
	 * Last used unprocessed transaction output, that can be used.
	 * Bitcoin, Litecoin, Bitcoin Cash only. If -1, it indicates prepared vOut with amount to be transferred to pool address.
	 * @type {string}
	 * @memberof WithdrawalResponseData
	 */
	VIn string
	/**
	 * Index of last used unprocessed transaction output in raw transaction, that can be used. Bitcoin, Litecoin, Bitcoin Cash only.
	 * @type {number}
	 * @memberof WithdrawalResponseData
	 */
	VInIndex uint32
	/**
	 * Script of last unprocessed UTXO. Bitcoin SV only.
	 * @type {string}
	 * @memberof WithdrawalResponseData
	 */
	ScriptPubKey string
}
