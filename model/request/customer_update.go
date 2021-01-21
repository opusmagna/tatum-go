package request

import (
	"github.com/tatumio/tatum-go/model/response/common"
)

type CustomerUpdate struct {
	CustomerCountry    Country
	AccountingCurrency common.Fiat
	ProviderCountry    Country
	ExternalId         string
}
