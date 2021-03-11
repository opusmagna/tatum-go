package request

type TransferEthErc20 struct {
	FromPrivateKey string    `validate:"required,min=66,max=66"`
	To             string    `validate:"required,min=42,max=42"`
	Amount         string    `validate:"required,numeric"`
	Data           string    `validate:"max=130000"`
	Currency       *Currency `validate:"required"` //@IsIn(ETH_BASED_CURRENCIES)
	Fee            *Fee      `validate:"omitempty"`
	Nonce          uint64    `validate:"min=0"`
}
