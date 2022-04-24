package kms

import (
	"github.com/opusmagna/tatum-go/model/request"
	"github.com/opusmagna/tatum-go/model/response/offchain"
)

type TransactionKMS struct {
	Id string

	Chain request.Currency

	SserializedTransaction string

	Hashes []string

	TxId string

	WithdrawalId string

	Index uint64

	WithdrawalResponses []offchain.WithdrawalResponseData
}
