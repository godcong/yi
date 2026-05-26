package core

var GuaStore = map[string]*Gua{}

var BaguaNames [8]string

var BaguaSymbols [8]string

var DayanList [81]Dayan

var WenYanStore map[string][]WenYanEntry

var JiaZiList [60]JiaZiInfo

var GuaPositionStore map[string]GuaPosition

var GuaGongStore map[string]Bagua

var GuaJieDuStore map[string]*GuaJieDu

var ShiYingTable [GuaPositionCount]shiYingEntry

var WuXingMapping map[Bagua]string

var WuXingAttributes map[string]WuXingInfo

var TianGanNames [TianGanCount]string

var DiZhiNames [DiZhiCount]string

var AllJieDuCategories = []JieDuCategory{
	JieDuShiYe, JieDuAiQing, JieDuCaiYun, JieDuKaoShi,
	JieDuJianKang, JieDuChuXing, JieDuGuanSi, JieDuJiaZhai,
}
