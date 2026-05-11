package yi

// ============================================================================
// Xiang (象辞/大象辞) - Image/Interpretation texts
// ============================================================================
//
// 象辞是《周易》中解释卦象和爻象的文字，分大象和小象：
//   - 大象：解释整个卦象（如"天行健，君子以自强不息"）
//   - 小象：解释各爻的爻辞
//
// 本模块存储的是大象辞（每卦一条），数据在 Gua.XiangText 字段中。
//
// 注意：Gua.GuaName 是卦名单字（如"乾"），不是象辞。
//
// Reference:
//   - 《周易正义》孔颖达
//   - 《周易本义》朱熹

// GetXiang returns the Xiang (大象辞) text for a given hexagram.
func (g *Gua) GetXiang() string {
	return g.XiangText
}

// GetAllXiang returns all Xiang data as a slice of Index+Text pairs.
func GetAllXiang() []struct {
	Index string
	Text  string
} {
	result := make([]struct {
		Index string
		Text  string
	}, 0, 64)
	for _, g := range guaStore {
		if g.XiangText != "" {
			result = append(result, struct {
				Index string
				Text  string
			}{Index: g.Index, Text: g.XiangText})
		}
	}
	return result
}
