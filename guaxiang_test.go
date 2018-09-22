package core

import "testing"

func qian(gx map[string]*GuaXiang) map[string]*GuaXiang {
	gx["乾坤"] = &GuaXiang{
		ShangGua: "乾为天",
		ShangShu: 0x00,
		XiaGua:   "坤为地",
		XiaShu:   0x07,
		GuaMing:  "天地否",
	}

	gx["乾艮"] = &GuaXiang{
		ShangGua: "乾为天",
		ShangShu: 0x00,
		XiaGua:   "艮为山",
		XiaShu:   0x06,
		GuaMing:  "天山遁",
	}
	gx["乾坎"] = &GuaXiang{
		ShangGua: "乾为天",
		ShangShu: 0x00,
		XiaGua:   "坎为水",
		XiaShu:   0x05,
		GuaMing:  "天水讼",
	}

	gx["乾巽"] = &GuaXiang{
		ShangGua: "乾为天",
		ShangShu: 0x00,
		XiaGua:   "巽为风",
		XiaShu:   0x04,
		GuaMing:  "天风姤",
	}

	gx["乾震"] = &GuaXiang{
		ShangGua: "乾为天",
		ShangShu: 0x00,
		XiaGua:   "震为雷",
		XiaShu:   0x03,
		GuaMing:  "天雷无妄",
	}

	gx["乾离"] = &GuaXiang{
		ShangGua: "乾为天",
		ShangShu: 0x00,
		XiaGua:   "离为火",
		XiaShu:   0x02,
		GuaMing:  "天火同人",
	}

	gx["乾兑"] = &GuaXiang{
		ShangGua: "乾为天",
		ShangShu: 0x00,
		XiaGua:   "兑为泽",
		XiaShu:   0x01,
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
		ShangGua: "坤为地",
		ShangShu: 0x07,
		XiaGua:   "坤为地",
		XiaShu:   0x07,
		GuaMing:  "坤为地",
	}

	gx["坤艮"] = &GuaXiang{
		ShangGua: "坤为地",
		ShangShu: 0x07,
		XiaGua:   "艮为山",
		XiaShu:   0x06,
		GuaMing:  "地山谦",
	}

	gx["坤坎"] = &GuaXiang{
		ShangGua: "坤为地",
		ShangShu: 0x07,
		XiaGua:   "坎为水",
		XiaShu:   0x05,
		GuaMing:  "地水师",
	}
	gx["坤巽"] = &GuaXiang{
		ShangGua: "坤为地",
		ShangShu: 0x07,
		XiaGua:   "巽为风",
		XiaShu:   0x04,
		GuaMing:  "地风升",
	}
	gx["坤震"] = &GuaXiang{
		ShangGua: "坤为地",
		ShangShu: 0x07,
		XiaGua:   "震为雷",
		XiaShu:   0x03,
		GuaMing:  "地雷复",
	}

	gx["坤离"] = &GuaXiang{
		ShangGua: "坤为地",
		ShangShu: 0x07,
		XiaGua:   "离为火",
		XiaShu:   0x02,
		GuaMing:  "地火明夷",
	}

	gx["坤兑"] = &GuaXiang{
		ShangGua: "坤为地",
		ShangShu: 0x07,
		XiaGua:   "兑为泽",
		XiaShu:   0x01,
		GuaMing:  "地泽临",
	}
	gx["坤乾"] = &GuaXiang{
		ShangGua: "坤为地",
		ShangShu: 0x07,
		XiaGua:   "乾为天",
		XiaShu:   0x00,
		GuaMing:  "地天泰",
	}
	return gx
}

func gen(gx map[string]*GuaXiang) map[string]*GuaXiang {
	gx["艮坤"] = &GuaXiang{
		ShangGua: "艮为山",
		ShangShu: 0x06,
		XiaGua:   "坤为地",
		XiaShu:   0x07,
		GuaMing:  "山地剥",
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
	gx["坎乾"] = &GuaXiang{
		ShangGua: "巽为风",
		ShangShu: 0x04,
		XiaGua:   "乾为天",
		XiaShu:   0x00,
		GuaMing:  "水天需",
	}
	return gx
}

func TestSetGuaXiang(t *testing.T) {
	gx := make(map[string]*GuaXiang)

	setGuaXiang(gx)

}

func TestGetGuaXiang(t *testing.T) {
	t.Log(getGuaXiang())
}
