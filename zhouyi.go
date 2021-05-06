package yi

import (
	"log"
	"math/bits"
	"strconv"
	"strings"
	"time"
)

const TimeFormat = "2006-01-02 15:04"

const (
	// BenGua 本卦
	BenGua = iota
	// BianGua 变卦
	BianGua
	// HuGua 互卦
	HuGua
	// CuoGua 错卦
	CuoGua
	// ZongGua 综卦
	ZongGua
	// GuaMax 卦最大值
	GuaMax
)

type BaGua = int

const (
	// QianGua 卦象:乾(0)
	QianGua BaGua = 0b000
	// DuiGua 卦象:兑(1)
	DuiGua BaGua = 0b001
	// LiGua 卦象:离(2)
	LiGua BaGua = 0b010
	// ZhenGua 卦象:震(3)
	ZhenGua BaGua = 0b011
	// XunGua 卦象:巽(4)
	XunGua BaGua = 0b100
	// KanGua 卦象:坎(5)
	KanGua BaGua = 0b101
	// GenGua 卦象:艮(6)
	GenGua BaGua = 0b110
	// KunGua 卦象:坤(7)
	KunGua BaGua = 0b111
)

//符号
var fu map[int]string

///卦象
var gua map[int]string

//Yi 周易卦象
type Yi struct {
	gua     [GuaMax]*GuaXiang
	bianShu []int
}

//Get 取卦
func (y *Yi) Get(m int) *GuaXiang {
	if m < 0 && m >= GuaMax {
		return nil
	}
	return y.gua[m]
}

func init() {
	fu = make(map[int]string)
	gua = make(map[int]string)

	file8gua, err := DataFiles.Open("data/8gua.csv")
	if err != nil {
		panic(err)
	}

	records, err := readData(file8gua)

	if err != nil {
		log.Fatal(err)
	}

	for _, record := range records {
		guashu, err := strconv.ParseInt(record[1], 10, bits.UintSize)
		if err != nil {
			if len(record[1]) == 0 {
				continue
			}
			panic(err)
		}
		fu[int(guashu)] = record[2]
		gua[int(guashu)] = record[0]
	}

}

//QiGua 起卦
func QiGua(xia, shang int) *Yi {
	return NumberQiGua(shang, xia)
}

//TimeQiGua 按时间起卦
func TimeQiGua(xia int, shang int, t time.Time) *Yi {
	bs := timeToBian(t)
	return NumberQiGua(xia, shang, bs)
}

func getBianShu(bs ...int) int {
	bNum := 0
	if bs != nil {
		bNum = bs[0]
	}
	return bNum
}

//NumberQiGua 按数起卦
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

//StringToTime trans string to time
func StringToTime(s string) time.Time {
	t, err := time.ParseInLocation(TimeFormat, s, time.Local)
	if err != nil {
		return time.Time{}
	}
	log.Println(t.String(), t.Unix())
	return t
}

func timeToBian(t time.Time) int {
	t, _ = time.ParseInLocation(TimeFormat, t.Format(TimeFormat), t.Location())
	bnum := t.Unix()
	if bnum != 0 {
		return int(bnum % 6)
	}
	return 0
}

//BianYao 变卦，爻
func (y *Yi) BianYao() int {
	return bianYao(y.bianShu...)
}

func bianYao(bs ...int) int {
	s := getBianShu(bs...)
	return getYao(s)
}

//由数取卦
func getGua(i int) string {
	if i = i % 8; i != 0 {
		return gua[i-1]
	}
	return gua[7]
}

// GetGua3Num 三数取卦数(1~8)
func GetGua3Num(shang, zhong, xia int) int {
	shangYao := shang % 2
	zhongYao := zhong % 2
	xiaYao := xia % 2
	return shangYao<<2 + zhongYao<<1 + xiaYao + 1
}

//取爻
func getYao(i int) int {
	if i = i % 6; i != 0 {
		return i - 1
	}
	return 5
}

//本卦
func benGua(shang, xia int) *GuaXiang {
	bg := strings.Join([]string{getGua(shang), getGua(xia)}, "")
	gx := getGuaXiangs()
	if v, b := gx[bg]; b {
		return v
	}
	return nil
}

//变卦
func bianGua(ben *GuaXiang, b ...int) *GuaXiang {
	gx := getGuaXiangs()
	bz := bianYao(b...)
	sg := gua[ben.ShangShu]
	xg := gua[ben.XiaShu]
	if bz > 2 {
		sg = gua[bian(ben.ShangShu, bz-3)]

	} else {
		xg = gua[bian(ben.XiaShu, bz)]
	}
	gua := strings.Join([]string{sg, xg}, "")

	if gx[gua] == nil {
		panic(gua)
	}

	return gx[gua]
}

//变
func bian(gua, bian int) int {
	idx := 1 << (2 - uint(bian))
	if gua&idx == 0 {
		gua = gua | idx
	} else {
		gua = gua ^ idx
	}
	return gua
}

//互
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

//交
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

//错
func cuo(gua int) int {
	gua ^= 0x7
	return gua
}

//错卦
func cuoGua(ben *GuaXiang) *GuaXiang {
	gx := getGuaXiangs()
	sg := gua[cuo(ben.ShangShu)]
	xg := gua[cuo(ben.XiaShu)]
	newGua := strings.Join([]string{sg, xg}, "")
	return gx[newGua]
}

//综
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

//综卦
func zongGua(ben *GuaXiang) *GuaXiang {
	sg, xg := zong(ben.ShangShu, ben.XiaShu)
	newGua := strings.Join([]string{gua[sg], gua[xg]}, "")
	return gx[newGua]
}

//互卦
func huGua(ben *GuaXiang) *GuaXiang {
	bg := strings.Join([]string{getGua(jiao(ben.ShangShu, ben.XiaShu)), getGua(hu(ben.ShangShu, ben.XiaShu))}, "")
	gx := getGuaXiangs()
	if v, b := gx[bg]; b {
		return v
	}
	return nil
}
