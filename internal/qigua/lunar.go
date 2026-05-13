package qigua

import (
	"yi/core"
	"yi/internal/gua"
)

func DivineByLunarTime(lunarYear, lunarMonth, lunarDay, shichenNum int) *core.ZhouYi {
	sum := lunarYear + lunarMonth + lunarDay
	shang := gua.ShaoYongToBagua(sum)
	lower := gua.ShaoYongToBagua(sum + shichenNum)
	bian := (sum + shichenNum) % 6
	return gua.DivineByNumber(int(shang), int(lower), bian)
}
