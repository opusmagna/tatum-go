package offchain

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/go-playground/validator"
	"github.com/opusmagna/tatum-go/model/request"
	"github.com/opusmagna/tatum-go/model/response/offchain"
	"github.com/opusmagna/tatum-go/transaction/bitcoin_tx_builder"
	"github.com/opusmagna/tatum-go/wallet"
)

type BitcoinOffchain struct {
}

var validate *validator.Validate

/**
 * Send Bitcoin transaction from Tatum Ledger account to the blockchain. This method broadcasts signed transaction to the blockchain.
 * This operation is irreversible.
 *
 * @param testnet mainnet or testnet version
 * @param body    content of the transaction to broadcast
 * @return the broadcast result
 * @throws Exception the exception
 * @returns transaction id of the transaction in the blockchain
 */
func (b *BitcoinOffchain) SendBitcoinOffchainTransaction(testnet bool, body request.TransferBtcBasedOffchain) (*offchain.BroadcastResult, error) {
	validate = validator.New()
	err := validate.Struct(body)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err.Field() + ": should have " + err.Tag() + " " + err.Param())
			fmt.Println(err.Value())
		}
		return nil, err
	}

	withdrawal := body.Withdrawal
	if len(withdrawal.Fee) == 0 {
		withdrawal.Fee = "0.0005"
	}

	withdrawalResponse := StoreWithdrawalOffchain(*withdrawal)
	id := withdrawalResponse.Id

	txData, err := b.PrepareBitcoinSignedOffchainTransaction(testnet,
		withdrawalResponse.Data,
		withdrawal.Amount,
		withdrawal.Address,
		body.Mnemonic,
		body.KeyPair,
		withdrawal.Attr,
		withdrawal.MultipleAmounts)

	if err != nil {
		_, _ = CancelWithdrawalOffchain(id, true)
		return nil, err
	}

	broadcastWithdrawal := request.BroadcastWithdrawal{}
	broadcastWithdrawal.TxData = txData
	broadcastWithdrawal.WithdrawalId = id
	broadcastWithdrawal.Currency = request.BTC.String()

	txHash, err := BroadcastOffchain(broadcastWithdrawal)
	if err != nil {
		_, err1 := CancelWithdrawalOffchain(id, true)
		if err1 == nil {
			return nil, err
		}
		return &offchain.BroadcastResult{TxHash: nil, Id: id}, nil
	}
	return &offchain.BroadcastResult{TxHash: txHash, Id: id}, nil

}

/**
 * Sign Bitcoin transaction with private keys locally. Nothing is broadcast to the blockchain.
 * @param testnet mainnet or testnet version
 * @param data data from Tatum system to prepare transaction from
 * @param amount amount to send
 * @param address recipient address, if multiple recipients are present, it should be string separated by ','
 * @param mnemonic mnemonic to sign transaction from. mnemonic or keyPair must be present
 * @param keyPair keyPair to sign transaction from. keyPair or mnemonic must be present
 * @param changeAddress address to send the rest of the unused coins
 * @param multipleAmounts if multiple recipients are present in the address separated by ',', this should be list of amounts to send
 * @returns transaction data to be broadcast to blockchain.
 */
func (b *BitcoinOffchain) PrepareBitcoinSignedOffchainTransaction(testnet bool, data []offchain.WithdrawalResponseData,
	amount string, address string, mnemonic string, keyPair []request.KeyPair,
	changeAddress string, multipleAmounts []string) (string, error) {

	var network *chaincfg.Params
	if testnet {
		network = &chaincfg.TestNet3Params
	} else {
		network = &chaincfg.MainNetParams
	}

	var (
		txBuilder = bitcoin_tx_builder.New().Init(network)
	)

	if len(multipleAmounts) > 0 {
		addresses := strings.Split(address, ",")
		for i := 0; i < len(multipleAmounts); i++ {
			_amount, err := strconv.ParseInt(multipleAmounts[i], 10, 64)
			if err != nil {
				return "", err
			}
			txBuilder.AddOutput(addresses[i], _amount)
		}
	} else {
		_amount, err := strconv.ParseInt(amount, 10, 64)
		if err != nil {
			return "", err
		}
		txBuilder.AddOutput(address, _amount)
	}

	lastVin := findVInFirst(data, func(s string) bool { return "-1" == s })
	last := lastVin.Amount
	if last > 0 {
		if len(mnemonic) > 0 && len(changeAddress) > 0 {
			xpub := wallet.GenerateWallet(request.BTC, testnet, mnemonic).Xpub
			txBuilder.AddOutput(wallet.GenerateAddressFromXPub(request.BTC, testnet, xpub, 0), last)
		} else if len(changeAddress) > 0 {
			txBuilder.AddOutput(changeAddress, last)
		} else {
			return "", errors.New("impossible to prepare transaction. Either mnemonic or keyPair and attr must be present")
		}
	}

	for _, input := range data {
		if input.VIn != "-1" {
			if len(mnemonic) > 0 {
				var derivationKey uint32
				if input.Address != nil {
					derivationKey = input.Address.DerivationKey
				} else {
					derivationKey = 0
				}
				privKey := wallet.GeneratePrivateKeyFromMnemonic(request.BTC, testnet, mnemonic, derivationKey)
				txBuilder.AddInput(input.VIn, input.VInIndex, privKey)
			} else if len(keyPair) > 0 {
				pair := findKeyPairFirst(keyPair, func(s string) bool { return input.Address.Address == s })
				txBuilder.AddInput(input.VIn, input.VInIndex, pair.PrivateKey)
			} else {
				return "", errors.New("impossible to prepare transaction. Either mnemonic or keyPair and attr must be present")
			}
		}
	}

	return txBuilder.Sign().ToHex(), nil
}

func findVInFirst(arr []offchain.WithdrawalResponseData, f func(string) bool) (ret *offchain.WithdrawalResponseData) {
	for _, s := range arr {
		if f(s.VIn) {
			return &s
		}
	}
	return nil
}

func findKeyPairFirst(keyPair []request.KeyPair, f func(string) bool) (ret *request.KeyPair) {
	for _, s := range keyPair {
		if f(s.Address) {
			return &s
		}
	}
	return nil
}
