package yi

import (
	"time"

	"github.com/godcong/yi/core"
	"github.com/godcong/yi/internal/gua"
	"github.com/godcong/yi/internal/i18n"
	"github.com/godcong/yi/internal/jiegua"
	"github.com/godcong/yi/internal/numerology"
	"github.com/godcong/yi/internal/qigua"
	"github.com/godcong/yi/internal/shiying"
	"github.com/godcong/yi/internal/wenyan"
	"github.com/godcong/yi/internal/wuxing"
)

type Gua = core.Gua

type Yao = core.Yao

type ZhouYi = core.ZhouYi

type Bagua = core.Bagua

type YaoPosition = core.YaoPosition

type Sex = core.Sex

type YinYang = core.YinYang

type WuXing = core.WuXing

type GuaPosition = core.GuaPosition

type ShiYingInfo = core.ShiYingInfo

type LiuQin = core.LiuQin

type CoinResult = core.CoinResult

type DayanStep = core.DayanStep

type DayanResult = core.DayanResult

type Dayan = core.Dayan

type TimeGuaParams = core.TimeGuaParams

type BirthdayParams = core.BirthdayParams

type WenYanData = core.WenYanData

type WenYanEntry = core.WenYanEntry

type JiaZiInfo = core.JiaZiInfo

type TianGan = core.TianGan

type DiZhi = core.DiZhi

type JieGuaResult = core.JieGuaResult

type GuaInfo = core.GuaInfo

type GuaFenXi = core.GuaFenXi

type FenXiCategory = core.FenXiCategory

type GuaJieDu = core.GuaJieDu

type JieDuCategory = core.JieDuCategory

type WuXingInfo = core.WuXingInfo

type Language = i18n.Language

type Hexagram = core.Gua

type Line = core.Yao

type IChing = core.ZhouYi

type Trigram = core.Bagua

const (
	Ben       = core.Ben
	Bian      = core.Bian
	Hu        = core.Hu
	Cuo       = core.Cuo
	Zong      = core.Zong
	GuaTypeMax = core.GuaTypeMax
)

const (
	Chu      YaoPosition = core.Chu
	Er       YaoPosition = core.Er
	San      YaoPosition = core.San
	Si       YaoPosition = core.Si
	Wu       YaoPosition = core.Wu
	Shang    YaoPosition = core.Shang
	YaoCount             = core.YaoCount
)

const (
	Qian Bagua = core.Qian
	Dui  Bagua = core.Dui
	Li   Bagua = core.Li
	Zhen Bagua = core.Zhen
	Xun  Bagua = core.Xun
	Kan  Bagua = core.Kan
	Gen  Bagua = core.Gen
	Kun  Bagua = core.Kun
)

const (
	Male   Sex = core.Male
	Female Sex = core.Female
)

const (
	Yang YinYang = core.Yang
	Yin  YinYang = core.Yin
)

const (
	Wood  WuXing = core.Wood
	Fire  WuXing = core.Fire
	Earth WuXing = core.Earth
	Metal WuXing = core.Metal
	Water WuXing = core.Water
)

const (
	BenGong GuaPosition = core.BenGong
	YiShi   GuaPosition = core.YiShi
	ErShi   GuaPosition = core.ErShi
	SanShi  GuaPosition = core.SanShi
	SiShi   GuaPosition = core.SiShi
	WuShi   GuaPosition = core.WuShi
	YouHun  GuaPosition = core.YouHun
	GuiHun  GuaPosition = core.GuiHun
)

const (
	LQFuMu    LiuQin = core.LQFuMu
	LQXiongDi LiuQin = core.LQXiongDi
	LQQiCai   LiuQin = core.LQQiCai
	LQZiSun   LiuQin = core.LQZiSun
	LQGuanGui LiuQin = core.LQGuanGui
)

const (
	CoinYinYinYin    CoinResult = core.CoinYinYinYin
	CoinYinYinYang   CoinResult = core.CoinYinYinYang
	CoinYinYangYang  CoinResult = core.CoinYinYangYang
	CoinYangYangYang CoinResult = core.CoinYangYangYang
)

const (
	LangZH Language = i18n.LangZH
	LangEN Language = i18n.LangEN
)

const (
	FenXiShiYe    FenXiCategory = core.FenXiShiYe
	FenXiAiQing   FenXiCategory = core.FenXiAiQing
	FenXiCaiYun   FenXiCategory = core.FenXiCaiYun
	FenXiKaoShi   FenXiCategory = core.FenXiKaoShi
	FenXiJianKang FenXiCategory = core.FenXiJianKang
	FenXiChuXing  FenXiCategory = core.FenXiChuXing
	FenXiGuanSi   FenXiCategory = core.FenXiGuanSi
	FenXiJiaZhai  FenXiCategory = core.FenXiJiaZhai
)

const (
	JieDuShiYe    JieDuCategory = core.JieDuShiYe
	JieDuAiQing   JieDuCategory = core.JieDuAiQing
	JieDuCaiYun   JieDuCategory = core.JieDuCaiYun
	JieDuKaoShi   JieDuCategory = core.JieDuKaoShi
	JieDuJianKang JieDuCategory = core.JieDuJianKang
	JieDuChuXing  JieDuCategory = core.JieDuChuXing
	JieDuGuanSi   JieDuCategory = core.JieDuGuanSi
	JieDuJiaZhai  JieDuCategory = core.JieDuJiaZhai
)

var (
	ErrInvalidGuaIndex = core.ErrInvalidGuaIndex
	ErrInvalidYaoIndex = core.ErrInvalidYaoIndex
	ErrInvalidDayanIdx = core.ErrInvalidDayanIdx
	ErrGuaNotFound     = core.ErrGuaNotFound
	ErrInvalidDayanNumber = core.ErrInvalidDayanNumber
	ErrInvalidJiaZiIndex  = core.ErrInvalidJiaZiIndex
	ErrInvalidTianGan     = core.ErrInvalidTianGan
	ErrInvalidDiZhi       = core.ErrInvalidDiZhi
	ErrInvalidJiaZiCombo  = core.ErrInvalidJiaZiCombo
)

func Divine(shang, xia Bagua, bianYao ...int) *ZhouYi {
	return gua.Divine(shang, xia, bianYao...)
}

func DivineByNumber(shang, xia int, bianYao ...int) *ZhouYi {
	return gua.DivineByNumber(shang, xia, bianYao...)
}

func DivineByTime(year, month, day, hour int, seed ...string) *ZhouYi {
	return gua.DivineByTime(year, month, day, hour, seed...)
}

func DivineByTimeGua(params TimeGuaParams, seed ...string) *ZhouYi {
	return qigua.DivineByTimeGua(params, seed...)
}

func DivineByCurrentTime(seed ...string) *ZhouYi {
	return qigua.DivineByCurrentTime(seed...)
}

func DivineByCoins(seed int64) (*ZhouYi, [6]CoinResult) {
	return qigua.DivineByCoins(seed)
}

func DivineByDayan(seed int64) (*ZhouYi, [6]DayanResult) {
	return qigua.DivineByDayan(seed)
}

func DivineByMeihua(upperNum, lowerNum, dongYao int) *ZhouYi {
	return qigua.DivineByMeihua(upperNum, lowerNum, dongYao)
}

func DivineByMeihuaTime(t time.Time, seeds ...string) (*ZhouYi, int, int, int, int) {
	return qigua.DivineByMeihuaTime(t, seeds...)
}

func DivineByDailyHexagram(year, month, day int, personalSeed string) *ZhouYi {
	return qigua.DivineByDailyHexagram(year, month, day, personalSeed)
}

func DivineByLunarTime(lunarYear, lunarMonth, lunarDay, shichenNum int) *ZhouYi {
	return qigua.DivineByLunarTime(lunarYear, lunarMonth, lunarDay, shichenNum)
}

func DivineByBirthday(params BirthdayParams) *ZhouYi {
	return qigua.DivineByBirthday(params)
}

func JieGua(zy *ZhouYi, sex Sex) *JieGuaResult {
	return jiegua.JieGua(zy, sex)
}

func JieGuaWithLang(zy *ZhouYi, sex Sex, lang Language) *JieGuaResult {
	return jiegua.JieGuaWithLang(zy, sex, lang)
}

func FormatJieGua(result *JieGuaResult) string {
	return jiegua.FormatJieGua(result)
}

func FormatJieGuaWithLang(result *JieGuaResult, lang Language) string {
	return jiegua.FormatJieGuaWithLang(result, lang)
}

func GetDayan(number int) (*Dayan, error) {
	return numerology.GetDayan(number)
}

func MustGetDayan(number int) Dayan {
	return numerology.MustGetDayan(number)
}

func GetWenYan(g *Gua) *WenYanData {
	return wenyan.GetWenYan(g)
}

func HasWenYan(g *Gua) bool {
	return wenyan.HasWenYan(g)
}

func GetWuXingByBagua(bagua Bagua) WuXing {
	return wuxing.GetWuXingByBagua(bagua)
}

func TranslateJiXiong(jx string, lang Language) string {
	return i18n.TranslateJiXiong(jx, lang)
}

func TranslateCategory(cat string, lang Language) string {
	return i18n.TranslateCategory(cat, lang)
}

func TranslateYaoPos(pos string, lang Language) string {
	return i18n.TranslateYaoPos(pos, lang)
}

func TranslateBaguaName(name string, lang Language) string {
	return i18n.TranslateBaguaName(name, lang)
}

func TranslateWuXing(name string, lang Language) string {
	return i18n.TranslateWuXing(name, lang)
}

func GetGuaByIndex(index string) (*Gua, error) {
	return gua.GetGuaByIndex(index)
}

func GetGuaByXu(xu int) (*Gua, error) {
	return gua.GetGuaByXu(xu)
}

func GetShiYing(g *Gua) *ShiYingInfo {
	return shiying.GetShiYing(g)
}

func GetGuaGong(g *Gua) Bagua {
	return shiying.GetGuaGong(g)
}

func GetGuaPosition(g *Gua) GuaPosition {
	return shiying.GetGuaPosition(g)
}

func IsJi(zy *ZhouYi, sex Sex) bool {
	return gua.IsJi(zy, sex)
}

func FilterYao(zy *ZhouYi, sex Sex, filters ...string) bool {
	return gua.FilterYao(zy, sex, filters...)
}
