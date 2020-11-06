package main

import (
	"fmt"
	"sort"
)

func main() {

	ages := map[string]int{
		"bogi":   40,
		"boldi":  10,
		"tibcsi": 45,
	}

	ages["bogi"]++
	fmt.Println(ages)

	ages2 := make(map[string]int)
	ages2["bogi"] = 40
	ages2["boldi"] = 10

	ages3 := make(map[string]int)
	ages3["bogi"] = 40
	ages3["boldi"] = 10

	for k, v := range ages {
		fmt.Printf("%s\t%d\n", k, v)
	}
	fmt.Println()

	names := sortMap(ages)
	for _, n := range names {
		fmt.Printf("%s\t%d\n", n, ages[n])
	}

	if _, ok := ages["boci"]; !ok {
		ages["boci"] = 5
	}
	fmt.Println(ages)

	fmt.Println(equal(ages3, ages2))

}

func sortMap(m map[string]int) []string {
	// sorting a map
	// var names []string
	names := make([]string, 0, len(m))

	for k := range m {
		names = append(names, k)
	}

	sort.Strings(names)

	return names
}

func equal(m1, m2 map[string]int) bool {
	if len(m1) != len(m2) {
		return false
	}

	for k, v := range m1 {
		if m2v, ok := m2[k]; !ok || m2v != v {
			return false
		}
	}
	return true
}
