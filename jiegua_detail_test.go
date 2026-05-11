package yi

import "testing"

func TestGetGuaJieDu(t *testing.T) {
	jd := GetGuaJieDu("乾乾")
	if jd == nil {
		t.Fatal("GetGuaJieDu returned nil for 乾为天")
	}
	if jd.Ming != "乾乾" {
		t.Errorf("Expected Ming=乾乾, got %s", jd.Ming)
	}

	t.Logf("GuaJieDu for %s:", jd.Ming)
	t.Logf("  事业: %s", jd.ShiYe)
	t.Logf("  爱情: %s", jd.AiQing)
	t.Logf("  财运: %s", jd.CaiYun)
	t.Logf("  考试: %s", jd.KaoShi)
	t.Logf("  健康: %s", jd.JianKang)
	t.Logf("  出行: %s", jd.ChuXing)
	t.Logf("  官司: %s", jd.GuanSi)
	t.Logf("  家宅: %s", jd.JiaZhai)
}

func TestGetGuaJieDu_NotFound(t *testing.T) {
	jd := GetGuaJieDu("不存在的卦")
	if jd != nil {
		t.Error("GetGuaJieDu should return nil for non-existent hexagram")
	}
}

func TestGetGuaJieDu_Empty(t *testing.T) {
	jd := GetGuaJieDu("")
	if jd != nil {
		t.Error("GetGuaJieDu should return nil for empty string")
	}
}

func TestGetGuaJieDuByCategory(t *testing.T) {
	tests := []struct {
		cat      JieDuCategory
		name     string
		expected string
	}{
		{JieDuShiYe, "事业", "事业"},
		{JieDuAiQing, "爱情", "爱情"},
		{JieDuCaiYun, "财运", "财运"},
		{JieDuKaoShi, "考试", "考试"},
		{JieDuJianKang, "健康", "健康"},
		{JieDuChuXing, "出行", "出行"},
		{JieDuGuanSi, "官司", "官司"},
		{JieDuJiaZhai, "家宅", "家宅"},
	}

	for _, tt := range tests {
		result := GetGuaJieDuByCategory("乾乾", tt.cat)
		if result == "" {
			t.Errorf("GetGuaJieDuByCategory(乾乾, %s) returned empty", tt.name)
		}
		t.Logf("%s: %s", tt.name, result)
	}
}

func TestGetGuaJieDuByCategory_NotFound(t *testing.T) {
	result := GetGuaJieDuByCategory("不存在的卦", JieDuShiYe)
	if result != "" {
		t.Errorf("Expected empty string for non-existent hexagram, got %s", result)
	}
}

func TestGetGuaJieDuByCategory_InvalidCategory(t *testing.T) {
	result := GetGuaJieDuByCategory("乾乾", JieDuCategory("未知类别"))
	if result != "" {
		t.Errorf("Expected empty string for invalid category, got %s", result)
	}
}

func TestAllJieDuEntries(t *testing.T) {
	entries := AllJieDuEntries()
	if entries == nil {
		t.Fatal("AllJieDuEntries returned nil")
	}
	if len(entries) == 0 {
		t.Error("AllJieDuEntries returned empty slice")
	}

	for _, e := range entries {
		if e.Ming == "" {
			t.Error("Entry has empty Ming")
		}
		t.Logf("卦: %s", e.Ming)
	}
	t.Logf("Total entries: %d", len(entries))
}

func TestAllJieDuCategories(t *testing.T) {
	cats := AllJieDuCategories
	if len(cats) != 8 {
		t.Errorf("Expected 8 categories, got %d", len(cats))
	}

	expected := []JieDuCategory{
		JieDuShiYe, JieDuAiQing, JieDuCaiYun, JieDuKaoShi,
		JieDuJianKang, JieDuChuXing, JieDuGuanSi, JieDuJiaZhai,
	}
	for i, cat := range cats {
		if cat != expected[i] {
			t.Errorf("Category[%d] = %s, want %s", i, cat, expected[i])
		}
	}
}

func TestGetGuaJieDu_SampleHexagrams(t *testing.T) {
	samples := []string{"乾乾", "坤坤", "坎坎", "离离", "震震", "艮艮"}

	for _, idx := range samples {
		jd := GetGuaJieDu(idx)
		if jd == nil {
			t.Errorf("GetGuaJieDu(%s) returned nil", idx)
			continue
		}
		t.Logf("%s: 事业=%s", idx, truncate(jd.ShiYe, 30))
	}
}
