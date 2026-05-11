package yi

import "testing"

func TestGetTuan(t *testing.T) {
	gua, err := GetGuaByIndex("乾乾")
	if err != nil {
		t.Fatalf("GetGuaByIndex failed: %v", err)
	}

	tuan := gua.GetTuan()
	if tuan == "" {
		t.Error("GetTuan returned empty string for 乾为天")
	}
	t.Logf("乾为天 彖辞: %s", truncate(tuan, 50))
}

func TestGetTuan_NotEmpty(t *testing.T) {
	gua, _ := GetGuaByIndex("乾乾")
	tuan := gua.GetTuan()
	if !contains(tuan, "乾") {
		t.Error("乾为天 彖辞 should contain '乾'")
	}
}

func TestGetAllTuan_Count(t *testing.T) {
	tuans := GetAllTuan()
	if len(tuans) < 64 {
		t.Errorf("GetAllTuan returned %d entries (expected 64)", len(tuans))
	}

	for _, entry := range tuans {
		if entry.Index == "" {
			t.Error("Tuan entry has empty Index")
		}
		if entry.Text == "" {
			t.Errorf("Tuan entry for %s has empty Text", entry.Index)
		}
	}
	t.Logf("Total Tuan entries: %d", len(tuans))
}

func TestGetTuan_SampleHexagrams(t *testing.T) {
	samples := []struct {
		index string
		name  string
	}{
		{"乾乾", "乾为天"},
		{"坤坤", "坤为地"},
		{"坎震", "水雷屯"},
		{"乾坎", "水天需"},
	}

	for _, s := range samples {
		gua, err := GetGuaByIndex(s.index)
		if err != nil {
			t.Errorf("GetGuaByIndex(%s) failed: %v", s.index, err)
			continue
		}
		tuan := gua.GetTuan()
		t.Logf("%s(%s) 彖辞: %s", s.name, s.index, truncate(tuan, 30))
	}
}

func truncate(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen] + "..."
}
