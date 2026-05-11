package yi

import (
	"errors"
	"time"
)

// ============================================================================
// Type aliases (for international users)
// ============================================================================

// Hexagram is an English alias for Gua
type Hexagram = Gua

// Line is an English alias for Yao
type Line = Yao

// Trigram is an English alias for Bagua
type Trigram = Bagua

// IChing is an English alias for ZhouYi
type IChing = ZhouYi

// ============================================================================
// Core type definitions
// ============================================================================

// Gua represents a hexagram (64 hexagrams)
type Gua struct {
	Xu          int     // hexagram number (1-64)
	Index       string  // index key (e.g., "乾乾", "坤坤")
	ShangMing   string  // upper trigram full name (e.g., "乾为天")
	ShangNum    int     // upper trigram number (0-7)
	XiaMing     string  // lower trigram full name (e.g., "坎为水")
	XiaNum      int     // lower trigram number (0-7)
	JiXiong     string  // fortune (吉/凶/半吉)
	GuaName     string  // hexagram single-char name (e.g., "乾", "屯")
	Ming        string  // hexagram full name (e.g., "乾为天", "水雷屯")
	GuaYi       string  // hexagram meaning (邵雍易学)
	GuaSymbol   string  // Unicode hexagram symbol (e.g., ䷀)
	TuanText    string  // 彖辞 (judgment commentary)
	XiangText   string  // 大象辞 (image commentary, e.g., "天行健，君子以自强不息")
	Yaos        [6]*Yao // six lines: first, second, third, fourth, fifth, top
	Yong        string  // 用九/用六 (only for Qian and Kun)
	YongJiXiong string  // Yong fortune
}

// Yao represents a line (one line in a hexagram)
type Yao struct {
	Ci      string // line text (爻辞)
	JiXiong string // line fortune (吉/凶/平)
	NvMing  string // female-specific judgment; empty means same as general (男女通用)
}

// HasNvMing returns whether this line has a female-specific judgment.
// When false, the general JiXiong applies to both male and female.
func (y *Yao) HasNvMing() bool {
	return y.NvMing != ""
}

// Bagua represents the eight trigrams (single trigram)
// Represented using int, binary: 0=Yang, 1=Yin
// 0b000 = Qian (three Yang), 0b111 = Kun (three Yin)
type Bagua = int

// Eight trigram constants
const (
	Qian Bagua = 0b000 // Qian ☰ Heaven
	Dui  Bagua = 0b001 // Dui  ☱ Lake
	Li   Bagua = 0b010 // Li   ☲ Fire
	Zhen Bagua = 0b011 // Zhen ☳ Thunder
	Xun  Bagua = 0b100 // Xun  ☴ Wind
	Kan  Bagua = 0b101 // Kan  ☵ Water
	Gen  Bagua = 0b110 // Gen  ☶ Mountain
	Kun  Bagua = 0b111 // Kun  ☷ Earth
)

// Hexagram types (original, changed, mutual, opposite, reversed)
const (
	Ben  = iota // Original hexagram
	Bian        // Changed hexagram
	Hu          // Mutual hexagram
	Cuo         // Opposite hexagram
	Zong        // Reversed hexagram
	GuaTypeMax
)

// YaoPosition represents line position
type YaoPosition int

const (
	Chu   YaoPosition = iota // First line
	Er                       // Second line
	San                      // Third line
	Si                       // Fourth line
	Wu                       // Fifth line
	Shang                    // Top line
	YaoCount
)

// ZhouYi is the main I Ching structure
type ZhouYi struct {
	gua     [GuaTypeMax]*Gua
	bianYao []int // changing line numbers
}

// ============================================================================
// Error definitions
// ============================================================================

var (
	ErrInvalidGuaIndex = errors.New("invalid gua index")
	ErrInvalidYaoIndex = errors.New("invalid yao index")
	ErrInvalidDayanIdx = errors.New("invalid dayan index, must be positive")
	ErrGuaNotFound     = errors.New("gua not found")
)

// ============================================================================
// Core API
// ============================================================================

// Divine performs divination (by numbers)
// xia: lower trigram number, shang: upper trigram number
func Divine(xia, shang int) *ZhouYi {
	return DivineByNumber(shang, xia)
}

// DivineByNumber performs divination by numbers
// shang: upper trigram number, xia: lower trigram number, bianYao: changing line numbers (optional)
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

// DivineByTime performs divination by time
func DivineByTime(shang, xia int, t time.Time) *ZhouYi {
	by := timeToBianYao(t)
	return DivineByNumber(shang, xia, by)
}

// GetGua returns the hexagram of specified type
func (zy *ZhouYi) GetGua(guaType int) *Gua {
	if guaType < 0 || guaType >= GuaTypeMax {
		return nil
	}
	return zy.gua[guaType]
}

// GetBianYao returns the changing line position (0-5)
func (zy *ZhouYi) GetBianYao() int {
	return calcBianYao(zy.bianYao...)
}

// GetAllBianYao returns all changing line numbers as a slice.
func (zy *ZhouYi) GetAllBianYao() []int {
	result := make([]int, len(zy.bianYao))
	copy(result, zy.bianYao)
	return result
}

// GetYao returns the specified line
func (g *Gua) GetYao(pos YaoPosition) *Yao {
	if pos < 0 || int(pos) >= int(YaoCount) {
		return nil
	}
	return g.Yaos[pos]
}

// ============================================================================
// Query API
// ============================================================================

// GetGuaByIndex returns hexagram by index
// index: e.g., "QianQian", "KunKun", "KanZhen"
func GetGuaByIndex(index string) (*Gua, error) {
	if g, ok := guaStore[index]; ok {
		return g, nil
	}
	return nil, ErrGuaNotFound
}

// GetGuaByXu returns hexagram by number (1-64)
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

// GetBaguaName returns the trigram name
func GetBaguaName(bg Bagua) string {
	if bg < 0 || bg > 7 {
		return ""
	}
	return baguaNames[bg]
}

// GetBaguaSymbol returns the trigram symbol
func GetBaguaSymbol(bg Bagua) string {
	if bg < 0 || bg > 7 {
		return ""
	}
	return baguaSymbols[bg]
}

// ============================================================================
// Judgment API
// ============================================================================

// IsJi returns whether it is auspicious
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

// FilterYao filters lines
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
// Pre-generated data (generated by go generate)
// ============================================================================

// baguaNames trigram names
var baguaNames = [8]string{
	Qian: "乾", // ☰ Heaven
	Dui:  "兑", // ☱ Lake
	Li:   "离", // ☲ Fire
	Zhen: "震", // ☳ Thunder
	Xun:  "巽", // ☴ Wind
	Kan:  "坎", // ☵ Water
	Gen:  "艮", // ☶ Mountain
	Kun:  "坤", // ☷ Earth
}

// baguaSymbols trigram symbols
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

// guaStore hexagram storage (populated by go generate)
var guaStore = map[string]*Gua{}

// ============================================================================
// Internal functions
// ============================================================================

// getBenGua returns the original hexagram
func getBenGua(shang, xia int) *Gua {
	idx := GetBaguaName(shang%8) + GetBaguaName(xia%8)
	if g, ok := guaStore[idx]; ok {
		return g
	}
	return nil
}

// getBianGua returns the changed hexagram
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

// bianYaoTransform transforms changing line
func bianYaoTransform(gua, pos int) int {
	mask := 1 << (2 - uint(pos))
	if gua&mask == 0 {
		return gua | mask
	}
	return gua ^ mask
}

// calcBianYao calculates changing line position
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

// getHuGua returns the mutual hexagram
func getHuGua(ben *Gua) *Gua {
	shang := calcHuShang(ben.ShangNum, ben.XiaNum)
	xia := calcHuXia(ben.ShangNum, ben.XiaNum)
	idx := GetBaguaName(shang) + GetBaguaName(xia)
	if g, ok := guaStore[idx]; ok {
		return g
	}
	return nil
}

// calcHuShang calculates mutual upper trigram
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

// calcHuXia calculates mutual lower trigram
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

// getCuoGua returns the opposite hexagram (all Yin/Yang reversed)
func getCuoGua(ben *Gua) *Gua {
	shang := ^ben.ShangNum & 0x7
	xia := ^ben.XiaNum & 0x7
	idx := GetBaguaName(shang) + GetBaguaName(xia)
	if g, ok := guaStore[idx]; ok {
		return g
	}
	return nil
}

// getZongGua returns the reversed hexagram (upside down)
func getZongGua(ben *Gua) *Gua {
	shang := reverseBits(ben.XiaNum)
	xia := reverseBits(ben.ShangNum)
	idx := GetBaguaName(shang) + GetBaguaName(xia)
	if g, ok := guaStore[idx]; ok {
		return g
	}
	return nil
}

// reverseBits reverses 3-bit binary
func reverseBits(n int) int {
	return ((n & 0x4) >> 2) | (n & 0x2) | ((n & 0x1) << 2)
}

// timeToBianYao converts time to changing line
func timeToBianYao(t time.Time) int {
	const layout = "2006-01-02 15:04"
	parsed, _ := time.ParseInLocation(layout, t.Format(layout), t.Location())
	if n := parsed.Unix(); n != 0 {
		return int(n % 6)
	}
	return 0
}

// contains checks if string contains substring
func contains(s, substr string) bool {
	return len(s) >= len(substr) && containsHelper(s, substr)
}

func containsHelper(s, substr string) bool {
	if len(substr) == 0 {
		return true
	}
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

// ============================================================================
// Initialization (load data)
// ============================================================================

func init() {
	// Data is generated and populated by go generate
	// This is an empty implementation, actual data is in data_generated.go
}
