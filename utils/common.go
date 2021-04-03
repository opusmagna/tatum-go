package utils

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/params"
	"io/ioutil"
	"math/big"
	"net/http"
	"strconv"
)

/**
 * Estimate Gas price for the transaction.
 * @param client
 */
func EthGetGasPriceInWei() *big.Int {
	var result map[string]interface{}
	res, err := http.Get("https://ethgasstation.info/json/ethgasAPI.json")
	if err != nil {
		return nil
	}

	bytes, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(bytes))
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		return nil
	}

	data, err := strconv.Atoi(fmt.Sprint(result["fast"]))
	if err != nil {
		return nil
	}
	inGWei := new(big.Int).Div(big.NewInt(int64(data)), big.NewInt(10))
	return new(big.Int).Mul(inGWei, big.NewInt(params.GWei)) // GWei to Wei
}
