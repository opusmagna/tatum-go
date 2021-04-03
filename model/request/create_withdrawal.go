package request

type CreateWithdrawal struct {
	SenderAccountId string `validate:"required,len=24"`
	Address         string `validate:"required,min=1,max=10000"`
	Amount          string `validate:"required,numeric,max=38"`
	Fee             string `validate:"numeric"`
	Compliant       bool
	PaymentId       string `validate:"min=1,max=100"`
	SenderNote      string `validate:"required,min=1,max=500"`
	MultipleAmounts []string
	Attr            string `validate:"max=64"`
}
