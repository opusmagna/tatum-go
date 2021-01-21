package request

import "github.com/tatumio/tatum-go/model/response/ledger"

type CreateAccount struct {
	Currency           string
	Xpub               string
	Compliant          bool
	AccountingCurrency ledger.Fiat
	AccountCode        string
	AccountNumber      string
	Customer           CustomerUpdate
}
