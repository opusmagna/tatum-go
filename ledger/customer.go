package ledger

import (
	"encoding/json"
	"fmt"
	"github.com/tatumio/tatum-go/model/request"
	"github.com/tatumio/tatum-go/model/response/common"
	"github.com/tatumio/tatum-go/model/response/ledger"
	"net/url"
	"strconv"
	"strings"
)

type CustomerLedger struct {
}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/getCustomerByExternalId" target="_blank">Tatum API documentation</a>
 */
func (c *CustomerLedger) GetCustomer(id string) *ledger.Customer {

	url := strings.Join([]string{"/v3/ledger/customer", id}, "/")
	res, err := sender.SendGet(url, nil)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	var customer ledger.Customer
	err = json.Unmarshal([]byte(res), &customer)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return &customer
}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/findAllCustomers" target="_blank">Tatum API documentation</a>
 */
func (c *CustomerLedger) GetAllCustomers(pageSize uint16, offset uint16) *[]ledger.Customer {

	url, _ := url.Parse("/v3/ledger/customer")
	q := url.Query()
	q.Add("offset", strconv.FormatUint(uint64(offset), 10))
	q.Add("pageSize", strconv.FormatUint(uint64(pageSize), 10))
	url.RawQuery = q.Encode()
	fmt.Println(url.String())

	var customers []ledger.Customer
	res, err := sender.SendGet(url.String(), nil)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	err = json.Unmarshal([]byte(res), &customers)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return &customers
}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/updateCustomer" target="_blank">Tatum API documentation</a>
 */
func (c *CustomerLedger) UpdateCustomer(id string, data request.CustomerUpdate) *common.Id {
	//await validateOrReject(data);

	url, _ := url.Parse("/v3/ledger/customer/" + id)

	requestJSON, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	res, err := sender.SendPut(url.String(), requestJSON)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	var idRes common.Id
	err = json.Unmarshal([]byte(res), &idRes)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return &idRes
}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/activateAccount" target="_blank">Tatum API documentation</a>
 */
func (c *CustomerLedger) ActivateCustomer(id string) {
	url, _ := url.Parse("/v3/ledger/customer/" + id + "/activate")
	sender.SendPut(url.String(), nil)
}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/deactivateCustomer" target="_blank">Tatum API documentation</a>
 */
func (c *CustomerLedger) DeactivateCustomer(id string) {
	url, _ := url.Parse("/v3/ledger/customer/" + id + "/deactivate")
	sender.SendPut(url.String(), nil)
}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/enableCustomer" target="_blank">Tatum API documentation</a>
 */
func (c *CustomerLedger) EnableCustomer(id string) {
	url, _ := url.Parse("/v3/ledger/customer/" + id + "/enable")
	sender.SendPut(url.String(), nil)
}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/disableCustomer" target="_blank">Tatum API documentation</a>
 */
func (c *CustomerLedger) DisableCustomer(id string) {
	url, _ := url.Parse("/v3/ledger/customer/" + id + "/disable")
	sender.SendPut(url.String(), nil)
}
