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

		fmt.Println("wrong type of attribute")
		return &idRes
	}

	_url, _ := url.Parse("/v3/subscription")

	requestJSON, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	res, err := sender.SendPost(_url.String(), requestJSON)
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

	_url, _ := url.Parse("/v3/subscription")
	q := _url.Query()
	q.Add("offset", strconv.FormatUint(uint64(offset), 10))
	q.Add("pageSize", strconv.FormatUint(uint64(pageSize), 10))
	_url.RawQuery = q.Encode()
	fmt.Println(_url.String())

	var subs []ledger.Subscription
	res, err := sender.SendGet(_url.String(), nil)
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
	_url, _ := url.Parse("/v3/subscription/" + id)
	_, err := sender.SendDel(_url.String(), nil)
	if err != nil {
		fmt.Println(err)
	}
}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/getSubscriptionReport" target="_blank">Tatum API documentation</a>
 */
func (s *Subcription) ObtainReportForSubscription(id string) interface{} {

	_url, _ := url.Parse("/v3/subscription/report" + id)

	res, err := sender.SendGet(_url.String(), nil)
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
