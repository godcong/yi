package yi

import "errors"

// ============================================================================
// JiaZi (甲子) - Sixty-year cycle information
// ============================================================================
//
// 甲子是中国古代的干支纪年系统，由十天干（甲乙丙丁戊己庚辛壬癸）
// 与十二地支（子丑寅卯辰巳午未申酉戌亥）按顺序组合，形成六十个单位，
// 称为"六十甲子"或"六十花甲子"。
//
// 天干地支与五行、纳音密切相关，广泛用于：
//   - 纪年、纪月、纪日、纪时
//   - 六爻纳甲法
//   - 命理学（四柱八字）
//
// Reference:
//   - 《渊海子平》
//   - 《三命通会》

// TianGan represents the Ten Heavenly Stems (天干).
type TianGan int

const (
	TGJia  TianGan = iota // 甲
	TGYi                  // 乙
	TGBing                // 丙
	TGDing                // 丁
	TGWu                  // 戊
	TGJi                  // 己
	TGGeng                // 庚
	TGXin                 // 辛
	TGRen                 // 壬
	TGGui                 // 癸
	TianGanCount
)

// DiZhi represents the Twelve Earthly Branches (地支).
type DiZhi int

const (
	DZZi   DiZhi = iota // 子
	DZChou              // 丑
	DZYin               // 寅
	DZMao               // 卯
	DZChen              // 辰
	DZSi                // 巳
	DZWu                // 午
	DZWei               // 未
	DZShen              // 申
	DZYou               // 酉
	DZXu                // 戌
	DZHai               // 亥
	DiZhiCount
)

// JiaZiInfo represents a single entry in the sixty-year cycle.
type JiaZiInfo struct {
	Index   int    // 1-60, position in the cycle
	Name    string // full name (e.g., "甲子", "乙丑")
	TianGan string // heavenly stem name
	DiZhi   string // earthly branch name
	NaYin   string // NaYin (纳音) name of the pair (e.g., "海中金", "炉中火")
	YinYang string // yin/yang of the heavenly stem
}

// WuXing returns the NaYin field for backward compatibility.
// Deprecated: use NaYin instead.
func (j *JiaZiInfo) WuXing() string {
	return j.NaYin
}

// GetNaYinWuXing extracts the base WuXing (五行) from the NaYin (纳音) name.
// For example, "海中金" → Metal, "炉中火" → Fire, "大林木" → Wood.
func (j *JiaZiInfo) GetNaYinWuXing() WuXing {
	return naYinToWuXing(j.NaYin)
}

// naYinToWuXing maps a NaYin name to its base WuXing element.
func naYinToWuXing(naYin string) WuXing {
	if naYin == "" {
		return 0
	}
	// NaYin names end with the element character: 金/木/水/火/土
	// Use utf8 to get the last rune
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

// GetJiaZi returns the JiaZi info for a given index (1-60).
func GetJiaZi(index int) (*JiaZiInfo, error) {
	if index < 1 || index > 60 {
		return nil, ErrInvalidJiaZiIndex
	}
	return &jiaZiList[index-1], nil
}

// MustGetJiaZi returns the JiaZi info (ignoring errors).
func MustGetJiaZi(index int) JiaZiInfo {
	if index < 1 || index > 60 {
		return JiaZiInfo{}
	}
	return jiaZiList[(index-1)%60]
}

// GetJiaZiByGanZhi returns the JiaZi info by TianGan and DiZhi combination.
func GetJiaZiByGanZhi(gan TianGan, zhi DiZhi) (*JiaZiInfo, error) {
	if int(gan) < 0 || int(gan) >= int(TianGanCount) {
		return nil, ErrInvalidTianGan
	}
	if int(zhi) < 0 || int(zhi) >= int(DiZhiCount) {
		return nil, ErrInvalidDiZhi
	}
	// Valid JiaZi: gan and zhi must have same parity (both even or both odd)
	if int(gan)%2 != int(zhi)%2 {
		return nil, ErrInvalidJiaZiCombo
	}
	idx := jiaZiComboIndex(gan, zhi)
	return &jiaZiList[idx], nil
}

// GetTianGanName returns the name of a Heavenly Stem.
func GetTianGanName(gan TianGan) string {
	if gan < 0 || int(gan) >= int(TianGanCount) {
		return ""
	}
	return tianGanNames[gan]
}

// GetDiZhiName returns the name of an Earthly Branch.
func GetDiZhiName(zhi DiZhi) string {
	if zhi < 0 || int(zhi) >= int(DiZhiCount) {
		return ""
	}
	return diZhiNames[zhi]
}

// GetTianGanWuXing returns the WuXing of a Heavenly Stem.
func GetTianGanWuXing(gan TianGan) WuXing {
	switch gan {
	case TGJia, TGYi:
		return Wood
	case TGBing, TGDing:
		return Fire
	case TGWu, TGJi:
		return Earth
	case TGGeng, TGXin:
		return Metal
	case TGRen, TGGui:
		return Water
	default:
		return 0
	}
}

// GetDiZhiWuXing returns the WuXing of an Earthly Branch.
func GetDiZhiWuXing(zhi DiZhi) WuXing {
	switch zhi {
	case DZYin, DZMao:
		return Wood
	case DZSi, DZWu:
		return Fire
	case DZChen, DZXu, DZChou, DZWei:
		return Earth
	case DZShen, DZYou:
		return Metal
	case DZZi, DZHai:
		return Water
	default:
		return 0
	}
}

// GetTianGanYinYang returns the YinYang of a Heavenly Stem.
func GetTianGanYinYang(gan TianGan) YinYang {
	return GetYinYangByNumber(int(gan) + 1)
}

// GetDiZhiYinYang returns the YinYang of an Earthly Branch.
func GetDiZhiYinYang(zhi DiZhi) YinYang {
	return GetYinYangByNumber(int(zhi) + 1)
}

// jiaZiComboIndex calculates the 0-based index in the 60-cycle
// from a given TianGan and DiZhi.
func jiaZiComboIndex(gan TianGan, zhi DiZhi) int {
	// The index satisfies: index % 10 == gan and index % 12 == zhi
	for i := 0; i < 60; i++ {
		if i%10 == int(gan) && i%12 == int(zhi) {
			return i
		}
	}
	return 0
}

// Error definitions for JiaZi
var (
	ErrInvalidJiaZiIndex = errors.New("jiazi index must be between 1 and 60")
	ErrInvalidTianGan    = errors.New("invalid tiangan index")
	ErrInvalidDiZhi      = errors.New("invalid dizhi index")
	ErrInvalidJiaZiCombo = errors.New("invalid jiazi combination: gan and zhi must have same parity")
)

// tianGanNames maps TianGan to its Chinese name.
var tianGanNames = [TianGanCount]string{
	TGJia: "甲", TGYi: "乙", TGBing: "丙", TGDing: "丁", TGWu: "戊",
	TGJi: "己", TGGeng: "庚", TGXin: "辛", TGRen: "壬", TGGui: "癸",
}

// diZhiNames maps DiZhi to its Chinese name.
var diZhiNames = [DiZhiCount]string{
	DZZi: "子", DZChou: "丑", DZYin: "寅", DZMao: "卯", DZChen: "辰", DZSi: "巳",
	DZWu: "午", DZWei: "未", DZShen: "申", DZYou: "酉", DZXu: "戌", DZHai: "亥",
}

// jiaZiList is populated in data_generated.go
var jiaZiList [60]JiaZiInfo
