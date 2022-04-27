package offchain

import (
	"encoding/json"
	"fmt"

	r "github.com/opusmagna/tatum-go/model/response/offchain"
)

/*
 */
func CreateNewDepositAddress(accountId string) (*r.Address, error) {

	_url := fmt.Sprintf("/v3/offchain/%s/address", accountId)

	var res string
	var body []byte
	var responseAddress r.Address
	res, err := sender.SendPost(_url, body)
	if err == nil {
		err = json.Unmarshal([]byte(res), &responseAddress)
	}

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return &responseAddress, nil

}
