package ledger

import (
	"fmt"
	"testing"

	"github.com/opusmagna/tatum-go/model/request"
	"github.com/opusmagna/tatum-go/model/response/common"
)

var customer = &CustomerLedger{}

func TestGetCustomer(t *testing.T) {
	cus := customer.GetCustomer("600d5497edfe394104f874c7")
	fmt.Println(cus)
	if cus.Id != "600d5497edfe394104f874c7" {
		t.Error("externalId must equal to externalid123456")
	}
}

func TestGetAllCustomers(t *testing.T) {
	cuss := customer.GetAllCustomers(25, 0)
	fmt.Println(cuss)
	if len(*cuss) < 1 {
		t.Error("length of customers must greater than zero")
	}
}

func TestUpdateCustomer(t *testing.T) {

	cus := customer.GetCustomer("600d5497edfe394104f874c7")
	fmt.Println(cus)
	update := request.CustomerUpdate{ExternalId: cus.ExternalId, AccountingCurrency: common.AFN,
		ProviderCountry: cus.ProviderCountry, CustomerCountry: cus.CustomerCountry}
	update.ExternalId = "updateId"
	id := customer.UpdateCustomer("600d5497edfe394104f874c7", update)
	fmt.Println(id)
	cus = customer.GetCustomer("600d5497edfe394104f874c7")
	fmt.Println(cus)
}
