package threads

import (
	"context"
	"time"
)

/**
 * Created by frankieci on 2022/4/9 12:24 pm
 */

func ExecuteFuncWithContext(timeout time.Duration, fn func() error) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	return GoFuncWithContext(ctx, fn)
}

func GoFuncWithContext(ctx context.Context, fn func() error) error {
	fnCh := make(chan error, 0)
	defer func() { close(fnCh) }()

	go func() {
		if err := fn(); err != nil {
			fnCh <- err
			return
		}
		fnCh <- nil
	}()

	select {
	case err := <-fnCh:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}
