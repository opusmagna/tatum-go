package wallet

import (
	"github.com/btcsuite/btcd/chaincfg"
	bch "github.com/gcash/bchd/chaincfg"
	"github.com/tatumio/tatum-go/model/request"
	networkcfg "github.com/tatumio/tatum-go/network"
	"github.com/tatumio/tatum-go/private_key"
	"github.com/tatumio/tatum-go/utils"
)

type Wallet struct {

	/**
	 * mnemonic seed
	 */
	Mnemonic string

	/**
	 * extended public key to derive addresses from
	 */
	Xpub string

	/**
	 * address
	 */
	Address string

	/**
	 * secret
	 */
	Secret string
}

/**
 * Generate Bitcoin io.tatum.wallet
 *
 * @param testnet testnet or mainnet version of address
 * @param mnem    mnemonic seed to use
 * @return the wallet
 * @throws ExecutionException   the execution exception
 * @throws InterruptedException the interrupted exception
 * @returns io.tatum.wallet
 */
func generateBtcWallet(testnet bool, mnemonic string) *Wallet {
	var network *chaincfg.Params
	var path string
	if testnet {
		network = &chaincfg.TestNet3Params
		path = utils.TestnetDerivationPath
	} else {
		network = &chaincfg.MainNetParams
		path = utils.BtcDerivationPath
	}

	xpub := private_key.NewBtcPrivateKey().Network(network).FromSeed(mnemonic).DerivePath(path).Xpub()
	return &Wallet{Mnemonic: mnemonic, Xpub: xpub}
}

/**
 * Generate Litecoin wallet
 *
 * @param testnet testnet or mainnet version of address
 * @param mnem    mnemonic seed to use
 * @return the wallet
 * @throws ExecutionException   the execution exception
 * @throws InterruptedException the interrupted exception
 * @returns wallet
 */
func generateLtcWallet(testnet bool, mnemonic string) *Wallet {
	var network *chaincfg.Params
	var (
		path string
	)
	if testnet {
		network = &networkcfg.LtcTestNet4Params
		path = utils.TestnetDerivationPath
	} else {
		network = &networkcfg.LtcMainNetParams
		path = utils.LtcDerivationPath
	}

	xpub := private_key.NewBtcPrivateKey().Network(network).FromSeed(mnemonic).DerivePath(path).Xpub()
	return &Wallet{Mnemonic: mnemonic, Xpub: xpub}
}

/**
 * Generate bch wallet wallet.
 *
 * @param testnet the testnet
 * @param mnem    the mnem
 * @return the wallet
 * @throws ExecutionException   the execution exception
 * @throws InterruptedException the interrupted exception
 */
func generateBchWallet(testnet bool, mnemonic string) *Wallet {
	var network *bch.Params
	var path string
	if testnet {
		network = &bch.TestNet3Params
	} else {
		network = &bch.MainNetParams
	}
	path = utils.BchDerivationPath

	xpub := private_key.NewBchPrivateKey().Network(network).FromSeed(mnemonic).DerivePath(path).Xpub()
	return &Wallet{Mnemonic: mnemonic, Xpub: xpub}
}

/**
 * Generate wallet
 *
 * @param currency blockchain to generate wallet for
 * @param testnet  testnet or mainnet version of address
 * @param mnemonic mnemonic seed to use. If not present, new one will be generated
 * @return the wallet
 * @throws Exception the exception
 * @returns wallet or a combination of address and private key
 */
func GenerateWallet(currency request.Currency, testnet bool, mnemonic string) *Wallet {
	switch currency {
	case request.BTC:
		return generateBtcWallet(testnet, mnemonic)
	case request.LTC:
		return generateLtcWallet(testnet, mnemonic)
	case request.BCH:
		return generateBchWallet(testnet, mnemonic)
	default:
		return &Wallet{}
	}
}
