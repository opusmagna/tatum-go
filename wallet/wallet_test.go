package wallet

import (
	"testing"

	"github.com/opusmagna/tatum-go/model/request"
	"github.com/stretchr/testify/assert"
)

func TestGenerateWallet_BTC(t *testing.T) {
	mnemonic := "quantum tobacco key they maid mean crime youth chief jungle mind design broken tilt bus shoulder leaf good forward erupt split divert bread kitten"

	testnet := GenerateWallet(request.BTC, true, mnemonic)
	assert.Equal(t, "tpubDFjLw3ykn4aB7fFt96FaqRjSnvtDsU2wpVr8GQk3Eo612LS9jo9JgMkQRfYVG248J3pTBsxGg3PYUXFd7pReNLTeUzxFcUDL3zCvrp3H34a", testnet.Xpub, "they should be equal")
	assert.Equal(t, mnemonic, testnet.Mnemonic, "they should be equal")

	mainnet := GenerateWallet(request.BTC, false, mnemonic)
	assert.Equal(t, "xpub6DtevPxud8AJUCqddtVqpqxAJvjFrYYvtt4co2kZF7ifPW3d5FJ3B9i5x7xL4CZirb2eNDEVqCtBgiGZR6Kkee5RdHsDoJVbtxi946axjXJ", mainnet.Xpub, "they should be equal")
	assert.Equal(t, mnemonic, mainnet.Mnemonic, "they should be equal")
}

func TestGenerateWallet_LTC(t *testing.T) {
	mnemonic := "quantum tobacco key they maid mean crime youth chief jungle mind design broken tilt bus shoulder leaf good forward erupt split divert bread kitten"

	testnet := GenerateWallet(request.LTC, true, mnemonic)
	assert.Equal(t, "ttub4giastL5S3AicjXRBEJt7uq22b611rJvVfTgJSRfYeyZkwXwKnZcctK3tEjMpqrgiNSnYAzkKPJDxGoKNWQzkzTJxSryHbaYxsYW9Vr6AYQ", testnet.Xpub, "they should be equal")
	assert.Equal(t, mnemonic, testnet.Mnemonic, "they should be equal")

	mainnet := GenerateWallet(request.LTC, false, mnemonic)
	assert.Equal(t, "Ltub2aXe9g8RPgAcY6jb6FftNJfQXHMV6UNBeZwrWH1K3vjpua9u8uj95xkZyCC4utdEbfYeh9TwxcUiFy2mGzBCJVBwW3ezHmLX2fHxv7HUt8J", mainnet.Xpub, "they should be equal")
	assert.Equal(t, mnemonic, mainnet.Mnemonic, "they should be equal")
}

func TestGenerateWallet_BCH(t *testing.T) {
	mnemonic := "quantum tobacco key they maid mean crime youth chief jungle mind design broken tilt bus shoulder leaf good forward erupt split divert bread kitten"

	testnet := GenerateWallet(request.BCH, true, mnemonic)
	assert.Equal(t, "tpubDExJFAGFe7NbFfXAtG1TRF19LDxq9JCFnHncz6mFjj2jabiNNVUiDUtpipbLSkNo74j2Rke82tkwzWEvDShudB7nT49mSimsF9gzFwTf4nw", testnet.Xpub, "they should be equal")
	assert.Equal(t, mnemonic, testnet.Mnemonic, "they should be equal")

	mainnet := GenerateWallet(request.BCH, false, mnemonic)
	assert.Equal(t, "xpub6EafivSZvqR8ysLKS52NDKfn16sB9uhCEfCKdYi7PpGqqK3fJGdd53DzUnWYvFRZKAC7pB8FVnvuJKkJparfjjfVPTQTmC7dfC6aVvw6f98", mainnet.Xpub, "they should be equal")
	assert.Equal(t, mnemonic, mainnet.Mnemonic, "they should be equal")
}

func TestGenerateWallet_VET(t *testing.T) {
	mnemonic := "quantum tobacco key they maid mean crime youth chief jungle mind design broken tilt bus shoulder leaf good forward erupt split divert bread kitten"

	testnet := GenerateWallet(request.VET, true, mnemonic)
	assert.Equal(t, "xpub6FMiQpA54nciqs52guGVdWQ5TonZt5XtGsFpurgtttL7H3mSfaJDXv5aBdThjX6tW9HYaJSQ8wZVnLm1ixaQUu1MRQCwvwZ6U2cX6mwWT25", testnet.Xpub, "they should be equal")
	assert.Equal(t, mnemonic, testnet.Mnemonic, "they should be equal")

	mainnet := GenerateWallet(request.VET, false, mnemonic)
	assert.Equal(t, "xpub6EzJLu3Hi5hEFAkiZAxCTaXqXoS95seTnG1tdYdF8fBcVZCfR8GQP8UGvfF52szpwZqiiGHJw5694emxSpYBE5qDxAZUgiHLzbVhb5ErRMa", mainnet.Xpub, "they should be equal")
	assert.Equal(t, mnemonic, mainnet.Mnemonic, "they should be equal")
}

func TestGenerateWallet_ETH(t *testing.T) {
	mnemonic := "quantum tobacco key they maid mean crime youth chief jungle mind design broken tilt bus shoulder leaf good forward erupt split divert bread kitten"

	testnet := GenerateWallet(request.ETH, true, mnemonic)
	assert.Equal(t, "xpub6FMiQpA54nciqs52guGVdWQ5TonZt5XtGsFpurgtttL7H3mSfaJDXv5aBdThjX6tW9HYaJSQ8wZVnLm1ixaQUu1MRQCwvwZ6U2cX6mwWT25", testnet.Xpub, "they should be equal")
	assert.Equal(t, mnemonic, testnet.Mnemonic, "they should be equal")

	mainnet := GenerateWallet(request.ETH, false, mnemonic)
	assert.Equal(t, "xpub6DtR524VQx3ENj2E9pNZnjqkVp47YN5sRCP5y4Gs6KZTwDhH9HTVX8shJPt74WaPZRftRXFfnsyPbMPh6DMEmrQ2WBxDJzGxriStAB36bQM", mainnet.Xpub, "they should be equal")
	assert.Equal(t, mnemonic, mainnet.Mnemonic, "they should be equal")
}
