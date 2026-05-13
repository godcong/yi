package shiying

import "yi/core"

func GetShiYing(g *core.Gua) *core.ShiYingInfo {
	pos := GetGuaPosition(g)
	if pos < 0 || pos >= core.GuaPositionCount {
		return nil
	}
	info := core.ShiYingTable[pos]
	return &core.ShiYingInfo{
		ShiPos:   info.ShiPos,
		YingPos:  info.YingPos,
		Position: pos,
	}
}

func GetGuaPosition(g *core.Gua) core.GuaPosition {
	if idx, ok := core.GuaPositionStore[g.Index]; ok {
		return idx
	}
	return core.GuaPositionCount
}

func GetGuaGong(g *core.Gua) core.Bagua {
	if bg, ok := core.GuaGongStore[g.Index]; ok {
		return bg
	}
	return -1
}
