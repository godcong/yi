package yi

import (
	"testing"
)

func TestGetGuaXiang1(t *testing.T) {
	gx := getGuaXiangs()
	if len(gx) != 64 {
		t.Fatal("not enough", len(gx))
	}
	for i := range gx {
		t.Logf("idx(%+v):%+v", i, gx[i])
	}
}

func TestGetGuaXiang(t *testing.T) {
	type args struct {
		guaIdx string
	}
	tests := []struct {
		name string
		args args
		want *GuaXiang
	}{
		// TODO: Add more cases.
		{
			name: "",
			args: args{
				guaIdx: "兑艮",
			},
			want: &GuaXiang{
				GuaXu:    31,
				ShangShu: 1,
				XiaShu:   6,
			},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := GetGuaXiang(tt.args.guaIdx)
				t.Logf("卦象：%+v", got)
				if got.ShangShu != tt.want.ShangShu {
					t.Errorf("GetGuaXiang() = %v, want %v", got, tt.want)
				}
				if got.XiaShu != tt.want.XiaShu {
					t.Errorf("GetGuaXiang() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
