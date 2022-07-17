package analyzer

import (
	"errors"
	"fmt"
	"testing"
)

func TestAnalyzeOne(t *testing.T) {
	_, err := analyzeOne(
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

func TestAnalyzeOneError(t *testing.T) {
	_, err := analyzeOne(
		func(i int) error {
			return func(threads int) error {
				for j := 1; j < threads; j++ {
					num := j
					go func() { fmt.Println(num) }()
				}
				return errors.New("some error")
			}(i)
		},
		2,
	)

	if err == nil {
		t.Errorf("Error expected")
	}
}

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
		"Test",
		[]int{2, 4, 8},
		1,
	)

	if err != nil {
		t.Errorf("Something went wrong: %v", err)
	}
}

func TestAnalyzeError(t *testing.T) {
	_, err := Analyze(
		func(i int) error {
			return func(threads int) error {
				for j := 1; j < threads; j++ {
					num := j
					go func() { fmt.Println(num) }()
				}
				return errors.New("some error")
			}(i)
		},
		"Test",
		[]int{2, 4, 8},
		1,
	)

	if err == nil {
		t.Errorf("Error expected")
	}
}
