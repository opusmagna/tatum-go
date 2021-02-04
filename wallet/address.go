package wallet

import (
	"fmt"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
	"github.com/btcsuite/btcutil/hdkeychain"
	bch "github.com/gcash/bchd/chaincfg"
	bchkeychain "github.com/gcash/bchutil/hdkeychain"
	ltc "github.com/ltcsuite/ltcd/chaincfg"
	"github.com/ltcsuite/ltcutil"
	ltckeychain "github.com/ltcsuite/ltcutil/hdkeychain"
	"github.com/tatumio/tatum-go/model/request"
	"github.com/tatumio/tatum-go/private_key"
	"github.com/tatumio/tatum-go/utils"
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
	var network *ltc.Params
	if testnet {
		network = &ltc.TestNet4Params
	} else {
		network = &ltc.MainNetParams
	}

	key, err := ltckeychain.NewKeyFromString(xpub)
	if err != nil {
		fmt.Println(err)
		return utils.EmptySpace
	}

	key, err = key.Child(i)
	if err != nil {
		fmt.Println(err)
		return utils.EmptySpace
	}

	pubKey, err := key.ECPubKey()
	if err != nil {
		fmt.Println(err)
		return utils.EmptySpace
	}

	addr, err := ltcutil.NewAddressPubKey(pubKey.SerializeCompressed(), network)
	if err != nil {
		fmt.Println(err)
		return utils.EmptySpace
	}

	return addr.EncodeAddress()
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
	var network *ltc.Params
	var path string
	if testnet {
		network = &ltc.TestNet4Params
		path = utils.TestnetDerivationPath
	} else {
		network = &ltc.MainNetParams
		path = utils.LtcDerivationPath
	}

	builder := private_key.NewLtcPrivateKey().
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
	default:
		return utils.WhiteSpace
	}
}
