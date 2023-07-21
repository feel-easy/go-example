package main

import (
	"context"
	"fmt"
	"time"
)

func A(ctx context.Context) {
	time.Sleep(1 * time.Second)
	fmt.Println("AAAA")
}

func B(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	go func() {
		time.Sleep(6 * time.Second)
		fmt.Println("BBBB")
	}()
	<-ctx.Done()
	return nil
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	for i := 0; i < 5; i++ {
		go A(ctx)
		go func() error {
			return B(ctx)
		}()
	}
	<-ctx.Done()
}
