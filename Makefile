build:
	go build -o bin/morsencoder ./cmd/morsencoder
run:
	./bin/morsencoder
start:
	go build -o bin/morsencoder ./cmd/morsencoder && ./bin/morsencoder
clean:
	rm -rf bin