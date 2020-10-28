package popcount

import "testing"

func BenchmarkPopCount(t *testing.B) {
	for n := 0; n < t.N; n++ {
		PopCount(uint64(n))
	}
}

func BenchmarkPopCount2(t *testing.B) {
	for n := 0; n < t.N; n++ {
		PopCount2(uint64(n))
	}
}

func BenchmarkPopCount3(t *testing.B) {
	for n := 0; n < t.N; n++ {
		PopCount3(uint64(n))
	}
}

func BenchmarkPopCount4(t *testing.B) {
	for n := 0; n < t.N; n++ {
		PopCount4(uint64(n))
	}
}
