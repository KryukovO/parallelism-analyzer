run:
	go run cmd/main.go

tests:
	rm -rf results && mkdir results
	rm -rf test && mkdir test
	go test parallelism-analyzer/internal/algorithms/fourier/
	go test parallelism-analyzer/internal/algorithms/dithering/
	go test parallelism-analyzer/pkg/analyzer/

testCover:
	rm -rf results && mkdir results
	rm -rf test && mkdir test
	go test parallelism-analyzer/internal/algorithms/fourier/ -cover -coverprofile=test/fourier_cover.out
	go tool cover -html=test/fourier_cover.out -o test/fourier_cover.html
	go test parallelism-analyzer/internal/algorithms/dithering/ -cover -coverprofile=test/dithering_cover.out
	go tool cover -html=test/dithering_cover.out -o test/dithering_cover.html
	go test parallelism-analyzer/pkg/analyzer/ -cover -coverprofile=test/analyzer_cover.out
	go tool cover -html=test/analyzer_cover.out -o test/analyzer_cover.html