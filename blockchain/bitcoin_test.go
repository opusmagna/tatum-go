package blockchain

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBtcBroadcast(t *testing.T) {

	txData := "62BD544D1B9031EFC330A3E855CC3A0D51CA5131455C1AB3BCAC6D243F65460D"
	bitcoin := &Bitcoin{}
	txHash := bitcoin.BtcBroadcast(txData, "1234-1234-1234-1234-1234")
	fmt.Println(txHash)
}

func TestBtcGetCurrentBlock(t *testing.T) {

	bitcoin := Bitcoin{}
	info := bitcoin.BtcGetCurrentBlock()
	fmt.Println(info.Chain)
	fmt.Println(info.Blocks)
	fmt.Println(info.Headers)
	fmt.Println(info.Bestblockhash)
	fmt.Println(info.Difficulty.String())
	assert.Equal(t, "test", info.Chain, "they should be equal")
}

func TestBtcGetBlock(t *testing.T) {
	//https://tbtc.bitaps.com/000000000000000f965504aec02813eef6cf7056a1e919d24e2d2675385865fa
	bitcoin := Bitcoin{}
	block := bitcoin.BtcGetBlock("000000000000000f965504aec02813eef6cf7056a1e919d24e2d2675385865fa")
	fmt.Println(block.Hash)
	fmt.Println(block.Txs[1])
	fmt.Println(block.Txs[0].Inputs[0].Sequence)
	fmt.Println(block.Txs[0].Inputs[0].Prevout.Index)
}

func TestBtcGetBlockHash(t *testing.T) {
	//https://tbtc.bitaps.com/000000000000000f965504aec02813eef6cf7056a1e919d24e2d2675385865fa
	bitcoin := Bitcoin{}
	hash := bitcoin.BtcGetBlockHash(1904473)
	fmt.Println(hash)
	assert.Equal(t, "000000000000000f965504aec02813eef6cf7056a1e919d24e2d2675385865fa", hash.Hash, "they should be equal")
}

func TestBtcGetUTXO(t *testing.T) {
	bitcoin := Bitcoin{}
	utxo := bitcoin.BtcGetUTXO("3c28b7a187cf7cde733139b6a49cf646fb430855aa2e5aad90b8711eeb62a562", 1)
	fmt.Println(utxo)
	assert.Equal(t, "n3DS6fdDBCva3AKz12YmDSeBnNCF7Tuwvv", utxo.Address, "they should be equal")
}

func TestBtcGetTxForAccount(t *testing.T) {
	bitcoin := Bitcoin{}
	txs := bitcoin.BtcGetTxForAccount("mfbPS2yrNc1fopS9aHwPNJeQHqrpFw9wLW", 4, 0)
	fmt.Println(len(txs))

	assert.Equal(t, "3c28b7a187cf7cde733139b6a49cf646fb430855aa2e5aad90b8711eeb62a562", txs[0].Hash, "they should be equal")

	assert.Equal(t, "mfbPS2yrNc1fopS9aHwPNJeQHqrpFw9wLW", txs[0].Inputs[0].Coin.Address, "they should be equal")
	assert.Equal(t, uint64(200000), txs[0].Inputs[0].Coin.Value, "they should be equal")

	assert.Equal(t, "n3XsLoopG4pbQRh2PBRJMqq5tU1zCB62kF", txs[0].Outputs[0].Address, "they should be equal")
	assert.Equal(t, uint64(50000), txs[0].Outputs[0].Value, "they should be equal")

	assert.Equal(t, "n3DS6fdDBCva3AKz12YmDSeBnNCF7Tuwvv", txs[0].Outputs[1].Address, "they should be equal")
	assert.Equal(t, uint64(148000), txs[0].Outputs[1].Value, "they should be equal")

	assert.Equal(t, uint64(2000), txs[0].Fee, "they should be equal")
}

func TestBtcGetTransaction(t *testing.T) {
	bitcoin := Bitcoin{}
	tx := bitcoin.BtcGetTransaction("3c28b7a187cf7cde733139b6a49cf646fb430855aa2e5aad90b8711eeb62a562")

	assert.Equal(t, "3c28b7a187cf7cde733139b6a49cf646fb430855aa2e5aad90b8711eeb62a562", tx.Hash, "they should be equal")

	assert.Equal(t, "mfbPS2yrNc1fopS9aHwPNJeQHqrpFw9wLW", tx.Inputs[0].Coin.Address, "they should be equal")
	assert.Equal(t, uint64(200000), tx.Inputs[0].Coin.Value, "they should be equal")

	assert.Equal(t, "n3XsLoopG4pbQRh2PBRJMqq5tU1zCB62kF", tx.Outputs[0].Address, "they should be equal")
	assert.Equal(t, uint64(50000), tx.Outputs[0].Value, "they should be equal")

	assert.Equal(t, "n3DS6fdDBCva3AKz12YmDSeBnNCF7Tuwvv", tx.Outputs[1].Address, "they should be equal")
	assert.Equal(t, uint64(148000), tx.Outputs[1].Value, "they should be equal")

	assert.Equal(t, uint64(2000), tx.Fee, "they should be equal")
}
