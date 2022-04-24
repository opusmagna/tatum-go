package wallet

import (
	"github.com/btcsuite/btcd/chaincfg"
	bch "github.com/gcash/bchd/chaincfg"
	"github.com/opusmagna/tatum-go/model/request"
	"github.com/opusmagna/tatum-go/network/ltc"
	"github.com/opusmagna/tatum-go/private_key"
	"github.com/opusmagna/tatum-go/utils"
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
		network = &ltc.LtcTestNet4Params
		path = utils.TestnetDerivationPath
	} else {
		network = &ltc.LtcMainNetParams
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
 * Generate VeChain wallet
 *
 * @param testnet testnet or mainnet version of address
 * @param mnem    mnemonic seed to use
 * @return the wallet
 * @throws ExecutionException   the execution exception
 * @throws InterruptedException the interrupted exception
 * @returns wallet
 */
func generateVetWallet(testnet bool, mnemonic string) *Wallet {
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

	xpub := private_key.NewBtcPrivateKey().Network(network).FromSeed(mnemonic).DerivePath(path).Xpub()
	return &Wallet{Mnemonic: mnemonic, Xpub: xpub}
}

/**
 * Generate Ethereum or any other ERC20 wallet
 * @param testnet testnet or mainnet version of address
 * @param mnem mnemonic seed to use
 * @returns wallet
 */
func generateEthWallet(testnet bool, mnemonic string) *Wallet {
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

	xpub := private_key.NewBtcPrivateKey().Network(network).FromSeed(mnemonic).DerivePath(path).Xpub()
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
	case request.VET:
		return generateVetWallet(testnet, mnemonic)
	case request.USDT:
	case request.WBTC:
	case request.LEO:
	case request.LINK:
	case request.UNI:
	case request.FREE:
	case request.MKR:
	case request.USDC:
	case request.BAT:
	case request.TUSD:
	case request.PAX:
	case request.PAXG:
	case request.PLTC:
	case request.XCON:
	case request.ETH:
	case request.MMY:
		return generateEthWallet(testnet, mnemonic)
	default:
		return nil
	}
	return nil
}
