package blockchain

import (
	"encoding/json"
	"fmt"
	"github.com/tatumio/tatum-go/model/response/bch"
	"github.com/tatumio/tatum-go/model/response/common"
	"net/url"
	"strconv"
	"strings"
)

type Bcash struct {
}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/BchBroadcast" target="_blank">Tatum API documentation</a>
 */
func (b *Bcash) BcashBroadcast(txData string, signatureId string) *common.TransactionHash {
	url := "/v3/bcash/broadcast"

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
 * For more details, see <a href="https://tatum.io/apidoc#operation/BchGetBlockChainInfo" target="_blank">Tatum API documentation</a>
 */
func (b *Bcash) BcashGetCurrentBlock() *bch.Info {
	url := "/v3/bcash/info"
	var info bch.Info
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
 * For more details, see <a href="https://tatum.io/apidoc#operation/BchGetBlock" target="_blank">Tatum API documentation</a>
 */
func (b *Bcash) BcashGetBlock(hash string) *bch.Block {
	url := "/v3/bcash/block/" + hash
	var block bch.Block
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
 * For more details, see <a href="https://tatum.io/apidoc#operation/BchGetBlockHash" target="_blank">Tatum API documentation</a>
 */
func (b *Bcash) BcashGetBlockHash(i uint64) *common.BlockHash {
	url := strings.Join([]string{"/v3/bcash/block/hash", strconv.FormatUint(i, 10)}, "/")
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
 * For more details, see <a href="https://tatum.io/apidoc#operation/BchGetTxByAddress" target="_blank">Tatum API documentation</a>
 */
func (b *Bitcoin) BcashGetTxForAccount(address string, skip uint32) *[]bch.Tx {
	url, _ := url.Parse("/v3/bcash/transaction/address/" + address)
	q := url.Query()
	q.Add("skip", strconv.FormatUint(uint64(skip), 10))
	url.RawQuery = q.Encode()
	fmt.Println(url.String())

	var txs []bch.Tx
	res, err := sender.SendGet(url.String(), nil)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	json.Unmarshal([]byte(res), &txs)
	return &txs
}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/BchGetRawTransaction" target="_blank">Tatum API documentation</a>
 */
func (b *Bcash) BcashGetTransaction(hash string) *bch.Tx {
	url := "/v3/bcash/transaction/" + hash
	var tx bch.Tx
	res, err := sender.SendGet(url, nil)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	json.Unmarshal([]byte(res), &tx)
	return &tx
}
