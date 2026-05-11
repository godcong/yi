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
	Ming     string // hexagram full name (e.g., "乾为天")
	ShiYe    string // 事业
	AiQing   string // 爱情
	CaiYun   string // 财运
	KaoShi   string // 考试
	JianKang string // 健康
	ChuXing  string // 出行
	GuanSi   string // 官司
	JiaZhai  string // 家宅
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
