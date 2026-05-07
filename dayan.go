package yi

import "errors"

// Dayan 大衍之数 (81数理)
// 用于姓名笔画数吉凶分析
type Dayan struct {
	Number     int    // 笔画数 (1-81)
	JiXiong    string // 吉凶
	NvMing     string // 女命判断
	IsMax      bool   // 是否最吉
	Gua        string // 对应卦象
	TianJiu    string // 天九 (如: 太极之数)
	YiXiang    string // 义象 (描述)
	Foundation string // 基业
	Family     string // 家庭
	Health     string // 健康
	Meaning    string // 含义详解
}

// 错误定义
var (
	ErrInvalidDayanNumber = errors.New("dayan number must be between 1 and 81")
)

// GetDayan 获取大衍之数
// number: 1-81 之间的整数
func GetDayan(number int) (*Dayan, error) {
	if number < 1 || number > 81 {
		return nil, ErrInvalidDayanNumber
	}
	return &dayanList[number-1], nil
}

// MustGetDayan 获取大衍之数 (忽略错误)
// 如果 number 不合法，返回零值
func MustGetDayan(number int) Dayan {
	if number < 1 || number > 81 {
		return Dayan{}
	}
	return dayanList[(number-1)%81]
}

// IsJi 是否为吉数
func (d *Dayan) IsJi() bool {
	return d.JiXiong == "吉" || d.JiXiong == "半吉"
}

// IsXiong 是否为凶数
func (d *Dayan) IsXiong() bool {
	return d.JiXiong == "凶"
}

// IsSuitableForFemale 是否适合女性
func (d *Dayan) IsSuitableForFemale() bool {
	return d.NvMing != "凶"
}

// IsBest 是否为最吉之数
func (d *Dayan) IsBest() bool {
	return d.IsMax
}

// dayanList 大衍之数列表 (81条)
// 由 go generate 生成
var dayanList [81]Dayan

// init 初始化 (数据由 go generate 填充)
func init() {
	// 实际数据在 data_generated.go 中
}
