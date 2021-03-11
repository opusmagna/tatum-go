package request

type TransferCustomErc20 struct {
	FromPrivateKey  string `validate:"required,min=66,max=66"`
	To              string `validate:"required,min=42,max=42"`
	Amount          string `validate:"required,numeric"`
	ContractAddress string `validate:"required,min=42,max=42"`
	Fee             *Fee   `validate:"omitempty"`
	Digits          int    `validate:"required,min=1,max=30"`
	Nonce           uint64 `validate:"required,min=0"`
}
