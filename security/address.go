package security

import (
	"encoding/json"
	"fmt"
	"github.com/tatumio/tatum-go/model/response/common"
	"github.com/tatumio/tatum-go/utils"
	"net/url"
)

var sender = &utils.Async{}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/CheckMalicousAddress" target="_blank">Tatum API documentation</a>
 */
func CheckMaliciousAddress(address string) *common.Status {

	url, _ := url.Parse("/v3/security/address/" + address)
	q := url.Query()
	url.RawQuery = q.Encode()

	var status common.Status
	res, err := sender.SendGet(url.String(), nil)
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
