package currency

import (
	"errors"
)

// https://www.xe.com/iso4217.php

type Currency string

const UnitedArabEmiratesDirham Currency = "AED"
const AfghanistanAfghani Currency = "AFN"
const AlbaniaLek Currency = "ALL"
const ArmeniaDram Currency = "AMD"
const NetherlandsAntillesGuilder Currency = "ANG"
const AngolaKwanza Currency = "AOA"
const ArgentinaPeso Currency = "ARS"
const AustraliaDollar Currency = "AUD"
const ArubaGuilder Currency = "AWG"
const AzerbaijanManat Currency = "AZN"
const BosniaandHerzegovinaConvertibleMark Currency = "BAM"
const BarbadosDollar Currency = "BBD"
const BangladeshTaka Currency = "BDT"
const BulgariaLev Currency = "BGN"
const BahrainDinar Currency = "BHD"
const BurundiFranc Currency = "BIF"
const BermudaDollar Currency = "BMD"
const BruneiDarussalamDollar Currency = "BND"
const BoliviaBolíviano Currency = "BOB"
const BrazilReal Currency = "BRL"
const BahamasDollar Currency = "BSD"
const BhutanNgultrum Currency = "BTN"
const BotswanaPula Currency = "BWP"
const BelarusRuble Currency = "BYN"
const BelizeDollar Currency = "BZD"
const CanadaDollar Currency = "CAD"
const CongoKinshasaFranc Currency = "CDF"
const SwitzerlandFranc Currency = "CHF"
const ChilePeso Currency = "CLP"
const ChinaYuanRenminbi Currency = "CNY"
const ColombiaPeso Currency = "COP"
const CostaRicaColon Currency = "CRC"
const CubaConvertiblePeso Currency = "CUC"
const CubaPeso Currency = "CUP"
const CapeVerdeEscudo Currency = "CVE"
const CzechRepublicKoruna Currency = "CZK"
const DjiboutiFranc Currency = "DJF"
const DenmarkKrone Currency = "DKK"
const DominicanRepublicPeso Currency = "DOP"
const AlgeriaDinar Currency = "DZD"
const EgyptPound Currency = "EGP"
const EritreaNakfa Currency = "ERN"
const EthiopiaBirr Currency = "ETB"
const EuroMemberCountries Currency = "EUR"
const FijiDollar Currency = "FJD"
const FalklandIslandsPound Currency = "FKP"
const UnitedKingdomPound Currency = "GBP"
const GeorgiaLari Currency = "GEL"
const GuernseyPound Currency = "GGP"
const GhanaCedi Currency = "GHS"
const GibraltarPound Currency = "GIP"
const GambiaDalasi Currency = "GMD"
const GuineaFranc Currency = "GNF"
const GuatemalaQuetzal Currency = "GTQ"
const GuyanaDollar Currency = "GYD"
const HongKongDollar Currency = "HKD"
const HondurasLempira Currency = "HNL"
const CroatiaKuna Currency = "HRK"
const HaitiGourde Currency = "HTG"
const HungaryForint Currency = "HUF"
const IndonesiaRupiah Currency = "IDR"
const IsraelShekel Currency = "ILS"
const IsleofManPound Currency = "IMP"
const IndiaRupee Currency = "INR"
const IraqDinar Currency = "IQD"
const IranRial Currency = "IRR"
const IcelandKrona Currency = "ISK"
const JerseyPound Currency = "JEP"
const JamaicaDollar Currency = "JMD"
const JordanDinar Currency = "JOD"
const JapanYen Currency = "JPY"
const KenyaShilling Currency = "KES"
const KyrgyzstanSom Currency = "KGS"
const CambodiaRiel Currency = "KHR"
const ComorianFranc Currency = "KMF"
const KoreaNorthWon Currency = "KPW"
const KoreaSouthWon Currency = "KRW"
const KuwaitDinar Currency = "KWD"
const CaymanIslandsDollar Currency = "KYD"
const KazakhstanTenge Currency = "KZT"
const LaosKip Currency = "LAK"
const LebanonPound Currency = "LBP"
const SriLankaRupee Currency = "LKR"
const LiberiaDollar Currency = "LRD"
const LesothoLoti Currency = "LSL"
const LibyaDinar Currency = "LYD"
const MoroccoDirham Currency = "MAD"
const MoldovaLeu Currency = "MDL"
const MadagascarAriary Currency = "MGA"
const MacedoniaDenar Currency = "MKD"
const MyanmarBurmaKyat Currency = "MMK"
const MongoliaTughrik Currency = "MNT"
const MacauPataca Currency = "MOP"
const MauritaniaOuguiya Currency = "MRU"
const MauritiusRupee Currency = "MUR"
const MaldivesRufiyaa Currency = "MVR"
const MalawiKwacha Currency = "MWK"
const MexicoPeso Currency = "MXN"
const MalaysiaRinggit Currency = "MYR"
const MozambiqueMetical Currency = "MZN"
const NamibiaDollar Currency = "NAD"
const NigeriaNaira Currency = "NGN"
const NicaraguaCordoba Currency = "NIO"
const NorwayKrone Currency = "NOK"
const NepalRupee Currency = "NPR"
const NewZealandDollar Currency = "NZD"
const OmanRial Currency = "OMR"
const PanamaBalboa Currency = "PAB"
const PeruSol Currency = "PEN"
const PapuaNewGuineaKina Currency = "PGK"
const PhilippinesPeso Currency = "PHP"
const PakistanRupee Currency = "PKR"
const PolandZloty Currency = "PLN"
const ParaguayGuarani Currency = "PYG"
const QatarRiyal Currency = "QAR"
const RomaniaLeu Currency = "RON"
const SerbiaDinar Currency = "RSD"
const RussiaRuble Currency = "RUB"
const RwandaFranc Currency = "RWF"
const SaudiArabiaRiyal Currency = "SAR"
const SolomonIslandsDollar Currency = "SBD"
const SeychellesRupee Currency = "SCR"
const SudanPound Currency = "SDG"
const SwedenKrona Currency = "SEK"
const SingaporeDollar Currency = "SGD"
const SaintHelenaPound Currency = "SHP"
const SierraLeoneLeone Currency = "SLL"
const SomaliaShilling Currency = "SOS"
const SeborgaLuigino Currency = "SPL"
const SurinameDollar Currency = "SRD"
const SãoToméandPríncipeDobra Currency = "STN"
const ElSalvadorColon Currency = "SVC"
const SyriaPound Currency = "SYP"
const eSwatiniLilangeni Currency = "SZL"
const ThailandBaht Currency = "THB"
const TajikistanSomoni Currency = "TJS"
const TurkmenistanManat Currency = "TMT"
const TunisiaDinar Currency = "TND"
const TongaPaanga Currency = "TOP"
const TurkeyLira Currency = "TRY"
const TrinidadandTobagoDollar Currency = "TTD"
const TuvaluDollar Currency = "TVD"
const TaiwanNewDollar Currency = "TWD"
const TanzaniaShilling Currency = "TZS"
const UkraineHryvnia Currency = "UAH"
const UgandaShilling Currency = "UGX"
const UnitedStatesDollar Currency = "USD"
const UruguayPeso Currency = "UYU"
const UzbekistanSom Currency = "UZS"
const VenezuelaBolívar Currency = "VEF"
const VietNamDong Currency = "VND"
const VanuatuVatu Currency = "VUV"
const SamoaTala Currency = "WST"
const CommunauteFinanciereAfricaineBEACCFAFrancBEAC Currency = "XAF"
const EastCaribbeanDollar Currency = "XCD"
const InternationalMonetaryFundSpecialDrawingRights Currency = "XDR"
const CommunauteFinanciereAfricaineBCEAOFranc Currency = "XOF"
const ComptoirsFrancaisduPacifique Currency = "XPF"
const YemenRial Currency = "YER"
const SouthAfricaRand Currency = "ZAR"
const ZambiaKwacha Currency = "ZMW"
const ZimbabweDollar Currency = "ZWD"

var validCodes = map[string]Currency{
	"AED": UnitedArabEmiratesDirham,
	"AFN": AfghanistanAfghani,
	"ALL": AlbaniaLek,
	"AMD": ArmeniaDram,
	"ANG": NetherlandsAntillesGuilder,
	"AOA": AngolaKwanza,
	"ARS": ArgentinaPeso,
	"AUD": AustraliaDollar,
	"AWG": ArubaGuilder,
	"AZN": AzerbaijanManat,
	"BAM": BosniaandHerzegovinaConvertibleMark,
	"BBD": BarbadosDollar,
	"BDT": BangladeshTaka,
	"BGN": BulgariaLev,
	"BHD": BahrainDinar,
	"BIF": BurundiFranc,
	"BMD": BermudaDollar,
	"BND": BruneiDarussalamDollar,
	"BOB": BoliviaBolíviano,
	"BRL": BrazilReal,
	"BSD": BahamasDollar,
	"BTN": BhutanNgultrum,
	"BWP": BotswanaPula,
	"BYN": BelarusRuble,
	"BZD": BelizeDollar,
	"CAD": CanadaDollar,
	"CDF": CongoKinshasaFranc,
	"CHF": SwitzerlandFranc,
	"CLP": ChilePeso,
	"CNY": ChinaYuanRenminbi,
	"COP": ColombiaPeso,
	"CRC": CostaRicaColon,
	"CUC": CubaConvertiblePeso,
	"CUP": CubaPeso,
	"CVE": CapeVerdeEscudo,
	"CZK": CzechRepublicKoruna,
	"DJF": DjiboutiFranc,
	"DKK": DenmarkKrone,
	"DOP": DominicanRepublicPeso,
	"DZD": AlgeriaDinar,
	"EGP": EgyptPound,
	"ERN": EritreaNakfa,
	"ETB": EthiopiaBirr,
	"EUR": EuroMemberCountries,
	"FJD": FijiDollar,
	"FKP": FalklandIslandsPound,
	"GBP": UnitedKingdomPound,
	"GEL": GeorgiaLari,
	"GGP": GuernseyPound,
	"GHS": GhanaCedi,
	"GIP": GibraltarPound,
	"GMD": GambiaDalasi,
	"GNF": GuineaFranc,
	"GTQ": GuatemalaQuetzal,
	"GYD": GuyanaDollar,
	"HKD": HongKongDollar,
	"HNL": HondurasLempira,
	"HRK": CroatiaKuna,
	"HTG": HaitiGourde,
	"HUF": HungaryForint,
	"IDR": IndonesiaRupiah,
	"ILS": IsraelShekel,
	"IMP": IsleofManPound,
	"INR": IndiaRupee,
	"IQD": IraqDinar,
	"IRR": IranRial,
	"ISK": IcelandKrona,
	"JEP": JerseyPound,
	"JMD": JamaicaDollar,
	"JOD": JordanDinar,
	"JPY": JapanYen,
	"KES": KenyaShilling,
	"KGS": KyrgyzstanSom,
	"KHR": CambodiaRiel,
	"KMF": ComorianFranc,
	"KPW": KoreaNorthWon,
	"KRW": KoreaSouthWon,
	"KWD": KuwaitDinar,
	"KYD": CaymanIslandsDollar,
	"KZT": KazakhstanTenge,
	"LAK": LaosKip,
	"LBP": LebanonPound,
	"LKR": SriLankaRupee,
	"LRD": LiberiaDollar,
	"LSL": LesothoLoti,
	"LYD": LibyaDinar,
	"MAD": MoroccoDirham,
	"MDL": MoldovaLeu,
	"MGA": MadagascarAriary,
	"MKD": MacedoniaDenar,
	"MMK": MyanmarBurmaKyat,
	"MNT": MongoliaTughrik,
	"MOP": MacauPataca,
	"MRU": MauritaniaOuguiya,
	"MUR": MauritiusRupee,
	"MVR": MaldivesRufiyaa,
	"MWK": MalawiKwacha,
	"MXN": MexicoPeso,
	"MYR": MalaysiaRinggit,
	"MZN": MozambiqueMetical,
	"NAD": NamibiaDollar,
	"NGN": NigeriaNaira,
	"NIO": NicaraguaCordoba,
	"NOK": NorwayKrone,
	"NPR": NepalRupee,
	"NZD": NewZealandDollar,
	"OMR": OmanRial,
	"PAB": PanamaBalboa,
	"PEN": PeruSol,
	"PGK": PapuaNewGuineaKina,
	"PHP": PhilippinesPeso,
	"PKR": PakistanRupee,
	"PLN": PolandZloty,
	"PYG": ParaguayGuarani,
	"QAR": QatarRiyal,
	"RON": RomaniaLeu,
	"RSD": SerbiaDinar,
	"RUB": RussiaRuble,
	"RWF": RwandaFranc,
	"SAR": SaudiArabiaRiyal,
	"SBD": SolomonIslandsDollar,
	"SCR": SeychellesRupee,
	"SDG": SudanPound,
	"SEK": SwedenKrona,
	"SGD": SingaporeDollar,
	"SHP": SaintHelenaPound,
	"SLL": SierraLeoneLeone,
	"SOS": SomaliaShilling,
	"SPL": SeborgaLuigino,
	"SRD": SurinameDollar,
	"STN": SãoToméandPríncipeDobra,
	"SVC": ElSalvadorColon,
	"SYP": SyriaPound,
	"SZL": eSwatiniLilangeni,
	"THB": ThailandBaht,
	"TJS": TajikistanSomoni,
	"TMT": TurkmenistanManat,
	"TND": TunisiaDinar,
	"TOP": TongaPaanga,
	"TRY": TurkeyLira,
	"TTD": TrinidadandTobagoDollar,
	"TVD": TuvaluDollar,
	"TWD": TaiwanNewDollar,
	"TZS": TanzaniaShilling,
	"UAH": UkraineHryvnia,
	"UGX": UgandaShilling,
	"USD": UnitedStatesDollar,
	"UYU": UruguayPeso,
	"UZS": UzbekistanSom,
	"VEF": VenezuelaBolívar,
	"VND": VietNamDong,
	"VUV": VanuatuVatu,
	"WST": SamoaTala,
	"XAF": CommunauteFinanciereAfricaineBEACCFAFrancBEAC,
	"XCD": EastCaribbeanDollar,
	"XDR": InternationalMonetaryFundSpecialDrawingRights,
	"XOF": CommunauteFinanciereAfricaineBCEAOFranc,
	"XPF": ComptoirsFrancaisduPacifique,
	"YER": YemenRial,
	"ZAR": SouthAfricaRand,
	"ZMW": ZambiaKwacha,
	"ZWD": ZimbabweDollar,
}

func GetByCode(code string) (Currency, error) {
	c, exists := validCodes[code]
	if !exists {
		return "", errors.New("Currency code not found")
	}
	return c, nil
}
