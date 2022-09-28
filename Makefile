build:
	go build -o bin/morsencoder ./cmd/morsencoder

run:
	./bin/morsencoder

start:
	go build -o bin/morsencoder ./cmd/morsencoder && ./bin/morsencoder

dockerize:
	docker build -t eneskzlcn/morsencoder:latest .

clean:
	rm -rf bin

unit-tests:
	go test -v ./internal/morsencoder

generate-mocks:
	mockgen -destination=internal/mocks/mock_logger.go -package mocks github.com/eneskzlcn/morsencoder/internal/morsencoder Logger
	mockgen -destination=internal/mocks/mock_morsencoder_service.go -package mocks github.com/eneskzlcn/morsencoder/internal/morsencoder MorseEncoderService