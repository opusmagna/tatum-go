package request

type DeployEthErc20 struct {
	Name           string `validate:"required,min=1,max=100"`
	Symbol         string `validate:"required,min=1,max=30"`
	Address        string `validate:"required,min=42,max=42"`
	Supply         string `validate:"required,numeric"`
	Digits         int    `validate:"required,min=1,max=30"`
	FromPrivateKey string `validate:"required,min=66,max=66"`
	Nonce          uint64 `validate:required,min=0`
	Fee            *Fee   `validate:"omitempty"`
}
