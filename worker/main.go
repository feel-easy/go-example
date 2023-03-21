package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/feel-easy/go-example/worker/tasks"
)

func main() {
	worker := tasks.NewWorker()
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		worker.Run(context.Background())
	}()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	s := <-sigs
	worker.Shutdown()
	fmt.Printf("Caught signal: %v\n", s)
	wg.Wait()
}
