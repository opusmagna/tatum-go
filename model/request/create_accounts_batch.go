package request

type CreateAccountsBatch struct {
	Accounts []CreateAccount `json:"accounts" validate:"required"`
}
