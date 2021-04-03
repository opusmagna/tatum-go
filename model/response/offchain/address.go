package offchain

type Address struct {

	/**
	 * Blockchain address.
	 * @type {string}
	 * @memberof Address
	 */

	Address string
	/**
	 * Currency of generated address. BTC, LTC, BCH, ETH, XRP, ERC20.
	 * @type {string}
	 * @memberof Address
	 */
	Currency string
	/**
	 * Derivation key index for given address.
	 * @type {number}
	 * @memberof Address
	 */
	DerivationKey uint32
	/**
	 * Extended public key to derive address from. In case of XRP, this is account address,
	 * since address is defined as DestinationTag, which is address field. In case of XLM, this is account address, since address is defined as message, which is address field.
	 * @type {string}
	 * @memberof Address
	 */
	Xpub string
	/**
	 * In case of XRP, destinationTag is the distinguisher of the account.
	 * @type {number}
	 * @memberof Address
	 */
	DestinatinTag uint64
	/**
	 * In case of XLM, message is the distinguisher of the account.
	 * @type {string}
	 * @memberof Address
	 */
	Message string
}
