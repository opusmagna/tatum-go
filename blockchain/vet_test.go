package blockchain

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/tatumio/tatum-go/model/request"
	"math/big"
	"testing"
)

var vetchain = Vet{}

func TestVet_VetGetTransaction(t *testing.T) {
	res := vetchain.VetGetTransaction("0x96f20206b4cc6b1458e8d599ad8c3df6941d7822477e7c39cf6fb31e7a138372")
	fmt.Println(res)
	assert.Equal(t, "0x96f20206b4cc6b1458e8d599ad8c3df6941d7822477e7c39cf6fb31e7a138372", (*res).Id, "they should be equal")
}

func TestVet_VetGetTransactionReceipt(t *testing.T) {
	res := vetchain.VetGetTransactionReceipt("0x96f20206b4cc6b1458e8d599ad8c3df6941d7822477e7c39cf6fb31e7a138372")
	fmt.Println(res)
	assert.Equal(t, "0x96f20206b4cc6b1458e8d599ad8c3df6941d7822477e7c39cf6fb31e7a138372", (*res).TransactionHash, "they should be equal")
}

func TestVet_VetGetBlock(t *testing.T) {
	res := vetchain.VetGetBlock("0x007c6eeb0b4fe4fd07067e9180384b9eec6507be52d157693526e6f25f63814d")
	fmt.Println(res)
	assert.Equal(t, "0x007c6eeb0b4fe4fd07067e9180384b9eec6507be52d157693526e6f25f63814d", (*res).Id, "they should be equal")
}

func TestVet_VetGetAccountEnergy(t *testing.T) {
	res := vetchain.VetGetAccountEnergy("0x0d7A1a48a8996Dd51F94d9e9118cCb028562B955")
	fmt.Println(res)
}

func TestVet_VetGetAccountBalance(t *testing.T) {
	res := vetchain.VetGetAccountBalance("0x0d7A1a48a8996Dd51F94d9e9118cCb028562B955")
	fmt.Println(res)
}

func TestVet_VetEstimateGas(t *testing.T) {
	body := request.EstimateGasVet{}
	body.Value = "123123"
	body.Data = "0x0d7A1a48a8996Dd51F94d9e9118cCb028562B955"
	body.From = "123456789012345678901234567890123456789012345678901234567890123456"
	body.To = "0x0d7A1a48a8996Dd51F94d9e9118cCb028562B955"
	body.Nonce, _ = big.NewInt(0).SetString("100", 10)
	res := vetchain.VetEstimateGas(body)
	fmt.Println(res)
}
