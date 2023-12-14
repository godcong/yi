package yi

import (
	"encoding/json"
	"os"

	"github.com/godcong/yi/source"
	_ "github.com/godcong/yi/statik"
)

// GuaXiangV1 卦象
type GuaXiangV1 struct {
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

type GuaXiang struct {
	ShangGua string      `bson:"shang_gua"` //上卦
	ShangShu int         `bson:"shang_shu"` //上卦数
	XiaGua   string      `bson:"xia_gua"`   //下卦
	XiaShu   int         `bson:"xia_shu"`   //下卦数
	JiXiong  string      `bson:"ji_xiong"`  //吉凶
	GuaXiang string      `bson:"gua_xiang"` //卦象
	GuaMing  string      `bson:"gua_ming"`  //卦名
	GuaYi    string      `bson:"gua_yi"`    //卦意
	GuaYun   string      `bson:"gua_yun"`   //卦云
	XiangYue string      `bson:"xiang_yue"` //象曰
	FuHao    string      `bson:"fu_hao"`    //符号
	Yaos     [YaoMax]Yao `bson:"yaos"`      //爻
	//ChuYao          string `bson:"chu_yao"`            //初爻
	//ChuYaoJiXiong   string `bson:"chu_yao_ji_xiong"`   //初爻吉凶
	//ErYao           string `bson:"er_yao"`             //二爻
	//ErYaoJiXiong    string `bson:"er_yao_ji_xiong"`    //二爻吉凶
	//SanYao          string `bson:"san_yao"`            //三爻
	//SanYaoJiXiong   string `bson:"san_yao_ji_xiong"`   //三爻吉凶
	//SiYao           string `bson:"si_yao"`             //四爻
	//SiYaoJiXiong    string `bson:"si_yao_ji_xiong"`    //四爻吉凶
	//WuYao           string `bson:"wu_yao"`             //五爻
	//WuYaoJiXiong    string `bson:"wu_yao_ji_xiong"`    //五爻吉凶
	//ShangYao        string `bson:"shang_yao"`          //上爻
	//ShangYaoJiXiong string `bson:"shang_yao_ji_xiong"` //上爻吉凶
	Yong        string `bson:"yong"`          //用九,用六
	YongJiXiong string `bson:"yong_ji_xiong"` //用九,用六吉凶
}

var gx map[string]*GuaXiangV1
var gxs map[string]*GuaXiang

func init() {

	gxs = make(map[string]*GuaXiang)
	//data, err := libDecompressStatik()
	//if err != nil {
	//	panic(err)
	//}
	err := json.Unmarshal(source.GuaxiangV2, &gxs)
	if err != nil {
		panic(err)
	}
}

func GetGuaXiang() map[string]*GuaXiangV1 {
	return gx
}

func GetGuaXiangV2() map[string]*GuaXiang {
	return gxs
}

func setGuaXiang(gx map[string]*GuaXiangV1) {
	if data, err := json.Marshal(gx); err == nil {
		libCompress(data)
	}
}

func setGuaXiangV2(gx map[string]*GuaXiangV1) {
	gxs := make(map[string]*GuaXiang)
	for k, v := range gx {
		gxs[k] = updateGX(v)
	}
	if data, err := json.Marshal(gxs); err == nil {
		err := os.WriteFile("source/guaxiang_v2.json", data, 0644)
		if err != nil {
			panic(err)
		}
	}
}

func updateGX(v1 *GuaXiangV1) *GuaXiang {
	return &GuaXiang{
		ShangGua: v1.ShangGua,
		ShangShu: v1.ShangShu,
		XiaGua:   v1.XiaGua,
		XiaShu:   v1.XiaShu,
		JiXiong:  v1.JiXiong,
		GuaXiang: v1.GuaXiang,
		GuaMing:  v1.GuaMing,
		GuaYi:    v1.GuaYi,
		GuaYun:   v1.GuaYun,
		XiangYue: v1.XiangYue,
		FuHao:    v1.FuHao,
		Yaos: [6]Yao{
			YaoChu: Yao{
				Index:   YaoChu,
				JiXiong: v1.ChuYaoJiXiong,
				YaoYue:  v1.ChuYao,
				YaoYi:   "",
			},
			YaoEr: Yao{
				Index:   YaoEr,
				JiXiong: v1.ErYaoJiXiong,
				YaoYue:  v1.ErYao,
				YaoYi:   "",
			},
			YaoSan: Yao{
				Index:   YaoSan,
				JiXiong: v1.SanYaoJiXiong,
				YaoYue:  v1.SanYao,
				YaoYi:   "",
			},
			YaoSi: Yao{
				Index:   YaoSi,
				JiXiong: v1.SiYaoJiXiong,
				YaoYue:  v1.SiYao,
				YaoYi:   "",
			},
			YaoWu: Yao{
				Index:   YaoWu,
				JiXiong: v1.WuYaoJiXiong,
				YaoYue:  v1.WuYao,
				YaoYi:   "",
			},
			YaoShang: Yao{
				Index:   YaoShang,
				JiXiong: v1.ShangYaoJiXiong,
				YaoYue:  v1.ShangYao,
				YaoYi:   "",
			},
		},
		Yong:        v1.Yong,
		YongJiXiong: v1.YongJiXiong,
	}

}
