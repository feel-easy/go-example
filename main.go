package main

import (
	"fmt"

	"google.golang.org/protobuf/types/known/wrapperspb"
)

type Item struct {
	Name    string
	List    []int
	B       bool
	Enabled *wrapperspb.BoolValue
}

type item Item

type Aa string

func (a Aa) ToItem() *Item {
	return &Item{
		Name: string(a),
	}
}

func demo() bool {
	fmt.Println("aaaa")
	return true
}

func main() {
	a := "17611111111"
	// fmt.Println(a[len(a)-4:])
	// if ok := demo(); false && ok {
	// 	fmt.Println("bbbb")
	// }
	for _, i := range a {
		j := string(i)
		flag := false
		switch j {
		case "6":
			flag = true

		}
		fmt.Println(j)
		if flag {
			break
		}

	}
}

// func main() {
// 	fmt.Println(path.Ext("aaa.zip"))
// 	b := 0.8
// 	c := int(1 << 30 * b)
// 	d := float64(c) / (1 << 30)
// 	fmt.Println(b, c, d, strconv.FormatFloat((float64(858993459)/(1<<30)), 'f', 2, 64))
// 	aa := []string{"aa", "bb", "cc"}
// 	for i, j := range aa {
// 		fmt.Println(i, j)
// 	}
// 	bb := map[string]string{"a": "1", "b": "2"}
// 	fmt.Println(bb["c"])
// 	cc := Item{Name: "zhangsan"}
// 	dd := cc.Name
// 	cc.Name = ""

// 	fmt.Println(dd)
// 	fmt.Printf("%q", time.Now().UTC().Unix())

// 	fmt.Println(cc.Enabled.GetValue())
// 	ddd := []int{1, 2, 3}
// 	fmt.Println(ddd[len(ddd)-1:])
// 	fmt.Println(ddd[:len(ddd)-1])
// 	fmt.Println(ddd[len(ddd)-1])
// 	ee := Aa("ssss")
// 	fmt.Printf("%q", ee.ToItem())

// 	ff := &item{
// 		Name: "name",
// 	}
// 	fmt.Printf("%q", Item(*ff))
// 	gg := make(map[string]bool)
// 	fmt.Println(gg["bb"])
// }
