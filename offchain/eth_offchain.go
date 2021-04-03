package offchain

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/go-playground/validator"
	"github.com/tatumio/tatum-go/ledger"
	"github.com/tatumio/tatum-go/model/request"
	"github.com/tatumio/tatum-go/model/response/offchain"
	"github.com/tatumio/tatum-go/transaction"
	"github.com/tatumio/tatum-go/utils"
	"github.com/tatumio/tatum-go/wallet"
	"golang.org/x/crypto/sha3"
	"log"
	"math/big"
	"os"
	"strings"
)

type EthereumOffchain struct {
}

/**
 * Send Ethereum transaction from Tatum Ledger account to the blockchain. This method broadcasts signed transaction to the blockchain.
 * This operation is irreversible.
 * @param testnet mainnet or testnet version
 * @param body content of the transaction to broadcast
 * @param provider url of the Ethereum Server to connect to. If not set, default public server will be used.
 * @returns transaction id of the transaction in the blockchain or id of the withdrawal, if it was not cancelled automatically
 */
func (e *EthereumOffchain) SendEthOffchainTransaction(testnet bool, body request.TransferEthErc20Offchain, provider string) (*offchain.BroadcastResult, error) {

	validate = validator.New()
	err := validate.Struct(body)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err.Field() + ": should have " + err.Tag() + " " + err.Param())
			fmt.Println(err.Value())
		}
		return nil, err
	}

	var rawurl string
	if len(provider) > 0 {
		rawurl = provider
	} else {
		rawurl = os.Getenv("TATUM_API_URL") + "/v3/ethereum/web3/" + os.Getenv("TATUM_API_KEY")
	}

	client, err := ethclient.Dial(rawurl)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	var withdrawal *request.CreateWithdrawal
	withdrawal.Amount = body.BaseTransferEthErc20Offchain.Amount
	withdrawal.Address = body.BaseTransferEthErc20Offchain.Address
	withdrawal.SenderAccountId = body.BaseTransferEthErc20Offchain.SenderAccountId
	withdrawal.Compliant = body.BaseTransferEthErc20Offchain.Compliant
	withdrawal.PaymentId = body.BaseTransferEthErc20Offchain.PaymentId
	withdrawal.SenderNote = body.BaseTransferEthErc20Offchain.SenderNote

	mnemonic := body.Mnemonic
	index := body.Index
	privateKey := body.PrivateKey
	var fromPriv string
	if len(mnemonic) > 0 && index > 0 {
		fromPriv = wallet.GeneratePrivateKeyFromMnemonic(request.ETH, testnet, mnemonic, index)
	} else if len(privateKey) > 0 {
		fromPriv = privateKey
	} else {
		return nil, errors.New("no mnemonic or private key is present")
	}

	var gasPrice *big.Int
	var ok bool
	if len(body.BaseTransferEthErc20Offchain.GasPrice) > 0 {
		gasPrice, ok = new(big.Int).SetString(body.BaseTransferEthErc20Offchain.GasPrice, 10)
		if !ok {
			return nil, errors.New("invalid gas price")
		}

		gasPrice = new(big.Int).Mul(gasPrice, big.NewInt(params.GWei)) // GWei to Wei
	} else {
		gasPrice = transaction.EthGetGasPriceInWei()
	}

	var accountLedger = ledger.AccountLedger{}
	account := accountLedger.GetAccountById(withdrawal.SenderAccountId)

	amount := withdrawal.Amount
	address := withdrawal.Address

	txData, gasLimit, err := e.PrepareEthSignedOffchainTransaction(testnet, amount, fromPriv, address, account.Currency, client,
		gasPrice.String(), body.BaseTransferEthErc20Offchain.Nonce)
	if err != nil {
		return nil, err
	}

	var _gasLimit *big.Int
	if len(body.BaseTransferEthErc20Offchain.GasLimit) > 0 {
		_gasLimit, ok = new(big.Int).SetString(body.BaseTransferEthErc20Offchain.GasLimit, 10)
		if !ok {
			return nil, err
		}
	} else {
		_gasLimit = new(big.Int).SetUint64(gasLimit)
	}

	feeInWei := new(big.Int).Mul(_gasLimit, gasPrice)
	feeInEther := new(big.Int).Div(feeInWei, big.NewInt(params.Ether))
	withdrawal.Fee = feeInEther.String()

	withdrawalResponse := OffchainStoreWithdrawal(*withdrawal)
	id := withdrawalResponse.Id

	broadcastWithdrawal := request.BroadcastWithdrawal{}
	broadcastWithdrawal.TxData = txData
	broadcastWithdrawal.WithdrawalId = id
	broadcastWithdrawal.Currency = request.BTC.String()

	txHash, err := OffchainBroadcast(broadcastWithdrawal)
	if err != nil {
		_, err1 := OffchainCancelWithdrawal(id, true)
		if err1 == nil {
			return nil, err
		}
		return &offchain.BroadcastResult{TxHash: nil, Id: id}, nil
	}
	return &offchain.BroadcastResult{TxHash: txHash, Id: id}, nil
}

/**
 * Send Ethereum ERC20 transaction from Tatum Ledger account to the blockchain. This method broadcasts signed transaction to the blockchain.
 * This operation is irreversible.
 * @param testnet mainnet or testnet version
 * @param body content of the transaction to broadcast
 * @param provider url of the Ethereum Server to connect to. If not set, default public server will be used.
 * @returns transaction id of the transaction in the blockchain or id of the withdrawal, if it was not cancelled automatically
 */
func (e *EthereumOffchain) SendEthErc20OffchainTransaction(testnet bool, body request.TransferEthErc20Offchain, provider string) (*offchain.BroadcastResult, error) {
	validate = validator.New()
	err := validate.Struct(body)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err.Field() + ": should have " + err.Tag() + " " + err.Param())
			fmt.Println(err.Value())
		}
		return nil, err
	}

	var rawurl string
	if len(provider) > 0 {
		rawurl = provider
	} else {
		rawurl = os.Getenv("TATUM_API_URL") + "/v3/ethereum/web3/" + os.Getenv("TATUM_API_KEY")
	}

	client, err := ethclient.Dial(rawurl)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	var withdrawal *request.CreateWithdrawal
	withdrawal.Amount = body.BaseTransferEthErc20Offchain.Amount
	withdrawal.Address = body.BaseTransferEthErc20Offchain.Address
	withdrawal.SenderAccountId = body.BaseTransferEthErc20Offchain.SenderAccountId
	withdrawal.Compliant = body.BaseTransferEthErc20Offchain.Compliant
	withdrawal.PaymentId = body.BaseTransferEthErc20Offchain.PaymentId
	withdrawal.SenderNote = body.BaseTransferEthErc20Offchain.SenderNote

	mnemonic := body.Mnemonic
	index := body.Index
	privateKey := body.PrivateKey
	var fromPriv string
	if len(mnemonic) > 0 && index > 0 {
		fromPriv = wallet.GeneratePrivateKeyFromMnemonic(request.ETH, testnet, mnemonic, index)
	} else if len(privateKey) > 0 {
		fromPriv = privateKey
	} else {
		return nil, errors.New("no mnemonic or private key is present")
	}

	var gasPrice *big.Int
	var ok bool
	if len(body.BaseTransferEthErc20Offchain.GasPrice) > 0 {
		gasPrice, ok = new(big.Int).SetString(body.BaseTransferEthErc20Offchain.GasPrice, 10)
		if !ok {
			return nil, errors.New("invalid gas price")
		}

		gasPrice = new(big.Int).Mul(gasPrice, big.NewInt(params.GWei)) // GWei to Wei
	} else {
		gasPrice = transaction.EthGetGasPriceInWei()
	}

	var accountLedger = ledger.AccountLedger{}
	account := accountLedger.GetAccountById(withdrawal.SenderAccountId)

	if findEthBased(request.EthBasedCurrencies, account.Currency) {
		return e.SendEthOffchainTransaction(testnet, body, provider)
	}

	v := ledger.VC{}
	vc := v.GetVirtualCurrencyByName(account.Currency)

	txData, gasLimit, err := e.PrepareEthErc20SignedOffchainTransaction(testnet, withdrawal.Amount, fromPriv,
		withdrawal.Address, client, vc.Erc20Address, gasPrice.String(), body.BaseTransferEthErc20Offchain.Nonce)

	if err != nil {
		return nil, err
	}

	var _gasLimit *big.Int
	if len(body.BaseTransferEthErc20Offchain.GasLimit) > 0 {
		_gasLimit, ok = new(big.Int).SetString(body.BaseTransferEthErc20Offchain.GasLimit, 10)
		if !ok {
			return nil, err
		}
	} else {
		_gasLimit = new(big.Int).SetUint64(gasLimit)
	}

	feeInWei := new(big.Int).Mul(_gasLimit, gasPrice)
	feeInEther := new(big.Int).Div(feeInWei, big.NewInt(params.Ether))
	withdrawal.Fee = feeInEther.String()

	withdrawalResponse := OffchainStoreWithdrawal(*withdrawal)
	id := withdrawalResponse.Id

	broadcastWithdrawal := request.BroadcastWithdrawal{}
	broadcastWithdrawal.TxData = txData
	broadcastWithdrawal.WithdrawalId = id
	broadcastWithdrawal.Currency = request.BTC.String()

	txHash, err := OffchainBroadcast(broadcastWithdrawal)
	if err != nil {
		_, err1 := OffchainCancelWithdrawal(id, true)
		if err1 == nil {
			return nil, err
		}
		return &offchain.BroadcastResult{TxHash: nil, Id: id}, nil
	}
	return &offchain.BroadcastResult{TxHash: txHash, Id: id}, nil
}

func findEthBased(arr []string, f string) bool {
	for _, s := range arr {
		return s == f
	}
	return false
}

/**
 * Sign Ethereum transaction with private keys locally. Nothing is broadcast to the blockchain.
 * @param amount amount to send
 * @param privateKey private key to sign transaction and send funds from
 * @param address recipient address
 * @param currency Ethereum or supported ERC20
 * @param web3 instance of the web3 client
 * @param gasPrice gas price of the blockchain fee
 * @param nonce nonce of the transaction. this is counter of the transactions from given address. should be + 1 from previous one.
 * @returns transaction data to be broadcast to blockchain.
 */
func (e *EthereumOffchain) PrepareEthSignedOffchainTransaction(testnet bool, amount string, privateKey string, address string,
	currency string, client *ethclient.Client, gasPrice string, nonce uint64) (string, uint64, error) {

	var to common.Address
	var value *big.Int
	var data []byte

	if currency == request.ETH {
		to = common.HexToAddress(address)
		value = utils.Ether2Wei(amount)

	} else {
		to = common.HexToAddress(utils.ContractAddresses()(currency))
		receiver := common.HexToAddress(address)
		// Token transfers don't require ETH to be transferred so set the value to 0
		value = big.NewInt(0)

		transferFnSignature := []byte("transfer(address,uint256)")
		hash := sha3.NewLegacyKeccak256()
		hash.Write(transferFnSignature)
		methodID := hash.Sum(nil)[:4]

		paddedAddress := common.LeftPadBytes(receiver.Bytes(), 32)

		digits := math.BigPow(10, int64(utils.ContractDecimals()(currency)))
		_amount := new(big.Int)
		_amount.SetString(amount, 10)
		_amount = new(big.Int).Mul(_amount, digits)
		paddedAmount := common.LeftPadBytes(_amount.Bytes(), 32)

		data = append(data, methodID...)
		data = append(data, paddedAddress...)
		data = append(data, paddedAmount...)
	}

	_gasPrice, ok := new(big.Int).SetString(gasPrice, 10)
	if !ok {
		return "", 0, errors.New("invalid gas price")
	}

	gasLimit, err := getGasLimit(client, to, data)
	if err != nil {
		return "", 0, err
	}

	rawTx, err := createRawTransaction(testnet, privateKey, nonce, to, value, gasLimit, _gasPrice, data)
	if err != nil {
		return "", 0, err
	}

	return rawTx, gasLimit, nil

}

/**
 * Sign Ethereum custom ERC20 transaction with private keys locally. Nothing is broadcast to the blockchain.
 * @param amount amount to send
 * @param privateKey private key to sign transaction and send funds from
 * @param address recipient address
 * @param tokenAddress blockchain address of the custom ERC20 token
 * @param web3 instance of the web3 client
 * @param gasPrice gas price of the blockchain fee
 * @param nonce nonce of the transaction. this is counter of the transactions from given address. should be + 1 from previous one.
 * @returns transaction data to be broadcast to blockchain.
 */

func (e *EthereumOffchain) PrepareEthErc20SignedOffchainTransaction(testnet bool, amount string, privateKey string, address string,
	client *ethclient.Client, tokenAddress string, gasPrice string, nonce uint64) (string, uint64, error) {

	var value *big.Int
	var data []byte

	to := common.HexToAddress(tokenAddress)
	receiver := common.HexToAddress(address)
	// Token transfers don't require ETH to be transferred so set the value to 0
	value = big.NewInt(0)

	transferFnSignature := []byte("transfer(address,uint256)")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]

	paddedAddress := common.LeftPadBytes(receiver.Bytes(), 32)

	digits := math.BigPow(10, 18)
	_amount := new(big.Int)
	_amount.SetString(amount, 10)
	_amount = new(big.Int).Mul(_amount, digits)
	paddedAmount := common.LeftPadBytes(_amount.Bytes(), 32)

	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	gasLimit, err := getGasLimit(client, to, data)
	if err != nil {
		return "", 0, err
	}

	_gasPrice, ok := new(big.Int).SetString(gasPrice, 10)
	if !ok {
		return "", 0, errors.New("invalid gas price")
	}

	rawTx, err := createRawTransaction(testnet, privateKey, nonce, to, value, gasLimit, _gasPrice, data)
	if err != nil {
		return "", 0, err
	}

	return rawTx, gasLimit, nil
}

func createRawTransaction(testnet bool, prv string,
	nonce uint64, to common.Address, amount *big.Int, gasLimit uint64, gasPrice *big.Int, data []byte) (string, error) {

	privateKey, err := crypto.HexToECDSA(strings.Replace(prv, "0x", "", 1))
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	tx := types.NewTransaction(nonce, to, amount, gasLimit, gasPrice, data)

	var chainID *big.Int
	if testnet {
		chainID = big.NewInt(3) // Ropsten
	} else {
		chainID = big.NewInt(1) // Mainnet
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	rawTxBytes, err := rlp.EncodeToBytes(signedTx)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	rawTxHex := hex.EncodeToString(rawTxBytes)

	fmt.Printf(rawTxHex)

	return rawTxHex, nil
}

func getGasLimit(client *ethclient.Client, to common.Address, data []byte) (uint64, error) {
	gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		To:   &to,
		Data: data,
	})

	if err != nil {
		return 0, err
	}

	return gasLimit + 5000, nil
}
