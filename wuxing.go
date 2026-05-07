package yi

// ============================================================================
// 阴阳 (Yin-Yang)
// ============================================================================

// YinYang 阴阳类型
type YinYang int

const (
	Yang YinYang = iota // 阳 (0)
	Yin                 // 阴 (1)
)

// String 返回阴阳字符串
func (yy YinYang) String() string {
	if yy == Yang {
		return "阳"
	}
	return "阴"
}

// IsYang 是否为阳
func (yy YinYang) IsYang() bool {
	return yy == Yang
}

// IsYin 是否为阴
func (yy YinYang) IsYin() bool {
	return yy == Yin
}

// ============================================================================
// 性别 (Sex)
// ============================================================================

// Sex 性别
type Sex int

const (
	Male   Sex = 0b01 // 男
	Female Sex = 0b10 // 女
)

// String 返回性别字符串
func (s Sex) String() string {
	switch s {
	case Male:
		return "男"
	case Female:
		return "女"
	default:
		return "未知"
	}
}

// IsMale 是否为男
func (s Sex) IsMale() bool {
	return s == Male
}

// IsFemale 是否为女
func (s Sex) IsFemale() bool {
	return s == Female
}

// ============================================================================
// 五行 (WuXing)
// ============================================================================

// WuXing 五行类型
type WuXing int

const (
	Wood  WuXing = iota + 1 // 木
	Fire                    // 火
	Earth                   // 土
	Metal                   // 金
	Water                   // 水
)

// String 返回五行字符串
func (wx WuXing) String() string {
	switch wx {
	case Wood:
		return "木"
	case Fire:
		return "火"
	case Earth:
		return "土"
	case Metal:
		return "金"
	case Water:
		return "水"
	default:
		return ""
	}
}

// WuXingList 五行列表 (用于计算)
// 顺序: 水木木火火土土金金水
const wuXingList = "水木木火火土土金金水"

// GetWuXingByNumber 根据数字获取五行
// 1-2木, 3-4火, 5-6土, 7-8金, 9-10水
// 返回: 阳木/阴木/阳火/阴火/阳土/阴土/阳金/阴金/阳水/阴水
func GetWuXingByNumber(n int) string {
	idx := n % 10
	if idx == 0 {
		idx = 10
	}
	base := string([]rune(wuXingList)[idx-1])

	// 判断阴阳
	if idx%2 == 1 {
		return "阳" + base
	}
	return "阴" + base
}

// GetYinYangByNumber 根据数字获取阴阳
// 奇数为阳，偶数为阴
func GetYinYangByNumber(n int) YinYang {
	if n%2 == 0 {
		return Yin
	}
	return Yang
}

// ============================================================================
// 五行生克关系
// ============================================================================

// Sheng 生 (我生者)
// 木生火, 火生土, 土生金, 金生水, 水生木
func (wx WuXing) Sheng() WuXing {
	switch wx {
	case Wood:
		return Fire
	case Fire:
		return Earth
	case Earth:
		return Metal
	case Metal:
		return Water
	case Water:
		return Wood
	default:
		return 0
	}
}

// Ke 克 (我克者)
// 木克土, 土克水, 水克火, 火克金, 金克木
func (wx WuXing) Ke() WuXing {
	switch wx {
	case Wood:
		return Earth
	case Earth:
		return Water
	case Water:
		return Fire
	case Fire:
		return Metal
	case Metal:
		return Wood
	default:
		return 0
	}
}

// BeiSheng 被生 (生我者)
func (wx WuXing) BeiSheng() WuXing {
	switch wx {
	case Wood:
		return Water
	case Fire:
		return Wood
	case Earth:
		return Fire
	case Metal:
		return Earth
	case Water:
		return Metal
	default:
		return 0
	}
}

// BeiKe 被克 (克我者)
func (wx WuXing) BeiKe() WuXing {
	switch wx {
	case Wood:
		return Metal
	case Fire:
		return Water
	case Earth:
		return Wood
	case Metal:
		return Fire
	case Water:
		return Earth
	default:
		return 0
	}
}
