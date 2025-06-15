package main

import (
	"database/sql/driver"
	"fmt"
	"math"
	"strings"
)

type Currency struct {
	s string
}

func (c *Currency) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	v, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("currency: assertion to string failed: %T", value)
	}
	if len(v) == 0 {
		return nil
	}
	c1, err := Parse(string(v))
	if err != nil {
		return fmt.Errorf("cannot scan currency: %w", err)
	}
	*c = c1
	return nil
}

func (c Currency) Value() (driver.Value, error) {
	return c.s, nil
}

func MustParse(cur string) Currency {
	c, err := Parse(cur)
	if err != nil {
		panic(err)
	}
	return c
}

func Parse(cur string) (Currency, error) {
	c := Currency{strings.ToLower(cur)}
	_, ok := currencies[c]
	if !ok {
		return XXX, fmt.Errorf("currency: unknown currency: %s", cur)
	}
	return c, nil
}

func (c Currency) Code() uint16 {
	i, ok := currencies[c]
	if !ok {
		return 999
	}
	return i.Code
}

func (c *Currency) UnmarshalText(b []byte) error {
	u, err := Parse(string(b))
	if err != nil {
		return err
	}
	*c = u
	return nil
}

func (c Currency) Equal(b Currency) bool {
	return c.s == b.s
}

func (c Currency) MarshalText() ([]byte, error) {
	return []byte(c.s), nil
}

func (c Currency) String() string {
	return c.s
}

type Money struct {
	amount   int64
	currency Currency
}

func NewMoney(a int64, cur Currency) Money {
	return Money{a, cur}
}

func NewMoneyFromFloat(f float64, cur Currency) Money {
	i, ok := currencies[cur]
	if !ok {
		return Money{}
	}
	scale := i.Scale
	amount := math.Round(f * math.Pow10(scale))
	a := int64(amount)
	return NewMoney(a, cur)
}

func (m Money) Amount() int64 {
	return m.amount
}

func (m Money) Currency() Currency {
	return m.currency
}

func (m Money) Value() (int64, Currency) {
	return m.amount, m.currency
}

func (m Money) ValueFloat() (float64, Currency) {
	i, ok := currencies[m.currency]
	if !ok {
		return 0.0, XXX
	}
	scale := i.Scale
	f := float64(m.amount) * math.Pow10(-scale)
	return f, m.currency
}

func (m Money) Equal(u Money) bool {
	return m.amount == u.amount && m.currency == u.currency
}

func (m Money) IsZero() bool {
	return m.amount == 0 && m.currency == (Currency{})
}

// List of values that Currency can take.
var (
	AED Currency = Currency{"aed"} // United Arab Emirates Dirham (784)
	AFN Currency = Currency{"afn"} // Afghan Afghani (971)
	ALL Currency = Currency{"all"} // Albanian Lek (8)
	AMD Currency = Currency{"amd"} // Armenian Dram (51)
	ANG Currency = Currency{"ang"} // Netherlands Antillean Gulden (532)
	AOA Currency = Currency{"aoa"} // Angolan Kwanza (973)
	ARS Currency = Currency{"ars"} // Argentine Peso (32)
	AUD Currency = Currency{"aud"} // Australian Dollar (36)
	AWG Currency = Currency{"awg"} // Aruban Florin (533)
	AZN Currency = Currency{"azn"} // Azerbaijani Manat (944)
	BAM Currency = Currency{"bam"} // Bosnia & Herzegovina Convertible Mark (977)
	BBD Currency = Currency{"bbd"} // Barbadian Dollar (52)
	BDT Currency = Currency{"bdt"} // Bangladeshi Taka (50)
	BGN Currency = Currency{"bgn"} // Bulgarian Lev (975)
	BHD Currency = Currency{"bhd"} // Bahraini Dinar (48)
	BIF Currency = Currency{"bif"} // Burundian Franc (108)
	BMD Currency = Currency{"bmd"} // Bermudian Dollar (60)
	BND Currency = Currency{"bnd"} // Brunei Dollar (96)
	BOB Currency = Currency{"bob"} // Bolivian Boliviano (68)
	BRL Currency = Currency{"brl"} // Brazilian Real (986)
	BSD Currency = Currency{"bsd"} // Bahamian Dollar (44)
	BTN Currency = Currency{"btn"} // Bhutanese Ngultrum (64)
	BWP Currency = Currency{"bwp"} // Botswana Pula (72)
	BYN Currency = Currency{"byn"} // Belarusian Ruble (933)
	BYR Currency = Currency{"byr"} // Belarusian Ruble (974)
	BZD Currency = Currency{"bzd"} // Belize Dollar (84)
	CAD Currency = Currency{"cad"} // Canadian Dollar (124)
	CDF Currency = Currency{"cdf"} // Congolese Franc (976)
	CHF Currency = Currency{"chf"} // Swiss Franc (756)
	CLF Currency = Currency{"clf"} // Chilean Unit of Account (UF) (990)
	CLP Currency = Currency{"clp"} // Chilean Peso (152)
	CNY Currency = Currency{"cny"} // Chinese Yuan (156)
	COP Currency = Currency{"cop"} // Colombian Peso (170)
	COU Currency = Currency{"cou"} // Unidad de Valor Real (970)
	CRC Currency = Currency{"crc"} // Costa Rican Colón (188)
	CUP Currency = Currency{"cup"} // Cuban Peso (192)
	CVE Currency = Currency{"cve"} // Cape Verdean Escudo (132)
	CZK Currency = Currency{"czk"} // Czech Koruna (203)
	DJF Currency = Currency{"djf"} // Djiboutian Franc (262)
	DKK Currency = Currency{"dkk"} // Danish Krone (208)
	DOP Currency = Currency{"dop"} // Dominican Peso (214)
	DZD Currency = Currency{"dzd"} // Algerian Dinar (12)
	EGP Currency = Currency{"egp"} // Egyptian Pound (818)
	ERN Currency = Currency{"ern"} // Eritrean Nakfa (232)
	ETB Currency = Currency{"etb"} // Ethiopian Birr (230)
	EUR Currency = Currency{"eur"} // Euro (978)
	FJD Currency = Currency{"fjd"} // Fijian Dollar (242)
	FKP Currency = Currency{"fkp"} // Falkland Islands Pound (238)
	GBP Currency = Currency{"gbp"} // British Pound Sterling (826)
	GEL Currency = Currency{"gel"} // Georgian Lari (981)
	GGP Currency = Currency{"ggp"} // Guernsey Pound (936)
	GHS Currency = Currency{"ghs"} // Ghanaian Cedi (936)
	GIP Currency = Currency{"gip"} // Gibraltar Pound (292)
	GMD Currency = Currency{"gmd"} // Gambian Dalasi (270)
	GNF Currency = Currency{"gnf"} // Guinean Franc (324)
	GTQ Currency = Currency{"gtq"} // Guatemalan Quetzal (320)
	GYD Currency = Currency{"gyd"} // Guyanaese Dollar (328)
	HKD Currency = Currency{"hkd"} // Hong Kong Dollar (344)
	HNL Currency = Currency{"hnl"} // Honduran Lempira (340)
	HRK Currency = Currency{"hrk"} // Croatian Kuna (191)
	HTG Currency = Currency{"htg"} // Haitian Gourde (332)
	HUF Currency = Currency{"huf"} // Hungarian Forint (348)
	IDR Currency = Currency{"idr"} // Indonesian Rupiah (360)
	ILS Currency = Currency{"ils"} // Israeli New Sheqel (376)
	IMP Currency = Currency{"imp"} // Isle of Man Pound (Nil)
	INR Currency = Currency{"inr"} // Indian Rupee (356)
	IQD Currency = Currency{"iqd"} // Iraqi Dinar (368)
	IRR Currency = Currency{"irr"} // Iranian Rial (364)
	ISK Currency = Currency{"isk"} // Icelandic Króna (352)
	JEP Currency = Currency{"jep"} // Jersey Pound (936)
	JMD Currency = Currency{"jmd"} // Jamaican Dollar (388)
	JOD Currency = Currency{"jod"} // Jordanian Dinar (400)
	JPY Currency = Currency{"jpy"} // Japanese Yen (392)
	KES Currency = Currency{"kes"} // Kenyan Shilling (404)
	KGS Currency = Currency{"kgs"} // Kyrgyzstani Som (417)
	KHR Currency = Currency{"khr"} // Cambodian Riel (116)
	KID Currency = Currency{"kid"} // Kiribati Dollar (Nil)
	KMF Currency = Currency{"kmf"} // Comorian Franc (174)
	KRW Currency = Currency{"krw"} // South Korean Won (410)
	KWD Currency = Currency{"kwd"} // Kuwaiti Dinar (414)
	KYD Currency = Currency{"kyd"} // Cayman Islands Dollar (136)
	KZT Currency = Currency{"kzt"} // Kazakhstani Tenge (398)
	LAK Currency = Currency{"lak"} // Laotian Kip (418)
	LBP Currency = Currency{"lbp"} // Lebanese Pound (422)
	LKR Currency = Currency{"lkr"} // Sri Lankan Rupee (144)
	LRD Currency = Currency{"lrd"} // Liberian Dollar (430)
	LSL Currency = Currency{"lsl"} // Basotho Loti (426)
	LYD Currency = Currency{"lyd"} // Libyan Dinar (434)
	MAD Currency = Currency{"mad"} // Moroccan Dirham (504)
	MDL Currency = Currency{"mdl"} // Moldovan Leu (498)
	LTL Currency = Currency{"ltl"} // Lithuanian Litas (440)
	LVL Currency = Currency{"lvl"} // Latvian Lats (428)
	MRU Currency = Currency{"mru"} // Mauritanian Ouguiya (929)
	STN Currency = Currency{"stn"} // São Tomé and Príncipe Dobra (930)
	MGA Currency = Currency{"mga"} // Malagasy Ariary (969)
	MKD Currency = Currency{"mkd"} // Macedonian Denar (807)
	MMK Currency = Currency{"mmk"} // Myanma Kyat (104)
	MNT Currency = Currency{"mnt"} // Mongolian Tugrik (496)
	MOP Currency = Currency{"mop"} // Macanese Pataca (446)
	MRO Currency = Currency{"mro"} // Mauritanian Ouguiya (478)
	MUR Currency = Currency{"mur"} // Mauritian Rupee (480)
	MVR Currency = Currency{"mvr"} // Maldivian Rufiyaa (462)
	MWK Currency = Currency{"mwk"} // Malawian Kwacha (454)
	MXN Currency = Currency{"mxn"} // Mexican Peso (484)
	MYR Currency = Currency{"myr"} // Malaysian Ringgit (458)
	MZN Currency = Currency{"mzn"} // Mozambican Metical (943)
	NAD Currency = Currency{"nad"} // Namibian Dollar (516)
	NGN Currency = Currency{"ngn"} // Nigerian Naira (566)
	NIO Currency = Currency{"nio"} // Nicaraguan Córdoba (558)
	NOK Currency = Currency{"nok"} // Norwegian Krone (578)
	NPR Currency = Currency{"npr"} // Nepalese Rupee (524)
	NZD Currency = Currency{"nzd"} // New Zealand Dollar (554)
	OMR Currency = Currency{"omr"} // Omani Rial (512)
	PAB Currency = Currency{"pab"} // Panamanian Balboa (590)
	PEN Currency = Currency{"pen"} // Peruvian Nuevo Sol (604)
	PGK Currency = Currency{"pgk"} // Papua New Guinean Kina (598)
	PHP Currency = Currency{"php"} // Philippine Peso (608)
	PKR Currency = Currency{"pkr"} // Pakistani Rupee (586)
	PLN Currency = Currency{"pln"} // Polish Złoty (985)
	PYG Currency = Currency{"pyg"} // Paraguayan Guarani (600)
	QAR Currency = Currency{"qar"} // Qatari Riyal (634)
	RON Currency = Currency{"ron"} // Romanian Leu (946)
	RSD Currency = Currency{"rsd"} // Serbian Dinar (941)
	RUB Currency = Currency{"rub"} // Russian Ruble (643)
	RWF Currency = Currency{"rwf"} // Rwandan Franc (646)
	SAR Currency = Currency{"sar"} // Saudi Riyal (682)
	SBD Currency = Currency{"sbd"} // Solomon Islands Dollar (90)
	SCR Currency = Currency{"scr"} // Seychellois Rupee (690)
	SDG Currency = Currency{"sdg"} // Sudanese Pound (938)
	SEK Currency = Currency{"sek"} // Swedish Krona (752)
	SGD Currency = Currency{"sgd"} // Singapore Dollar (702)
	SHP Currency = Currency{"shp"} // Saint Helena Pound (654)
	SLL Currency = Currency{"sll"} // Sierra Leonean Leone (694)
	SOS Currency = Currency{"sos"} // Somali Shilling (706)
	SRD Currency = Currency{"srd"} // Surinamese Dollar (968)
	SSP Currency = Currency{"ssp"} // South Sudanese Pound (728)
	STD Currency = Currency{"std"} // São Tomé and Príncipe Dobra (678)
	SVC Currency = Currency{"svc"} // Salvadoran Colón (222)
	SYP Currency = Currency{"syp"} // Syrian Pound (760)
	SZL Currency = Currency{"szl"} // Swazi Lilangeni (748)
	THB Currency = Currency{"thb"} // Thai Baht (764)
	TJS Currency = Currency{"tjs"} // Tajikistani Somoni (972)
	TMT Currency = Currency{"tmt"} // Turkmenistani Manat (934)
	TND Currency = Currency{"tnd"} // Tunisian Dinar (788)
	TOP Currency = Currency{"top"} // Tongan Paʻanga (776)
	TRY Currency = Currency{"try"} // Turkish Lira (949)
	TTD Currency = Currency{"ttd"} // Trinidad and Tobago Dollar (780)
	TWD Currency = Currency{"twd"} // New Taiwan Dollar (901)
	TZS Currency = Currency{"tzs"} // Tanzanian Shilling (834)
	UAH Currency = Currency{"uah"} // Ukrainian Hryvnia (980)
	UGX Currency = Currency{"ugx"} // Ugandan Shilling (800)
	USD Currency = Currency{"usd"} // United States Dollar (840)
	UYU Currency = Currency{"uyu"} // Uruguayan Peso (858)
	UZS Currency = Currency{"uzs"} // Uzbekistani Som (860)
	VEF Currency = Currency{"vef"} // Venezuelan Bolívar (937)
	VES Currency = Currency{"ves"} // Venezuelan Bolívar Soberano (928)
	VND Currency = Currency{"vnd"} // Vietnamese Dong (704)
	VUV Currency = Currency{"vuv"} // Vanuatu Vatu (548)
	WST Currency = Currency{"wst"} // Samoan Tala (882)
	XAF Currency = Currency{"xaf"} // Central African CFA Franc BEAC (950)
	XAG Currency = Currency{"xag"} // Silver Ounce (961)
	XAU Currency = Currency{"xau"} // Gold Ounce (959)
	XBA Currency = Currency{"xba"} // European Composite Unit (955)
	XBB Currency = Currency{"xbb"} // European Monetary Unit (956)
	XBC Currency = Currency{"xbc"} // European Unit of Account 9 (957)
	XBD Currency = Currency{"xbd"} // European Unit of Account 17 (958)
	XCD Currency = Currency{"xcd"} // East Caribbean Dollar (951)
	XDR Currency = Currency{"xdr"} // International Monetary Fund (IMF) Special Drawing Rights (960)
	XOF Currency = Currency{"xof"} // CFA Franc BCEAO (952)
	XPD Currency = Currency{"xpd"} // Palladium Ounce (964)
	XPF Currency = Currency{"xpf"} // CFP Franc (953)
	XPT Currency = Currency{"xpt"} // Platinum Ounce (962)
	XSU Currency = Currency{"xsu"} // SUCRE (994)
	XTS Currency = Currency{"xts"} // Code for the Testing of Payments (963)
	XUA Currency = Currency{"xua"} // ADB Unit of Account (965)
	XXX Currency = Currency{"xxx"} // No Currency (999)
	YER Currency = Currency{"yer"} // Yemeni Rial (886)
	ZAR Currency = Currency{"zar"} // South African Rand (710)
	ZMK Currency = Currency{"zmk"} // Zambian Kwacha (900)
	ZMW Currency = Currency{"zmw"} // Zambian Kwacha (967)
	ZWL Currency = Currency{"zwl"} // Zimbabwean Dollar (932)

	// Stable currencies
	USDT Currency = Currency{"usdt"} // USDT
	USDC Currency = Currency{"usdc"} // USDC
)

type currencyInfo struct {
	Code     uint16
	Scale    int
	IsStable bool
	IsCrypto bool
	Backing  Currency
}

var currencies = map[Currency]currencyInfo{
	AED: {Code: 784, Scale: 2}, // United Arab Emirates Dirham
	AFN: {Code: 971, Scale: 2}, // Afghan Afghani
	ALL: {Code: 8, Scale: 2},   // Albanian Lek
	AMD: {Code: 51, Scale: 2},  // Armenian Dram
	ANG: {Code: 532, Scale: 2}, // Netherlands Antillean Gulden
	AOA: {Code: 973, Scale: 2}, // Angolan Kwanza
	ARS: {Code: 32, Scale: 2},  // Argentine Peso
	AUD: {Code: 36, Scale: 2},  // Australian Dollar
	AWG: {Code: 533, Scale: 2}, // Aruban Florin
	AZN: {Code: 944, Scale: 2}, // Azerbaijani Manat
	BAM: {Code: 977, Scale: 2}, // Bosnia & Herzegovina Convertible Mark
	BBD: {Code: 52, Scale: 2},  // Barbadian Dollar
	BDT: {Code: 50, Scale: 2},  // Bangladeshi Taka
	BGN: {Code: 975, Scale: 2}, // Bulgarian Lev
	BHD: {Code: 48, Scale: 3},  // Bahraini Dinar
	BIF: {Code: 108, Scale: 0}, // Burundian Franc
	BMD: {Code: 60, Scale: 2},  // Bermudian Dollar
	BND: {Code: 96, Scale: 2},  // Brunei Dollar
	BOB: {Code: 68, Scale: 2},  // Bolivian Boliviano
	BRL: {Code: 986, Scale: 2}, // Brazilian Real
	BSD: {Code: 44, Scale: 2},  // Bahamian Dollar
	BTN: {Code: 64, Scale: 2},  // Bhutanese Ngultrum
	BWP: {Code: 72, Scale: 2},  // Botswana Pula
	BYN: {Code: 933, Scale: 2}, // Belarusian Ruble
	BYR: {Code: 974, Scale: 0}, // Belarusian Ruble (Obsolete)
	BZD: {Code: 84, Scale: 2},  // Belize Dollar
	CAD: {Code: 124, Scale: 2}, // Canadian Dollar
	CDF: {Code: 976, Scale: 2}, // Congolese Franc
	CHF: {Code: 756, Scale: 2}, // Swiss Franc
	CLP: {Code: 152, Scale: 0}, // Chilean Peso
	CNY: {Code: 156, Scale: 2}, // Chinese Yuan
	COP: {Code: 170, Scale: 2}, // Colombian Peso
	CRC: {Code: 188, Scale: 2}, // Costa Rican Colón
	CUP: {Code: 192, Scale: 2}, // Cuban Peso
	CVE: {Code: 132, Scale: 0}, // Cape Verdean Escudo
	CZK: {Code: 203, Scale: 2}, // Czech Koruna
	DJF: {Code: 262, Scale: 0}, // Djiboutian Franc
	DKK: {Code: 208, Scale: 2}, // Danish Krone
	DOP: {Code: 214, Scale: 2}, // Dominican Peso
	DZD: {Code: 12, Scale: 2},  // Algerian Dinar
	EGP: {Code: 818, Scale: 2}, // Egyptian Pound
	ERN: {Code: 232, Scale: 2}, // Eritrean Nakfa
	ETB: {Code: 230, Scale: 2}, // Ethiopian Birr
	EUR: {Code: 978, Scale: 2}, // Euro
	FJD: {Code: 242, Scale: 2}, // Fijian Dollar
	FKP: {Code: 238, Scale: 2}, // Falkland Islands Pound
	GBP: {Code: 826, Scale: 2}, // British Pound Sterling
	GEL: {Code: 981, Scale: 2}, // Georgian Lari
	GIP: {Code: 292, Scale: 2}, // Gibraltar Pound
	GMD: {Code: 270, Scale: 2}, // Gambian Dalasi
	GNF: {Code: 324, Scale: 0}, // Guinean Franc
	GTQ: {Code: 320, Scale: 2}, // Guatemalan Quetzal
	GYD: {Code: 328, Scale: 2}, // Guyanaese Dollar
	HKD: {Code: 344, Scale: 2}, // Hong Kong Dollar
	HNL: {Code: 340, Scale: 2}, // Honduran Lempira
	HRK: {Code: 191, Scale: 2}, // Croatian Kuna
	HTG: {Code: 332, Scale: 2}, // Haitian Gourde
	HUF: {Code: 348, Scale: 2}, // Hungarian Forint
	IDR: {Code: 360, Scale: 0}, // Indonesian Rupiah
	ILS: {Code: 376, Scale: 2}, // Israeli New Sheqel
	INR: {Code: 356, Scale: 2}, // Indian Rupee
	IQD: {Code: 368, Scale: 3}, // Iraqi Dinar
	IRR: {Code: 364, Scale: 2}, // Iranian Rial
	ISK: {Code: 352, Scale: 0}, // Icelandic Króna
	JMD: {Code: 388, Scale: 2}, // Jamaican Dollar
	JOD: {Code: 400, Scale: 3}, // Jordanian Dinar
	JPY: {Code: 392, Scale: 0}, // Japanese Yen
	KES: {Code: 404, Scale: 2}, // Kenyan Shilling
	KGS: {Code: 417, Scale: 2}, // Kyrgyzstani Som
	KHR: {Code: 116, Scale: 2}, // Cambodian Riel
	KID: {Code: 296, Scale: 2}, // Kiribati Dollar
	KMF: {Code: 174, Scale: 0}, // Comorian Franc
	KRW: {Code: 410, Scale: 0}, // South Korean Won
	KWD: {Code: 414, Scale: 3}, // Kuwaiti Dinar
	KYD: {Code: 136, Scale: 2}, // Cayman Islands Dollar
	KZT: {Code: 398, Scale: 2}, // Kazakhstani Tenge
	LAK: {Code: 418, Scale: 0}, // Laotian Kip
	LBP: {Code: 422, Scale: 2}, // Lebanese Pound
	LKR: {Code: 144, Scale: 2}, // Sri Lankan Rupee
	LRD: {Code: 430, Scale: 2}, // Liberian Dollar
	LSL: {Code: 426, Scale: 2}, // Lesotho Loti
	LTL: {Code: 440, Scale: 2}, // Lithuanian Litas
	LVL: {Code: 428, Scale: 2}, // Latvian Lats
	LYD: {Code: 434, Scale: 3}, // Libyan Dinar
	MAD: {Code: 504, Scale: 2}, // Moroccan Dirham
	MDL: {Code: 498, Scale: 2}, // Moldovan Leu
	MGA: {Code: 969, Scale: 0}, // Malagasy Ariary
	MKD: {Code: 807, Scale: 2}, // Macedonian Denar
	MMK: {Code: 104, Scale: 2}, // Myanma Kyat
	MNT: {Code: 496, Scale: 2}, // Mongolian Tugrik
	MOP: {Code: 446, Scale: 2}, // Macanese Pataca
	MRU: {Code: 929, Scale: 2}, // Mauritanian Ouguiya
	MUR: {Code: 480, Scale: 2}, // Mauritian Rupee
	MVR: {Code: 462, Scale: 2}, // Maldivian Rufiyaa
	MWK: {Code: 454, Scale: 2}, // Malawian Kwacha
	MXN: {Code: 484, Scale: 2}, // Mexican Peso
	MYR: {Code: 458, Scale: 2}, // Malaysian Ringgit
	MZN: {Code: 943, Scale: 2}, // Mozambican Metical
	NAD: {Code: 516, Scale: 2}, // Namibian Dollar
	NGN: {Code: 566, Scale: 2}, // Nigerian Naira
	NIO: {Code: 558, Scale: 2}, // Nicaraguan Córdoba
	NOK: {Code: 578, Scale: 2}, // Norwegian Krone
	NPR: {Code: 524, Scale: 2}, // Nepalese Rupee
	NZD: {Code: 554, Scale: 2}, // New Zealand Dollar
	OMR: {Code: 512, Scale: 3}, // Omani Rial
	PAB: {Code: 590, Scale: 2}, // Panamanian Balboa
	PEN: {Code: 604, Scale: 2}, // Peruvian Nuevo Sol
	PGK: {Code: 598, Scale: 2}, // Papua New Guinean Kina
	PHP: {Code: 608, Scale: 2}, // Philippine Peso
	PKR: {Code: 586, Scale: 2}, // Pakistani Rupee
	PLN: {Code: 985, Scale: 2}, // Polish Złoty
	PYG: {Code: 600, Scale: 0}, // Paraguayan Guarani
	QAR: {Code: 634, Scale: 2}, // Qatari Rial
	RON: {Code: 946, Scale: 2}, // Romanian Leu
	RSD: {Code: 941, Scale: 2}, // Serbian Dinar
	RUB: {Code: 643, Scale: 2}, // Russian Ruble
	RWF: {Code: 646, Scale: 0}, // Rwandan Franc
	SAR: {Code: 682, Scale: 2}, // Saudi Riyal
	SBD: {Code: 90, Scale: 2},  // Solomon Islands Dollar
	SCR: {Code: 690, Scale: 2}, // Seychellois Rupee
	SDG: {Code: 938, Scale: 2}, // Sudanese Pound
	SEK: {Code: 752, Scale: 2}, // Swedish Krona
	SGD: {Code: 702, Scale: 2}, // Singapore Dollar
	SHP: {Code: 654, Scale: 2}, // Saint Helena Pound
	SLL: {Code: 694, Scale: 2}, // Sierra Leonean Leone
	SOS: {Code: 706, Scale: 2}, // Somali Shilling
	SRD: {Code: 968, Scale: 2}, // Surinamese Dollar
	SSP: {Code: 728, Scale: 2}, // South Sudanese Pound
	STN: {Code: 930, Scale: 2}, // São Tomé and Príncipe Dobra
	SVC: {Code: 222, Scale: 2}, // Salvadoran Colón
	SYP: {Code: 760, Scale: 2}, // Syrian Pound
	SZL: {Code: 748, Scale: 2}, // Swazi Lilangeni
	THB: {Code: 764, Scale: 2}, // Thai Baht
	TJS: {Code: 972, Scale: 2}, // Tajikistani Somoni
	TMT: {Code: 934, Scale: 2}, // Turkmenistani Manat
	TND: {Code: 788, Scale: 3}, // Tunisian Dinar
	TOP: {Code: 776, Scale: 2}, // Tongan Paʻanga
	TRY: {Code: 949, Scale: 2}, // Turkish Lira
	TTD: {Code: 780, Scale: 2}, // Trinidad and Tobago Dollar
	TWD: {Code: 901, Scale: 2}, // New Taiwan Dollar
	TZS: {Code: 834, Scale: 2}, // Tanzanian Shilling
	UAH: {Code: 980, Scale: 2}, // Ukrainian Hryvnia
	UGX: {Code: 800, Scale: 0}, // Ugandan Shilling
	USD: {Code: 840, Scale: 2}, // United States Dollar
	UYU: {Code: 858, Scale: 2}, // Uruguayan Peso
	UZS: {Code: 860, Scale: 2}, // Uzbekistan Som
	VES: {Code: 928, Scale: 2}, // Venezuelan Bolívar
	VND: {Code: 704, Scale: 0}, // Vietnamese Dong
	VUV: {Code: 548, Scale: 0}, // Vanuatu Vatu
	WST: {Code: 882, Scale: 2}, // Samoan Tala
	XAF: {Code: 950, Scale: 0}, // Central African CFA Franc
	XCD: {Code: 951, Scale: 2}, // East Caribbean Dollar
	XOF: {Code: 952, Scale: 0}, // West African CFA Franc
	XPF: {Code: 953, Scale: 0}, // CFP Franc
	YER: {Code: 886, Scale: 2}, // Yemeni Rial
	ZAR: {Code: 710, Scale: 2}, // South African Rand
	ZMW: {Code: 967, Scale: 2}, // Zambian Kwacha
	ZWL: {Code: 932, Scale: 2}, // Zimbabwean Dollar

	// Stable crypto

	USDT: {Code: 1000, Scale: 2, IsStable: true, Backing: USD, IsCrypto: true}, // Tether
	USDC: {Code: 1001, Scale: 2, IsStable: true, Backing: USD, IsCrypto: true}, // USD Coin
}
