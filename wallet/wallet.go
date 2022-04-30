package wallet

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
