package util

import (
	"errors"
	"time"
)

var RetryableError = errors.New("retryable error")

func Retry(maxRetry int, f func() error) error {
	for i := 0; i < maxRetry; i++ {
		time.Sleep(time.Duration(i*i) * time.Second)
		if err := f(); err != nil {
			if errors.Is(err, RetryableError) {
				continue
			}
			return err
		}
		return nil
	}
	return nil
}
