package request

type VetFee struct {
	GasLimit string `validate:"required,numeric"`
}

type TransferVet struct {
	FromPrivateKey string  `validate:"required,min=66,max=66"`
	To             string  `validate:"required,min=42,max=42"`
	Amount         string  `validate:"required,numeric"`
	Data           string  `validate:"required,max=10000"`
	Fee            *VetFee `validate:"omitempty"`
}
