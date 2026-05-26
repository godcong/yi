package qigua

import (
	"time"

	"github.com/godcong/yi/core"
)

func DivineByDailyHexagram(year, month, day int, personalSeed string) *core.ZhouYi {
	if personalSeed == "" {
		return DivineByTimeGua(core.TimeGuaParams{Year: year, Month: month, Day: day, Hour: 0})
	}

	seed := int64(year*10000+month*100+day) * 31
	for _, c := range personalSeed {
		seed = seed*31 + int64(c)
	}

	zy, _ := DivineByCoins(seed)
	return zy
}

func DivineByDailyHexagramNow(personalSeed string) *core.ZhouYi {
	now := time.Now()
	return DivineByDailyHexagram(now.Year(), int(now.Month()), now.Day(), personalSeed)
}
