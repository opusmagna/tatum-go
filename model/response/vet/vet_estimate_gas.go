package vet

/**
 *
 * @export
 * @interface VetEstimateGas
 */
type VetEstimateGas struct {
	/**
	 * Sender
	 * @type {string}
	 * @memberof VetEstimateGas
	 */
	From string
	/**
	 * Recipient
	 * @type {string}
	 * @memberof VetEstimateGas
	 */
	To string
	/**
	 * Amount to send
	 * @type {string}
	 * @memberof VetEstimateGas
	 */
	Value string
	/**
	 * Data to send to Smart Contract
	 * @type {string}
	 * @memberof VetEstimateGas
	 */
	Data string
	/**
	 * Nonce
	 * @type {number}
	 * @memberof VetEstimateGas
	 */
	Nonce uint64
}
