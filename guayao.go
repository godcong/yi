package yi

import "strings"

// GuaYao ...
type GuaYao struct {
	Yao     string //二爻
	JiXiong string //二爻吉凶
	NvMing  string //女命
}

func getGuaYao(xiang *GuaXiang, yao int) GuaYao {
	if yao < 0 || yao > 5 {
		panic("wrong yao")
	}

	return *xiang.GuaYaos[yao]
}

func (y *Yi) FilterYao(sex Sex, fs ...string) bool {
	yao := getGuaYao(y.Get(BianGua), y.BianYao())
	for _, s := range fs {
		if sex == SexGirl && yao.NvMing != "" {
			if yao.NvMing == s {
				return false
			} else {
				return true
			}
		}

		if yao.JiXiong == s {
			return false
		}
	}
	return true
}

func (y *Yi) IsLucky(sex Sex) bool {
	yao := getGuaYao(y.Get(BianGua), y.BianYao())

	if sex == SexGirl && yao.NvMing != "" {
		if strings.Contains(yao.NvMing, "凶") {
			return false
		}
	} else if strings.Contains(yao.JiXiong, "凶") {
		return false
	}

	return true
}
