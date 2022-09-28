build:
	go build -o bin/morsencoder ./cmd/morsencoder
run:
	./bin/morsencoder
start:
	go build -o bin/morsencoder ./cmd/morsencoder && ./bin/morsencoder
clean:
	rm -rf bin

generate-mocks:
	mockgen -destination=internal/mocks/mock_logger.go -package mocks github.com/eneskzlcn/morsencoder/internal/morsencoder Logger
	mockgen -destination=internal/mocks/mock_morsencoder_service.go -package mocks github.com/eneskzlcn/morsencoder/internal/morsencoder MorseEncoderService