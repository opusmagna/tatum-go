package blockchain

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/opusmagna/tatum-go/model/response/common"
	"github.com/opusmagna/tatum-go/model/response/ltc"
)

type Litecoin struct {
}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/LtcBroadcast" target="_blank">Tatum API documentation</a>
 */
func (b *Litecoin) LtcBroadcast(txData string, signatureId string) *common.TransactionHash {
	_url := "/v3/litecoin/broadcast"

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
 * For more details, see <a href="https://tatum.io/apidoc#operation/LtcGetBlockChainInfo" target="_blank">Tatum API documentation</a>
 */
func (b *Litecoin) LtcGetCurrentBlock() *ltc.Info {
	_url := "/v3/litecoin/info"
	var info ltc.Info
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
 * For more details, see <a href="https://tatum.io/apidoc#operation/LtcGetBlock" target="_blank">Tatum API documentation</a>
 */
func (b *Litecoin) LtcGetBlock(hash string) *ltc.Block {
	_url := "/v3/litecoin/block/" + hash

	var block ltc.Block
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
 * For more details, see <a href="https://tatum.io/apidoc#operation/LtcGetBlockHash" target="_blank">Tatum API documentation</a>
 */
func (b *Litecoin) LtcGetBlockHash(i uint64) *common.BlockHash {
	_url := strings.Join([]string{"/v3/litecoin/block/hash", strconv.FormatUint(i, 10)}, "/")
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
 * For more details, see <a href="https://tatum.io/apidoc#operation/LtcGetUTXO" target="_blank">Tatum API documentation</a>
 */
func (b *Litecoin) LtcGetUTXO(hash string, i uint64) *ltc.LtcUTXO {
	_url := strings.Join([]string{"/v3/litecoin/utxo", hash, strconv.FormatUint(i, 10)}, "/")
	var utxo ltc.LtcUTXO
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
 * For more details, see <a href="https://tatum.io/apidoc#operation/LtcGetTxByAddress" target="_blank">Tatum API documentation</a>
 */
func (b *Litecoin) LtcGetTxForAccount(address string, pageSize uint16, offset uint16) []ltc.Tx {
	_url, _ := url.Parse("/v3/litecoin/transaction/address/" + address)
	q := _url.Query()
	q.Add("offset", strconv.FormatUint(uint64(offset), 10))
	q.Add("pageSize", strconv.FormatUint(uint64(pageSize), 10))
	_url.RawQuery = q.Encode()
	fmt.Println(_url.String())

	var txs []ltc.Tx
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
 * For more details, see <a href="https://tatum.io/apidoc#operation/LtcGetRawTransaction" target="_blank">Tatum API documentation</a>
 */
func (b *Litecoin) LtcGetTransaction(hash string) *ltc.Tx {
	_url := "/v3/litecoin/transaction/" + hash
	var tx ltc.Tx
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
