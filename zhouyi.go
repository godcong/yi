package yi

import (
	"log"
	"strings"
	"time"
)

type GuaMing int

const (
	//本卦
	BenGua GuaMing = iota
	//变卦
	BianGua
	//互卦
	HuGua
	//错卦
	CuoGua
	//综卦
	ZongGua
	GuaMax
)

const (
	//卦象:乾
	QianGua = 0x00
	//卦象:兑
	DuiGua = 0x01
	//卦象:离
	LiGua = 0x02
	//卦象:震
	ZhenGua = 0x03
	//卦象:巽
	XunGua = 0x04
	//卦象:坎
	KanGua = 0x05
	//卦象:艮
	GenGua = 0x06
	//卦象:坤
	KunGua = 0x07
)

var fu = []string{"☰", "☱", "☲", "☳", "☴", "☵", "☶", "☷"}

var gua = [...]string{
	KunGua:  "坤",
	GenGua:  "艮",
	KanGua:  "坎",
	XunGua:  "巽",
	ZhenGua: "震",
	LiGua:   "离",
	DuiGua:  "兑",
	QianGua: "乾",
}

// Yi 周易卦象
type Yi struct {
	gua     [GuaMax]*GuaXiang
	bianShu []int
}

func (y *Yi) Get(m GuaMing) *GuaXiang {

	if m < 0 && m >= GuaMax {
		return nil
	}

	return y.gua[m]
}

// TimeQiGua
func TimeQiGua(xia int, shang int, t time.Time) *Yi {
	ben := benGua(shang, xia)
	bs := timeToBian(t)
	bian := bianGua(ben, bs)
	hu := ben
	if ben.ShangShu == KunGua && ben.XiaShu == KunGua ||
		ben.ShangShu == QianGua && ben.XiaShu == QianGua {
		hu = bian
	}
	hu = huGua(hu)
	cuo := cuoGua(ben)
	zong := zongGua(ben)
	return &Yi{
		gua: [GuaMax]*GuaXiang{
			BenGua:  ben,
			BianGua: bian,
			HuGua:   hu,
			CuoGua:  cuo,
			ZongGua: zong,
		},
		bianShu: []int{bs},
	}
}

func getBianShu(bs ...int) int {
	bNum := 0
	if bs != nil {
		bNum = bs[0]
	}
	return bNum
}

// NumberQiGua
func NumberQiGua(xia int, shang int, bs ...int) *Yi {
	ben := benGua(shang, xia)
	bian := bianGua(ben, bs...)
	hu := ben
	if ben.ShangShu == KunGua && ben.XiaShu == KunGua ||
		ben.ShangShu == QianGua && ben.XiaShu == QianGua {
		hu = bian
	}
	hu = huGua(hu)
	cuo := cuoGua(ben)
	zong := zongGua(ben)
	return &Yi{
		gua: [GuaMax]*GuaXiang{
			BenGua:  ben,
			BianGua: bian,
			HuGua:   hu,
			CuoGua:  cuo,
			ZongGua: zong,
		},
		bianShu: bs,
	}
}

// StringToTime trans string to time
func StringToTime(s string) time.Time {
	t, err := time.ParseInLocation("2006-01-02 15:04", s, time.Local)
	if err != nil {
		return time.Time{}
	}
	log.Println(t.String(), t.Unix())
	return t
}

func timeToBian(t time.Time) int {
	t, _ = time.ParseInLocation("2006-01-02 15:04", t.Format("2006-01-02 15:04"), t.Location())
	bnum := t.Unix()
	if bnum != 0 {
		return int(bnum % 6)
	}
	return 0
}

// set 设定卦象
func (y *Yi) set(idx int, xiang *GuaXiang) {
	y.gua[idx] = xiang
}

// BianYao 变卦，爻
func (y *Yi) BianYao() int {
	return bianYao(y.bianShu...)
}

func bianYao(bs ...int) int {
	s := getBianShu(bs...)
	return getYao(s)
}

func getGua(i int) string {
	if i = i % 8; i != 0 {
		return gua[i-1]
	}
	return gua[7]
}

// 取爻
func getYao(i int) int {
	if i = i % 6; i != 0 {
		return i - 1
	}
	return 5
}

// 本卦
func benGua(x, m int) *GuaXiang {
	bg := strings.Join([]string{getGua(x), getGua(m)}, "")
	gx := GetGuaXiangV2()
	if v, b := gx[bg]; b {
		return v
	}
	return nil
}

// 变卦
func bianGua(ben *GuaXiang, b ...int) *GuaXiang {
	gxs := GetGuaXiangV2()
	bz := bianYao(b...)
	sg := gua[ben.ShangShu]
	xg := gua[ben.XiaShu]
	if bz > 2 {
		sg = gua[bian(ben.ShangShu, bz-3)]

	} else {
		xg = gua[bian(ben.XiaShu, bz)]
	}
	gua := strings.Join([]string{sg, xg}, "")

	return gxs[gua]
}

// 变
func bian(gua, bian int) int {
	idx := 1 << (2 - uint(bian))
	if gua&idx == 0 {
		gua = gua | idx
	} else {
		gua = gua ^ idx
	}
	return gua
}

// 互
func hu(shang, xia int) int {
	huXia := 0
	er := 1 << 1
	san := 1 << 0
	si := 1 << 2
	if xia&er > 0 {
		huXia |= 1 << 2
	}
	if xia&san > 0 {
		huXia |= 1 << 1
	}
	if shang&si > 0 {
		huXia |= 1 << 0
	}
	return huXia
}

// 交
func jiao(shang, xia int) int {
	jiaoShang := 0
	san := 1 << 0
	si := 1 << 2
	wu := 1 << 1
	if xia&san > 0 {
		//位移2
		jiaoShang |= 1 << 2
	}
	if shang&si > 0 {
		//位移1
		jiaoShang |= 1 << 1
	}
	if shang&wu > 0 {
		//位不动
		jiaoShang |= 1 << 0
	}
	return jiaoShang
}

// 错
func cuo(gua int) int {
	gua ^= 0x7
	return gua
}

// 错卦
func cuoGua(ben *GuaXiang) *GuaXiang {
	gx := GetGuaXiangV2()
	sg := gua[cuo(ben.ShangShu)]
	xg := gua[cuo(ben.XiaShu)]
	newGua := strings.Join([]string{sg, xg}, "")
	return gx[newGua]
}

// 综
func zong(shang, xia int) (int, int) {
	zShang := 0
	zXia := 0
	if (shang & (1 << 2)) > 0 {
		zXia |= 1 << 0
	}
	if (shang & (1 << 1)) > 0 {
		zXia |= 1 << 1
	}
	if (shang & (1 << 0)) > 0 {
		zXia |= 1 << 2
	}
	if (xia & (1 << 2)) > 0 {
		zShang |= 1 << 0
	}
	if (xia & (1 << 1)) > 0 {
		zShang |= 1 << 1
	}
	if (xia & (1 << 0)) > 0 {
		zShang |= 1 << 2
	}
	return zShang, zXia
}

// 综卦
func zongGua(ben *GuaXiang) *GuaXiang {
	sg, xg := zong(ben.ShangShu, ben.XiaShu)
	newGua := strings.Join([]string{gua[sg], gua[xg]}, "")
	return gxs[newGua]
}

// 互卦
func huGua(ben *GuaXiang) *GuaXiang {
	bg := strings.Join([]string{getGua(jiao(ben.ShangShu, ben.XiaShu)), getGua(hu(ben.ShangShu, ben.XiaShu))}, "")
	gx := GetGuaXiangV2()
	if v, b := gx[bg]; b {
		return v
	}
	return nil
}
