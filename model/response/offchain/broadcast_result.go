package offchain

import "github.com/opusmagna/tatum-go/model/response/common"

type BroadcastResult struct {
	TxHash *common.TxHash
	Id     string
}
