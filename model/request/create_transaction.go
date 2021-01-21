package request

type CreateTransaction struct {
	SenderAccountId    string
	RecipientAccountId string
	Amount             string
	PaymentId          string
	TransactionCode    string
	SenderNote         string
	RecipientNote      string
	BaseRate           uint32
	Anonymous          bool
	Compliant          bool
}
