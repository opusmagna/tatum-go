package record

import (
	"encoding/json"
	"fmt"
	"github.com/tatumio/tatum-go/model/request"
	"github.com/tatumio/tatum-go/model/response/common"
	"github.com/tatumio/tatum-go/utils"
	"net/url"
)

var sender = &utils.Async{}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/GetLog" target="_blank">Tatum API documentation</a>
 */
func GetLogRecord(chain request.Currency, id string) *common.Rate {

	_url, _ := url.Parse("/v3/record")
	q := _url.Query()
	q.Add("chain", string(chain))
	q.Add("id", id)
	_url.RawQuery = q.Encode()
	fmt.Println(_url.String())

	var rate common.Rate
	res, err := sender.SendGet(_url.String(), nil)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	err = json.Unmarshal([]byte(res), &rate)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return &rate
}
