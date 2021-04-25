package transaction

import (
	"encoding/json"
	"fmt"
	"github.com/btcsuite/btcutil"
	bchcfg "github.com/gcash/bchd/chaincfg"
	"github.com/go-playground/validator"
	"github.com/tatumio/tatum-go/model/request"
	"github.com/tatumio/tatum-go/model/response/bch"
	"github.com/tatumio/tatum-go/transaction/bcash_tx_builder"
	"github.com/tatumio/tatum-go/utils"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
)

func (b *BitcoinTx) prepareBitcoinCashSignedTransaction(testnet bool, body request.TransferBchBlockchain) (string, error) {
	validate = validator.New()
	err := validate.Struct(body)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err.Field() + ": should have " + err.Tag() + " " + err.Param())
			fmt.Println(err.Value())
		}
		return "", err
	}

	fromUTXO := body.FromUTXO
	to := body.To

	var network *bchcfg.Params
	if testnet {
		network = &bchcfg.TestNet3Params
	} else {
		network = &bchcfg.MainNetParams
	}

	var txHash = make([]string, 0)
	for _, utxo := range fromUTXO {
		txHash = append(txHash, utxo.FromUTXO.TxHash)
	}

	var txs = getTransactions(txHash)

	var (
		transactionBuilder = bcash_tx_builder.New().Init(network)
	)

	for i := range to {
		var value, err = btcutil.NewAmount(to[i].Value)
		if err != nil {
			return "", err
		}
		transactionBuilder.AddOutput(to[i].Address, int64(value.ToUnit(btcutil.AmountSatoshi)))
	}

	for i, utxo := range fromUTXO {
		var value, err = btcutil.NewAmount(txs[i].Vout[utxo.FromUTXO.Index].Value)
		if err != nil {
			return "", err
		}
		transactionBuilder.AddInput(utxo.FromUTXO.TxHash, utxo.FromUTXO.Index, utxo.FromUTXO.PrivateKey, int64(value.ToUnit(btcutil.AmountSatoshi)))
	}

	return transactionBuilder.Sign().ToHex(), nil
}

func getTransactions(txHash []string) []bch.Tx {

	var baseUrl = utils.TATUM_API_URL
	if len(baseUrl) == 0 {
		baseUrl = os.Getenv("TATUM_API_URL")
	}

	c := make(chan string)
	var wg sync.WaitGroup

	for _, hash := range txHash {
		wg.Add(1)
		go getTx(baseUrl+"/v3/bcash/transaction/"+hash, c, &wg)
	}

	go func() {
		wg.Wait()
		close(c)
	}()

	var txs = make([]bch.Tx, 0)
	var tx bch.Tx
	for msg := range c {
		fmt.Println(msg)
		err := json.Unmarshal([]byte(msg), &tx)
		if err != nil {
			return nil
		}
		txs = append(txs, tx)
	}

	return txs

}

func getTx(url string, c chan string, wg *sync.WaitGroup) {
	defer (*wg).Done()
	res, err := http.Get(url)

	if err == nil {
		bytes, _ := ioutil.ReadAll(res.Body)
		fmt.Println(string(bytes))
		c <- string(bytes)
	} else {
		c <- ""
	}
}
