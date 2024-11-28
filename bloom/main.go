package main

import (
	"fmt"

	"github.com/bits-and-blooms/bloom/v3"
)

func main() {
	// 创建布隆过滤器，参数为预计元素数量和误报率
	expectedElements := uint(1000) // 预计存储的元素数量
	falsePositiveRate := 0.01      // 误报率（1%）
	filter := bloom.NewWithEstimates(expectedElements, falsePositiveRate)

	// 添加一些元素
	elements := []string{"apple", "banana", "cherry", "date", "fig"}
	for _, elem := range elements {
		filter.Add([]byte(elem))
	}

	// 测试是否可能存在
	testItems := []string{"apple", "grape", "banana", "kiwi", "fig"}
	for _, item := range testItems {
		if filter.Test([]byte(item)) {
			fmt.Printf("%s might be in the set.\n", item)
		} else {
			fmt.Printf("%s is definitely not in the set.\n", item)
		}
	}

	// 检查是否添加过元素
	item := "grape"
	if filter.TestAndAdd([]byte(item)) {
		fmt.Printf("%s was already in the set.\n", item)
	} else {
		fmt.Printf("%s was not in the set, but now added.\n", item)
	}
}
