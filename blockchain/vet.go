package blockchain

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator"
	"github.com/tatumio/tatum-go/model/request"
	"github.com/tatumio/tatum-go/model/response/common"
	"github.com/tatumio/tatum-go/model/response/vet"
	"strconv"
)

type Vet struct {
}

var validate *validator.Validate

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/VetBroadcast" target="_blank">Tatum API documentation</a>
 */
func (v *Vet) VetBroadcast(txData string, signatureId string) *common.TransactionHash {
	url := "/v3/vet/broadcast"

	payload := make(map[string]interface{})
	payload["txData"] = txData
	if len(signatureId) > 0 {
		payload["signatureId"] = signatureId
	}

	requestJSON, err := json.Marshal(payload)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	fmt.Println(string(requestJSON))

	txHash := common.TransactionHash{}
	var result map[string]interface{}
	res, err := sender.SendPost(url, requestJSON)
	if err == nil {
		if err = json.Unmarshal([]byte(res), &result); err == nil {
			txHash.TxId = fmt.Sprint(result["txId"])
		}
	}
	return &txHash
}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/VetEstimateGas" target="_blank">Tatum API documentation</a>
 */
func (v *Vet) VetEstimateGas(body request.EstimateGasVet) *vet.VetEstimateGas {
	validate = validator.New()
	err := validate.Struct(body)

	url := "/v3/vet/transaction/gas"
	requestJSON, err := json.Marshal(body)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	estGas := vet.VetEstimateGas{}
	res, err := sender.SendPost(url, requestJSON)
	if err == nil {
		err = json.Unmarshal([]byte(res), &estGas)
	}

	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return &estGas

}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/VetGetCurrentBlock" target="_blank">Tatum API documentation</a>
 */
func (v *Vet) VetGetCurrentBlock() uint64 {
	url := "/v3/vet/current"

	res, err := sender.SendGet(url, nil)
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}

	current, err := strconv.ParseUint(res, 10, 64)
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}
	return current
}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/VetGetBlock" target="_blank">Tatum API documentation</a>
 */
func (v *Vet) VetGetBlock(hash string) *vet.VetBlock {
	url := "/v3/vet/block/" + hash

	res, err := sender.SendGet(url, nil)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	var block vet.VetBlock
	if err == nil {
		err = json.Unmarshal([]byte(res), &block)
	}

	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return &block
}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/VetGetBalance" target="_blank">Tatum API documentation</a>
 */
func (v *Vet) VetGetAccountBalance(address string) *common.Balance {
	url := "/v3/vet/account/balance/" + address

	var balance common.Balance
	res, err := sender.SendGet(url, nil)
	if err != nil {
		fmt.Println(err.Error())
		return &balance
	}

	err = json.Unmarshal([]byte(res), &balance)
	if err != nil {
		fmt.Println(err.Error())
		return &balance
	}
	return &balance
}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/VetGetEnergy" target="_blank">Tatum API documentation</a>
 */
func (v *Vet) VetGetAccountEnergy(address string) *vet.Energy {
	url := "/v3/vet/account/energy/" + address

	var energy vet.Energy
	res, err := sender.SendGet(url, nil)
	if err != nil {
		fmt.Println(err.Error())
		return &energy
	}

	err = json.Unmarshal([]byte(res), &energy)
	if err != nil {
		fmt.Println(err.Error())
		return &energy
	}
	return &energy
}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/VetGetTransaction" target="_blank">Tatum API documentation</a>
 */
func (v *Vet) VetGetTransaction(hash string) *vet.Tx {
	url := "/v3/vet/transaction/" + hash

	res, err := sender.SendGet(url, nil)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	var tx vet.Tx
	if err == nil {
		err = json.Unmarshal([]byte(res), &tx)
	}

	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return &tx
}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/VetGetTransactionReceipt" target="_blank">Tatum API documentation</a>
 */
func (v *Vet) VetGetTransactionReceipt(hash string) *vet.VetTxReceipt {
	url := "/v3/vet/transaction/" + hash + "/receipt"

	res, err := sender.SendGet(url, nil)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	var receipt vet.VetTxReceipt
	if err == nil {
		err = json.Unmarshal([]byte(res), &receipt)
	}

	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return &receipt

}
