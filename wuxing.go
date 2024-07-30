package yi

// DAO 道分阴阳
// ENUM(YING,YANG)
type DAO bool

const (
	Yin  DAO = false // 阴
	Yang DAO = true  // 阳
)

type Gender int

// 性别
const (
	GenderBoy  Gender = 0b01 // "男"
	GenderGirl Gender = 0b10 // "女"
)

const wuXingList string = "水木木火火土土金金水"

// 五行
const (
	MU   int = iota + 1 // 木
	HUO                 // 火
	TU                  // 土
	JIN                 // 金
	SHUI                // 水
)

func NumberToDAO(i int) DAO {
	if i%2 == 0 {
		return Yang
	}
	return i%2 == 0
}

func (yy DAO) String() string {
	if yy {
		return "阳"
	}
	return "阴"
}

// NumberWuXing 计算字符的五行属性
// 1-2木：1为阳木，2为阴木
// 3-4火：3为阳火，4为阴火
// 5-6土：5为阳土，6为阴土
// 7-8金：7为阳金，8为阴金
// 9-10水：9为阳水，10为阴水
func NumberWuXing(i uint) string {
	return string([]rune(wuXingList)[i%10])
}
