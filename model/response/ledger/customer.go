package ledger

import (
	"github.com/opusmagna/tatum-go/model/request"
	"github.com/opusmagna/tatum-go/model/response/common"
)

type Customer struct {
	Id string

	ExternalId string

	CustomerCountry request.Country

	AccountingCurrency common.Fiat

	ProviderCountry request.Country

	Active bool

	Enabled bool
}
