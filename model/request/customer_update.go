package request

import "github.com/tatumio/tatum-go/model/response/ledger"

type CustomerUpdate struct {
	CustomerCountry    Country
	AccountingCurrency ledger.Fiat
	ProviderCountry    Country
	ExternalId         string
}
