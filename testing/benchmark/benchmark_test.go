package benchmark

import (
	"testing"
)

// to use a benchmark the name should start with Benchmark
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 4)
	}
}
