package yi

import (
	"encoding/json"
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
	data, err := libDecompressStatik()
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &gx)
	if err != nil {
		panic(err)
	}
}

func GetGuaXiang() map[string]*GuaXiang {
	return gx
}

func setGuaXiang(gx map[string]*GuaXiang) {
	if data, err := json.Marshal(gx); err == nil {
		libCompress(data)
	}
}
