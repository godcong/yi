package yi

import "testing"

func TestShiYingInfo(t *testing.T) {
	gua, err := GetGuaByIndex("乾乾")
	if err != nil {
		t.Fatalf("GetGuaByIndex failed: %v", err)
	}

	sy := gua.GetShiYing()
	if sy == nil {
		t.Fatal("GetShiYing returned nil")
	}

	t.Logf("乾为天: 世=%d, 应=%d, 宫位=%s",
		sy.ShiPos, sy.YingPos, sy.Position)
}

func TestShiYingInfo_BenGong(t *testing.T) {
	gua, _ := GetGuaByIndex("乾乾")
	sy := gua.GetShiYing()

	if sy.ShiPos != Shang {
		t.Errorf("BenGong ShiPos should be Shang(5), got %d", sy.ShiPos)
	}
	if sy.YingPos != San {
		t.Errorf("BenGong YingPos should be San(2), got %d", sy.YingPos)
	}
	if sy.Position != BenGong {
		t.Errorf("BenGong Position should be BenGong, got %s", sy.Position)
	}
}

func TestShiYingInfo_YiShi(t *testing.T) {
	gua, _ := GetGuaByIndex("乾巽")
	sy := gua.GetShiYing()

	if sy.Position != YiShi {
		t.Errorf("乾巽 should be YiShi, got %s", sy.Position)
	}
	if sy.ShiPos != Chu {
		t.Errorf("YiShi ShiPos should be Chu(0), got %d", sy.ShiPos)
	}
	if sy.YingPos != Si {
		t.Errorf("YiShi YingPos should be Si(3), got %d", sy.YingPos)
	}
}

func TestShiYingInfo_ErShi(t *testing.T) {
	gua, _ := GetGuaByIndex("乾艮")
	sy := gua.GetShiYing()

	if sy.Position != ErShi {
		t.Errorf("乾艮 should be ErShi, got %s", sy.Position)
	}
	if sy.ShiPos != Er {
		t.Errorf("ErShi ShiPos should be Er(1), got %d", sy.ShiPos)
	}
	if sy.YingPos != Wu {
		t.Errorf("ErShi YingPos should be Wu(4), got %d", sy.YingPos)
	}
}

func TestShiYingInfo_YouHun(t *testing.T) {
	gua, _ := GetGuaByIndex("离坤")
	sy := gua.GetShiYing()

	if sy.Position != YouHun {
		t.Errorf("离坤 should be YouHun, got %s", sy.Position)
	}
	if sy.ShiPos != Si {
		t.Errorf("YouHun ShiPos should be Si(3), got %d", sy.ShiPos)
	}
	if sy.YingPos != Chu {
		t.Errorf("YouHun YingPos should be Chu(0), got %d", sy.YingPos)
	}
}

func TestShiYingInfo_GuiHun(t *testing.T) {
	gua, _ := GetGuaByIndex("离乾")
	sy := gua.GetShiYing()

	if sy.Position != GuiHun {
		t.Errorf("离乾 should be GuiHun, got %s", sy.Position)
	}
	if sy.ShiPos != San {
		t.Errorf("GuiHun ShiPos should be San(2), got %d", sy.ShiPos)
	}
	if sy.YingPos != Shang {
		t.Errorf("GuiHun YingPos should be Shang(5), got %d", sy.YingPos)
	}
}

func TestGuaPosition_String(t *testing.T) {
	tests := []struct {
		pos   GuaPosition
		names []string
	}{
		{BenGong, []string{"本宫", "BenGong"}},
		{YiShi, []string{"一世", "YiShi"}},
		{ErShi, []string{"二世", "ErShi"}},
		{SanShi, []string{"三世", "SanShi"}},
		{SiShi, []string{"四世", "SiShi"}},
		{WuShi, []string{"五世", "WuShi"}},
		{YouHun, []string{"游魂", "YouHun"}},
		{GuiHun, []string{"归魂", "GuiHun"}},
	}

	for _, tt := range tests {
		if s := tt.pos.String(); s == "" {
			t.Errorf("GuaPosition(%d).String() returned empty", tt.pos)
		} else {
			t.Logf("%s: %s", tt.names[1], s)
		}
	}
}

func TestGuaPosition_Invalid(t *testing.T) {
	pos := GuaPosition(100)
	if pos.String() != "" {
		t.Error("Invalid GuaPosition should return empty string")
	}
}

func TestGetGuaGong(t *testing.T) {
	gua, _ := GetGuaByIndex("乾乾")
	gong := gua.GetGuaGong()
	if gong != Qian {
		t.Errorf("乾为天 should be in Qian palace, got %d", gong)
	}

	gua2, _ := GetGuaByIndex("乾巽")
	gong2 := gua2.GetGuaGong()
	if gong2 != Qian {
		t.Errorf("乾巽 should be in Qian palace, got %d", gong2)
	}
}

func TestGetGuaGong_AllHexagrams(t *testing.T) {
	for xu := 1; xu <= 64; xu++ {
		gua, err := GetGuaByXu(xu)
		if err != nil {
			t.Fatalf("GetGuaByXu(%d) failed: %v", xu, err)
		}
		gong := gua.GetGuaGong()
		if gong < 0 || gong > 7 {
			t.Errorf("xu=%d: invalid GuaGong %d", xu, gong)
		}
	}
}

func TestGetGuaPosition_AllHexagrams(t *testing.T) {
	for xu := 1; xu <= 64; xu++ {
		gua, err := GetGuaByXu(xu)
		if err != nil {
			t.Fatalf("GetGuaByXu(%d) failed: %v", xu, err)
		}
		pos := gua.GetGuaPosition()
		if pos < 0 || pos >= GuaPositionCount {
			t.Errorf("xu=%d (%s): invalid GuaPosition %d", xu, gua.Ming, pos)
		}
		g := gua.GetGuaGong()
		t.Logf("xu=%d: %s -> %s宫(%s)", xu, gua.Ming, GetBaguaName(g), pos)
	}
}

func TestShiYingTable(t *testing.T) {
	expected := []struct {
		pos     GuaPosition
		shiPos  YaoPosition
		yingPos YaoPosition
	}{
		{BenGong, Shang, San},
		{YiShi, Chu, Si},
		{ErShi, Er, Wu},
		{SanShi, San, Shang},
		{SiShi, Si, Chu},
		{WuShi, Wu, Er},
		{YouHun, Si, Chu},
		{GuiHun, San, Shang},
	}

	for _, e := range expected {
		if shiYingTable[e.pos].ShiPos != e.shiPos {
			t.Errorf("Position %d: ShiPos should be %d, got %d",
				e.pos, e.shiPos, shiYingTable[e.pos].ShiPos)
		}
		if shiYingTable[e.pos].YingPos != e.yingPos {
			t.Errorf("Position %d: YingPos should be %d, got %d",
				e.pos, e.yingPos, shiYingTable[e.pos].YingPos)
		}
	}
}
