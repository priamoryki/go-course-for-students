package storage

import (
	"context"
)

// Result represents the Size function result
type Result struct {
	// Total Size of File objects
	Size int64
	// Count is a count of File objects processed
	Count int64
}

type DirSizer interface {
	// Size calculate a size of given Dir, receive a ctx and the root Dir instance
	// will return Result or error if happened
	Size(ctx context.Context, d Dir) (Result, error)
}

// sizer implement the DirSizer interface
type sizer struct {
	// maxWorkersCount number of workers for asynchronous run
	// maxWorkersCount int
}

// NewSizer returns new DirSizer instance
func NewSizer() DirSizer {
	return &sizer{}
}

func (a *sizer) Size(ctx context.Context, d Dir) (Result, error) {
	result := Result{}
	dirs, files, err := d.Ls(ctx)
	if err != nil {
		return result, err
	}
	for _, file := range files {
		stat, err := file.Stat(ctx)
		if err != nil {
			return result, err
		}
		result.Size += stat
		result.Count += 1
	}
	c := make(chan Result, len(dirs))
	for _, dir := range dirs {
		go func(ctx context.Context, c chan Result, dir Dir) {
			res, _ := a.Size(ctx, dir)
			c <- res
		}(ctx, c, dir)
	}
	for i := 0; i < len(dirs); i++ {
		res := <-c
		result.Size += res.Size
		result.Count += res.Count
	}
	return result, nil
}
