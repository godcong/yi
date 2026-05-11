package yi

import (
	"strings"
	"testing"
)

func TestJieGuaByCoins(t *testing.T) {
	zy, coins := DivineByCoins(42)
	result := JieGua(zy, Male)
	if result == nil {
		t.Fatal("JieGua returned nil")
	}
	if result.ZhouYi == nil {
		t.Error("JieGuaResult.ZhouYi is nil")
	}
	if result.BenGuaInfo == nil {
		t.Error("JieGuaResult.BenGuaInfo is nil")
	}
	if result.BenGuaInfo.Ming == "" {
		t.Error("BenGuaInfo.Ming is empty")
	}
	if result.BianGuaInfo == nil {
		t.Error("JieGuaResult.BianGuaInfo is nil")
	}

	// Verify changing lines were detected
	t.Logf("Coins: %v", coins)
	t.Logf("本卦: %s (%s)", result.BenGuaInfo.Ming, result.BenGuaInfo.JiXiong)
	t.Logf("变卦: %s (%s)", result.BianGuaInfo.Ming, result.BianGuaInfo.JiXiong)
}

func TestJieGuaByNumber(t *testing.T) {
	zy := DivineByNumber(Qian, Qian, 1) // Qian with changing line at position 1
	result := JieGua(zy, Male)

	if result.BenGuaInfo == nil {
		t.Fatal("BenGuaInfo is nil")
	}
	if result.BenGuaInfo.GuaName != "乾" {
		t.Errorf("Expected GuaName=乾, got %s", result.BenGuaInfo.GuaName)
	}
	if result.BenGuaInfo.JiXiong == "" {
		t.Error("BenGuaInfo.JiXiong is empty")
	}
	if result.DongYaoText == "" {
		t.Error("DongYaoText is empty — expected yao ci for changing line")
	}
	t.Logf("动爻: %s", result.DongYaoText)
	t.Logf("吉凶: %v, 原因: %s", result.IsJi, result.JiXiongReason)
}

func TestJieGuaFemale(t *testing.T) {
	zy := DivineByNumber(Qian, Qian, 0)
	resultMale := JieGua(zy, Male)
	resultFemale := JieGua(zy, Female)

	// Female may get different NvMing judgment
	t.Logf("Male: Ji=%v, DongYaoJiXiong=%s", resultMale.IsJi, resultMale.DongYaoJiXiong)
	t.Logf("Female: Ji=%v, DongYaoJiXiong=%s", resultFemale.IsJi, resultFemale.DongYaoJiXiong)
}

func TestJieGuaNilZhouYi(t *testing.T) {
	result := JieGua(nil, Male)
	if result != nil {
		t.Error("Expected nil for nil ZhouYi")
	}
}

func TestJieGuaAllFiveGua(t *testing.T) {
	zy := DivineByNumber(Zhen, Kan, 2)
	result := JieGua(zy, Male)

	if result.BenGuaInfo == nil {
		t.Error("BenGuaInfo is nil")
	}
	if result.BianGuaInfo == nil {
		t.Error("BianGuaInfo is nil")
	}
	if result.HuGuaInfo == nil {
		t.Error("HuGuaInfo is nil")
	}
	if result.CuoGuaInfo == nil {
		t.Error("CuoGuaInfo is nil")
	}
	if result.ZongGuaInfo == nil {
		t.Error("ZongGuaInfo is nil")
	}

	t.Logf("本卦: %s %s", result.BenGuaInfo.Symbol, result.BenGuaInfo.Ming)
	t.Logf("变卦: %s %s", result.BianGuaInfo.Symbol, result.BianGuaInfo.Ming)
	t.Logf("互卦: %s %s", result.HuGuaInfo.Symbol, result.HuGuaInfo.Ming)
	t.Logf("错卦: %s %s", result.CuoGuaInfo.Symbol, result.CuoGuaInfo.Ming)
	t.Logf("综卦: %s %s", result.ZongGuaInfo.Symbol, result.ZongGuaInfo.Ming)
}

func TestJieGuaShiYingInfo(t *testing.T) {
	zy := DivineByNumber(Qian, Qian)
	result := JieGua(zy, Male)

	if result.BenGuaInfo.GuaGong == "" {
		t.Error("GuaGong is empty")
	}
	if result.BenGuaInfo.Position == "" {
		t.Error("Position is empty")
	}
	if result.BenGuaInfo.ShiYao == "" {
		t.Error("ShiYao is empty")
	}
	if result.BenGuaInfo.YingYao == "" {
		t.Error("YingYao is empty")
	}

	// Qian is BenGong (本宫), so Shi at Shang, Ying at San
	if result.BenGuaInfo.ShiYao != "上爻" {
		t.Errorf("Expected ShiYao=上爻 for BenGong, got %s", result.BenGuaInfo.ShiYao)
	}
	if result.BenGuaInfo.YingYao != "三爻" {
		t.Errorf("Expected YingYao=三爻 for BenGong, got %s", result.BenGuaInfo.YingYao)
	}

	t.Logf("归属: %s·%s, 世: %s, 应: %s",
		result.BenGuaInfo.GuaGong, result.BenGuaInfo.Position,
		result.BenGuaInfo.ShiYao, result.BenGuaInfo.YingYao)
}

func TestJieGuaJiXiongReason(t *testing.T) {
	zy := DivineByNumber(Qian, Qian, 0)
	result := JieGua(zy, Male)

	if result.JiXiongReason == "" {
		t.Error("JiXiongReason is empty")
	}
	// Should contain 本卦 and 变卦
	if !strings.Contains(result.JiXiongReason, "本卦") {
		t.Error("JiXiongReason should contain '本卦'")
	}
	if !strings.Contains(result.JiXiongReason, "变卦") {
		t.Error("JiXiongReason should contain '变卦'")
	}
	t.Logf("Reason: %s", result.JiXiongReason)
}

func TestFormatJieGua(t *testing.T) {
	zy, _ := DivineByCoins(123)
	result := JieGua(zy, Male)
	output := FormatJieGua(result)

	if output == "" {
		t.Error("FormatJieGua returned empty string")
	}
	if !strings.Contains(output, "解  卦  结  果") {
		t.Error("Output should contain header")
	}
	if !strings.Contains(output, "本卦") {
		t.Error("Output should contain '本卦'")
	}
	if !strings.Contains(output, "变卦") {
		t.Error("Output should contain '变卦'")
	}
	if !strings.Contains(output, "综合判断") {
		t.Error("Output should contain '综合判断'")
	}

	t.Logf("FormatJieGua output:\n%s", output)
}

func TestFormatJieGuaNil(t *testing.T) {
	output := FormatJieGua(nil)
	if output != "解卦结果为空" {
		t.Errorf("Expected '解卦结果为空', got %q", output)
	}
}

func TestJieGuaByMeihuaTime(t *testing.T) {
	zy, _, _, _, _ := DivineByMeihuaTime(testTime(2024, 6, 15, 10))
	result := JieGua(zy, Male)
	if result == nil {
		t.Fatal("JieGua returned nil for MeihuaTime")
	}
	if result.BenGuaInfo == nil {
		t.Fatal("BenGuaInfo is nil")
	}
	t.Logf("梅花易数: 本卦=%s, 变卦=%s", result.BenGuaInfo.Ming, result.BianGuaInfo.Ming)
	t.Logf("解卦:\n%s", FormatJieGua(result))
}

func TestJieGuaFenXi(t *testing.T) {
	zy := DivineByNumber(Qian, Qian, 0) // 乾为天，初爻动
	result := JieGua(zy, Male)

	if len(result.FenXi) == 0 {
		t.Fatal("FenXi is empty")
	}

	// Check that at least some categories were extracted
	foundCats := make(map[FenXiCategory]bool)
	for _, fx := range result.FenXi {
		foundCats[fx.Category] = true
		if fx.Content == "" {
			t.Errorf("FenXi %s has empty Content", fx.Category)
		}
		if fx.JiXiong == "" {
			t.Errorf("FenXi %s has empty JiXiong", fx.Category)
		}
		if fx.Source == "" {
			t.Errorf("FenXi %s has empty Source", fx.Category)
		}
		t.Logf("%s（%s）%s [来源: %s]", fx.Category, fx.JiXiong, fx.Content, fx.Source)
	}

	// 乾为天 has 事业 and other keywords in GuaYi and Yao.Ci
	if !foundCats[FenXiShiYe] {
		t.Error("Expected 事业 category for 乾为天")
	}
}

func TestJieGuaFenXiFormat(t *testing.T) {
	zy, _, _, _, _ := DivineByMeihuaTime(testTime(2024, 6, 15, 10))
	result := JieGua(zy, Male)
	output := FormatJieGua(result)

	if !strings.Contains(output, "释义") {
		t.Error("FormatJieGua should contain '释义' section")
	}

	// Check that FenXi categories appear in formatted output
	for _, fx := range result.FenXi {
		if !strings.Contains(output, string(fx.Category)) {
			t.Errorf("Output should contain category %s", fx.Category)
		}
	}
}

func TestJieGuaByDayan(t *testing.T) {
	zy, _ := DivineByDayan(999)
	result := JieGua(zy, Female)
	if result == nil {
		t.Fatal("JieGua returned nil for Dayan")
	}
	t.Logf("大衍法: 本卦=%s (%s)", result.BenGuaInfo.Ming, result.BenGuaInfo.JiXiong)
}
