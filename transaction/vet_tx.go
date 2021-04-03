package transaction

import (
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/go-playground/validator"
	"github.com/tatumio/tatum-go/blockchain"
	"github.com/tatumio/tatum-go/model/request"
	"github.com/tatumio/tatum-go/model/response/common"
	"github.com/vechain/thor/thor"
	"github.com/vechain/thor/tx"
	"math/big"
	"strconv"
)

type VetTx struct {
}

/**
 * Send VeChain transaction to the blockchain. This method broadcasts signed transaction to the blockchain.
 * This operation is irreversible.
 * @param testnet mainnet or testnet version
 * @param body content of the transaction to broadcast
 * @param provider url of the VeChain Server to connect to. If not set, default public server will be used.
 * @returns transaction id of the transaction in the blockchain
 */
func (v *VetTx) sendVetTransaction(testnet bool, body request.TransferVet, provider string) *common.TransactionHash {
	vet := &blockchain.Vet{}
	txData, err := v.PrepareVetSignedTransaction(testnet, body, provider)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return vet.VetBroadcast(txData, "")
}

/**
 * Sign VeChain transaction with private keys locally. Nothing is broadcast to the blockchain.
 * @param testnet mainnet or testnet version
 * @param body content of the transaction to broadcast
 * @param provider url of the VeChain Server to connect to. If not set, default public server will be used.
 * @returns transaction data to be broadcast to blockchain.
 */
func (v *VetTx) PrepareVetSignedTransaction(testnet bool, body request.TransferVet, provider string) (string, error) {
	validate = validator.New()
	err := validate.Struct(body)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err.Field() + ": should have " + err.Tag() + " " + err.Param())
			fmt.Println(err.Value())
		}
		return "", err
	}

	to, _ := thor.ParseAddress(body.To)
	value, ok := new(big.Int).SetString(body.Amount, 10)
	if !ok {
		return "", errors.New("invalid amount")
	}

	var gas uint64
	if body.Fee != nil {
		gas, err = strconv.ParseUint(body.Fee.GasLimit, 10, 64)
	} else {
		gas = 21000
	}

	trx := new(tx.Builder).ChainTag(1).
		BlockRef(tx.BlockRef{0, 0, 0, 0, 0, 0, 0, 0}).
		Expiration(32).
		Clause(tx.NewClause(&to).WithValue(value).WithData([]byte(body.Data))).
		GasPriceCoef(128).
		Gas(gas).
		DependsOn(nil).Build()

	k, err := hex.DecodeString(body.FromPrivateKey)
	if err != nil {
		return "", err
	}

	priv, err := crypto.ToECDSA(k)
	if err != nil {
		return "", err
	}

	sig, err := crypto.Sign(trx.SigningHash().Bytes(), priv)
	if err != nil {
		return "", err
	}

	trx = trx.WithSignature(sig)

	d, err := rlp.EncodeToBytes(trx)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(d), nil
}
