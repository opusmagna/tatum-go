package blockchain

import (
	"encoding/json"
	"fmt"
	"github.com/shopspring/decimal"
	"github.com/tatumio/tatum-go/model/response/common"
	"github.com/tatumio/tatum-go/model/response/eth"
	"net/url"
	"strconv"
)

type Ethereum struct {
}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/EthBroadcast" target="_blank">Tatum API documentation</a>
 */
func (e *Ethereum) EthBroadcast(txData string, signatureId string) *common.TransactionHash {
	url := "/v3/ethereum/broadcast"

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
 * For more details, see <a href="https://tatum.io/apidoc#operation/EthGetTransactionCount" target="_blank">Tatum API documentation</a>
 */
func (e *Ethereum) EthGetTransactionsCount(address string) uint32 {
	url := "/v3/ethereum/transaction/count/" + address
	res, err := sender.SendGet(url, nil)
	if err != nil {
		fmt.Println(err.Error())
		count, err := strconv.Atoi(res)
		if err == nil {
			return uint32(count)
		}
	}
	return 0
}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/EthGetCurrentBlock" target="_blank">Tatum API documentation</a>
 */
func (e *Ethereum) EthGetCurrentBlock() uint32 {
	url := "/v3/ethereum/block/current"
	res, err := sender.SendGet(url, nil)
	if err != nil {
		fmt.Println(err.Error())
		count, err := strconv.Atoi(res)
		if err == nil {
			return uint32(count)
		}
	}
	return 0
}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/EthGetBlock" target="_blank">Tatum API documentation</a>
 */
func (e *Ethereum) EthGetBlock(hash string) *eth.Block {
	url := "/v3/ethereum/block/" + hash
	var block eth.Block
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
 * For more details, see <a href="https://tatum.io/apidoc#operation/EthGetBalance" target="_blank">Tatum API documentation</a>
 */
func (e *Ethereum) EthGetAccountBalance(address string) decimal.Decimal {
	url := "/v3/ethereum/account/balance/" + address
	res, err := sender.SendGet(url, nil)
	if err != nil {
		fmt.Println(err.Error())
		balance, err := decimal.NewFromString(res)
		if err != nil {
			return balance
		}
	}
	return decimal.Zero
}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/EthErc20GetBalance" target="_blank">Tatum API documentation</a>
 */
func (e *Ethereum) EthGetAccountErc20Address(address string, contractAddress string) *common.Balance {
	url, _ := url.Parse("/v3/ethereum/account/balance/erc20/" + address)
	q := url.Query()
	q.Add("contractAddress", contractAddress)
	url.RawQuery = q.Encode()
	fmt.Println(url.String())

	var balance common.Balance
	res, err := sender.SendGet(url.String(), nil)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	json.Unmarshal([]byte(res), &balance)
	return &balance
}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/EthGetTransaction" target="_blank">Tatum API documentation</a>
 */
func (e *Ethereum) EthGetTransaction(hash string) *eth.Tx {
	url := "/v3/ethereum/transaction/" + hash
	var tx eth.Tx
	res, err := sender.SendGet(url, nil)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	json.Unmarshal([]byte(res), &tx)
	return &tx

}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/EthGetTransactionByAddress" target="_blank">Tatum API documentation</a>
 */
func (e *Ethereum) EthGetAccountTransactions(address string, pageSize uint32, offset uint32) *[]eth.Tx {
	url, _ := url.Parse("/v3/ethereum/account/transaction/" + address)
	q := url.Query()
	q.Add("offset", strconv.FormatUint(uint64(offset), 10))
	q.Add("pageSize", strconv.FormatUint(uint64(pageSize), 10))
	url.RawQuery = q.Encode()
	fmt.Println(url.String())

	var txs []eth.Tx
	res, err := sender.SendGet(url.String(), nil)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	json.Unmarshal([]byte(res), &txs)
	return &txs
}
