package request

import (
	"github.com/opusmagna/tatum-go/model/response/common"
)

type SubscriptionAttrAccountBalanceLimit struct {
	Limit         string
	TypeOfBalance string
}

type SubscriptionAttrOffchainWithdrawal struct {
	Currency string `json:"currency"`
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
	Type common.SubscriptionType `json:"type"`
	Attr interface{}             `json:"attr"`
}
