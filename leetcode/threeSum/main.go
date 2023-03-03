package main

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	numMaps := make(map[int][]int, len(nums))
	for i, num := range nums {
		numMaps[num] = append(numMaps[num], i)
	}
	keys := make([]int, 0, len(numMaps))
	for k := range numMaps {
		keys = append(keys, k)
	}
	sort.IntSlice(keys).Sort()

	var left, right int
	for i, num := range keys {
		if num >= 0 {
			left = i
			right = i
			break
		}
	}
	ret := make([][]int, 0, len(keys))
	for i := left; i >= 0; i-- {
		num1 := keys[i]
		if num1 == 0 {
			if len(numMaps[num1]) >= 3 {
				ret = append(ret, []int{0, 0, 0})
			}
			continue
		}
		for j := right; j < len(keys); j++ {
			num3 := keys[j]
			num2 := -(num1 + num3)
			if num2 > num3 {
				continue
			}
			if num2 < num1 {
				break
			}
			switch num2 {
			case num1:
				if len(numMaps[num1]) >= 2 {
					ret = append(ret, []int{num1, num2, num3})
					break
				}
			case num3:
				if len(numMaps[num3]) >= 2 {
					ret = append(ret, []int{num1, num2, num3})
					break
				}
			default:
				if _, ok := numMaps[num2]; ok {
					ret = append(ret, []int{num1, num2, num3})
					break
				}
			}
		}
	}
	return ret
}
