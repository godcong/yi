package yi

import "testing"

func TestGetWenYan(t *testing.T) {
	gua, err := GetGuaByIndex("乾乾")
	if err != nil {
		t.Fatalf("GetGuaByIndex failed: %v", err)
	}

	wenyan := gua.GetWenYan()
	if wenyan == nil {
		t.Fatal("GetWenYan returned nil for 乾为天")
	}
	if wenyan.Index != "乾乾" {
		t.Errorf("Expected Index=乾乾, got %s", wenyan.Index)
	}
	if len(wenyan.Entries) == 0 {
		t.Error("WenYan entries are empty")
	}

	t.Logf("乾为天 文言 (%d entries):", len(wenyan.Entries))
	for _, e := range wenyan.Entries {
		t.Logf("  [%s] %s", e.Title, truncate(e.Text, 30))
	}
}

func TestGetWenYan_Kun(t *testing.T) {
	gua, err := GetGuaByIndex("坤坤")
	if err != nil {
		t.Fatalf("GetGuaByIndex failed: %v", err)
	}

	wenyan := gua.GetWenYan()
	if wenyan == nil {
		t.Fatal("GetWenYan returned nil for 坤为地")
	}

	t.Logf("坤为地 文言 (%d entries):", len(wenyan.Entries))
	for _, e := range wenyan.Entries {
		t.Logf("  [%s] %s", e.Title, truncate(e.Text, 30))
	}
}

func TestGetWenYan_NotExist(t *testing.T) {
	gua, err := GetGuaByIndex("坎震")
	if err != nil {
		t.Fatalf("GetGuaByIndex failed: %v", err)
	}

	wenyan := gua.GetWenYan()
	if wenyan != nil {
		t.Error("GetWenYan should return nil for hexagrams without 文言")
	}
}

func TestHasWenYan(t *testing.T) {
	gua, _ := GetGuaByIndex("乾乾")
	if !gua.HasWenYan() {
		t.Error("HasWenYan should return true for 乾为天")
	}

	gua2, _ := GetGuaByIndex("坤坤")
	if !gua2.HasWenYan() {
		t.Error("HasWenYan should return true for 坤为地")
	}
}

func TestHasWenYan_OtherHexagrams(t *testing.T) {
	samples := []string{"坎震", "艮坎", "乾坎", "乾离"}

	for _, idx := range samples {
		gua, err := GetGuaByIndex(idx)
		if err != nil {
			t.Errorf("GetGuaByIndex(%s) failed: %v", idx, err)
			continue
		}
		if gua.HasWenYan() {
			t.Errorf("HasWenYan should return false for %s", idx)
		}
	}
}

func TestWenYanEntryStructure(t *testing.T) {
	gua, _ := GetGuaByIndex("乾乾")
	wenyan := gua.GetWenYan()
	if wenyan == nil {
		t.Skip("WenYan not available")
	}

	for i, e := range wenyan.Entries {
		if e.Title == "" {
			t.Errorf("Entry %d has empty Title", i)
		}
		if e.Text == "" {
			t.Errorf("Entry %d has empty Text", i)
		}
	}
}
