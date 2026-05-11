package yi

// ============================================================================
// Yin-Yang
// ============================================================================

// YinYang represents Yin-Yang type
type YinYang int

const (
	Yang YinYang = iota // Yang (0)
	Yin                 // Yin (1)
)

// String returns Yin-Yang string
func (yy YinYang) String() string {
	if yy == Yang {
		return "阳"
	}
	return "阴"
}

// IsYang returns whether it is Yang
func (yy YinYang) IsYang() bool {
	return yy == Yang
}

// IsYin returns whether it is Yin
func (yy YinYang) IsYin() bool {
	return yy == Yin
}

// ============================================================================
// Sex
// ============================================================================

// Sex represents gender
type Sex int

const (
	Male   Sex = 0b01 // Male
	Female Sex = 0b10 // Female
)

// String returns gender string
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

// IsMale returns whether it is male
func (s Sex) IsMale() bool {
	return s == Male
}

// IsFemale returns whether it is female
func (s Sex) IsFemale() bool {
	return s == Female
}

// ============================================================================
// WuXing (Five Elements)
// ============================================================================

// WuXing represents Five Elements type
type WuXing int

const (
	Wood  WuXing = iota + 1 // Wood
	Fire                    // Fire
	Earth                   // Earth
	Metal                   // Metal
	Water                   // Water
)

// String returns Five Elements string
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

// WuXingList Five Elements list (for calculation)
// Order: Wood, Wood, Fire, Fire, Earth, Earth, Metal, Metal, Water, Water
// 1-2 Wood, 3-4 Fire, 5-6 Earth, 7-8 Metal, 9-10 Water
const wuXingList = "木木火火土土金金水水"

// GetWuXingByNumber returns Five Elements by number
// 1-2 Wood, 3-4 Fire, 5-6 Earth, 7-8 Metal, 9-10 Water
// Returns: Yang Wood / Yin Wood / Yang Fire / Yin Fire / Yang Earth / Yin Earth / Yang Metal / Yin Metal / Yang Water / Yin Water
func GetWuXingByNumber(n int) string {
	idx := n % 10
	if idx == 0 {
		idx = 10
	}
	base := string([]rune(wuXingList)[idx-1])

	// Determine Yin-Yang
	if idx%2 == 1 {
		return "阳" + base
	}
	return "阴" + base
}

// GetYinYangByNumber returns Yin-Yang by number
// Odd numbers are Yang, even numbers are Yin
func GetYinYangByNumber(n int) YinYang {
	if n%2 == 0 {
		return Yin
	}
	return Yang
}

// ============================================================================
// Five Elements generation and control relationships
// ============================================================================

// Sheng returns what I generate (generating)
// Wood generates Fire, Fire generates Earth, Earth generates Metal, Metal generates Water, Water generates Wood
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

// Ke returns what I control (controlling)
// Wood controls Earth, Earth controls Water, Water controls Fire, Fire controls Metal, Metal controls Wood
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

// BeiSheng returns what generates me (being generated)
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

// BeiKe returns what controls me (being controlled)
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
