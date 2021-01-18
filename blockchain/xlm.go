package blockchain

import (
	"encoding/json"
	"fmt"
	"github.com/tatumio/tatum-go/model/response/common"
	"github.com/tatumio/tatum-go/model/response/xlm"
	"strconv"
)

type Xlm struct {
}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/XlmGetAccountInfo" target="_blank">Tatum API documentation</a>
 */
func (x *Xlm) XlmGetAccountInfo(account string) *xlm.Sequence {
	url := "/v3/xlm/account/" + account
	res, err := sender.SendGet(url, nil)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	var result map[string]interface{}
	json.Unmarshal([]byte(res), &result)
	return &xlm.Sequence{Sequence: result["sequence"].(string)}
}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/XlmBroadcast" target="_blank">Tatum API documentation</a>
 */
func (x *Xlm) XlmBroadcast(txData string, signatureId string) *common.TransactionHash {
	url := "/v3/xlm/broadcast"

	payload := make(map[string]interface{})
	payload["txData"] = txData
	if len(signatureId) > 0 {
		payload["signatureId"] = signatureId
	}

	requestJSON, err := json.Marshal(payload)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	fmt.Println(string(requestJSON))

	txHash := common.TransactionHash{}
	var result map[string]interface{}
	res, err := sender.SendPost(url, requestJSON)
	if err == nil {
		json.Unmarshal([]byte(res), &result)
		txHash.TxId = fmt.Sprint(result["txId"])
	}
	return &txHash
}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/XlmGetLastClosedLedger" target="_blank">Tatum API documentation</a>
 */
func (x *Xlm) XlmGetCurrentLedger() string {
	url := "/v3/xlm/info"
	res, err := sender.SendGet(url, nil)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	return res
}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/XlmGetFee" target="_blank">Tatum API documentation</a>
 */
func (x *Xlm) XlmGetFee() uint32 {
	url := "/v3/xlm/fee"
	res, err := sender.SendGet(url, nil)
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}
	fee, err := strconv.ParseUint(res, 10, 32)
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}
	return uint32(fee)
}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/XlmGetLedger" target="_blank">Tatum API documentation</a>
 */
func (x *Xlm) XlmGetLedger(i uint32) string {
	url := "/v3/xlm/ledger/" + strconv.FormatUint(uint64(i), 10)
	res, err := sender.SendGet(url, nil)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	return res
}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/XlmGetLedgerTx" target="_blank">Tatum API documentation</a>
 */
func (x *Xlm) XlmGetLedgerTx(i uint32) string {
	url := "/v3/xlm/ledger/" + strconv.FormatUint(uint64(i), 10) + "/transaction"
	res, err := sender.SendGet(url, nil)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	return res
}
