package request

import "github.com/tatumio/tatum-go/model/response/ledger"

type SubscriptionAttrAccountBalanceLimit struct {
	Limit         string
	TypeOfBalance string
}

type SubscriptionAttrOffchainWithdrawal struct {
	Currency string
}

type SubscriptionAttrTxHistoryReport struct {
	Interval uint32
}

type SubscriptionAttrIncomingBlockchainTx struct {
	Id  string
	Url string
}

type SubscriptionAttrCompleteBlockchainTx struct {
	Currency string
}

type CreateSubscription struct {
	Type ledger.SubscriptionType
	attr interface{}
}
