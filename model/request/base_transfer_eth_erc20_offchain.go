package request

type BaseTransferEthErc20Offchain struct {
	SenderAccountId string `validate:"required,min=24,max=24"`
	Address         string `validate:"required,min=42,max=42"`
	Amount          string `validate:"required,numeric"`
	Compliant       bool
	PaymentId       string `validate:"min=1,max=100"`
	SenderNote      string `validate:"min=1,max=500"`
	Nonce           uint64 `validate:"min=0"`
	GasPrice        string `validate:"numeric"`
	GasLimit        string `validate:"numeric"`
}
