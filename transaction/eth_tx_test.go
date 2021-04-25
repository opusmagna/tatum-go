package transaction

import (
	"encoding/hex"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/tatumio/tatum-go/model/request"
	"log"
	"testing"
)

func TestEthTx_PrepareStoreDataTransaction(t *testing.T) {
	fmt.Println("test")

}

func TestEthTx_PrepareEthOrErc20SignedTransaction(t *testing.T) {
	body := request.TransferEthErc20{}
	body.FromPrivateKey = "0x4874827a55d87f2309c55b835af509e3427aa4d52321eeb49a2b93b5c0f8edfb"
	body.Amount = "0"
	body.Currency = request.ETH
	body.To = "0x8cb76aed9c5e336ef961265c6079c14e9cd3d2ea"
	eth := EthTx{}
	tx, err := eth.PrepareEthOrErc20SignedTransaction(true, body, "")
	if err != nil {
		fmt.Println(err.Error())
		t.Fail()
	}
	fmt.Println(tx)
}

func TestEthTx_PrepareEthOrErc20SignedTransaction_Local(t *testing.T) {
	body := request.TransferEthErc20{}
	body.FromPrivateKey = "0x20642cff687d1953264d3c4e7e4e942edbf3602709056fa6d6c1947447aab35d"
	body.Amount = "0.1"
	body.Currency = request.ETH
	body.To = "0x68075d7d281300F3486aaA5C4b7E43F38afA3b72"
	body.Nonce = 1
	eth := EthTx{}
	tx, err := eth.PrepareEthOrErc20SignedTransaction(true, body, "")
	if err != nil || tx == "" {
		fmt.Println(err)
		t.Fail()
	}
	fmt.Println(tx)

	//rawTx := "f86d8202b28477359400825208944592d8f8d7b001e72cb26a73e4fa1806a51ac79d880de0b6b3a7640000802ca05924bde7ef10aa88db9c66dd4f5fb16b46dff2319b9968be983118b57bb50562a001b24b31010004f13d9a26b320845257a6cfc2bf819a3d55e3fc86263c5f0772"

	_tx := new(types.Transaction)
	rawTxBytes, err := hex.DecodeString(tx)
	rlp.DecodeBytes(rawTxBytes, &_tx)
	spew.Dump(_tx)

	msg, err := _tx.AsMessage(types.HomesteadSigner{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(msg.From().Hex())
	fmt.Println(msg.Nonce())
	fmt.Println(msg.Gas())
	fmt.Println(msg.GasPrice())
	fmt.Println(msg.To())
	fmt.Println(msg.Value())
	fmt.Println(msg.Data())
}

func TestEthTx_RawTransaction_Java(t *testing.T) {
	rawTx := "f86c01852e555306008265949468075d7d281300f3486aaa5c4b7e43f38afa3b7288016345785d8a0000001ba09863128dc1a075442a58180044596d5ccb983352b08c5e53a7d7005a753fe180a03c1354078a59e0a87a0d878b36a1c4a78fc29c917771b7a6862804a28517ce33"

	tx := new(types.Transaction)
	rawTxBytes, err := hex.DecodeString(rawTx)
	rlp.DecodeBytes(rawTxBytes, &tx)
	spew.Dump(tx)

	msg, err := tx.AsMessage(types.HomesteadSigner{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(msg.From().Hex())
	fmt.Println(msg.Nonce())
	fmt.Println(msg.Gas())
	fmt.Println(msg.GasPrice())
	fmt.Println(msg.To())
	fmt.Println(msg.Value())
	fmt.Println(msg.Data())
}

func TestEthTx_RawTransaction_JS(t *testing.T) {
	rawTx := ""

	tx := new(types.Transaction)
	rawTxBytes, err := hex.DecodeString(rawTx)
	rlp.DecodeBytes(rawTxBytes, &tx)
	spew.Dump(tx)

	msg, err := tx.AsMessage(types.HomesteadSigner{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(msg.From().Hex())
	fmt.Println(msg.Nonce())
	fmt.Println(msg.Gas())
	fmt.Println(msg.GasPrice())
	fmt.Println(msg.To())
	fmt.Println(msg.Value())
	fmt.Println(msg.Data())
}
