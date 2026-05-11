package yi

import "testing"

func TestGetXiang(t *testing.T) {
	gua, err := GetGuaByIndex("乾乾")
	if err != nil {
		t.Fatalf("GetGuaByIndex failed: %v", err)
	}

	xiang := gua.GetXiang()
	if xiang == "" {
		t.Error("GetXiang returned empty string for 乾为天")
	}
	t.Logf("乾为天 象辞: %s", xiang)
}

func TestGetXiang_KnownText(t *testing.T) {
	gua, _ := GetGuaByIndex("乾乾")
	xiang := gua.GetXiang()
	if !contains(xiang, "天") && !contains(xiang, "行") {
		t.Error("乾为天 象辞 should mention '天' or '行'")
	}
}

func TestGetAllXiang_Count(t *testing.T) {
	xiangs := GetAllXiang()
	if len(xiangs) < 64 {
		t.Errorf("GetAllXiang returned %d entries (expected 64)", len(xiangs))
	}

	for _, entry := range xiangs {
		if entry.Index == "" {
			t.Error("Xiang entry has empty Index")
		}
		if entry.Text == "" {
			t.Errorf("Xiang entry for %s has empty Text", entry.Index)
		}
	}
	t.Logf("Total Xiang entries: %d", len(xiangs))
}

func TestGetXiang_SampleHexagrams(t *testing.T) {
	samples := []string{"乾乾", "坤坤", "坎震", "乾离"}

	for _, idx := range samples {
		gua, err := GetGuaByIndex(idx)
		if err != nil {
			t.Errorf("GetGuaByIndex(%s) failed: %v", idx, err)
			continue
		}
		xiang := gua.GetXiang()
		t.Logf("%s 象辞: %s", idx, xiang)
	}
}
