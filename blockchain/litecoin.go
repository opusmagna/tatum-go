package blockchain

import (
	"encoding/json"
	"fmt"
	"github.com/tatumio/tatum-go/model/response/common"
	"github.com/tatumio/tatum-go/model/response/ltc"
	"net/url"
	"strconv"
	"strings"
)

type Litecoin struct {
}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/LtcBroadcast" target="_blank">Tatum API documentation</a>
 */
func (b *Litecoin) LtcBroadcast(txData string, signatureId string) *common.TransactionHash {
	url := "/v3/litecoin/broadcast"

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
 * For more details, see <a href="https://tatum.io/apidoc#operation/LtcGetBlockChainInfo" target="_blank">Tatum API documentation</a>
 */
func (b *Litecoin) LtcGetCurrentBlock() *ltc.Info {
	url := "/v3/litecoin/info"
	var info ltc.Info
	res, err := sender.SendGet(url, nil)
	fmt.Println(res)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	json.Unmarshal([]byte(res), &info)
	return &info
}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/LtcGetBlock" target="_blank">Tatum API documentation</a>
 */
func (b *Litecoin) LtcGetBlock(hash string) *ltc.Block {
	url := "/v3/litecoin/block/" + hash

	var block ltc.Block
	res, err := sender.SendGet(url, nil)
	fmt.Println(res)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	json.Unmarshal([]byte(res), &block)
	return &block
}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/LtcGetBlockHash" target="_blank">Tatum API documentation</a>
 */
func (b *Litecoin) LtcGetBlockHash(i uint64) *common.BlockHash {
	url := strings.Join([]string{"/v3/litecoin/block/hash", strconv.FormatUint(i, 10)}, "/")
	var hash common.BlockHash
	res, err := sender.SendGet(url, nil)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	json.Unmarshal([]byte(res), &hash)
	return &hash
}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/LtcGetUTXO" target="_blank">Tatum API documentation</a>
 */
func (b *Litecoin) LtcGetUTXO(hash string, i uint64) *ltc.LtcUTXO {
	url := strings.Join([]string{"/v3/litecoin/utxo", hash, strconv.FormatUint(i, 10)}, "/")
	var utxo ltc.LtcUTXO
	res, err := sender.SendGet(url, nil)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	json.Unmarshal([]byte(res), &utxo)
	return &utxo
}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/LtcGetTxByAddress" target="_blank">Tatum API documentation</a>
 */
func (b *Litecoin) LtcGetTxForAccount(address string, pageSize uint16, offset uint16) *[]ltc.Tx {
	url, _ := url.Parse("/v3/litecoin/transaction/address/" + address)
	q := url.Query()
	q.Add("offset", strconv.FormatUint(uint64(offset), 10))
	q.Add("pageSize", strconv.FormatUint(uint64(pageSize), 10))
	url.RawQuery = q.Encode()
	fmt.Println(url.String())

	var txs []ltc.Tx
	res, err := sender.SendGet(url.String(), nil)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	json.Unmarshal([]byte(res), &txs)
	return &txs
}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/LtcGetRawTransaction" target="_blank">Tatum API documentation</a>
 */
func (b *Litecoin) LtcGetTransaction(hash string) *ltc.Tx {
	url := "/v3/litecoin/transaction/" + hash
	var tx ltc.Tx
	res, err := sender.SendGet(url, nil)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	json.Unmarshal([]byte(res), &tx)
	return &tx
}
