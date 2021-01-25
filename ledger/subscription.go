package ledger

import (
	"encoding/json"
	"fmt"
	"github.com/tatumio/tatum-go/model/request"
	"github.com/tatumio/tatum-go/model/response/common"
	"github.com/tatumio/tatum-go/model/response/ledger"
	"net/url"
	"reflect"
	"strconv"
)

type Subcription struct {
}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/createSubscription" target="_blank">Tatum API documentation</a>
 */
func (s *Subcription) CreateNewSubscription(data request.CreateSubscription) *common.Id {
	var idRes common.Id

	xType := reflect.TypeOf(data.Attr)
	if xType.String() != "request.SubscriptionAttrAccountBalanceLimit" &&
		xType.String() != "request.SubscriptionAttrOffchainWithdrawal" &&
		xType.String() != "request.SubscriptionAttrTxHistoryReport" &&
		xType.String() != "request.SubscriptionAttrIncomingBlockchainTx" &&
		xType.String() != "request.SubscriptionAttrCompleteBlockchainTx" {

		fmt.Errorf("wrong type of attribute")
		return &idRes
	}

	url, _ := url.Parse("/v3/subscription")

	requestJSON, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	res, err := sender.SendPost(url.String(), requestJSON)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	err = json.Unmarshal([]byte(res), &idRes)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return &idRes
}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/getSubscriptions" target="_blank">Tatum API documentation</a>
 */
func (s *Subcription) ListActiveSubscriptions(pageSize uint16, offset uint16) *[]ledger.Subscription {

	url, _ := url.Parse("/v3/subscription")
	q := url.Query()
	q.Add("offset", strconv.FormatUint(uint64(offset), 10))
	q.Add("pageSize", strconv.FormatUint(uint64(pageSize), 10))
	url.RawQuery = q.Encode()
	fmt.Println(url.String())

	var subs []ledger.Subscription
	res, err := sender.SendGet(url.String(), nil)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	err = json.Unmarshal([]byte(res), &subs)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return &subs
}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/deleteSubscription" target="_blank">Tatum API documentation</a>
 */
func (s *Subcription) CancelExistingSubscription(id string) {
	url, _ := url.Parse("/v3/subscription/" + id)
	sender.SendDel(url.String(), nil)
}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/getSubscriptionReport" target="_blank">Tatum API documentation</a>
 */
func (s *Subcription) ObtainReportForSubscription(id string) interface{} {

	url, _ := url.Parse("/v3/subscription/report" + id)

	res, err := sender.SendGet(url.String(), nil)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	var txs []ledger.Transaction
	err = json.Unmarshal([]byte(res), &txs)
	if err != nil {
		var accs []ledger.Account
		err = json.Unmarshal([]byte(res), &accs)
		if err != nil {
			fmt.Println(err.Error())
			return nil
		}
		return &accs
	}

	return &txs
}
