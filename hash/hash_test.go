package hash

import (
	"testing"
)

func BenchmarkUint128(b *testing.B) {
	u := Uint128([2]uint64{151515, 15541561})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Uint128Base62.Encoding(u)
	}
}

func BenchmarkUint64(b *testing.B) {
	u := uint64(151515)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Uint64Base62.Encoding(u)
	}
}

func BenchmarkUint32(b *testing.B) {
	u := uint32(151515)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Uint32Base62.Encoding(u)
	}
}
