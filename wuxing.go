package yi

const (
	Yin  = 1
	Yang = 0
)

const (
	YinStr  string = "阴"
	YangStr string = "阳"
)

type Sex int

// 性别
const (
	SexBoy  Sex = 0b01 //"男"
	SexGirl Sex = 0b10 //"女"
)

const wuXingList string = "水木木火火土土金金水"

//五行
const (
	MU int = iota + 1
	HUO
	TU
	JIN
	SHUI
)

func YinYang(i int) int {
	if i%2 == Yang {
		return Yang
	}
	return Yin
}

func YinYangAttr(i int) string {
	if YinYang(i) == Yang {
		return YangStr
	}

	return YinStr
}

// GenerateThreeTalent 计算字符的三才属性
// 1-2木：1为阳木，2为阴木
// 3-4火：3为阳火，4为阴火
// 5-6土：5为阳土，6为阴土
// 7-8金：7为阳金，8为阴金
// 9-10水：9为阳水，10为阴水
func NumberWuXing(i int) string {
	return string([]rune(wuXingList)[i%10])
}
