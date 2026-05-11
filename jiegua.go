package yi

import "strings"

// ============================================================================
// JieGua (解卦) - Hexagram Interpretation
// ============================================================================
//
// 解卦是根据起卦结果，综合本卦、变卦、互卦、错卦、综卦等信息，
// 生成结构化解读的模块。解卦不同于简单的卦象查询，它会：
//   - 综合分析本卦与变卦的关系
//   - 提取动爻爻辞作为核心判断依据
//   - 结合彖辞、象辞给出整体判断
//   - 提供事业、爱情、财运等分类解读
//   - 考虑六亲世应系统（可选）
//
// 解卦流程（传统六爻法）：
//   1. 确定本卦与变卦
//   2. 找出动爻，以动爻爻辞为主断
//   3. 参考本卦卦辞（彖辞）和象辞
//   4. 结合互卦看中间过程，错综卦看对立面
//   5. 综合六亲世应判断细节
//   6. 提取分类解读（事业/爱情/财运/等）
//
// Reference:
//   - 《增删卜易》野鹤老人
//   - 《卜筮正宗》王洪绪
//   - 《周易本义》朱熹

// FenXiCategory defines a category for hexagram aspect analysis.
type FenXiCategory string

const (
	FenXiShiYe    FenXiCategory = "事业" // career/business
	FenXiAiQing   FenXiCategory = "爱情" // love/marriage
	FenXiCaiYun   FenXiCategory = "财运" // wealth/finance
	FenXiKaoShi   FenXiCategory = "考试" // exams/academic
	FenXiJianKang FenXiCategory = "健康" // health
	FenXiChuXing  FenXiCategory = "出行" // travel
	FenXiGuanSi   FenXiCategory = "官司" // litigation
	FenXiJiaZhai  FenXiCategory = "家宅" // home/family
)

// GuaFenXi contains aspect-based interpretation for one category.
type GuaFenXi struct {
	Category FenXiCategory // aspect category
	Content  string        // interpretation text
	JiXiong  string        // fortune for this aspect (吉/凶/平)
	Source   string        // source of the interpretation ("卦义"/"爻辞")
}

// JieGuaResult represents the complete interpretation result of a divination.
type JieGuaResult struct {
	ZhouYi *ZhouYi // the divination result
	Sex    Sex     // sex of the querent

	// Core interpretation
	BenGuaInfo  *GuaInfo // original hexagram interpretation
	BianGuaInfo *GuaInfo // changed hexagram interpretation
	HuGuaInfo   *GuaInfo // mutual hexagram interpretation
	CuoGuaInfo  *GuaInfo // opposite hexagram interpretation
	ZongGuaInfo *GuaInfo // reversed hexagram interpretation

	// Changing line analysis
	DongYaoPos     YaoPosition // changing line position (if single)
	DongYaoText    string      // changing line text (爻辞)
	DongYaoJiXiong string      // changing line fortune
	IsJi           bool        // overall auspicious judgment
	JiXiongReason  string      // reason for the judgment

	// Aspect-based analysis
	FenXi []GuaFenXi // categorized interpretations

	// Detailed interpretation (释义) - pre-written multi-aspect explanations
	JieDu *GuaJieDu // from jiegua.json data, nil if not available
}

// GuaInfo contains interpretation details for a single hexagram.
type GuaInfo struct {
	Gua       *Gua   // the hexagram
	Ming      string // full name (e.g., "乾为天")
	GuaName   string // single-char name (e.g., "乾")
	GuaYi     string // hexagram meaning (邵雍易学)
	TuanText  string // 彖辞 (judgment commentary)
	XiangText string // 大象辞 (image commentary)
	JiXiong   string // fortune (吉/凶/半吉)
	Symbol    string // Unicode hexagram symbol
	GuaGong   string // palace name
	Position  string // palace position name
	ShiYao    string // 世爻 position description
	YingYao   string // 应爻 position description
}

// JieGua performs hexagram interpretation from a ZhouYi divination result.
// This is the main entry point for interpretation after any divination method.
func JieGua(zy *ZhouYi, sex Sex) *JieGuaResult {
	if zy == nil {
		return nil
	}

	result := &JieGuaResult{
		ZhouYi: zy,
		Sex:    sex,
	}

	// Build interpretation for each hexagram type
	result.BenGuaInfo = buildGuaInfo(zy.GetGua(Ben))
	result.BianGuaInfo = buildGuaInfo(zy.GetGua(Bian))
	result.HuGuaInfo = buildGuaInfo(zy.GetGua(Hu))
	result.CuoGuaInfo = buildGuaInfo(zy.GetGua(Cuo))
	result.ZongGuaInfo = buildGuaInfo(zy.GetGua(Zong))

	// Analyze changing line
	bianYaoPos := YaoPosition(zy.GetBianYao())
	result.DongYaoPos = bianYaoPos

	bianGua := zy.GetGua(Bian)
	if bianGua != nil && int(bianYaoPos) >= 0 && int(bianYaoPos) < int(YaoCount) {
		yao := bianGua.GetYao(bianYaoPos)
		if yao != nil {
			result.DongYaoText = yao.Ci
			if sex == Female && yao.HasNvMing() {
				result.DongYaoJiXiong = yao.NvMing
			} else {
				result.DongYaoJiXiong = yao.JiXiong
			}
		}
	}

	// Overall judgment
	result.IsJi = zy.IsJi(sex)
	result.JiXiongReason = buildJiXiongReason(zy, sex, result)

	// Detailed interpretation (释义) from pre-written data
	result.JieDu = GetGuaJieDu(result.BenGuaInfo.Ming)

	// Aspect-based analysis
	result.FenXi = buildFenXi(zy, sex, result)

	return result
}

// buildGuaInfo creates a GuaInfo from a Gua.
func buildGuaInfo(g *Gua) *GuaInfo {
	if g == nil {
		return nil
	}

	info := &GuaInfo{
		Gua:       g,
		Ming:      g.Ming,
		GuaName:   g.GuaName,
		GuaYi:     g.GuaYi,
		TuanText:  g.TuanText,
		XiangText: g.XiangText,
		JiXiong:   g.JiXiong,
		Symbol:    g.GuaSymbol,
	}

	// Add ShiYing info
	sy := g.GetShiYing()
	if sy != nil {
		info.GuaGong = GetBaguaName(g.GetGuaGong()) + "宫"
		info.Position = sy.Position.String()
		info.ShiYao = yaoPosName(sy.ShiPos)
		info.YingYao = yaoPosName(sy.YingPos)
	}

	return info
}

// yaoPosName returns the Chinese name for a YaoPosition.
func yaoPosName(pos YaoPosition) string {
	switch pos {
	case Chu:
		return "初爻"
	case Er:
		return "二爻"
	case San:
		return "三爻"
	case Si:
		return "四爻"
	case Wu:
		return "五爻"
	case Shang:
		return "上爻"
	default:
		return ""
	}
}

// fenXiRules maps categories to keyword sets used to extract aspect-specific
// interpretations from GuaYi and Yao.Ci text.
var fenXiRules = []struct {
	Cat  FenXiCategory
	Keys []string // keywords in source text that indicate this category
}{
	{FenXiShiYe, []string{"事业", "做官", "经商", "营谋", "经营", "谋事", "生意"}},
	{FenXiAiQing, []string{"爱情", "婚姻", "夫妻", "妇人", "女命", "求婚", "配婚", "婚嫁"}},
	{FenXiCaiYun, []string{"财运", "发财", "获利", "进财", "财利", "积蓄", "损财"}},
	{FenXiKaoShi, []string{"考试", "读书人", "学子", "功名", "科举", "佳绩"}},
	{FenXiJianKang, []string{"健康", "疾", "病", "寿", "丧", "灾疾", "足疾"}},
	{FenXiChuXing, []string{"出行", "远行", "出家", "出差", "离家", "行路"}},
	{FenXiGuanSi, []string{"官司", "争诉", "诉讼", "谗言", "刑克"}},
	{FenXiJiaZhai, []string{"家宅", "家业", "家庭", "家破", "家门", "六亲", "子嗣", "骨肉"}},
}

// fenXiToJieDuCategory maps a FenXiCategory to the corresponding JieDuCategory.
func fenXiToJieDuCategory(cat FenXiCategory) JieDuCategory {
	switch cat {
	case FenXiShiYe:
		return JieDuShiYe
	case FenXiAiQing:
		return JieDuAiQing
	case FenXiCaiYun:
		return JieDuCaiYun
	case FenXiKaoShi:
		return JieDuKaoShi
	case FenXiJianKang:
		return JieDuJianKang
	case FenXiChuXing:
		return JieDuChuXing
	case FenXiGuanSi:
		return JieDuGuanSi
	case FenXiJiaZhai:
		return JieDuJiaZhai
	default:
		return ""
	}
}

// buildFenXi extracts categorized interpretations. It prioritizes
// pre-written JieDu data (from jiegua.json) and falls back to keyword
// extraction from GuaYi and Yao.Ci when JieDu is unavailable.
func buildFenXi(zy *ZhouYi, sex Sex, result *JieGuaResult) []GuaFenXi {
	var fenXi []GuaFenXi

	benGua := zy.GetGua(Ben)
	if benGua == nil {
		return fenXi
	}

	// Collect all relevant text sources for fallback keyword extraction
	guaYi := benGua.GuaYi
	dongYaoCi := result.DongYaoText

	for _, rule := range fenXiRules {
		var content string
		var source string
		var jiXiong string

		// Primary: use pre-written detailed interpretation (释义)
		jieDuCat := fenXiToJieDuCategory(rule.Cat)
		if jieDuCat != "" {
			jieDuText := GetGuaJieDuByCategory(benGua.Index, jieDuCat)
			if jieDuText != "" {
				content = jieDuText
				source = "释义"
				jiXiong = aspectJiXiong(jieDuText)
				fenXi = append(fenXi, GuaFenXi{
					Category: rule.Cat,
					Content:  content,
					JiXiong:  jiXiong,
					Source:   source,
				})
				continue
			}
		}

		// Fallback: keyword extraction from GuaYi (general hexagram meaning)
		guayiMatch := extractAspectText(guaYi, rule.Keys)
		if guayiMatch != "" {
			content = guayiMatch
			source = "卦义"
		}

		// Extract from changing line Yao.Ci (more specific)
		yaoMatch := extractAspectText(dongYaoCi, rule.Keys)
		if yaoMatch != "" {
			if content != "" {
				content += "；" + yaoMatch
				source = "卦义+爻辞"
			} else {
				content = yaoMatch
				source = "爻辞"
			}
		}

		// Extract from all 6 Yao.Ci if still no content
		if content == "" {
			for i := 0; i < int(YaoCount); i++ {
				yao := benGua.GetYao(YaoPosition(i))
				if yao == nil {
					continue
				}
				match := extractAspectText(yao.Ci, rule.Keys)
				if match != "" {
					content = match
					source = yaoPosName(YaoPosition(i)) + "爻辞"
					break
				}
			}
		}

		if content == "" {
			continue
		}

		// Determine JiXiong for this aspect
		jiXiong = aspectJiXiong(content)

		fenXi = append(fenXi, GuaFenXi{
			Category: rule.Cat,
			Content:  content,
			JiXiong:  jiXiong,
			Source:   source,
		})
	}

	return fenXi
}

// extractAspectText extracts sentences containing any of the keywords from text.
// It splits on Chinese punctuation and returns matching sentences joined.
func extractAspectText(text string, keywords []string) string {
	if text == "" {
		return ""
	}

	// Split into sentences on common Chinese punctuation
	seps := []string{"。", "，", "；", "！", "？"}
	sentences := []string{text}
	for _, sep := range seps {
		var next []string
		for _, s := range sentences {
			for _, part := range strings.Split(s, sep) {
				part = strings.TrimSpace(part)
				if part != "" {
					next = append(next, part)
				}
			}
		}
		sentences = next
	}

	var matches []string
	for _, s := range sentences {
		for _, kw := range keywords {
			if strings.Contains(s, kw) {
				matches = append(matches, s)
				break
			}
		}
	}

	if len(matches) == 0 {
		return ""
	}
	return strings.Join(matches, "；")
}

// aspectJiXiong determines fortune level from aspect text content.
func aspectJiXiong(text string) string {
	jiWords := []string{"吉", "如意", "称意", "获利", "升迁", "佳绩", "成功", "和睦", "长寿", "健康", "得利"}
	xiongWords := []string{"凶", "灾", "不利", "不遂", "祸", "疾", "丧", "贬", "退", "难", "破"}

	jiScore := 0
	xiongScore := 0

	for _, w := range jiWords {
		if strings.Contains(text, w) {
			jiScore++
		}
	}
	for _, w := range xiongWords {
		if strings.Contains(text, w) {
			xiongScore++
		}
	}

	switch {
	case jiScore > xiongScore:
		return "吉"
	case xiongScore > jiScore:
		return "凶"
	default:
		return "平"
	}
}

// buildJiXiongReason constructs the reason string for the judgment.
func buildJiXiongReason(zy *ZhouYi, sex Sex, result *JieGuaResult) string {
	benGua := zy.GetGua(Ben)
	bianGua := zy.GetGua(Bian)
	if benGua == nil || bianGua == nil {
		return "卦象数据不完整，无法判断"
	}

	var reason string

	// 1. Original hexagram fortune
	reason += "本卦【" + benGua.Ming + "】" + benGua.JiXiong + "；"

	// 2. Changed hexagram fortune
	reason += "变卦【" + bianGua.Ming + "】" + bianGua.JiXiong + "；"

	// 3. Changing line text
	if result.DongYaoText != "" {
		posName := yaoPosName(result.DongYaoPos)
		reason += "动爻在" + posName + "，爻辞「" + result.DongYaoText + "」，" + result.DongYaoJiXiong + "；"
	}

	// 4. Xiang commentary (大象辞)
	if benGua.XiangText != "" {
		reason += "象曰：" + benGua.XiangText + "；"
	}

	// 5. Overall
	if result.IsJi {
		reason += "综合判断：吉"
	} else {
		reason += "综合判断：凶"
	}

	return reason
}

// FormatJieGua formats the interpretation result as a readable string.
// Structure:
//   【解卦】 - 卦象的解读（卦象分析、卦义、彖辞、象辞）
//   【释义】 - 解读内容的释义（事业/爱情/财运等分类详解）
//   【综合判断】 - 整体吉凶判断
func FormatJieGua(result *JieGuaResult) string {
	if result == nil {
		return "解卦结果为空"
	}

	var output string

	// Header
	output += "═══════════════════════════════════\n"
	output += "           解  卦  结  果\n"
	output += "═══════════════════════════════════\n\n"

	// ─────────────────────────────────────────
	// Section 1: 解卦 — Hexagram interpretation
	// ─────────────────────────────────────────
	output += "━━━━━━━━ 【解卦】 卦象解读 ━━━━━━━━\n\n"

	// Original hexagram
	if result.BenGuaInfo != nil {
		output += "【本卦】" + result.BenGuaInfo.Symbol + " " + result.BenGuaInfo.Ming
		output += "（" + result.BenGuaInfo.JiXiong + "）\n"
		if result.BenGuaInfo.GuaGong != "" {
			output += "  归属：" + result.BenGuaInfo.GuaGong + "·" + result.BenGuaInfo.Position + "\n"
			output += "  世爻：" + result.BenGuaInfo.ShiYao + "  应爻：" + result.BenGuaInfo.YingYao + "\n"
		}
		if result.BenGuaInfo.GuaYi != "" {
			output += "  卦义：" + result.BenGuaInfo.GuaYi + "\n"
		}
		if result.BenGuaInfo.TuanText != "" {
			output += "  彖曰：" + result.BenGuaInfo.TuanText + "\n"
		}
		if result.BenGuaInfo.XiangText != "" {
			output += "  象曰：" + result.BenGuaInfo.XiangText + "\n"
		}
	}

	// Changed hexagram
	if result.BianGuaInfo != nil {
		output += "\n【变卦】" + result.BianGuaInfo.Symbol + " " + result.BianGuaInfo.Ming
		output += "（" + result.BianGuaInfo.JiXiong + "）\n"
	}

	// Related hexagrams
	if result.HuGuaInfo != nil {
		output += "【互卦】" + result.HuGuaInfo.Symbol + " " + result.HuGuaInfo.Ming + "\n"
	}
	if result.CuoGuaInfo != nil {
		output += "【错卦】" + result.CuoGuaInfo.Symbol + " " + result.CuoGuaInfo.Ming + "\n"
	}
	if result.ZongGuaInfo != nil {
		output += "【综卦】" + result.ZongGuaInfo.Symbol + " " + result.ZongGuaInfo.Ming + "\n"
	}

	// Changing line analysis
	if result.DongYaoText != "" {
		output += "\n【动爻分析】\n"
		output += "  动爻位置：" + yaoPosName(result.DongYaoPos) + "\n"
		output += "  爻辞：" + result.DongYaoText + "\n"
		output += "  吉凶：" + result.DongYaoJiXiong + "\n"
	}

	// ─────────────────────────────────────────
	// Section 2: 释义 — Detailed explanations
	// ─────────────────────────────────────────
	output += "\n━━━━━━ 【释义】 解读内容的释义 ━━━━━━\n\n"

	if len(result.FenXi) > 0 {
		for _, fx := range result.FenXi {
			output += "【" + string(fx.Category) + "】"
			if fx.JiXiong != "" {
				output += "（" + fx.JiXiong + "）"
			}
			output += "\n  " + fx.Content + "\n"
		}
	} else {
		output += "  暂无详解数据。\n"
	}

	// ─────────────────────────────────────────
	// Section 3: 综合判断 — Overall judgment
	// ─────────────────────────────────────────
	output += "\n━━━━━━━━ 【综合判断】 ━━━━━━━━\n\n"
	if result.IsJi {
		output += "  ✦ 吉 ✦\n"
	} else {
		output += "  ✧ 凶 ✧\n"
	}
	if result.JiXiongReason != "" {
		output += "  " + result.JiXiongReason + "\n"
	}

	output += "\n═══════════════════════════════════\n"

	return output
}
