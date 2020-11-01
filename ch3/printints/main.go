// intsToString is like fmt.Sprint(values) but adds commas.
package main

import (
	"bytes"
	"fmt"
	"os"
	"reflect"
	"sort"
	"unicode"
)

func intsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')

	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}

func cleanUpWords(s string) string {
	// ns := strings.ReplaceAll(s, " ", "")
	// ns = strings.ToLower(ns)
	// return ns

	// var t []rune
	// for _, r := range s {
	// 	if unicode.IsLetter(r) {
	// 		r = unicode.ToLower(r)
	// 		t = append(t, r)
	// 	}
	// }
	// return string(t)

	r := []rune(s)
	for i := 0; i < len(r); i++ {
		if unicode.IsLetter(r[i]) {
			r[i] = unicode.ToLower(r[i])
			continue
		}
		r[i] = r[len(r)-1]
		r = r[:len(r)-1]
		i--
	}
	return string(r)
}

func isAnagrams(word1 string, word2 string) bool {
	w1, w2 := []rune(word1), []rune(word2)
	if len(w1) != len(w2) {
		return false
	}

	sort.Slice(w1, func(i int, j int) bool { return w1[i] < w1[j] })
	sort.Slice(w2, func(i int, j int) bool { return w2[i] < w2[j] })
	for i, r := range w1 {
		if r != w2[i] {
			// fmt.Printf("%q != %q\n", r, w2[i])
			return false
		}
	}
	return true
}

func isAnagrams2(word1 string, word2 string) bool {
	w1, w2 := []rune(word1), []rune(word2)
	if len(w1) != len(w2) {
		//fmt.Println("Different length")
		return false
	}

	sort.Slice(w1, func(i int, j int) bool { return w1[i] < w1[j] })
	sort.Slice(w2, func(i int, j int) bool { return w2[i] < w2[j] })
	return reflect.DeepEqual(w1, w2)
}

func isAnagrams3(word1 string, word2 string) bool {
	w1, w2 := []rune(word1), []rune(word2)
	if len(w1) != len(w2) {
		// fmt.Println("Different length")
		return false
	}

	sort.Slice(w1, func(i int, j int) bool { return w1[i] < w1[j] })
	sort.Slice(w2, func(i int, j int) bool { return w2[i] < w2[j] })

	var a1 [1000]rune
	var a2 [1000]rune

	copy(a1[:], w1)
	copy(a2[:], w2)
	return a1 == a2
}

func isAnagrams4(word1 string, word2 string) bool {
	w1, w2 := []rune(word1), []rune(word2)
	if len(w1) != len(w2) {
		// fmt.Println("Different length")
		return false
	}

	for _, r1 := range w1 {
		for i, r2 := range w2 {
			if r1 == r2 {
				w2[i] = w2[len(w2)-1]
				w2 = w2[:len(w2)-1]
				break
			}
		}
	}
	return len(w2) == 0
}

func main() {
	fmt.Println(intsToString([]int{23, 2, 45, 666, 212, 213232, 333, 321, 32323}))
	s1, s2 := os.Args[1], os.Args[2]
	s1, s2 = cleanUpWords(s1), cleanUpWords(s2)
	fmt.Printf("Is %q anagram of %q? Answer: %v\n", s1, s2, isAnagrams(s1, s2))
	fmt.Printf("Is %q anagram of %q? Answer: %v\n", s1, s2, isAnagrams2(s1, s2))
	fmt.Printf("Is %q anagram of %q? Answer: %v\n", s1, s2, isAnagrams3(s1, s2))
	fmt.Printf("Is %q anagram of %q? Answer: %v\n", s1, s2, isAnagrams4(s1, s2))
	// fmt.Println(s1 == s2)
	// s1 = "anya áéű"
	// fmt.Println(s1 == s2)
	// sb := []byte(s1)
	// sr := []rune(s1)
	// fmt.Println(sb)
	// fmt.Println(sr)

	// sort.Slice(sr, func(i int, j int) bool { return sr[i] < sr[j] })
	// fmt.Println(sr)
}
