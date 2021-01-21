package ledger

import "github.com/tatumio/tatum-go/model/response/common"

/**
 *
 * @export
 * @interface Subscription
 */

type Subscription struct {

	/**
	 * Type of the subscription.
	 * @type {string}
	 * @memberof Subscription
	 */
	Type common.SubscriptionType
	/**
	 * ID of the subscription.
	 * @type {string}
	 * @memberof Subscription
	 */
	Id string
	/**
	 * Additional attributes based on the subscription type.
	 * @type {object}
	 * @memberof Subscription
	 */
	Attr interface{}
}
