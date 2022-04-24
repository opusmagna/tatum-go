package transaction

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/go-playground/validator"
	"github.com/opusmagna/tatum-go/blockchain"
	"github.com/opusmagna/tatum-go/contracts/erc20"
	"github.com/opusmagna/tatum-go/model/request"
	response "github.com/opusmagna/tatum-go/model/response/common"
	"github.com/opusmagna/tatum-go/utils"
	sha3 "golang.org/x/crypto/sha3"
)

type EthTx struct {
}

/**
 * Sign Ethereum Store data transaction with private keys locally. Nothing is broadcast to the blockchain.
 * @param testnet mainnet or testnet version
 * @param body content of the transaction to broadcast
 * @param provider url of the Ethereum Server to connect to. If not set, default public server will be used.
 * @returns transaction data to be broadcast to blockchain.
 */
func (b *EthTx) PrepareStoreDataTransaction(testnet bool, body request.CreateRecord, provider string) (string, error) {
	validate = validator.New()
	err := validate.Struct(body)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err.Field() + ": should have " + err.Tag() + " " + err.Param())
			fmt.Println(err.Value())
		}
		return "", err
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
		return "", err
	}

	privateKey, err := crypto.HexToECDSA(strings.Replace(body.FromPrivateKey, "0x", "", -1))
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	var address common.Address
	if len(body.To) > 0 {
		address = common.HexToAddress(body.To)
	} else {
		publicKey := privateKey.Public()
		publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
		if !ok {
			log.Fatal("error casting public key to ECDSA")
			return "", err
		}
		address = crypto.PubkeyToAddress(*publicKeyECDSA)
	}

	nonce := body.Nonce
	if nonce == 0 {
		var ethr = blockchain.Ethereum{}
		nonce = ethr.EthGetTransactionsCount(address.String())
	}

	gasLimit := uint64(len(body.Data)*68 + 21000)
	gasPrice := utils.EthGetGasPriceInWei()

	return createRawTransaction(client, body.FromPrivateKey, body.Nonce, address, big.NewInt(0), gasLimit, gasPrice, []byte(body.Data))
}

/**
 * Sign Ethereum or supported ERC20 transaction with private keys locally. Nothing is broadcast to the blockchain.
 * @param testnet mainnet or testnet version
 * @param body content of the transaction to broadcast
 * @param provider url of the Ethereum Server to connect to. If not set, default public server will be used.
 * @returns transaction data to be broadcast to blockchain.
 */
func (b *EthTx) PrepareEthOrErc20SignedTransaction(testnet bool, body request.TransferEthErc20, provider string) (string, error) {
	validate = validator.New()
	err := validate.Struct(body)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err.Field() + ": should have " + err.Tag() + " " + err.Param())
			fmt.Println(err.Value())
		}
		return "", err
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
		return "", err
	}

	var to common.Address
	var value *big.Int
	var data []byte

	if body.Currency == request.ETH {
		to = common.HexToAddress(body.To)
		value = utils.Ether2Wei(body.Amount)
		data = []byte(hex.EncodeToString([]byte(body.Data)))

	} else {
		to = common.HexToAddress(utils.ContractAddresses()(body.Currency))
		receiver := common.HexToAddress(body.To)
		// Token transfers don't require ETH to be transferred so set the value to 0
		value = big.NewInt(0)

		transferFnSignature := []byte("transfer(address,uint256)")
		hash := sha3.NewLegacyKeccak256()
		hash.Write(transferFnSignature)
		methodID := hash.Sum(nil)[:4]

		paddedAddress := common.LeftPadBytes(receiver.Bytes(), 32)

		digits := math.BigPow(10, int64(utils.ContractDecimals()(body.Currency)))
		amount := new(big.Int)
		amount.SetString(body.Amount, 10)
		amount = new(big.Int).Mul(amount, digits)
		paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)

		data = append(data, methodID...)
		data = append(data, paddedAddress...)
		data = append(data, paddedAmount...)
	}

	gasPrice, gasLimit, err := getGasPriceAndGasLimit(client, body.Fee, to, data)
	if err != nil {
		return "", err
	}
	return createRawTransaction(client, body.FromPrivateKey, body.Nonce, to, value, gasLimit, gasPrice, data)
}

/**
 * Sign Ethereum custom ERC20 transaction with private keys locally. Nothing is broadcast to the blockchain.
 * @param testnet mainnet or testnet version
 * @param body content of the transaction to broadcast
 * @param provider url of the Ethereum Server to connect to. If not set, default public server will be used.
 * @returns transaction data to be broadcast to blockchain.
 */
func (b *EthTx) PrepareCustomErc20SignedTransaction(testnet bool, body request.TransferCustomErc20, provider string) (string, error) {

	validate = validator.New()
	err := validate.Struct(body)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err.Field() + ": should have " + err.Tag() + " " + err.Param())
			fmt.Println(err.Value())
		}
		return "", err
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
		return "", err
	}

	var to common.Address
	var value *big.Int
	var data []byte

	to = common.HexToAddress(body.ContractAddress)
	receiver := common.HexToAddress(body.To)
	value = big.NewInt(0)

	transferFnSignature := []byte("transfer(address,uint256)")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]
	fmt.Println(hexutil.Encode(methodID)) // 0xa9059cbb

	paddedAddress := common.LeftPadBytes(receiver.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedAddress))

	digits := math.BigPow(10, int64(body.Digits))
	amount := new(big.Int)
	amount.SetString(body.Amount, 10)
	amount = new(big.Int).Mul(amount, digits)

	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedAmount))

	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	gasPrice, gasLimit, err := getGasPriceAndGasLimit(client, body.Fee, to, data)
	if err != nil {
		return "", err
	}
	return createRawTransaction(client, body.FromPrivateKey, body.Nonce, to, value, gasLimit, gasPrice, data)
}

/**
 * Sign Ethereum deploy ERC20 transaction with private keys locally. Nothing is broadcast to the blockchain.
 * @param testnet mainnet or testnet version
 * @param body content of the transaction to broadcast
 * @param provider url of the Ethereum Server to connect to. If not set, default public server will be used.
 * @returns transaction data to be broadcast to blockchain.
 */
func (b *EthTx) PrepareDeployErc20SignedTransaction(testnet bool, body request.DeployEthErc20, provider string) (string, error) {

	validate = validator.New()
	err := validate.Struct(body)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err.Field() + ": should have " + err.Tag() + " " + err.Param())
			fmt.Println(err.Value())
		}
		return "", err
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
		return "", err
	}

	var to common.Address
	var value *big.Int
	var data []byte

	data = append(data, common.LeftPadBytes([]byte(erc20.TokenByteCode), 32)...)
	data = append(data, common.LeftPadBytes([]byte(body.Name), 32)...)
	data = append(data, common.LeftPadBytes([]byte(body.Supply), 32)...)

	receiver := common.HexToAddress(body.Address)
	data = append(data, common.LeftPadBytes(receiver.Bytes(), 32)...)
	data = append(data, common.LeftPadBytes([]byte(strconv.Itoa(body.Digits)), 32)...)

	digits := math.BigPow(10, int64(body.Digits))
	supply := new(big.Int)
	supply.SetString(body.Supply, 10)
	supply = new(big.Int).Mul(supply, digits)
	data = append(data, common.LeftPadBytes(supply.Bytes(), 32)...)
	data = append(data, common.LeftPadBytes(supply.Bytes(), 32)...)

	gasPrice, gasLimit, err := getGasPriceAndGasLimit(client, body.Fee, to, data)
	if err != nil {
		return "", err
	}
	return createRawTransaction(client, body.FromPrivateKey, body.Nonce, to, value, gasLimit, gasPrice, data)
}

func createRawTransaction(client *ethclient.Client, prv string,
	nonce uint64, to common.Address, amount *big.Int, gasLimit uint64, gasPrice *big.Int, data []byte) (string, error) {

	privateKey, err := crypto.HexToECDSA(strings.Replace(prv, "0x", "", 1))
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	tx := types.NewTransaction(nonce, to, amount, gasLimit, gasPrice, data)

	//chainID, err := client.NetworkID(context.Background())
	//	log.Fatal(err)
	//if err != nil {
	//return "", err
	//}

	//signedTx, err := types.SignTx(tx, types.NewEIP155Signer(big.NewInt(1)), privateKey)
	signedTx, err := types.SignTx(tx, types.HomesteadSigner{}, privateKey)
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

func getGasPriceAndGasLimit(client *ethclient.Client, fee *request.Fee, to common.Address, data []byte) (*big.Int, uint64, error) {
	var gasLimit uint64
	var gasPrice *big.Int
	var err error
	var ok bool

	if fee != nil {
		gasPrice, ok = new(big.Int).SetString(fee.GasPrice, 10)
		if !ok {
			return nil, 0, err
		}
	} else {
		gasPrice = utils.EthGetGasPriceInWei()
	}

	if fee != nil {
		gasLimit, err = strconv.ParseUint(fee.GasLimit, 10, 64)
	} else {
		gasLimit, err = client.EstimateGas(context.Background(), ethereum.CallMsg{
			To:   &to,
			Data: data,
		})
	}

	return gasPrice, gasLimit + 5000, nil
}

/**
 * Send Ethereum store data transaction to the blockchain. This method broadcasts signed transaction to the blockchain.
 * This operation is irreversible.
 * @param testnet mainnet or testnet version
 * @param body content of the transaction to broadcast
 * @param provider url of the Ethereum Server to connect to. If not set, default public server will be used.
 * @returns transaction id of the transaction in the blockchain
 */
func (b *EthTx) SendStoreDataTransaction(testnet bool, body request.CreateRecord, provider string) *response.TransactionHash {
	txData, err := b.PrepareStoreDataTransaction(testnet, body, provider)
	if err != nil {
		e := blockchain.Ethereum{}
		return e.EthBroadcast(txData, "")
	}
	return nil
}

/**
 * Send Ethereum or supported ERC20 transaction to the blockchain. This method broadcasts signed transaction to the blockchain.
 * This operation is irreversible.
 * @param testnet mainnet or testnet version
 * @param body content of the transaction to broadcast
 * @param provider url of the Ethereum Server to connect to. If not set, default public server will be used.
 * @returns transaction id of the transaction in the blockchain
 */
func (b *EthTx) SendEthOrErc20Transaction(testnet bool, body request.TransferEthErc20, provider string) *response.TransactionHash {
	txData, err := b.PrepareEthOrErc20SignedTransaction(testnet, body, provider)
	if err != nil {
		e := blockchain.Ethereum{}
		return e.EthBroadcast(txData, "")
	}
	return nil
}

/**
 * Send Ethereum custom ERC20 transaction to the blockchain. This method broadcasts signed transaction to the blockchain.
 * This operation is irreversible.
 * @param testnet mainnet or testnet version
 * @param body content of the transaction to broadcast
 * @param provider url of the Ethereum Server to connect to. If not set, default public server will be used.
 * @returns transaction id of the transaction in the blockchain
 */
func (b *EthTx) SendCustomErc20Transaction(testnet bool, body request.TransferCustomErc20, provider string) *response.TransactionHash {
	txData, err := b.PrepareCustomErc20SignedTransaction(testnet, body, provider)
	if err != nil {
		e := blockchain.Ethereum{}
		return e.EthBroadcast(txData, "")
	}
	return nil
}

/**
 * Send Ethereum deploy ERC20 transaction to the blockchain. This method broadcasts signed transaction to the blockchain.
 * This operation is irreversible.
 * @param testnet mainnet or testnet version
 * @param body content of the transaction to broadcast
 * @param provider url of the Ethereum Server to connect to. If not set, default public server will be used.
 * @returns transaction id of the transaction in the blockchain
 */
func (b *EthTx) SendDeployErc20Transaction(testnet bool, body request.DeployEthErc20, provider string) *response.TransactionHash {
	txData, err := b.PrepareDeployErc20SignedTransaction(testnet, body, provider)
	if err != nil {
		e := blockchain.Ethereum{}
		return e.EthBroadcast(txData, "")
	}
	return nil
}
