package request

type BroadcastWithdrawal struct {
	/**
	 * Currency of signed transaction to be broadcast, BTC, LTC, BCH, ETH, XRP, ERC20
	 * @type {string}
	 * @memberof BroadcastWithdrawal
	 */
	Currency string
	/**
	 * Raw signed transaction to be published to network.
	 * @type {string}
	 * @memberof BroadcastWithdrawal
	 */
	TxData string
	/**
	 * Withdrawal ID to be completed by transaction broadcast
	 * @type {string}
	 * @memberof BroadcastWithdrawal
	 */
	WithdrawalId string
	/**
	 * Signature ID to be completed by transaction broadcast
	 * @type {string}
	 * @memberof BroadcastWithdrawal
	 */
	SignatureId string
}
