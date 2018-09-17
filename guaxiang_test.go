package core

import "testing"

func TestSetGuaXiang(t *testing.T) {
	gx := make(map[string]*GuaXiang)

	gx["乾乾"] = &GuaXiang{
		ShangGua: "乾为天",
		ShangShu: 0x07,
		XiaGua:   "乾为天",
		XiaShu:   0x07,
		GuaXiang: "乾",
		GuaMing:  "乾为天",
		GuaYi:    "元亨，利贞",
		GuaYun:   "大哉乾元，万物资始，乃统天。云行雨施，品物流形。大明始终，六位时成，时乘六龙以御天。乾道变化，各正性命，保合大和，乃利贞。首出庶物，万国咸宁。",
		XiangYue: "天行健，君子以自强不息",
		FuHao:    "䷀",
		ChuYao:   "潜龙勿用，阳在下也。",
		ErYao:    "见龙在田，德施普也。",
		SanYao:   "终日乾乾，反复道也。",
		SiYao:    "或跃在渊，进无咎也。",
		WuYao:    "飞龙在天，大人造也。",
		ShangYao: "亢龙有悔，盈不可久也。",
		Yong:     "天德不可为首也。",
	}

	setGuaXiang(gx)

}

func TestGetGuaXiang(t *testing.T) {
	t.Log(getGuaXiang())
}
