package security

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"

	"github.com/opusmagna/tatum-go/model/request"
	"github.com/opusmagna/tatum-go/model/response/kms"
)

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/GetPendingTransactionToSign" target="_blank">Tatum API documentation</a>
 */
func GetTransactionKMS(id string) *kms.TransactionKMS {

	_url, _ := url.Parse("/v3/kms/" + id)
	q := _url.Query()
	_url.RawQuery = q.Encode()
	fmt.Println(_url.String())

	var tx kms.TransactionKMS
	res, err := sender.SendGet(_url.String(), nil)
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
 * For more details, see <a href="https://tatum.io/apidoc#operation/DeletePendingTransactionToSign" target="_blank">Tatum API documentation</a>
 */
func DeleteTransactionKMS(id string, revert bool) {
	_url, _ := url.Parse("/v3/kms/" + id + "/" + strconv.FormatBool(revert))
	_, err := sender.SendDel(_url.String(), nil)
	if err != nil {
		fmt.Println(err)
	}
}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/CompletePendingSignature" target="_blank">Tatum API documentation</a>
 */
func CompletePendingTransactionKMS(id string, txId string) {
	_url, _ := url.Parse("/v3/kms/" + id + "/" + txId)
	_, err := sender.SendPut(_url.String(), nil)
	if err != nil {
		fmt.Println(err)
	}
}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/GetPendingTransactionsToSign" target="_blank">Tatum API documentation</a>
 */
func GetPendingTransactionsKMSByChain(chain request.Currency) *[]kms.TransactionKMS {

	_url, _ := url.Parse("/v3/kms/pending/" + chain.String())
	q := _url.Query()
	_url.RawQuery = q.Encode()
	fmt.Println(_url.String())

	var txs []kms.TransactionKMS
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
