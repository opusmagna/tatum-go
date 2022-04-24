package tatum

import (
	"fmt"
	"testing"

	"github.com/opusmagna/tatum-go/model/request"
	"github.com/opusmagna/tatum-go/model/response/common"
)

func TestGetExchangeRate(t *testing.T) {
	var currency = request.BTC
	var pair common.Fiat = common.EUR
	res := GetExchangeRate(currency.String(), pair)
	fmt.Println(res)
}
