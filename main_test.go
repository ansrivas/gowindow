package main

import (
	"testing"
)

func Benchmark_processFile(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		handleProcessing("test_data.txt", false)
	}
}
