package ledger

/**
 * @export
 * @enum {string}
 */

type OperationType string

const (
	PAYMENT                OperationType = "PAYMENT"
	WITHDRAWAL                           = "WITHDRAWAL"
	BLOCKCHAIN_TRANSACTION               = "BLOCKCHAIN_TRANSACTION"
	EXCHANGE                             = "EXCHANGE"
	//FAILED = "FAILED" // TO-DO
	DEPOSIT = "DEPOSIT"
	MINT    = "MINT"
	REVOKE  = "REVOKE"
)
