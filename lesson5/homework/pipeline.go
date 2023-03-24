package executor

import (
	"context"
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
		in = stage(inWithTimeout)
	}
	return in
}
