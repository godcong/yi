package yi

import (
	"math/bits"
	"strconv"
)

type Yao int

const (
	ChuYao Yao = iota
	ErYao
	SanYao
	SiYao
	WuYao
	ShangYao
	YaoMax
)

//GuaXiang 卦象
type GuaXiang struct {
	GuaXu       int             //卦序
	ShangGua    string          //上卦
	ShangShu    int             //上卦数
	XiaGua      string          //下卦
	XiaShu      int             //下卦数
	JiXiong     string          //吉凶（？）
	GuaXiang    string          //卦象
	GuaMing     string          //卦名
	GuaYi       string          //卦意（邵雍）
	FuHao       string          //符号
	GuaYaos     [YaoMax]*GuaYao //初，二，三，四，五，上
	Yong        string          //用九,用六
	YongJiXiong string          //用九,用六吉凶
}

var gx map[string]*GuaXiang

func init() {
	gx = make(map[string]*GuaXiang)

	file64gua, err := DataFiles.Open("data/64gua.csv")
	if err != nil {
		panic(err)
	}

	records, err := readData(file64gua)

	if err != nil {
		panic(err)
	}

	for _, record := range records {
		guaxu, err := strconv.ParseInt(record[0], 10, bits.UintSize)
		if err != nil {
			panic(err)
		}
		shangshu, err := strconv.ParseInt(record[3], 10, bits.UintSize)
		if err != nil {
			panic(err)
		}
		xiashu, err := strconv.ParseInt(record[5], 10, bits.UintSize)
		if err != nil {
			panic(err)
		}

		guaxiang := GuaXiang{
			GuaXu:       int(guaxu),
			ShangGua:    record[2],
			ShangShu:    int(shangshu),
			XiaGua:      record[4],
			XiaShu:      int(xiashu),
			JiXiong:     record[6],
			GuaXiang:    record[7],
			GuaMing:     record[8],
			GuaYi:       record[9],
			FuHao:       record[10],
			Yong:        record[29],
			YongJiXiong: record[30],
		}

		guyao1 := GuaYao{
			Yao:     record[11], //初爻
			JiXiong: record[12], //初爻吉凶
			NvMing:  record[13], //初爻女命
		}
		guaxiang.GuaYaos[0] = &guyao1

		guyao2 := GuaYao{
			Yao:     record[14], //二爻
			JiXiong: record[15], //二爻吉凶
			NvMing:  record[16], //二爻女命
		}
		guaxiang.GuaYaos[1] = &guyao2

		guyao3 := GuaYao{
			Yao:     record[17], //三爻
			JiXiong: record[18], //三爻吉凶
			NvMing:  record[19], //三爻女命
		}
		guaxiang.GuaYaos[2] = &guyao3

		guyao4 := GuaYao{
			Yao:     record[20], //四爻
			JiXiong: record[21], //四爻吉凶
			NvMing:  record[22], //四爻女命
		}
		guaxiang.GuaYaos[3] = &guyao4

		guyao5 := GuaYao{
			Yao:     record[23], //五爻
			JiXiong: record[24], //五爻吉凶
			NvMing:  record[25], //五爻女命
		}
		guaxiang.GuaYaos[4] = &guyao5

		guyao6 := GuaYao{
			Yao:     record[26], //上爻
			JiXiong: record[27], //上爻吉凶
			NvMing:  record[28], //上爻女命
		}
		guaxiang.GuaYaos[5] = &guyao6

		gxIndex := record[1]
		if len(gxIndex) < 1 {
			panic("index is wrong")
		}

		gx[gxIndex] = &guaxiang
	}
}

func getGuaXiangs() map[string]*GuaXiang {
	return gx
}

func GetGuaXiang(guaIdx string) *GuaXiang {
	return gx[guaIdx]
}
