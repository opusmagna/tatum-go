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
	_url := "/v3/ethereum/broadcast"

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
		fmt.Println(err)
		return nil
	}
	txHash.TxId = fmt.Sprint(result["txId"])

	return &txHash
}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/EthGetTransactionCount" target="_blank">Tatum API documentation</a>
 */
func (e *Ethereum) EthGetTransactionsCount(address string) uint64 {
	_url := "/v3/ethereum/transaction/count/" + address
	res, err := sender.SendGet(_url, nil)
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}

	count, err := strconv.ParseUint(res, 10, 64)
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}

	return count
}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/EthGetCurrentBlock" target="_blank">Tatum API documentation</a>
 */
func (e *Ethereum) EthGetCurrentBlock() uint32 {
	_url := "/v3/ethereum/block/current"
	res, err := sender.SendGet(_url, nil)
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
	_url := "/v3/ethereum/block/" + hash
	var block eth.Block
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
 * For more details, see <a href="https://tatum.io/apidoc#operation/EthGetBalance" target="_blank">Tatum API documentation</a>
 */
func (e *Ethereum) EthGetAccountBalance(address string) decimal.Decimal {
	_url := "/v3/ethereum/account/balance/" + address
	res, err := sender.SendGet(_url, nil)
	if err != nil {
		fmt.Println(err.Error())
		return decimal.Zero
	}

	var result map[string]interface{}
	err = json.Unmarshal([]byte(res), &result)
	if err != nil {
		fmt.Println(err.Error())
		return decimal.Zero
	}

	balance, err := decimal.NewFromString(result["balance"].(string))
	if err != nil {
		fmt.Println(err.Error())
		return decimal.Zero
	}

	return balance
}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/EthErc20GetBalance" target="_blank">Tatum API documentation</a>
 */
func (e *Ethereum) EthGetAccountErc20Address(address string, contractAddress string) *common.Balance {
	_url, _ := url.Parse("/v3/ethereum/account/balance/erc20/" + address)
	q := _url.Query()
	q.Add("contractAddress", contractAddress)
	_url.RawQuery = q.Encode()
	fmt.Println(_url.String())

	var balance common.Balance
	res, err := sender.SendGet(_url.String(), nil)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	err = json.Unmarshal([]byte(res), &balance)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return &balance
}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/EthGetTransaction" target="_blank">Tatum API documentation</a>
 */
func (e *Ethereum) EthGetTransaction(hash string) *eth.Tx {
	_url := "/v3/ethereum/transaction/" + hash
	var tx eth.Tx
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

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/EthGetTransactionByAddress" target="_blank">Tatum API documentation</a>
 */
func (e *Ethereum) EthGetAccountTransactions(address string, pageSize uint32, offset uint32) *[]eth.Tx {
	_url, _ := url.Parse("/v3/ethereum/account/transaction/" + address)
	q := _url.Query()
	q.Add("offset", strconv.FormatUint(uint64(offset), 10))
	q.Add("pageSize", strconv.FormatUint(uint64(pageSize), 10))
	_url.RawQuery = q.Encode()
	fmt.Println(_url.String())

	var txs []eth.Tx
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

	return &txs
}
