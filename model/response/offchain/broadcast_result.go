package offchain

import "github.com/tatumio/tatum-go/model/response/common"

type BroadcastResult struct {
	TxHash *common.TxHash
	Id     string
}
