package core

import (
	"encoding/json"
)

//GuaXiang 卦象
type GuaXiang struct {
	ShangGua string `bson:"shang_gua"` //上卦
	ShangShu int    `bson:"shang_shu"` //上卦数
	XiaGua   string `bson:"xia_gua"`   //下卦
	XiaShu   int    `bson:"xia_shu"`   //下卦数
	HeGua    string `bson:"he_gua"`    //
	GuaXiang string `bson:"gua_xiang"` //卦象
	GuaMing  string `bson:"gua_ming"`  //卦名
	GuaYi    string `bson:"gua_yi"`    //卦意
	GuaYun   string `bson:"gua_yun"`   //卦云
	XiangYue string `bson:"xiang_yue"` //象曰
	FuHao    string `bson:"fu_hao"`    //符号
	ChuYao   string `bson:"chu_yao"`   //初爻
	ErYao    string `bson:"er_yao"`    //二爻
	SanYao   string `bson:"san_yao"`   //三爻
	SiYao    string `bson:"si_yao"`    //四爻
	WuYao    string `bson:"wu_yao"`    //五爻
	ShangYao string `bson:"shang_yao"` //上爻
	Yong     string `bson:"yong"`      //用九,用六
}

var gx map[string]*GuaXiang

func GetGuaXiang() map[string]*GuaXiang {
	return getGuaXiang()
}
func getGuaXiang() map[string]*GuaXiang {
	if gx == nil {
		gx = make(map[string]*GuaXiang)
		data, err := libDecompress()
		if err == nil {
			json.Unmarshal(data, &gx)
		}
	}
	return gx
}

func setGuaXiang(gx map[string]*GuaXiang) {
	if data, err := json.Marshal(gx); err == nil {
		libCompress(data)
	}
}
