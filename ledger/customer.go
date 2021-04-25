package ledger

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator"
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

	_url := strings.Join([]string{"/v3/ledger/customer", id}, "/")
	res, err := sender.SendGet(_url, nil)
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

	_url, _ := url.Parse("/v3/ledger/customer")
	q := _url.Query()
	q.Add("offset", strconv.FormatUint(uint64(offset), 10))
	q.Add("pageSize", strconv.FormatUint(uint64(pageSize), 10))
	_url.RawQuery = q.Encode()
	fmt.Println(_url.String())

	var customers []ledger.Customer
	res, err := sender.SendGet(_url.String(), nil)
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
	validate = validator.New()
	err := validate.Struct(data)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err.Field() + ": should have " + err.Tag() + " " + err.Param())
			fmt.Println(err.Value())
		}
		return nil
	}

	_url, _ := url.Parse("/v3/ledger/customer/" + id)

	requestJSON, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	res, err := sender.SendPut(_url.String(), requestJSON)
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
	_url, _ := url.Parse("/v3/ledger/customer/" + id + "/activate")
	_, err := sender.SendPut(_url.String(), nil)
	if err != nil {
		fmt.Println(err)
	}
}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/deactivateCustomer" target="_blank">Tatum API documentation</a>
 */
func (c *CustomerLedger) DeactivateCustomer(id string) {
	_url, _ := url.Parse("/v3/ledger/customer/" + id + "/deactivate")
	_, err := sender.SendPut(_url.String(), nil)
	if err != nil {
		fmt.Println(err)
	}
}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/enableCustomer" target="_blank">Tatum API documentation</a>
 */
func (c *CustomerLedger) EnableCustomer(id string) {
	_url, _ := url.Parse("/v3/ledger/customer/" + id + "/enable")
	_, err := sender.SendPut(_url.String(), nil)
	if err != nil {
		fmt.Println(err)
	}
}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/disableCustomer" target="_blank">Tatum API documentation</a>
 */
func (c *CustomerLedger) DisableCustomer(id string) {
	_url, _ := url.Parse("/v3/ledger/customer/" + id + "/disable")
	_, err := sender.SendPut(_url.String(), nil)
	if err != nil {
		fmt.Println(err)
	}
}
