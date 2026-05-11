package yi

import (
	"math/rand"
	"time"
)

// ============================================================================
// 起卦算法 (QiGua) - Divination Methods
// ============================================================================
//
// 本模块实现了四种经典起卦方法：
//   1. 铜钱法（三钱法）  — 最常用，六次抛铜钱
//   2. 蓍草法（大衍法）  — 最古老，《周易》本法，50根蓍草
//   3. 梅花易数法         — 邵雍创，以数起卦
//   4. 时间起卦法         — 按年月日时自动计算上下卦和动爻
//
// 参考：《周易》、邵雍《梅花易数》、京房纳甲

// ============================================================================
// 1. 铜钱法（三钱法）
// ============================================================================
//
// 算法：抛3枚铜钱6次，记录每次的阴面数量
//   3阴（3个背）→ 老阴(6) → 变爻
//   2阴1阳（2个背）→ 少阴(7)
//   1阴2阳（1个背）→ 少阳(8)
//   0阴3阳（0个背）→ 老阳(9) → 变爻
//
// 爻值：6=老阴(变), 7=少阴, 8=少阳, 9=老阳(变)
// 八卦映射：1阳爻=1，0阴爻=0，从初爻到上爻构成二进制数

// CoinResult represents the result of one coin throw (3 coins).
type CoinResult int

const (
	CoinYinYinYin    CoinResult = 6 // 三阴 = 老阴(动)
	CoinYinYinYang   CoinResult = 7
	CoinYinYangYang  CoinResult = 8
	CoinYangYangYang CoinResult = 9 // 三阳 = 老阳(动)
)

// String implements fmt.Stringer.
func (c CoinResult) String() string {
	switch c {
	case CoinYinYinYin:
		return "老阴"
	case CoinYinYinYang:
		return "少阴"
	case CoinYinYangYang:
		return "少阳"
	case CoinYangYangYang:
		return "老阳"
	default:
		return "未知"
	}
}

// IsChanging returns whether this is a changing line (老阴 or 老阳).
func (c CoinResult) IsChanging() bool {
	return c == CoinYinYinYin || c == CoinYangYangYang
}

// IsYang returns true if this is a yang line (少阳 or 老阳).
func (c CoinResult) IsYang() bool {
	return c == CoinYinYangYang || c == CoinYangYangYang
}

// YaoValue converts CoinResult to yao value (6/7/8/9).
func (c CoinResult) YaoValue() int {
	return int(c)
}

// throwCoins simulates throwing 3 coins and returns the number of yin (背) sides.
// A real coin has 1 yang (字) side and 1 yin (背) side.
// This simulates: 0-1 yin=1阳2阴, 2 yin=1阳2阴, 3 yin=全阴, -1=全阳 (0 yin)
func throwCoins(rng *rand.Rand) CoinResult {
	yin := rng.Intn(4) // 0,1,2,3
	switch yin {
	case 0:
		return CoinYangYangYang // 0 yin = 3阳 = 老阳(9)
	case 1:
		return CoinYinYangYang // 1 yin = 少阳(8)
	case 2:
		return CoinYinYinYang // 2 yin = 少阴(7)
	case 3:
		return CoinYinYinYin // 3 yin = 老阴(6)
	}
	return CoinYinYinYang
}

// DivineByCoins performs divination using the coin (三钱) method.
// It takes a source of randomness for simulating coin throws.
// Returns the hexagram result and the 6 coin throws for reference.
func DivineByCoins(seed int64) (*ZhouYi, [6]CoinResult) {
	rng := rand.New(rand.NewSource(seed))
	var results [6]CoinResult
	var shangBits, xiaBits int

	for i := 0; i < 6; i++ {
		cr := throwCoins(rng)
		results[i] = cr
		// 9/老阳=阳, 8/少阳=阳, 7/少阴=阴, 6/老阴=阴
		// 在卦象中: 阴爻为断(1), 阳爻为连(0)
		isYin := cr == CoinYinYinYang || cr == CoinYinYinYin
		if isYin {
			// 从初爻到上爻: 第0次=初爻(bit0), 第5次=上爻(bit5)
			xiaBits |= 1 << i
		}
	}

	shangBits = xiaBits >> 3
	xiaBits = xiaBits & 0x7

	// Calculate changing lines
	var bianYao []int
	for i, cr := range results {
		if cr.IsChanging() {
			bianYao = append(bianYao, i)
		}
	}

	return DivineByNumber(shangBits, xiaBits, bianYao...), results
}

// DivineByCoinsRand is DivineByCoins using current time as seed.
func DivineByCoinsRand() (*ZhouYi, [6]CoinResult) {
	return DivineByCoins(time.Now().UnixNano())
}

// CoinThrowSimulate simulates one coin throw (for manual input).
// Pass the actual yin count: 0=yang, 1=yang, 2=yin, 3=yin.
func CoinThrowSimulate(yinCount int) CoinResult {
	switch yinCount % 4 {
	case 0:
		return CoinYangYangYang // 0 yin
	case 1:
		return CoinYinYangYang // 1 yin
	case 2:
		return CoinYinYinYang // 2 yin
	case 3:
		return CoinYinYinYin // 3 yin
	}
	return CoinYinYinYang
}

// DivineByCoinValues performs divination from 6 coin results (6/7/8/9 values).
// Values: 6=老阴, 7=少阴, 8=少阳, 9=老阳.
// Changing lines are detected automatically from 6 and 9.
func DivineByCoinValues(values [6]int) (*ZhouYi, error) {
	if len(values) != 6 {
		return nil, ErrInvalidCoinValues
	}
	var shangBits, xiaBits int
	var bianYao []int

	for i, v := range values {
		if v < 6 || v > 9 {
			return nil, ErrInvalidCoinValues
		}
		isChanging := v == 6 || v == 9
		isYin := v == 6 || v == 7
		if isYin {
			xiaBits |= 1 << i
		}
		if isChanging {
			bianYao = append(bianYao, i)
		}
	}

	shangBits = xiaBits >> 3
	xiaBits = xiaBits & 0x7

	return DivineByNumber(shangBits, xiaBits, bianYao...), nil
}

// ============================================================================
// 2. 蓍草法（大衍法）
// ============================================================================
//
// 算法（《周易·系辞》）：
//   "大衍之数五十，其用四十有九。分二以象两，卦一以象三，
//    揲之以四以象四时，归奇于扐以象闰，五岁再闰，故再扐而后挂。"
//
//   步骤（每得一爻，重复3次）：
//     1. 49根蓍草，随机分为两堆 ← "分二以象两"
//     2. 从右边一堆取出1根挂于左手小指间 ← "卦一"
//     3. 右手持左边一堆，以4根为单位分组，余数（1-4）放在一旁 ← "揲之以四"
//     4. 右手持右边一堆（不含挂的那根），同样以4根分组，余数放在一旁 ← "揲之以四"
//     5. 挂的1根 + 左边余数 + 右边余数 = 挂扐数（必为5或9）
//     6. 49 - 挂扐数 = 本次剩余（必为40或44）
//     7. 重复步骤1-6，共3次，得到剩余数 R
//        R=36 → 9（一阳爻），R=40 → 8（少阳）
//        R=44 → 7（少阴），R=37 → 6（老阴/动）
//     8. 6次得到6爻，构成卦象
//
// 爻值：9=老阳(变), 8=少阳, 7=少阴, 6=老阴(变)

// DayanStep represents one step in the Dayan (yarrow) algorithm.
type DayanStep struct {
	Total     int // total remaining before this step
	Left      int // left pile count
	Right     int // right pile count (after removing 1)
	Remainder int // total remainder (left_rem + right_rem + 1)
	Final     int // final remaining count after this step
}

// DayanResult represents the result of one yarrow line calculation.
type DayanResult struct {
	Steps      [3]DayanStep // the 3 sub-steps
	Remaining  int          // final remaining count (36/37/40/44)
	YaoValue   int          // resulting yao value (6/7/8/9)
	IsChanging bool         // whether this is a changing line
}

// dayanDivide simulates the first division of yarrow stalks.
func dayanDivide(rng *rand.Rand, total int) (left, right int) {
	// Divide into two piles, each at least 1
	minLeft := 1
	maxLeft := total - 1
	left = minLeft + rng.Intn(maxLeft-minLeft+1)
	right = total - left
	return
}

// dayanYuQi calculates remainder when dividing by 4.
func dayanYuQi(count int) int {
	return count % 4
}

// dayanSimOneYao simulates one yarrow line (揲蓍9次 → 1爻).
func dayanSimOneYao(rng *rand.Rand) DayanResult {
	const totalInitial = 49

	remaining := totalInitial
	var steps [3]DayanStep

	for s := 0; s < 3; s++ {
		// Step 1: divide into two piles
		left, right := dayanDivide(rng, remaining)
		// Step 2: remove 1 for the "挂"
		right -= 1
		if right < 0 {
			right = 0
		}
		// Step 3 & 4: count remainder by 4
		leftRem := 1 + dayanYuQi(left-1) // 1-4
		if leftRem > left {
			leftRem = left
		}
		rightRem := 1 + dayanYuQi(right-1) // 1-4
		if rightRem > right {
			rightRem = right
		}
		// Total remainder (挂扐): always 5 or 9
		remainder := 1 + leftRem + rightRem // 1=挂, leftRem+rightRem=4 or 8
		steps[s] = DayanStep{
			Total:     remaining,
			Left:      left,
			Right:     right,
			Remainder: remainder,
			Final:     remaining - remainder,
		}
		remaining = steps[s].Final
	}

	// Map remaining count to yao value
	var yaoValue int
	var isChanging bool
	switch remaining {
	case 36:
		yaoValue, isChanging = 9, true // 老阳(动)
	case 40:
		yaoValue, isChanging = 8, false // 少阳
	case 44:
		yaoValue, isChanging = 7, false // 少阴
	case 37:
		yaoValue, isChanging = 6, true // 老阴(动)
	default:
		// In practice this shouldn't happen with correct random division
		// But for safety, map any unexpected value
		if remaining > 40 {
			yaoValue, isChanging = 8, false
		} else {
			yaoValue, isChanging = 7, false
		}
	}

	return DayanResult{Steps: steps, Remaining: remaining, YaoValue: yaoValue, IsChanging: isChanging}
}

// DivineByDayan performs divination using the Dayan (yarrow stalk) method.
// It takes a source of randomness.
// Returns the hexagram result and 6 DayanResult for reference.
func DivineByDayan(seed int64) (*ZhouYi, [6]DayanResult) {
	rng := rand.New(rand.NewSource(seed))
	var yaoValues [6]int
	var results [6]DayanResult
	var shangBits, xiaBits int

	for i := 0; i < 6; i++ {
		dr := dayanSimOneYao(rng)
		results[i] = dr
		yaoValues[i] = dr.YaoValue

		isYin := dr.YaoValue == 6 || dr.YaoValue == 7
		if isYin {
			xiaBits |= 1 << i
		}
	}

	shangBits = xiaBits >> 3
	xiaBits = xiaBits & 0x7

	var bianYao []int
	for i, v := range yaoValues {
		if v == 6 || v == 9 {
			bianYao = append(bianYao, i)
		}
	}

	return DivineByNumber(shangBits, xiaBits, bianYao...), results
}

// DivineByDayanRand is DivineByDayan using current time as seed.
func DivineByDayanRand() (*ZhouYi, [6]DayanResult) {
	return DivineByDayan(time.Now().UnixNano())
}

// DivineByDayanValues performs divination from 6 Dayan yao values (6/7/8/9).
func DivineByDayanValues(values [6]int) (*ZhouYi, error) {
	if len(values) != 6 {
		return nil, ErrInvalidDayanValues
	}
	var shangBits, xiaBits int
	var bianYao []int

	for i, v := range values {
		if v < 6 || v > 9 {
			return nil, ErrInvalidDayanValues
		}
		isYin := v == 6 || v == 7
		if isYin {
			xiaBits |= 1 << i
		}
		if v == 6 || v == 9 {
			bianYao = append(bianYao, i)
		}
	}

	shangBits = xiaBits >> 3
	xiaBits = xiaBits & 0x7

	return DivineByNumber(shangBits, xiaBits, bianYao...), nil
}

// ============================================================================
// 3. 梅花易数法 (Plum Blossom Numerology by Shao Yong)
// ============================================================================
//
// 邵雍《梅花易数》核心公式：
//   上卦 = (上数 + 下数) mod 8   （上卦数从1起算）
//   下卦 = (上数 + 下数 + 变数) mod 8  （从1起算，0视为8）
//   动爻 = (上数 + 下数 + 变数 + 时辰) mod 6  （从1起算，0视为6）
//
// 传统梅花以"数"起卦，数可以是任意数（字数、年龄、金额等）。
// 常见用法：
//   - 报数起卦：直接用报出的两个数（上卦数、下卦数）
//   - 时间起卦：(年+月+日) mod 8 = 上卦, (年+月+日+时) mod 8 = 下卦
//   - 字数起卦：总字数分上下卦
//   - 方位起卦：方位数分上下卦
//
// toBagua: 将1-8映射到 Bagua 常量 (0=Qian, 1=Dui, ..., 7=Kun)
func shaoYongToBagua(n int) Bagua {
	// n从1起算: 1-8 → 0-7 (Qian-Kun)
	return Bagua((n - 1 + 8) % 8)
}

// DivineByMeihua performs divination using the Plum Blossom Numerology method.
// shang, xia: the two numbers for upper and lower trigrams (any positive integers).
// bian: optional changing line number (0-5), pass -1 for auto from time.
func DivineByMeihua(shang, xia int, bian int) *ZhouYi {
	// Apply Shao Yong formula: 上卦/下卦从1起算
	upper := shaoYongToBagua(shang)
	lower := shaoYongToBagua(xia)

	var bYao []int
	if bian >= 0 && bian <= 5 {
		bYao = []int{bian}
	}

	return DivineByNumber(int(upper), int(lower), bYao...)
}

// DivineByMeihuaTime performs Plum Blossom divination using current time.
// Formula: 上卦=(年+月+日)mod8, 下卦=(年+月+日+时)mod8, 动爻=(年+月+日+时)mod6
// All values are taken as the calendar year/month/day/hour number.
// An optional personalSeed can be provided for per-user personalization.
// Returns hexagram and the 4 time components for reference.
func DivineByMeihuaTime(t time.Time, personalSeed ...string) (*ZhouYi, int, int, int, int) {
	year := t.Year()
	month := int(t.Month())
	day := t.Day()

	// 梅花易数时辰: 子=1, 丑=2, ..., 亥=12
	dzHour := shichen(t.Hour())

	sum := year + month + day
	upper := shaoYongToBagua(sum)
	lower := shaoYongToBagua(sum + dzHour)
	bian := (sum + dzHour) % 6

	// If a personal seed is provided, use coin divination with a mixed seed
	if len(personalSeed) > 0 && personalSeed[0] != "" {
		seed := int64(sum*10000 + dzHour)
		for _, c := range personalSeed[0] {
			seed = seed*31 + int64(c)
		}
		zy, _ := DivineByCoins(seed)
		return zy, year, month, day, dzHour
	}

	return DivineByNumber(int(upper), int(lower), bian), year, month, day, dzHour
}

// shichen converts hour to traditional Chinese hour (时辰), 1-12.
func shichen(hour int) int {
	// 子时=23:00-01:00 → 1, 丑时=1-3 → 2, ...
	// In 24-hour: 23-0 → 子(1), 1-3 → 丑(2), ...
	if hour >= 23 || hour < 1 {
		return 1 // 子时
	}
	// hour=1,2 → 丑(2), hour=3,4 → 寅(3), ..., hour=21,22 → 亥(12)
	return (hour+1)/2 + 1
}

// ============================================================================
// 4. 时间起卦法
// ============================================================================
//
// 时间起卦是最常用的自动起卦方法之一，只需提供年月日时即可：
//   上卦 = 年数 + 月数 + 日数 的和除以8的余数
//   下卦 = 上数 + 下数 + 时数 的和除以8的余数
//   动爻 = 上数 + 下数 + 时数 的和除以6的余数
//
// 所有数取农历或公历均可，只要前后一致。
// 京房易学中，年数取地支数（1-12），本方法取年本身数字（更简单）。

// TimeGuaParams holds the parameters for time-based divination.
type TimeGuaParams struct {
	Year  int // year (any positive integer, e.g., 2024)
	Month int // month (1-12)
	Day   int // day (1-31)
	Hour  int // hour (0-23 or 1-12 for traditional)
}

// DivineByTimeGua performs divination from time parameters.
// Uses the formula:
//   shang = (year + month + day) % 8
//   xia   = (year + month + day + hour) % 8
//   bian  = (year + month + day + hour) % 6
//
// An optional personalSeed can be provided to personalize the result
// (e.g., user ID, name, or question hash), ensuring different users
// or questions get different hexagrams even at the same time.
func DivineByTimeGua(params TimeGuaParams, personalSeed ...string) *ZhouYi {
	sum := params.Year + params.Month + params.Day
	shang := shaoYongToBagua(sum)
	lower := shaoYongToBagua(sum + params.Hour)
	bian := (sum + params.Hour) % 6

	// If a personal seed is provided, use coin divination with a mixed seed
	// to ensure per-user uniqueness while preserving the time-based structure
	if len(personalSeed) > 0 && personalSeed[0] != "" {
		seed := int64(sum*10000 + params.Hour)
		for _, c := range personalSeed[0] {
			seed = seed*31 + int64(c)
		}
		zy, _ := DivineByCoins(seed)
		return zy
	}

	return DivineByNumber(int(shang), int(lower), bian)
}

// DivineByCurrentTime performs divination using the current time.
// An optional personalSeed can be provided to personalize the result.
// Shortcut for DivineByTimeGua(TimeGuaParams{Year: time.Now().Year(), ...}, personalSeed...).
func DivineByCurrentTime(personalSeed ...string) *ZhouYi {
	now := time.Now()
	return DivineByTimeGua(TimeGuaParams{
		Year:  now.Year(),
		Month: int(now.Month()),
		Day:   now.Day(),
		Hour:  now.Hour(),
	}, personalSeed...)
}

// DivineByDailyHexagram performs daily hexagram divination that is stable
// for the same person on the same day.
// Unlike time-based methods that include the hour, this uses only the date
// (year+month+day) combined with a personal seed, ensuring the same person
// gets the same hexagram all day regardless of when they ask.
// The personalSeed is required and should be a stable user identifier
// (e.g., name, phone number hash, user ID). It is hashed with the date
// to produce a deterministic daily result.
func DivineByDailyHexagram(year, month, day int, personalSeed string) *ZhouYi {
	if personalSeed == "" {
		// Fallback to time-based without hour if no seed provided
		params := TimeGuaParams{Year: year, Month: month, Day: day, Hour: 0}
		return DivineByTimeGua(params)
	}

	// Hash the seed with date for deterministic daily result
	// Uses FNV-1a-like mixing: date provides the "time" component,
	// personalSeed provides the "person" component.
	seed := int64(year*10000+month*100+day) * 31
	for _, c := range personalSeed {
		seed = seed*31 + int64(c)
	}

	// Use coin method for rich randomness while being deterministic
	zy, _ := DivineByCoins(seed)
	return zy
}

// DivineByDailyHexagramNow is a shortcut for DivineByDailyHexagram using today's date.
func DivineByDailyHexagramNow(personalSeed string) *ZhouYi {
	now := time.Now()
	return DivineByDailyHexagram(now.Year(), int(now.Month()), now.Day(), personalSeed)
}

// DivineByLunarTime performs divination using lunar calendar time parameters.
// lunarYear: 农历年, lunarMonth: 农历月(1-12), lunarDay: 农历日(1-30), hour: 时辰(1-12)
func DivineByLunarTime(lunarYear, lunarMonth, lunarDay, shichen int) *ZhouYi {
	sum := lunarYear + lunarMonth + lunarDay
	shang := shaoYongToBagua(sum)
	lower := shaoYongToBagua(sum + shichen)
	bian := (sum + shichen) % 6

	return DivineByNumber(int(shang), int(lower), bian)
}

// ============================================================================
// Error definitions
// ============================================================================

var (
	ErrInvalidCoinValues  = newError("coin values must be 6 integers each in range 6-9")
	ErrInvalidDayanValues = newError("dayan values must be 6 integers each in range 6-9")
)

func newError(msg string) error {
	return &divinationError{msg: msg}
}

type divinationError struct{ msg string }

func (e *divinationError) Error() string { return e.msg }
