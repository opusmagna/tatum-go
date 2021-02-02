package request

type FromUTXOBcash struct {
	Value    float64  `validate:"required,min=0"`
	FromUTXO FromUTXO `validate:"required"`
}

type TransferBchBlockchain struct {
	FromUTXO []FromUTXOBcash `validate:"required"`
	To       []To            `validate:"required"`
}
