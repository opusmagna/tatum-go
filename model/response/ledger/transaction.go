package ledger

import "github.com/tatumio/tatum-go/model/response/common"

/**
 *
 * @export
 * @interface Transaction
 */
type Transaction struct {

	/**
	 * Source account - source of transaction(s)
	 * @type {string}
	 * @memberof Transaction
	 */
	AccountId string
	/**
	 * Amount in account's currency
	 * @type {string}
	 * @memberof Transaction
	 */
	Amount string
	/**
	 * Whether the transaction is anonymous. If true, counter account owner does not see source account.
	 * @type {boolean}
	 * @memberof Transaction
	 */
	Anonymous bool
	/**
	 * Counter account - transaction(s) destination account. In case of blockchain recipient, this is addess of blockchain account.
	 * @type {string}
	 * @memberof Transaction
	 */
	CounterAccountId string
	/**
	 * Transaction currency
	 * @type {string}
	 * @memberof Transaction
	 */
	Currency string
	/**
	 * Time in UTC of transaction.
	 * @type {number}
	 * @memberof Transaction
	 */
	Created uint32
	/**
	 * List of market values of given transaction with all supported base pairs.
	 * @type {Array<MarketValue>}
	 * @memberof Transaction
	 */
	MarketValue []MarketValue
	/**
	 * Type of operation.
	 * @type {string}
	 * @memberof Transaction
	 */
	OperationType common.OperationType
	/**
	 * Payment ID defined in payment order by sender.
	 * @type {string}
	 * @memberof Transaction
	 */
	PaymentId string
	/**
	 * Present only for operationType WITHDRAWAL and XLM / XRP based accounts it represents message or destinationTag of the recipient, if present.
	 * @type {string}
	 * @memberof Transaction
	 */
	Attr string
	/**
	 * For operationType DEPOSIT it represents address, on which was deposit credited for the account.
	 * @type {string}
	 * @memberof Transaction
	 */
	Address string
	/**
	 * Note visible for both sender and recipient.
	 * @type {string}
	 * @memberof Transaction
	 */
	RecipientNote string
	/**
	 * Transaction internal reference - unique identifier within Tatum ledger. In order of failure, use this value to search for problems.
	 * @type {string}
	 * @memberof Transaction
	 */
	Reference string
	/**
	 * For operationType DEPOSIT, BLOCKCHAIN_TRANSACTION it represents transaction id, for which deposit occured.
	 * @type {string}
	 * @memberof Transaction
	 */
	TxId string
	/**
	 * Note visible for sender.
	 * @type {string}
	 * @memberof Transaction
	 */
	SenderNote string
	/**
	 * For bookkeeping to distinct transaction purpose.
	 * @type {string}
	 * @memberof Transaction
	 */
	TransactionCode string
	/**
	 * Type of payment.
	 * @type {string}
	 * @memberof Transaction
	 */
	TransactionType common.TransactionType
}
