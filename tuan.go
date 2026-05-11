package yi

// ============================================================================
// Tuan (彖辞) - Judgment texts
// ============================================================================
//
// 彖辞是《周易》中解释卦名和卦辞的文字，传为孔子所作。
// 每一卦都有一段彖辞，用来断定一卦的吉凶及其所由来。
// "彖者，言乎象者也" —— 彖辞论断一卦之整体含义。
//
// 彖辞数据存储在 Gua.TuanText 字段中。
//
// Reference:
//   - 《周易正义》孔颖达
//   - 《周易本义》朱熹

// GetTuan returns the Tuan (彖辞) text for a given hexagram.
func (g *Gua) GetTuan() string {
	return g.TuanText
}

// GetAllTuan returns all Tuan data as a slice of Index+Text pairs.
func GetAllTuan() []struct {
	Index string
	Text  string
} {
	result := make([]struct {
		Index string
		Text  string
	}, 0, 64)
	for _, g := range guaStore {
		if g.TuanText != "" {
			result = append(result, struct {
				Index string
				Text  string
			}{Index: g.Index, Text: g.TuanText})
		}
	}
	return result
}
