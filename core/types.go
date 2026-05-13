package core

type Hexagram = Gua

type Line = Yao

type Trigram = Bagua

type IChing = ZhouYi

type Gua struct {
	Xu          int
	Index       string
	ShangMing   string
	ShangNum    int
	XiaMing     string
	XiaNum      int
	JiXiong     string
	GuaName     string
	Ming        string
	GuaYi       string
	GuaSymbol   string
	TuanText    string
	XiangText   string
	Yaos        [6]*Yao
	Yong        string
	YongJiXiong string
}

type Yao struct {
	Ci      string
	JiXiong string
	NvMing  string
}

func (y *Yao) HasNvMing() bool {
	return y.NvMing != ""
}

type Bagua = int

type YaoPosition int

type ZhouYi struct {
	Gua     [GuaTypeMax]*Gua
	BianYao []int
}

func (zy *ZhouYi) GetGua(guaType int) *Gua {
	if guaType < 0 || guaType >= GuaTypeMax {
		return nil
	}
	return zy.Gua[guaType]
}

func (zy *ZhouYi) GetBianYao() int {
	return calcBianYao(zy.BianYao...)
}

func (zy *ZhouYi) GetAllBianYao() []int {
	result := make([]int, len(zy.BianYao))
	copy(result, zy.BianYao)
	return result
}

func (g *Gua) GetYao(pos YaoPosition) *Yao {
	if pos < 0 || int(pos) >= int(YaoCount) {
		return nil
	}
	return g.Yaos[pos]
}

func calcBianYao(bianYao ...int) int {
	if len(bianYao) == 0 {
		return 0
	}
	n := bianYao[0] % 6
	if n == 0 {
		return 5
	}
	return n - 1
}

type Sex int

func (s Sex) String() string {
	switch s {
	case Male:
		return "男"
	case Female:
		return "女"
	default:
		return "未知"
	}
}

func (s Sex) IsMale() bool {
	return s == Male
}

func (s Sex) IsFemale() bool {
	return s == Female
}

type YinYang int

func (yy YinYang) String() string {
	if yy == Yang {
		return "阳"
	}
	return "阴"
}

func (yy YinYang) IsYang() bool {
	return yy == Yang
}

func (yy YinYang) IsYin() bool {
	return yy == Yin
}

type WuXing int

func (wx WuXing) String() string {
	switch wx {
	case Wood:
		return "木"
	case Fire:
		return "火"
	case Earth:
		return "土"
	case Metal:
		return "金"
	case Water:
		return "水"
	default:
		return ""
	}
}

func (wx WuXing) Sheng() WuXing {
	switch wx {
	case Wood:
		return Fire
	case Fire:
		return Earth
	case Earth:
		return Metal
	case Metal:
		return Water
	case Water:
		return Wood
	default:
		return 0
	}
}

func (wx WuXing) Ke() WuXing {
	switch wx {
	case Wood:
		return Earth
	case Earth:
		return Water
	case Water:
		return Fire
	case Fire:
		return Metal
	case Metal:
		return Wood
	default:
		return 0
	}
}

func (wx WuXing) BeiSheng() WuXing {
	switch wx {
	case Wood:
		return Water
	case Fire:
		return Wood
	case Earth:
		return Fire
	case Metal:
		return Earth
	case Water:
		return Metal
	default:
		return 0
	}
}

func (wx WuXing) BeiKe() WuXing {
	switch wx {
	case Wood:
		return Metal
	case Fire:
		return Water
	case Earth:
		return Wood
	case Metal:
		return Fire
	case Water:
		return Earth
	default:
		return 0
	}
}

type GuaPosition int

func (gp GuaPosition) String() string {
	switch gp {
	case BenGong:
		return "本宫"
	case YiShi:
		return "一世"
	case ErShi:
		return "二世"
	case SanShi:
		return "三世"
	case SiShi:
		return "四世"
	case WuShi:
		return "五世"
	case YouHun:
		return "游魂"
	case GuiHun:
		return "归魂"
	default:
		return ""
	}
}

type ShiYingInfo struct {
	ShiPos   YaoPosition
	YingPos  YaoPosition
	Position GuaPosition
}

type shiYingEntry struct {
	ShiPos  YaoPosition
	YingPos YaoPosition
}

type LiuQin int

func (lq LiuQin) String() string {
	switch lq {
	case LQFuMu:
		return "父母"
	case LQXiongDi:
		return "兄弟"
	case LQQiCai:
		return "妻财"
	case LQZiSun:
		return "子孙"
	case LQGuanGui:
		return "官鬼"
	default:
		return ""
	}
}

func (lq LiuQin) Description() string {
	switch lq {
	case LQFuMu:
		return "生我者，主庇护文书"
	case LQXiongDi:
		return "同我者，主竞争劫财"
	case LQQiCai:
		return "我克者，主财富妻室"
	case LQZiSun:
		return "我生者，主子嗣福德"
	case LQGuanGui:
		return "克我者，主官非祸患"
	default:
		return ""
	}
}

type CoinResult int

func (c CoinResult) String() string {
	switch c {
	case CoinYinYinYin:
		return "老阴"
	case CoinYinYinYang:
		return "少阴"
	case CoinYinYangYang:
		return "少阳"
	case CoinYangYangYang:
		return "老阳"
	default:
		return "未知"
	}
}

func (c CoinResult) IsChanging() bool {
	return c == CoinYinYinYin || c == CoinYangYangYang
}

func (c CoinResult) IsYang() bool {
	return c == CoinYinYangYang || c == CoinYangYangYang
}

func (c CoinResult) YaoValue() int {
	return int(c)
}

type DayanStep struct {
	Total     int
	Left      int
	Right     int
	Remainder int
	Final     int
}

type DayanResult struct {
	Steps      [3]DayanStep
	Remaining  int
	YaoValue   int
	IsChanging bool
}

type TimeGuaParams struct {
	Year  int
	Month int
	Day   int
	Hour  int
}

type WenYanData struct {
	Index   string
	Entries []WenYanEntry
}

type WenYanEntry struct {
	Title string
	Text  string
}

type TianGan int

type DiZhi int

type JiaZiInfo struct {
	Index   int
	Name    string
	TianGan string
	DiZhi   string
	NaYin   string
	YinYang string
}

func (j *JiaZiInfo) WuXing() string {
	return j.NaYin
}

func (j *JiaZiInfo) GetNaYinWuXing() WuXing {
	return naYinToWuXing(j.NaYin)
}

func naYinToWuXing(naYin string) WuXing {
	if naYin == "" {
		return 0
	}
	runes := []rune(naYin)
	last := runes[len(runes)-1]
	switch last {
	case '金':
		return Metal
	case '木':
		return Wood
	case '水':
		return Water
	case '火':
		return Fire
	case '土':
		return Earth
	default:
		return 0
	}
}

type Dayan struct {
	Number     int
	JiXiong    string
	NvMing     string
	IsMax      bool
	Gua        string
	TianJiu    string
	YiXiang    string
	Foundation string
	Family     string
	Health     string
	Meaning    string
}

func (d *Dayan) IsJi() bool {
	return d.JiXiong == "吉" || d.JiXiong == "半吉"
}

func (d *Dayan) IsXiong() bool {
	return d.JiXiong == "凶"
}

func (d *Dayan) IsSuitableForFemale() bool {
	return d.NvMing != "凶"
}

func (d *Dayan) IsBest() bool {
	return d.IsMax
}

type FenXiCategory string

type GuaFenXi struct {
	Category FenXiCategory
	Content  string
	JiXiong  string
	Source   string
}

type JieGuaResult struct {
	ZhouYi *ZhouYi
	Sex    Sex

	BenGuaInfo  *GuaInfo
	BianGuaInfo *GuaInfo
	HuGuaInfo   *GuaInfo
	CuoGuaInfo  *GuaInfo
	ZongGuaInfo *GuaInfo

	DongYaoPos     YaoPosition
	DongYaoText    string
	DongYaoJiXiong string
	IsJi           bool
	JiXiongReason  string

	FenXi []GuaFenXi

	JieDu *GuaJieDu

	WuXingInfo *WuXingInfo
}

type GuaInfo struct {
	Gua       *Gua
	Ming      string
	GuaName   string
	GuaYi     string
	TuanText  string
	XiangText string
	JiXiong   string
	Symbol    string
	GuaGong   string
	Position  string
	ShiYao    string
	YingYao   string
}

type JieDuCategory string

type GuaJieDu struct {
	Ming      string
	ShiYe     string
	AiQing    string
	CaiYun    string
	KaoShi    string
	JianKang  string
	ChuXing   string
	GuanSi    string
	JiaZhai   string
	CoreImage string
	Yi        []string
	Ji        []string
}

type WuXingInfo struct {
	WuXing      string
	Direction   string
	LuckyNumber string
	LuckyColor  string
}
