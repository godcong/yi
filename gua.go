package yi

type Gua struct {
	Name            string // 卦名
	Number          uint   // 数
	Symbol          string // 符
	XianTianNumber  uint   // 先天卦数
	HouTianNumber   uint   // 后天卦数
	SymbolName      string // 符号名称
	BinaryYang      uint   // 二进（阳1）
	XianTianRemain  uint   // 先天余数
	WuXing          string // 五行
	SymbolRep       string // 卦象代表
	Attribute       string // 属性
	Position        string // 五位
	HumanAffair     string // 人事现象
	BodyPart        string // 身体部位
	Animal          string // 动物
	GeoLocation     string // 地理位置
	StillObject     string // 静物
	Season          string // 季节
	EightDoorDir    string // 八门方位
	PurpleWhiteStar uint   // 紫白九星
	TianYunStar     uint   // 天运九星
	HouseStar       uint   // 宅局九星
	FourFortunePos  string // 四吉凶位
}
