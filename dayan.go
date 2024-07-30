package yi

import (
	"fmt"
	"math/bits"
	"strconv"
	"strings"
)

type DaYanIndex uint

// DaYan ...
type DaYan struct {
	Index   DaYanIndex
	Lucky   string
	nvMing  string
	max     bool
	SkyNine string
	Comment string
}

func loadDaYan() ([DaYanMax]DaYan, error) {
	var dayanData [DaYanMax]DaYan

	file81shu, err := DataFiles.Open("data/81shu.csv")
	if err != nil {
		return dayanData, err
	}

	records, err := readData(file81shu)
	if err != nil {
		return dayanData, err
	}

	for _, record := range records {
		bihua, _ := strconv.ParseUint(record[0], 10, bits.UintSize)
		idx := DaYanIndex(bihua)
		if !idx.IsValid() {
			return dayanData, fmt.Errorf("dayan number out of range: %d", idx)
		}

		dayanData[bihua-1] = DaYan{
			Index:   idx,
			Lucky:   record[1],
			nvMing:  record[2],
			max:     strings.TrimSpace(record[3]) == "最吉",
			SkyNine: record[5],
			Comment: record[6],
		}
	}
	return dayanData, nil
}

// IsNvMing 女性不宜此数
func (dy DaYan) IsNvMing() bool {
	return dy.nvMing == "凶"
}

// IsMax 是否最大好运数
func (dy DaYan) IsMax() bool {
	return dy.max
}

// DaYanByIndex 获取大衍之数
func DaYanByIndex(idx DaYanIndex) DaYan {
	if !idx.IsValid() {
		panic(fmt.Errorf("dayan number out of range: %d", idx))
	}
	return dayanData[idx.index()]
}

func (idx DaYanIndex) IsValid() bool {
	return idx >= 1 && idx.index() < DaYanMax
}

func (idx DaYanIndex) DaYan() DaYan {
	return DaYanByIndex(idx)
}

func (idx DaYanIndex) index() uint {
	return uint((idx - 1) % 81)
}
