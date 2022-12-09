// Code generated by microgen 0.9.0. DO NOT EDIT.

package service

import (
	"context"
	"fmt"
	service "github.com/valerylobachev/microgen/examples/generated"
	log "github.com/go-kit/kit/log"
)

// RecoveringMiddleware recovers panics from method calls, writes to provided logger and returns the error of panic as method error.
func RecoveringMiddleware(logger log.Logger) Middleware {
	return func(next service.StringService) service.StringService {
		return &recoveringMiddleware{
			logger: logger,
			next:   next,
		}
	}
}

type recoveringMiddleware struct {
	logger log.Logger
	next   service.StringService
}

func (M recoveringMiddleware) Uppercase(ctx context.Context, stringsMap map[string]string) (ans string, err error) {
	defer func() {
		if r := recover(); r != nil {
			M.logger.Log("method", "Uppercase", "message", r)
			err = fmt.Errorf("%v", r)
		}
	}()
	return M.next.Uppercase(ctx, stringsMap)
}

func (M recoveringMiddleware) Count(ctx context.Context, text string, symbol string) (count int, positions []int, err error) {
	defer func() {
		if r := recover(); r != nil {
			M.logger.Log("method", "Count", "message", r)
			err = fmt.Errorf("%v", r)
		}
	}()
	return M.next.Count(ctx, text, symbol)
}

func (M recoveringMiddleware) TestCase(ctx context.Context, comments []*service.Comment) (tree map[string]int, err error) {
	defer func() {
		if r := recover(); r != nil {
			M.logger.Log("method", "TestCase", "message", r)
			err = fmt.Errorf("%v", r)
		}
	}()
	return M.next.TestCase(ctx, comments)
}

func (M recoveringMiddleware) DummyMethod(ctx context.Context) (err error) {
	defer func() {
		if r := recover(); r != nil {
			M.logger.Log("method", "DummyMethod", "message", r)
			err = fmt.Errorf("%v", r)
		}
	}()
	return M.next.DummyMethod(ctx)
}

func (M recoveringMiddleware) IgnoredMethod() {
	M.next.IgnoredMethod()
}

func (M recoveringMiddleware) IgnoredErrorMethod() error {
	return M.next.IgnoredErrorMethod()
}
