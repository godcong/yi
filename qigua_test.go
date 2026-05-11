package yi

import (
	"fmt"
	"testing"
	"time"
)

func TestCoinResult(t *testing.T) {
	tests := []struct {
		cr     CoinResult
		want   string
		isYang bool
		isChg  bool
	}{
		{CoinYangYangYang, "\u8001\u9633", true, true},
		{CoinYinYangYang, "\u5c11\u9633", true, false},
		{CoinYinYinYang, "\u5c11\u9634", false, false},
		{CoinYinYinYin, "\u8001\u9634", false, true},
	}

	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := tt.cr.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
			if got := tt.cr.IsYang(); got != tt.isYang {
				t.Errorf("IsYang() = %v, want %v", got, tt.isYang)
			}
			if got := tt.cr.IsChanging(); got != tt.isChg {
				t.Errorf("IsChanging() = %v, want %v", got, tt.isChg)
			}
		})
	}
}

func TestDivineByCoinValues(t *testing.T) {
	// Test: all 9s = all changing = Qian->Kun
	values := [6]int{9, 9, 9, 9, 9, 9}
	zy, err := DivineByCoinValues(values)
	if err != nil {
		t.Fatalf("DivineByCoinValues failed: %v", err)
	}
	if zy == nil {
		t.Fatal("DivineByCoinValues returned nil")
	}
	gua := zy.GetGua(Ben)
	if gua == nil {
		t.Fatal("GetGua(Ben) returned nil")
	}
	bianGua := zy.GetGua(Bian)
	fmt.Printf("Coin: %s -> %s\n", gua.Ming, bianGua.Ming)
	if gua.Index != "乾乾" {
		t.Errorf("want 乾乾, got %s", gua.Index)
	}

	// Test: all 8s = no changing
	values8 := [6]int{8, 8, 8, 8, 8, 8}
	zy8, err := DivineByCoinValues(values8)
	if err != nil {
		t.Fatalf("DivineByCoinValues 8s failed: %v", err)
	}
	gua8 := zy8.GetGua(Ben)
	bian := zy8.GetAllBianYao()
	fmt.Printf("Coin8: %s (bianYao: %v)\n", gua8.Ming, bian)
	if len(bian) != 0 {
		t.Errorf("all 8s should have no changing lines, got %v", bian)
	}

	// Test invalid value
	bad := [6]int{5, 8, 8, 8, 8, 8}
	_, err = DivineByCoinValues(bad)
	if err == nil {
		t.Error("expected error for invalid value 5, got nil")
	}
}

func TestDivineByCoins(t *testing.T) {
	zy, results := DivineByCoins(42)
	if zy == nil {
		t.Fatal("DivineByCoins returned nil")
	}
	gua := zy.GetGua(Ben)
	fmt.Printf("Coins seed=42: %s, bianYao=%v\n", gua.Ming, zy.GetAllBianYao())
	for i, cr := range results {
		fmt.Printf("  第%d爻: %s (值=%d, 动=%v)\n", i+1, cr, cr.YaoValue(), cr.IsChanging())
	}
}

func TestDivineByCoinsDeterministic(t *testing.T) {
	zy1, _ := DivineByCoins(12345)
	zy2, _ := DivineByCoins(12345)
	n1 := zy1.GetGua(Ben).Ming
	n2 := zy2.GetGua(Ben).Ming
	if n1 != n2 {
		t.Errorf("same seed should produce same result: %s vs %s", n1, n2)
	}
}

func TestShichen(t *testing.T) {
	tests := []struct {
		hour int
		want int
	}{
		{0, 1},  // 子时
		{1, 2},  // 丑时
		{2, 2},  // 丑时
		{3, 3},  // 寅时
		{12, 7}, // 午时
		{13, 8}, // 未时
		{23, 1}, // 子时
	}
	for _, tt := range tests {
		got := shichen(tt.hour)
		if got != tt.want {
			t.Errorf("shichen(%d) = %d, want %d", tt.hour, got, tt.want)
		}
	}
}

func TestDivineByMeihua(t *testing.T) {
	// 上卦=1(乾), 下卦=1(乾) -> 乾为天
	zy := DivineByMeihua(1, 1, -1)
	gua := zy.GetGua(Ben)
	fmt.Printf("Meihua(1,1): %s\n", gua.Ming)
	if gua.Index != "乾乾" {
		t.Errorf("want 乾乾, got %s", gua.Index)
	}

	// Test with larger numbers (wrap-around)
	zy2 := DivineByMeihua(9, 9, 2)
	gua2 := zy2.GetGua(Ben)
	fmt.Printf("Meihua(9,9,2): %s (bianYao: %v)\n", gua2.Ming, zy2.GetAllBianYao())
}

func TestDivineByMeihuaTime(t *testing.T) {
	zy, y, m, d, h := DivineByMeihuaTime(testTime(2024, 1, 15, 14))
	fmt.Printf("MeihuaTime(2024-01-15-14): %s, 年=%d 月=%d 日=%d 时辰=%d\n",
		zy.GetGua(Ben).Ming, y, m, d, h)
}

func TestDivineByTimeGua(t *testing.T) {
	zy := DivineByTimeGua(TimeGuaParams{Year: 2024, Month: 1, Day: 15, Hour: 14})
	fmt.Printf("TimeGua(2024,1,15,14): %s, bianYao=%v\n",
		zy.GetGua(Ben).Ming, zy.GetAllBianYao())

	zy2 := DivineByCurrentTime()
	fmt.Printf("CurrentTime: %s\n", zy2.GetGua(Ben).Ming)
}

func TestDivineByLunarTime(t *testing.T) {
	// 农历2024年正月初一子时
	zy := DivineByLunarTime(2024, 1, 1, 1)
	fmt.Printf("LunarTime(2024-1-1-1): %s\n", zy.GetGua(Ben).Ming)
}

func TestDivineByDayanValues(t *testing.T) {
	// Test: all 7s = no changing
	values := [6]int{7, 7, 7, 7, 7, 7}
	zy, err := DivineByDayanValues(values)
	if err != nil {
		t.Fatalf("DivineByDayanValues failed: %v", err)
	}
	fmt.Printf("Dayan(7s): %s (bianYao: %v)\n",
		zy.GetGua(Ben).Ming, zy.GetAllBianYao())

	// Test invalid
	bad := [6]int{5, 7, 7, 7, 7, 7}
	_, err = DivineByDayanValues(bad)
	if err == nil {
		t.Error("expected error for invalid Dayan value 5")
	}
}

func TestDivineByDayan(t *testing.T) {
	zy, results := DivineByDayan(42)
	if zy == nil {
		t.Fatal("DivineByDayan returned nil")
	}
	fmt.Printf("Dayan seed=42: %s, bianYao=%v\n",
		zy.GetGua(Ben).Ming, zy.GetAllBianYao())
	for i, dr := range results {
		fmt.Printf("  第%d爻: 值=%d (动=%v), 三揲剩余=%d\n",
			i+1, dr.YaoValue, dr.IsChanging, dr.Remaining)
	}
}

// testTime creates a time.Time for testing.
func testTime(year, month, day, hour int) time.Time {
	tt, _ := time.Parse("2006-01-02 15:04", fmt.Sprintf("%04d-%02d-%02d %02d:00", year, month, day, hour))
	return tt
}
