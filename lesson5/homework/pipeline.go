package executor

import (
	"context"
	"sync"
)

type (
	In  <-chan any
	Out = In
)

type Stage func(in In) (out Out)

func ExecutePipeline(ctx context.Context, in In, stages ...Stage) Out {
	for _, stage := range stages {
		inWithTimeout := make(chan any)
		go func(in In) {
			defer close(inWithTimeout)
			for {
				select {
				case <-ctx.Done():
					return
				case value, ok := <-in:
					if !ok {
						return
					}
					inWithTimeout <- value
				}
			}
		}(in)
		wg := sync.WaitGroup{}
		wg.Add(1)
		go func() {
			defer wg.Done()
			in = stage(inWithTimeout)
		}()
		wg.Wait()
	}
	return in
}
