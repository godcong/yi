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
	gua := NumberQiGua(17, 19, StringToTime("2018-09-27 18:00"))
	log.Println(gua.Get(BenGua), gua.Get(BianGua))
}
