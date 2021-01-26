package security

import (
	"fmt"
	"github.com/tatumio/tatum-go/model/request"
	"testing"
)

func TestGetPendingTransactionsKMSByChain(t *testing.T) {
	res := GetPendingTransactionsKMSByChain(request.BTC)
	fmt.Println(res)
}
