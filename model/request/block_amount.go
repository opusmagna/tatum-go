package request

type BlockAmount struct {
	Amount      *string `json:"amount"`
	Type        *string `json:"type"`
	Description *string `json:"description"`
}
