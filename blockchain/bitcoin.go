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
func (b *Bitcoin) BtcBroadcast(txData string, signatureId string) *common.TransactionHash {
	_url := "/v3/bitcoin/broadcast"

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
	res, err := sender.SendPost(_url, requestJSON)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	err = json.Unmarshal([]byte(res), &result)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	txHash.TxId = fmt.Sprint(result["txId"])

	return &txHash
}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/BtcGetBlockChainInfo" target="_blank">Tatum API documentation</a>
 *
 * @return the btc info
 */
func (b *Bitcoin) BtcGetCurrentBlock() *btc.Info {
	_url := "/v3/bitcoin/info"
	var info btc.Info
	res, err := sender.SendGet(_url, nil)
	fmt.Println(res)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	err = json.Unmarshal([]byte(res), &info)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return &info
}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/BtcGetBlock" target="_blank">Tatum API documentation</a>
 *
 * @param hash the hash
 * @return the btc block
 */
func (b *Bitcoin) BtcGetBlock(hash string) *btc.Block {
	_url := "/v3/bitcoin/block/" + hash

	var block btc.Block
	res, err := sender.SendGet(_url, nil)
	fmt.Println(res)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	err = json.Unmarshal([]byte(res), &block)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return &block
}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/BtcGetBlockHash" target="_blank">Tatum API documentation</a>
 */
func (b *Bitcoin) BtcGetBlockHash(i uint64) *common.BlockHash {
	_url := strings.Join([]string{"/v3/bitcoin/block/hash", strconv.FormatUint(i, 10)}, "/")
	var hash common.BlockHash
	res, err := sender.SendGet(_url, nil)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	err = json.Unmarshal([]byte(res), &hash)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return &hash
}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/BtcGetUTXO" target="_blank">Tatum API documentation</a>
 */
func (b *Bitcoin) BtcGetUTXO(hash string, i uint64) *btc.UTXO {
	_url := strings.Join([]string{"/v3/bitcoin/utxo", hash, strconv.FormatUint(i, 10)}, "/")
	var utxo btc.UTXO
	res, err := sender.SendGet(_url, nil)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	err = json.Unmarshal([]byte(res), &utxo)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return &utxo
}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/BtcGetTxByAddress" target="_blank">Tatum API documentation</a>
 */
func (b *Bitcoin) BtcGetTxForAccount(address string, pageSize uint16, offset uint16) []btc.Tx {
	_url, _ := url.Parse("/v3/bitcoin/transaction/address/" + address)
	q := _url.Query()
	q.Add("offset", strconv.FormatUint(uint64(offset), 10))
	q.Add("pageSize", strconv.FormatUint(uint64(pageSize), 10))
	_url.RawQuery = q.Encode()
	fmt.Println(_url.String())

	var txs []btc.Tx
	res, err := sender.SendGet(_url.String(), nil)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	err = json.Unmarshal([]byte(res), &txs)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return txs
}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/BtcGetRawTransaction" target="_blank">Tatum API documentation</a>
 */
func (b *Bitcoin) BtcGetTransaction(hash string) *btc.Tx {
	_url := "/v3/bitcoin/transaction/" + hash
	var tx btc.Tx
	res, err := sender.SendGet(_url, nil)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	err = json.Unmarshal([]byte(res), &tx)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return &tx
}
