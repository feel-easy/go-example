package main

import (
	"fmt"
	"reflect"
)

type Entity struct {
	Name  string `tag:"name"`
	Age   int    `tag:"age"`
	Title string
}

func (ce Entity) toMap(category string) map[string]interface{} {
	obj1 := reflect.TypeOf(ce)
	obj2 := reflect.ValueOf(ce)
	data := make(map[string]interface{})
	for i := 0; i < obj1.NumField(); i++ {
		tag := obj1.Field(i).Tag.Get("tag")
		if tag == "" || tag == category {
			data[obj1.Field(i).Name] = obj2.Field(i).Interface()
		}
	}
	return data
}
func main() {
	aa := Entity{Name: "zhangsan", Age: 10, Title: "heihei"}
	fmt.Println(aa.toMap("name"))
	fmt.Println(aa.toMap("age"))
}
