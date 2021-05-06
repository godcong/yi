package yi

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestQiGua(t *testing.T) {
	yi := QiGua(7, 7)
	get := yi.Get(0)
	t.Log(get)
}

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

func TestQiGua1(t *testing.T) {
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
				xia:   7,
				shang: 7,
			},
			want: &Yi{
				gua: [GuaMax]*GuaXiang{
					{
						GuaXu:       0,
						ShangGua:    "",
						ShangShu:    0,
						XiaGua:      "",
						XiaShu:      0,
						JiXiong:     "",
						GuaXiang:    "",
						GuaMing:     "",
						GuaYi:       "",
						FuHao:       "",
						GuaYaos:     [YaoMax]*GuaYao{},
						Yong:        "",
						YongJiXiong: "",
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
					for _, xiang := range got.gua {
						fmt.Printf("%+v", xiang)
					}

					t.Errorf("QiGua() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
