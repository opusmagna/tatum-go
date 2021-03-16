package utils

import (
	"github.com/ethereum/go-ethereum/params"
	"math/big"
	"strconv"
)

func Ether2Wei(ether string) *big.Int {
	f, _, err := big.ParseFloat(ether, 0, 100, big.ToNearestEven)
	if err != nil {
		return big.NewInt(0)
	}
	amount, _ := new(big.Float).Mul(f, big.NewFloat(params.Ether)).Float64() // Ether to GWei
	s := strconv.FormatFloat(amount, 'f', 0, 64)
	v, ok := new(big.Int).SetString(s, 10)
	if !ok {
		return big.NewInt(0)
	}
	return v
}
