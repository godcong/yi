package yi

import (
	"reflect"
	"testing"
	"time"
)

func TestJiaoHu(t *testing.T) {
	for i, j := 0, 0; i <= 7; j++ {
		t.Log(fu[i], fu[j], fu[hu(i, j)])
		t.Log(fu[i], fu[j], fu[jiao(i, j)])
		t.Log(fu[i], fu[cuo(i)])
		if j == 7 {
			i++
			j = 0
		}
	}
}

func TestCuo(t *testing.T) {
	for i := 0; i <= 7; i++ {
		t.Log(fu[i], fu[cuo(i)])

	}
}

func TestZong(t *testing.T) {
	for i, j := 0, 0; i <= 7; j++ {
		shang, xia := zong(i, j)
		t.Log(fu[i], fu[j], fu[shang], fu[xia])
		if j == 7 {
			i++
			j = 0
		}
	}

}

func TestTimeToBian(t *testing.T) {
	t.Log(timeToBian(time.Now()))
}

func TestNumberQiGua(t *testing.T) {
	//t.Log(bian(7, 0))
	//gua1 := NumberQiGua(23, 20, int(StringToTime("2018-01-01 00:00").Unix()/1000000%6))
	//log.Println(gua1.Get(BenGua), gua1.Get(BianGua))
	//gua2 := NumberQiGua(23, 20, int(StringToTime("2019-01-01 00:00").Unix()/1000000%6))
	//log.Println(gua2.Get(BenGua), gua2.Get(BianGua))
	//gua3 := NumberQiGua(23, 20, int(StringToTime("2020-01-01 00:00").Unix()/1000000%6))
	//log.Println(gua3.Get(BenGua), gua3.Get(BianGua))
	//gua4 := NumberQiGua(23, 20, int(StringToTime("2021-01-01 00:00").Unix()/1000000%6))
	//log.Println(gua4.Get(BenGua), gua4.Get(BianGua))
	//gua5 := NumberQiGua(23, 20, int(StringToTime("2022-01-01 00:00").Unix()/1000000%6))
	//log.Println(gua5.Get(BenGua), gua5.Get(BianGua))
	//gua6 := NumberQiGua(23, 20, int(StringToTime("2023-01-01 00:00").Unix()/1000000%6))
	//log.Println(gua6.Get(BenGua), gua6.Get(BianGua))
	//gua7 := NumberQiGua(23, 20, int(StringToTime("2024-01-01 00:00").Unix()/1000000%6))
	//log.Println(gua7.Get(BenGua), gua7.Get(BianGua))
}

func TestQiGua(t *testing.T) {
	type args struct {
		xia   int
		shang int
	}
	tests := []struct {
		name string
		args args
		want *Yi
	}{
		// TODO: Add more cases.
		{
			name: "",
			args: args{
				xia:   8,
				shang: 8,
			},
			want: &Yi{
				gua: [GuaMax]*GuaXiang{
					BenGua: {
						GuaXu:    2,
						ShangShu: 7,
						XiaShu:   7,
					},
				},
				bianShu: nil,
			},
		},
		{
			name: "",
			args: args{
				xia:   7,
				shang: 7,
			},
			want: &Yi{
				gua: [GuaMax]*GuaXiang{
					BenGua: {
						GuaXu:    52,
						ShangShu: 6,
						XiaShu:   6,
					},
				},
				bianShu: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := QiGua(tt.args.xia, tt.args.shang); !reflect.DeepEqual(got, tt.want) {
					t.Logf("本卦：%+v", got.gua[BenGua])
					if got.gua[BenGua].ShangShu != tt.want.gua[BenGua].ShangShu {
						t.Errorf("Shang QiGua() = %v, want %v", got.gua[BenGua], tt.want.gua[BenGua])
					}

					if got.gua[BenGua].XiaShu != tt.want.gua[BenGua].XiaShu {
						t.Errorf("Xia QiGua() = %v, want %v", got.gua[BenGua], tt.want.gua[BenGua])
					}

				}
			},
		)
	}
}
