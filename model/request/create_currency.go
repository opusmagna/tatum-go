package request

import "github.com/tatumio/tatum-go/model/response/ledger"

type CreateCurrency struct {
	Name               string
	Supply             string
	Description        string
	AccountCode        string
	BasePair           interface{}
	BaseRate           uint32
	AccountingCurrency ledger.Fiat
	Customer           CustomerUpdate
}
