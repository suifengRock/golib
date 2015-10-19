package utils

import (
	"testing"
)

func BenchmarkGetCurGoroutineId(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetCurGoroutineId()
	}
}
