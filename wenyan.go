package yi

// ============================================================================
// WenYan (文言) - Wenyan commentaries
// ============================================================================
//
// 文言传是《易传》之一，专门解释乾坤两卦的卦辞和爻辞。
// "文言"意为"文饰其言"，即以华美的文辞阐释乾坤两卦的深意。
//
// 文言传仅见于乾坤两卦：
//   - 乾卦文言：详细阐释乾卦各爻的道德含义
//   - 坤卦文言：详细阐释坤卦各爻的柔顺之道
//
// Reference:
//   - 《周易正义》孔颖达
//   - 《周易本义》朱熹

// WenYanData stores the WenYan (文言) text for a hexagram.
type WenYanData struct {
	Index   string        // hexagram index ("乾乾" or "坤坤")
	Entries []WenYanEntry // ordered entries
}

// WenYanEntry represents a single WenYan commentary passage.
type WenYanEntry struct {
	Title string // section title (e.g., "元者善之长也")
	Text  string // commentary text
}

// GetWenYan returns the WenYan (文言) data for a given hexagram.
// Only Qian (乾) and Kun (坤) hexagrams have WenYan texts.
func (g *Gua) GetWenYan() *WenYanData {
	if entries, ok := wenYanStore[g.Index]; ok {
		return &WenYanData{
			Index:   g.Index,
			Entries: entries,
		}
	}
	return nil
}

// HasWenYan returns whether the hexagram has WenYan text.
func (g *Gua) HasWenYan() bool {
	_, ok := wenYanStore[g.Index]
	return ok
}

// wenYanStore maps hexagram index to WenYan entries.
// Only "乾乾" and "坤坤" have entries.
var wenYanStore map[string][]WenYanEntry
