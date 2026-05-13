package gua

import "yi/core"

func IsJi(zy *core.ZhouYi, sex core.Sex) bool {
	bianGua := zy.Gua[core.Bian]
	bianYaoPos := calcBianYao(zy.BianYao...)
	if bianGua == nil || bianYaoPos < 0 || bianYaoPos >= int(core.YaoCount) {
		return false
	}
	yao := bianGua.Yaos[bianYaoPos]
	if yao == nil {
		return false
	}
	if sex == core.Female && yao.NvMing != "" {
		return !contains(yao.NvMing, "凶")
	}
	return !contains(yao.JiXiong, "凶")
}

func FilterYao(zy *core.ZhouYi, sex core.Sex, filters ...string) bool {
	bianGua := zy.Gua[core.Bian]
	bianYaoPos := calcBianYao(zy.BianYao...)
	if bianGua == nil || bianYaoPos < 0 || bianYaoPos >= int(core.YaoCount) {
		return true
	}
	yao := bianGua.Yaos[bianYaoPos]
	if yao == nil {
		return true
	}
	for _, f := range filters {
		if sex == core.Female && yao.NvMing != "" {
			if yao.NvMing == f {
				return false
			}
			return true
		}
		if yao.JiXiong == f {
			return false
		}
	}
	return true
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && containsHelper(s, substr)
}

func containsHelper(s, substr string) bool {
	if len(substr) == 0 {
		return true
	}
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
