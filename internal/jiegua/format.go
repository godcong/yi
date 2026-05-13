package jiegua

import (
	"yi/core"
	"yi/internal/i18n"
)

func FormatJieGua(result *core.JieGuaResult) string {
	return FormatJieGuaWithLang(result, i18n.LangZH)
}

func FormatJieGuaWithLang(result *core.JieGuaResult, lang i18n.Language) string {
	if result == nil {
		return i18n.GetI18n(lang).EmptyResult
	}

	t := i18n.GetI18n(lang)
	var output string

	output += "═══════════════════════════════════\n"
	output += "           " + t.JieGuaResult + "\n"
	output += "═══════════════════════════════════\n\n"

	output += "━━━━━━━━ " + t.JieGuaSection + " ━━━━━━━━\n\n"

	if result.BenGuaInfo != nil {
		output += "【" + t.BenGua + "】" + result.BenGuaInfo.Symbol + " " + result.BenGuaInfo.Ming
		output += "（" + i18n.TranslateJiXiong(result.BenGuaInfo.JiXiong, lang) + "）\n"
		if result.BenGuaInfo.GuaGong != "" {
			guaGong := result.BenGuaInfo.GuaGong
			if lang == i18n.LangEN {
				guaGong = i18n.TranslateBaguaName(guaGong, lang)
			}
			output += "  " + t.GuiShu + ": " + guaGong + "·" + result.BenGuaInfo.Position + "\n"
			output += "  " + t.ShiYao + ": " + i18n.TranslateYaoPos(result.BenGuaInfo.ShiYao, lang) + "  " + t.YingYao + ": " + i18n.TranslateYaoPos(result.BenGuaInfo.YingYao, lang) + "\n"
		}
		if result.BenGuaInfo.GuaYi != "" {
			output += "  " + t.GuaYi + ": " + result.BenGuaInfo.GuaYi + "\n"
		}
		if result.BenGuaInfo.TuanText != "" {
			output += "  " + t.TuanYue + ": " + result.BenGuaInfo.TuanText + "\n"
		}
		if result.BenGuaInfo.XiangText != "" {
			output += "  " + t.XiangYue + ": " + result.BenGuaInfo.XiangText + "\n"
		}
	}

	if result.BianGuaInfo != nil {
		output += "\n【" + t.BianGua + "】" + result.BianGuaInfo.Symbol + " " + result.BianGuaInfo.Ming
		output += "（" + i18n.TranslateJiXiong(result.BianGuaInfo.JiXiong, lang) + "）\n"
	}

	if result.HuGuaInfo != nil {
		output += "【" + t.HuGua + "】" + result.HuGuaInfo.Symbol + " " + result.HuGuaInfo.Ming + "\n"
	}
	if result.CuoGuaInfo != nil {
		output += "【" + t.CuoGua + "】" + result.CuoGuaInfo.Symbol + " " + result.CuoGuaInfo.Ming + "\n"
	}
	if result.ZongGuaInfo != nil {
		output += "【" + t.ZongGua + "】" + result.ZongGuaInfo.Symbol + " " + result.ZongGuaInfo.Ming + "\n"
	}

	if result.DongYaoText != "" {
		output += "\n【" + t.DongYaoAnalysis + "】\n"
		output += "  " + t.DongYaoPos + ": " + i18n.TranslateYaoPos(yaoPosName(result.DongYaoPos), lang) + "\n"
		output += "  " + t.YaoCi + ": " + result.DongYaoText + "\n"
		output += "  " + t.JiXiong + ": " + i18n.TranslateJiXiong(result.DongYaoJiXiong, lang) + "\n"
	}

	output += "\n━━━━━━ " + t.ShiYiSection + " ━━━━━━\n\n"

	if len(result.FenXi) > 0 {
		for _, fx := range result.FenXi {
			output += "【" + i18n.TranslateCategory(string(fx.Category), lang) + "】"
			if fx.JiXiong != "" {
				output += "（" + i18n.TranslateJiXiong(fx.JiXiong, lang) + "）"
			}
			output += "\n  " + fx.Content + "\n"
		}
	} else {
		output += "  " + t.NoData + "\n"
	}

	output += "\n━━━━━━━━ " + t.ZongHeSection + " ━━━━━━━━\n\n"
	if result.IsJi {
		output += "  ✦ " + t.Ji + " ✦\n"
	} else {
		output += "  ✧ " + t.Xiong + " ✧\n"
	}
	if result.JiXiongReason != "" {
		output += "  " + result.JiXiongReason + "\n"
	}

	output += "\n═══════════════════════════════════\n"

	return output
}
