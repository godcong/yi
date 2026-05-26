package i18n

import "github.com/godcong/yi/core"

type Language string

const (
	LangZH Language = "zh"
	LangEN Language = "en"
)

type I18n struct {
	Ji              string
	Xiong           string
	Ping            string
	BanJi           string
	BenGua          string
	BianGua         string
	HuGua           string
	CuoGua          string
	ZongGua         string
	JieGuaResult    string
	JieGuaSection   string
	ShiYiSection    string
	ZongHeSection   string
	DongYaoAnalysis string
	GuiShu          string
	ShiYao          string
	YingYao         string
	GuaYi           string
	TuanYue         string
	XiangYue        string
	DongYaoPos      string
	YaoCi           string
	JiXiong         string
	ZongHePanduan   string
	Career          string
	Love            string
	Wealth          string
	Exams           string
	Health          string
	Travel          string
	Lawsuit         string
	Home            string
	WuXing          string
	Metal           string
	Wood            string
	Water           string
	Fire            string
	Earth           string
	Qian            string
	Dui             string
	Li              string
	Zhen            string
	Xun             string
	Kan             string
	Gen             string
	Kun             string
	ChuYao          string
	ErYao           string
	SanYao          string
	SiYao           string
	WuYao           string
	ShangYao        string
	Gong            string
	NoData          string
	EmptyResult     string
	Jia             string
	Xiang           string
	BenGong         string
	YiShi           string
	ErShi           string
	SanShi          string
	SiShi           string
	WuShi           string
	YouHun          string
	GuiHun          string
	Male            string
	Female          string
	FuMu            string
	XiongDi         string
	QiCai           string
	ZiSun           string
	GuanGui         string
}

var ZH = I18n{
	Ji: "吉", Xiong: "凶", Ping: "平", BanJi: "半吉",
	BenGua: "本卦", BianGua: "变卦", HuGua: "互卦",
	CuoGua: "错卦", ZongGua: "综卦",
	JieGuaResult:    "解  卦  结  果",
	JieGuaSection:   "【解卦】 卦象解读",
	ShiYiSection:    "【释义】 解读内容的释义",
	ZongHeSection:   "【综合判断】",
	DongYaoAnalysis: "动爻分析",
	GuiShu:   "归属",
	ShiYao:   "世爻",
	YingYao:  "应爻",
	GuaYi:    "卦义",
	TuanYue:  "彖曰",
	XiangYue: "象曰",
	DongYaoPos: "动爻位置",
	YaoCi:    "爻辞",
	JiXiong:  "吉凶",
	ZongHePanduan: "综合判断",
	Career: "事业", Love: "爱情", Wealth: "财运",
	Exams: "考试", Health: "健康", Travel: "出行",
	Lawsuit: "官司", Home: "家宅",
	WuXing: "五行", Metal: "金", Wood: "木", Water: "水", Fire: "火", Earth: "土",
	Qian: "乾", Dui: "兑", Li: "离", Zhen: "震",
	Xun: "巽", Kan: "坎", Gen: "艮", Kun: "坤",
	ChuYao: "初爻", ErYao: "二爻", SanYao: "三爻",
	SiYao: "四爻", WuYao: "五爻", ShangYao: "上爻",
	Gong: "宫", NoData: "暂无详解数据。", EmptyResult: "解卦结果为空",
	Jia: "甲", Xiang: "象",
	BenGong: "本宫", YiShi: "一世", ErShi: "二世", SanShi: "三世",
	SiShi: "四世", WuShi: "五世", YouHun: "游魂", GuiHun: "归魂",
	Male: "男", Female: "女",
	FuMu: "父母", XiongDi: "兄弟", QiCai: "妻财", ZiSun: "子孙", GuanGui: "官鬼",
}

var EN = I18n{
	Ji: "Auspicious", Xiong: "Inauspicious", Ping: "Neutral", BanJi: "Semi-Auspicious",
	BenGua: "Primary", BianGua: "Transformed", HuGua: "Nuclear",
	CuoGua: "Inverse", ZongGua: "Reverse",
	JieGuaResult:    "INTERPRETATION RESULT",
	JieGuaSection:   "Hexagram Interpretation",
	ShiYiSection:    "Detailed Explanations",
	ZongHeSection:   "Overall Judgment",
	DongYaoAnalysis: "Moving Line Analysis",
	GuiShu:   "Palace",
	ShiYao:   "Shi",
	YingYao:  "Ying",
	GuaYi:    "Meaning",
	TuanYue:  "Judgment",
	XiangYue: "Image",
	DongYaoPos: "Moving Line Position",
	YaoCi:    "Line Text",
	JiXiong:  "Fortune",
	ZongHePanduan: "Overall",
	Career: "Career", Love: "Love", Wealth: "Wealth",
	Exams: "Exams", Health: "Health", Travel: "Travel",
	Lawsuit: "Lawsuit", Home: "Home",
	WuXing: "Element", Metal: "Metal", Wood: "Wood", Water: "Water", Fire: "Fire", Earth: "Earth",
	Qian: "Qian (Heaven)", Dui: "Dui (Lake)", Li: "Li (Fire)", Zhen: "Zhen (Thunder)",
	Xun: "Xun (Wind)", Kan: "Kan (Water)", Gen: "Gen (Mountain)", Kun: "Kun (Earth)",
	ChuYao: "1st Line", ErYao: "2nd Line", SanYao: "3rd Line",
	SiYao: "4th Line", WuYao: "5th Line", ShangYao: "6th Line",
	Gong: "Palace", NoData: "No detailed interpretation available.", EmptyResult: "Interpretation result is empty",
	Jia: "Jia", Xiang: "Xiang",
	BenGong: "Pure", YiShi: "1st Change", ErShi: "2nd Change", SanShi: "3rd Change",
	SiShi: "4th Change", WuShi: "5th Change", YouHun: "Wandering Soul", GuiHun: "Returning Soul",
	Male: "Male", Female: "Female",
	FuMu: "Parents", XiongDi: "Siblings", QiCai: "Wife/Wealth", ZiSun: "Children", GuanGui: "Officer/Ghost",
}

func GetI18n(lang Language) I18n {
	switch lang {
	case LangEN:
		return EN
	default:
		return ZH
	}
}

func TranslateJiXiong(jx string, lang Language) string {
	t := GetI18n(lang)
	switch jx {
	case "吉":
		return t.Ji
	case "凶":
		return t.Xiong
	case "平":
		return t.Ping
	case "半吉":
		return t.BanJi
	default:
		return jx
	}
}

func TranslateCategory(cat string, lang Language) string {
	t := GetI18n(lang)
	switch cat {
	case "事业":
		return t.Career
	case "爱情":
		return t.Love
	case "财运":
		return t.Wealth
	case "考试":
		return t.Exams
	case "健康":
		return t.Health
	case "出行":
		return t.Travel
	case "官司":
		return t.Lawsuit
	case "家宅":
		return t.Home
	default:
		return cat
	}
}

func TranslateYaoPos(pos string, lang Language) string {
	t := GetI18n(lang)
	switch pos {
	case "初爻":
		return t.ChuYao
	case "二爻":
		return t.ErYao
	case "三爻":
		return t.SanYao
	case "四爻":
		return t.SiYao
	case "五爻":
		return t.WuYao
	case "上爻":
		return t.ShangYao
	default:
		return pos
	}
}

func TranslateBaguaName(name string, lang Language) string {
	t := GetI18n(lang)
	switch name {
	case "乾":
		return t.Qian
	case "兑":
		return t.Dui
	case "离":
		return t.Li
	case "震":
		return t.Zhen
	case "巽":
		return t.Xun
	case "坎":
		return t.Kan
	case "艮":
		return t.Gen
	case "坤":
		return t.Kun
	default:
		return name
	}
}

func TranslateWuXing(name string, lang Language) string {
	t := GetI18n(lang)
	switch name {
	case "金":
		return t.Metal
	case "木":
		return t.Wood
	case "水":
		return t.Water
	case "火":
		return t.Fire
	case "土":
		return t.Earth
	default:
		return name
	}
}

func TranslateGuaPosition(pos string, lang Language) string {
	t := GetI18n(lang)
	switch pos {
	case "本宫":
		return t.BenGong
	case "一世":
		return t.YiShi
	case "二世":
		return t.ErShi
	case "三世":
		return t.SanShi
	case "四世":
		return t.SiShi
	case "五世":
		return t.WuShi
	case "游魂":
		return t.YouHun
	case "归魂":
		return t.GuiHun
	default:
		return pos
	}
}

func TranslateSex(sex core.Sex, lang Language) string {
	t := GetI18n(lang)
	switch sex {
	case core.Male:
		return t.Male
	case core.Female:
		return t.Female
	default:
		return ""
	}
}

func TranslateFenXiCategory(cat core.FenXiCategory, lang Language) string {
	return TranslateCategory(string(cat), lang)
}

func TranslateLiuQin(lq core.LiuQin, lang Language) string {
	t := GetI18n(lang)
	switch lq {
	case core.LQFuMu:
		return t.FuMu
	case core.LQXiongDi:
		return t.XiongDi
	case core.LQQiCai:
		return t.QiCai
	case core.LQZiSun:
		return t.ZiSun
	case core.LQGuanGui:
		return t.GuanGui
	default:
		return ""
	}
}
