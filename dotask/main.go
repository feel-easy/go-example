package main

import (
	"fmt"
	"sync"
)

type Work struct {
	Ch chan int
	sync.Mutex
}

func main() {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 0; i < 100; i++ {
			ch <- i
		}
	}()

	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for c := range ch {
				fmt.Println(c)
			}
		}()
	}
	wg.Wait()
}
