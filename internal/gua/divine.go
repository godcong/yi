package gua

import "yi/core"

func Divine(shang, xia core.Bagua, bianYao ...int) *core.ZhouYi {
	return DivineByNumber(int(shang), int(xia), bianYao...)
}

func DivineByNumber(shang, xia int, bianYao ...int) *core.ZhouYi {
	ben := getBenGua(shang, xia)
	if ben == nil {
		return nil
	}
	return newZhouYi(ben, bianYao...)
}

func DivineByTime(year, month, day, hour int, seed ...string) *core.ZhouYi {
	sum := year + month + day
	shang := ShaoYongToBagua(sum)
	lower := ShaoYongToBagua(sum + hour)
	bian := (sum + hour) % 6
	if len(seed) > 0 && seed[0] != "" {
		s := int64(sum*10000 + hour)
		for _, c := range seed[0] {
			s = s*31 + int64(c)
		}
		bian = int(s % 6)
	}
	return DivineByNumber(int(shang), int(lower), bian)
}

func ShaoYongToBagua(n int) core.Bagua {
	return core.Bagua((n - 1 + 8) % 8)
}

func GetBianYao(zy *core.ZhouYi) int {
	return calcBianYao(zy.BianYao...)
}

func GetAllBianYao(zy *core.ZhouYi) []int {
	result := make([]int, len(zy.BianYao))
	copy(result, zy.BianYao)
	return result
}

func GetYao(g *core.Gua, pos core.YaoPosition) *core.Yao {
	if pos < 0 || int(pos) >= int(core.YaoCount) {
		return nil
	}
	return g.Yaos[pos]
}
