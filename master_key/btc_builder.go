package master_key

import (
	"fmt"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil/hdkeychain"
	"github.com/tatumio/tatum-go/private_key"
	"github.com/tatumio/tatum-go/utils"
	"regexp"
)

type BtcMasterKey interface {
	Network(net *chaincfg.Params) BtcMasterKey
	FromSeed(mnemonic string) BtcMasterKey
	String() string
}

type btcMasterKey struct {
	net  *chaincfg.Params
	seed []byte
	key  *hdkeychain.ExtendedKey
}

func NewBtcMasterKey() BtcMasterKey {
	return &btcMasterKey{}
}

func (p *btcMasterKey) Network(net *chaincfg.Params) BtcMasterKey {
	p.net = net
	return p
}

func (p *btcMasterKey) FromSeed(mnemonic string) BtcMasterKey {
	re := regexp.MustCompile("\\s+")
	words := re.Split(mnemonic, -1)
	p.seed = private_key.ToSeed(words, utils.EmptySpace)
	return p
}

func (p *btcMasterKey) String() string {
	var err error
	p.key, err = hdkeychain.NewMaster(p.seed, p.net)
	if err != nil {
		fmt.Println(err)
		return utils.EmptySpace
	}

	pubKey, err := p.key.Neuter()
	if err != nil {
		fmt.Println(err)
		return utils.EmptySpace
	}

	return pubKey.String()
}
