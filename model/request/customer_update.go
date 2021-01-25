package request

import (
	"github.com/tatumio/tatum-go/model/response/common"
)

type CustomerUpdate struct {
	CustomerCountry    Country     `json:"customerCountry" validate:"min=2,max=2"`
	AccountingCurrency common.Fiat `json:"accountingCurrency" validate:"min=3,max=3"`
	ProviderCountry    Country     `json:"providerCountry" validate:"min=2,max=2"`
	ExternalId         string      `json:"externalId" validate:"min=1,max=100"`
}
