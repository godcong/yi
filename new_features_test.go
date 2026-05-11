package yi

import "testing"

func TestTuanData(t *testing.T) {
	g, err := GetGuaByIndex("乾乾")
	if err != nil {
		t.Fatal(err)
	}
	tuan := g.GetTuan()
	if tuan == "" {
		t.Error("Expected non-empty Tuan for 乾卦")
	}
	t.Logf("乾卦彖辞: %s...", truncate(tuan, 30))
}

func TestXiangData(t *testing.T) {
	g, err := GetGuaByIndex("乾乾")
	if err != nil {
		t.Fatal(err)
	}
	xiang := g.GetXiang()
	if xiang == "" {
		t.Error("Expected non-empty Xiang for 乾卦")
	}
	t.Logf("乾卦象辞: %s", xiang)

	g2, _ := GetGuaByIndex("坤坤")
	t.Logf("坤卦象辞: %s", g2.GetXiang())
}

func TestWenYanData(t *testing.T) {
	g, err := GetGuaByIndex("乾乾")
	if err != nil {
		t.Fatal(err)
	}
	if !g.HasWenYan() {
		t.Error("Expected WenYan for 乾卦")
	}
	wy := g.GetWenYan()
	if wy == nil || len(wy.Entries) == 0 {
		t.Error("Expected non-empty WenYan entries for 乾卦")
	}
	t.Logf("乾卦文言: %d entries", len(wy.Entries))

	g2, _ := GetGuaByIndex("坎坎")
	if g2.HasWenYan() {
		t.Error("Expected no WenYan for 坎卦")
	}
}

func TestJiaZiData(t *testing.T) {
	jz, err := GetJiaZi(1)
	if err != nil {
		t.Fatal(err)
	}
	if jz.Name != "甲子" {
		t.Errorf("Expected 甲子, got %s", jz.Name)
	}
	t.Logf("甲子: %s, 纳音: %s, 阴阳: %s", jz.Name, jz.NaYin, jz.YinYang)

	jz60, _ := GetJiaZi(60)
	t.Logf("癸亥: %s, 纳音: %s", jz60.Name, jz60.NaYin)
}

func TestLiuQin(t *testing.T) {
	gongWX := GetGuaGongWuXing(Qian)
	if gongWX != Metal {
		t.Errorf("Expected Metal for Qian palace, got %v", gongWX)
	}

	lq := GetLiuQin(Metal, Metal)
	if lq != LQXiongDi {
		t.Errorf("Expected 兄弟 for Metal-Metal, got %s", lq)
	}
	lq2 := GetLiuQin(Metal, Water)
	if lq2 != LQZiSun {
		t.Errorf("Expected 子孙 for Metal-Water, got %s", lq2)
	}
	t.Logf("乾宫六亲: 金→%s, 水→%s, 木→%s, 土→%s, 火→%s",
		GetLiuQin(Metal, Metal), GetLiuQin(Metal, Water),
		GetLiuQin(Metal, Wood), GetLiuQin(Metal, Earth), GetLiuQin(Metal, Fire))
}

func TestShiYing(t *testing.T) {
	g, err := GetGuaByIndex("乾乾")
	if err != nil {
		t.Fatal(err)
	}
	sy := g.GetShiYing()
	if sy == nil {
		t.Fatal("Expected ShiYing for 乾卦")
	}
	if sy.ShiPos != Shang {
		t.Errorf("Expected Shi at Shang for BenGong, got %d", sy.ShiPos)
	}
	if sy.YingPos != San {
		t.Errorf("Expected Ying at San for BenGong, got %d", sy.YingPos)
	}
	t.Logf("乾卦: 世%d应%d, 位置: %s", sy.ShiPos, sy.YingPos, sy.Position)

	g2, _ := GetGuaByIndex("坤坤")
	sy2 := g2.GetShiYing()
	t.Logf("坤卦: 世%d应%d, 位置: %s", sy2.ShiPos, sy2.YingPos, sy2.Position)
}

func TestTianGanDiZhi(t *testing.T) {
	if GetTianGanName(TGJia) != "甲" {
		t.Error("Expected 甲 for TGJia")
	}
	if GetDiZhiName(DZZi) != "子" {
		t.Error("Expected 子 for DZZi")
	}
	if GetTianGanWuXing(TGJia) != Wood {
		t.Error("Expected Wood for 甲")
	}
	if GetDiZhiWuXing(DZSi) != Fire {
		t.Error("Expected Fire for 巳")
	}
	t.Logf("甲→%s, 子→%s, 甲五行→%s, 巳五行→%s",
		GetTianGanName(TGJia), GetDiZhiName(DZZi),
		GetTianGanWuXing(TGJia), GetDiZhiWuXing(DZSi))
}
