run:
	go run cmd/main.go

test:
	go test internal/algorithms/fourier/fourier_test.go internal/algorithms/fourier/fourier.go internal/algorithms/fourier/fourier_parallel.go 
	go test pkg/analyzer/analyzer_test.go pkg/analyzer/analyzer.go 