package bitcoin_tx_builder

import (
	"bytes"
	"encoding/hex"
	"fmt"

	"github.com/btcsuite/btcutil"

	"github.com/btcsuite/btcd/wire"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/txscript"
)

type Transaction interface {
	ToHex() string
}

type TransactionBuilder interface {
	Init(net *chaincfg.Params) TransactionBuilder
	AddOutput(address string, amount int64) TransactionBuilder
	AddInput(txHash string, index uint32, key string) TransactionBuilder
	Sign() Transaction
}

type transactionBuilder struct {
	net               *chaincfg.Params
	privateKeysToSign []string
	amountsToSign     []int64
	tx                *wire.MsgTx
}

func (tb *transactionBuilder) Init(net *chaincfg.Params) TransactionBuilder {
	tb.net = net
	tb.tx = wire.NewMsgTx(wire.TxVersion)
	tb.privateKeysToSign = make([]string, 0)
	return tb
}

func (tb *transactionBuilder) AddOutput(address string, amount int64) TransactionBuilder {
	fmt.Println(amount)
	var (
		toAddress, err = btcutil.DecodeAddress(address, tb.net)
	)
	if err != nil {
		fmt.Println(err.Error())
		return tb
	}

	pkScript, err := txscript.PayToAddrScript(toAddress)
	if err != nil {
		fmt.Println(err.Error())
		return tb
	}

	// Add an output paying to the address.
	tb.tx.AddTxOut(&wire.TxOut{
		Value:    amount,
		PkScript: pkScript,
	})

	return tb
}

func (tb *transactionBuilder) AddInput(txHash string, index uint32, key string) TransactionBuilder {

	utxoHash, err := chainhash.NewHashFromStr(txHash)
	if err != nil {
		fmt.Println(err.Error())
		return tb
	}

	// Add the input(s) the redeeming transaction will spend.
	// The second argument is vout or Tx-index, which is the index
	// of spending UTXO in the transaction that Txid referred to
	prevOut := wire.NewOutPoint(utxoHash, index)
	txIn := wire.NewTxIn(prevOut, nil, nil)
	tb.tx.AddTxIn(txIn)

	tb.privateKeysToSign = append(tb.privateKeysToSign, key)

	return tb
}

func (tb *transactionBuilder) Sign() Transaction {

	// Sign the new transaction.
	for i := range tb.tx.TxIn {
		key := tb.privateKeysToSign[i]
		pkScript, wif, err := createPkScript(key, tb.net)
		if err != nil {
			return &transaction{}
		}

		disasm, _ := txscript.DisasmString(pkScript)
		fmt.Println("Script Disassembly:", disasm)

		sigScript, err := txscript.SignatureScript(tb.tx, i, pkScript, txscript.SigHashAll, wif.PrivKey, false)
		if err != nil {
			return &transaction{}
		}
		tb.tx.TxIn[i].SignatureScript = sigScript

		// Prove that the transaction has been validly signed by executing the
		// script pair.
		flags := txscript.StandardVerifyFlags
		vm, err := txscript.NewEngine(pkScript, tb.tx, i, flags, nil, nil, tb.tx.TxOut[i].Value)
		if err != nil {
			return &transaction{}
		}
		if err := vm.Execute(); err != nil {
			return &transaction{}
		}
	}

	var signedTx bytes.Buffer

	err := tb.tx.Serialize(&signedTx)
	if err != nil {
		return &transaction{}
	}

	return &transaction{signedTx}
}

func createPkScript(privKey string, net *chaincfg.Params) ([]byte, *btcutil.WIF, error) {

	wif, err := btcutil.DecodeWIF(privKey)
	if err != nil {
		return nil, nil, err
	}

	addressPubkey, err := btcutil.NewAddressPubKey(wif.PrivKey.PubKey().SerializeUncompressed(), net)
	if err != nil {
		return nil, nil, err
	}

	fromAddress, err := btcutil.DecodeAddress(addressPubkey.EncodeAddress(), net)
	if err != nil {
		return nil, nil, err
	}

	pkScript, err := txscript.PayToAddrScript(fromAddress)
	if err != nil {
		return nil, nil, err
	}

	extractPkScriptAddrs(hex.EncodeToString(pkScript))

	return pkScript, wif, nil
}

func extractPkScriptAddrs(scriptHex string) {
	// Start with a standard pay-to-pubkey-hash script.
	script, err := hex.DecodeString(scriptHex)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Extract and print details from the script.
	scriptClass, addresses, reqSigs, err := txscript.ExtractPkScriptAddrs(
		script, &chaincfg.MainNetParams)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Script Class:", scriptClass)
	fmt.Println("Addresses:", addresses)
	fmt.Println("Required Signatures:", reqSigs)
}

func New() TransactionBuilder {
	return &transactionBuilder{}
}

type transaction struct {
	bitcoinSerialize bytes.Buffer
}

func (t *transaction) ToHex() string {
	return hex.EncodeToString(t.bitcoinSerialize.Bytes())
}
