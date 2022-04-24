package request

import (
	"github.com/opusmagna/tatum-go/model/response/common"
)

type TransactionFilter struct {
	Id              *string                 `json:"id" validate:"min=1,max=50"`
	From            *uint64                 `json:"id" validate:"min=0"`
	To              *uint64                 `json:"id" validate:"min=0"`
	Account         *string                 `json:"account" validate:"min=1,max=50"`
	CounterAccount  *string                 `json:"counterAccount" validate:"min=1,max=50"`
	Currency        *string                 `json:"currency" validate:"min=1,max=50"`
	PaymentId       *string                 `json:"paymentId" validate:"min=1,max=100"`
	TransactionCode *string                 `json:"transactionCode" validate:"min=1,max=100"`
	SenderNote      *string                 `json:"SenderNote" validate:"min=1,max=500"`
	RecipientNote   *string                 `json:"RecipientNote" validate:"min=1,max=500"`
	OpType          *common.OperationType   `json:"opType"`
	TransactionType *common.TransactionType `json:"transactionType"`
}
