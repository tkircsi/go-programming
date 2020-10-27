package main

import (
	"testing"
)

var result string

func BenchmarkEcho1(b *testing.B) {
	var s string
	for n := 0; n < b.N; n++ {
		s = echo1()
	}
	result = s
}

func BenchmarkEcho1_2(b *testing.B) {
	var s string
	for n := 0; n < b.N; n++ {
		s = echo1_2()
	}
	result = s
}

func BenchmarkEcho2(b *testing.B) {
	var s string
	for n := 0; n < b.N; n++ {
		s = echo2()
	}
	result = s
}

func BenchmarkEcho3(b *testing.B) {
	var s string
	for n := 0; n < b.N; n++ {
		s = echo3()
	}
	result = s
}
