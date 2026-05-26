package qigua

import (
	"time"

	"github.com/godcong/yi/core"
	"github.com/godcong/yi/internal/gua"
)

func DivineByTimeGua(params core.TimeGuaParams, seed ...string) *core.ZhouYi {
	sum := params.Year + params.Month + params.Day
	shang := gua.ShaoYongToBagua(sum)
	lower := gua.ShaoYongToBagua(sum + params.Hour)
	bian := (sum + params.Hour) % 6

	if len(seed) > 0 && seed[0] != "" {
		s := int64(sum*10000 + params.Hour)
		for _, c := range seed[0] {
			s = s*31 + int64(c)
		}
		zy, _ := DivineByCoins(s)
		return zy
	}

	return gua.DivineByNumber(int(shang), int(lower), bian)
}

func DivineByCurrentTime(seed ...string) *core.ZhouYi {
	now := time.Now()
	return DivineByTimeGua(core.TimeGuaParams{
		Year:  now.Year(),
		Month: int(now.Month()),
		Day:   now.Day(),
		Hour:  now.Hour(),
	}, seed...)
}
