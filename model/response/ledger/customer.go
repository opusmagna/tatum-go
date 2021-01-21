package ledger

import "github.com/tatumio/tatum-go/model/request"

type Customer struct {
	Id string

	ExternalId string

	CustomerCountry request.Country

	AccountingCurrency Fiat

	ProviderCountry request.Country

	Active bool

	Enabled bool
}
