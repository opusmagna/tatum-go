//import axios from 'axios';
//import {TATUM_API_URL} from '../constants';
//import {Currency} from '../model/request';
//import {Fiat} from '../model/response';
//import {Rate} from '../model/response/common/Rate';

package tatum

import (
	"encoding/json"
	"fmt"
	"net/url"
	"reflect"

	"github.com/opusmagna/tatum-go/model/request"
	"github.com/opusmagna/tatum-go/model/response/common"
	"github.com/opusmagna/tatum-go/utils"
)

var sender = &utils.Async{}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/getExchangeRate" target="_blank">Tatum API documentation</a>
 */
func GetExchangeRate(currency string, basePair common.Fiat) *common.Rate {

	xType := reflect.TypeOf(currency)
	fmt.Println(xType)

	if _, ok := request.Currency(currency).IsValid(); !ok {
		if _, ok := common.Fiat(currency).IsValid(); !ok {
			fmt.Println("invalid currencry or fiat")
			return nil
		}
	}

	_url, _ := url.Parse("/v3/tatum/rate/" + currency)
	q := _url.Query()
	q.Add("basePair", string(basePair))
	_url.RawQuery = q.Encode()
	fmt.Println(_url.String())

	var rate common.Rate
	res, err := sender.SendGet(_url.String(), nil)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	err = json.Unmarshal([]byte(res), &rate)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return &rate
}
