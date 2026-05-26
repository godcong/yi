package qigua

import (
	"github.com/godcong/chronos/v2"
	"github.com/godcong/yi/core"
	"github.com/godcong/yi/internal/gua"
)

const defaultShichenHour = 12

func DivineByBirthday(params core.BirthdayParams) *core.ZhouYi {
	var cal chronos.Calendar
	if params.IsLunar {
		cal = chronos.ParseLunarDate(params.Year, params.Month, params.Day, params.Hour, 0, 0, params.IsLeapMonth)
	} else {
		cal = chronos.ParseSolarDate(params.Year, params.Month, params.Day, params.Hour, 0, 0)
	}
	if cal == nil {
		return nil
	}

	solar := cal.Solar()
	year := solar.GetYear()
	month := solar.GetMonth()
	day := solar.GetDay()
	hour := solar.GetHour()
	if hour == 0 {
		hour = defaultShichenHour
	}

	sum := year + month + day
	shang := gua.ShaoYongToBagua(sum)
	lower := gua.ShaoYongToBagua(sum + shichen(hour))
	bian := (sum + shichen(hour)) % 6

	return gua.DivineByNumber(int(shang), int(lower), bian)
}
