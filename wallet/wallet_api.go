package wallet

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/opusmagna/tatum-go/utils"
)

const (
	//TODO add all the rest here matey
	USDT = "tron"
)

var sender = &utils.Async{}

var CurrencyBasedPaths = map[string]string{
	"USDT": "tron",
}

type WalletApi struct {
	BasePath string
	PathSlug string
}

func GetInstance(slug string) (*WalletApi, error) {
	key, ok := CurrencyBasedPaths[slug]

	if ok == false {
		return nil, errors.New("Unsupported currency slug")
	}

	return &WalletApi{
		BasePath: key,
	}, nil
}

func (wapi *WalletApi) GenerateWallet(mnemonic string) (*Wallet, error) {

	url := wapi.setUrl(mnemonic)
	res, err := sender.SendGet(url, nil)

	var w Wallet
	err = json.Unmarshal([]byte(res), &w)
	if err != nil {
		return nil, err
	}
	return &w, nil
}

func (wapi *WalletApi) setUrl(mnemonic string) string {
	return fmt.Sprintf("%v/%v?mnemonic=%v", wapi.BasePath, "wallet", mnemonic)
}
