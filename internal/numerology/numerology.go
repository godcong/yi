package numerology

import "github.com/godcong/yi/core"

func GetDayan(number int) (*core.Dayan, error) {
	if number < 1 || number > 81 {
		return nil, core.ErrInvalidDayanNumber
	}
	return &core.DayanList[number-1], nil
}

func MustGetDayan(number int) core.Dayan {
	if number < 1 || number > 81 {
		return core.Dayan{}
	}
	return core.DayanList[(number-1)%81]
}

func DayanIsJi(d *core.Dayan) bool {
	return d.JiXiong == "吉" || d.JiXiong == "半吉"
}
