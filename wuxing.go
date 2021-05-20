package yi

import "sort"

type YingYang int

const (
	Yin  YingYang = 1 //阴
	Yang YingYang = 0 //阳
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
	MU   int = iota + 1 //木
	HUO                 //火
	TU                  //土
	JIN                 //金
	SHUI                //水
)

func ModeYinYang(i int) YingYang {
	if i%2 == 0 {
		return Yang
	}
	return Yin
}

func (yy YingYang) String() string {
	if yy == 0 {
		return YangStr
	}
	return YinStr
}

// NumberWuXing 计算字符的三才属性
// 1-2木：1为阳木，2为阴木
// 3-4火：3为阳火，4为阴火
// 5-6土：5为阳土，6为阴土
// 7-8金：7为阳金，8为阴金
// 9-10水：9为阳水，10为阴水
func NumberWuXing(i int) string {
	return string([]rune(wuXingList)[i%10])
}

type WuXingStr = string

const (
	WuXingJin  WuXingStr = "金"
	WuXingMu   WuXingStr = "木"
	WuXingShui WuXingStr = "水"
	WuXingHuo  WuXingStr = "火"
	WuXingTu   WuXingStr = "土"
)

type WuXingChar = rune

const (
	WuXingJinChar  WuXingChar = '金'
	WuXingMuChar   WuXingChar = '木'
	WuXingShuiChar WuXingChar = '水'
	WuXingHuoChar  WuXingChar = '火'
	WuXingTuChar   WuXingChar = '土'
)

//单个五行的计数信息
type WuXingValue struct {
	WuXingChr rune
	Count     int
}

//五行的计数信息
type WuXingCount struct {
	Index   map[rune]*WuXingValue
	Content []*WuXingValue //用于排序
}

//从输入的字符串数出五行的计数信息
func CountWuxing(data *WuXingCount, input string) (output *WuXingCount) {
	if data == nil {
		data_jin := &WuXingValue{
			WuXingChr: WuXingJinChar,
			Count:     0,
		}
		data_mu := &WuXingValue{
			WuXingChr: WuXingMuChar,
			Count:     0,
		}
		data_shui := &WuXingValue{
			WuXingChr: WuXingShuiChar,
			Count:     0,
		}
		data_huo := &WuXingValue{
			WuXingChr: WuXingHuoChar,
			Count:     0,
		}
		data_tu := &WuXingValue{
			WuXingChr: WuXingTuChar,
			Count:     0,
		}
		data_index := map[rune]*WuXingValue{}
		data_index[data_jin.WuXingChr] = data_jin
		data_index[data_mu.WuXingChr] = data_mu
		data_index[data_shui.WuXingChr] = data_shui
		data_index[data_huo.WuXingChr] = data_huo
		data_index[data_tu.WuXingChr] = data_tu

		data = &WuXingCount{
			Index:   data_index,
			Content: []*WuXingValue{data_jin, data_mu, data_shui, data_huo, data_tu},
		}
	}

	for _, input_ele := range input {
		switch input_ele {
		case WuXingJinChar:
			data.Index[WuXingJinChar].Count += 1
		case WuXingMuChar:
			data.Index[WuXingMuChar].Count += 1
		case WuXingShuiChar:
			data.Index[WuXingShuiChar].Count += 1
		case WuXingHuoChar:
			data.Index[WuXingHuoChar].Count += 1
		case WuXingTuChar:
			data.Index[WuXingTuChar].Count += 1
		}
	}

	return data
}

//排序五行的计数信息，倒序排列
func SortWuxingCount(data *WuXingCount) (output *WuXingCount) {
	if data == nil {
		panic("empty input")
	}

	sort.SliceStable(data.Content,
		func(i, j int) bool {
			return data.Content[rune(i)].Count > data.Content[rune(j)].Count
		})

	return data
}
