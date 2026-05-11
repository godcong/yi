package yi

// ============================================================================
// GuaJieDu (卦解读) - Detailed Hexagram Interpretation Data
// ============================================================================
//
// Each of the 64 hexagrams has pre-written detailed interpretations for
// 8 life aspects: 事业, 爱情, 财运, 考试, 健康, 出行, 官司, 家宅.
//
// These interpretations are based on traditional I Ching scholarship:
//   - 《增删卜易》野鹤老人
//   - 《卜筮正宗》王洪绪
//   - 《周易本义》朱熹
//   - 《断易天机》
//   - 《周易折中》李光地

// JieDuCategory identifies one of the 8 life-aspect categories.
type JieDuCategory string

const (
	JieDuShiYe    JieDuCategory = "事业" // career / business
	JieDuAiQing   JieDuCategory = "爱情" // love / marriage
	JieDuCaiYun   JieDuCategory = "财运" // wealth / finance
	JieDuKaoShi   JieDuCategory = "考试" // exams / academic
	JieDuJianKang JieDuCategory = "健康" // health
	JieDuChuXing  JieDuCategory = "出行" // travel
	JieDuGuanSi   JieDuCategory = "官司" // litigation
	JieDuJiaZhai  JieDuCategory = "家宅" // home / family
)

// AllJieDuCategories returns all 8 interpretation categories in order.
var AllJieDuCategories = []JieDuCategory{
	JieDuShiYe, JieDuAiQing, JieDuCaiYun, JieDuKaoShi,
	JieDuJianKang, JieDuChuXing, JieDuGuanSi, JieDuJiaZhai,
}

// GuaJieDu holds the detailed interpretation for one hexagram.
type GuaJieDu struct {
	Ming      string   // hexagram full name (e.g., "乾为天")
	ShiYe     string   // 事业
	AiQing    string   // 爱情
	CaiYun    string   // 财运
	KaoShi    string   // 考试
	JianKang  string   // 健康
	ChuXing   string   // 出行
	GuanSi    string   // 官司
	JiaZhai   string   // 家宅
	CoreImage string   // 核心意象（一句话概括卦象本质）
	Yi        []string // 宜（适合做的事）
	Ji        []string // 忌（应避免的事）
}

// guaJieDuStore maps hexagram Ming (e.g., "乾为天") to its interpretation.
var guaJieDuStore map[string]*GuaJieDu

// GetGuaJieDu returns the detailed interpretation for a hexagram by its Ming.
// Returns nil if not found.
func GetGuaJieDu(ming string) *GuaJieDu {
	if guaJieDuStore == nil {
		return nil
	}
	return guaJieDuStore[ming]
}

// GetGuaJieDuByIndex returns the detailed interpretation by hexagram index (e.g., "乾乾").
// This is the primary lookup method since the store keys are index-based.
func GetGuaJieDuByIndex(index string) *GuaJieDu {
	if guaJieDuStore == nil {
		return nil
	}
	return guaJieDuStore[index]
}

// WuXingInfo holds the five-element derived lucky attributes for a hexagram.
type WuXingInfo struct {
	WuXing      string // 五行（金/木/水/火/土）
	Direction   string // 幸运方位
	LuckyNumber string // 幸运数字
	LuckyColor  string // 幸运颜色
}

// wuXingMapping maps Bagua number to its Wu Xing (five element).
var wuXingMapping = map[Bagua]string{
	0: "金", // 乾
	1: "金", // 兑
	2: "火", // 离
	3: "木", // 震
	4: "木", // 巽
	5: "水", // 坎
	6: "土", // 艮
	7: "土", // 坤
}

// wuXingAttributes maps Wu Xing to its direction, lucky number, and color.
var wuXingAttributes = map[string]WuXingInfo{
	"金": {WuXing: "金", Direction: "西、西北", LuckyNumber: "4、9", LuckyColor: "白、金、银"},
	"木": {WuXing: "木", Direction: "东、东南", LuckyNumber: "3、8", LuckyColor: "绿、青"},
	"水": {WuXing: "水", Direction: "北", LuckyNumber: "1、6", LuckyColor: "黑、蓝"},
	"火": {WuXing: "火", Direction: "南", LuckyNumber: "2、7", LuckyColor: "红、紫"},
	"土": {WuXing: "土", Direction: "东北、西南", LuckyNumber: "5、10", LuckyColor: "黄、棕"},
}

// GetWuXingInfo returns the five-element info derived from a hexagram's upper trigram.
func GetWuXingInfo(shangNum int) *WuXingInfo {
	wx, ok := wuXingMapping[Bagua(shangNum)]
	if !ok {
		wx = "土" // default
	}
	info, ok := wuXingAttributes[wx]
	if !ok {
		return nil
	}
	return &info
}

// GetGuaJieDuByCategory returns one category's interpretation text.
// Returns empty string if not found.
func GetGuaJieDuByCategory(ming string, cat JieDuCategory) string {
	jd := GetGuaJieDu(ming)
	if jd == nil {
		return ""
	}
	switch cat {
	case JieDuShiYe:
		return jd.ShiYe
	case JieDuAiQing:
		return jd.AiQing
	case JieDuCaiYun:
		return jd.CaiYun
	case JieDuKaoShi:
		return jd.KaoShi
	case JieDuJianKang:
		return jd.JianKang
	case JieDuChuXing:
		return jd.ChuXing
	case JieDuGuanSi:
		return jd.GuanSi
	case JieDuJiaZhai:
		return jd.JiaZhai
	default:
		return ""
	}
}

// AllJieDuEntries returns all 64 hexagram interpretations as a slice.
// Iterates over guaStore keys to maintain consistency with hexagram order.
func AllJieDuEntries() []*GuaJieDu {
	if guaJieDuStore == nil {
		return nil
	}
	result := make([]*GuaJieDu, 0, len(guaJieDuStore))
	for _, gua := range guaStore {
		if gua != nil {
			if jd, ok := guaJieDuStore[gua.Index]; ok {
				result = append(result, jd)
			}
		}
	}
	return result
}
