package yi

type Gua struct {
	GuaMing         any //卦名
	Shu             any //数
	FuHao           any //符
	XianTianGuaShu  any //先天卦数
	HouTianGuaShu   any //后天卦数
	FuHaoName       any //符号名称
	ErJing          any //二进
	XianTianYuShu   any //先天余数
	WuXing          any //五行
	GuaXiangDaiBiao any //卦象代表
	ShuXing         any //属性
	WuWei           any //五位
	RenShiXianXiang any //人事现象
	ShenTiBuWei     any //身体部位
	DongWu          any //动物
	DiLiWeiZhi      any //地理位置
	JingWu          any //静物
	JiJie           any //季节
	BanMenFangWei   any //八门方位
	ZiBaiJiuXing    any //紫白九星
	TianYunJiuXing  any //天运九星
	ZhaiJuJiuXing   any //宅局九星
	SiJiXiongWei    any //四吉凶位
}

type Gua64 struct {
	GuaXu           any //卦序
	SuoYin          any //索引
	ShangGua        any //上卦
	ShangShu        any //上数
	XiaGua          any //下卦
	XiaShu          any //下数
	JiXiong         any //吉凶
	GuaXiang        any //卦象
	GuaMing         any //卦名
	GuaYi           any //卦意
	FuHao           any //符号
	ChuYao          any //初爻
	ChuYaoJiXiong   any //初爻吉凶
	ChuNvMing       any //初女命
	ErYao           any //二爻
	ErYaoJiXiong    any //二爻吉凶
	ErNvMing        any //二女命
	SanYao          any //三爻
	SanYaoJiXiong   any //三爻吉凶
	SanNvMing       any //三女命
	SiYao           any //四爻
	SiYaoJiXiong    any //四爻吉凶
	SiNvMing        any //四女命
	WuYao           any //五爻
	WuYaoJiXiong    any //五爻吉凶
	WuNvMing        any //五女命
	ShangYao        any //上爻
	ShangYaoJiXiong any //上爻吉凶
	ShangNvMing     any //六女命
	YongJiuLiu      any //用九六
	YongJiXiong     any //用吉凶
}
