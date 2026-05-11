package yi

import (
	"testing"
)

func BenchmarkGetGuaByXu(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for xu := 1; xu <= 64; xu++ {
			GetGuaByXu(xu)
		}
	}
}

func BenchmarkGetGuaByIndex(b *testing.B) {
	indices := []string{"乾乾", "坤坤", "坎坎", "离离", "震震", "艮艮", "巽巽", "兑兑"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, idx := range indices {
			GetGuaByIndex(idx)
		}
	}
}

func BenchmarkDivine(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Divine(0, 0)
	}
}

func BenchmarkDivineByNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DivineByNumber(0, 0, 3)
	}
}