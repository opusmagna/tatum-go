package blockchain

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var lite = Litecoin{}

func TestLitecoin_LtcGetBlock(t *testing.T) {
	res := lite.LtcGetBlock("5064134a0d83c16f35ed0082b5f31394de67335f9c62744339aebe59f87f5c5d")
	fmt.Println(res)
	assert.Equal(t, uint64(470558643), res.Bits, "they should be equal")
	assert.Equal(t, "5064134a0d83c16f35ed0082b5f31394de67335f9c62744339aebe59f87f5c5d", res.Hash, "they should be equal")
}

func TestLitecoin_LtcGetBlockHash(t *testing.T) {
	res := lite.LtcGetBlockHash(uint64(1793409))
	fmt.Println(res)
	assert.Equal(t, "5064134a0d83c16f35ed0082b5f31394de67335f9c62744339aebe59f87f5c5d", res.Hash, "they should be equal")
}

func TestLitecoin_LtcGetCurrentBlock(t *testing.T) {
	res := lite.LtcGetCurrentBlock()
	fmt.Println(res)
}

func TestLitecoin_LtcGetTransaction(t *testing.T) {
	res := lite.LtcGetTransaction("68e98e8a75bc00883a2ee9340943832d61695de568f2f17529ea9c412b5b7be0")
	fmt.Println(res)
	assert.Equal(t, "68e98e8a75bc00883a2ee9340943832d61695de568f2f17529ea9c412b5b7be0", res.Hash, "they should be equal")
	assert.Equal(t, "892f7c8898429e885fdf0f241884a86ba8a21e16db23b5ef7f6779192b85f713", res.Inputs[0].Prevout.Hash, "they should be equal")
}

func TestLitecoin_LtcGetTxForAccount(t *testing.T) {
	res := lite.LtcGetTxForAccount("QaAqKiTwm5qpYyjuSLRXhuAHtpBuWn6vFU", 20, 1)
	fmt.Println(res)
}

func TestLitecoin_LtcGetUTXO(t *testing.T) {
	res := lite.LtcGetUTXO("68e98e8a75bc00883a2ee9340943832d61695de568f2f17529ea9c412b5b7be0", uint64(1))
	fmt.Println(res)
}
