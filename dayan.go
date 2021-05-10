package yi

import (
	"math/bits"
	"strconv"
	"strings"
)

// DaYan ...
type DaYan struct {
	Number  int
	Lucky   string
	NvMing  string
	Max     bool
	SkyNine string
	Comment string
}

var daYanList map[int]*DaYan

func init() {
	daYanList = make(map[int]*DaYan)

	file81shu, err := DataFiles.Open("data/81shu.csv")
	if err != nil {
		panic(err)
	}

	records, err := readData(file81shu)

	if err != nil {
		panic(err)
	}

	for _, record := range records {
		bihua, err := strconv.ParseInt(record[0], 10, bits.UintSize)
		if err != nil {
			panic(err)
		}

		var max bool = false
		maxstr := strings.TrimSpace(record[3])
		if maxstr == "最吉" {
			max = true
		}

		dayan := DaYan{
			Number:  int(bihua),
			Lucky:   record[1],
			NvMing:  record[2],
			Max:     max,
			SkyNine: record[5],
			Comment: record[6],
		}

		daYanList[dayan.Number-1] = &dayan
	}
}

// IsNotSuitableGirl 女性不宜此数
func (dy DaYan) IsNotSuitableGirl() bool {
	return dy.NvMing == "凶"
}

//IsMax 是否最大好运数
func (dy DaYan) IsMax() bool {
	return dy.Max
}

//GetDaYan 获取大衍之数
func GetDaYan(idx int) DaYan {
	if idx <= 0 {
		panic("wrong idx")
	}
	i := (idx - 1) % 81

	return *daYanList[i]
}
