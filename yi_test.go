package yi

import (
	"testing"
	"time"
)

func TestDivineByNumber(t *testing.T) {
	zy := DivineByNumber(0, 0)
	if zy == nil {
		t.Fatal("DivineByNumber returned nil")
	}
	if zy.GetGua(Ben) == nil {
		t.Fatal("Ben gua is nil")
	}
	if zy.GetGua(Ben).Ming != "乾为天" {
		t.Errorf("Expected 乾为天, got %s", zy.GetGua(Ben).Ming)
	}
}

func TestDivineByTime(t *testing.T) {
	zy := DivineByTime(2024, 6, 15, 10)
	if zy == nil {
		t.Fatal("DivineByTime returned nil")
	}
	if zy.GetGua(Ben) == nil {
		t.Fatal("Ben gua is nil")
	}
}

func TestDivineByTimeGua(t *testing.T) {
	params := TimeGuaParams{Year: 2024, Month: 6, Day: 15, Hour: 10}
	zy := DivineByTimeGua(params)
	if zy == nil {
		t.Fatal("DivineByTimeGua returned nil")
	}
}

func TestDivineByTimeGuaWithSeed(t *testing.T) {
	params := TimeGuaParams{Year: 2024, Month: 6, Day: 15, Hour: 10}
	zy := DivineByTimeGua(params, "testuser")
	if zy == nil {
		t.Fatal("DivineByTimeGua with seed returned nil")
	}
}

func TestDivineByCoins(t *testing.T) {
	zy, results := DivineByCoins(12345)
	if zy == nil {
		t.Fatal("DivineByCoins returned nil zy")
	}
	for i, r := range results {
		if r < CoinYinYinYin || r > CoinYangYangYang {
			t.Errorf("Coin result %d out of range: %d", i, r)
		}
	}
}

func TestDivineByDayan(t *testing.T) {
	zy, results := DivineByDayan(12345)
	if zy == nil {
		t.Fatal("DivineByDayan returned nil zy")
	}
	for i, r := range results {
		if r.YaoValue < 6 || r.YaoValue > 9 {
			t.Errorf("Dayan result %d yao value out of range: %d", i, r.YaoValue)
		}
	}
}

func TestDivineByMeihua(t *testing.T) {
	zy := DivineByMeihua(10, 5, 2)
	if zy == nil {
		t.Fatal("DivineByMeihua returned nil")
	}
}

func TestDivineByMeihuaTime(t *testing.T) {
	tm := time.Date(2024, 6, 15, 10, 0, 0, 0, time.Local)
	zy, y, m, d, h := DivineByMeihuaTime(tm)
	if zy == nil {
		t.Fatal("DivineByMeihuaTime returned nil")
	}
	if y != 2024 || m != 6 || d != 15 {
		t.Errorf("Unexpected date: %d-%d-%d", y, m, d)
	}
	_ = h
}

func TestDivineByDailyHexagram(t *testing.T) {
	zy := DivineByDailyHexagram(2024, 6, 15, "user123")
	if zy == nil {
		t.Fatal("DivineByDailyHexagram returned nil")
	}
}

func TestJieGua(t *testing.T) {
	zy := DivineByNumber(0, 0)
	result := JieGua(zy, Male)
	if result == nil {
		t.Fatal("JieGua returned nil")
	}
	if result.BenGuaInfo == nil {
		t.Fatal("BenGuaInfo is nil")
	}
	if result.BenGuaInfo.Ming != "乾为天" {
		t.Errorf("Expected 乾为天, got %s", result.BenGuaInfo.Ming)
	}
}

func TestJieGuaWithLang(t *testing.T) {
	zy := DivineByNumber(0, 0)
	result := JieGuaWithLang(zy, Male, LangEN)
	if result == nil {
		t.Fatal("JieGuaWithLang returned nil")
	}
}

func TestFormatJieGua(t *testing.T) {
	zy := DivineByNumber(0, 0)
	result := JieGua(zy, Male)
	output := FormatJieGua(result)
	if output == "" {
		t.Fatal("FormatJieGua returned empty string")
	}
}

func TestFormatJieGuaWithLang(t *testing.T) {
	zy := DivineByNumber(0, 0)
	result := JieGuaWithLang(zy, Male, LangEN)
	output := FormatJieGuaWithLang(result, LangEN)
	if output == "" {
		t.Fatal("FormatJieGuaWithLang returned empty string")
	}
}

func TestGetDayan(t *testing.T) {
	d, err := GetDayan(1)
	if err != nil {
		t.Fatalf("GetDayan(1) error: %v", err)
	}
	if d.Number != 1 {
		t.Errorf("Expected number 1, got %d", d.Number)
	}
}

func TestMustGetDayan(t *testing.T) {
	d := MustGetDayan(1)
	if d.Number != 1 {
		t.Errorf("Expected number 1, got %d", d.Number)
	}
}

func TestGetDayanOutOfRange(t *testing.T) {
	_, err := GetDayan(0)
	if err == nil {
		t.Error("Expected error for GetDayan(0)")
	}
	_, err = GetDayan(82)
	if err == nil {
		t.Error("Expected error for GetDayan(82)")
	}
}

func TestGetWenYan(t *testing.T) {
	zy := DivineByNumber(0, 0)
	g := zy.GetGua(Ben)
	wd := GetWenYan(g)
	if wd == nil {
		t.Log("No WenYan data for this gua (expected for some)")
	}
}

func TestWuXingByBagua(t *testing.T) {
	wx := GetWuXingByBagua(Qian)
	if wx != Metal {
		t.Errorf("Expected Metal for Qian, got %d", wx)
	}
	wx = GetWuXingByBagua(Li)
	if wx != Fire {
		t.Errorf("Expected Fire for Li, got %d", wx)
	}
}

func TestTypeAliases(t *testing.T) {
	var _ Hexagram = Gua{}
	var _ Line = Yao{}
	var _ Trigram = Bagua(0)
}

func TestConstants(t *testing.T) {
	if Ben != 0 || Bian != 1 || Hu != 2 || Cuo != 3 || Zong != 4 {
		t.Error("Gua type constants mismatch")
	}
	if Qian != 0 || Kun != 7 {
		t.Error("Bagua constants mismatch")
	}
	if Male == Female {
		t.Error("Male and Female should be different")
	}
	if Wood == Fire {
		t.Error("Wood and Fire should be different")
	}
}

func TestSexMethods(t *testing.T) {
	if !Male.IsMale() {
		t.Error("Male should be male")
	}
	if Male.IsFemale() {
		t.Error("Male should not be female")
	}
	if !Female.IsFemale() {
		t.Error("Female should be female")
	}
}

func TestIsJi(t *testing.T) {
	zy := DivineByNumber(0, 0)
	result := JieGua(zy, Male)
	_ = result.IsJi
}

func TestAllGuaTypes(t *testing.T) {
	zy := DivineByNumber(0, 0)
	for _, gt := range []int{Ben, Bian, Hu, Cuo, Zong} {
		g := zy.GetGua(gt)
		if g == nil {
			t.Errorf("Gua type %d is nil", gt)
		}
	}
}

func TestDivineByLunarTime(t *testing.T) {
	zy := DivineByLunarTime(2024, 6, 15, 5)
	if zy == nil {
		t.Fatal("DivineByLunarTime returned nil")
	}
}

func TestTranslateJiXiong(t *testing.T) {
	if TranslateJiXiong("吉", LangEN) != "Auspicious" {
		t.Error("Translation mismatch for 吉")
	}
	if TranslateJiXiong("凶", LangEN) != "Inauspicious" {
		t.Error("Translation mismatch for 凶")
	}
}

func TestTranslateCategory(t *testing.T) {
	if TranslateCategory("事业", LangEN) != "Career" {
		t.Error("Translation mismatch for 事业")
	}
}
