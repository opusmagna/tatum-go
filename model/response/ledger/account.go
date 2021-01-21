package ledger

/**
 *
 * @export
 * @interface Account
 */
type Account struct {

	/**
	 * For bookkeeping to distinct account purpose.
	 * @type {string}
	 * @memberof Account
	 */
	AccountCode string
	/**
	 * Account ID.
	 * @type {string}
	 * @memberof Account
	 */
	Id string
	/**
	 *
	 * @type {AccountBalance}
	 * @memberof Account
	 */
	Balance AccountBalance
	/**
	 * Time of account creation.
	 * @type {string}
	 * @memberof Account
	 */
	Created string
	/**
	 * Account currency. Supported values are BTC, LTC, BCH, ETH, XRP, Tatum Virtual Currencies started with VC_ prefix or ERC20 customer token created via Tatum Platform.
	 * @type {string}
	 * @memberof Account
	 */
	Currency string
	/**
	 * ID of newly created customer or existing customer associated with account.
	 * @type {string}
	 * @memberof Account
	 */
	CustomerId string
	/**
	 * Indicates whether account is frozen or not.
	 * @type {boolean}
	 * @memberof Account
	 */
	Frozen bool
	/**
	 * Indicates whether account is active or not.
	 * @type {boolean}
	 * @memberof Account
	 */
	Active bool
	/**
	 * Extended public key to derive address from.
	 * In case of XRP, this is account address, since address is defined as DestinationTag, which is address field.
	 * In case of XLM, this is account address, since address is defined as message, which is address field.
	 * @type {string}
	 * @memberof Account
	 */
	Xpub string
}
