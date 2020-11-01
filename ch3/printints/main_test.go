package main

import (
	"os"
	"testing"
)

func BenchmarkCleanUpWords1(t *testing.B) {
	w1 := os.Args[len(os.Args)-2]
	// w2 := os.Args[len(os.Args)-1]
	for n := 0; n < t.N; n++ {
		cleanUpWords(w1)
	}
}

func BenchmarkCleanUpWords2(t *testing.B) {
	// w1 := os.Args[len(os.Args)-2]
	w2 := os.Args[len(os.Args)-1]
	for n := 0; n < t.N; n++ {
		cleanUpWords(w2)
	}
}

func BenchmarkIsAnagrams(t *testing.B) {
	w1 := os.Args[len(os.Args)-2]
	w2 := os.Args[len(os.Args)-1]
	w1, w2 = cleanUpWords(w1), cleanUpWords(w2)
	t.ResetTimer()
	for n := 0; n < t.N; n++ {
		isAnagrams(w1, w2)
	}
}

func BenchmarkIsAnagrams2(t *testing.B) {
	w1 := os.Args[len(os.Args)-2]
	w2 := os.Args[len(os.Args)-1]
	w1, w2 = cleanUpWords(w1), cleanUpWords(w2)
	t.ResetTimer()
	for n := 0; n < t.N; n++ {
		isAnagrams2(w1, w2)
	}
}

func BenchmarkIsAnagrams3(t *testing.B) {
	w1 := os.Args[len(os.Args)-2]
	w2 := os.Args[len(os.Args)-1]
	w1, w2 = cleanUpWords(w1), cleanUpWords(w2)
	t.ResetTimer()
	for n := 0; n < t.N; n++ {
		isAnagrams3(w1, w2)
	}
}

func BenchmarkIsAnagrams4(t *testing.B) {
	w1 := os.Args[len(os.Args)-2]
	w2 := os.Args[len(os.Args)-1]
	w1, w2 = cleanUpWords(w1), cleanUpWords(w2)
	t.ResetTimer()
	for n := 0; n < t.N; n++ {
		isAnagrams4(w1, w2)
	}
}
