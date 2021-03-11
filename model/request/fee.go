package request

type Fee struct {
	GasLimit string `validate:"required,numeric"`
	GasPrice string `validate:"required,numeric"`
}
