package liuqin

import (
	"github.com/godcong/yi/core"
	"github.com/godcong/yi/internal/wuxing"
)

func GetLiuQin(guaGongWX, yaoWX core.WuXing) core.LiuQin {
	if guaGongWX == yaoWX {
		return core.LQXiongDi
	}
	if wuxing.Sheng(guaGongWX) == yaoWX {
		return core.LQZiSun
	}
	if wuxing.Ke(guaGongWX) == yaoWX {
		return core.LQQiCai
	}
	if wuxing.Sheng(yaoWX) == guaGongWX {
		return core.LQFuMu
	}
	if wuxing.Ke(yaoWX) == guaGongWX {
		return core.LQGuanGui
	}
	return core.LQCount
}

func GetGuaGongWuXing(bagua core.Bagua) core.WuXing {
	return wuxing.GetWuXingByBagua(bagua)
}

func GetLiuQinForGua(guaGongWX core.WuXing) map[core.WuXing]core.LiuQin {
	result := make(map[core.WuXing]core.LiuQin, 5)
	for _, wx := range []core.WuXing{core.Wood, core.Fire, core.Earth, core.Metal, core.Water} {
		result[wx] = GetLiuQin(guaGongWX, wx)
	}
	return result
}
