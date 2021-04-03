package utils

import "github.com/tatumio/tatum-go/model/request"

const TATUM_API_URL = "https://api-eu1.tatum.io"

const WhiteSpace = " "
const EmptySpace = ""
const SEPARATOR = "/"

const TestnetDerivationPath = "M/44H/1H/0H/0"
const BtcDerivationPath = "M/44H/0H/0H/0"
const LtcDerivationPath = "M/44H/2H/0H/0"
const BchDerivationPath = "M/44H/145H/0H/0"
const VetDerivationPath = "M/44H/818H/0H/0"
const EthDerivationPath = "M/44H/60H/0H/0"
const TronDerivationPath = "M/44H/195H/0H/0"

/**
 * The constant TRANSFER_METHOD_ABI.
 */
const TRANSFER_METHOD_ABI = "{" +
	"   \"constant\":false," +
	"   \"inputs\":[" +
	"      {" +
	"         \"name\":\"to\"," +
	"         \"type\":\"address\"" +
	"      }," +
	"      {" +
	"         \"name\":\"value\"," +
	"         \"type\":\"uint256\"" +
	"      }" +
	"   ]," +
	"   \"name\":\"transfer\"," +
	"   \"outputs\":[" +
	"      {" +
	"         \"name\":\"\"," +
	"         \"type\":\"bool\"" +
	"      }" +
	"   ]," +
	"   \"payable\":false," +
	"   \"stateMutability\":\"nonpayable\"," +
	"   \"type\":\"function\"" +
	"}"

func ContractAddresses() func(string) string {
	innerMap := map[string]string{
		request.BTC.String(): "0xdac17f958d2ee523a2206206994597c13d831ec7",
		request.LEO:          "0x2af5d2ad76741191d15dfe7bf6ac92d4bd912ca3",
		request.UNI:          "0x1f9840a85d5af5bf1d1762f925bdaddc4201f984",
		request.LINK:         "0x514910771af9ca656af840dff83e8264ecf986ca",
		request.FREE:         "0x2f141ce366a2462f02cea3d12cf93e4dca49e4fd",
		request.MKR:          "0x9f8f72aa9304c8b593d555f12ef6589cc3a579a2",
		request.USDC:         "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
		request.BAT:          "0x0d8775f648430679a709e98d2b0cb6250d2887ef",
		request.TUSD:         "0x0000000000085d4780B73119b644AE5ecd22b376",
		request.PAX:          "0x8e870d67f660d95d5be530380d0ec0bd388289e1",
		request.PAXG:         "0x45804880de22913dafe09f4980848ece6ecbaf78",
		request.PLTC:         "0x429d83bb0dcb8cdd5311e34680adc8b12070a07f",
		request.MMY:          "0x385ddf50c3de724f6b8ecb41745c29f9dd3c6d75",
		request.XCON:         "0x0f237d5ea7876e0e2906034d98fdb20d43666ad4",
	}

	return func(key string) string {
		return innerMap[key]
	}
}

func ContractDecimals() func(string) int {
	innerMap := map[string]int{
		request.USDT: 6,
		//request.USDT_TRON: 6,
		//request.WBTC: 8,
		request.LEO:  18,
		request.LINK: 18,
		request.UNI:  18,
		request.FREE: 18,
		request.MKR:  18,
		request.USDC: 6,
		request.BAT:  18,
		request.TUSD: 18,
		request.PAX:  18,
		request.PAXG: 18,
		request.PLTC: 18,
		request.MMY:  18,
		request.XCON: 18,
	}

	return func(key string) int {
		return innerMap[key]
	}
}
