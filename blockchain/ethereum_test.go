package blockchain

import (
	"fmt"
	"testing"
)

var ethr = Ethereum{}

func TestEthereum_EthGetAccountBalance(t *testing.T) {
	res := ethr.EthGetAccountBalance("0x78c115F1c8B7D0804FbDF3CF7995B030c512ee78")
	fmt.Println(res)
}

func TestEthereum_EthGetTransactionsCount(t *testing.T) {
	res := ethr.EthGetTransactionsCount("0x78c115F1c8B7D0804FbDF3CF7995B030c512ee78")
	fmt.Println(res)
}

func TestEthereum_EthGetBlock(t *testing.T) {
	res := ethr.EthGetBlock("0xff952cb0b8314fea61b5980a540a27807dddcbf0bf0e2169fc327abd7b56ffd5")
	fmt.Println(res)
	fmt.Println(res.Difficulty)
	fmt.Println(res.ExtraData)
}

func TestEthereum_EthGetAccountErc20Address(t *testing.T) {
	res := ethr.EthGetAccountErc20Address("0x080d43DE2C3059c30f95AB2EdDcee0B0a0Ddb539",
		"0xdac17f958d2ee523a2206206994597c13d831ec7")
	fmt.Println(res)
}
