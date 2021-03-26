package private_key

import (
	"fmt"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
	btc "github.com/btcsuite/btcutil/hdkeychain"
	"github.com/tatumio/tatum-go/utils"
	"regexp"
	"strconv"
	"strings"
)

type BtcPrivateKey interface {
	Network(net *chaincfg.Params) BtcPrivateKey
	FromSeed(mnemonic string) BtcPrivateKey
	DerivePath(path string) BtcPrivateKey
	Derive(i uint32) BtcPrivateKey
	ToWIF() string
	Xpub() string
}

type btcPrivateKey struct {
	net  *chaincfg.Params
	seed []byte
	key  *btc.ExtendedKey
}

func NewBtcPrivateKey() BtcPrivateKey {
	return &btcPrivateKey{}
}

func (p *btcPrivateKey) Network(net *chaincfg.Params) BtcPrivateKey {
	p.net = net
	return p
}

func (p *btcPrivateKey) FromSeed(mnemonic string) BtcPrivateKey {
	re := regexp.MustCompile("\\s+")
	words := re.Split(mnemonic, -1)
	p.seed = ToSeed(words, utils.EmptySpace)
	return p
}

func (p *btcPrivateKey) DerivePath(path string) BtcPrivateKey {
	var err error
	p.key, err = btc.NewMaster(p.seed, p.net)
	if err != nil {
		fmt.Println(err)
		return &btcPrivateKey{}
	}

	trimPrefix := strings.Replace(path, "M", utils.EmptySpace, -1)
	parsedNodes := strings.Split(trimPrefix, utils.SEPARATOR)
	for _, n := range parsedNodes {
		n = strings.Replace(n, utils.WhiteSpace, utils.EmptySpace, -1)
		if len(n) == 0 {
			continue
		}
		isHard := strings.HasSuffix(n, "H")
		if isHard {
			n = strings.TrimSuffix(n, "H")
		}

		nodeNumber, err := strconv.Atoi(n)
		if err != nil {
			fmt.Println(err)
			return &btcPrivateKey{}
		}

		var r uint32
		if isHard {
			r = btc.HardenedKeyStart + uint32(nodeNumber)
		} else {
			r = uint32(nodeNumber)
		}
		p.key, err = p.key.Derive(r)
		if err != nil {
			fmt.Println(err)
			return &btcPrivateKey{}
		}

		//pubKey, _ := p.key.Neuter()
		//fmt.Println(pubKey.String())
	}
	return p
}

func (p *btcPrivateKey) Derive(i uint32) BtcPrivateKey {
	var err error
	p.key, err = p.key.Derive(i)
	if err != nil {
		fmt.Println(err)
		return &btcPrivateKey{}
	}
	return p
}

func (p *btcPrivateKey) ToWIF() string {
	privKey, err := p.key.ECPrivKey()
	if err != nil {
		fmt.Println(err)
		return utils.EmptySpace
	}

	wif, err := btcutil.NewWIF(privKey, p.net, true)
	if err != nil {
		fmt.Println(err)
		return utils.EmptySpace
	}

	return wif.String()
}

func (p *btcPrivateKey) Xpub() string {
	pubKey, err := p.key.Neuter()
	if err != nil {
		fmt.Println(err)
		return utils.EmptySpace
	}

	return pubKey.String()
}
