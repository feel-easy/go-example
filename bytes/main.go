package main

import (
	"fmt"

	"github.com/feel-easy/go-example/bytes/bytefmt"
)

func main() {
	// fmt.Println(bytefmt.ByteSize(100.5 * bytefmt.MEGABYTE)) // "100.5M"
	// fmt.Println(int64(1<<30) * 0.8) // "100.5M"
	// fmt.Println(bytefmt.ByteSize(uint64(858993459))) // "1K"
	fmt.Println(bytefmt.ToBytes("1GB"))
	a := make([]int, 0)
	a = append(a, 1)
	fmt.Println(a)
	val, _ := bytefmt.ToBytes(fmt.Sprintf("%fGB", 0.8))
	fmt.Println(bytefmt.ByteSize(val))
	// bytefmt.ToBytes(fmt.Sprintf("%fGB", 0.8))
}
