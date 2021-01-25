package ledger

import (
	"fmt"
	"github.com/tatumio/tatum-go/model/request"
	"github.com/tatumio/tatum-go/model/response/common"
	"testing"
)

var sub = Subcription{}

func TestSubcription_CreateNewSubscription(t *testing.T) {
	data := request.CreateSubscription{}
	data.Type = common.OFFCHAIN_WITHDRAWAL

	attr := request.SubscriptionAttrOffchainWithdrawal{Currency: "BTC"}
	data.Attr = attr

	res := sub.CreateNewSubscription(data)
	fmt.Println(res)
}

func TestSubcription_ListActiveSubscriptions(t *testing.T) {
	res := sub.ListActiveSubscriptions(45, 0)
	fmt.Println(res)
}
