package ledger

import (
	"fmt"
	"testing"

	"github.com/opusmagna/tatum-go/model/request"
	"github.com/opusmagna/tatum-go/model/response/common"
	"github.com/stretchr/testify/assert"
)

const (
	xpub = "xpub6DYE4EtQyRirEhSujw8fCtpAHK4jxNepFiVWBy85x5Nes6vC2b2uS2dgMo22pULAHt7k79DDJnAa3DhXQnKEvc8j1s45cMC42Co7pXpnVQd"
)

var accountLedger = AccountLedger{}

func TestCreateAccount(t *testing.T) {
	createAccount := request.CreateAccount{}
	createAccount.Currency = "ETH"

	account := accountLedger.CreateAccount(createAccount)
	assert.Equal(t, "ETH", account.Currency, "they should be equal")
	assert.NotNil(t, account, "it should not be empty")
	assert.NotNil(t, account.Id, "it should not be empty")
}

func TestCreateAccountWrongCurrency(t *testing.T) {
	createAccount := request.CreateAccount{}
	createAccount.Currency = "E"

	account := accountLedger.CreateAccount(createAccount)
	assert.Nil(t, account, "it should be empty")
}

func TestCreateAccountWrongExternalId(t *testing.T) {
	createAccount := request.CreateAccount{}
	createAccount.Currency = "ETH"

	createAccount.Customer = &request.CustomerUpdate{}

	account := accountLedger.CreateAccount(createAccount)
	assert.Nil(t, account, "it should be empty")
}

func TestCreateAccountExternalId(t *testing.T) {
	createAccount := request.CreateAccount{}
	createAccount.Currency = "ETH"
	accCur := common.EUR
	createAccount.AccountingCurrency = &accCur
	createAccount.Xpub = xpub

	createAccount.Customer = &request.CustomerUpdate{CustomerCountry: request.SZ,
		AccountingCurrency: common.EUR,
		ProviderCountry:    request.SZ,
		ExternalId:         "externalid123456"}

	account := accountLedger.CreateAccount(createAccount)
	assert.Equal(t, "ETH", account.Currency, "they should be equal")
	assert.NotNil(t, account, "it should not be empty")
	assert.NotNil(t, account.Id, "it should not be empty")
}

func TestGetAccountById(t *testing.T) {
	account := accountLedger.GetAccountById("600d385009779640d6fc4bf3")
	assert.Equal(t, "xpub6DYE4EtQyRirEhSujw8fCtpAHK4jxNepFiVWBy85x5Nes6vC2b2uS2dgMo22pULAHt7k79DDJnAa3DhXQnKEvc8j1s45cMC42Co7pXpnVQd", account.Xpub, "they should be equal")
}

func TestCreateAccounts(t *testing.T) {
	createAccountsBatch := request.CreateAccountsBatch{}

	createAccount1 := request.CreateAccount{}
	createAccount1.Currency = "BTC"

	createAccount2 := request.CreateAccount{}
	createAccount2.Currency = "LTC"

	createAccountsBatch.Accounts = make([]request.CreateAccount, 0)
	createAccountsBatch.Accounts = append(createAccountsBatch.Accounts, createAccount1)
	createAccountsBatch.Accounts = append(createAccountsBatch.Accounts, createAccount2)

	res := accountLedger.CreateAccounts(createAccountsBatch)
	fmt.Println((*res)[0].Id)
	fmt.Println((*res)[1].Id)
	assert.NotNil(t, (*res)[0].Id, "it should not be empty")
	assert.NotNil(t, (*res)[1].Id, "it should not be empty")
}

func TestBlockAmount(t *testing.T) {
	blk := request.BlockAmount{}
	amount := "5"
	blk.Amount = &amount
	tp := "TYPE"
	blk.Type = &tp
	id := accountLedger.BlockAmount("600d385009779640d6fc4bf3", blk)
	fmt.Println((*id).Id)
	assert.NotNil(t, (*id).Id, "it should not be empty")
}

func TestGetBlockedAmountsByAccountId(t *testing.T) {
	blks := accountLedger.GetBlockedAmountsByAccountId("600d385009779640d6fc4bf3", 20, 0)
	fmt.Println(len(*blks))
	if len(*blks) < 1 {
		t.Error("length of blks should be greater than zero")
	}
}

func TestGetAccountsByCustomerId(t *testing.T) {
	accs := accountLedger.GetAccountsByCustomerId("600d312ae9d07754e0402792", 20, 0)
	if len(*accs) < 1 {
		t.Error("length of accs should be greater than zero")
	}
}

func TestGetAllAccounts(t *testing.T) {
	accs := accountLedger.GetAllAccounts(20, 0)
	if len(*accs) < 1 {
		t.Error("length of accs should be greater than zero")
	}
}

func TestGetAccountBalance(t *testing.T) {
	balance := accountLedger.GetAccountBalance("5ffffa1c0af53a806ff657d9")
	fmt.Println(balance)
	if balance == nil {
		t.Error("balance must be not nil")
	}
}
