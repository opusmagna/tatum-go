package ledger

import (
	"encoding/json"
	"fmt"
	"github.com/tatumio/tatum-go/model/request"
	"github.com/tatumio/tatum-go/model/response/ledger"
	"net/url"
)

type VC struct {
}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/getCurrency" target="_blank">Tatum API documentation</a>
 */
func (v *VC) GetVirtualCurrencyByName(name string) *ledger.VC {

	_url, _ := url.Parse("/v3/ledger/virtualCurrency/" + name)
	var vc ledger.VC
	res, err := sender.SendGet(_url.String(), nil)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	err = json.Unmarshal([]byte(res), &vc)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return &vc
}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/createCurrency" target="_blank">Tatum API documentation</a>
 */
func (v *VC) CreateVirtualCurrency(data request.CreateCurrency) *ledger.Account {
	//await validateOrReject(data);
	_url := "/v3/ledger/virtualCurrency"

	requestJSON, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	var acc ledger.Account
	res, err := sender.SendPost(_url, requestJSON)
	if err == nil {
		err = json.Unmarshal([]byte(res), &acc)
	}

	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return &acc
}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/updateCurrency" target="_blank">Tatum API documentation</a>
 */
func (v *VC) UpdateVirtualCurrency(data request.UpdateCurrency) {
	//await validateOrReject(data);

	_url := "/v3/ledger/virtualCurrency"

	requestJSON, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	_, err = sender.SendPut(_url, requestJSON)
	if err != nil {
		fmt.Println(err)
	}
}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/mintCurrency" target="_blank">Tatum API documentation</a>
 */
func (v *VC) MintVirtualCurrency(data request.CurrencyOperation) *ledger.Reference {
	//await validateOrReject(data);
	_url := "/v3/ledger/virtualCurrency/mint"

	requestJSON, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	var ref ledger.Reference
	res, err := sender.SendPost(_url, requestJSON)
	if err == nil {
		err = json.Unmarshal([]byte(res), &ref)
	}

	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return &ref
}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/revokeCurrency" target="_blank">Tatum API documentation</a>
 */
func (v *VC) RevokeVirtualCurrency(data request.CurrencyOperation) *ledger.Reference {
	//await validateOrReject(data);
	_url := "/v3/ledger/virtualCurrency/revoke"

	requestJSON, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	var ref ledger.Reference
	res, err := sender.SendPost(_url, requestJSON)
	if err == nil {
		err = json.Unmarshal([]byte(res), &ref)
	}

	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return &ref
}
