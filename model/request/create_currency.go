package request

import (
	"github.com/tatumio/tatum-go/model/response/common"
)

type CreateCurrency struct {
	Name               string
	Supply             string
	Description        string
	AccountCode        string
	BasePair           interface{}
	BaseRate           uint32
	AccountingCurrency common.Fiat
	Customer           CustomerUpdate
}
