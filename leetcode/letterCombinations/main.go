package main

func letterCombinations(digits string) []string {
	data := map[string][]string{
		"1": {},
		"2": {"a", "b", "c"},
		"3": {"d", "e", "f"},
		"4": {"g", "h", "i"},
		"5": {"j", "k", "l"},
		"6": {"m", "n", "o"},
		"7": {"p", "q", "r", "s"},
		"8": {"t", "u", "v"},
		"9": {"w", "x", "y", "z"},
	}
	join := func(arr1, arr2 []string) []string {
		ret := make([]string, 0, len(arr1)*len(arr2))
		for _, s1 := range arr1 {
			for _, s2 := range arr2 {
				ret = append(ret, s1+s2)
			}
		}
		return ret
	}
	ret := make([]string, 0, 1024)
	for i, ch := range digits {
		if i == 0 {
			ret = append(ret, data[string(ch)]...)
			continue
		}
		ret = join(ret, data[string(ch)])
	}
	return ret
}

func main() {
	ret := letterCombinations("23")
	_ = ret
}
