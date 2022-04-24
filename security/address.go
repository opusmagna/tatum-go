package security

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/opusmagna/tatum-go/model/response/common"
	"github.com/opusmagna/tatum-go/utils"
)

var sender = &utils.Async{}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/CheckMalicousAddress" target="_blank">Tatum API documentation</a>
 */
func CheckMaliciousAddress(address string) *common.Status {

	_url, _ := url.Parse("/v3/security/address/" + address)
	q := _url.Query()
	_url.RawQuery = q.Encode()

	var status common.Status
	res, err := sender.SendGet(_url.String(), nil)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	err = json.Unmarshal([]byte(res), &status)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return &status
}
