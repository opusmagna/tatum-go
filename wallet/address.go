package wallet

import (
	"fmt"

	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/btcutil/hdkeychain"
	"github.com/btcsuite/btcd/chaincfg"
	bch "github.com/gcash/bchd/chaincfg"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	bchkeychain "github.com/gcash/bchutil/hdkeychain"
	"github.com/opusmagna/tatum-go/model/request"
	"github.com/opusmagna/tatum-go/network/ltc"
	"github.com/opusmagna/tatum-go/private_key"
	"github.com/opusmagna/tatum-go/utils"
	sha3 "golang.org/x/crypto/sha3"
)

/**
 * Generate Bitcoin address
 *
 * @param testnet testnet or mainnet version of address
 * @param xpub    extended public key to generate address from
 * @param i       derivation index of address to generate. Up to 2^32 addresses can be generated.
 * @return the string
 * @throws ExecutionException   the execution exception
 * @throws InterruptedException the interrupted exception
 * @returns blockchain address
 */
func generateBtcAddress(testnet bool, xpub string, i uint32) string {
	var network *chaincfg.Params
	if testnet {
		network = &chaincfg.TestNet3Params
	} else {
		network = &chaincfg.MainNetParams
	}
	return generateAddress(network, xpub, i)
}

/**
 * Generate Bitcoin address
 *
 * @param testnet testnet or mainnet version of address
 * @param xpub    extended public key to generate address from
 * @param i       derivation index of address to generate. Up to 2^32 addresses can be generated.
 * @return the string
 * @throws ExecutionException   the execution exception
 * @throws InterruptedException the interrupted exception
 * @returns blockchain address
 */
func generateLtcAddress(testnet bool, xpub string, i uint32) string {
	var network *chaincfg.Params
	if testnet {
		network = &ltc.LtcTestNet4Params
	} else {
		network = &ltc.LtcMainNetParams
	}
	return generateAddress(network, xpub, i)
}

/**
 * Generate Bitcoin Cash address
 *
 * @param testnet testnet or mainnet version of address
 * @param xpub    extended public key to generate address from
 * @param i       derivation index of address to generate. Up to 2^32 addresses can be generated.
 * @return the string
 * @throws ExecutionException   the execution exception
 * @throws InterruptedException the interrupted exception
 * @returns blockchain address
 */
func generateBchAddress(testnet bool, xpub string, i uint32) string {
	var network *bch.Params
	if testnet {
		network = &bch.TestNet3Params
	} else {
		network = &bch.MainNetParams
	}

	key, err := bchkeychain.NewKeyFromString(xpub)
	if err != nil {
		fmt.Println(err)
		return utils.EmptySpace
	}

	key, err = key.Child(i)
	if err != nil {
		fmt.Println(err)
		return utils.EmptySpace
	}

	addressPubKeyHash, err := key.Address(network)
	if err != nil {
		fmt.Println(err)
		return utils.EmptySpace
	}

	return network.CashAddressPrefix + ":" + addressPubKeyHash.EncodeAddress()
}

/**
 * Generate Ethereum or any other ERC20 address
 *
 * @param testnet testnet or mainnet version of address
 * @param xpub    extended public key to generate address from
 * @param i       derivation index of address to generate. Up to 2^32 addresses can be generated.
 * @return the string
 * @throws ExecutionException   the execution exception
 * @throws InterruptedException the interrupted exception
 * @returns blockchain address
 */
func generateEthAddress(testnet bool, xpub string, i uint32) string {
	key, err := hdkeychain.NewKeyFromString(xpub)
	if err != nil {
		fmt.Println(err)
		return utils.WhiteSpace
	}

	key, err = key.Derive(i)
	if err != nil {
		fmt.Println(err)
		return utils.EmptySpace
	}

	pubKey, err := key.ECPubKey()
	if err != nil {
		fmt.Println(err)
		return utils.EmptySpace
	}

	publicKeyBytes := crypto.FromECDSAPub(pubKey.ToECDSA())
	hash := sha3.NewLegacyKeccak256()
	hash.Write(publicKeyBytes[1:])

	return hexutil.Encode(hash.Sum(nil)[12:])
}

func generateAddress(network *chaincfg.Params, xpub string, i uint32) string {
	key, err := hdkeychain.NewKeyFromString(xpub)
	if err != nil {
		fmt.Println(err)
		return utils.WhiteSpace
	}

	key, err = key.Derive(i)
	if err != nil {
		fmt.Println(err)
		return utils.EmptySpace
	}

	pubKey, err := key.ECPubKey()
	if err != nil {
		fmt.Println(err)
		return utils.EmptySpace
	}

	addr, err := btcutil.NewAddressPubKey(pubKey.SerializeCompressed(), network)
	if err != nil {
		fmt.Println(err)
		return utils.EmptySpace
	}

	return addr.EncodeAddress()
}

/**
 * Generate Bitcoin private key from mnemonic seed
 * @param testnet testnet or mainnet version of address
 * @param mnemonic mnemonic to generate private key from
 * @param i derivation index of private key to generate.
 * @returns blockchain private key to the address
 */
func generateBtcPrivateKey(testnet bool, mnemonic string, i uint32) string {
	var network *chaincfg.Params
	var path string
	if testnet {
		network = &chaincfg.TestNet3Params
		path = utils.TestnetDerivationPath
	} else {
		network = &chaincfg.MainNetParams
		path = utils.BtcDerivationPath
	}

	builder := private_key.NewBtcPrivateKey().
		Network(network).
		FromSeed(mnemonic).
		DerivePath(path).
		Derive(i)

	return builder.ToWIF()
}

/**
 * Generate Litecoin private key from mnemonic seed
 * @param testnet testnet or mainnet version of address
 * @param mnemonic mnemonic to generate private key from
 * @param i derivation index of private key to generate.
 * @returns blockchain private key to the address
 */
func generateLtcPrivateKey(testnet bool, mnemonic string, i uint32) string {
	var network *chaincfg.Params
	var path string
	if testnet {
		network = &ltc.LtcTestNet4Params
		path = utils.TestnetDerivationPath
	} else {
		network = &ltc.LtcMainNetParams
		path = utils.LtcDerivationPath
	}

	builder := private_key.NewBtcPrivateKey().
		Network(network).
		FromSeed(mnemonic).
		DerivePath(path).
		Derive(i)

	return builder.ToWIF()
}

/**
 * Generate Bitcoin Cash private key from mnemonic seed
 * @param testnet testnet or mainnet version of address
 * @param mnemonic mnemonic to generate private key from
 * @param i derivation index of private key to generate.
 * @returns blockchain private key to the address
 */
func generateBchPrivateKey(testnet bool, mnemonic string, i uint32) string {
	var network *bch.Params
	var path string
	if testnet {
		network = &bch.TestNet3Params
	} else {
		network = &bch.MainNetParams
	}
	path = utils.BchDerivationPath

	builder := private_key.NewBchPrivateKey().
		Network(network).
		FromSeed(mnemonic).
		DerivePath(path).
		Derive(i)

	return builder.ToWIF()
}

/**
 * Generate Ethereum or any other ERC20 private key from mnemonic seed
 *
 * @param testnet  testnet or mainnet version of address
 * @param mnemonic mnemonic to generate private key from
 * @param i        derivation index of private key to generate.
 * @return the string
 * @throws ExecutionException   the execution exception
 * @throws InterruptedException the interrupted exception
 * @returns blockchain private key to the address
 */
func generateEthPrivateKey(testnet bool, mnemonic string, i uint32) string {
	var network *chaincfg.Params
	network = &chaincfg.MainNetParams

	var (
		path string
	)

	if testnet {
		path = utils.TestnetDerivationPath
	} else {
		path = utils.EthDerivationPath
	}

	builder := private_key.NewBtcPrivateKey().
		Network(network).
		FromSeed(mnemonic).
		DerivePath(path).
		Derive(i)

	return "0x" + builder.ToHex()
}

/**
 * Generate VeChain private key from mnemonic seed
 * @param testnet testnet or mainnet version of address
 * @param mnemonic mnemonic to generate private key from
 * @param i derivation index of private key to generate.
 * @returns blockchain private key to the address
 */
func generateVetPrivateKey(testnet bool, mnemonic string, i uint32) string {
	var network *chaincfg.Params
	network = &chaincfg.MainNetParams

	var (
		path string
	)

	if testnet {
		path = utils.TestnetDerivationPath
	} else {
		path = utils.VetDerivationPath
	}

	builder := private_key.NewBtcPrivateKey().
		Network(network).
		FromSeed(mnemonic).
		DerivePath(path).
		Derive(i)

	return "0x" + builder.ToHex()
}

/**
 * Generate Tron private key from mnemonic seed
 *
 * @param mnemonic mnemonic to generate private key from
 * @param i        derivation index of private key to generate.
 * @returns blockchain private key to the address
 */
func generateTronPrivateKey(testnet bool, mnemonic string, i uint32) string {
	var network *chaincfg.Params
	network = &chaincfg.MainNetParams

	var (
		path string
	)

	path = utils.TronDerivationPath

	builder := private_key.NewBtcPrivateKey().
		Network(network).
		FromSeed(mnemonic).
		DerivePath(path).
		Derive(i)

	return builder.ToHex()
}

/**
 * Generate address
 *
 * @param currency type of blockchain
 * @param testnet  testnet or mainnet version of address
 * @param xpub     extended public key to generate address from
 * @param i        derivation index of address to generate. Up to 2^32 addresses can be generated.
 * @return the string
 * @throws Exception the exception
 * @returns blockchain address
 */
func GenerateAddressFromXPub(currency request.Currency, testnet bool, xpub string, i uint32) string {
	switch currency {
	case request.BTC:
		return generateBtcAddress(testnet, xpub, i)
	case request.LTC:
		return generateLtcAddress(testnet, xpub, i)
	case request.BCH:
		return generateBchAddress(testnet, xpub, i)
	case request.USDT:
		fallthrough
	case request.WBTC:
		fallthrough
	case request.LEO:
		fallthrough
	case request.LINK:
		fallthrough
	case request.UNI:
		fallthrough
	case request.FREE:
		fallthrough
	case request.MKR:
		fallthrough
	case request.USDC:
		fallthrough
	case request.BAT:
		fallthrough
	case request.TUSD:
		fallthrough
	case request.PAX:
		fallthrough
	case request.PAXG:
		fallthrough
	case request.PLTC:
		fallthrough
	case request.XCON:
		fallthrough
	case request.ETH:
		fallthrough
	case request.MMY:
		fallthrough
	case request.VET:
		return generateEthAddress(testnet, xpub, i)
	default:
		return utils.EmptySpace
	}
}

/**
 * Generate private key from mnemonic seed
 *
 * @param currency type of blockchain
 * @param testnet  testnet or mainnet version of address
 * @param mnemonic mnemonic to generate private key from
 * @param i        derivation index of private key to generate.
 * @return the string
 * @throws Exception the exception
 * @returns blockchain private key to the address
 */
func GeneratePrivateKeyFromMnemonic(currency request.Currency, testnet bool, mnemonic string, i uint32) string {
	switch currency {
	case request.BTC:
		return generateBtcPrivateKey(testnet, mnemonic, i)
	case request.LTC:
		return generateLtcPrivateKey(testnet, mnemonic, i)
	case request.BCH:
		return generateBchPrivateKey(testnet, mnemonic, i)
	case request.USDT:
		fallthrough
	case request.WBTC:
		fallthrough
	case request.LEO:
		fallthrough
	case request.LINK:
		fallthrough
	case request.UNI:
		fallthrough
	case request.FREE:
		fallthrough
	case request.MKR:
		fallthrough
	case request.USDC:
		fallthrough
	case request.BAT:
		fallthrough
	case request.TUSD:
		fallthrough
	case request.PAX:
		fallthrough
	case request.PAXG:
		fallthrough
	case request.PLTC:
		fallthrough
	case request.XCON:
		fallthrough
	case request.ETH:
		fallthrough
	case request.MMY:
		return generateEthPrivateKey(testnet, mnemonic, i)
	case request.VET:
		return generateVetPrivateKey(testnet, mnemonic, i)
	case request.TRON:
		return generateTronPrivateKey(testnet, mnemonic, i)
	default:
		return utils.WhiteSpace
	}
}
