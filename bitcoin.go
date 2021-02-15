package main

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
	"github.com/btcsuite/btcutil"
)

type Bitcoin struct {
}

type Transaction struct {
	TxId               string `json:"txid"`
	SourceAddress      string `json:"source_address"`
	DestinationAddress string `json:"destination_address"`
	Amount             int64  `json:"amount"`
	//UnsignedTx         string `json:"unsignedtx"`
	SignedTx string `json:"signedtx"`
}

func (b *Bitcoin) prepareSignedTransaction() {

}

func CreateTransaction(secret string, destination string, amount int64, txHash string) (Transaction, error) {
	var transaction Transaction
	wif, err := btcutil.DecodeWIF(secret) // 10
	if err != nil {
		return Transaction{}, err
	}
	addresspubkey, _ := btcutil.NewAddressPubKey(wif.PrivKey.PubKey().SerializeUncompressed(), &chaincfg.TestNet3Params) // 9

	sourceTx := wire.NewMsgTx(wire.TxVersion)
	sourceUtxoHash, _ := chainhash.NewHashFromStr(txHash)                                                // 14
	sourceUtxo := wire.NewOutPoint(sourceUtxoHash, 0)                                                    // 13
	sourceTxIn := wire.NewTxIn(sourceUtxo, nil, nil)                                                     // 12
	destinationAddress, err := btcutil.DecodeAddress(destination, &chaincfg.MainNetParams)               // 4
	sourceAddress, err := btcutil.DecodeAddress(addresspubkey.EncodeAddress(), &chaincfg.TestNet3Params) // 8
	if err != nil {
		return Transaction{}, err
	}

	destinationPkScript, _ := txscript.PayToAddrScript(destinationAddress) // 3

	sourcePkScript, _ := txscript.PayToAddrScript(sourceAddress) // 7

	sourceTxOut := wire.NewTxOut(amount, sourcePkScript) // 6

	sourceTx.AddTxIn(sourceTxIn)   // 11
	sourceTx.AddTxOut(sourceTxOut) // 5

	sourceTxHash := sourceTx.TxHash() // 4
	fmt.Println("==============")
	fmt.Println(sourceTxHash)
	fmt.Println("=======")

	redeemTx := wire.NewMsgTx(wire.TxVersion)
	prevOut := wire.NewOutPoint(&sourceTxHash, 0) // 3

	redeemTxIn := wire.NewTxIn(prevOut, nil, nil) // 2
	redeemTx.AddTxIn(redeemTxIn)                  // 1

	redeemTxOut := wire.NewTxOut(amount, destinationPkScript) // 2
	redeemTx.AddTxOut(redeemTxOut)                            // 1

	sigScript, err := txscript.SignatureScript(redeemTx, 0, sourceTx.TxOut[0].PkScript, txscript.SigHashAll, wif.PrivKey, false)
	if err != nil {
		return Transaction{}, err
	}
	fmt.Println(hex.EncodeToString(sigScript))
	redeemTx.TxIn[0].SignatureScript = sigScript
	flags := txscript.StandardVerifyFlags
	vm, err := txscript.NewEngine(sourceTx.TxOut[0].PkScript, redeemTx, 0, flags, nil, nil, amount)
	if err != nil {
		return Transaction{}, err
	}
	if err := vm.Execute(); err != nil {
		return Transaction{}, err
	}
	//var unsignedTx bytes.Buffer
	var signedTx bytes.Buffer
	//sourceTx.Serialize(&unsignedTx)

	redeemTx.Serialize(&signedTx)
	transaction.TxId = sourceTxHash.String()

	//transaction.UnsignedTx = hex.EncodeToString(unsignedTx.Bytes())

	disasm, _ := txscript.DisasmString(signedTx.Bytes())

	fmt.Println("Script Disassembly:", disasm)

	transaction.Amount = amount
	transaction.SignedTx = hex.EncodeToString(signedTx.Bytes())
	transaction.SourceAddress = sourceAddress.EncodeAddress()
	transaction.DestinationAddress = destinationAddress.EncodeAddress()
	return transaction, nil
}

func main() {
	transaction, err := CreateTransaction("5HusYj2b2x4nroApgfvaSfKYZhRbKFH41bVyPooymbC6KfgSXdD", "1KKKK6N21XKo48zWKuQKXdvSsCf95ibHFa", 91234, "81b4c832d70cb56ff957589752eb4125a4cab78a25a8fc52d6a09e5bd4404d48")
	if err != nil {
		fmt.Println(err)
		return
	}
	data, _ := json.Marshal(transaction)
	fmt.Println(string(data))
}
