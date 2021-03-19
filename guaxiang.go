package yi

import (
	"encoding/csv"
	"math/bits"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
)

//GuaXiang 卦象
type GuaXiang struct {
	GuaXu           int    //卦序
	ShangGua        string //上卦
	ShangShu        int    //上卦数
	XiaGua          string //下卦
	XiaShu          int    //下卦数
	JiXiong         string //吉凶（？）
	GuaXiang        string //卦象
	GuaMing         string //卦名
	GuaYi           string //卦意（邵雍）
	FuHao           string //符号
	ChuYao          string //初爻
	ChuYaoJiXiong   string //初爻吉凶
	ChuYaoNvMing    string //初爻女命
	ErYao           string //二爻
	ErYaoJiXiong    string //二爻吉凶
	ErYaoNvMing     string //二爻女命
	SanYao          string //三爻
	SanYaoJiXiong   string //三爻吉凶
	SanYaoNvMing    string //三爻女命
	SiYao           string //四爻
	SiYaoJiXiong    string //四爻吉凶
	SiYaoNvMing     string //四爻女命
	WuYao           string //五爻
	WuYaoJiXiong    string //五爻吉凶
	WuYaoNvMing     string //五爻女命
	ShangYao        string //上爻
	ShangYaoJiXiong string //上爻吉凶
	ShangYaoNvMing  string //上爻女命
	Yong            string //用九,用六
	YongJiXiong     string //用九,用六吉凶
}

var gx map[string]*GuaXiang

func init() {
	gx = make(map[string]*GuaXiang)

	var (
		_, b, _, _ = runtime.Caller(0)
		basepath   = filepath.Dir(b)
	)
	records, err := readData(filepath.Join(basepath, "data/64gua.csv"))

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
			GuaXu:           int(guaxu),
			ShangGua:        record[2],
			ShangShu:        int(shangshu),
			XiaGua:          record[4],
			XiaShu:          int(xiashu),
			JiXiong:         record[6],
			GuaXiang:        record[7],
			GuaMing:         record[8],
			GuaYi:           record[9],
			FuHao:           record[10],
			ChuYao:          record[11],
			ChuYaoJiXiong:   record[12],
			ChuYaoNvMing:    record[13],
			ErYao:           record[14],
			ErYaoJiXiong:    record[15],
			ErYaoNvMing:     record[16],
			SanYao:          record[17],
			SanYaoJiXiong:   record[18],
			SanYaoNvMing:    record[19],
			SiYao:           record[20],
			SiYaoJiXiong:    record[21],
			SiYaoNvMing:     record[22],
			WuYao:           record[23],
			WuYaoJiXiong:    record[24],
			WuYaoNvMing:     record[25],
			ShangYao:        record[26],
			ShangYaoJiXiong: record[27],
			ShangYaoNvMing:  record[28],
			Yong:            record[29],
			YongJiXiong:     record[30],
		}

		gx_index := record[1]
		if len(gx_index) < 1 {
			panic("index is wrong")
		}

		gx[gx_index] = &guaxiang
	}
}

func GetGuaXiang() map[string]*GuaXiang {
	return gx
}

func readData(fileName string) ([][]string, error) {

	f, err := os.Open(fileName)

	if err != nil {
		return [][]string{}, err
	}
	defer f.Close()

	r := csv.NewReader(f)
	r.Comma = ','
	r.Comment = '#'

	// skip first line
	if _, err := r.Read(); err != nil {
		return [][]string{}, err
	}

	records, err := r.ReadAll()

	if err != nil {
		return [][]string{}, err
	}

	return records, nil
}
