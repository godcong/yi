package core

const (
	Ben = iota
	Bian
	Hu
	Cuo
	Zong
	GuaTypeMax
)

const (
	Chu   YaoPosition = iota
	Er
	San
	Si
	Wu
	Shang
	YaoCount
)

const (
	Qian Bagua = 0b000
	Dui  Bagua = 0b001
	Li   Bagua = 0b010
	Zhen Bagua = 0b011
	Xun  Bagua = 0b100
	Kan  Bagua = 0b101
	Gen  Bagua = 0b110
	Kun  Bagua = 0b111
)

const (
	Male   Sex = 0b01
	Female Sex = 0b10
)

const (
	Yang YinYang = iota
	Yin
)

const (
	Wood  WuXing = iota + 1
	Fire
	Earth
	Metal
	Water
)

const wuXingList = "木木火火土土金金水水"

const (
	BenGong GuaPosition = iota
	YiShi
	ErShi
	SanShi
	SiShi
	WuShi
	YouHun
	GuiHun
	GuaPositionCount
)

const (
	LQFuMu    LiuQin = iota
	LQXiongDi
	LQQiCai
	LQZiSun
	LQGuanGui
	LQCount
)

const (
	CoinYinYinYin    CoinResult = 6
	CoinYinYinYang   CoinResult = 7
	CoinYinYangYang  CoinResult = 8
	CoinYangYangYang CoinResult = 9
)

const (
	TGJia  TianGan = iota
	TGYi
	TGBing
	TGDing
	TGWu
	TGJi
	TGGeng
	TGXin
	TGRen
	TGGui
	TianGanCount
)

const (
	DZZi   DiZhi = iota
	DZChou
	DZYin
	DZMao
	DZChen
	DZSi
	DZWu
	DZWei
	DZShen
	DZYou
	DZXu
	DZHai
	DiZhiCount
)

const (
	FenXiShiYe    FenXiCategory = "事业"
	FenXiAiQing   FenXiCategory = "爱情"
	FenXiCaiYun   FenXiCategory = "财运"
	FenXiKaoShi   FenXiCategory = "考试"
	FenXiJianKang FenXiCategory = "健康"
	FenXiChuXing  FenXiCategory = "出行"
	FenXiGuanSi   FenXiCategory = "官司"
	FenXiJiaZhai  FenXiCategory = "家宅"
)

const (
	JieDuShiYe    JieDuCategory = "事业"
	JieDuAiQing   JieDuCategory = "爱情"
	JieDuCaiYun   JieDuCategory = "财运"
	JieDuKaoShi   JieDuCategory = "考试"
	JieDuJianKang JieDuCategory = "健康"
	JieDuChuXing  JieDuCategory = "出行"
	JieDuGuanSi   JieDuCategory = "官司"
	JieDuJiaZhai  JieDuCategory = "家宅"
)
