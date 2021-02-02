package request

type FromAddress struct {
	Address    string `json:"address" validate:"required,min=30,max=50"`
	PrivateKey string `json:"privateKey" validate:"required,min=52,max=52"`
}

type FromUTXO struct {
	TxHash     string `json:"txHash" validate:"required,min=64,max=64"`
	Index      uint32 `json:"index" validate:"required"`
	PrivateKey string `json:"privateKey" validate:"required,min=52,max=52"`
}

type To struct {
	Address string  `json:"address" validate:"required,min=30,max=50"`
	Value   float64 `json:"value" validate:"required,min=0"`
}

type TransferBtcBasedBlockchain struct {
	FromAddress []FromAddress `json:"fromAddress" validate:"omitempty"`
	FromUTXO    []FromUTXO    `json:"fromUtxo" validate:"omitempty"`
	To          []To          `json:"to" validate:"required"`
}
