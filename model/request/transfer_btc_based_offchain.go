package request

type KeyPair struct {
	Address    string `validate:"required,min=30,max=50"`
	PrivateKey string `validate:"required,min=52,max=52"`
}

type TransferBtcBasedOffchain struct {
	Withdrawal *CreateWithdrawal `validate:"required"`
	Mnemonic   string            `validate:"min=1,max=500"`
	KeyPair    []KeyPair         `validate:"omitempty"`
}
