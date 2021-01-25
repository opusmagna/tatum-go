package ledger

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator"
	"github.com/tatumio/tatum-go/model/request"
	"github.com/tatumio/tatum-go/model/response/common"
	"github.com/tatumio/tatum-go/model/response/ledger"
	"github.com/tatumio/tatum-go/utils"
	"net/url"
	"strconv"
	"strings"
)

type AccountLedger struct {
}

var sender = &utils.Async{}

var validate *validator.Validate

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/getAccountByAccountId" target="_blank">Tatum API documentation</a>
 */
func (a *AccountLedger) GetAccountById(id string) *ledger.Account {
	url := strings.Join([]string{"/v3/ledger/account", id}, "/")
	res, err := sender.SendGet(url, nil)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	var account ledger.Account
	err = json.Unmarshal([]byte(res), &account)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return &account
}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/createAccount" target="_blank">Tatum API documentation</a>
 */
func (a *AccountLedger) CreateAccount(account request.CreateAccount) *ledger.Account {
	validate = validator.New()
	err := validate.Struct(account)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err.Field() + ": should have " + err.Tag() + " " + err.Param())
			fmt.Println(err.Value())
		}
		return nil
	}

	url := "/v3/ledger/account"

	requestJSON, err := json.Marshal(account)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	fmt.Println(string(requestJSON))
	var accLedger ledger.Account
	res, err := sender.SendPost(url, requestJSON)
	if err == nil {
		err = json.Unmarshal([]byte(res), &accLedger)
	}

	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return &accLedger
}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/createAccountBatch" target="_blank">Tatum API documentation</a>
 */
func (a *AccountLedger) CreateAccounts(accounts request.CreateAccountsBatch) *[]ledger.Account {
	validate = validator.New()
	err := validate.Struct(accounts)

	url := "/v3/ledger/account/batch"

	requestJSON, err := json.Marshal(accounts)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	var accsLedger []ledger.Account
	res, err := sender.SendPost(url, requestJSON)
	if err == nil {
		err = json.Unmarshal([]byte(res), &accsLedger)
	}

	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return &accsLedger
}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/getBlockAmount" target="_blank">Tatum API documentation</a>
 */
func (a *AccountLedger) GetBlockedAmountsByAccountId(id string, pageSize uint16, offset uint16) *[]ledger.Blockage {

	url, _ := url.Parse("/v3/ledger/account/block/" + id)
	q := url.Query()
	q.Add("offset", strconv.FormatUint(uint64(offset), 10))
	q.Add("pageSize", strconv.FormatUint(uint64(pageSize), 10))
	url.RawQuery = q.Encode()
	fmt.Println(url.String())

	var blks []ledger.Blockage
	res, err := sender.SendGet(url.String(), nil)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	err = json.Unmarshal([]byte(res), &blks)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return &blks
}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/blockAmount" target="_blank">Tatum API documentation</a>
 */
func (a *AccountLedger) BlockAmount(id string, block request.BlockAmount) *common.Id {
	validate = validator.New()
	err := validate.Struct(block)

	url, _ := url.Parse("/v3/ledger/account/block/" + id)

	requestJSON, err := json.Marshal(block)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	res, err := sender.SendPost(url.String(), requestJSON)
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
 * For more details, see <a href="https://tatum.io/apidoc#operation/deleteBlockAmount" target="_blank">Tatum API documentation</a>
 */
func (a *AccountLedger) DeleteBlockedAmount(id string) {
	url, _ := url.Parse("/v3/ledger/account/block/" + id)
	sender.SendDel(url.String(), nil)
}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/deleteAllBlockAmount" target="_blank">Tatum API documentation</a>
 */
func (a *AccountLedger) DeleteBlockedAmountForAccount(id string) {
	url, _ := url.Parse("/v3/ledger/account/block/account/" + id)
	sender.SendDel(url.String(), nil)
}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/activateAccount" target="_blank">Tatum API documentation</a>
 */
func (a *AccountLedger) ActivateAccount(id string) {
	url, _ := url.Parse("/v3/ledger/account/" + id + "/activate")
	sender.SendPut(url.String(), nil)
}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/deactivateAccount" target="_blank">Tatum API documentation</a>
 */
func (a *AccountLedger) DeactivateAccount(id string) {
	url, _ := url.Parse("/v3/ledger/account/" + id + "/deactivate")
	sender.SendPut(url.String(), nil)
}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/freezeAccount" target="_blank">Tatum API documentation</a>
 */
func (a *AccountLedger) FreezeAccount(id string) {
	url, _ := url.Parse("/v3/ledger/account/" + id + "/freeze")
	sender.SendPut(url.String(), nil)
}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/unfreezeAccount" target="_blank">Tatum API documentation</a>
 */
func (a *AccountLedger) UnfreezeAccount(id string) {
	url, _ := url.Parse("/v3/ledger/account/" + id + "/unfreeze")
	sender.SendPut(url.String(), nil)
}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/getAccountsByCustomerId" target="_blank">Tatum API documentation</a>
 */
func (a *AccountLedger) GetAccountsByCustomerId(id string, pageSize uint16, offset uint16) *[]ledger.Account {

	url, _ := url.Parse("/v3/ledger/account/customer/" + id)
	q := url.Query()
	q.Add("offset", strconv.FormatUint(uint64(offset), 10))
	q.Add("pageSize", strconv.FormatUint(uint64(pageSize), 10))
	url.RawQuery = q.Encode()
	fmt.Println(url.String())

	var accs []ledger.Account
	res, err := sender.SendGet(url.String(), nil)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	err = json.Unmarshal([]byte(res), &accs)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return &accs
}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/getAllAccounts" target="_blank">Tatum API documentation</a>
 */
func (a *AccountLedger) GetAllAccounts(pageSize uint16, offset uint16) *[]ledger.Account {

	url, _ := url.Parse("/v3/ledger/account")
	q := url.Query()
	q.Add("offset", strconv.FormatUint(uint64(offset), 10))
	q.Add("pageSize", strconv.FormatUint(uint64(pageSize), 10))
	url.RawQuery = q.Encode()
	fmt.Println(url.String())

	var accs []ledger.Account
	res, err := sender.SendGet(url.String(), nil)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	err = json.Unmarshal([]byte(res), &accs)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return &accs
}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/getAccountBalance" target="_blank">Tatum API documentation</a>
 */
func (a *AccountLedger) GetAccountBalance(id string) *ledger.AccountBalance {

	url, _ := url.Parse("/v3/ledger/account/" + id + "/balance")

	res, err := sender.SendGet(url.String(), nil)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	var balance ledger.AccountBalance
	err = json.Unmarshal([]byte(res), &balance)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return &balance
}
