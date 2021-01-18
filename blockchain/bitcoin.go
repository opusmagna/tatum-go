package blockchain

import (
	"encoding/json"
	"fmt"
	"github.com/tatumio/tatum-go/model/response/btc"
	"github.com/tatumio/tatum-go/model/response/common"
	"github.com/tatumio/tatum-go/utils"
	"net/url"
	"strconv"
	"strings"
)

type Bitcoin struct {
}

var sender = &utils.Async{}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/BtcBroadcast" target="_blank">Tatum API documentation</a>
 */
func (b *Bitcoin) BtcBroadcast(txData string, signatureId string) common.TransactionHash {
	url := "/v3/bitcoin/broadcast"

	payload := make(map[string]interface{})
	payload["txData"] = txData
	if len(signatureId) > 0 {
		payload["signatureId"] = signatureId
	}

	requestJSON, err := json.Marshal(payload)
	if err != nil {
		fmt.Println(err.Error())
		return common.TransactionHash{}
	}
	fmt.Println(string(requestJSON))

	txHash := common.TransactionHash{}
	var result map[string]interface{}
	res, err := sender.SendPost(url, requestJSON)
	if err == nil {
		json.Unmarshal([]byte(res), &result)
		txHash.TxId = fmt.Sprint(result["txId"])
	}
	return txHash
}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/BtcGetBlockChainInfo" target="_blank">Tatum API documentation</a>
 *
 * @return the btc info
 */
func (b *Bitcoin) BtcGetCurrentBlock() btc.Info {
	url := "/v3/bitcoin/info"
	var info btc.Info
	res, err := sender.SendGet(url, nil)
	fmt.Println(res)
	if err != nil {
		fmt.Println(err.Error())
		return info
	}
	json.Unmarshal([]byte(res), &info)
	return info
}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/BtcGetBlock" target="_blank">Tatum API documentation</a>
 *
 * @param hash the hash
 * @return the btc block
 */
func (b *Bitcoin) BtcGetBlock(hash string) btc.Block {
	url := "/v3/bitcoin/block/" + hash

	var block btc.Block
	res, err := sender.SendGet(url, nil)
	fmt.Println(res)
	if err != nil {
		fmt.Println(err.Error())
		return block
	}
	json.Unmarshal([]byte(res), &block)
	return block
}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/BtcGetBlockHash" target="_blank">Tatum API documentation</a>
 */
func (b *Bitcoin) BtcGetBlockHash(i uint64) common.BlockHash {
	url := strings.Join([]string{"/v3/bitcoin/block/hash", strconv.FormatUint(i, 10)}, "/")
	var hash common.BlockHash
	res, err := sender.SendGet(url, nil)
	if err != nil {
		fmt.Println(err.Error())
		return hash
	}
	json.Unmarshal([]byte(res), &hash)
	return hash
}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/BtcGetUTXO" target="_blank">Tatum API documentation</a>
 */
func (b *Bitcoin) BtcGetUTXO(hash string, i uint64) btc.UTXO {
	url := strings.Join([]string{"/v3/bitcoin/utxo", hash, strconv.FormatUint(i, 10)}, "/")
	var utxo btc.UTXO
	res, err := sender.SendGet(url, nil)
	if err != nil {
		fmt.Println(err.Error())
		return utxo
	}
	json.Unmarshal([]byte(res), &utxo)
	return utxo
}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/BtcGetTxByAddress" target="_blank">Tatum API documentation</a>
 */
func (b *Bitcoin) BtcGetTxForAccount(address string, pageSize uint16, offset uint16) []btc.Tx {
	url, _ := url.Parse("/v3/bitcoin/transaction/address/" + address)
	q := url.Query()
	q.Add("offset", strconv.FormatUint(uint64(offset), 10))
	q.Add("pageSize", strconv.FormatUint(uint64(pageSize), 10))
	url.RawQuery = q.Encode()
	fmt.Println(url.String())

	var txs []btc.Tx
	res, err := sender.SendGet(url.String(), nil)
	if err != nil {
		fmt.Println(err.Error())
		return txs
	}
	json.Unmarshal([]byte(res), &txs)
	return txs
}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/BtcGetRawTransaction" target="_blank">Tatum API documentation</a>
 */
func (b *Bitcoin) BtcGetTransaction(hash string) btc.Tx {
	url := "/v3/bitcoin/transaction/" + hash
	var tx btc.Tx
	res, err := sender.SendGet(url, nil)
	if err != nil {
		fmt.Println(err.Error())
		return tx
	}
	json.Unmarshal([]byte(res), &tx)
	return tx
}
