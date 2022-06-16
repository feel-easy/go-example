package main

import (
	"fmt"
	"path"
	"strconv"

	"google.golang.org/protobuf/types/known/wrapperspb"
)

type Item struct {
	List    []int
	B       bool
	Enabled *wrapperspb.BoolValue
}

func main() {
	fmt.Println(path.Ext("aaa.zip"))
	b := 0.8
	c := int(1 << 30 * b)
	d := float64(c) / (1 << 30)
	fmt.Println(b, c, d, strconv.FormatFloat((float64(858993459)/(1<<30)), 'f', 2, 64))
	aa := []string{"aa", "bb", "cc"}
	for i, j := range aa {
		fmt.Println(i, j)
	}
	bb := map[string]string{"a": "1", "b": "2"}
	fmt.Println(bb["c"])
	cc := Item{}

	fmt.Println(cc.Enabled.GetValue())
	fmt.Printf("%q", bb)
}
