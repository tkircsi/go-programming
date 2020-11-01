// This package present some type conversion
package main

import (
	"fmt"
	"strconv"
)

func main() {
	// number to string
	i, pi, neg := 1974, 3.14, -20
	var s string
	s = fmt.Sprintf("%d", i)
	fmt.Println(s)
	s = fmt.Sprintf("%f", pi)
	fmt.Println(s)
	s = fmt.Sprintf("%d", neg)
	fmt.Println(s)

	s = strconv.Itoa(i)
	fmt.Println(s)
	s = strconv.FormatFloat(pi, 'f', 3, 64)
	fmt.Println(s)

	// string to number
	i, _ = strconv.Atoi("1976")
	fmt.Printf("Type: %T,\t%[1]v\n", i)
	pi, _ = strconv.ParseFloat("3.14", 64)
	fmt.Printf("Type: %T,\t%[1]v\n", pi)

}
