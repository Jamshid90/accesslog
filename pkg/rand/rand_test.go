package rand

import (
	"testing"
)

func BenchmarkString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		String(7)
	}
}

func BenchmarkStringNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringNumber(7)
	}
}
