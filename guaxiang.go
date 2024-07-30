package yi

import (
	"fmt"
	"math/bits"
	"strconv"
)

// Yao 爻: 初，二，三，四，五，上
type Yao uint

const (
	ChuYao Yao = iota
	ErYao
	SanYao
	SiYao
	WuYao
	ShangYao
	YaoMax
)

type GuaXiangIndex uint

func (idx GuaXiangIndex) IsValid() bool {
	return idx >= 0 && idx.index() < SixtyFourMax
}

func (idx GuaXiangIndex) index() uint {
	return uint(idx)
}

func (idx GuaXiangIndex) String() string {
	return ""
}

// GuaXiang 卦象
type GuaXiang struct {
	GuaXu       uint           // 卦序
	ShangGua    string         // 上卦
	ShangShu    uint           // 上卦数
	XiaGua      string         // 下卦
	XiaShu      uint           // 下卦数
	JiXiong     string         // 吉凶（？）
	GuaXiang    string         // 卦象
	GuaMing     string         // 卦名
	GuaYi       string         // 卦意（邵雍）
	FuHao       string         // 符号
	GuaYaos     [YaoMax]GuaYao // 初，二，三，四，五，上
	Yong        string         // 用九,用六
	YongJiXiong string         // 用九,用六吉凶
}

func loadGuaXiang() (map[string]GuaXiang, error) {
	var data map[string]GuaXiang

	file64gua, err := DataFiles.Open("data/64gua.csv")
	if err != nil {
		return data, err
	}

	records, err := readData(file64gua)
	if err != nil {
		return data, err
	}

	for _, record := range records {
		// gxIndex, _ := strconv.ParseUint(record[1], 10, bits.UintSize)
		// if !GuaXiangIndex(gxIndex).IsValid() {
		// 	return data, fmt.Errorf("guaxiang index out of range: %d", gxIndex)
		// }

		guaxu, _ := strconv.ParseUint(record[0], 10, bits.UintSize)
		shangshu, _ := strconv.ParseUint(record[3], 10, bits.UintSize)
		xiashu, _ := strconv.ParseUint(record[5], 10, bits.UintSize)

		guaxiang := GuaXiang{
			GuaXu:       uint(guaxu),
			ShangGua:    record[2],
			ShangShu:    uint(shangshu),
			XiaGua:      record[4],
			XiaShu:      uint(xiashu),
			JiXiong:     record[6],
			GuaXiang:    record[7],
			GuaMing:     record[8],
			GuaYi:       record[9],
			FuHao:       record[10],
			Yong:        record[29],
			YongJiXiong: record[30],
		}

		for i := Yao(0); i < YaoMax; i++ {
			guyao := GuaYao{
				Yao:     record[11+i*3],
				JiXiong: record[12+i*3],
				NvMing:  record[13+i*3],
			}
			guaxiang.GuaYaos[i] = guyao
		}

		data[record[1]] = guaxiang
	}
	return data, nil
}
func getGuaXiangs() map[string]GuaXiang {
	return guaxiangData
}

func GetGuaXiang(name string) GuaXiang {
	// if !index.IsValid() {
	// 	panic(fmt.Errorf("guaxiang index out of range: %d", index))
	// }
	if gx, ok := guaxiangData[name]; ok {
		return gx
	}
	panic(fmt.Errorf("guaxiang not found: %s", name))
}
