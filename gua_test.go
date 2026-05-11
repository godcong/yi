package yi

import (
	"testing"
	"time"
)

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

func TestDivineByNumber(t *testing.T) {
	yi := DivineByNumber(1, 0, 1)
	if yi == nil {
		t.Fatal("DivineByNumber returned nil")
	}

	ben := yi.GetGua(Ben)
	if ben == nil {
		t.Fatal("GetGua(Ben) returned nil")
	}

	bian := yi.GetGua(Bian)
	if bian == nil {
		t.Fatal("GetGua(Bian) returned nil")
	}

	t.Logf("Ben: %s, Bian: %s", ben.Ming, bian.Ming)
}

func TestDivineByTime(t *testing.T) {
	yi := DivineByTime(1, 0, time.Now())
	if yi == nil {
		t.Fatal("DivineByTime returned nil")
	}
	t.Logf("DivineByTime success")
}

func TestGetGua(t *testing.T) {
	yi := Divine(3, 5)

	tests := []struct {
		guaType int
		name    string
	}{
		{Ben, "Ben"},
		{Bian, "Bian"},
		{Hu, "Hu"},
		{Cuo, "Cuo"},
		{Zong, "Zong"},
	}

	for _, tt := range tests {
		gua := yi.GetGua(tt.guaType)
		if gua == nil {
			t.Errorf("GetGua(%s) returned nil", tt.name)
		}
	}
}

func TestGetGuaByIndex_NotFound(t *testing.T) {
	_, err := GetGuaByIndex("不存在")
	if err != ErrGuaNotFound {
		t.Errorf("expected ErrGuaNotFound, got %v", err)
	}
}

func TestGetGuaByIndex_Empty(t *testing.T) {
	_, err := GetGuaByIndex("")
	if err != ErrGuaNotFound {
		t.Errorf("expected ErrGuaNotFound for empty index, got %v", err)
	}
}

func TestGetGuaByIndex_BoundaryValues(t *testing.T) {
	tests := []struct {
		index string
	}{
		{"0"},
		{"-1"},
		{"100"},
		{"999"},
	}

	for _, tt := range tests {
		_, err := GetGuaByIndex(tt.index)
		if err != ErrGuaNotFound {
			t.Errorf("GetGuaByIndex(%q) = %v, want ErrGuaNotFound", tt.index, err)
		}
	}
}

func TestGetGuaByXu_InvalidIndex(t *testing.T) {
	_, err := GetGuaByXu(0)
	if err != ErrInvalidGuaIndex {
		t.Errorf("expected ErrInvalidGuaIndex, got %v", err)
	}

	_, err = GetGuaByXu(-1)
	if err != ErrInvalidGuaIndex {
		t.Errorf("expected ErrInvalidGuaIndex for xu=-1, got %v", err)
	}

	_, err = GetGuaByXu(65)
	if err != ErrInvalidGuaIndex {
		t.Errorf("expected ErrInvalidGuaIndex for xu=65, got %v", err)
	}

	_, err = GetGuaByXu(100)
	if err != ErrInvalidGuaIndex {
		t.Errorf("expected ErrInvalidGuaIndex for xu=100, got %v", err)
	}
}

func TestGetBaguaName(t *testing.T) {
	tests := []struct {
		bagua Bagua
		name  string
	}{
		{Qian, "乾"},
		{Dui, "兑"},
		{Li, "离"},
		{Zhen, "震"},
		{Xun, "巽"},
		{Kan, "坎"},
		{Gen, "艮"},
		{Kun, "坤"},
	}

	for _, tt := range tests {
		if name := GetBaguaName(tt.bagua); name != tt.name {
			t.Errorf("GetBaguaName(%d) = %s, want %s", tt.bagua, name, tt.name)
		}
	}

	if name := GetBaguaName(8); name != "" {
		t.Errorf("GetBaguaName(8) = %s, want empty", name)
	}
}

func TestGetBaguaSymbol(t *testing.T) {
	tests := []struct {
		bagua  Bagua
		symbol string
	}{
		{Qian, "☰"},
		{Dui, "☱"},
		{Li, "☲"},
		{Zhen, "☳"},
		{Xun, "☴"},
		{Kan, "☵"},
		{Gen, "☶"},
		{Kun, "☷"},
	}

	for _, tt := range tests {
		if symbol := GetBaguaSymbol(tt.bagua); symbol != tt.symbol {
			t.Errorf("GetBaguaSymbol(%d) = %s, want %s", tt.bagua, symbol, tt.symbol)
		}
	}
}

func TestBaguaConstants(t *testing.T) {
	if Qian != 0 {
		t.Errorf("Qian = %d, want 0", Qian)
	}
	if Dui != 1 {
		t.Errorf("Dui = %d, want 1", Dui)
	}
	if Li != 2 {
		t.Errorf("Li = %d, want 2", Li)
	}
	if Zhen != 3 {
		t.Errorf("Zhen = %d, want 3", Zhen)
	}
	if Xun != 4 {
		t.Errorf("Xun = %d, want 4", Xun)
	}
	if Kan != 5 {
		t.Errorf("Kan = %d, want 5", Kan)
	}
	if Gen != 6 {
		t.Errorf("Gen = %d, want 6", Gen)
	}
	if Kun != 7 {
		t.Errorf("Kun = %d, want 7", Kun)
	}
}

func TestYaoPosition(t *testing.T) {
	if Chu != 0 {
		t.Errorf("Chu = %d, want 0", Chu)
	}
	if Er != 1 {
		t.Errorf("Er = %d, want 1", Er)
	}
	if San != 2 {
		t.Errorf("San = %d, want 2", San)
	}
	if Si != 3 {
		t.Errorf("Si = %d, want 3", Si)
	}
	if Wu != 4 {
		t.Errorf("Wu = %d, want 4", Wu)
	}
	if Shang != 5 {
		t.Errorf("Shang = %d, want 5", Shang)
	}
}

func TestGuaType(t *testing.T) {
	if Ben != 0 {
		t.Errorf("Ben = %d, want 0", Ben)
	}
	if Bian != 1 {
		t.Errorf("Bian = %d, want 1", Bian)
	}
	if Hu != 2 {
		t.Errorf("Hu = %d, want 2", Hu)
	}
	if Cuo != 3 {
		t.Errorf("Cuo = %d, want 3", Cuo)
	}
	if Zong != 4 {
		t.Errorf("Zong = %d, want 4", Zong)
	}
}

func TestGetYao(t *testing.T) {
	gua, _ := GetGuaByIndex("乾乾")
	if gua == nil {
		t.Fatal("GetGuaByIndex returned nil")
	}

	yao := gua.GetYao(Chu)
	if yao == nil {
		t.Error("GetYao(Chu) returned nil")
	}

	yao = gua.GetYao(6)
	if yao != nil {
		t.Error("GetYao(6) should return nil")
	}
}

func TestGetBianYao(t *testing.T) {
	yi := DivineByNumber(1, 0, 2)
	if yi == nil {
		t.Fatal("DivineByNumber returned nil")
	}

	bianYao := yi.GetBianYao()
	if bianYao != 1 {
		t.Errorf("GetBianYao() = %d, want 1", bianYao)
	}
}

func TestIsJi(t *testing.T) {
	yi := DivineByNumber(1, 0, 1)
	if yi == nil {
		t.Fatal("DivineByNumber returned nil")
	}

	isJi := yi.IsJi(Male)
	t.Logf("IsJi(Male) = %v", isJi)

	isJi = yi.IsJi(Female)
	t.Logf("IsJi(Female) = %v", isJi)
}

func TestContains(t *testing.T) {
	tests := []struct {
		s      string
		substr string
		want   bool
	}{
		{"吉凶", "凶", true},
		{"吉", "凶", false},
		{"大吉", "吉", true},
		{"", "", true},
		{"凶", "", true},
		{"大吉大利", "凶", false},
	}

	for _, tt := range tests {
		got := contains(tt.s, tt.substr)
		if got != tt.want {
			t.Errorf("contains(%q, %q) = %v, want %v", tt.s, tt.substr, got, tt.want)
		}
	}
}

func TestGetDayan(t *testing.T) {
	for n := 1; n <= 81; n++ {
		dayan, err := GetDayan(n)
		if err != nil {
			t.Fatalf("GetDayan(%d) failed: %v", n, err)
		}
		if dayan.Number != n {
			t.Errorf("GetDayan(%d).Number = %d", n, dayan.Number)
		}
	}
}

func TestGetDayan_Invalid(t *testing.T) {
	_, err := GetDayan(0)
	if err != ErrInvalidDayanNumber {
		t.Errorf("expected ErrInvalidDayanNumber, got %v", err)
	}

	_, err = GetDayan(-1)
	if err != ErrInvalidDayanNumber {
		t.Errorf("expected ErrInvalidDayanNumber for n=-1, got %v", err)
	}

	_, err = GetDayan(82)
	if err != ErrInvalidDayanNumber {
		t.Errorf("expected ErrInvalidDayanNumber for n=82, got %v", err)
	}

	_, err = GetDayan(100)
	if err != ErrInvalidDayanNumber {
		t.Errorf("expected ErrInvalidDayanNumber for n=100, got %v", err)
	}
}

func TestGetDayan_BoundaryValues(t *testing.T) {
	_, err := GetDayan(0)
	if err != ErrInvalidDayanNumber {
		t.Errorf("GetDayan(0) should return ErrInvalidDayanNumber")
	}

	_, err = GetDayan(-1)
	if err != ErrInvalidDayanNumber {
		t.Errorf("GetDayan(-1) should return ErrInvalidDayanNumber")
	}

	_, err = GetDayan(82)
	if err != ErrInvalidDayanNumber {
		t.Errorf("GetDayan(82) should return ErrInvalidDayanNumber")
	}

	_, err = GetDayan(999)
	if err != ErrInvalidDayanNumber {
		t.Errorf("GetDayan(999) should return ErrInvalidDayanNumber")
	}
}

func TestMustGetDayan(t *testing.T) {
	dayan := MustGetDayan(1)
	if dayan.Number != 1 {
		t.Errorf("MustGetDayan(1).Number = %d, want 1", dayan.Number)
	}

	dayan = MustGetDayan(0)
	if dayan.Number != 0 {
		t.Errorf("MustGetDayan(0).Number = %d, want 0", dayan.Number)
	}
}

func TestDayan_IsJi(t *testing.T) {
	tests := []struct {
		number int
		jiXiong string
		want    bool
	}{
		{1, "吉", true},
		{2, "凶", false},
		{8, "半吉", true},
	}

	for _, tt := range tests {
		d := Dayan{Number: tt.number, JiXiong: tt.jiXiong}
		if got := d.IsJi(); got != tt.want {
			t.Errorf("Dayan{Number: %d, JiXiong: %q}.IsJi() = %v, want %v", tt.number, tt.jiXiong, got, tt.want)
		}
	}
}

func TestDayan_IsXiong(t *testing.T) {
	d := Dayan{JiXiong: "凶"}
	if !d.IsXiong() {
		t.Error("IsXiong() should return true for 凶")
	}

	d = Dayan{JiXiong: "吉"}
	if d.IsXiong() {
		t.Error("IsXiong() should return false for 吉")
	}
}

func TestYinYang(t *testing.T) {
	if Yang.String() != "阳" {
		t.Errorf("Yang.String() = %s, want 阳", Yang.String())
	}
	if Yin.String() != "阴" {
		t.Errorf("Yin.String() = %s, want 阴", Yin.String())
	}
	if !Yang.IsYang() {
		t.Error("Yang.IsYang() should return true")
	}
	if !Yin.IsYin() {
		t.Error("Yin.IsYin() should return true")
	}
}

func TestSex(t *testing.T) {
	if Male.String() != "男" {
		t.Errorf("Male.String() = %s, want 男", Male.String())
	}
	if Female.String() != "女" {
		t.Errorf("Female.String() = %s, want 女", Female.String())
	}
	if !Male.IsMale() {
		t.Error("Male.IsMale() should return true")
	}
	if !Female.IsFemale() {
		t.Error("Female.IsFemale() should return true")
	}
}

func TestWuXing(t *testing.T) {
	if Wood.String() != "木" {
		t.Errorf("Wood.String() = %s, want 木", Wood.String())
	}
	if Fire.String() != "火" {
		t.Errorf("Fire.String() = %s, want 火", Fire.String())
	}
	if Earth.String() != "土" {
		t.Errorf("Earth.String() = %s, want 土", Earth.String())
	}
	if Metal.String() != "金" {
		t.Errorf("Metal.String() = %s, want 金", Metal.String())
	}
	if Water.String() != "水" {
		t.Errorf("Water.String() = %s, want 水", Water.String())
	}
}

func TestWuXing_Sheng(t *testing.T) {
	if Wood.Sheng() != Fire {
		t.Errorf("Wood.Sheng() = %s, want 火", Wood.Sheng().String())
	}
	if Fire.Sheng() != Earth {
		t.Errorf("Fire.Sheng() = %s, want 土", Fire.Sheng().String())
	}
	if Earth.Sheng() != Metal {
		t.Errorf("Earth.Sheng() = %s, want 金", Earth.Sheng().String())
	}
	if Metal.Sheng() != Water {
		t.Errorf("Metal.Sheng() = %s, want 水", Metal.Sheng().String())
	}
	if Water.Sheng() != Wood {
		t.Errorf("Water.Sheng() = %s, want 木", Water.Sheng().String())
	}
}

func TestWuXing_Ke(t *testing.T) {
	if Wood.Ke() != Earth {
		t.Errorf("Wood.Ke() = %s, want 土", Wood.Ke().String())
	}
	if Earth.Ke() != Water {
		t.Errorf("Earth.Ke() = %s, want 水", Earth.Ke().String())
	}
	if Water.Ke() != Fire {
		t.Errorf("Water.Ke() = %s, want 火", Water.Ke().String())
	}
	if Fire.Ke() != Metal {
		t.Errorf("Fire.Ke() = %s, want 金", Fire.Ke().String())
	}
	if Metal.Ke() != Wood {
		t.Errorf("Metal.Ke() = %s, want 木", Metal.Ke().String())
	}
}

func TestGetWuXingByNumber(t *testing.T) {
	tests := []struct {
		n    int
		want string
	}{
		{1, "阳木"},
		{2, "阴木"},
		{3, "阳火"},
		{4, "阴火"},
		{5, "阳土"},
		{6, "阴土"},
		{7, "阳金"},
		{8, "阴金"},
		{9, "阳水"},
		{10, "阴水"},
	}

	for _, tt := range tests {
		got := GetWuXingByNumber(tt.n)
		if got != tt.want {
			t.Errorf("GetWuXingByNumber(%d) = %s, want %s", tt.n, got, tt.want)
		}
	}
}

func TestGetYinYangByNumber(t *testing.T) {
	if GetYinYangByNumber(1) != Yang {
		t.Error("GetYinYangByNumber(1) should be Yang")
	}
	if GetYinYangByNumber(2) != Yin {
		t.Error("GetYinYangByNumber(2) should be Yin")
	}
}
