package ledger

import (
	"fmt"
	"testing"

	"github.com/opusmagna/tatum-go/model/request"
)

var tx = Transaction{}

func TestTransaction_StoreTransaction(t *testing.T) {
	req := request.CreateTransaction{}
	req.SenderAccountId = "600e6e06c89eaa4111616aae"
	req.RecipientAccountId = "600e6e30edfe394104f874ca"
	req.Amount = "10.1"
	ref := tx.StoreTransaction(req)
	fmt.Println(ref)
}

func TestTransaction_GetTransactionsByAccount(t *testing.T) {
	filter := request.TransactionFilter{}
	res := tx.GetTransactionsByAccount(filter, 20, 0)
	fmt.Println(res)
}

func TestTransaction_CountTransactionsByAccount(t *testing.T) {
	filter := request.TransactionFilter{}
	res := tx.CountTransactionsByAccount(filter)
	fmt.Println(res)
}

func TestTransaction_GetTransactionsByCustomer(t *testing.T) {
	filter := request.TransactionFilter{}
	res := tx.GetTransactionsByCustomer(filter, 20, 0)
	fmt.Println(res)
}

func TestTransaction_CountTransactionsByCustomer(t *testing.T) {
	filter := request.TransactionFilter{}
	res := tx.CountTransactionsByCustomer(filter)
	fmt.Println(res)
}

func TestTransaction_GetTransactionsByLedger(t *testing.T) {
	filter := request.TransactionFilter{}
	res := tx.GetTransactionsByLedger(filter, 20, 0)
	fmt.Println(res)
}

func TestTransaction_CountTransactionsByLedger(t *testing.T) {
	filter := request.TransactionFilter{}
	res := tx.CountTransactionsByLedger(filter)
	fmt.Println(res)
}

func TestTransaction_GetTransactionsByReference(t *testing.T) {
	res := tx.GetTransactionsByReference("abc")
	fmt.Println(res)
}
