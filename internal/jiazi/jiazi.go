package jiazi

import "github.com/godcong/yi/core"

func GetJiaZiList() []core.JiaZiInfo {
	result := make([]core.JiaZiInfo, len(core.JiaZiList))
	copy(result, core.JiaZiList[:])
	return result
}

func GetJiaZiByIndex(index int) (*core.JiaZiInfo, error) {
	if index < 1 || index > 60 {
		return nil, core.ErrInvalidJiaZiIndex
	}
	return &core.JiaZiList[index-1], nil
}

func GetJiaZiByGanZhi(gan core.TianGan, zhi core.DiZhi) (*core.JiaZiInfo, error) {
	if int(gan) < 0 || int(gan) >= int(core.TianGanCount) {
		return nil, core.ErrInvalidTianGan
	}
	if int(zhi) < 0 || int(zhi) >= int(core.DiZhiCount) {
		return nil, core.ErrInvalidDiZhi
	}
	if int(gan)%2 != int(zhi)%2 {
		return nil, core.ErrInvalidJiaZiCombo
	}
	idx := jiaZiComboIndex(gan, zhi)
	return &core.JiaZiList[idx], nil
}

func GetJiaZiWuXing(info *core.JiaZiInfo) core.WuXing {
	return info.GetNaYinWuXing()
}

func GetJiaZiNaYinWuXing(info *core.JiaZiInfo) core.WuXing {
	return info.GetNaYinWuXing()
}

func jiaZiComboIndex(gan core.TianGan, zhi core.DiZhi) int {
	for i := 0; i < 60; i++ {
		if i%10 == int(gan) && i%12 == int(zhi) {
			return i
		}
	}
	return 0
}
