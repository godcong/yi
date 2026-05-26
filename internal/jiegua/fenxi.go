package jiegua

import (
	"strings"

	"github.com/godcong/yi/core"
)

var fenXiRules = []struct {
	Cat  core.FenXiCategory
	Keys []string
}{
	{core.FenXiShiYe, []string{"事业", "做官", "经商", "营谋", "经营", "谋事", "生意"}},
	{core.FenXiAiQing, []string{"爱情", "婚姻", "夫妻", "妇人", "女命", "求婚", "配婚", "婚嫁"}},
	{core.FenXiCaiYun, []string{"财运", "发财", "获利", "进财", "财利", "积蓄", "损财"}},
	{core.FenXiKaoShi, []string{"考试", "读书人", "学子", "功名", "科举", "佳绩"}},
	{core.FenXiJianKang, []string{"健康", "疾", "病", "寿", "丧", "灾疾", "足疾"}},
	{core.FenXiChuXing, []string{"出行", "远行", "出家", "出差", "离家", "行路"}},
	{core.FenXiGuanSi, []string{"官司", "争诉", "诉讼", "谗言", "刑克"}},
	{core.FenXiJiaZhai, []string{"家宅", "家业", "家庭", "家破", "家门", "六亲", "子嗣", "骨肉"}},
}

func fenXiToJieDuCategory(cat core.FenXiCategory) core.JieDuCategory {
	switch cat {
	case core.FenXiShiYe:
		return core.JieDuShiYe
	case core.FenXiAiQing:
		return core.JieDuAiQing
	case core.FenXiCaiYun:
		return core.JieDuCaiYun
	case core.FenXiKaoShi:
		return core.JieDuKaoShi
	case core.FenXiJianKang:
		return core.JieDuJianKang
	case core.FenXiChuXing:
		return core.JieDuChuXing
	case core.FenXiGuanSi:
		return core.JieDuGuanSi
	case core.FenXiJiaZhai:
		return core.JieDuJiaZhai
	default:
		return ""
	}
}

func buildFenXi(zy *core.ZhouYi, sex core.Sex, result *core.JieGuaResult) []core.GuaFenXi {
	var fenXi []core.GuaFenXi

	benGua := zy.Gua[core.Ben]
	if benGua == nil {
		return fenXi
	}

	guaYi := benGua.GuaYi
	dongYaoCi := result.DongYaoText

	for _, rule := range fenXiRules {
		var content string
		var source string
		var jiXiong string

		var kwContent string
		var kwSource string

		guayiMatch := extractAspectText(guaYi, rule.Keys)
		if guayiMatch != "" {
			kwContent = guayiMatch
			kwSource = "卦义"
		}

		yaoMatch := extractAspectText(dongYaoCi, rule.Keys)
		if yaoMatch != "" {
			if kwContent != "" {
				kwContent += "；" + yaoMatch
				kwSource = "卦义+爻辞"
			} else {
				kwContent = yaoMatch
				kwSource = "爻辞"
			}
		}

		if kwContent == "" {
			for i := 0; i < int(core.YaoCount); i++ {
				yao := benGua.Yaos[core.YaoPosition(i)]
				if yao == nil {
					continue
				}
				match := extractAspectText(yao.Ci, rule.Keys)
				if match != "" {
					kwContent = match
					kwSource = yaoPosName(core.YaoPosition(i)) + "爻辞"
					break
				}
			}
		}

		var jieDuText string
		jieDuCat := fenXiToJieDuCategory(rule.Cat)
		if jieDuCat != "" {
			jieDuText = GetGuaJieDuByCategory(benGua.Index, jieDuCat)
		}

		switch {
		case jieDuText != "" && kwContent != "" && jieDuText != kwContent:
			content = jieDuText + "；" + kwContent
			source = "释义+" + kwSource
		case jieDuText != "":
			content = jieDuText
			source = "释义"
		case kwContent != "":
			content = kwContent
			source = kwSource
		default:
			continue
		}

		jiXiong = aspectJiXiong(content)

		fenXi = append(fenXi, core.GuaFenXi{
			Category: rule.Cat,
			Content:  content,
			JiXiong:  jiXiong,
			Source:   source,
		})
	}

	return fenXi
}

func extractAspectText(text string, keywords []string) string {
	if text == "" {
		return ""
	}

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
