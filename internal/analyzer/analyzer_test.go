package analyzer

import (
	"fmt"
	"testing"
)

func TestAnalyze(t *testing.T) {
	_, err := Analyze(
		func(i int) error {
			func(threads int) {
				for j := 1; j < threads; j++ {
					num := j
					go func() { fmt.Println(num) }()
				}
			}(i)
			return nil
		},
		2,
	)

	if err != nil {
		t.Errorf("Something went wrong: %v", err)
	}
}
