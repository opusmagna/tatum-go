package offchain

import (
	"errors"
	"fmt"
	bch "github.com/gcash/bchd/chaincfg"
	"github.com/go-playground/validator"
	"github.com/tatumio/tatum-go/model/request"
	"github.com/tatumio/tatum-go/model/response/offchain"
	"github.com/tatumio/tatum-go/transaction/bcash_tx_builder"
	"github.com/tatumio/tatum-go/wallet"
	"strconv"
	"strings"
)

type BcashOffchain struct {
}

/**
 * Send Bitcoin Cash transaction from Tatum Ledger account to the blockchain. This method broadcasts signed transaction to the blockchain.
 * This operation is irreversible.
 * @param testnet mainnet or testnet version
 * @param body content of the transaction to broadcast
 * @returns transaction id of the transaction in the blockchain or id of the withdrawal, if it was not cancelled automatically
 */
func (bc *BcashOffchain) SendBitcoinCashOffchainTransaction(testnet bool, body request.TransferBtcBasedOffchain) (*offchain.BroadcastResult, error) {
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

	txData, err := bc.PrepareBitcoinCashSignedOffchainTransaction(testnet,
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
 * Sign Bitcoin Cash transaction with private keys locally. Nothing is broadcast to the blockchain.
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
func (bc *BcashOffchain) PrepareBitcoinCashSignedOffchainTransaction(testnet bool, data []offchain.WithdrawalResponseData,
	amount string, address string, mnemonic string, keyPair []request.KeyPair,
	changeAddress string, multipleAmounts []string) (string, error) {

	var network *bch.Params
	if testnet {
		network = &bch.TestNet3Params
	} else {
		network = &bch.MainNetParams
	}

	var (
		txBuilder = bcash_tx_builder.New().Init(network)
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
			xpub := wallet.GenerateWallet(request.BCH, testnet, mnemonic).Xpub
			txBuilder.AddOutput(wallet.GenerateAddressFromXPub(request.BCH, testnet, xpub, 0), last)
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
				privKey := wallet.GeneratePrivateKeyFromMnemonic(request.BCH, testnet, mnemonic, derivationKey)
				txBuilder.AddInput(input.VIn, input.VInIndex, privKey, input.Amount)
			} else if len(keyPair) > 0 {
				pair := findKeyPairFirst(keyPair, func(s string) bool { return input.Address.Address == s })
				txBuilder.AddInput(input.VIn, input.VInIndex, pair.PrivateKey, input.Amount)
			} else {
				return "", errors.New("impossible to prepare transaction. Either mnemonic or keyPair and attr must be present")
			}
		}
	}

	return txBuilder.Sign().ToHex(), nil
}
