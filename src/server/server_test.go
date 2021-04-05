package main

import (
	"context"
	"testing"
	"time"
)

func TestSetver(t *testing.T) {
	serve := func(ctx context.Context) {
		go func() {
			newServer()
			for {
				select {
				case <-ctx.Done():
					return // returning not to leak the goroutine
				}
			}
		}()
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	serve(ctx)

	sleepDuration := 10
	timeout := time.Duration(sleepDuration) * time.Second
	time.Sleep(timeout)

	t.Errorf("Exit. Test duration %d seconds", sleepDuration)
}
