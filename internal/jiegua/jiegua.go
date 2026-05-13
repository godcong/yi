package jiegua

import (
	"yi/core"
	"yi/internal/gua"
	"yi/internal/i18n"
	"yi/internal/shiying"
)

func JieGua(zy *core.ZhouYi, sex core.Sex) *core.JieGuaResult {
	return JieGuaWithLang(zy, sex, i18n.LangZH)
}

func JieGuaWithLang(zy *core.ZhouYi, sex core.Sex, lang i18n.Language) *core.JieGuaResult {
	if zy == nil {
		return nil
	}

	result := &core.JieGuaResult{
		ZhouYi: zy,
		Sex:    sex,
	}

	result.BenGuaInfo = buildGuaInfo(zy.Gua[core.Ben])
	result.BianGuaInfo = buildGuaInfo(zy.Gua[core.Bian])
	result.HuGuaInfo = buildGuaInfo(zy.Gua[core.Hu])
	result.CuoGuaInfo = buildGuaInfo(zy.Gua[core.Cuo])
	result.ZongGuaInfo = buildGuaInfo(zy.Gua[core.Zong])

	bianYaoPos := core.YaoPosition(gua.GetBianYao(zy))
	result.DongYaoPos = bianYaoPos

	bianGua := zy.Gua[core.Bian]
	if bianGua != nil && int(bianYaoPos) >= 0 && int(bianYaoPos) < int(core.YaoCount) {
		yao := bianGua.Yaos[bianYaoPos]
		if yao != nil {
			result.DongYaoText = yao.Ci
			if sex == core.Female && yao.NvMing != "" {
				result.DongYaoJiXiong = yao.NvMing
			} else {
				result.DongYaoJiXiong = yao.JiXiong
			}
		}
	}

	result.IsJi = gua.IsJi(zy, sex)
	result.JiXiongReason = buildJiXiongReasonWithLang(zy, sex, result, lang)

	result.JieDu = GetGuaJieDuByIndex(zy.Gua[core.Ben].Index)

	if benGua := zy.Gua[core.Ben]; benGua != nil {
		result.WuXingInfo = GetWuXingInfo(benGua.ShangNum)
	}

	result.FenXi = buildFenXi(zy, sex, result)

	return result
}

func buildGuaInfo(g *core.Gua) *core.GuaInfo {
	if g == nil {
		return nil
	}

	info := &core.GuaInfo{
		Gua:       g,
		Ming:      g.Ming,
		GuaName:   g.GuaName,
		GuaYi:     g.GuaYi,
		TuanText:  g.TuanText,
		XiangText: g.XiangText,
		JiXiong:   g.JiXiong,
		Symbol:    g.GuaSymbol,
	}

	sy := shiying.GetShiYing(g)
	if sy != nil {
		guaGong := shiying.GetGuaGong(g)
		if guaGong >= 0 && guaGong <= 7 {
			info.GuaGong = core.BaguaNames[guaGong] + "宫"
		}
		info.Position = sy.Position.String()
		info.ShiYao = yaoPosName(sy.ShiPos)
		info.YingYao = yaoPosName(sy.YingPos)
	}

	return info
}

func yaoPosName(pos core.YaoPosition) string {
	switch pos {
	case core.Chu:
		return "初爻"
	case core.Er:
		return "二爻"
	case core.San:
		return "三爻"
	case core.Si:
		return "四爻"
	case core.Wu:
		return "五爻"
	case core.Shang:
		return "上爻"
	default:
		return ""
	}
}

func buildJiXiongReasonWithLang(zy *core.ZhouYi, sex core.Sex, result *core.JieGuaResult, lang i18n.Language) string {
	t := i18n.GetI18n(lang)
	benGua := zy.Gua[core.Ben]
	bianGua := zy.Gua[core.Bian]
	if benGua == nil || bianGua == nil {
		if lang == i18n.LangEN {
			return "Hexagram data incomplete, cannot judge"
		}
		return "卦象数据不完整，无法判断"
	}

	var reason string

	reason += "【" + t.BenGua + "】" + benGua.Ming + " " + i18n.TranslateJiXiong(benGua.JiXiong, lang) + "; "
	reason += "【" + t.BianGua + "】" + bianGua.Ming + " " + i18n.TranslateJiXiong(bianGua.JiXiong, lang) + "; "

	if result.DongYaoText != "" {
		posName := i18n.TranslateYaoPos(yaoPosName(result.DongYaoPos), lang)
		reason += t.DongYaoPos + " " + posName + ", " + t.YaoCi + "「" + result.DongYaoText + "」, " + i18n.TranslateJiXiong(result.DongYaoJiXiong, lang) + "; "
	}

	if benGua.XiangText != "" {
		reason += t.XiangYue + ": " + benGua.XiangText + "; "
	}

	if result.IsJi {
		reason += t.ZongHePanduan + ": " + t.Ji
	} else {
		reason += t.ZongHePanduan + ": " + t.Xiong
	}

	return reason
}
