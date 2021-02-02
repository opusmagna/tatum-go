package bitcoin_tx_builder

import (
	"fmt"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTransaction_ToHex_MainNet(t *testing.T) {
	builder := New()
	builder = builder.Init(&chaincfg.MainNetParams).AddOutput("1KKKK6N21XKo48zWKuQKXdvSsCf95ibHFa", 91234)
	builder = builder.AddInput("4e8378675bcf6a389c8cfe246094aafa44249e48ab88a40e6fda3bf0f44f916a", 0, "5HusYj2b2x4nroApgfvaSfKYZhRbKFH41bVyPooymbC6KfgSXdD")
	tx := builder.Sign()

	if tx == nil {
		fmt.Errorf("error")
	}

	//fmt.Println(tx.ToHex())
	assert.Equal(t, "01000000016a914ff4f03bda6f0ea488ab489e2444faaa946024fe8c9c386acf5b6778834e000000008b483045022100904dbeddeecccf6391ac92922381ae006bf244c002f42e195daa0a9837a4b5820220461677f9dbb7d9580e268ac486cfeb4b9d87bfdd6d4e7b1be09b8e6f5cc0a70701410414e301b2328f17442c0b8310d787bf3d8a404cfbd0704f135b6ad4b2d3ee751310f981926e53a6e8c39bd7d3fefd576c543cce493cbac06388f2651d1aacbfcdffffffff0162640100000000001976a914c8e90996c7c6080ee06284600c684ed904d14c5c88ac00000000", tx.ToHex(), "they should be equal")

}

func TestTransaction_ToHex_TestNet(t *testing.T) {
	builder := New()

	builder = builder.Init(&chaincfg.TestNet3Params).AddOutput("mkYvnmm3KUBkvVqUAYsG6A6amt5Dva4jzX", 60000)
	builder = builder.AddInput("16688d2946c3e029ca91ce730109994c2bcafb859d580a6f7c820fb2bb5b6afc",
		1,
		"91izeJtyQ1DNGkiRtMGRKBEKYQTX46Ug8mGtKWpX9mDKqArsLpH")
	tx := builder.Sign()

	if tx == nil {
		fmt.Errorf("error")
	}

	//fmt.Println(tx.ToHex())
	assert.Equal(t, "01000000"+ // version
		"01"+ // Number of inputs
		"fc6a5bbbb20f827c6f0a589d85fbca2b4c99090173ce91ca29e0c346298d6816"+ // Outpoint TXID
		"01000000"+ // Outpoint index number
		"8a"+
		"47"+
		"304402207772d91e633259fb0cbd35427c4d24806877c5f52e8a8d032505ce3f8b73aa2302201731382284469f8bad48ad2457e1cf445b23c158922f26fb3e3c4fad6298cb5d0141044739edd9fc850cf5db037ecd839ba09f699765d0b13fe8c949688ed3b7ef9291a038729e0c70d6802e3adf1458550922012ebd1e9a979775578eefa867557506"+
		"ffffffff"+
		"01"+ // Number of outputs
		"60ea000000000000"+ // Satoshis
		"19"+ // Bytes in pubkey script: 25
		"76"+ // OP_DUP
		"a9"+ // OP_HASH160
		"14"+ // Push 20 bytes as data
		"37383464376cccaf2ae4c8a121805f45bf544e44"+ // PubKey hash
		"88"+ // OP_EQUALVERIFY
		"ac"+ // OP_CHECKSIG
		"00000000", // locktime
		tx.ToHex(), "they should be equal")

}
