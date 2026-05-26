package wenyan

import "github.com/godcong/yi/core"

func GetWenYan(g *core.Gua) *core.WenYanData {
	if entries, ok := core.WenYanStore[g.Index]; ok {
		return &core.WenYanData{
			Index:   g.Index,
			Entries: entries,
		}
	}
	return nil
}

func HasWenYan(g *core.Gua) bool {
	_, ok := core.WenYanStore[g.Index]
	return ok
}

func GetAllWenYan() []*core.WenYanData {
	result := make([]*core.WenYanData, 0, len(core.WenYanStore))
	for index, entries := range core.WenYanStore {
		result = append(result, &core.WenYanData{
			Index:   index,
			Entries: entries,
		})
	}
	return result
}
