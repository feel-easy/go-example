package tasks

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func (w *Worker) demo(ctx context.Context) {
	proc := demoProc{}
	if err := proc.do(); err != nil {
		fmt.Println(err.Error())
	}
}

type demoProc struct {
}

func (proc *demoProc) do() error {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			proc.doSameThing(j)
		}(i)
	}
	wg.Wait()
	return nil
}

func (proc *demoProc) doSameThing(i int) error {
	fmt.Println("old", i)
	time.Sleep(10 * time.Second)
	fmt.Println("new", i)
	return nil
}
