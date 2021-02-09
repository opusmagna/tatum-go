package wallet

import (
	"github.com/stretchr/testify/assert"
	"github.com/tatumio/tatum-go/model/request"
	"testing"
)

func TestGenerateAddressFromXPub_BTC(t *testing.T) {
	xpub_test := "tpubDFjLw3ykn4aB7fFt96FaqRjSnvtDsU2wpVr8GQk3Eo612LS9jo9JgMkQRfYVG248J3pTBsxGg3PYUXFd7pReNLTeUzxFcUDL3zCvrp3H34a"
	address_test := GenerateAddressFromXPub(request.BTC, true, xpub_test, 1)
	assert.Equal(t, "mjJotvHmzEuyXZJGJXXknS6N3PWQnw6jf5", address_test, "they should be equal")

	xpub := "xpub6EsCk1uU6cJzqvP9CdsTiJwT2rF748YkPnhv5Qo8q44DG7nn2vbyt48YRsNSUYS44jFCW9gwvD9kLQu9AuqXpTpM1c5hgg9PsuBLdeNncid"
	address := GenerateAddressFromXPub(request.BTC, false, xpub, 1)
	assert.Equal(t, "1HWYaP13JKtaW2Mhq69NVeSLjRYGpD3aKv", address, "they should be equal")
}

func TestGenerateAddressFromXPub_LTC(t *testing.T) {
	xpub_test := "ttub4giastL5S3AicjXRBEJt7uq22b611rJvVfTgJSRfYeyZkwXwKnZcctK3tEjMpqrgiNSnYAzkKPJDxGoKNWQzkzTJxSryHbaYxsYW9Vr6AYQ"
	address_test := GenerateAddressFromXPub(request.LTC, true, xpub_test, 1)
	assert.Equal(t, "mjJotvHmzEuyXZJGJXXknS6N3PWQnw6jf5", address_test, "they should be equal")

	xpub := "Ltub2aXe9g8RPgAcY6jb6FftNJfQXHMV6UNBeZwrWH1K3vjpua9u8uj95xkZyCC4utdEbfYeh9TwxcUiFy2mGzBCJVBwW3ezHmLX2fHxv7HUt8J"
	address := GenerateAddressFromXPub(request.LTC, false, xpub, 1)
	assert.Equal(t, "LepMzqfXSgQommH2qu3fk7Gf5xgoHQsP1b", address, "they should be equal")
}

func TestGenerateAddressFromXPub_BCH(t *testing.T) {
	xpub_test := "tpubDExJFAGFe7NbFfXAtG1TRF19LDxq9JCFnHncz6mFjj2jabiNNVUiDUtpipbLSkNo74j2Rke82tkwzWEvDShudB7nT49mSimsF9gzFwTf4nw"
	address_test := GenerateAddressFromXPub(request.BCH, true, xpub_test, 1)
	assert.Equal(t, "bchtest:qr9wgjtyjd4q60323gd2ytsv5w3thl92rcms83akcc", address_test, "they should be equal")

	xpub := "xpub6EafivSZvqR8ysLKS52NDKfn16sB9uhCEfCKdYi7PpGqqK3fJGdd53DzUnWYvFRZKAC7pB8FVnvuJKkJparfjjfVPTQTmC7dfC6aVvw6f98"
	address := GenerateAddressFromXPub(request.BCH, false, xpub, 1)
	assert.Equal(t, "bitcoincash:qr9wgjtyjd4q60323gd2ytsv5w3thl92rclzrklply", address, "they should be equal")
}

func TestGeneratePrivateKeyFromMnemonic_BTC(t *testing.T) {
	mnemonic := "quantum tobacco key they maid mean crime youth chief jungle mind design broken tilt bus shoulder leaf good forward erupt split divert bread kitten"
	//privKey := GeneratePrivateKeyFromMnemonic(request.BTC, true, mnemonic, 1)
	//assert.Equal(t, "cQ1YZMep3CiAnMTA9y62ha6BjGaaTFsTvtDuGmucGvpAVmS89khV", privKey, "they should be equal")

	privKey := GeneratePrivateKeyFromMnemonic(request.BTC, false, mnemonic, 1)
	assert.Equal(t, "KwrYonf8pFfyQR87NTn124Ep9zoJsZMBCoVUi7mjMc1eTHDyLyBN", privKey, "they should be equal")
}

func TestGeneratePrivateKeyFromMnemonic_LTC(t *testing.T) {
	mnemonic := "quantum tobacco key they maid mean crime youth chief jungle mind design broken tilt bus shoulder leaf good forward erupt split divert bread kitten"
	privKey := GeneratePrivateKeyFromMnemonic(request.LTC, true, mnemonic, 1)
	assert.Equal(t, "cQ1YZMep3CiAnMTA9y62ha6BjGaaTFsTvtDuGmucGvpAVmS89khV", privKey, "they should be equal")

	privKey = GeneratePrivateKeyFromMnemonic(request.LTC, false, mnemonic, 1)
	assert.Equal(t, "T63MUovVt5GN5rmfwYMr4M6YqFmisjbrZrfZYZ53qWmCwiP6xCHa", privKey, "they should be equal")
}

func TestGeneratePrivateKeyFromMnemonic_BCH(t *testing.T) {
	mnemonic := "quantum tobacco key they maid mean crime youth chief jungle mind design broken tilt bus shoulder leaf good forward erupt split divert bread kitten"
	privKey := GeneratePrivateKeyFromMnemonic(request.BCH, true, mnemonic, 1)
	assert.Equal(t, "cRCLa2kAZ4XpSF62HaqbBEWKA2aVquTGX5sRmFuu2SpZ4s72vi5Y", privKey, "they should be equal")

	privKey = GeneratePrivateKeyFromMnemonic(request.BCH, false, mnemonic, 1)
	assert.Equal(t, "KzqM77kK7zqZGockuB2Tov1FXoH6BTMaT3ixeqTPXLAYp838W3KT", privKey, "they should be equal")
}
