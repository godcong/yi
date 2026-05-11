package yi

// ============================================================================
// ShiYing (世应) - Shi (Self) and Ying (Response) system
// ============================================================================
//
// 世应是六爻预测中的核心定位系统，用于确定卦中各爻的归属和关系：
//
//   - 世爻（Shi）: 代表求测者本人，是卦中的核心定位点
//   - 应爻（Ying）: 代表所测之人、事、物，与世爻相对应
//
// 世应的位置由卦所在的宫（本宫）和卦在宫中的排列顺序决定。
// 每宫八卦，按照"本宫→一世→二世→三世→四世→五世→游魂→归魂"排列。
//
// 世应规则：
//   - 本宫卦（纯卦）：世在第六爻（上爻），应在第三爻
//   - 一世卦：世在第一爻（初爻），应在第四爻
//   - 二世卦：世在第二爻，应在第五爻
//   - 三世卦：世在第三爻，应在第六爻
//   - 四世卦：世在第四爻，应在第一爻
//   - 五世卦：世在第五爻，应在第二爻
//   - 游魂卦：世在第四爻，应在第一爻
//   - 归魂卦：世在第三爻，应在第六爻
//
// Reference:
//   - 《增删卜易》野鹤老人
//   - 《卜筮正宗》王洪绪

// GuaPosition represents the position of a hexagram within its palace.
type GuaPosition int

const (
	BenGong GuaPosition = iota // 本宫卦 (Pure hexagram)
	YiShi                      // 一世卦 (1st change)
	ErShi                      // 二世卦 (2nd change)
	SanShi                     // 三世卦 (3rd change)
	SiShi                      // 四世卦 (4th change)
	WuShi                      // 五世卦 (5th change)
	YouHun                     // 游魂卦 (Wandering soul)
	GuiHun                     // 归魂卦 (Returning soul)
	GuaPositionCount
)

// String returns the Chinese name of GuaPosition.
func (gp GuaPosition) String() string {
	switch gp {
	case BenGong:
		return "本宫"
	case YiShi:
		return "一世"
	case ErShi:
		return "二世"
	case SanShi:
		return "三世"
	case SiShi:
		return "四世"
	case WuShi:
		return "五世"
	case YouHun:
		return "游魂"
	case GuiHun:
		return "归魂"
	default:
		return ""
	}
}

// ShiYingInfo represents the Shi (世) and Ying (应) positions for a hexagram.
type ShiYingInfo struct {
	ShiPos   YaoPosition // 世爻 position (0-5)
	YingPos  YaoPosition // 应爻 position (0-5)
	Position GuaPosition // position in palace
}

// GetShiYing returns the Shi and Ying positions for a hexagram.
// It determines the hexagram's palace position and calculates accordingly.
func (g *Gua) GetShiYing() *ShiYingInfo {
	pos := g.GetGuaPosition()
	if pos < 0 || pos >= GuaPositionCount {
		return nil
	}
	info := &shiYingTable[pos]
	return &ShiYingInfo{
		ShiPos:   info.ShiPos,
		YingPos:  info.YingPos,
		Position: pos,
	}
}

// GetGuaPosition determines the palace position of a hexagram.
// This uses the standard 8-palace system to classify the hexagram.
func (g *Gua) GetGuaPosition() GuaPosition {
	if idx, ok := guaPositionStore[g.Index]; ok {
		return idx
	}
	return GuaPositionCount
}

// GetGuaGong returns the palace (宫) that a hexagram belongs to.
// Each palace is identified by its pure hexagram (本宫卦).
func (g *Gua) GetGuaGong() Bagua {
	if bg, ok := guaGongStore[g.Index]; ok {
		return bg
	}
	return -1
}

// shiYingTable maps GuaPosition to Shi/Ying positions.
// Index corresponds to GuaPosition constants.
var shiYingTable = [GuaPositionCount]struct {
	ShiPos  YaoPosition
	YingPos YaoPosition
}{
	BenGong: {Shang, San}, // 本宫: 世六应三
	YiShi:   {Chu, Si},    // 一世: 世初应四
	ErShi:   {Er, Wu},     // 二世: 世二应五
	SanShi:  {San, Shang}, // 三世: 世三应六
	SiShi:   {Si, Chu},    // 四世: 世四应初
	WuShi:   {Wu, Er},     // 五世: 世五应二
	YouHun:  {Si, Chu},    // 游魂: 世四应初
	GuiHun:  {San, Shang}, // 归魂: 世三应六
}

// guaPositionStore maps hexagram index to its GuaPosition.
// Populated in data_generated.go
var guaPositionStore map[string]GuaPosition

// guaGongStore maps hexagram index to its palace (Bagua).
// Populated in data_generated.go
var guaGongStore map[string]Bagua
