package yi

import "testing"

func TestGetGuaByIndex(t *testing.T) {
	gua, err := GetGuaByIndex("乾乾")
	if err != nil {
		t.Fatalf("GetGuaByIndex failed: %v", err)
	}
	if gua.Xu != 1 {
		t.Errorf("expected Xu=1, got %d", gua.Xu)
	}
	t.Logf("Gua: %+v", gua)
}

func TestGetGuaByXu(t *testing.T) {
	for xu := 1; xu <= 64; xu++ {
		gua, err := GetGuaByXu(xu)
		if err != nil {
			t.Fatalf("GetGuaByXu(%d) failed: %v", xu, err)
		}
		if gua.Xu != xu {
			t.Errorf("expected Xu=%d, got %d", xu, gua.Xu)
		}
		t.Logf("Xu %d: %s", xu, gua.Ming)
	}
}

func TestDivine(t *testing.T) {
	yi := Divine(8, 8)
	if yi == nil {
		t.Fatal("Divine returned nil")
	}
	
	ben := yi.GetGua(Ben)
	if ben == nil {
		t.Fatal("GetGua(Ben) returned nil")
	}
	
	t.Logf("Ben Gua: %+v", ben)
}
