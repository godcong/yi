package yi

import (
	"encoding/csv"
	"math/bits"
	"os"
	"strconv"
)

//GuaXiang 卦象
type GuaXiang struct {
	ShangGua        string `bson:"shang_gua"`          //上卦
	ShangShu        int    `bson:"shang_shu"`          //上卦数
	XiaGua          string `bson:"xia_gua"`            //下卦
	XiaShu          int    `bson:"xia_shu"`            //下卦数
	JiXiong         string `bson:"ji_xiong"`           //吉凶
	GuaXiang        string `bson:"gua_xiang"`          //卦象
	GuaMing         string `bson:"gua_ming"`           //卦名
	GuaYi           string `bson:"gua_yi"`             //卦意
	GuaYun          string `bson:"gua_yun"`            //卦云
	XiangYue        string `bson:"xiang_yue"`          //象曰
	FuHao           string `bson:"fu_hao"`             //符号
	ChuYao          string `bson:"chu_yao"`            //初爻
	ChuYaoJiXiong   string `bson:"chu_yao_ji_xiong"`   //初爻吉凶
	ErYao           string `bson:"er_yao"`             //二爻
	ErYaoJiXiong    string `bson:"er_yao_ji_xiong"`    //二爻吉凶
	SanYao          string `bson:"san_yao"`            //三爻
	SanYaoJiXiong   string `bson:"san_yao_ji_xiong"`   //三爻吉凶
	SiYao           string `bson:"si_yao"`             //四爻
	SiYaoJiXiong    string `bson:"si_yao_ji_xiong"`    //四爻吉凶
	WuYao           string `bson:"wu_yao"`             //五爻
	WuYaoJiXiong    string `bson:"wu_yao_ji_xiong"`    //五爻吉凶
	ShangYao        string `bson:"shang_yao"`          //上爻
	ShangYaoJiXiong string `bson:"shang_yao_ji_xiong"` //上爻吉凶
	Yong            string `bson:"yong"`               //用九,用六
	YongJiXiong     string `bson:"yong_ji_xiong"`      //用九,用六吉凶
}

var gx map[string]*GuaXiang

func init() {
	gx = make(map[string]*GuaXiang)

	records, err := readData("data/64gua.csv")

	if err != nil {
		panic(err)
	}

	for _, record := range records {
		shangshu, err := strconv.ParseInt(record[2], 10, bits.UintSize)
		if err != nil {
			panic(err)
		}
		xiashu, err := strconv.ParseInt(record[4], 10, bits.UintSize)
		if err != nil {
			panic(err)
		}
		guaxiang := GuaXiang{
			ShangGua:        record[1],
			ShangShu:        int(shangshu),
			XiaGua:          record[3],
			XiaShu:          int(xiashu),
			JiXiong:         record[5],
			GuaXiang:        record[6],
			GuaMing:         record[7],
			GuaYi:           record[8],
			GuaYun:          record[9],
			XiangYue:        record[10],
			FuHao:           record[11],
			ChuYao:          record[12],
			ChuYaoJiXiong:   record[13],
			ErYao:           record[14],
			ErYaoJiXiong:    record[15],
			SanYao:          record[16],
			SanYaoJiXiong:   record[17],
			SiYao:           record[18],
			SiYaoJiXiong:    record[19],
			WuYao:           record[20],
			WuYaoJiXiong:    record[21],
			ShangYao:        record[22],
			ShangYaoJiXiong: record[23],
			Yong:            record[24],
			YongJiXiong:     record[25],
		}

		gx_index := record[0]
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
