package main

import "fmt"

type Ints []int

func (i Ints) Iterator() *Iterator {
	return &Iterator{
		data:  i,
		index: 0,
	}
}

type Iterator struct {
	data  Ints
	index int
}

func (i *Iterator) HasNext() bool {
	return i.index < len(i.data)
}

func (i *Iterator) Next() (v int) {
	v = i.data[i.index]
	i.index++
	return v
}

func main() {
	ints := Ints{1, 2, 3}
	for it := ints.Iterator(); it.HasNext(); {
		fmt.Println(it.Next())
	}
}
