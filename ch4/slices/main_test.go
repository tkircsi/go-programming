package main

import "testing"

func BenchmarkEqual1(t *testing.B) {
	s1 := makeMonths()
	s2 := makeMonths()
	for n := 0; n < t.N; n++ {
		equal1(s1, s2)
	}
}

func BenchmarkEqual2(t *testing.B) {
	s1 := makeMonths()
	s2 := makeMonths()
	for n := 0; n < t.N; n++ {
		equal2(s1, s2)
	}
}
