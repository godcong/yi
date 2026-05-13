package gua

import "yi/core"

func getBenGua(shang, xia int) *core.Gua {
	idx := GetBaguaName(core.Bagua(shang%8)) + GetBaguaName(core.Bagua(xia%8))
	if g, ok := core.GuaStore[idx]; ok {
		return g
	}
	return nil
}

func getBianGua(ben *core.Gua, bianYao ...int) *core.Gua {
	bz := calcBianYao(bianYao...)
	shang := ben.ShangNum
	xia := ben.XiaNum
	if bz > 2 {
		shang = bianYaoTransform(shang, bz-3)
	} else {
		xia = bianYaoTransform(xia, bz)
	}
	idx := GetBaguaName(core.Bagua(shang)) + GetBaguaName(core.Bagua(xia))
	if g, ok := core.GuaStore[idx]; ok {
		return g
	}
	return nil
}

func bianYaoTransform(gua, pos int) int {
	mask := 1 << (2 - uint(pos))
	if gua&mask == 0 {
		return gua | mask
	}
	return gua ^ mask
}

func calcBianYao(bianYao ...int) int {
	if len(bianYao) == 0 {
		return 0
	}
	n := bianYao[0] % 6
	if n == 0 {
		return 5
	}
	return n - 1
}

func getHuGua(ben *core.Gua) *core.Gua {
	shang := calcHuShang(ben.ShangNum, ben.XiaNum)
	xia := calcHuXia(ben.ShangNum, ben.XiaNum)
	idx := GetBaguaName(core.Bagua(shang)) + GetBaguaName(core.Bagua(xia))
	if g, ok := core.GuaStore[idx]; ok {
		return g
	}
	return nil
}

func calcHuShang(shang, xia int) int {
	result := 0
	if xia&(1<<0) > 0 {
		result |= 1 << 2
	}
	if shang&(1<<2) > 0 {
		result |= 1 << 1
	}
	if shang&(1<<1) > 0 {
		result |= 1 << 0
	}
	return result
}

func calcHuXia(shang, xia int) int {
	result := 0
	if xia&(1<<1) > 0 {
		result |= 1 << 2
	}
	if xia&(1<<0) > 0 {
		result |= 1 << 1
	}
	if shang&(1<<2) > 0 {
		result |= 1 << 0
	}
	return result
}

func getCuoGua(ben *core.Gua) *core.Gua {
	shang := flipBits(ben.ShangNum)
	xia := flipBits(ben.XiaNum)
	idx := GetBaguaName(core.Bagua(shang)) + GetBaguaName(core.Bagua(xia))
	if g, ok := core.GuaStore[idx]; ok {
		return g
	}
	return nil
}

func getZongGua(ben *core.Gua) *core.Gua {
	shang := reverseBits(ben.XiaNum)
	xia := reverseBits(ben.ShangNum)
	idx := GetBaguaName(core.Bagua(shang)) + GetBaguaName(core.Bagua(xia))
	if g, ok := core.GuaStore[idx]; ok {
		return g
	}
	return nil
}

func bitLen(n int) int {
	count := 0
	for n > 0 {
		count++
		n >>= 1
	}
	return count
}

func reverseBits(n int) int {
	return ((n & 0x4) >> 2) | (n & 0x2) | ((n & 0x1) << 2)
}

func flipBit(n, pos int) int {
	return n ^ (1 << uint(pos))
}

func flipBits(n int) int {
	return ^n & 0x7
}

func newZhouYi(ben *core.Gua, bianYao ...int) *core.ZhouYi {
	bian := getBianGua(ben, bianYao...)
	hu := ben
	if (ben.ShangNum == core.Kun && ben.XiaNum == core.Kun) ||
		(ben.ShangNum == core.Qian && ben.XiaNum == core.Qian) {
		hu = bian
	}
	hu = getHuGua(hu)
	cuo := getCuoGua(ben)
	zong := getZongGua(ben)
	return &core.ZhouYi{
		Gua: [core.GuaTypeMax]*core.Gua{
			core.Ben:  ben,
			core.Bian: bian,
			core.Hu:   hu,
			core.Cuo:  cuo,
			core.Zong: zong,
		},
		BianYao: bianYao,
	}
}
