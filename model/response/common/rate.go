package common

type Rate struct {
	Id        string
	Value     string
	BasePair  Fiat
	Timestamp uint64
	Source    string
}
