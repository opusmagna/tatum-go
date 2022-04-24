package blockchain

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/opusmagna/tatum-go/model/response/bch"
	"github.com/opusmagna/tatum-go/model/response/common"
)

type Bcash struct {
}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/BchBroadcast" target="_blank">Tatum API documentation</a>
 */
func (b *Bcash) BcashBroadcast(txData string, signatureId string) *common.TransactionHash {
	_url := "/v3/bcash/broadcast"

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
 * For more details, see <a href="https://tatum.io/apidoc#operation/BchGetBlockChainInfo" target="_blank">Tatum API documentation</a>
 */
func (b *Bcash) BcashGetCurrentBlock() *bch.Info {
	_url := "/v3/bcash/info"
	var info bch.Info
	res, err := sender.SendGet(_url, nil)
	fmt.Println(res)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	err = json.Unmarshal([]byte(res), &info)
	if err != nil {
		return nil
	}
	return &info
}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/BchGetBlock" target="_blank">Tatum API documentation</a>
 */
func (b *Bcash) BcashGetBlock(hash string) *bch.Block {
	_url := "/v3/bcash/block/" + hash
	var block bch.Block
	res, err := sender.SendGet(_url, nil)
	fmt.Println(res)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	err = json.Unmarshal([]byte(res), &block)
	if err != nil {
		return nil
	}
	return &block
}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/BchGetBlockHash" target="_blank">Tatum API documentation</a>
 */
func (b *Bcash) BcashGetBlockHash(i uint64) *common.BlockHash {
	_url := strings.Join([]string{"/v3/bcash/block/hash", strconv.FormatUint(i, 10)}, "/")
	var hash common.BlockHash
	res, err := sender.SendGet(_url, nil)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	err = json.Unmarshal([]byte(res), &hash)
	if err != nil {
		return nil
	}
	return &hash
}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/BchGetTxByAddress" target="_blank">Tatum API documentation</a>
 */
func (b *Bitcoin) BcashGetTxForAccount(address string, skip uint32) *[]bch.Tx {
	_url, err := url.Parse("/v3/bcash/transaction/address/" + address)
	if err != nil {
		return nil
	}

	q := _url.Query()
	q.Add("skip", strconv.FormatUint(uint64(skip), 10))
	_url.RawQuery = q.Encode()

	var txs []bch.Tx
	res, err := sender.SendGet(_url.String(), nil)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	err = json.Unmarshal([]byte(res), &txs)
	if err != nil {
		return nil
	}
	return &txs
}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/BchGetRawTransaction" target="_blank">Tatum API documentation</a>
 */
func (b *Bcash) BcashGetTransaction(hash string) *bch.Tx {
	_url := "/v3/bcash/transaction/" + hash
	var tx bch.Tx
	res, err := sender.SendGet(_url, nil)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	err = json.Unmarshal([]byte(res), &tx)
	if err != nil {
		return nil
	}
	return &tx
}
