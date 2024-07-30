package yi

import (
	"sync"
)

var (
	once         sync.Once
	guaData      [EightMax]Gua
	dayanData    [DaYanMax]DaYan
	guaxiangData map[string]GuaXiang
)

func init() {
	var err error
	once.Do(func() {
		guaData, err = loadGua()
		if err != nil {
			panic(err)
		}
		dayanData, err = loadDaYan()
		if err != nil {
			panic(err)
		}
		guaxiangData, err = loadGuaXiang()
		if err != nil {
			panic(err)
		}
	})
}

func loadGua() ([EightMax]Gua, error) {
	// TODO: need implement
	return [EightMax]Gua{}, nil
}
