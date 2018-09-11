package core

import "testing"

func TestSetGuaXiang(t *testing.T) {
	gx := make(map[string]*GuaXiang)

	gx["乾乾"] = &GuaXiang{
		ShangGua: "乾",
		ShangShu: 0x07,
		XiaGua:   "乾",
		XiaShu:   0x07,
		GuaXiang: "1",
		GuaMing:  "1",
		GuaYi:    "1",
		GuaYun:   "1",
		XiangYue: "1",
		FuHao:    "1",
	}

	setGuaXiang(gx)

}

func TestGetGuaXiang(t *testing.T) {
	t.Log(getGuaXiang())
}
