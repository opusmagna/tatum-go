package transaction

import (
	"fmt"

	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/go-playground/validator"
	"github.com/opusmagna/tatum-go/blockchain"
	"github.com/opusmagna/tatum-go/model/request"
	"github.com/opusmagna/tatum-go/model/response/common"
	"github.com/opusmagna/tatum-go/transaction/bitcoin_tx_builder"
)

type BitcoinTx struct {
}

var validate *validator.Validate

/**
 * Sign Bitcoin Cash transaction with private keys locally. Nothing is broadcast to the blockchain.
 * @param testnet mainnet or testnet version
 * @param body content of the transaction to broadcast
 * @returns transaction data to be broadcast to blockchain.
 */
func (b *BitcoinTx) prepareSignedTransaction(network *chaincfg.Params, body request.TransferBtcBasedBlockchain) (string, error) {
	validate = validator.New()
	err := validate.Struct(body)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err.Field() + ": should have " + err.Tag() + " " + err.Param())
			fmt.Println(err.Value())
		}
		return "", err
	}

	if len(body.FromAddress) == 0 && len(body.FromUTXO) == 0 {
		return "", fmt.Errorf("both addresses and utxo must not be empty")
	}

	if len(body.FromAddress) > 0 && len(body.FromUTXO) > 0 {
		return "", fmt.Errorf("only accept from either addresses or utxo")
	}

	fromUTXO := body.FromUTXO
	to := body.To
	fromAddress := body.FromAddress

	var bitcoin = blockchain.Bitcoin{}
	var (
		transactionBuilder = bitcoin_tx_builder.New().Init(network)
	)

	for i := range to {
		var value, err = btcutil.NewAmount(to[i].Value)
		if err != nil {
			return "", err
		}
		transactionBuilder.AddOutput(to[i].Address, int64(value.ToUnit(btcutil.AmountSatoshi)))
	}

	if len(fromAddress) > 0 {
		for i := range fromAddress {
			txs := bitcoin.BtcGetTxForAccount(fromAddress[i].Address, 50, 0)
			for j := range txs {
				outputs := txs[j].Outputs
				for k := range outputs {
					if outputs[k].Address == fromAddress[i].Address {
						utxo := bitcoin.BtcGetUTXO(txs[j].Hash, outputs[k].Value)
						fmt.Println(utxo)
						transactionBuilder.AddInput(txs[j].Hash, uint32(k), fromAddress[i].PrivateKey)
					}
				}
			}
		}
	} else if len(fromUTXO) > 0 {
		for i := range fromUTXO {
			transactionBuilder.AddInput(fromUTXO[i].TxHash, fromUTXO[i].Index, fromUTXO[i].PrivateKey)
		}
	}

	return transactionBuilder.Sign().ToHex(), nil
}

/**
 * Sign Bitcoin transaction with private keys locally. Nothing is broadcast to the blockchain.
 * @param testnet mainnet or testnet version
 * @param body content of the transaction to broadcast
 * @returns transaction data to be broadcast to blockchain.
 */
func (b *BitcoinTx) PrepareBitcoinSignedTransaction(testnet bool, body request.TransferBtcBasedBlockchain) (string, error) {
	var network *chaincfg.Params
	if testnet {
		network = &chaincfg.TestNet3Params
	} else {
		network = &chaincfg.MainNetParams
	}
	return b.prepareSignedTransaction(network, body)
}

/**
 * Send Bitcoin transaction to the blockchain. This method broadcasts signed transaction to the blockchain.
 * This operation is irreversible.
 * @param testnet mainnet or testnet version
 * @param body content of the transaction to broadcast
 * @returns transaction id of the transaction in the blockchain
 */
func (b *BitcoinTx) SendBitcoinTransaction(testnet bool, body request.TransferBtcBasedBlockchain) *common.TransactionHash {
	bitcoin := &blockchain.Bitcoin{}
	txData, err := b.PrepareBitcoinSignedTransaction(testnet, body)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return bitcoin.BtcBroadcast(txData, "")
}
