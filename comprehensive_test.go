package yi

import (
	"testing"
)

// ============================================================================
// Gua struct field completeness tests (renamed fields)
// ============================================================================

func TestAllGuaHaveRequiredFields(t *testing.T) {
	for xu := 1; xu <= 64; xu++ {
		g, err := GetGuaByXu(xu)
		if err != nil {
			t.Fatalf("GetGuaByXu(%d) failed: %v", xu, err)
		}
		if g.Index == "" {
			t.Errorf("xu=%d: Index is empty", xu)
		}
		if g.Ming == "" {
			t.Errorf("xu=%d: Ming is empty", xu)
		}
		if g.GuaName == "" {
			t.Errorf("xu=%d (%s): GuaName is empty", xu, g.Ming)
		}
		if g.ShangNum < 0 || g.ShangNum > 7 {
			t.Errorf("xu=%d: ShangNum=%d out of range [0,7]", xu, g.ShangNum)
		}
		if g.XiaNum < 0 || g.XiaNum > 7 {
			t.Errorf("xu=%d: XiaNum=%d out of range [0,7]", xu, g.XiaNum)
		}
		if g.ShangMing == "" {
			t.Errorf("xu=%d (%s): ShangMing is empty", xu, g.Ming)
		}
		if g.XiaMing == "" {
			t.Errorf("xu=%d (%s): XiaMing is empty", xu, g.Ming)
		}
		if g.GuaSymbol == "" {
			t.Errorf("xu=%d (%s): GuaSymbol is empty", xu, g.Ming)
		}
		if g.GuaYi == "" {
			t.Errorf("xu=%d (%s): GuaYi is empty", xu, g.Ming)
		}
	}
}

func TestAllGuaHaveJiXiong(t *testing.T) {
	// After fix: all 64 should have JiXiong
	emptyCount := 0
	for xu := 1; xu <= 64; xu++ {
		g, _ := GetGuaByXu(xu)
		if g.JiXiong == "" {
			emptyCount++
			t.Errorf("xu=%d (%s): JiXiong is empty", xu, g.Ming)
		}
		// Validate JiXiong values
		validJX := map[string]bool{"吉": true, "凶": true, "半吉": true}
		if g.JiXiong != "" && !validJX[g.JiXiong] {
			t.Errorf("xu=%d (%s): JiXiong=%q is not a valid value", xu, g.Ming, g.JiXiong)
		}
	}
	if emptyCount > 0 {
		t.Errorf("%d/64 gua have empty JiXiong", emptyCount)
	}
}

func TestJiXiongDistribution(t *testing.T) {
	// Should have both 吉 and 凶 in the 64 gua
	ji, xiong, banji := 0, 0, 0
	for xu := 1; xu <= 64; xu++ {
		g, _ := GetGuaByXu(xu)
		switch g.JiXiong {
		case "吉":
			ji++
		case "凶":
			xiong++
		case "半吉":
			banji++
		}
	}
	t.Logf("JiXiong distribution: 吉=%d, 凶=%d, 半吉=%d", ji, xiong, banji)
	if ji == 0 || xiong == 0 {
		t.Errorf("Unexpected distribution: 吉=%d, 凶=%d — both should exist", ji, xiong)
	}
}

func TestAllGuaHaveYaoData(t *testing.T) {
	for xu := 1; xu <= 64; xu++ {
		g, err := GetGuaByXu(xu)
		if err != nil {
			t.Fatalf("GetGuaByXu(%d) failed: %v", xu, err)
		}
		for pos := Chu; pos < YaoCount; pos++ {
			yao := g.GetYao(pos)
			if yao == nil {
				t.Errorf("xu=%d pos=%d: GetYao returned nil", xu, pos)
				continue
			}
			if yao.Ci == "" {
				t.Errorf("xu=%d (%s) pos=%d: Yao.Ci is empty", xu, g.Ming, pos)
			}
			if yao.JiXiong == "" {
				t.Errorf("xu=%d (%s) pos=%d: Yao.JiXiong is empty", xu, g.Ming, pos)
			}
		}
	}
}

func TestGuaEmptyDataReport(t *testing.T) {
	// Report empty fields (informational)
	var emptyYong, emptyYongJiXiong []string
	emptyNvMing := 0

	for xu := 1; xu <= 64; xu++ {
		g, _ := GetGuaByXu(xu)
		if g.Yong == "" {
			emptyYong = append(emptyYong, g.Ming)
		}
		if g.YongJiXiong == "" {
			emptyYongJiXiong = append(emptyYongJiXiong, g.Ming)
		}
		for pos := Chu; pos < YaoCount; pos++ {
			if yao := g.GetYao(pos); yao != nil && !yao.HasNvMing() {
				emptyNvMing++
			}
		}
	}

	t.Logf("Gua.Yong empty: %d/64 — %v (only Qian/Kun have 用九/用六)", len(emptyYong), emptyYong)
	t.Logf("Gua.YongJiXiong empty: %d/64", len(emptyYongJiXiong))
	t.Logf("Yao.NvMing empty: %d/384 (empty = general/unisex, not missing)", emptyNvMing)
}

// ============================================================================
// TuanText/XiangText data completeness (merged into Gua struct)
// ============================================================================

func TestAllGuaHaveTuanText(t *testing.T) {
	emptyCount := 0
	for xu := 1; xu <= 64; xu++ {
		g, _ := GetGuaByXu(xu)
		if g.TuanText == "" {
			emptyCount++
			t.Errorf("xu=%d (%s): TuanText is empty", xu, g.Ming)
		}
	}
	if emptyCount > 0 {
		t.Errorf("%d/64 gua have empty TuanText", emptyCount)
	}
}

func TestAllGuaHaveXiangText(t *testing.T) {
	emptyCount := 0
	for xu := 1; xu <= 64; xu++ {
		g, _ := GetGuaByXu(xu)
		if g.XiangText == "" {
			emptyCount++
			t.Errorf("xu=%d (%s): XiangText is empty", xu, g.Ming)
		}
	}
	if emptyCount > 0 {
		t.Errorf("%d/64 gua have empty XiangText", emptyCount)
	}
}

func TestGetTuanMethodMatchesField(t *testing.T) {
	for xu := 1; xu <= 64; xu++ {
		g, _ := GetGuaByXu(xu)
		if g.GetTuan() != g.TuanText {
			t.Errorf("xu=%d: GetTuan()=%q != TuanText=%q", xu, g.GetTuan(), g.TuanText)
		}
	}
}

func TestGetXiangMethodMatchesField(t *testing.T) {
	for xu := 1; xu <= 64; xu++ {
		g, _ := GetGuaByXu(xu)
		if g.GetXiang() != g.XiangText {
			t.Errorf("xu=%d: GetXiang()=%q != XiangText=%q", xu, g.GetXiang(), g.XiangText)
		}
	}
}

// ============================================================================
// WenYan tests
// ============================================================================

func TestWenYanOnlyQianKun(t *testing.T) {
	qian, _ := GetGuaByIndex("乾乾")
	kun, _ := GetGuaByIndex("坤坤")
	li, _ := GetGuaByIndex("离离")

	if !qian.HasWenYan() {
		t.Error("Qian should have WenYan")
	}
	if !kun.HasWenYan() {
		t.Error("Kun should have WenYan")
	}
	if li.HasWenYan() {
		t.Error("Li should NOT have WenYan")
	}
}

// ============================================================================
// GuaName (renamed from Xiang) tests
// ============================================================================

func TestGuaNameValues(t *testing.T) {
	// GuaName should be single-char hexagram name
	samples := []struct {
		index string
		want  string
	}{
		{"乾乾", "乾"}, {"坤坤", "坤"}, {"坎震", "屯"},
		{"艮坎", "蒙"}, {"坎乾", "需"}, {"乾坎", "讼"},
	}
	for _, tt := range samples {
		g, _ := GetGuaByIndex(tt.index)
		if g.GuaName != tt.want {
			t.Errorf("%s.GuaName=%q, want %q", tt.index, g.GuaName, tt.want)
		}
	}
}

// ============================================================================
// Bagua consistency tests
// ============================================================================

func TestBaguaIndexConsistency(t *testing.T) {
	for xu := 1; xu <= 64; xu++ {
		g, _ := GetGuaByXu(xu)
		shangName := GetBaguaName(g.ShangNum)
		xiaName := GetBaguaName(g.XiaNum)
		if shangName == "" {
			t.Errorf("xu=%d: ShangNum=%d has no BaguaName", xu, g.ShangNum)
		}
		if xiaName == "" {
			t.Errorf("xu=%d: XiaNum=%d has no BaguaName", xu, g.XiaNum)
		}
		expectedIndex := shangName + xiaName
		if g.Index != expectedIndex {
			t.Errorf("xu=%d: Index=%q, expected %q", xu, g.Index, expectedIndex)
		}
	}
}

func TestAllBaguaNames(t *testing.T) {
	expected := map[Bagua]string{
		Qian: "乾", Dui: "兑", Li: "离", Zhen: "震",
		Xun: "巽", Kan: "坎", Gen: "艮", Kun: "坤",
	}
	for bg, name := range expected {
		if got := GetBaguaName(bg); got != name {
			t.Errorf("GetBaguaName(%d) = %q, want %q", bg, got, name)
		}
	}
}

func TestAllBaguaSymbols(t *testing.T) {
	expected := map[Bagua]string{
		Qian: "☰", Dui: "☱", Li: "☲", Zhen: "☳",
		Xun: "☴", Kan: "☵", Gen: "☶", Kun: "☷",
	}
	for bg, sym := range expected {
		if got := GetBaguaSymbol(bg); got != sym {
			t.Errorf("GetBaguaSymbol(%d) = %q, want %q", bg, got, sym)
		}
	}
}

// ============================================================================
// WuXing relationship tests
// ============================================================================

func TestWuXingShengCycle(t *testing.T) {
	cycle := []WuXing{Wood, Fire, Earth, Metal, Water}
	for i, wx := range cycle {
		next := cycle[(i+1)%len(cycle)]
		if wx.Sheng() != next {
			t.Errorf("%s.Sheng() = %s, want %s", wx, wx.Sheng(), next)
		}
	}
}

func TestWuXingKeCycle(t *testing.T) {
	kePairs := []struct{ from, to WuXing }{
		{Wood, Earth}, {Earth, Water}, {Water, Fire}, {Fire, Metal}, {Metal, Wood},
	}
	for _, p := range kePairs {
		if p.from.Ke() != p.to {
			t.Errorf("%s.Ke() = %s, want %s", p.from, p.from.Ke(), p.to)
		}
	}
}

func TestWuXingBeiShengBeiKe(t *testing.T) {
	for _, wx := range []WuXing{Wood, Fire, Earth, Metal, Water} {
		generator := wx.BeiSheng()
		if generator.Sheng() != wx {
			t.Errorf("%s.BeiSheng()=%s, but %s.Sheng()=%s", wx, generator, generator, generator.Sheng())
		}
		controller := wx.BeiKe()
		if controller.Ke() != wx {
			t.Errorf("%s.BeiKe()=%s, but %s.Ke()=%s", wx, controller, controller, controller.Ke())
		}
	}
}

// ============================================================================
// JiaZi tests
// ============================================================================

func TestJiaZiAll60(t *testing.T) {
	for i := 1; i <= 60; i++ {
		jz, err := GetJiaZi(i)
		if err != nil {
			t.Fatalf("GetJiaZi(%d) failed: %v", i, err)
		}
		if jz.Name == "" {
			t.Errorf("JiaZi[%d]: Name is empty", i)
		}
		if jz.TianGan == "" {
			t.Errorf("JiaZi[%d]: TianGan is empty", i)
		}
		if jz.DiZhi == "" {
			t.Errorf("JiaZi[%d]: DiZhi is empty", i)
		}
	}
}

func TestJiaZiFirstAndLast(t *testing.T) {
	first, _ := GetJiaZi(1)
	if first.Name != "甲子" {
		t.Errorf("JiaZi[1].Name = %q, want 甲子", first.Name)
	}
	last, _ := GetJiaZi(60)
	if last.Name != "癸亥" {
		t.Errorf("JiaZi[60].Name = %q, want 癸亥", last.Name)
	}
}

func TestJiaZiInvalidIndex(t *testing.T) {
	for _, idx := range []int{0, -1, 61, 100} {
		_, err := GetJiaZi(idx)
		if err == nil {
			t.Errorf("GetJiaZi(%d) should return error", idx)
		}
	}
}

// ============================================================================
// LiuQin (六亲) tests
// ============================================================================

func TestLiuQinAllPalaces(t *testing.T) {
	palaceWX := map[Bagua]WuXing{
		Qian: Metal, Dui: Metal,
		Li:   Fire,
		Zhen: Wood, Xun: Wood,
		Kan: Water,
		Gen: Earth, Kun: Earth,
	}
	for bg, wx := range palaceWX {
		got := GetGuaGongWuXing(bg)
		if got != wx {
			t.Errorf("GetGuaGongWuXing(%s) = %s, want %s", GetBaguaName(bg), got, wx)
		}
	}
}

func TestLiuQinRelationships(t *testing.T) {
	tests := []struct {
		gongWX, yaoWX WuXing
		want          LiuQin
	}{
		{Metal, Metal, LQXiongDi},
		{Metal, Water, LQZiSun},
		{Metal, Wood, LQQiCai},
		{Metal, Earth, LQFuMu},
		{Metal, Fire, LQGuanGui},
		{Wood, Wood, LQXiongDi},
		{Wood, Fire, LQZiSun},
		{Fire, Fire, LQXiongDi},
		{Water, Water, LQXiongDi},
		{Earth, Earth, LQXiongDi},
	}
	for _, tt := range tests {
		got := GetLiuQin(tt.gongWX, tt.yaoWX)
		if got != tt.want {
			t.Errorf("GetLiuQin(%s,%s) = %s, want %s", tt.gongWX, tt.yaoWX, got, tt.want)
		}
	}
}

func TestLiuQinStrings(t *testing.T) {
	tests := []struct {
		lq   LiuQin
		want string
	}{
		{LQFuMu, "父母"}, {LQXiongDi, "兄弟"}, {LQQiCai, "妻财"},
		{LQZiSun, "子孙"}, {LQGuanGui, "官鬼"},
	}
	for _, tt := range tests {
		if tt.lq.String() != tt.want {
			t.Errorf("LiuQin(%d).String() = %q, want %q", tt.lq, tt.lq.String(), tt.want)
		}
	}
}

// ============================================================================
// ShiYing (世应) tests
// ============================================================================

func TestShiYingAllGua(t *testing.T) {
	for xu := 1; xu <= 64; xu++ {
		g, _ := GetGuaByXu(xu)
		sy := g.GetShiYing()
		if sy == nil {
			t.Errorf("xu=%d (%s): GetShiYing returned nil", xu, g.Ming)
			continue
		}
		if sy.ShiPos < 0 || sy.ShiPos >= YaoCount {
			t.Errorf("xu=%d (%s): ShiPos=%d out of range", xu, g.Ming, sy.ShiPos)
		}
		if sy.YingPos < 0 || sy.YingPos >= YaoCount {
			t.Errorf("xu=%d (%s): YingPos=%d out of range", xu, g.Ming, sy.YingPos)
		}
	}
}

func TestShiYingRule(t *testing.T) {
	for xu := 1; xu <= 64; xu++ {
		g, _ := GetGuaByXu(xu)
		sy := g.GetShiYing()
		if sy == nil {
			continue
		}
		diff := int(sy.ShiPos) - int(sy.YingPos)
		if diff < 0 {
			diff = -diff
		}
		if diff != 3 {
			t.Errorf("xu=%d (%s): ShiPos=%d YingPos=%d, diff=%d, want 3",
				xu, g.Ming, sy.ShiPos, sy.YingPos, diff)
		}
	}
}

func TestShiYingBenGong(t *testing.T) {
	pureGua := []string{"乾乾", "坤坤", "离离", "坎坎", "震震", "艮艮", "巽巽", "兑兑"}
	for _, idx := range pureGua {
		g, _ := GetGuaByIndex(idx)
		sy := g.GetShiYing()
		if sy.ShiPos != Shang {
			t.Errorf("%s: ShiPos=%d, want Shang(5)", idx, sy.ShiPos)
		}
		if sy.YingPos != San {
			t.Errorf("%s: YingPos=%d, want San(2)", idx, sy.YingPos)
		}
		if sy.Position != BenGong {
			t.Errorf("%s: Position=%s, want BenGong", idx, sy.Position)
		}
	}
}

// ============================================================================
// GuaGong (八宫) coverage test
// ============================================================================

func TestGuaGongEightPalaces(t *testing.T) {
	palaceCounts := make(map[Bagua]int)
	for xu := 1; xu <= 64; xu++ {
		g, _ := GetGuaByXu(xu)
		gong := g.GetGuaGong()
		if gong < 0 || gong > 7 {
			t.Errorf("xu=%d: invalid gong=%d", xu, gong)
			continue
		}
		palaceCounts[gong]++
	}
	if len(palaceCounts) != 8 {
		t.Errorf("expected 8 palaces, got %d", len(palaceCounts))
	}
	for gong, count := range palaceCounts {
		if count != 8 {
			t.Errorf("palace %s has %d gua, want 8", GetBaguaName(gong), count)
		}
	}
}

func TestGuaPositionCoverage(t *testing.T) {
	posCounts := make(map[GuaPosition]int)
	for xu := 1; xu <= 64; xu++ {
		g, _ := GetGuaByXu(xu)
		sy := g.GetShiYing()
		if sy != nil {
			posCounts[sy.Position]++
		}
	}
	for pos := BenGong; pos < GuaPositionCount; pos++ {
		if posCounts[pos] != 8 {
			t.Errorf("position %s has %d gua, want 8", pos, posCounts[pos])
		}
	}
}

func TestGuaPositionStrings(t *testing.T) {
	tests := []struct {
		pos  GuaPosition
		want string
	}{
		{BenGong, "本宫"}, {YiShi, "一世"}, {ErShi, "二世"},
		{SanShi, "三世"}, {SiShi, "四世"}, {WuShi, "五世"},
		{YouHun, "游魂"}, {GuiHun, "归魂"},
	}
	for _, tt := range tests {
		if tt.pos.String() != tt.want {
			t.Errorf("GuaPosition(%d).String() = %q, want %q", tt.pos, tt.pos.String(), tt.want)
		}
	}
}

// ============================================================================
// Transform tests
// ============================================================================

func TestCuoGuaOpposite(t *testing.T) {
	yi := Divine(0, 0)
	cuo := yi.GetGua(Cuo)
	if cuo == nil {
		t.Fatal("CuoGua is nil")
	}
	if cuo.Index != "坤坤" {
		t.Errorf("乾之错卦=%s, want 坤坤", cuo.Index)
	}
}

func TestZongGuaReversed(t *testing.T) {
	yi := DivineByNumber(5, 3)
	zong := yi.GetGua(Zong)
	if zong == nil {
		t.Fatal("ZongGua is nil")
	}
	if zong.GuaName != "蒙" {
		t.Errorf("屯之综卦 GuaName=%q, want 蒙", zong.GuaName)
	}
}

func TestHuGua(t *testing.T) {
	yi := Divine(0, 0)
	hu := yi.GetGua(Hu)
	if hu == nil {
		t.Fatal("HuGua is nil")
	}
}

func TestBianGua(t *testing.T) {
	yi := DivineByNumber(0, 0, 1)
	bian := yi.GetGua(Bian)
	if bian == nil {
		t.Fatal("BianGua is nil")
	}
}

// ============================================================================
// GetAllTuan / GetAllXiang tests
// ============================================================================

func TestGetAllTuanCompleteness(t *testing.T) {
	all := GetAllTuan()
	if len(all) != 64 {
		t.Errorf("GetAllTuan() returned %d, want 64", len(all))
	}
	for _, td := range all {
		if td.Index == "" {
			t.Error("Tuan entry has empty Index")
		}
		if td.Text == "" {
			t.Errorf("Tuan entry for %s has empty Text", td.Index)
		}
	}
}

func TestGetAllXiangCompleteness(t *testing.T) {
	all := GetAllXiang()
	if len(all) != 64 {
		t.Errorf("GetAllXiang() returned %d, want 64", len(all))
	}
	for _, xd := range all {
		if xd.Index == "" {
			t.Error("Xiang entry has empty Index")
		}
		if xd.Text == "" {
			t.Errorf("Xiang entry for %s has empty Text", xd.Index)
		}
	}
}

// ============================================================================
// Dayan tests
// ============================================================================

func TestDayanAll81(t *testing.T) {
	for n := 1; n <= 81; n++ {
		d, err := GetDayan(n)
		if err != nil {
			t.Fatalf("GetDayan(%d) failed: %v", n, err)
		}
		if d.Number != n {
			t.Errorf("Dayan[%d].Number = %d", n, d.Number)
		}
		if d.JiXiong == "" {
			t.Errorf("Dayan[%d].JiXiong is empty", n)
		}
	}
}

// ============================================================================
// Yao.HasNvMing tests
// ============================================================================

func TestYaoHasNvMing(t *testing.T) {
	g, _ := GetGuaByIndex("乾乾")
	// 乾初爻有女命
	yao0 := g.GetYao(Chu)
	if !yao0.HasNvMing() {
		t.Error("乾初爻 should have NvMing")
	}
}

func TestYaoNoNvMing(t *testing.T) {
	g, _ := GetGuaByIndex("坎震")
	// 坎震二爻无女命
	yao1 := g.GetYao(Er)
	if yao1.HasNvMing() {
		t.Error("水雷屯二爻 should NOT have NvMing")
	}
}

func TestNvMingCount(t *testing.T) {
	count := 0
	for xu := 1; xu <= 64; xu++ {
		g, _ := GetGuaByXu(xu)
		for pos := Chu; pos < YaoCount; pos++ {
			if yao := g.GetYao(pos); yao != nil && yao.HasNvMing() {
				count++
			}
		}
	}
	t.Logf("Yao with NvMing: %d/384", count)
	if count == 0 {
		t.Error("Expected at least some Yao to have NvMing")
	}
}

// ============================================================================
// Edge case tests
// ============================================================================

func TestGetGuaByXuBoundaries(t *testing.T) {
	_, err := GetGuaByXu(0)
	if err == nil {
		t.Error("GetGuaByXu(0) should error")
	}
	_, err = GetGuaByXu(65)
	if err == nil {
		t.Error("GetGuaByXu(65) should error")
	}
	_, err = GetGuaByXu(1)
	if err != nil {
		t.Errorf("GetGuaByXu(1) should succeed: %v", err)
	}
	_, err = GetGuaByXu(64)
	if err != nil {
		t.Errorf("GetGuaByXu(64) should succeed: %v", err)
	}
}

func TestGetYaoInvalidPosition(t *testing.T) {
	g, _ := GetGuaByIndex("乾乾")
	if yao := g.GetYao(-1); yao != nil {
		t.Error("GetYao(-1) should return nil")
	}
	if yao := g.GetYao(6); yao != nil {
		t.Error("GetYao(6) should return nil")
	}
}

func TestMustGetJiaZi(t *testing.T) {
	valid := MustGetJiaZi(1)
	if valid.Name != "甲子" {
		t.Errorf("MustGetJiaZi(1) = %q, want 甲子", valid.Name)
	}
}

// ============================================================================
// Gua struct renamed field backward compatibility
// ============================================================================

func TestGuaNameNotXiang(t *testing.T) {
	// Ensure GuaName is the single-char name, NOT the 大象辞
	g, _ := GetGuaByIndex("乾乾")
	if g.GuaName != "乾" {
		t.Errorf("GuaName=%q, want 乾 (single-char name)", g.GuaName)
	}
	if g.XiangText == "乾" {
		t.Error("XiangText should NOT be single-char name, it should be 大象辞")
	}
	if g.GetXiang() != g.XiangText {
		t.Error("GetXiang() should return XiangText")
	}
}

func TestShangMingXiaMingFormat(t *testing.T) {
	// ShangMing/XiaMing should be full Bagua names like "乾为天"
	g, _ := GetGuaByIndex("乾乾")
	if g.ShangMing != "乾为天" {
		t.Errorf("ShangMing=%q, want 乾为天", g.ShangMing)
	}
	if g.XiaMing != "乾为天" {
		t.Errorf("XiaMing=%q, want 乾为天", g.XiaMing)
	}
}

func TestGuaSymbolFormat(t *testing.T) {
	// GuaSymbol should be Unicode hexagram symbol
	g, _ := GetGuaByIndex("乾乾")
	if len(g.GuaSymbol) < 1 {
		t.Error("GuaSymbol is empty")
	}
}

// ============================================================================
// JiaZiInfo NaYin tests
// ============================================================================

func TestJiaZiNaYinField(t *testing.T) {
	jz, err := GetJiaZi(1)
	if err != nil {
		t.Fatal(err)
	}
	if jz.NaYin != "海中金" {
		t.Errorf("JiaZi[1].NaYin = %q, want 海中金", jz.NaYin)
	}

	jz30, _ := GetJiaZi(30)
	if jz30.NaYin != "长流水" {
		t.Errorf("JiaZi[30].NaYin = %q, want 长流水", jz30.NaYin)
	}

	jz60, _ := GetJiaZi(60)
	if jz60.NaYin != "大海水" {
		t.Errorf("JiaZi[60].NaYin = %q, want 大海水", jz60.NaYin)
	}
}

func TestJiaZiNaYinBackwardCompat(t *testing.T) {
	// WuXing() method should still work for backward compatibility
	jz, _ := GetJiaZi(1)
	if jz.WuXing() != "海中金" {
		t.Errorf("JiaZi[1].WuXing() = %q, want 海中金", jz.WuXing())
	}
}

func TestJiaZiGetNaYinWuXing(t *testing.T) {
	// GetNaYinWuXing extracts the base element from NaYin name
	tests := []struct {
		index int
		want  WuXing
	}{
		{1, Metal},  // 海中金 → Metal
		{3, Fire},   // 炉中火 → Fire
		{5, Wood},   // 大林木 → Wood
		{9, Metal},  // 剑锋金 → Metal
		{13, Water}, // 涧下水 → Water
		{15, Earth}, // 城头土 → Earth
		{60, Water}, // 大海水 → Water
	}
	for _, tt := range tests {
		jz, _ := GetJiaZi(tt.index)
		got := jz.GetNaYinWuXing()
		if got != tt.want {
			t.Errorf("JiaZi[%d].GetNaYinWuXing() = %v, want %v (NaYin=%q)", tt.index, got, tt.want, jz.NaYin)
		}
	}
}

func TestJiaZiAllNaYinWuXing(t *testing.T) {
	// All 60 entries should return a valid WuXing
	for i := 1; i <= 60; i++ {
		jz, err := GetJiaZi(i)
		if err != nil {
			t.Fatal(err)
		}
		wx := jz.GetNaYinWuXing()
		if wx != Metal && wx != Wood && wx != Water && wx != Fire && wx != Earth {
			t.Errorf("JiaZi[%d].GetNaYinWuXing() = %v, invalid (NaYin=%q)", i, wx, jz.NaYin)
		}
	}
}
