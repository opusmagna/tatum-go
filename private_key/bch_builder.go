package private_key

import (
	"fmt"
	"github.com/gcash/bchd/chaincfg"
	"github.com/gcash/bchutil"
	"github.com/gcash/bchutil/hdkeychain"
	"github.com/tatumio/tatum-go/utils"
	"regexp"
	"strconv"
	"strings"
)

type BchPrivateKey interface {
	Network(net *chaincfg.Params) BchPrivateKey
	FromSeed(mnemonic string) BchPrivateKey
	DerivePath(path string) BchPrivateKey
	Derive(i uint32) BchPrivateKey
	ToWIF() string
}

type bchPrivateKey struct {
	net  *chaincfg.Params
	seed []byte
	key  *hdkeychain.ExtendedKey
}

func NewBchPrivateKey() BchPrivateKey {
	return &bchPrivateKey{}
}

func (p *bchPrivateKey) Network(net *chaincfg.Params) BchPrivateKey {
	p.net = net
	return p
}

func (p *bchPrivateKey) FromSeed(mnemonic string) BchPrivateKey {
	re := regexp.MustCompile("\\s+")
	words := re.Split(mnemonic, -1)
	p.seed = ToSeed(words, utils.EmptySpace)
	return p
}

func (p *bchPrivateKey) DerivePath(path string) BchPrivateKey {
	var err error
	p.key, err = hdkeychain.NewMaster(p.seed, p.net)
	if err != nil {
		fmt.Println(err)
		return &bchPrivateKey{}
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
			return &bchPrivateKey{}
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
			return &bchPrivateKey{}
		}
	}
	return p
}

func (p *bchPrivateKey) Derive(i uint32) BchPrivateKey {
	var err error
	p.key, err = p.key.Child(i)
	if err != nil {
		fmt.Println(err)
		return &bchPrivateKey{}
	}
	return p
}

func (p *bchPrivateKey) ToWIF() string {
	privKey, err := p.key.ECPrivKey()
	if err != nil {
		fmt.Println(err)
		return utils.EmptySpace
	}

	wif, err := bchutil.NewWIF(privKey, p.net, true)
	if err != nil {
		fmt.Println(err)
		return utils.EmptySpace
	}

	return wif.String()
}
