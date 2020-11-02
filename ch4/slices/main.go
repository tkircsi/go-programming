// This package is about practice Go's sliece
package main

import "fmt"

func main() {
	m := makeMonths()

	fmt.Println("=== Months ===")
	fmt.Printf("Len: %d, Cap: %d\n", len(m), cap(m))
	fmt.Printf("Data: %v\n\n", m)

	fmt.Println("=== Summer ===")
	summer := m[6:9]
	fmt.Printf("Len: %d, Cap: %d\n", len(summer), cap(summer))
	fmt.Printf("Data: %v\n\n", summer)

	fmt.Println("=== Q2 ===")
	q2 := m[4:7]
	fmt.Printf("Len: %d, Cap: %d\n", len(q2), cap(q2))
	fmt.Printf("Data: %v\n\n", q2)

	fmt.Println("=== Share \"Június\" ===")
	summer[0] = "*Június*"
	fmt.Printf("Months: %v\n", m)
	fmt.Printf("Summer: %v\n", summer)
	fmt.Printf("Q2: %v\n\n", q2)

	fmt.Println("=== Common in Q2 and Summer ===")
	res := findCommonElements(q2, summer)
	for _, item := range res {
		fmt.Printf("%q is in q2 and summe\n", item)
	}
	fmt.Println()

	fmt.Println("=== Reverse months ===")
	reverse(m)
	fmt.Printf("Data: %v\n\n", m)

	fmt.Println("=== Rotate months ===")
	m = makeMonths()
	rotate(m[1:], 5)
	fmt.Printf("Data: %v\n\n", m)

	fmt.Println("=== Slice equality ===")
	s1 := makeMonths()
	s2 := makeMonths()
	// s2[1] = "Januar"
	fmt.Println(equal1(s1, s2))
	fmt.Println(equal2(s1, s2))
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println()

	fmt.Println("=== Append []int ===")
	i := []int{1, 2, 3, 4}
	fmt.Printf("len: %d, cap: %d, data:%v\n", len(i), cap(i), i)
	j := appendInt(i, 10)
	fmt.Printf("len: %d, cap: %d, data:%v\n\n", len(j), cap(j), j)

	fmt.Println("=== Nonempty ===")
	s := []string{"One", "", "Two", "Three", "", "", "Four", ""}
	fmt.Printf("%q\n", s)
	s = nonempty(s)
	fmt.Printf("%q\n\n", s)

	fmt.Println("=== Reverse array pointer ===")
	p := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("%v\n", p)
	reversePtr(&p)
	fmt.Printf("%v\n\n", p)

}

func makeMonths() []string {
	var months = []string{
		0:  "",
		1:  "Január",
		2:  "Február",
		3:  "Március",
		4:  "Április",
		5:  "Május",
		6:  "Június",
		7:  "Július",
		8:  "Augusztus",
		9:  "Szeptember",
		10: "Október",
		11: "November",
		12: "December",
	}

	return months
}

func findCommonElements(s, d []string) []string {
	var common []string
	for _, si := range s {
		for _, di := range d {
			if si == di {
				common = append(common, di)
			}
		}
	}
	return common
}

func reverse(s []string) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func reversePtr(p *[5]int) {
	for i, j := 0, len(*p)-1; i < j; i, j = i+1, j-1 {
		p[i], p[j] = p[j], p[i]
	}
}

func rotate(s []string, n int) {
	if n > len(s)-1 {
		reverse(s)
		return
	}
	reverse(s[:n])
	reverse(s[n:])
	reverse(s)
}

func equal1(s1, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

func equal2(s1, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i, j := 0, len(s1)/2+1; i < len(s1)/2; i, j = i+1, j+1 {
		if s1[i] != s2[i] || s1[j] != s2[j] {
			return false
		}
	}
	return true
}

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen < cap(x) {
		z = x[:zlen]
	} else {
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	z[len(x)] = y
	return z
}

// nonpempty filter empty strings in-place
func nonempty(s []string) []string {
	i := 0
	for _, item := range s {
		if item != "" {
			s[i] = item
			i++
		}
	}
	return s[:i]
}
