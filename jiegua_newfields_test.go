package yi

import (
	"strings"
	"testing"
)

// ============================================================================
// GuaJieDu new fields (CoreImage, Yi, Ji) tests
// ============================================================================

func TestAllGuaHaveJieDu(t *testing.T) {
	missing := 0
	for xu := 1; xu <= 64; xu++ {
		g, err := GetGuaByXu(xu)
		if err != nil {
			t.Fatalf("GetGuaByXu(%d) failed: %v", xu, err)
		}
		jd := GetGuaJieDuByIndex(g.Index)
		if jd == nil {
			missing++
			t.Errorf("xu=%d (%s, index=%s): no JieDu entry", xu, g.Ming, g.Index)
		}
	}
	if missing > 0 {
		t.Errorf("%d/64 gua missing JieDu", missing)
	}
}

func TestAllGuaJieDuCoreImage(t *testing.T) {
	emptyCount := 0
	for xu := 1; xu <= 64; xu++ {
		g, _ := GetGuaByXu(xu)
		jd := GetGuaJieDuByIndex(g.Index)
		if jd == nil {
			continue
		}
		if jd.CoreImage == "" {
			emptyCount++
			t.Errorf("xu=%d (%s): CoreImage is empty", xu, g.Ming)
		}
	}
	if emptyCount > 0 {
		t.Errorf("%d/64 gua have empty CoreImage", emptyCount)
	}
}

func TestAllGuaJieDuYi(t *testing.T) {
	emptyCount := 0
	for xu := 1; xu <= 64; xu++ {
		g, _ := GetGuaByXu(xu)
		jd := GetGuaJieDuByIndex(g.Index)
		if jd == nil {
			continue
		}
		if len(jd.Yi) == 0 {
			emptyCount++
			t.Errorf("xu=%d (%s): Yi list is empty", xu, g.Ming)
		}
	}
	if emptyCount > 0 {
		t.Errorf("%d/64 gua have empty Yi list", emptyCount)
	}
}

func TestAllGuaJieDuJi(t *testing.T) {
	emptyCount := 0
	for xu := 1; xu <= 64; xu++ {
		g, _ := GetGuaByXu(xu)
		jd := GetGuaJieDuByIndex(g.Index)
		if jd == nil {
			continue
		}
		if len(jd.Ji) == 0 {
			emptyCount++
			t.Errorf("xu=%d (%s): Ji list is empty", xu, g.Ming)
		}
	}
	if emptyCount > 0 {
		t.Errorf("%d/64 gua have empty Ji list", emptyCount)
	}
}

func TestJieDuEightDimensions(t *testing.T) {
	// All 64 gua should have non-empty content for all 8 dimensions
	dims := []struct {
		name string
		get  func(*GuaJieDu) string
	}{
		{"事业", func(j *GuaJieDu) string { return j.ShiYe }},
		{"爱情", func(j *GuaJieDu) string { return j.AiQing }},
		{"财运", func(j *GuaJieDu) string { return j.CaiYun }},
		{"考试", func(j *GuaJieDu) string { return j.KaoShi }},
		{"健康", func(j *GuaJieDu) string { return j.JianKang }},
		{"出行", func(j *GuaJieDu) string { return j.ChuXing }},
		{"官司", func(j *GuaJieDu) string { return j.GuanSi }},
		{"家宅", func(j *GuaJieDu) string { return j.JiaZhai }},
	}
	for _, dim := range dims {
		emptyCount := 0
		for xu := 1; xu <= 64; xu++ {
			g, _ := GetGuaByXu(xu)
			jd := GetGuaJieDuByIndex(g.Index)
			if jd == nil {
				continue
			}
			if dim.get(jd) == "" {
				emptyCount++
			}
		}
		if emptyCount > 0 {
			t.Errorf("%s: %d/64 gua have empty content", dim.name, emptyCount)
		}
	}
}

func TestJieDuDimensionMinLength(t *testing.T) {
	// Each dimension should have meaningful content (>= 20 chars)
	minLen := 20
	dims := []struct {
		name string
		get  func(*GuaJieDu) string
	}{
		{"事业", func(j *GuaJieDu) string { return j.ShiYe }},
		{"爱情", func(j *GuaJieDu) string { return j.AiQing }},
		{"财运", func(j *GuaJieDu) string { return j.CaiYun }},
		{"考试", func(j *GuaJieDu) string { return j.KaoShi }},
		{"健康", func(j *GuaJieDu) string { return j.JianKang }},
		{"出行", func(j *GuaJieDu) string { return j.ChuXing }},
		{"官司", func(j *GuaJieDu) string { return j.GuanSi }},
		{"家宅", func(j *GuaJieDu) string { return j.JiaZhai }},
	}
	for _, dim := range dims {
		shortCount := 0
		for xu := 1; xu <= 64; xu++ {
			g, _ := GetGuaByXu(xu)
			jd := GetGuaJieDuByIndex(g.Index)
			if jd == nil {
				continue
			}
			text := dim.get(jd)
			if len([]rune(text)) < minLen {
				shortCount++
				t.Errorf("xu=%d (%s) %s: only %d chars (min %d)", xu, g.Ming, dim.name, len([]rune(text)), minLen)
			}
		}
		if shortCount > 0 {
			t.Errorf("%s: %d/64 gua below minimum length", dim.name, shortCount)
		}
	}
}

func TestJieDuYiCount(t *testing.T) {
	// Each gua should have at least 3 Yi items
	for xu := 1; xu <= 64; xu++ {
		g, _ := GetGuaByXu(xu)
		jd := GetGuaJieDuByIndex(g.Index)
		if jd == nil {
			continue
		}
		if len(jd.Yi) < 3 {
			t.Errorf("xu=%d (%s): only %d Yi items (min 3)", xu, g.Ming, len(jd.Yi))
		}
	}
}

func TestJieDuJiCount(t *testing.T) {
	// Each gua should have at least 3 Ji items
	for xu := 1; xu <= 64; xu++ {
		g, _ := GetGuaByXu(xu)
		jd := GetGuaJieDuByIndex(g.Index)
		if jd == nil {
			continue
		}
		if len(jd.Ji) < 3 {
			t.Errorf("xu=%d (%s): only %d Ji items (min 3)", xu, g.Ming, len(jd.Ji))
		}
	}
}

func TestJieDuNoDuplicateYiJi(t *testing.T) {
	// Yi and Ji should not have duplicate items within each list
	for xu := 1; xu <= 64; xu++ {
		g, _ := GetGuaByXu(xu)
		jd := GetGuaJieDuByIndex(g.Index)
		if jd == nil {
			continue
		}
		seen := make(map[string]bool)
		for _, y := range jd.Yi {
			if seen[y] {
				t.Errorf("xu=%d (%s): duplicate Yi item %q", xu, g.Ming, y)
			}
			seen[y] = true
		}
		seenJi := make(map[string]bool)
		for _, j := range jd.Ji {
			if seenJi[j] {
				t.Errorf("xu=%d (%s): duplicate Ji item %q", xu, g.Ming, j)
			}
			seenJi[j] = true
		}
	}
}

func TestJieDuDimensionNoDuplication(t *testing.T) {
	// Dimension content should not have obvious duplicated sentences
	// (same 8+ char sequence appearing twice separated by punctuation)
	// Using 8-char threshold to avoid false positives on common short phrases
	for xu := 1; xu <= 64; xu++ {
		g, _ := GetGuaByXu(xu)
		jd := GetGuaJieDuByIndex(g.Index)
		if jd == nil {
			continue
		}
		for _, text := range []string{jd.ShiYe, jd.AiQing, jd.CaiYun, jd.KaoShi, jd.JianKang, jd.ChuXing, jd.GuanSi, jd.JiaZhai} {
			if text == "" {
				continue
			}
			// Split by sentence separators and check for exact duplicates
			sentences := strings.FieldsFunc(text, func(r rune) bool {
				return r == '。' || r == '，' || r == '；' || r == '、'
			})
			seen := make(map[string]bool)
			for _, s := range sentences {
				s = strings.TrimSpace(s)
				if len([]rune(s)) < 8 {
					continue // skip short fragments that may be common phrases
				}
				if seen[s] {
					t.Errorf("xu=%d (%s): duplicate sentence in dimension: %q", xu, g.Ming, s)
				}
				seen[s] = true
			}
		}
	}
}

// ============================================================================
// WuXingInfo tests
// ============================================================================

func TestWuXingInfoAllBagua(t *testing.T) {
	// All 8 bagua should return valid WuXingInfo
	for bg := Bagua(0); bg <= 7; bg++ {
		info := GetWuXingInfo(int(bg))
		if info == nil {
			t.Errorf("GetWuXingInfo(%d/%s) returned nil", bg, GetBaguaName(bg))
			continue
		}
		if info.WuXing == "" {
			t.Errorf("Bagua %d: WuXing is empty", bg)
		}
		if info.Direction == "" {
			t.Errorf("Bagua %d: Direction is empty", bg)
		}
		if info.LuckyNumber == "" {
			t.Errorf("Bagua %d: LuckyNumber is empty", bg)
		}
		if info.LuckyColor == "" {
			t.Errorf("Bagua %d: LuckyColor is empty", bg)
		}
	}
}

func TestWuXingInfoMapping(t *testing.T) {
	// Verify specific bagua → wuxing mappings
	tests := []struct {
		bagua   Bagua
		wantWX  string
		wantDir string
	}{
		{0, "金", "西、西北"},  // 乾
		{1, "金", "西、西北"},  // 兑
		{2, "火", "南"},       // 离
		{3, "木", "东、东南"},  // 震
		{4, "木", "东、东南"},  // 巽
		{5, "水", "北"},       // 坎
		{6, "土", "东北、西南"}, // 艮
		{7, "土", "东北、西南"}, // 坤
	}
	for _, tt := range tests {
		info := GetWuXingInfo(int(tt.bagua))
		if info == nil {
			t.Errorf("GetWuXingInfo(%d) returned nil", tt.bagua)
			continue
		}
		if info.WuXing != tt.wantWX {
			t.Errorf("Bagua %d (%s): WuXing=%q, want %q", tt.bagua, GetBaguaName(tt.bagua), info.WuXing, tt.wantWX)
		}
		if info.Direction != tt.wantDir {
			t.Errorf("Bagua %d (%s): Direction=%q, want %q", tt.bagua, GetBaguaName(tt.bagua), info.Direction, tt.wantDir)
		}
	}
}

func TestWuXingInfoInvalidInput(t *testing.T) {
	// Out-of-range bagua number should default to 土
	info := GetWuXingInfo(99)
	if info == nil {
		t.Error("GetWuXingInfo(99) should default to 土, not nil")
	} else if info.WuXing != "土" {
		t.Errorf("GetWuXingInfo(99).WuXing=%q, want 土 (default)", info.WuXing)
	}
}

// ============================================================================
// JieGua result new fields integration test
// ============================================================================

func TestJieGuaResultHasJieDu(t *testing.T) {
	// Divine should produce a result with JieDu populated
	result := JieGua(DivineByNumber(0, 0, 1), Male)
	if result.JieDu == nil {
		t.Fatal("JieGua result.JieDu is nil")
	}
	if result.JieDu.CoreImage == "" {
		t.Error("JieGua result.JieDu.CoreImage is empty")
	}
	if len(result.JieDu.Yi) == 0 {
		t.Error("JieGua result.JieDu.Yi is empty")
	}
	if len(result.JieDu.Ji) == 0 {
		t.Error("JieGua result.JieDu.Ji is empty")
	}
}

func TestJieGuaResultHasWuXingInfo(t *testing.T) {
	// Divine should produce a result with WuXingInfo populated
	result := JieGua(DivineByNumber(0, 0, 1), Male)
	if result.WuXingInfo == nil {
		t.Fatal("JieGua result.WuXingInfo is nil")
	}
	if result.WuXingInfo.WuXing == "" {
		t.Error("JieGua result.WuXingInfo.WuXing is empty")
	}
	if result.WuXingInfo.Direction == "" {
		t.Error("JieGua result.WuXingInfo.Direction is empty")
	}
	if result.WuXingInfo.LuckyNumber == "" {
		t.Error("JieGua result.WuXingInfo.LuckyNumber is empty")
	}
	if result.WuXingInfo.LuckyColor == "" {
		t.Error("JieGua result.WuXingInfo.LuckyColor is empty")
	}
}

func TestJieGuaResultWuXingMatchesBenGua(t *testing.T) {
	// WuXing should be derived from BenGua's upper trigram
	result := JieGua(DivineByNumber(0, 0, 1), Male) // 乾为天, ShangNum=0 → 金
	if result.WuXingInfo == nil {
		t.Fatal("WuXingInfo is nil")
	}
	// 乾(ShangNum=0) → 金
	benGua := result.BenGuaInfo
	expectedWX := GetWuXingInfo(benGua.Gua.ShangNum)
	if result.WuXingInfo.WuXing != expectedWX.WuXing {
		t.Errorf("WuXing=%q, expected %q (from ShangNum=%d)",
			result.WuXingInfo.WuXing, expectedWX.WuXing, benGua.Gua.ShangNum)
	}
}

// ============================================================================
// AllJieDuEntries completeness
// ============================================================================

func TestAllJieDuEntriesComplete(t *testing.T) {
	entries := AllJieDuEntries()
	if len(entries) != 64 {
		t.Errorf("AllJieDuEntries() returned %d entries, want 64", len(entries))
	}
	for _, jd := range entries {
		if jd.CoreImage == "" {
			t.Errorf("JieDu entry %s: CoreImage is empty", jd.Ming)
		}
	}
}

// ============================================================================
// GetGuaJieDuByIndex vs GetGuaJieDu
// ============================================================================

func TestGetGuaJieDuByIndexConsistency(t *testing.T) {
	// GetGuaJieDuByIndex should find entries since keys are index-based
	for xu := 1; xu <= 64; xu++ {
		g, _ := GetGuaByXu(xu)
		jd := GetGuaJieDuByIndex(g.Index)
		if jd == nil {
			t.Errorf("GetGuaJieDuByIndex(%q) returned nil for xu=%d", g.Index, xu)
		}
	}
}

func TestGetGuaJieDuByMingMayFail(t *testing.T) {
	// GetGuaJieDu uses Ming (e.g., "乾为天") but store keys are Index (e.g., "乾乾")
	// This test documents that GetGuaJieDu(Ming) returns nil for the current data format
	g, _ := GetGuaByXu(1) // 乾为天
	jd := GetGuaJieDu(g.Ming) // "乾为天" — not a key
	// Store uses Index as key, so Ming lookup should fail
	if jd != nil {
		t.Logf("GetGuaJieDu(%q) = non-nil (unexpected but ok)", g.Ming)
	}
	// But Index lookup should work
	jdIdx := GetGuaJieDuByIndex(g.Index) // "乾乾" — the actual key
	if jdIdx == nil {
		t.Errorf("GetGuaJieDuByIndex(%q) = nil, expected non-nil", g.Index)
	}
}

// ============================================================================
// DivineByDailyHexagram tests
// ============================================================================

func TestDailyHexagramSamePersonSameDay(t *testing.T) {
	// Same person + same day should always produce the same hexagram
	zy1 := DivineByDailyHexagram(2026, 5, 12, "张三")
	zy2 := DivineByDailyHexagram(2026, 5, 12, "张三")
	if zy1.GetGua(Ben).Index != zy2.GetGua(Ben).Index {
		t.Errorf("Same person same day: %s != %s", zy1.GetGua(Ben).Index, zy2.GetGua(Ben).Index)
	}
}

func TestDailyHexagramDifferentPerson(t *testing.T) {
	// Different people should (almost always) get different hexagrams
	zy1 := DivineByDailyHexagram(2026, 5, 12, "张三")
	zy2 := DivineByDailyHexagram(2026, 5, 12, "李四")
	// We can't guarantee different results (hash collision possible), but log it
	t.Logf("张三: %s, 李四: %s", zy1.GetGua(Ben).Ming, zy2.GetGua(Ben).Ming)
}

func TestDailyHexagramSamePersonDifferentDay(t *testing.T) {
	// Same person on different days should get different hexagrams
	zy1 := DivineByDailyHexagram(2026, 5, 12, "张三")
	zy2 := DivineByDailyHexagram(2026, 5, 13, "张三")
	t.Logf("5/12: %s, 5/13: %s", zy1.GetGua(Ben).Ming, zy2.GetGua(Ben).Ming)
}

func TestDailyHexagramEmptySeed(t *testing.T) {
	// Empty seed should still work (falls back to date-based)
	zy := DivineByDailyHexagram(2026, 5, 12, "")
	if zy == nil {
		t.Error("DivineByDailyHexagram with empty seed returned nil")
	}
}

func TestDailyHexagramNow(t *testing.T) {
	zy := DivineByDailyHexagramNow("测试用户")
	if zy == nil {
		t.Error("DivineByDailyHexagramNow returned nil")
	}
}

func TestDailyHexagramStability(t *testing.T) {
	// Run 100 times to verify determinism
	first := DivineByDailyHexagram(2026, 5, 12, "稳定性测试")
	for i := 0; i < 100; i++ {
		zy := DivineByDailyHexagram(2026, 5, 12, "稳定性测试")
		if zy.GetGua(Ben).Index != first.GetGua(Ben).Index {
			t.Fatalf("Iteration %d: got %s, want %s", i, zy.GetGua(Ben).Index, first.GetGua(Ben).Index)
		}
	}
}
