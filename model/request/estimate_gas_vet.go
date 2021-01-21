package request

type EstimateGasVet struct {
	From  string
	To    string
	Value string
	Data  string
	Nonce uint64
}
