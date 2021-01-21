package request

import (
	"github.com/tatumio/tatum-go/model/response/common"
)

type CreateAccount struct {
	Currency           string
	Xpub               string
	Compliant          bool
	AccountingCurrency common.Fiat
	AccountCode        string
	AccountNumber      string
	Customer           CustomerUpdate
}
