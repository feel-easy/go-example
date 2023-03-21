package tasks

import (
	"context"
	"time"
)

const (
	restTime = 10 * time.Second //time.Minute
	timeout  = 5 * time.Second
)

type Worker struct {
	close chan struct{}
}

func NewWorker() *Worker {
	return &Worker{
		close: make(chan struct{}),
	}
}

func (worker *Worker) Run(ctx context.Context) {
	ticker := time.Tick(restTime)
loop:
	for {
		select {
		case <-worker.close:
			break loop
		case <-ticker:
		}
		c, cancel := context.WithTimeout(ctx, timeout)
		defer cancel()
		worker.demo(c)
	}
}

func (worker *Worker) Shutdown() {
	close(worker.close)
}
