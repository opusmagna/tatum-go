package tatum

import (
	"fmt"
	"github.com/tatumio/tatum-go/model/request"
	"github.com/tatumio/tatum-go/model/response/common"
	"testing"
)

func TestGetExchangeRate(t *testing.T) {
	var currency = request.BTC
	var pair common.Fiat = common.EUR
	res := GetExchangeRate(currency.String(), pair)
	fmt.Println(res)
}
