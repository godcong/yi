package qigua

import (
	"math/rand"

	"yi/core"
	"yi/internal/gua"
)

func DivineByDayan(seed int64) (*core.ZhouYi, [6]core.DayanResult) {
	rng := rand.New(rand.NewSource(seed))
	var yaoValues [6]int
	var results [6]core.DayanResult
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

	return gua.DivineByNumber(shangBits, xiaBits, bianYao...), results
}

func dayanSimOneYao(rng *rand.Rand) core.DayanResult {
	const totalInitial = 49
	remaining := totalInitial
	var steps [3]core.DayanStep

	for s := 0; s < 3; s++ {
		left, right := dayanDivide(rng, remaining)
		right -= 1
		if right < 0 {
			right = 0
		}
		leftRem := 1 + dayanYuQi(left-1)
		if leftRem > left {
			leftRem = left
		}
		rightRem := 1 + dayanYuQi(right-1)
		if rightRem > right {
			rightRem = right
		}
		remainder := 1 + leftRem + rightRem
		steps[s] = core.DayanStep{
			Total:     remaining,
			Left:      left,
			Right:     right,
			Remainder: remainder,
			Final:     remaining - remainder,
		}
		remaining = steps[s].Final
	}

	var yaoValue int
	var isChanging bool
	switch remaining {
	case 36:
		yaoValue, isChanging = 9, true
	case 40:
		yaoValue, isChanging = 8, false
	case 44:
		yaoValue, isChanging = 7, false
	case 37:
		yaoValue, isChanging = 6, true
	default:
		if remaining > 40 {
			yaoValue, isChanging = 8, false
		} else {
			yaoValue, isChanging = 7, false
		}
	}

	return core.DayanResult{Steps: steps, Remaining: remaining, YaoValue: yaoValue, IsChanging: isChanging}
}

func dayanDivide(rng *rand.Rand, total int) (left, right int) {
	minLeft := 1
	maxLeft := total - 1
	left = minLeft + rng.Intn(maxLeft-minLeft+1)
	right = total - left
	return
}

func dayanYuQi(count int) int {
	return count % 4
}
