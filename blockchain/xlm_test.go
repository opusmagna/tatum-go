package blockchain

import (
	"fmt"
	"testing"
)

var xlmChain = Xlm{}

func TestXlm_XlmGetAccountInfo(t *testing.T) {
	res := xlmChain.XlmGetAccountInfo("GBYDRZAFZVHPNX27XK35G7WVVG4OVHBRFNXZJEU6B63UJHGGVESLNB37")
	fmt.Println(res)
}

func TestXlm_XlmGetCurrentLedger(t *testing.T) {
	res := xlmChain.XlmGetCurrentLedger()
	fmt.Println(res)
}

func TestXlm_XlmGetFee(t *testing.T) {
	res := xlmChain.XlmGetFee()
	fmt.Println(res)
}

func TestXlm_XlmGetLedger(t *testing.T) {
	res := xlmChain.XlmGetLedger(uint32(1466790))
	fmt.Println(res)
}

func TestXlm_XlmGetLedgerTx(t *testing.T) {
	res := xlmChain.XlmGetLedgerTx(uint32(1466790))
	fmt.Println(res)
}
