package wuxing

import "github.com/godcong/yi/core"

func WuXingString(wx core.WuXing) string {
	switch wx {
	case core.Wood:
		return "木"
	case core.Fire:
		return "火"
	case core.Earth:
		return "土"
	case core.Metal:
		return "金"
	case core.Water:
		return "水"
	default:
		return ""
	}
}

func Sheng(wx core.WuXing) core.WuXing {
	switch wx {
	case core.Wood:
		return core.Fire
	case core.Fire:
		return core.Earth
	case core.Earth:
		return core.Metal
	case core.Metal:
		return core.Water
	case core.Water:
		return core.Wood
	default:
		return 0
	}
}

func Ke(wx core.WuXing) core.WuXing {
	switch wx {
	case core.Wood:
		return core.Earth
	case core.Earth:
		return core.Water
	case core.Water:
		return core.Fire
	case core.Fire:
		return core.Metal
	case core.Metal:
		return core.Wood
	default:
		return 0
	}
}

func BeiSheng(wx core.WuXing) core.WuXing {
	switch wx {
	case core.Wood:
		return core.Water
	case core.Fire:
		return core.Wood
	case core.Earth:
		return core.Fire
	case core.Metal:
		return core.Earth
	case core.Water:
		return core.Metal
	default:
		return 0
	}
}

func BeiKe(wx core.WuXing) core.WuXing {
	switch wx {
	case core.Wood:
		return core.Metal
	case core.Fire:
		return core.Water
	case core.Earth:
		return core.Wood
	case core.Metal:
		return core.Fire
	case core.Water:
		return core.Earth
	default:
		return 0
	}
}

var ShengMap = map[core.WuXing]core.WuXing{
	core.Wood:  core.Fire,
	core.Fire:  core.Earth,
	core.Earth: core.Metal,
	core.Metal: core.Water,
	core.Water: core.Wood,
}

var KeMap = map[core.WuXing]core.WuXing{
	core.Wood:  core.Earth,
	core.Earth: core.Water,
	core.Water: core.Fire,
	core.Fire:  core.Metal,
	core.Metal: core.Wood,
}

func GetWuXingByBagua(bagua core.Bagua) core.WuXing {
	switch bagua {
	case core.Qian, core.Dui:
		return core.Metal
	case core.Li:
		return core.Fire
	case core.Zhen, core.Xun:
		return core.Wood
	case core.Kan:
		return core.Water
	case core.Gen, core.Kun:
		return core.Earth
	default:
		return 0
	}
}
