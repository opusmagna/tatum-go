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

func TestBitcoinTx_PrepareBitcoinSignedTransaction_2(t *testing.T) {

	body := request.TransferBtcBasedBlockchain{}
	var utxo = request.FromUTXO{TxHash: "53faa103e8217e1520f5149a4e8c84aeb58e55bdab11164a95e69a8ca50f8fcc",
		Index: 0, PrivateKey: "cVX7YtgL5muLTPncHFhP95oitV1mqUUA5VeSn8HeCRJbPqipzobf"}

	body.FromUTXO = make([]request.FromUTXO, 0)
	body.FromUTXO = append(body.FromUTXO, utxo)
	var to = request.To{Address: "2MzNGwuKvMEvKMQogtgzSqJcH2UW3Tc5oc7", Value: 0.02969944}

	body.To = make([]request.To, 0)
	body.To = append(body.To, to)

	btc := BitcoinTx{}

	txData, _ := btc.PrepareBitcoinSignedTransaction(true, body)
	fmt.Println(txData)

	//assert.Equal(t, "0100000001cc8f0fa58c9ae6954a1611abbd558eb5ae848c4e9a14f520157e21e803a1fa53000000006a47304402205e49848369acc41719b669dcc9ba486c570f1ca4974f61a4321329fe35e3ff36022007485588ede47e17db992ba41aef35c72cb292f9889d471f2c592fb7f252672e012103b17a162956975765aa6951f6349f9ab5bf510584c5df9f6065924bfd94a08513ffffffff0158512d000000000017a9144e1e4321307c88ecd4ddd6aeec040c6f01e53c998700000000", txData, "they should be equal")

}
