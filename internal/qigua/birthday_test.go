package qigua

import (
	"testing"

	"github.com/godcong/yi/core"
)

func TestDivineByBirthday_SolarInput(t *testing.T) {
	params := core.BirthdayParams{
		Year:    1990,
		Month:   6,
		Day:     15,
		Hour:    12,
		IsLunar: false,
	}
	zy := DivineByBirthday(params)
	if zy == nil {
		t.Fatal("expected non-nil ZhouYi")
	}
	benGua := zy.GetGua(core.Ben)
	if benGua == nil {
		t.Fatal("expected non-nil ben gua")
	}
}

func TestDivineByBirthday_LunarInput(t *testing.T) {
	params := core.BirthdayParams{
		Year:    1990,
		Month:   5,
		Day:     23,
		Hour:    12,
		IsLunar: true,
	}
	zy := DivineByBirthday(params)
	if zy == nil {
		t.Fatal("expected non-nil ZhouYi")
	}
}

func TestDivineByBirthday_SameResult_LunarAndSolar(t *testing.T) {
	solarParams := core.BirthdayParams{
		Year:    1990,
		Month:   6,
		Day:     15,
		Hour:    12,
		IsLunar: false,
	}

	lunarParams := core.BirthdayParams{
		Year:    1990,
		Month:   5,
		Day:     23,
		Hour:    12,
		IsLunar: true,
	}

	zySolar := DivineByBirthday(solarParams)
	zyLunar := DivineByBirthday(lunarParams)

	if zySolar == nil || zyLunar == nil {
		t.Fatal("expected non-nil results")
	}

	benSolar := zySolar.GetGua(core.Ben)
	benLunar := zyLunar.GetGua(core.Ben)
	if benSolar == nil || benLunar == nil {
		t.Fatal("expected non-nil ben gua")
	}

	if benSolar.Ming != benLunar.Ming {
		t.Errorf("same birthday should produce same gua: solar=%s, lunar=%s", benSolar.Ming, benLunar.Ming)
	}
}

func TestDivineByBirthday_DefaultHour(t *testing.T) {
	params := core.BirthdayParams{
		Year:    1990,
		Month:   6,
		Day:     15,
		Hour:    0,
		IsLunar: false,
	}
	zy := DivineByBirthday(params)
	if zy == nil {
		t.Fatal("expected non-nil ZhouYi for zero hour (should default to wu shi)")
	}
}

func TestDivineByBirthday_LeapMonth(t *testing.T) {
	params := core.BirthdayParams{
		Year:        2023,
		Month:       2,
		Day:         15,
		Hour:        12,
		IsLunar:     true,
		IsLeapMonth: true,
	}
	zy := DivineByBirthday(params)
	if zy == nil {
		t.Fatal("expected non-nil ZhouYi for leap month")
	}
}

func TestShichen(t *testing.T) {
	tests := []struct {
		hour int
		want int
	}{
		{0, 1},
		{1, 2},
		{3, 3},
		{5, 4},
		{7, 5},
		{9, 6},
		{11, 7},
		{12, 7},
		{13, 8},
		{15, 9},
		{17, 10},
		{19, 11},
		{21, 12},
		{23, 1},
	}
	for _, tt := range tests {
		got := shichen(tt.hour)
		if got != tt.want {
			t.Errorf("shichen(%d) = %d, want %d", tt.hour, got, tt.want)
		}
	}
}
