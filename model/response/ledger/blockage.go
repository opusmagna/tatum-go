package ledger

/**
 *
 * @export
 * @interface Blockage
 */
type Blockage struct {
	Id string
	/**
	 * ID of the account this blockage is for.
	 * @type {string}
	 * @memberof Blockage
	 */
	AccountId string
	/**
	 * Amount blocked on account.
	 * @type {string}
	 * @memberof Blockage
	 */
	Amount string
	/**
	 * Type of blockage.
	 * @type {string}
	 * @memberof Blockage
	 */
	Type string
	/**
	 * Description of blockage.
	 * @type {string}
	 * @memberof Blockage
	 */
	Description string
}
