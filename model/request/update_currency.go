package request

type UpdateCurrency struct {
	Name     string
	BasePair interface{}
	BaseRate uint32
}
