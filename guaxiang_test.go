package yi

import "testing"

func qian(gx map[string]*GuaXiang) map[string]*GuaXiang {
	gx["乾坤"] = &GuaXiang{
		ShangGua: "乾为天",
		ShangShu: 0x00,
		XiaGua:   "坤为地",
		XiaShu:   0x07,
		GuaXiang: "否",
		GuaMing:  "天地否",
	}

	gx["乾艮"] = &GuaXiang{
		ShangGua: "乾为天",
		ShangShu: 0x00,
		XiaGua:   "艮为山",
		XiaShu:   0x06,
		GuaXiang: "遁",
		GuaMing:  "天山遁",
	}
	gx["乾坎"] = &GuaXiang{
		ShangGua: "乾为天",
		ShangShu: 0x00,
		XiaGua:   "坎为水",
		XiaShu:   0x05,
		GuaXiang: "讼",
		GuaMing:  "天水讼",
	}

	gx["乾巽"] = &GuaXiang{
		ShangGua:        "乾为天",
		ShangShu:        0x00,
		XiaGua:          "巽为风",
		XiaShu:          0x04,
		JiXiong:         "",
		GuaXiang:        "姤",
		GuaMing:         "天风姤",
		GuaYi:           "",
		GuaYun:          "",
		XiangYue:        "",
		FuHao:           "",
		ChuYao:          "系于金柅，柔道牵也。",
		ChuYaoJiXiong:   "平",
		ErYao:           "包有鱼，义不及宾也。",
		ErYaoJiXiong:    "平",
		SanYao:          "其行次且，行未牵也。",
		SanYaoJiXiong:   "凶",
		SiYao:           "无鱼之凶，远民也。",
		SiYaoJiXiong:    "凶",
		WuYao:           "九五含章，中正也；有陨自天，志不舍命也。",
		WuYaoJiXiong:    "平",
		ShangYao:        "姤其角，上穷吝也。",
		ShangYaoJiXiong: "凶",
		Yong:            "",
		YongJiXiong:     "",
	}

	gx["乾震"] = &GuaXiang{
		ShangGua: "乾为天",
		ShangShu: 0x00,
		XiaGua:   "震为雷",
		XiaShu:   0x03,
		GuaXiang: "无妄",
		GuaMing:  "天雷无妄",
	}

	gx["乾离"] = &GuaXiang{
		ShangGua:        "乾为天",
		ShangShu:        0x00,
		XiaGua:          "离为火",
		XiaShu:          0x02,
		JiXiong:         "",
		GuaXiang:        "同人",
		GuaMing:         "天火同人",
		GuaYi:           "",
		GuaYun:          "",
		XiangYue:        "",
		FuHao:           "",
		ChuYao:          "出门同人，又谁咎也。",
		ChuYaoJiXiong:   "平",
		ErYao:           "同人于宗，吝道也。",
		ErYaoJiXiong:    "凶",
		SanYao:          "伏戎于莽，敌刚也，三岁不兴，安行也。",
		SanYaoJiXiong:   "凶",
		SiYao:           "乘其墉，义弗克也，其吉，则困而反则也。",
		SiYaoJiXiong:    "平",
		WuYao:           "同人之先，以中直也。大师相遇，言相克也。",
		WuYaoJiXiong:    "平",
		ShangYao:        "同人于郊，志未得也。",
		ShangYaoJiXiong: "平",
		Yong:            "",
		YongJiXiong:     "",
	}

	gx["乾兑"] = &GuaXiang{
		ShangGua: "乾为天",
		ShangShu: 0x00,
		XiaGua:   "兑为泽",
		XiaShu:   0x01,
		GuaXiang: "履",
		GuaMing:  "天泽履",
	}

	gx["乾乾"] = &GuaXiang{
		ShangGua:        "乾为天",
		ShangShu:        0x00,
		XiaGua:          "乾为天",
		XiaShu:          0x00,
		JiXiong:         "吉",
		GuaXiang:        "乾",
		GuaMing:         "乾为天",
		GuaYi:           "元亨，利贞",
		GuaYun:          "大哉乾元，万物资始，乃统天。云行雨施，品物流形。大明始终，六位时成，时乘六龙以御天。乾道变化，各正性命，保合大和，乃利贞。首出庶物，万国咸宁。",
		XiangYue:        "天行健，君子以自强不息",
		FuHao:           "䷀",
		ChuYao:          "潜龙勿用，阳在下也。",
		ChuYaoJiXiong:   "平",
		ErYao:           "见龙在田，德施普也。",
		ErYaoJiXiong:    "吉",
		SanYao:          "终日乾乾，反复道也。",
		SanYaoJiXiong:   "平",
		SiYao:           "或跃在渊，进无咎也。",
		SiYaoJiXiong:    "平",
		WuYao:           "飞龙在天，大人造也。",
		WuYaoJiXiong:    "吉",
		ShangYao:        "亢龙有悔，盈不可久也。",
		ShangYaoJiXiong: "平",
		Yong:            "天德不可为首也。",
		YongJiXiong:     "-",
	}

	return gx
}

func kun(gx map[string]*GuaXiang) map[string]*GuaXiang {

	gx["坤坤"] = &GuaXiang{
		ShangGua:        "坤为地",
		ShangShu:        0x07,
		XiaGua:          "坤为地",
		XiaShu:          0x07,
		JiXiong:         "",
		GuaXiang:        "坤",
		GuaMing:         "坤为地",
		GuaYi:           "",
		GuaYun:          "",
		XiangYue:        "",
		FuHao:           "",
		ChuYao:          "履霜坚冰，阴始凝也。驯致其道，至坚冰也。",
		ChuYaoJiXiong:   "平",
		ErYao:           "六二之动，直以方也；不习无不利，地道光也。",
		ErYaoJiXiong:    "吉",
		SanYao:          "含章可贞，以时发也；或从王事，知光大也。",
		SanYaoJiXiong:   "平",
		SiYao:           "括囊无咎，慎不害也。",
		SiYaoJiXiong:    "平",
		WuYao:           "黄裳元吉，文在中也。",
		WuYaoJiXiong:    "吉",
		ShangYao:        "龙战于野，其道穷也。",
		ShangYaoJiXiong: "凶",
		Yong:            "",
		YongJiXiong:     "",
	}

	gx["坤艮"] = &GuaXiang{
		ShangGua:        "坤为地",
		ShangShu:        0x07,
		XiaGua:          "艮为山",
		XiaShu:          0x06,
		JiXiong:         "",
		GuaXiang:        "谦",
		GuaMing:         "地山谦",
		GuaYi:           "",
		GuaYun:          "",
		XiangYue:        "",
		FuHao:           "",
		ChuYao:          "",
		ChuYaoJiXiong:   "",
		ErYao:           "",
		ErYaoJiXiong:    "",
		SanYao:          "",
		SanYaoJiXiong:   "",
		SiYao:           "",
		SiYaoJiXiong:    "",
		WuYao:           "",
		WuYaoJiXiong:    "",
		ShangYao:        "",
		ShangYaoJiXiong: "",
		Yong:            "",
		YongJiXiong:     "",
	}

	gx["坤坎"] = &GuaXiang{
		ShangGua: "坤为地",
		ShangShu: 0x07,
		XiaGua:   "坎为水",
		XiaShu:   0x05,
		GuaXiang: "师",
		GuaMing:  "地水师",
	}

	gx["坤巽"] = &GuaXiang{
		ShangGua: "坤为地",
		ShangShu: 0x07,
		XiaGua:   "巽为风",
		XiaShu:   0x04,
		GuaXiang: "升",
		GuaMing:  "地风升",
	}

	gx["坤震"] = &GuaXiang{
		ShangGua:        "坤为地",
		ShangShu:        0x07,
		XiaGua:          "震为雷",
		XiaShu:          0x03,
		JiXiong:         "",
		GuaXiang:        "复",
		GuaMing:         "地雷复",
		GuaYi:           "",
		GuaYun:          "",
		XiangYue:        "",
		FuHao:           "",
		ChuYao:          "不远之复，以修身也。",
		ChuYaoJiXiong:   "吉",
		ErYao:           "休复之吉，以下仁也。",
		ErYaoJiXiong:    "吉",
		SanYao:          "频复之厉，义无咎也。",
		SanYaoJiXiong:   "平",
		SiYao:           "中行独复，以从道也。",
		SiYaoJiXiong:    "平",
		WuYao:           "敦复无悔，中以自考也。",
		WuYaoJiXiong:    "吉",
		ShangYao:        "迷复之凶，反君道也。",
		ShangYaoJiXiong: "凶",
		Yong:            "",
		YongJiXiong:     "",
	}

	gx["坤离"] = &GuaXiang{
		ShangGua: "坤为地",
		ShangShu: 0x07,
		XiaGua:   "离为火",
		XiaShu:   0x02,
		GuaXiang: "明夷",
		GuaMing:  "地火明夷",
	}

	gx["坤兑"] = &GuaXiang{
		ShangGua: "坤为地",
		ShangShu: 0x07,
		XiaGua:   "兑为泽",
		XiaShu:   0x01,
		GuaXiang: "临",
		GuaMing:  "地泽临",
	}

	gx["坤乾"] = &GuaXiang{
		ShangGua: "坤为地",
		ShangShu: 0x07,
		XiaGua:   "乾为天",
		XiaShu:   0x00,
		GuaXiang: "泰",
		GuaMing:  "地天泰",
	}
	return gx
}

func gen(gx map[string]*GuaXiang) map[string]*GuaXiang {
	gx["艮坤"] = &GuaXiang{
		ShangGua:        "艮为山",
		ShangShu:        0x06,
		XiaGua:          "坤为地",
		XiaShu:          0x07,
		JiXiong:         "",
		GuaXiang:        "",
		GuaMing:         "山地剥",
		GuaYi:           "",
		GuaYun:          "",
		XiangYue:        "",
		FuHao:           "",
		ChuYao:          "剥床以足，以灭下也。",
		ChuYaoJiXiong:   "凶",
		ErYao:           "剥床以辨，未有与也。",
		ErYaoJiXiong:    "凶",
		SanYao:          "剥之无咎，失上下也。",
		SanYaoJiXiong:   "平",
		SiYao:           "剥床以肤，切近灾也。",
		SiYaoJiXiong:    "凶",
		WuYao:           "以宫人宠，终无尤也。",
		WuYaoJiXiong:    "吉",
		ShangYao:        "君子得舆，民所载也；小人剥庐，终不可用也。",
		ShangYaoJiXiong: "平",
		Yong:            "",
		YongJiXiong:     "",
	}

	gx["艮艮"] = &GuaXiang{
		ShangGua: "艮为山",
		ShangShu: 0x06,
		XiaGua:   "艮为山",
		XiaShu:   0x06,
		GuaMing:  "艮为山",
	}

	gx["艮坎"] = &GuaXiang{
		ShangGua: "艮为山",
		ShangShu: 0x06,
		XiaGua:   "坎为水",
		XiaShu:   0x05,
		GuaMing:  "山水蒙",
	}
	gx["艮巽"] = &GuaXiang{
		ShangGua: "艮为山",
		ShangShu: 0x06,
		XiaGua:   "巽为风",
		XiaShu:   0x04,
		GuaMing:  "山风蛊",
	}
	gx["艮震"] = &GuaXiang{
		ShangGua: "艮为山",
		ShangShu: 0x06,
		XiaGua:   "震为雷",
		XiaShu:   0x03,
		GuaMing:  "山雷颐",
	}

	gx["艮离"] = &GuaXiang{
		ShangGua: "艮为山",
		ShangShu: 0x06,
		XiaGua:   "离为火",
		XiaShu:   0x02,
		GuaMing:  "山火贲",
	}

	gx["艮兑"] = &GuaXiang{
		ShangGua: "艮为山",
		ShangShu: 0x06,
		XiaGua:   "兑为泽",
		XiaShu:   0x01,
		GuaMing:  "山泽损",
	}
	gx["艮乾"] = &GuaXiang{
		ShangGua: "艮为山",
		ShangShu: 0x06,
		XiaGua:   "乾为天",
		XiaShu:   0x00,
		GuaMing:  "山天大畜",
	}

	gx["坎坤"] = &GuaXiang{
		ShangGua: "坎为水",
		ShangShu: 0x05,
		XiaGua:   "坤为地",
		XiaShu:   0x07,
		GuaMing:  "水地比",
	}
	return gx
}

func kan(gx map[string]*GuaXiang) map[string]*GuaXiang {

	gx["坎艮"] = &GuaXiang{
		ShangGua: "坎为水",
		ShangShu: 0x05,
		XiaGua:   "艮为山",
		XiaShu:   0x06,
		GuaMing:  "水山蹇",
	}

	gx["坎坎"] = &GuaXiang{
		ShangGua: "坎为水",
		ShangShu: 0x05,
		XiaGua:   "坎为水",
		XiaShu:   0x05,
		GuaMing:  "坎为水",
	}
	gx["坎巽"] = &GuaXiang{
		ShangGua: "坎为水",
		ShangShu: 0x05,
		XiaGua:   "巽为风",
		XiaShu:   0x04,
		GuaMing:  "水风井",
	}
	gx["坎震"] = &GuaXiang{
		ShangGua: "坎为水",
		ShangShu: 0x05,
		XiaGua:   "震为雷",
		XiaShu:   0x03,
		GuaMing:  "水雷屯",
	}

	gx["坎离"] = &GuaXiang{
		ShangGua: "坎为水",
		ShangShu: 0x05,
		XiaGua:   "离为火",
		XiaShu:   0x02,
		GuaMing:  "水火既济",
	}

	gx["坎兑"] = &GuaXiang{
		ShangGua: "坎为水",
		ShangShu: 0x05,
		XiaGua:   "兑为泽",
		XiaShu:   0x01,
		GuaMing:  "水泽节",
	}
	gx["坎乾"] = &GuaXiang{
		ShangGua: "坎为水",
		ShangShu: 0x05,
		XiaGua:   "乾为天",
		XiaShu:   0x00,
		GuaMing:  "水天需",
	}

	gx["坎坤"] = &GuaXiang{
		ShangGua: "坎为水",
		ShangShu: 0x05,
		XiaGua:   "坤为地",
		XiaShu:   0x07,
		GuaMing:  "水地比",
	}

	gx["坎艮"] = &GuaXiang{
		ShangGua: "坎为水",
		ShangShu: 0x05,
		XiaGua:   "艮为山",
		XiaShu:   0x06,
		GuaMing:  "水山蹇",
	}

	gx["坎坎"] = &GuaXiang{
		ShangGua: "坎为水",
		ShangShu: 0x05,
		XiaGua:   "坎为水",
		XiaShu:   0x05,
		GuaMing:  "坎为水",
	}
	gx["坎巽"] = &GuaXiang{
		ShangGua: "坎为水",
		ShangShu: 0x05,
		XiaGua:   "巽为风",
		XiaShu:   0x04,
		GuaMing:  "水风井",
	}
	gx["坎震"] = &GuaXiang{
		ShangGua: "坎为水",
		ShangShu: 0x05,
		XiaGua:   "震为雷",
		XiaShu:   0x03,
		GuaMing:  "水雷屯",
	}

	gx["坎离"] = &GuaXiang{
		ShangGua: "坎为水",
		ShangShu: 0x05,
		XiaGua:   "离为火",
		XiaShu:   0x02,
		GuaMing:  "水火既济",
	}

	gx["坎兑"] = &GuaXiang{
		ShangGua: "坎为水",
		ShangShu: 0x05,
		XiaGua:   "兑为泽",
		XiaShu:   0x01,
		GuaMing:  "水泽节",
	}
	gx["坎乾"] = &GuaXiang{
		ShangGua: "坎为水",
		ShangShu: 0x05,
		XiaGua:   "乾为天",
		XiaShu:   0x00,
		GuaMing:  "水天需",
	}
	return gx
}

func xun(gx map[string]*GuaXiang) map[string]*GuaXiang {

	gx["巽坤"] = &GuaXiang{
		ShangGua: "巽为风",
		ShangShu: 0x04,
		XiaGua:   "坤为地",
		XiaShu:   0x07,
		GuaMing:  "风地观",
	}

	gx["巽艮"] = &GuaXiang{
		ShangGua: "巽为风",
		ShangShu: 0x04,
		XiaGua:   "艮为山",
		XiaShu:   0x06,
		GuaMing:  "风山渐",
	}

	gx["巽坎"] = &GuaXiang{
		ShangGua: "巽为风",
		ShangShu: 0x04,
		XiaGua:   "坎为水",
		XiaShu:   0x05,
		GuaMing:  "风水涣",
	}
	gx["巽巽"] = &GuaXiang{
		ShangGua: "巽为风",
		ShangShu: 0x04,
		XiaGua:   "巽为风",
		XiaShu:   0x04,
		GuaMing:  "巽为风",
	}
	gx["巽震"] = &GuaXiang{
		ShangGua: "巽为风",
		ShangShu: 0x04,
		XiaGua:   "震为雷",
		XiaShu:   0x03,
		GuaMing:  "风雷益",
	}

	gx["巽离"] = &GuaXiang{
		ShangGua: "巽为风",
		ShangShu: 0x04,
		XiaGua:   "离为火",
		XiaShu:   0x02,
		GuaMing:  "风火家人",
	}

	gx["巽兑"] = &GuaXiang{
		ShangGua: "巽为风",
		ShangShu: 0x04,
		XiaGua:   "兑为泽",
		XiaShu:   0x01,
		GuaMing:  "风泽中孚",
	}
	gx["巽乾"] = &GuaXiang{
		ShangGua: "巽为风",
		ShangShu: 0x04,
		XiaGua:   "乾为天",
		XiaShu:   0x00,
		GuaMing:  "风天小畜",
	}
	return gx
}

func zhen(gx map[string]*GuaXiang) map[string]*GuaXiang {

	gx["震坤"] = &GuaXiang{
		ShangGua: "震为雷",
		ShangShu: 0x03,
		XiaGua:   "坤为地",
		XiaShu:   0x07,
		GuaMing:  "雷地豫",
	}

	gx["震艮"] = &GuaXiang{
		ShangGua: "震为雷",
		ShangShu: 0x03,
		XiaGua:   "艮为山",
		XiaShu:   0x06,
		GuaMing:  "雷山小过",
	}

	gx["震坎"] = &GuaXiang{
		ShangGua: "震为雷",
		ShangShu: 0x03,
		XiaGua:   "坎为水",
		XiaShu:   0x05,
		GuaMing:  "雷水解",
	}
	gx["震巽"] = &GuaXiang{
		ShangGua: "震为雷",
		ShangShu: 0x03,
		XiaGua:   "巽为风",
		XiaShu:   0x04,
		GuaMing:  "雷风恒",
	}
	gx["震震"] = &GuaXiang{
		ShangGua: "震为雷",
		ShangShu: 0x03,
		XiaGua:   "震为雷",
		XiaShu:   0x03,
		GuaMing:  "震为雷",
	}

	gx["震离"] = &GuaXiang{
		ShangGua: "震为雷",
		ShangShu: 0x03,
		XiaGua:   "离为火",
		XiaShu:   0x02,
		GuaMing:  "雷火丰",
	}

	gx["震兑"] = &GuaXiang{
		ShangGua: "震为雷",
		ShangShu: 0x03,
		XiaGua:   "兑为泽",
		XiaShu:   0x01,
		GuaMing:  "雷泽归妹",
	}
	gx["震乾"] = &GuaXiang{
		ShangGua: "震为雷",
		ShangShu: 0x03,
		XiaGua:   "乾为天",
		XiaShu:   0x00,
		GuaMing:  "雷天大壮",
	}
	return gx
}

func li(gx map[string]*GuaXiang) map[string]*GuaXiang {

	gx["离坤"] = &GuaXiang{
		ShangGua: "离为火",
		ShangShu: 0x02,
		XiaGua:   "坤为地",
		XiaShu:   0x07,
		GuaMing:  "火地晋",
	}

	gx["离艮"] = &GuaXiang{
		ShangGua: "离为火",
		ShangShu: 0x02,
		XiaGua:   "艮为山",
		XiaShu:   0x06,
		GuaMing:  "火山旅",
	}

	gx["离坎"] = &GuaXiang{
		ShangGua: "离为火",
		ShangShu: 0x02,
		XiaGua:   "坎为水",
		XiaShu:   0x05,
		GuaMing:  "火水未济",
	}

	gx["离巽"] = &GuaXiang{
		ShangGua:        "离为火",
		ShangShu:        0x02,
		XiaGua:          "巽为风",
		XiaShu:          0x04,
		JiXiong:         "",
		GuaXiang:        "",
		GuaMing:         "火风鼎",
		GuaYi:           "",
		GuaYun:          "",
		XiangYue:        "",
		FuHao:           "",
		ChuYao:          "鼎颠趾，未悖也。利出否，以从贵也。",
		ChuYaoJiXiong:   "平",
		ErYao:           "鼎有实，慎所之也。我仇有疾，终无尤也。",
		ErYaoJiXiong:    "吉",
		SanYao:          "鼎耳革，失其义也。",
		SanYaoJiXiong:   "平",
		SiYao:           "覆公餗，信如何也。",
		SiYaoJiXiong:    "凶",
		WuYao:           "鼎黄耳，中以为实也。",
		WuYaoJiXiong:    "吉",
		ShangYao:        "玉铉在上，刚柔节也。",
		ShangYaoJiXiong: "吉 ",
		Yong:            "",
		YongJiXiong:     "",
	}

	gx["离震"] = &GuaXiang{
		ShangGua: "离为火",
		ShangShu: 0x02,
		XiaGua:   "震为雷",
		XiaShu:   0x03,
		GuaMing:  "火雷噬嗑",
	}

	gx["离离"] = &GuaXiang{
		ShangGua: "离为火",
		ShangShu: 0x02,
		XiaGua:   "离为火",
		XiaShu:   0x02,
		GuaMing:  "离为火",
	}

	gx["离兑"] = &GuaXiang{
		ShangGua: "离为火",
		ShangShu: 0x02,
		XiaGua:   "兑为泽",
		XiaShu:   0x01,
		GuaMing:  "火泽睽",
	}

	gx["离乾"] = &GuaXiang{
		ShangGua: "离为火",
		ShangShu: 0x02,
		XiaGua:   "乾为天",
		XiaShu:   0x00,
		GuaMing:  "火天大有",
	}
	return gx
}

func dui(gx map[string]*GuaXiang) map[string]*GuaXiang {
	gx["兑坤"] = &GuaXiang{
		ShangGua: "兑为泽",
		ShangShu: 0x01,
		XiaGua:   "坤为地",
		XiaShu:   0x07,
		GuaMing:  "火地晋",
	}

	gx["兑艮"] = &GuaXiang{
		ShangGua: "兑为泽",
		ShangShu: 0x01,
		XiaGua:   "艮为山",
		XiaShu:   0x06,
		GuaMing:  "泽山咸",
	}

	gx["兑坎"] = &GuaXiang{
		ShangGua: "兑为泽",
		ShangShu: 0x01,
		XiaGua:   "坎为水",
		XiaShu:   0x05,
		GuaMing:  "火水未济",
	}

	gx["兑巽"] = &GuaXiang{
		ShangGua: "兑为泽",
		ShangShu: 0x01,
		XiaGua:   "巽为风",
		XiaShu:   0x04,
		GuaMing:  "火风鼎",
	}

	gx["兑震"] = &GuaXiang{
		ShangGua: "兑为泽",
		ShangShu: 0x01,
		XiaGua:   "震为雷",
		XiaShu:   0x03,
		GuaMing:  "火雷噬嗑",
	}

	gx["兑离"] = &GuaXiang{
		ShangGua:        "兑为泽",
		ShangShu:        0x01,
		XiaGua:          "离为火",
		XiaShu:          0x02,
		JiXiong:         "",
		GuaXiang:        "",
		GuaMing:         "泽火革",
		GuaYi:           "",
		GuaYun:          "",
		XiangYue:        "",
		FuHao:           "",
		ChuYao:          "巩用黄牛，不可以有为也。",
		ChuYaoJiXiong:   "平",
		ErYao:           "己日革之，行有嘉也。",
		ErYaoJiXiong:    "吉",
		SanYao:          "革言三就，又何之矣。",
		SanYaoJiXiong:   "凶",
		SiYao:           "改命之吉，信志也。",
		SiYaoJiXiong:    "平",
		WuYao:           "大人虎变，其文炳也。",
		WuYaoJiXiong:    "吉",
		ShangYao:        "君子豹变，其文蔚也。小人革面，顺以从君也。",
		ShangYaoJiXiong: "平",
		Yong:            "",
		YongJiXiong:     "",
	}

	gx["兑兑"] = &GuaXiang{
		ShangGua: "兑为泽",
		ShangShu: 0x01,
		XiaGua:   "兑为泽",
		XiaShu:   0x01,
		GuaMing:  "火泽睽",
	}

	gx["兑乾"] = &GuaXiang{
		ShangGua:        "兑为泽",
		ShangShu:        0x01,
		XiaGua:          "乾为天",
		XiaShu:          0x00,
		JiXiong:         "",
		GuaXiang:        "",
		GuaMing:         "火天大有",
		GuaYi:           "",
		GuaYun:          "",
		XiangYue:        "",
		FuHao:           "",
		ChuYao:          "大有初九，无交害也。",
		ChuYaoJiXiong:   "平",
		ErYao:           "大车以载，积中不败也。",
		ErYaoJiXiong:    "吉",
		SanYao:          "公用亨于天子，小人害也。",
		SanYaoJiXiong:   "平",
		SiYao:           "匪其彭，无咎，明辩晰也。",
		SiYaoJiXiong:    "平",
		WuYao:           "厥孚交如，信以发志也；威如之吉，易而无备也。",
		WuYaoJiXiong:    "吉",
		ShangYao:        "大有上吉，自天祐也。",
		ShangYaoJiXiong: "吉",
		Yong:            "",
		YongJiXiong:     "",
	}

	return gx
}

func TestSetGuaXiang(t *testing.T) {
	gx := make(map[string]*GuaXiang)
	gx = qian(gx)
	gx = kun(gx)
	gx = gen(gx)
	gx = kan(gx)
	gx = xun(gx)
	gx = zhen(gx)
	gx = li(gx)
	gx = dui(gx)
	setGuaXiang(gx)

}

func TestGetGuaXiang(t *testing.T) {
	if len(getGuaXiang()) != 64 {
		t.Log("not enough", getGuaXiang())
	}

}
