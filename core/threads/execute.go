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
	successCh := make(chan struct{}, 0)
	failedCh := make(chan error, 0)
	defer func() { close(successCh); close(failedCh) }()

	go func() {
		if err := fn(); err != nil {
			failedCh <- err
			return
		}
		successCh <- struct{}{}
	}()

	select {
	case err := <-failedCh:
		return err
	case <-successCh:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
