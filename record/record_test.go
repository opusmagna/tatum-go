package record

import (
	"fmt"
	"testing"

	"github.com/opusmagna/tatum-go/model/request"
)

func TestGetLogRecord(t *testing.T) {
	res := GetLogRecord(request.ETH, "1")
	fmt.Println(res)
}
