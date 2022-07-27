package analyzer

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
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
	anl := New([]int{2, 4, 8}, 1, "Test")
	err := anl.Analyze(
		func(i int) error {
			func(threads int) {
				for j := 1; j < threads; j++ {
					num := j
					go func() { fmt.Println(num) }()
				}
			}(i)
			return nil
		},
	)

	if err != nil {
		t.Errorf("Something went wrong: %v", err)
	}
}

func TestAnalyzeEmptyName(t *testing.T) {
	anl := New([]int{2, 4, 8}, 1)
	err := anl.Analyze(
		func(i int) error {
			func(threads int) {
				for j := 1; j < threads; j++ {
					num := j
					go func() { fmt.Println(num) }()
				}
			}(i)
			return nil
		},
	)

	if err != nil {
		t.Errorf("Something went wrong: %v", err)
	}
}

func TestAnalyzeError(t *testing.T) {
	anl := New([]int{2, 4, 8}, 1)
	err := anl.Analyze(
		func(i int) error {
			return func(threads int) error {
				for j := 1; j < threads; j++ {
					num := j
					go func() { fmt.Println(num) }()
				}
				return errors.New("some error")
			}(i)
		},
	)

	if err == nil {
		t.Errorf("Error expected")
	}
}

func TestSaveXLSX(t *testing.T) {
	table := []struct {
		name        string
		dstFilePath string
		isError     bool
	}{
		{
			name:        "analyzer.SaveXLSX(): Wrong destination file",
			dstFilePath: "",
			isError:     true,
		},
		{
			name:        "analyzer.SaveXLSX(): Check result",
			dstFilePath: "../../results/test_result.xlsx",
			isError:     false,
		},
	}

	for _, test := range table {
		t.Run(test.name, func(t *testing.T) {
			anl := New([]int{2, 4, 8}, 1)
			err := anl.Analyze(
				func(i int) error {
					func(threads int) {
						for j := 1; j < threads; j++ {
							num := j
							go func() { fmt.Println(num) }()
						}
					}(i)
					return nil
				},
			)
			if err != nil {
				t.Errorf("something went wrong: %v", err)
			}

			err = anl.SaveXLSX(test.dstFilePath)

			if (err != nil) != test.isError {
				t.Errorf("error received = %v, expected = %v", err, test.isError)
			}
			if err == nil && !test.isError {
				assert.FileExists(t, test.dstFilePath)
			}
		})
	}
}
