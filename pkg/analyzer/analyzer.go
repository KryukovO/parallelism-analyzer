package analyzer

import (
	"fmt"
	"time"
)

func Analyze(f func(int) error, threadCount int) (elapsed time.Duration, errF error) {
	defer func() {
		if msg := recover(); msg != nil {
			elapsed = 0
			errF = fmt.Errorf("%v", msg)
		}
	}()

	start := time.Now()
	err := f(threadCount)
	if err != nil {
		return 0, err
	}
	now := time.Now()
	elapsed = now.Sub(start)

	return elapsed, nil
}
