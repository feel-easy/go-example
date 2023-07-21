package tasks

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func (w *Worker) demo(ctx context.Context, paymentOrder string) {
	proc := demoProc{}
	if err := proc.do(ctx, paymentOrder); err != nil {
		fmt.Println(err.Error())
	}
}

type demoProc struct {
}

func (proc *demoProc) do(ctx context.Context, paymentOrder string) error {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			go proc.doSameThing(j)
			select {
			case <-ctx.Done():

			}
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
