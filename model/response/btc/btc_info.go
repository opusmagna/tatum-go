package btc

import "github.com/shopspring/decimal"

type Info struct {
	Chain         string          `json:"chain"`
	Blocks        uint64          `json:"blocks"`
	Headers       uint64          `json:"headers"`
	Bestblockhash string          `json:"bestblockhash"`
	Difficulty    decimal.Decimal `json:"difficulty"`
}
