package request

type CurrencyOperation struct {
	AccountId       string
	Amount          string
	PaymentId       string
	TransactionCode string
	SenderNote      string
	RecipientNote   string
	CounterAccount  string
	Reference       string
}
