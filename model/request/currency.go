package request

type Currency string

const (
	BTC  Currency = "BTC"
	BCH           = "BCH"
	LTC           = "LTC"
	ETH           = "ETH"
	XRP           = "XRP"
	XLM           = "XLM"
	VET           = "VET"
	NEO           = "NEO"
	BNB           = "BNB"
	USDT          = "USDT"
	LEO           = "LEO"
	LINK          = "LINK"
	UNI           = "UNI"
	FREE          = "FREE"
	MKR           = "MKR"
	USDC          = "USDC"
	BAT           = "BAT"
	TUSD          = "TUSD"
	PAX           = "PAX"
	PLTC          = "PLTC"
	ADA           = "ADA"
	XCON          = "XCON"
	MMY           = "MMY"
	PAXG          = "PAXG"
)

var currencies = [...]string{"BTC",
	"BCH",
	"ETH",
	"XRP",
	"XLM",
	"VET",
	"NEO",
	"BNB",
	"USDT",
	"LEO",
	"LINK",
	"UNI",
	"FREE",
	"MKR",
	"USDC",
	"BAT",
	"TUSD",
	"PAX",
	"PLTC",
	"ADA",
	"XCON",
	"MMY",
	"PAXG"}

func (c Currency) String() string {
	s, ok := c.IsValid()
	if ok {
		return *s
	}
	return *s
}

func (c Currency) IsValid() (*string, bool) {
	x := string(c)
	for _, v := range currencies {
		if v == x {
			return &x, true
		}
	}
	return nil, false
}

//export const ETH_BASED_CURRENCIES = [
//    Currency.USDT.toString(),
//    Currency.LEO.toString(),
//    Currency.LINK.toString(),
//    Currency.UNI.toString(),
//    Currency.FREE.toString(),
//    Currency.MKR.toString(),
//    Currency.USDC.toString(),
//    Currency.BAT.toString(),
//    Currency.TUSD.toString(),
//    Currency.PAX.toString(),
//    Currency.PAXG.toString(),
//    Currency.PLTC.toString(),
//    Currency.XCON.toString(),
//    Currency.ETH.toString(),
//    Currency.MMY.toString(),
//];
