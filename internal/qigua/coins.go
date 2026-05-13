package qigua

import (
	"math/rand"

	"yi/core"
	"yi/internal/gua"
)

func DivineByCoins(seed int64) (*core.ZhouYi, [6]core.CoinResult) {
	rng := rand.New(rand.NewSource(seed))
	var results [6]core.CoinResult
	var shangBits, xiaBits int

	for i := 0; i < 6; i++ {
		cr := throwCoins(rng)
		results[i] = cr
		isYin := cr == core.CoinYinYinYang || cr == core.CoinYinYinYin
		if isYin {
			xiaBits |= 1 << i
		}
	}

	shangBits = xiaBits >> 3
	xiaBits = xiaBits & 0x7

	var bianYao []int
	for i, cr := range results {
		if cr.IsChanging() {
			bianYao = append(bianYao, i)
		}
	}

	return gua.DivineByNumber(shangBits, xiaBits, bianYao...), results
}

func throwCoins(rng *rand.Rand) core.CoinResult {
	yin := rng.Intn(4)
	switch yin {
	case 0:
		return core.CoinYangYangYang
	case 1:
		return core.CoinYinYangYang
	case 2:
		return core.CoinYinYinYang
	case 3:
		return core.CoinYinYinYin
	}
	return core.CoinYinYinYang
}
