package request

type WalletPrivateKeyRequest struct {
	Mnemonic string `json:"mnemonic"`
	Index    int    `json:"index"`
}
