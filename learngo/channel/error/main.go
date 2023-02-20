package main

import (
	"errors"
	"fmt"
	"sync"
)

type calculate func(int, int) (int, error)

type results struct {
	result int
	err    error
}

func main() {
	wg := &sync.WaitGroup{}
	cals := []calculate{sum, sub}
	resultchan := make(chan results, len(cals))
	for _, v := range cals {
		wg.Add(1)
		go func(cal calculate) {
			result, err := cal(3, 2)
			resultchan <- results{result, err}
			wg.Done()
		}(v)
	}
	wg.Wait()
	close(resultchan)

	for v := range resultchan {
		if v.err != nil {
			fmt.Println(v.err)
			continue
		}
		fmt.Println(v.result)
	}
}

func sum(a int, b int) (int, error) {
	if a > b {
		return a + b, nil
	}
	return 0, errors.New("a is smaller than b")
}

func sub(a int, b int) (int, error) {
	if a > b {
		return a - b, nil
	}
	return 0, errors.New("a is smaller than b")
}
