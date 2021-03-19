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
