package private_key

import (
	"fmt"
	"github.com/nvxtien/ltcutil"
	"github.com/nvxtien/ltcutil/hdkeychain"
	ltctatum "github.com/tatumio/ltcd/chaincfg"
	"github.com/tatumio/tatum-go/utils"
	"regexp"
	"strconv"
	"strings"
)

type LtcPrivateKey interface {
	Network(net *ltctatum.Params) LtcPrivateKey
	FromSeed(mnemonic string) LtcPrivateKey
	DerivePath(path string) LtcPrivateKey
	Derive(i uint32) LtcPrivateKey
	ToWIF() string
	Xpub() string
}

type ltcPrivateKey struct {
	net  *ltctatum.Params
	seed []byte
	key  *hdkeychain.ExtendedKey
}

func NewLtcPrivateKey() LtcPrivateKey {
	return &ltcPrivateKey{}
}

func (p *ltcPrivateKey) Network(net *ltctatum.Params) LtcPrivateKey {
	p.net = net
	return p
}

func (p *ltcPrivateKey) FromSeed(mnemonic string) LtcPrivateKey {
	re := regexp.MustCompile("\\s+")
	words := re.Split(mnemonic, -1)
	p.seed = ToSeed(words, utils.EmptySpace)
	return p
}

func (p *ltcPrivateKey) DerivePath(path string) LtcPrivateKey {
	var err error
	p.key, err = hdkeychain.NewMaster(p.seed, p.net)
	if err != nil {
		fmt.Println(err)
		return &ltcPrivateKey{}
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
			return &ltcPrivateKey{}
		}

		var r uint32
		if isHard {
			r = hdkeychain.HardenedKeyStart + uint32(nodeNumber)
		} else {
			r = uint32(nodeNumber)
		}
		p.key, err = p.key.Child(r)
		if err != nil {
			fmt.Println(err)
			return &ltcPrivateKey{}
		}
		pubKey, _ := p.key.Neuter()
		fmt.Println(pubKey.String())
	}
	return p
}

func (p *ltcPrivateKey) Derive(i uint32) LtcPrivateKey {
	var err error
	p.key, err = p.key.Child(i)
	if err != nil {
		fmt.Println(err)
		return &ltcPrivateKey{}
	}
	return p
}

func (p *ltcPrivateKey) ToWIF() string {
	privKey, err := p.key.ECPrivKey()
	if err != nil {
		fmt.Println(err)
		return utils.EmptySpace
	}

	wif, err := ltcutil.NewWIF(privKey, p.net, true)
	if err != nil {
		fmt.Println(err)
		return utils.EmptySpace
	}

	return wif.String()
}

func (p *ltcPrivateKey) Xpub() string {
	pubKey, err := p.key.Neuter()
	if err != nil {
		fmt.Println(err)
		return utils.EmptySpace
	}

	return pubKey.String()
}
