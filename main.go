package main

import (
	"encoding/json"
	"fmt"

	"google.golang.org/protobuf/types/known/wrapperspb"
)

type Body struct {
	ItemID string `json:"item_id"`
	Filter struct {
		Sort struct {
			CommentSmartHot string `json:"comment_smart_hot"`
		} `json:"sort"`
	} `json:"filter"`
}
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
	switch num := 1; num {
	case 1:
		fmt.Println("Case 1")
		fallthrough
	case 2:
		fmt.Println("Case 2")
		fallthrough
	case 3:
		fmt.Println("Case 3")
	default:
		fmt.Println("Default case")
	}
	aa := map[string][]string{"aa": {"aa"}}
	fmt.Printf("%v", aa["a1a"])
	fmt.Print(FormatMobileStar("13933730122"))
	ret, _ := json.Marshal(map[string]any{
		"item_id": 121,
		"filter": map[string]any{
			"sort": map[string]string{
				"comment_smart_hot": "2",
			},
		},
	})

	fmt.Printf("%v", string(ret))
}

func FormatMobileStar(mobile string) string {
	if len(mobile) <= 10 {
		return mobile
	}

	return fmt.Sprintf("%s****%s", mobile[:3], mobile[7:])
}

// func main() {
// 	a := "17611111111"
// 	// fmt.Println(a[len(a)-4:])
// 	// if ok := demo(); false && ok {
// 	// 	fmt.Println("bbbb")
// 	// }
// 	var sts []string
// 	for _, i := range a {
// 		j := string(i)
// 		sts = append(sts, j)
// 		flag := false
// 		switch j {
// 		case "6":
// 			flag = true

// 		}
// 		fmt.Println(j)
// 		if flag {
// 			break
// 		}
// 	}
// 	fmt.Println(sts)
// }

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
