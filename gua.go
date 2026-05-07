package yi

import (
	"errors"
	"time"
)

// ============================================================================
// 类型别名 (便于国际用户理解)
// ============================================================================

// Hexagram 是 Gua 的英文别名
type Hexagram = Gua

// Line 是 Yao 的英文别名
type Line = Yao

// Trigram 是 Bagua 的英文别名
type Trigram = Bagua

// IChing 是 ZhouYi 的英文别名
type IChing = ZhouYi

// ============================================================================
// 核心类型定义
// ============================================================================

// Gua 卦象 (六十四卦)
type Gua struct {
	Xu        int     // 卦序 (1-64)
	Index     string  // 索引 (如: "乾乾", "坤坤")
	Shang     string  // 上卦名称
	ShangNum  int     // 上卦数 (0-7)
	Xia       string  // 下卦名称
	XiaNum    int     // 下卦数 (0-7)
	JiXiong   string  // 吉凶
	Xiang     string  // 卦象符号 (如: ䷀)
	Ming      string  // 卦名 (如: 乾为天)
	Yi        string  // 卦意 (邵雍)
	Symbol    string  // 符号
	Yaos      [6]*Yao // 六爻: 初，二，三，四，五，上
	Yong        string // 用九/用六
	YongJiXiong string // 用九/用六吉凶
}

// Yao 爻 (卦中的一爻)
type Yao struct {
	Ci      string // 爻辞
	JiXiong string // 爻吉凶
	NvMing  string // 女命判断
}

// Bagua 八卦 (单卦)
// 使用 int 表示，二进制: 0=阳, 1=阴
// 0b000 = 乾(三阳), 0b111 = 坤(三阴)
type Bagua = int

// 八卦常量
const (
	Qian Bagua = 0b000 // 乾 ☰ 天
	Dui  Bagua = 0b001 // 兑 ☱ 泽
	Li   Bagua = 0b010 // 离 ☲ 火
	Zhen Bagua = 0b011 // 震 ☳ 雷
	Xun  Bagua = 0b100 // 巽 ☴ 风
	Kan  Bagua = 0b101 // 坎 ☵ 水
	Gen  Bagua = 0b110 // 艮 ☶ 山
	Kun  Bagua = 0b111 // 坤 ☷ 地
)

// 卦类型 (本卦、变卦、互卦、错卦、综卦)
const (
	Ben  = iota // 本卦 (Original)
	Bian        // 变卦 (Changed)
	Hu          // 互卦 (Mutual)
	Cuo         // 错卦 (Opposite)
	Zong        // 综卦 (Reversed)
	GuaTypeMax
)

// YaoPosition 爻位
type YaoPosition int

const (
	Chu   YaoPosition = iota // 初爻 (First)
	Er                       // 二爻 (Second)
	San                      // 三爻 (Third)
	Si                       // 四爻 (Fourth)
	Wu                       // 五爻 (Fifth)
	Shang                    // 上爻 (Top)
	YaoCount
)

// ZhouYi 周易主结构
type ZhouYi struct {
	gua     [GuaTypeMax]*Gua
	bianYao []int // 变爻数
}

// ============================================================================
// 错误定义
// ============================================================================

var (
	ErrInvalidGuaIndex = errors.New("invalid gua index")
	ErrInvalidYaoIndex = errors.New("invalid yao index")
	ErrInvalidDayanIdx = errors.New("invalid dayan index, must be positive")
	ErrGuaNotFound     = errors.New("gua not found")
)

// ============================================================================
// 核心 API
// ============================================================================

// Divine 起卦 (按数起卦)
// xia: 下卦数, shang: 上卦数
func Divine(xia, shang int) *ZhouYi {
	return DivineByNumber(shang, xia)
}

// DivineByNumber 按数起卦
// shang: 上卦数, xia: 下卦数, bianYao: 变爻数(可选)
func DivineByNumber(shang, xia int, bianYao ...int) *ZhouYi {
	ben := getBenGua(shang, xia)
	bian := getBianGua(ben, bianYao...)
	hu := ben
	if (ben.ShangNum == Kun && ben.XiaNum == Kun) ||
		(ben.ShangNum == Qian && ben.XiaNum == Qian) {
		hu = bian
	}
	hu = getHuGua(hu)
	cuo := getCuoGua(ben)
	zong := getZongGua(ben)

	return &ZhouYi{
		gua: [GuaTypeMax]*Gua{
			Ben:  ben,
			Bian: bian,
			Hu:   hu,
			Cuo:  cuo,
			Zong: zong,
		},
		bianYao: bianYao,
	}
}

// DivineByTime 按时间起卦
func DivineByTime(shang, xia int, t time.Time) *ZhouYi {
	by := timeToBianYao(t)
	return DivineByNumber(shang, xia, by)
}

// GetGua 获取指定类型的卦
func (zy *ZhouYi) GetGua(guaType int) *Gua {
	if guaType < 0 || guaType >= GuaTypeMax {
		return nil
	}
	return zy.gua[guaType]
}

// GetBianYao 获取变爻位置 (0-5)
func (zy *ZhouYi) GetBianYao() int {
	return calcBianYao(zy.bianYao...)
}

// GetYao 获取指定爻
func (g *Gua) GetYao(pos YaoPosition) *Yao {
	if pos < 0 || int(pos) >= int(YaoCount) {
		return nil
	}
	return g.Yaos[pos]
}

// ============================================================================
// 查询 API
// ============================================================================

// GetGuaByIndex 按索引获取卦象
// index: 如 "乾乾", "坤坤", "坎震" 等
func GetGuaByIndex(index string) (*Gua, error) {
	if g, ok := guaStore[index]; ok {
		return g, nil
	}
	return nil, ErrGuaNotFound
}

// GetGuaByXu 按卦序获取卦象 (1-64)
func GetGuaByXu(xu int) (*Gua, error) {
	if xu < 1 || xu > 64 {
		return nil, ErrInvalidGuaIndex
	}
	for _, g := range guaStore {
		if g.Xu == xu {
			return g, nil
		}
	}
	return nil, ErrGuaNotFound
}

// GetBaguaName 获取八卦名称
func GetBaguaName(bg Bagua) string {
	if bg < 0 || bg > 7 {
		return ""
	}
	return baguaNames[bg]
}

// GetBaguaSymbol 获取八卦符号
func GetBaguaSymbol(bg Bagua) string {
	if bg < 0 || bg > 7 {
		return ""
	}
	return baguaSymbols[bg]
}

// ============================================================================
// 判断 API
// ============================================================================

// IsJi 是否为吉
func (zy *ZhouYi) IsJi(sex Sex) bool {
	yao := zy.GetGua(Bian).Yaos[zy.GetBianYao()]
	if yao == nil {
		return false
	}

	if sex == Female && yao.NvMing != "" {
		return !contains(yao.NvMing, "凶")
	}
	return !contains(yao.JiXiong, "凶")
}

// FilterYao 过滤爻
func (zy *ZhouYi) FilterYao(sex Sex, filters ...string) bool {
	yao := zy.GetGua(Bian).Yaos[zy.GetBianYao()]
	if yao == nil {
		return true
	}

	for _, f := range filters {
		if sex == Female && yao.NvMing != "" {
			if yao.NvMing == f {
				return false
			}
			return true
		}
		if yao.JiXiong == f {
			return false
		}
	}
	return true
}

// ============================================================================
// 预生成数据 (由 go generate 生成)
// ============================================================================

// baguaNames 八卦名称
var baguaNames = [8]string{
	Qian: "乾", // ☰ 天
	Dui:  "兑", // ☱ 泽
	Li:   "离", // ☲ 火
	Zhen: "震", // ☳ 雷
	Xun:  "巽", // ☴ 风
	Kan:  "坎", // ☵ 水
	Gen:  "艮", // ☶ 山
	Kun:  "坤", // ☷ 地
}

// baguaSymbols 八卦符号
var baguaSymbols = [8]string{
	Qian: "☰",
	Dui:  "☱",
	Li:   "☲",
	Zhen: "☳",
	Xun:  "☴",
	Kan:  "☵",
	Gen:  "☶",
	Kun:  "☷",
}

// guaStore 卦象存储 (由 go generate 填充)
var guaStore = map[string]*Gua{}

// ============================================================================
// 内部函数
// ============================================================================

// getBenGua 本卦
func getBenGua(shang, xia int) *Gua {
	idx := GetBaguaName(shang%8) + GetBaguaName(xia%8)
	if g, ok := guaStore[idx]; ok {
		return g
	}
	return nil
}

// getBianGua 变卦
func getBianGua(ben *Gua, bianYao ...int) *Gua {
	bz := calcBianYao(bianYao...)
	shang := ben.ShangNum
	xia := ben.XiaNum

	if bz > 2 {
		shang = bianYaoTransform(shang, bz-3)
	} else {
		xia = bianYaoTransform(xia, bz)
	}

	idx := GetBaguaName(shang) + GetBaguaName(xia)
	if g, ok := guaStore[idx]; ok {
		return g
	}
	return nil
}

// bianYaoTransform 变爻变换
func bianYaoTransform(gua, pos int) int {
	mask := 1 << (2 - uint(pos))
	if gua&mask == 0 {
		return gua | mask
	}
	return gua ^ mask
}

// calcBianYao 计算变爻位置
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

// getHuGua 互卦
func getHuGua(ben *Gua) *Gua {
	shang := calcHuShang(ben.ShangNum, ben.XiaNum)
	xia := calcHuXia(ben.ShangNum, ben.XiaNum)
	idx := GetBaguaName(shang) + GetBaguaName(xia)
	if g, ok := guaStore[idx]; ok {
		return g
	}
	return nil
}

// calcHuShang 计算互卦上卦
func calcHuShang(shang, xia int) int {
	result := 0
	if xia&(1<<0) > 0 {
		result |= 1 << 2
	}
	if shang&(1<<2) > 0 {
		result |= 1 << 1
	}
	if shang&(1<<1) > 0 {
		result |= 1 << 0
	}
	return result
}

// calcHuXia 计算互卦下卦
func calcHuXia(shang, xia int) int {
	result := 0
	if xia&(1<<1) > 0 {
		result |= 1 << 2
	}
	if xia&(1<<0) > 0 {
		result |= 1 << 1
	}
	if shang&(1<<2) > 0 {
		result |= 1 << 0
	}
	return result
}

// getCuoGua 错卦 (阴阳全反)
func getCuoGua(ben *Gua) *Gua {
	shang := ^ben.ShangNum & 0x7
	xia := ^ben.XiaNum & 0x7
	idx := GetBaguaName(shang) + GetBaguaName(xia)
	if g, ok := guaStore[idx]; ok {
		return g
	}
	return nil
}

// getZongGua 综卦 (上下颠倒)
func getZongGua(ben *Gua) *Gua {
	shang := reverseBits(ben.XiaNum)
	xia := reverseBits(ben.ShangNum)
	idx := GetBaguaName(shang) + GetBaguaName(xia)
	if g, ok := guaStore[idx]; ok {
		return g
	}
	return nil
}

// reverseBits 反转3位二进制
func reverseBits(n int) int {
	return ((n & 0x4) >> 2) | (n & 0x2) | ((n & 0x1) << 2)
}

// timeToBianYao 时间转变爻
func timeToBianYao(t time.Time) int {
	const layout = "2006-01-02 15:04"
	parsed, _ := time.ParseInLocation(layout, t.Format(layout), t.Location())
	if n := parsed.Unix(); n != 0 {
		return int(n % 6)
	}
	return 0
}

// contains 字符串包含检查
func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > 0 && containsHelper(s, substr))
}

func containsHelper(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

// ============================================================================
// 初始化 (加载数据)
// ============================================================================

func init() {
	// 数据由 go generate 生成并填充
	// 这里保留空实现，实际数据在 data_generated.go 中
}
