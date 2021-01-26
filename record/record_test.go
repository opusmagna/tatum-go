package record

import (
	"fmt"
	"github.com/tatumio/tatum-go/model/request"
	"testing"
)

func TestGetLogRecord(t *testing.T) {
	res := GetLogRecord(request.ETH, "1")
	fmt.Println(res)
}
