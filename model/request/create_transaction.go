package request

type CreateTransaction struct {
	SenderAccountId    string  `json:"senderAccountId" validate:"required,len=24"`
	RecipientAccountId string  `json:"recipientAccountId" validate:"required,len=24"`
	Amount             string  `json:"amount" validate:"required,numeric,max=38"`
	PaymentId          *string `json:"paymentId" validate:"min=1,max=100"`
	TransactionCode    *string `json:"transactionCode" validate:"min=1,max=100"`
	SenderNote         *string `json:"senderNote" validate:"min=1,max=500"`
	RecipientNote      *string `json:"recipientNote" validate:"min=1,max=500"`
	BaseRate           uint32  `json:"baseRate" validate:"min=0"`
	Anonymous          bool    `json:"anonymous"`
	Compliant          bool    `json:"compliant"`
}
