package executor

import (
	"context"
)

type (
	In  <-chan any
	Out = In
)

type Stage func(in In) (out Out)

func chanToChanWithTimeout(ctx context.Context, in In, out chan any) {
	defer close(out)
	for {
		select {
		case <-ctx.Done():
			return
		case value, ok := <-in:
			if !ok {
				return
			}
			out <- value
		}
	}
}

func ExecutePipeline(ctx context.Context, in In, stages ...Stage) Out {
	for _, stage := range stages {
		inWithTimeout := make(chan any)
		go chanToChanWithTimeout(ctx, in, inWithTimeout)
		in = stage(inWithTimeout)
	}
	out := make(chan any)
	go chanToChanWithTimeout(ctx, in, out)
	return out
}
