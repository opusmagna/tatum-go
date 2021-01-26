package security

import (
	"encoding/json"
	"fmt"
	"github.com/tatumio/tatum-go/model/request"
	"github.com/tatumio/tatum-go/model/response/kms"
	"net/url"
	"strconv"
)

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/GetPendingTransactionToSign" target="_blank">Tatum API documentation</a>
 */
func GetTransactionKMS(id string) *kms.TransactionKMS {

	url, _ := url.Parse("/v3/kms/" + id)
	q := url.Query()
	url.RawQuery = q.Encode()
	fmt.Println(url.String())

	var tx kms.TransactionKMS
	res, err := sender.SendGet(url.String(), nil)
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
	url, _ := url.Parse("/v3/kms/" + id + "/" + strconv.FormatBool(revert))
	sender.SendDel(url.String(), nil)
}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/CompletePendingSignature" target="_blank">Tatum API documentation</a>
 */
func CompletePendingTransactionKMS(id string, txId string) {
	url, _ := url.Parse("/v3/kms/" + id + "/" + txId)
	sender.SendPut(url.String(), nil)
}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/GetPendingTransactionsToSign" target="_blank">Tatum API documentation</a>
 */
func GetPendingTransactionsKMSByChain(chain request.Currency) *[]kms.TransactionKMS {

	url, _ := url.Parse("/v3/kms/pending/" + chain.String())
	q := url.Query()
	url.RawQuery = q.Encode()
	fmt.Println(url.String())

	var txs []kms.TransactionKMS
	res, err := sender.SendGet(url.String(), nil)
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
