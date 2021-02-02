package transaction

import (
	"fmt"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
	"github.com/go-playground/validator"
	"github.com/tatumio/tatum-go/blockchain"
	"github.com/tatumio/tatum-go/model/request"
	"github.com/tatumio/tatum-go/model/response/common"
	"github.com/tatumio/tatum-go/transaction/bitcoin_tx_builder"
)

type LitecoinTx struct {
}

func (b *LitecoinTx) prepareSignedTransaction(network *chaincfg.Params, body request.TransferBtcBasedBlockchain) (string, error) {
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

	var litecoin = blockchain.Litecoin{}
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
			txs := litecoin.LtcGetTxForAccount(fromAddress[i].Address, 50, 0)
			for j := range txs {
				outputs := txs[j].Outputs
				for k := range outputs {
					if outputs[k].Address == fromAddress[i].Address {
						utxo := litecoin.LtcGetUTXO(txs[j].Hash, outputs[k].Value)
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
 * Sign Litcoin transaction with private keys locally. Nothing is broadcast to the blockchain.
 * @param testnet mainnet or testnet version
 * @param body content of the transaction to broadcast
 * @returns transaction data to be broadcast to blockchain.
 */
func (b *LitecoinTx) PrepareLitecoinSignedTransaction(testnet bool, body request.TransferBtcBasedBlockchain) (string, error) {
	var network *chaincfg.Params
	if testnet {
		network = &chaincfg.TestNet3Params
	} else {
		network = &chaincfg.MainNetParams
	}
	return b.prepareSignedTransaction(network, body)
}

/**
 * Send Litecoin transaction to the blockchain. This method broadcasts signed transaction to the blockchain.
 * This operation is irreversible.
 * @param testnet mainnet or testnet version
 * @param body content of the transaction to broadcast
 * @returns transaction id of the transaction in the blockchain
 */
func (b *LitecoinTx) SendLitecoinTransaction(testnet bool, body request.TransferBtcBasedBlockchain) *common.TransactionHash {
	bitcoin := &blockchain.Bitcoin{}
	txData, err := b.PrepareLitecoinSignedTransaction(testnet, body)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return bitcoin.BtcBroadcast(txData, "")
}
