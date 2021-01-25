package blockchain

import (
	"fmt"
	"testing"
)

var xrpChain = Xrp{}

func TestXrp_XrpGetLedger(t *testing.T) {
	res := xrpChain.XrpGetLedger(uint64(14374504))
	fmt.Println(res)
}

func TestXrp_XrpGetCurrentLedger(t *testing.T) {
	res := xrpChain.XrpGetCurrentLedger()
	fmt.Println(res)
}

func TestXrp_XrpGetFee(t *testing.T) {
	res := xrpChain.XrpGetFee()
	fmt.Println(res)
}

func TestXrp_XrpGetAccountTransactions(t *testing.T) {
	marker := "{\"ledger\": 14375189,\"seq\": 14375189}"
	res := xrpChain.XrpGetAccountTransactions("rHb9CJAWyB4rj91VRWn96DkukG4bwdtyTh", 0, marker)
	fmt.Println(res)
}

func TestXrp_XrpGetAccountInfo(t *testing.T) {
	res := xrpChain.XrpGetAccountInfo("rHb9CJAWyB4rj91VRWn96DkukG4bwdtyTh")
	fmt.Println(res)
}
