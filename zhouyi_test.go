package core

import (
	"log"
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

	gua1 := TimeQiGua(23, 20, StringToTime("2018-01-01 00:00"))
	log.Println(gua1.Get(BenGua), gua1.Get(BianGua))
	gua2 := TimeQiGua(23, 20, StringToTime("2019-01-01 00:00"))
	log.Println(gua2.Get(BenGua), gua2.Get(BianGua))
	gua3 := TimeQiGua(23, 20, StringToTime("2020-01-01 00:00"))
	log.Println(gua3.Get(BenGua), gua3.Get(BianGua))
	gua4 := TimeQiGua(23, 20, StringToTime("2021-01-01 00:00"))
	log.Println(gua4.Get(BenGua), gua4.Get(BianGua))
	gua5 := TimeQiGua(23, 20, StringToTime("2022-01-01 00:00"))
	log.Println(gua5.Get(BenGua), gua5.Get(BianGua))
	gua6 := TimeQiGua(23, 20, StringToTime("2023-01-01 00:00"))
	log.Println(gua6.Get(BenGua), gua6.Get(BianGua))
}
