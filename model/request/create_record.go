package request

type CreateRecord struct {
	Data           string `validate:"required,min=1,max=130000"`
	FromPrivateKey string `validate:"required,min=32,max=66"`
	To             string `validate:"min=42,max=42"`
	Nonce          uint64 `validate:"min=0"`
	EthFee         *Fee   `validate:"omitempty"`
}
