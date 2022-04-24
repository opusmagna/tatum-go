package private_key

import (
	"crypto/sha512"
	"strings"

	"github.com/opusmagna/tatum-go/utils"
	"golang.org/x/crypto/pbkdf2"
)

/**
 * Convert mnemonic word list to seed.
 */
func ToSeed(words []string, passphrase string) []byte {
	pass := strings.Join(words, utils.WhiteSpace)
	salt := "mnemonic" + passphrase
	seed := pbkdf2.Key([]byte(pass), []byte(salt), 2048, 64, sha512.New)
	return seed
}
