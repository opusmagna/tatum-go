package request

import "math/big"

type EstimateGasVet struct {
	From  string   `json:"from" validate:"required,len=66"`
	To    string   `json:"to" validate:"required,len=42"`
	Value string   `json:"value" validate:"required,numeric"`
	Data  string   `json:"data" validate:"required,len=10000"`
	Nonce *big.Int `json:"nonce" validate:"omitempty,min=0"`
}
