package qigua

import (
	"time"

	"github.com/godcong/yi/core"
	"github.com/godcong/yi/internal/gua"
)

func DivineByMeihua(upperNum, lowerNum, dongYao int) *core.ZhouYi {
	upper := gua.ShaoYongToBagua(upperNum)
	lower := gua.ShaoYongToBagua(lowerNum)
	var bYao []int
	if dongYao >= 0 && dongYao <= 5 {
		bYao = []int{dongYao}
	}
	return gua.DivineByNumber(int(upper), int(lower), bYao...)
}

func DivineByMeihuaTime(t time.Time, seeds ...string) (*core.ZhouYi, int, int, int, int) {
	year := t.Year()
	month := int(t.Month())
	day := t.Day()
	dzHour := shichen(t.Hour())

	sum := year + month + day
	upper := gua.ShaoYongToBagua(sum)
	lower := gua.ShaoYongToBagua(sum + dzHour)
	bian := (sum + dzHour) % 6

	if len(seeds) > 0 && seeds[0] != "" {
		seed := int64(sum*10000 + dzHour)
		for _, c := range seeds[0] {
			seed = seed*31 + int64(c)
		}
		zy, _ := DivineByCoins(seed)
		return zy, year, month, day, dzHour
	}

	return gua.DivineByNumber(int(upper), int(lower), bian), year, month, day, dzHour
}

func shichen(hour int) int {
	if hour >= 23 || hour < 1 {
		return 1
	}
	return (hour+1)/2 + 1
}
