build:
	go build -o bin/client_test cmd/main.go

run:
	./bin/client_test

clean:
	rm -rf bin/
