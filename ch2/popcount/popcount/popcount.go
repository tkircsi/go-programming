// The package defines a function PopCount that
// returns the number of set bits, that is, bits whose value is 1
package popcount

import (
	"fmt"
	"strings"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PrintPC() {
	// for _, v := range pc { // uses only the index. Equal to for i, _ :=
	// 	fmt.Printf("%08b\n", v)
	// }
	fmt.Println(pc)
}

func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCount2(x uint64) int {
	var result int
	for i := 0; i < 8; i++ {
		result += int(pc[byte(x>>(i*8))])
	}
	return result
}

func PopCount3(x uint64) int {
	var r int
	for i := 0; i < 64; i++ {
		r += int((x >> i) & 1)
	}
	return r
}

func PopCount4(x uint64) int {
	s := fmt.Sprintf("%b", x)
	s = strings.ReplaceAll(s, "0", "")
	return len(s)
}
