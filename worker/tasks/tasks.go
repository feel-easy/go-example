package tasks

import (
	"context"
	"sync"
	"time"
)

const (
	restTime    = 10 * time.Second //time.Minute
	timeout     = 5 * time.Second
	paymentTime = 10 * time.Minute
)

var worker *Worker

type Worker struct {
	close chan struct{}
	tasks chan Task
	wg    sync.WaitGroup
}

func InitWorker() {
	worker = NewWorker()
}

func NewWorker() *Worker {
	return &Worker{
		tasks: make(chan Task, 100),
		close: make(chan struct{}),
		wg:    sync.WaitGroup{},
	}
}

func AddTask(task Task) {
	worker.tasks <- task
}

type Task interface {
	Do(context.Context) error
}

func Run(ctx context.Context) {
	// ticker := time.Tick(restTime)
loop:
	for {
		select {
		case <-worker.close:
			close(worker.tasks)
			worker.wg.Wait()
			break loop
		case task := <-worker.tasks:
			worker.wg.Add(1)
			go func() {
				defer worker.wg.Done()
				c, cancel := context.WithTimeout(ctx, paymentTime)
				defer cancel()
				if err := task.Do(c); err != nil {

				}
			}()
		}
	}
}

func Shutdown() {
	close(worker.close)
}
