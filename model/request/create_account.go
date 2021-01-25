package request

import (
	_ "github.com/go-playground/validator"
)

type CreateAccount struct {
	Currency           string          `json:"currency" validate:"required,min=2,max=40"`
	Xpub               string          `json:"xpub" validate:"omitempty,max=192"`
	Compliant          bool            `json:"compliant"`
	AccountingCurrency *string         `json:"accountingCurrency" validate:"omitempty"`
	AccountCode        *string         `json:"accountCode" validate:"omitempty,min=1,max=50"`
	AccountNumber      *string         `json:"accountNumber" validate:"omitempty,min=1,max=20"`
	Customer           *CustomerUpdate `json:"customer" validate:"omitempty"`
}
