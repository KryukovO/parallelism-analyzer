run:
	go run cmd/main.go

tests:
	go test internal/algorithms/fourier/fourier_test.go internal/algorithms/fourier/fourier.go internal/algorithms/fourier/fourier_parallel.go 
	go test pkg/analyzer/analyzer_test.go pkg/analyzer/analyzer.go 

testCover:
	go test internal/algorithms/fourier/fourier_test.go internal/algorithms/fourier/fourier.go internal/algorithms/fourier/fourier_parallel.go -cover -coverprofile=test/fourier_cover.out
	go tool cover -html=test/fourier_cover.out -o test/fourier_cover.html
	go test pkg/analyzer/analyzer_test.go pkg/analyzer/analyzer.go -cover -coverprofile=test/analyzer_cover.out
	go tool cover -html=test/analyzer_cover.out -o test/analyzer_cover.html