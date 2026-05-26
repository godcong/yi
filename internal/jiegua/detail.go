package jiegua

import "github.com/godcong/yi/core"

func GetGuaJieDuByIndex(index string) *core.GuaJieDu {
	if core.GuaJieDuStore == nil {
		return nil
	}
	return core.GuaJieDuStore[index]
}

func GetGuaJieDuByCategory(index string, category core.JieDuCategory) string {
	jd := GetGuaJieDuByIndex(index)
	if jd == nil {
		return ""
	}
	switch category {
	case core.JieDuShiYe:
		return jd.ShiYe
	case core.JieDuAiQing:
		return jd.AiQing
	case core.JieDuCaiYun:
		return jd.CaiYun
	case core.JieDuKaoShi:
		return jd.KaoShi
	case core.JieDuJianKang:
		return jd.JianKang
	case core.JieDuChuXing:
		return jd.ChuXing
	case core.JieDuGuanSi:
		return jd.GuanSi
	case core.JieDuJiaZhai:
		return jd.JiaZhai
	default:
		return ""
	}
}

func GetWuXingInfo(bagua core.Bagua) *core.WuXingInfo {
	wx, ok := core.WuXingMapping[bagua]
	if !ok {
		wx = "土"
	}
	info, ok := core.WuXingAttributes[wx]
	if !ok {
		return nil
	}
	return &info
}
