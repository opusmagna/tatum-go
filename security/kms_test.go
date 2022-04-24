package security

import (
	"fmt"
	"testing"

	"github.com/opusmagna/tatum-go/model/request"
)

func TestGetPendingTransactionsKMSByChain(t *testing.T) {
	res := GetPendingTransactionsKMSByChain(request.BTC)
	fmt.Println(res)
}
