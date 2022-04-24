package offchain

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"

	"github.com/opusmagna/tatum-go/model/request"
	"github.com/opusmagna/tatum-go/model/response/common"
	"github.com/opusmagna/tatum-go/model/response/offchain"
	"github.com/opusmagna/tatum-go/utils"
)

var sender = &utils.Async{}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/storeWithdrawal" target="_blank">Tatum API documentation</a>
 */
func StoreWithdrawalOffchain(data request.CreateWithdrawal) *offchain.WithdrawalResponse {
	_url := "/v3/offchain/withdrawal"

	requestJSON, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	fmt.Println(string(requestJSON))

	withdrawal := offchain.WithdrawalResponse{}
	res, err := sender.SendPost(_url, requestJSON)
	if err == nil {
		err = json.Unmarshal([]byte(res), &withdrawal)
	}

	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return &withdrawal
}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/cancelInProgressWithdrawal" target="_blank">Tatum API documentation</a>
 */
func CancelWithdrawalOffchain(id string, revert bool) (string, error) {
	_url, _ := url.Parse("/v3/offchain/withdrawal/" + id + "?revert=" + strconv.FormatBool(revert))
	return sender.SendDel(_url.String(), nil)
}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/broadcastBlockchainTransaction" target="_blank">Tatum API documentation</a>
 */
func BroadcastOffchain(data request.BroadcastWithdrawal) (*common.TxHash, error) {
	_url := "/v3/offchain/withdrawal/broadcast"

	requestJSON, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	fmt.Println(string(requestJSON))

	txHash := common.TxHash{}
	res, err := sender.SendPost(_url, requestJSON)
	if err == nil {
		err = json.Unmarshal([]byte(res), &txHash)
	}

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return &txHash, nil
}
