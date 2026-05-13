package gua

import "yi/core"

func GetGuaByIndex(index string) (*core.Gua, error) {
	if g, ok := core.GuaStore[index]; ok {
		return g, nil
	}
	return nil, core.ErrGuaNotFound
}

func GetGuaByXu(xu int) (*core.Gua, error) {
	if xu < 1 || xu > 64 {
		return nil, core.ErrInvalidGuaIndex
	}
	for _, g := range core.GuaStore {
		if g.Xu == xu {
			return g, nil
		}
	}
	return nil, core.ErrGuaNotFound
}

func GetBaguaName(b core.Bagua) string {
	if b < 0 || b > 7 {
		return ""
	}
	return core.BaguaNames[b]
}

func GetBaguaSymbol(b core.Bagua) string {
	if b < 0 || b > 7 {
		return ""
	}
	return core.BaguaSymbols[b]
}

func GetAllGua() map[string]*core.Gua {
	result := make(map[string]*core.Gua, len(core.GuaStore))
	for k, v := range core.GuaStore {
		result[k] = v
	}
	return result
}
