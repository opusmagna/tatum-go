package transaction

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/tatumio/tatum-go/model/request"
	"testing"
)

func TestBitcoinTx_PrepareBitcoinSignedTransaction(t *testing.T) {

	body := request.TransferBtcBasedBlockchain{}
	var utxo = request.FromUTXO{TxHash: "16688d2946c3e029ca91ce730109994c2bcafb859d580a6f7c820fb2bb5b6afc",
		Index: 1, PrivateKey: "91izeJtyQ1DNGkiRtMGRKBEKYQTX46Ug8mGtKWpX9mDKqArsLpH"}

	body.FromUTXO = make([]request.FromUTXO, 0)
	body.FromUTXO = append(body.FromUTXO, utxo)
	var to = request.To{Address: "mkYvnmm3KUBkvVqUAYsG6A6amt5Dva4jzX", Value: 0.00060000}

	body.To = make([]request.To, 0)
	body.To = append(body.To, to)

	/*utxo = request.FromUTXO{TxHash: "4e8378675bcf6a389c8cfe246094aafa44249e48ab88a40e6fda3bf0f44f916a",
		Index: 0, PrivateKey: "5HusYj2b2x4nroApgfvaSfKYZhRbKFH41bVyPooymbC6KfgSXdD"}
	body.FromUTXO = append(body.FromUTXO, utxo)

	to = request.To{Address: "1KKKK6N21XKo48zWKuQKXdvSsCf95ibHFa", Value: 0.00091234}
	body.To = append(body.To, to)*/

	btc := BitcoinTx{}

	txData, _ := btc.PrepareBitcoinSignedTransaction(true, body)
	fmt.Println(txData)

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
		"00000000", txData, "they should be equal")

}
