package ledger

import "github.com/opusmagna/tatum-go/model/response/common"

/**
 *
 * @export
 * @interface MarketValue
 */
type MarketValue struct {

	/**
	 * Value of transaction in given base pair.
	 * @type {string}
	 * @memberof MarketValue
	 */
	Amount string
	/**
	 * Base pair.
	 * @type {string}
	 * @memberof MarketValue
	 */
	Currency common.Fiat
	/**
	 * Date of validity of rate in UTC.
	 * @type {number}
	 * @memberof MarketValue
	 */
	SourceDate uint32
	/**
	 * Source of base pair.
	 * @type {string}
	 * @memberof MarketValue
	 */
	Source string
}
