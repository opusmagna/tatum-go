package offchain

import (
	"encoding/json"
	"fmt"
	"github.com/tatumio/tatum-go/model/request"
	"github.com/tatumio/tatum-go/model/response/common"
	"github.com/tatumio/tatum-go/model/response/offchain"
	"github.com/tatumio/tatum-go/utils"
	"net/url"
	"strconv"
)

var sender = &utils.Async{}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/storeWithdrawal" target="_blank">Tatum API documentation</a>
 */
func OffchainStoreWithdrawal(data request.CreateWithdrawal) *offchain.WithdrawalResponse {
	url := "/v3/offchain/withdrawal"

	requestJSON, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	fmt.Println(string(requestJSON))

	withdrawal := offchain.WithdrawalResponse{}
	res, err := sender.SendPost(url, requestJSON)
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
func OffchainCancelWithdrawal(id string, revert bool) {
	url, _ := url.Parse("/v3/offchain/withdrawal/" + id + "?revert=" + strconv.FormatBool(revert))
	sender.SendDel(url.String(), nil)
}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/broadcastBlockchainTransaction" target="_blank">Tatum API documentation</a>
 */
func OffchainBroadcast(data request.BroadcastWithdrawal) (*common.TxHash, error) {
	url := "/v3/offchain/withdrawal/broadcast"

	requestJSON, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	fmt.Println(string(requestJSON))

	txHash := common.TxHash{}
	res, err := sender.SendPost(url, requestJSON)
	if err == nil {
		err = json.Unmarshal([]byte(res), &txHash)
	}

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return &txHash
}
