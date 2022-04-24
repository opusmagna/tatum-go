// Copyright (c) 2014-2016 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package ltc

import (
	"math/big"
	"time"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/wire"
)

// These variables are the chain proof-of-work limit parameters for each default
// network.
var (

	// mainPowLimit is the highest proof of work value a Litecoin block can
	// have for the main network.
	mainPowLimit, _ = new(big.Int).SetString("0x0fffff000000000000000000000000000000000000000000000000000000", 0)

	// testNet4PowLimit is the highest proof of work value a Litecoin block
	// can have for the test network (version 4).
	testNet4PowLimit, _ = new(big.Int).SetString("0x0fffff000000000000000000000000000000000000000000000000000000", 0)
)

// ConsensusDeployment defines details related to a specific consensus rule
// change that is voted in.  This is part of BIP0009.
type ConsensusDeployment struct {
	// BitNumber defines the specific bit number within the block version
	// this particular soft-fork deployment refers to.
	BitNumber uint8

	// StartTime is the median block time after which voting on the
	// deployment starts.
	StartTime uint64

	// ExpireTime is the median block time after which the attempted
	// deployment expires.
	ExpireTime uint64
}

// Constants that define the deployment offset in the deployments field of the
// parameters for each deployment.  This is useful to be able to get the details
// of a specific deployment by name.
const (
	// DeploymentTestDummy defines the rule change deployment ID for testing
	// purposes.
	DeploymentTestDummy = iota

	// DeploymentCSV defines the rule change deployment ID for the CSV
	// soft-fork package. The CSV package includes the deployment of BIPS
	// 68, 112, and 113.
	DeploymentCSV

	// DeploymentSegwit defines the rule change deployment ID for the
	// Segregated Witness (segwit) soft-fork package. The segwit package
	// includes the deployment of BIPS 141, 142, 144, 145, 147 and 173.
	DeploymentSegwit

	// NOTE: DefinedDeployments must always come last since it is used to
	// determine how many defined deployments there currently are.

	// DefinedDeployments is the number of currently defined deployments.
	DefinedDeployments
)

// genesisBlock defines the genesis block of the block chain which serves as the
// public transaction ledger for the main network.
var genesisBlock = wire.MsgBlock{
	Header: wire.BlockHeader{
		Version:    1,
		PrevBlock:  chainhash.Hash{},  // 0000000000000000000000000000000000000000000000000000000000000000
		MerkleRoot: genesisMerkleRoot, // 97ddfbbae6be97fd6cdf3e7ca13232a3afff2353e29badfab7f73011edd4ced9
		Timestamp:  time.Unix(1317972665, 0),
		Bits:       0x1e0ffff0,
		Nonce:      2084524493,
	},
	Transactions: []*wire.MsgTx{&genesisCoinbaseTx},
}

// genesisHash is the hash of the first block in the block chain for the main
// network (genesis block).
var genesisHash = chainhash.Hash([chainhash.HashSize]byte{ // Make go vet happy.
	0xe2, 0xbf, 0x04, 0x7e, 0x7e, 0x5a, 0x19, 0x1a,
	0xa4, 0xef, 0x34, 0xd3, 0x14, 0x97, 0x9d, 0xc9,
	0x98, 0x6e, 0x0f, 0x19, 0x25, 0x1e, 0xda, 0xba,
	0x59, 0x40, 0xfd, 0x1f, 0xe3, 0x65, 0xa7, 0x12,
})

// genesisMerkleRoot is the hash of the first transaction in the genesis block
// for the main network.
var genesisMerkleRoot = chainhash.Hash([chainhash.HashSize]byte{ // Make go vet happy.
	0xd9, 0xce, 0xd4, 0xed, 0x11, 0x30, 0xf7, 0xb7,
	0xfa, 0xad, 0x9b, 0xe2, 0x53, 0x23, 0xff, 0xaf,
	0xa3, 0x32, 0x32, 0xa1, 0x7c, 0x3e, 0xdf, 0x6c,
	0xfd, 0x97, 0xbe, 0xe6, 0xba, 0xfb, 0xdd, 0x97,
})

// testNet4GenesisHash is the hash of the first block in the block chain for the
// test network (version 4).
var testNet4GenesisHash = chainhash.Hash([chainhash.HashSize]byte{ // Make go vet happy.
	0xa0, 0x29, 0x3e, 0x4e, 0xeb, 0x3d, 0xa6, 0xe6,
	0xf5, 0x6f, 0x81, 0xed, 0x59, 0x5f, 0x57, 0x88,
	0xd, 0x1a, 0x21, 0x56, 0x9e, 0x13, 0xee, 0xfd,
	0xd9, 0x51, 0x28, 0x4b, 0x5a, 0x62, 0x66, 0x49,
})

// testNet4GenesisBlock defines the genesis block of the block chain which
// serves as the public transaction ledger for the test network (version 4).
var testNet4GenesisBlock = wire.MsgBlock{
	Header: wire.BlockHeader{
		Version:    1,
		PrevBlock:  chainhash.Hash{},          // 0000000000000000000000000000000000000000000000000000000000000000
		MerkleRoot: testNet4GenesisMerkleRoot, // 97ddfbbae6be97fd6cdf3e7ca13232a3afff2353e29badfab7f73011edd4ced9
		Timestamp:  time.Unix(1486949366, 0),
		Bits:       0x1e0ffff0,
		Nonce:      293345,
	},
	Transactions: []*wire.MsgTx{&genesisCoinbaseTx},
}

// genesisCoinbaseTx is the coinbase transaction for the genesis blocks for
// the main network, regression test network, and test network (version 3).
var genesisCoinbaseTx = wire.MsgTx{
	Version: 1,
	TxIn: []*wire.TxIn{
		{
			PreviousOutPoint: wire.OutPoint{
				Hash:  chainhash.Hash{},
				Index: 0xffffffff,
			},
			SignatureScript: []byte{
				0x04, 0xff, 0xff, 0x00, 0x1d, 0x01, 0x04, 0x40, 0x4e, 0x59, 0x20, 0x54, 0x69, 0x6d, 0x65, 0x73, // |.......@NY Times|
				0x20, 0x30, 0x35, 0x2f, 0x4f, 0x63, 0x74, 0x2f, 0x32, 0x30, 0x31, 0x31, 0x20, 0x53, 0x74, 0x65, // | 05/Oct/2011 Ste|
				0x76, 0x65, 0x20, 0x4a, 0x6f, 0x62, 0x73, 0x2c, 0x20, 0x41, 0x70, 0x70, 0x6c, 0x65, 0xe2, 0x80, // |ve Jobs, Apple..|
				0x99, 0x73, 0x20, 0x56, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x61, 0x72, 0x79, 0x2c, 0x20, 0x44, 0x69, // |.s Visionary, Di|
				0x65, 0x73, 0x20, 0x61, 0x74, 0x20, 0x35, 0x36, // |es at 56|

			},
			Sequence: 0xffffffff,
		},
	},
	TxOut: []*wire.TxOut{
		{
			Value: 0x12a05f200,
			PkScript: []byte{
				0x41, 0x4, 0x1, 0x84, 0x71, 0xf, 0xa6, 0x89,
				0xad, 0x50, 0x23, 0x69, 0xc, 0x80, 0xf3, 0xa4,
				0x9c, 0x8f, 0x13, 0xf8, 0xd4, 0x5b, 0x8c, 0x85,
				0x7f, 0xbc, 0xbc, 0x8b, 0xc4, 0xa8, 0xe4, 0xd3,
				0xeb, 0x4b, 0x10, 0xf4, 0xd4, 0x60, 0x4f, 0xa0,
				0x8d, 0xce, 0x60, 0x1a, 0xaf, 0xf, 0x47, 0x2,
				0x16, 0xfe, 0x1b, 0x51, 0x85, 0xb, 0x4a, 0xcf,
				0x21, 0xb1, 0x79, 0xc4, 0x50, 0x70, 0xac, 0x7b,
				0x3, 0xa9, 0xac,
			},
		},
	},
	LockTime: 0,
}

// testNet4GenesisMerkleRoot is the hash of the first transaction in the genesis
// block for the test network (version 4).  It is the same as the merkle root
// for the main network.
var testNet4GenesisMerkleRoot = chainhash.Hash([chainhash.HashSize]byte{ // Make go vet happy.
	0xd9, 0xce, 0xd4, 0xed, 0x11, 0x30, 0xf7, 0xb7,
	0xfa, 0xad, 0x9b, 0xe2, 0x53, 0x23, 0xff, 0xaf,
	0xa3, 0x32, 0x32, 0xa1, 0x7c, 0x3e, 0xdf, 0x6c,
	0xfd, 0x97, 0xbe, 0xe6, 0xba, 0xfb, 0xdd, 0x97,
})
var (
	// LtcMainNetParams defines the network parameters for the main Litecoin network.
	LtcMainNetParams = chaincfg.Params{
		Name:        "mainnet",
		Net:         LtcMainNet,
		DefaultPort: "9333",
		DNSSeeds: []chaincfg.DNSSeed{
			{Host: "seed-a.litecoin.loshan.co.uk", HasFiltering: true},
			{Host: "dnsseed.thrasher.io", HasFiltering: true},
			{Host: "dnsseed.litecointools.com", HasFiltering: false},
			{Host: "dnsseed.litecoinpool.org", HasFiltering: false},
			{Host: "dnsseed.koin-project.com", HasFiltering: false},
		},

		// Chain parameters
		GenesisBlock:             &genesisBlock,
		GenesisHash:              &genesisHash,
		PowLimit:                 mainPowLimit,
		PowLimitBits:             504365055,
		BIP0034Height:            710000,
		BIP0065Height:            918684,
		BIP0066Height:            811879,
		CoinbaseMaturity:         100,
		SubsidyReductionInterval: 840000,
		TargetTimespan:           (time.Hour * 24 * 3) + (time.Hour * 12), // 3.5 days
		TargetTimePerBlock:       (time.Minute * 2) + (time.Second * 30),  // 2.5 minutes
		RetargetAdjustmentFactor: 4,                                       // 25% less, 400% more
		ReduceMinDifficulty:      false,
		MinDiffReductionTime:     0,
		GenerateSupported:        false,

		// Checkpoints ordered from oldest to newest.
		Checkpoints: []chaincfg.Checkpoint{
			{Height: 1500, Hash: newHashFromStr("841a2965955dd288cfa707a755d05a54e45f8bd476835ec9af4402a2b59a2967")},
			{Height: 4032, Hash: newHashFromStr("9ce90e427198fc0ef05e5905ce3503725b80e26afd35a987965fd7e3d9cf0846")},
			{Height: 8064, Hash: newHashFromStr("eb984353fc5190f210651f150c40b8a4bab9eeeff0b729fcb3987da694430d70")},
			{Height: 16128, Hash: newHashFromStr("602edf1859b7f9a6af809f1d9b0e6cb66fdc1d4d9dcd7a4bec03e12a1ccd153d")},
			{Height: 23420, Hash: newHashFromStr("d80fdf9ca81afd0bd2b2a90ac3a9fe547da58f2530ec874e978fce0b5101b507")},
			{Height: 50000, Hash: newHashFromStr("69dc37eb029b68f075a5012dcc0419c127672adb4f3a32882b2b3e71d07a20a6")},
			{Height: 80000, Hash: newHashFromStr("4fcb7c02f676a300503f49c764a89955a8f920b46a8cbecb4867182ecdb2e90a")},
			{Height: 120000, Hash: newHashFromStr("bd9d26924f05f6daa7f0155f32828ec89e8e29cee9e7121b026a7a3552ac6131")},
			{Height: 161500, Hash: newHashFromStr("dbe89880474f4bb4f75c227c77ba1cdc024991123b28b8418dbbf7798471ff43")},
			{Height: 179620, Hash: newHashFromStr("2ad9c65c990ac00426d18e446e0fd7be2ffa69e9a7dcb28358a50b2b78b9f709")},
			{Height: 240000, Hash: newHashFromStr("7140d1c4b4c2157ca217ee7636f24c9c73db39c4590c4e6eab2e3ea1555088aa")},
			{Height: 383640, Hash: newHashFromStr("2b6809f094a9215bafc65eb3f110a35127a34be94b7d0590a096c3f126c6f364")},
			{Height: 409004, Hash: newHashFromStr("487518d663d9f1fa08611d9395ad74d982b667fbdc0e77e9cf39b4f1355908a3")},
			{Height: 456000, Hash: newHashFromStr("bf34f71cc6366cd487930d06be22f897e34ca6a40501ac7d401be32456372004")},
			{Height: 638902, Hash: newHashFromStr("15238656e8ec63d28de29a8c75fcf3a5819afc953dcd9cc45cecc53baec74f38")},
			{Height: 721000, Hash: newHashFromStr("198a7b4de1df9478e2463bd99d75b714eab235a2e63e741641dc8a759a9840e5")},
		},

		// Consensus rule change deployments.
		//
		// The miner confirmation window is defined as:
		//   target proof of work timespan / target proof of work spacing
		RuleChangeActivationThreshold: 6048, // 75% of MinerConfirmationWindow
		MinerConfirmationWindow:       8064, //
		Deployments: [4]chaincfg.ConsensusDeployment{
			DeploymentTestDummy: {
				BitNumber:  28,
				StartTime:  1199145601, // January 1, 2008 UTC
				ExpireTime: 1230767999, // December 31, 2008 UTC
			},
			DeploymentCSV: {
				BitNumber:  0,
				StartTime:  1485561600, // January 28, 2017 UTC
				ExpireTime: 1517356801, // January 31st, 2018 UTC
			},
			DeploymentSegwit: {
				BitNumber:  1,
				StartTime:  1485561600, // January 28, 2017 UTC
				ExpireTime: 1517356801, // January 31st, 2018 UTC.
			},
		},

		// Mempool parameters
		RelayNonStdTxs: false,

		// Human-readable part for Bech32 encoded segwit addresses, as defined in
		// BIP 173.
		Bech32HRPSegwit: "ltc", // always ltc for main net

		// Address encoding magics
		PubKeyHashAddrID:        0x30, // starts with L
		ScriptHashAddrID:        0x32, // starts with M
		PrivateKeyID:            0xB0, // starts with 6 (uncompressed) or T (compressed)
		WitnessPubKeyHashAddrID: 0x06, // starts with p2
		WitnessScriptHashAddrID: 0x0A, // starts with 7Xh

		// BIP32 hierarchical deterministic extended key magics
		HDPrivateKeyID: [4]byte{0x01, 0x9d, 0x9c, 0xfe}, // starts with Ltpv
		HDPublicKeyID:  [4]byte{0x01, 0x9d, 0xa4, 0x62}, // starts with Ltub

		// BIP44 coin type used in the hierarchical deterministic path for
		// address generation.
		HDCoinType: 2,
	}
)

const (
	// MainNet represents the main litecoin network.
	LtcMainNet wire.BitcoinNet = 0xdbb6c0fb

	// TestNet4 represents the test network (version 4).
	LtcTestNet4 wire.BitcoinNet = 0xf1c8d2fd
)

// LtcTestNet4Params defines the network parameters for the test Litecoin network
// (version 4).  Not to be confused with the regression test network, this
// network is sometimes simply called "testnet".
var LtcTestNet4Params = chaincfg.Params{
	Name:        "testnet4",
	Net:         LtcTestNet4,
	DefaultPort: "19335",
	DNSSeeds: []chaincfg.DNSSeed{
		{Host: "testnet-seed.litecointools.com", HasFiltering: false},
		{Host: "seed-b.litecoin.loshan.co.uk", HasFiltering: true},
		{Host: "dnsseed-testnet.thrasher.io", HasFiltering: true},
	},

	// Chain parameters
	GenesisBlock:             &testNet4GenesisBlock,
	GenesisHash:              &testNet4GenesisHash,
	PowLimit:                 testNet4PowLimit,
	PowLimitBits:             504365055,
	BIP0034Height:            76,
	BIP0065Height:            76,
	BIP0066Height:            76,
	CoinbaseMaturity:         100,
	SubsidyReductionInterval: 840000,
	TargetTimespan:           (time.Hour * 24 * 3) + (time.Hour * 12), // 3.5 days
	TargetTimePerBlock:       (time.Minute * 2) + (time.Second * 30),  // 2.5 minutes
	RetargetAdjustmentFactor: 4,                                       // 25% less, 400% more
	ReduceMinDifficulty:      true,
	MinDiffReductionTime:     time.Minute * 5, // TargetTimePerBlock * 2
	GenerateSupported:        false,

	// Checkpoints ordered from oldest to newest.
	Checkpoints: []chaincfg.Checkpoint{
		{Height: 26115, Hash: newHashFromStr("817d5b509e91ab5e439652eee2f59271bbc7ba85021d720cdb6da6565b43c14f")},
		{Height: 43928, Hash: newHashFromStr("7d86614c153f5ef6ad878483118ae523e248cd0dd0345330cb148e812493cbb4")},
		{Height: 69296, Hash: newHashFromStr("66c2f58da3cfd282093b55eb09c1f5287d7a18801a8ff441830e67e8771010df")},
		{Height: 99949, Hash: newHashFromStr("8dd471cb5aecf5ead91e7e4b1e932c79a0763060f8d93671b6801d115bfc6cde")},
		{Height: 159256, Hash: newHashFromStr("ab5b0b9968842f5414804591119d6db829af606864b1959a25d6f5c114afb2b7")},
	},

	// Consensus rule change deployments.
	//
	// The miner confirmation window is defined as:
	//   target proof of work timespan / target proof of work spacing
	RuleChangeActivationThreshold: 1512, // 75% of MinerConfirmationWindow
	MinerConfirmationWindow:       2016,
	Deployments: [4]chaincfg.ConsensusDeployment{
		DeploymentTestDummy: {
			BitNumber:  28,
			StartTime:  1199145601, // January 1, 2008 UTC
			ExpireTime: 1230767999, // December 31, 2008 UTC
		},
		DeploymentCSV: {
			BitNumber:  0,
			StartTime:  1483228800, // January 1, 2017
			ExpireTime: 1517356801, // January 31st, 2018
		},
		DeploymentSegwit: {
			BitNumber:  1,
			StartTime:  1483228800, // January 1, 2017
			ExpireTime: 1517356801, // January 31st, 2018
		},
	},

	// Mempool parameters
	RelayNonStdTxs: true,

	// Human-readable part for Bech32 encoded segwit addresses, as defined in
	// BIP 173.
	Bech32HRPSegwit: "tltc", // always tltc for test net

	// Address encoding magics
	PubKeyHashAddrID:        0x6f, // starts with m or n
	ScriptHashAddrID:        0x3a, // starts with Q
	WitnessPubKeyHashAddrID: 0x52, // starts with QW
	WitnessScriptHashAddrID: 0x31, // starts with T7n
	PrivateKeyID:            0xef, // starts with 9 (uncompressed) or c (compressed)

	// BIP32 hierarchical deterministic extended key magics
	HDPrivateKeyID: [4]byte{0x04, 0x36, 0xef, 0x7d}, // starts with ttpv
	HDPublicKeyID:  [4]byte{0x04, 0x36, 0xf6, 0xe1}, // starts with ttub

	// BIP44 coin type used in the hierarchical deterministic path for
	// address generation.
	HDCoinType: 1,
}

// mustRegister performs the same function as Register except it panics if there
// is an error.  This should only be called from package init functions.
func mustRegister(params *chaincfg.Params) {
	if err := chaincfg.Register(params); err != nil {
		panic("failed to register network: " + err.Error())
	}
}

// newHashFromStr converts the passed big-endian hex string into a
// chainhash.Hash.  It only differs from the one available in chainhash in that
// it panics on an error since it will only (and must only) be called with
// hard-coded, and therefore known good, hashes.
func newHashFromStr(hexStr string) *chainhash.Hash {
	hash, err := chainhash.NewHashFromStr(hexStr)
	if err != nil {
		// Ordinarily I don't like panics in library code since it
		// can take applications down without them having a chance to
		// recover which is extremely annoying, however an exception is
		// being made in this case because the only way this can panic
		// is if there is an error in the hard-coded hashes.  Thus it
		// will only ever potentially panic on init and therefore is
		// 100% predictable.
		panic(err)
	}
	return hash
}

func init() {
	// Register all default networks when the package is initialized.
	mustRegister(&LtcMainNetParams)
	mustRegister(&LtcTestNet4Params)
}
